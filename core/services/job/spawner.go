package job

import (
	"context"
	"fmt"
	"math"
	"reflect"
	"sync"

	"github.com/pkg/errors"
	"github.com/smartcontractkit/sqlx"

	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services"
	"github.com/smartcontractkit/chainlink/core/services/pg"
	"github.com/smartcontractkit/chainlink/core/utils"
)

//go:generate mockery --name Spawner --output ./mocks/ --case=underscore
//go:generate mockery --name Delegate --output ./mocks/ --case=underscore

type (
	// Spawner manages the spinning up and down of the long-running
	// services that perform the work described by job specs.  Each active job spec
	// has 1 or more of these services associated with it.
	Spawner interface {
		services.ServiceCtx

		// CreateJob creates a new job and starts services.
		// All services must start without errors for the job to be active.
		CreateJob(jb *Job, qopts ...pg.QOpt) (err error)
		// DeleteJob deletes a job and stops any active services.
		DeleteJob(jobID int32, qopts ...pg.QOpt) error
		// ActiveJobs returns a map of jobs with active services (started without error).
		ActiveJobs() map[int32]Job

		// StartService starts services for the given job spec.
		// NOTE: Prefer to use CreateJob, this is only publicly exposed for use in tests
		// to start a job that was previously manually inserted into DB
		StartService(ctx context.Context, spec Job) error
	}

	spawner struct {
		orm              ORM
		config           Config
		jobTypeDelegates map[Type]Delegate
		activeJobs       map[int32]activeJob
		activeJobsMu     sync.RWMutex
		q                pg.Q
		lggr             logger.Logger

		utils.StartStopOnce
		chStop              chan struct{}
		lbDependentAwaiters []utils.DependentAwaiter
	}

	// TODO(spook): I can't wait for Go generics
	Delegate interface {
		JobType() Type
		// ServicesForSpec returns services to be started and stopped for this
		// job. In case a given job type relies upon well-defined startup/shutdown
		// ordering for services, they are started in the order they are given
		// and stopped in reverse order.
		BeforeJobCreated(spec Job)
		ServicesForSpec(spec Job) ([]ServiceCtx, error)
		AfterJobCreated(spec Job)
		BeforeJobDeleted(spec Job)
	}

	activeJob struct {
		delegate Delegate
		spec     Job
		services []ServiceCtx
	}
)

var _ Spawner = (*spawner)(nil)

func NewSpawner(orm ORM, config Config, jobTypeDelegates map[Type]Delegate, db *sqlx.DB, lggr logger.Logger, lbDependentAwaiters []utils.DependentAwaiter) *spawner {
	namedLogger := lggr.Named("JobSpawner")
	s := &spawner{
		orm:                 orm,
		config:              config,
		jobTypeDelegates:    jobTypeDelegates,
		q:                   pg.NewQ(db, namedLogger, config),
		lggr:                namedLogger,
		activeJobs:          make(map[int32]activeJob),
		chStop:              make(chan struct{}),
		lbDependentAwaiters: lbDependentAwaiters,
	}
	return s
}

// Start starts Spawner.
func (js *spawner) Start(ctx context.Context) error {
	return js.StartOnce("JobSpawner", func() error {
		js.startAllServices(ctx)
		return nil

	})
}

func (js *spawner) Close() error {
	return js.StopOnce("JobSpawner", func() error {
		close(js.chStop)
		js.stopAllServices()
		return nil

	})
}

func (js *spawner) startAllServices(ctx context.Context) {
	// TODO: rename to find AllJobs
	specs, _, err := js.orm.FindJobs(0, math.MaxUint32)
	if err != nil {
		js.lggr.Criticalf("Couldn't fetch unclaimed jobs: %v", err)
		return
	}

	for _, spec := range specs {
		if err = js.StartService(ctx, spec); err != nil {
			js.lggr.Errorf("Couldn't start service %v: %v", spec.Name, err)
		}
	}
	// Log Broadcaster fully starts after all initial Register calls are done from other starting services
	// to make sure the initial backfill covers those subscribers.
	for _, lbd := range js.lbDependentAwaiters {
		lbd.DependentReady()
	}
}

func (js *spawner) stopAllServices() {
	jobIDs := js.activeJobIDs()
	for _, jobID := range jobIDs {
		js.stopService(jobID)
	}
}

// stopService removes the job from memory and stop the services.
// It will always delete the job from memory even if closing the services fail.
func (js *spawner) stopService(jobID int32) {
	js.lggr.Debugw("Stopping services for job", "jobID", jobID)
	js.activeJobsMu.Lock()
	defer js.activeJobsMu.Unlock()

	aj := js.activeJobs[jobID]

	for i := len(aj.services) - 1; i >= 0; i-- {
		service := aj.services[i]
		err := service.Close()
		if err != nil {
			js.lggr.Criticalw("Error stopping job service", "jobID", jobID, "error", err, "subservice", i, "serviceType", reflect.TypeOf(service))
		} else {
			js.lggr.Debugw("Stopped job service", "jobID", jobID, "subservice", i, "serviceType", fmt.Sprintf("%T", service))
		}
	}
	js.lggr.Debugw("Stopped all services for job", "jobID", jobID)

	delete(js.activeJobs, jobID)
}

func (js *spawner) StartService(ctx context.Context, jb Job) error {
	js.activeJobsMu.Lock()
	defer js.activeJobsMu.Unlock()

	delegate, exists := js.jobTypeDelegates[jb.Type]
	if !exists {
		js.lggr.Errorw("Job type has not been registered with job.Spawner", "type", jb.Type, "jobID", jb.ID)
		return errors.Errorf("unregistered type %q for job: %d", jb.Type, jb.ID)
	}
	// We always add the active job in the activeJob map, even in the case
	// that it fails to start. That way we have access to the delegate to call
	// OnJobDeleted before deleting. However, the activeJob will only have services
	// that it was able to start without an error.
	aj := activeJob{delegate: delegate, spec: jb}

	jb.PipelineSpec.JobName = jb.Name.ValueOrZero()
	jb.PipelineSpec.JobID = jb.ID
	jb.PipelineSpec.JobType = string(jb.Type)
	jb.PipelineSpec.ForwardingAllowed = jb.ForwardingAllowed
	if jb.GasLimit.Valid {
		jb.PipelineSpec.GasLimit = &jb.GasLimit.Uint32
	}

	srvs, err := delegate.ServicesForSpec(jb)
	if err != nil {
		js.lggr.Errorw("Error creating services for job", "jobID", jb.ID, "error", err)
		cctx, cancel := utils.ContextFromChan(js.chStop)
		defer cancel()
		js.orm.TryRecordError(jb.ID, err.Error(), pg.WithParentCtx(cctx))
		js.activeJobs[jb.ID] = aj
		return errors.Wrapf(err, "failed to create services for job: %d", jb.ID)
	}

	js.lggr.Debugw("JobSpawner: Starting services for job", "jobID", jb.ID, "count", len(srvs))

	var ms services.MultiStart
	for _, srv := range srvs {
		err = ms.Start(ctx, srv)
		if err != nil {
			js.lggr.Criticalw("Error starting service for job", "jobID", jb.ID, "error", err)
			return errors.Wrapf(err, "failed to start service for job %d", jb.ID)
		}
		aj.services = append(aj.services, srv)
	}
	js.lggr.Debugw("JobSpawner: Finished starting services for job", "jobID", jb.ID, "count", len(srvs))
	js.activeJobs[jb.ID] = aj
	return nil
}

// Should not get called before Start()
func (js *spawner) CreateJob(jb *Job, qopts ...pg.QOpt) (err error) {
	delegate, exists := js.jobTypeDelegates[jb.Type]
	if !exists {
		js.lggr.Errorf("job type '%s' has not been registered with the job.Spawner", jb.Type)
		err = errors.Errorf("job type '%s' has not been registered with the job.Spawner", jb.Type)
		return
	}

	q := js.q.WithOpts(qopts...)
	if q.ParentCtx != nil {
		ctx, cancel := utils.WithCloseChan(q.ParentCtx, js.chStop)
		defer cancel()
		q.ParentCtx = ctx
	} else {
		ctx, cancel := utils.ContextFromChan(js.chStop)
		defer cancel()
		q.ParentCtx = ctx
	}
	ctx, cancel := q.Context()
	defer cancel()

	err = js.orm.CreateJob(jb, pg.WithQueryer(q.Queryer), pg.WithParentCtx(ctx))
	if err != nil {
		js.lggr.Errorw("Error creating job", "type", jb.Type, "error", err)
		return
	}
	js.lggr.Infow("Created job", "type", jb.Type, "jobID", jb.ID)

	delegate.BeforeJobCreated(*jb)
	err = js.StartService(q.ParentCtx, *jb)
	if err != nil {
		js.lggr.Errorw("Error starting job services", "type", jb.Type, "jobID", jb.ID, "error", err)
	} else {
		js.lggr.Infow("Started job services", "type", jb.Type, "jobID", jb.ID)
	}

	delegate.AfterJobCreated(*jb)

	return err
}

// Should not get called before Start()
func (js *spawner) DeleteJob(jobID int32, qopts ...pg.QOpt) error {
	if jobID == 0 {
		return errors.New("will not delete job with 0 ID")
	}

	lggr := js.lggr.With("jobID", jobID)
	lggr.Debugw("Deleting job")

	var aj activeJob
	var exists bool
	func() {
		js.activeJobsMu.RLock()
		defer js.activeJobsMu.RUnlock()
		aj, exists = js.activeJobs[jobID]
	}()

	ctx, cancel := utils.ContextFromChan(js.chStop)
	defer cancel()

	if !exists { // inactive, so look up the spec and delegate
		jb, err := js.orm.FindJob(ctx, jobID)
		if err != nil {
			return errors.Wrapf(err, "job %d not found", jobID)
		}
		aj.spec = jb
		if !func() (ok bool) {
			js.activeJobsMu.RLock()
			defer js.activeJobsMu.RUnlock()
			aj.delegate, ok = js.jobTypeDelegates[jb.Type]
			return ok
		}() {
			js.lggr.Errorw("Job type has not been registered with job.Spawner", "type", jb.Type, "jobID", jb.ID)
			return errors.Errorf("unregistered type %q for job: %d", jb.Type, jb.ID)
		}
	}
	lggr.Debugw("Callback: BeforeJobDeleted")
	aj.delegate.BeforeJobDeleted(aj.spec)
	lggr.Debugw("Callback: BeforeJobDeleted done")

	err := js.orm.DeleteJob(jobID, append(qopts, pg.WithParentCtx(ctx))...)
	if err != nil {
		js.lggr.Errorw("Error deleting job", "jobID", jobID, "error", err)
		return err
	}

	if exists {
		// Stop the service and remove the job from memory, which will always happen even if closing the services fail.
		js.stopService(jobID)
	}

	lggr.Infow("Stopped and deleted job")

	return nil
}

func (js *spawner) ActiveJobs() map[int32]Job {
	js.activeJobsMu.RLock()
	defer js.activeJobsMu.RUnlock()

	m := make(map[int32]Job, len(js.activeJobs))
	for jobID := range js.activeJobs {
		m[jobID] = js.activeJobs[jobID].spec
	}
	return m
}

func (js *spawner) activeJobIDs() []int32 {
	js.activeJobsMu.RLock()
	defer js.activeJobsMu.RUnlock()

	ids := make([]int32, 0, len(js.activeJobs))
	for jobID := range js.activeJobs {
		ids = append(ids, jobID)
	}
	return ids
}

var _ Delegate = &NullDelegate{}

type NullDelegate struct {
	Type Type
}

func (n *NullDelegate) JobType() Type {
	return n.Type
}

// ServicesForSpec does no-op.
func (n *NullDelegate) ServicesForSpec(spec Job) (s []ServiceCtx, err error) {
	return
}

func (*NullDelegate) BeforeJobCreated(spec Job) {}
func (*NullDelegate) AfterJobCreated(spec Job)  {}
func (*NullDelegate) BeforeJobDeleted(spec Job) {}
