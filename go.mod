module github.com/smartcontractkit/chainlink/v2

go 1.20

require (
	github.com/CosmWasm/wasmd v0.30.0
	github.com/Depado/ginprom v1.7.11
	github.com/Masterminds/semver/v3 v3.2.1
	github.com/Masterminds/sprig/v3 v3.2.3
	github.com/avast/retry-go/v4 v4.3.4
	github.com/btcsuite/btcd v0.23.4
	github.com/cosmos/cosmos-sdk v0.45.11
	github.com/danielkov/gin-helmet v0.0.0-20171108135313-1387e224435e
	github.com/ethereum/go-ethereum v1.12.0
	github.com/fatih/color v1.15.0
	github.com/fxamacker/cbor/v2 v2.4.0
	github.com/gagliardetto/solana-go v1.4.1-0.20220428092759-5250b4abbb27
	github.com/getsentry/sentry-go v0.19.0
	github.com/gin-contrib/cors v1.4.0
	github.com/gin-contrib/expvar v0.0.1
	github.com/gin-contrib/sessions v0.0.5
	github.com/gin-contrib/size v0.0.0-20230212012657-e14a14094dc4
	github.com/gin-gonic/gin v1.9.1
	github.com/go-webauthn/webauthn v0.8.2
	github.com/gogo/protobuf v1.3.3
	github.com/google/pprof v0.0.0-20230602150820-91b7bce49751
	github.com/google/uuid v1.3.0
	github.com/gorilla/securecookie v1.1.1
	github.com/gorilla/sessions v1.2.1
	github.com/gorilla/websocket v1.5.0
	github.com/graph-gophers/dataloader v5.0.0+incompatible
	github.com/graph-gophers/graphql-go v1.3.0
	github.com/hashicorp/go-plugin v1.4.10
	github.com/hdevalence/ed25519consensus v0.0.0-20220222234857-c00d1f31bab3
	github.com/jackc/pgconn v1.14.0
	github.com/jackc/pgtype v1.14.0
	github.com/jackc/pgx/v4 v4.18.1
	github.com/jpillora/backoff v1.0.0
	github.com/kylelemons/godebug v1.1.0
	github.com/leanovate/gopter v0.2.10-0.20210127095200-9abe2343507a
	github.com/lib/pq v1.10.9
	github.com/libp2p/go-libp2p-core v0.8.5
	github.com/libp2p/go-libp2p-peerstore v0.2.7
	github.com/manyminds/api2go v0.0.0-20171030193247-e7b693844a6f
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/mapstructure v1.5.0
	github.com/mr-tron/base58 v1.2.0
	github.com/multiformats/go-multiaddr v0.3.3
	github.com/olekukonko/tablewriter v0.0.5
	github.com/onsi/gomega v1.27.8
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pelletier/go-toml v1.9.5
	github.com/pelletier/go-toml/v2 v2.0.8
	github.com/pkg/errors v0.9.1
	github.com/pressly/goose/v3 v3.5.3
	github.com/prometheus/client_golang v1.16.0
	github.com/prometheus/client_model v0.4.0
	github.com/prometheus/common v0.44.0
	github.com/prometheus/prometheus v0.45.0
	github.com/pyroscope-io/client v0.7.1
	github.com/robfig/cron/v3 v3.0.1
	github.com/rogpeppe/go-internal v1.10.0
	github.com/rs/zerolog v1.29.1
	github.com/scylladb/go-reflectx v1.0.1
	github.com/shirou/gopsutil/v3 v3.22.12
	github.com/shopspring/decimal v1.3.1
	github.com/smartcontractkit/caigo v0.0.0-20230621050857-b29a4ca8c704
	github.com/smartcontractkit/chainlink-cosmos v0.4.1-0.20230525203711-20bed74ac906
	github.com/smartcontractkit/chainlink-relay v0.1.7-0.20230620171700-bbcb3a99b7d3
	github.com/smartcontractkit/chainlink-solana v1.0.3-0.20230612131011-369bfb503592
	github.com/smartcontractkit/chainlink-starknet/relayer v0.0.1-beta-test.0.20230622060316-7ce48476dd7d
	github.com/smartcontractkit/chainlink/integration-tests v0.0.0-20230601092347-ed7c04f0085f
	github.com/smartcontractkit/libocr v0.0.0-20230606215712-82b910bef5c1
	github.com/smartcontractkit/ocr2keepers v0.7.0
	github.com/smartcontractkit/ocr2vrf v0.0.0-20230616201444-d8b4222aff3c
	github.com/smartcontractkit/sqlx v1.3.5-0.20210805004948-4be295aacbeb
	github.com/smartcontractkit/tdh2/go/ocr2/decryptionplugin v0.0.0-20230616141813-ca0ecf03ca5c
	github.com/smartcontractkit/tdh2/go/tdh2 v0.0.0-20230625174811-4934857d1c4a
	github.com/smartcontractkit/wsrpc v0.7.2
	github.com/spf13/cast v1.5.1
	github.com/stretchr/testify v1.8.4
	github.com/tendermint/tendermint v0.34.23
	github.com/test-go/testify v1.1.4
	github.com/theodesp/go-heaps v0.0.0-20190520121037-88e35354fe0a
	github.com/tidwall/gjson v1.14.4
	github.com/ugorji/go/codec v1.2.11
	github.com/ulule/limiter/v3 v3.11.2
	github.com/umbracle/ethgo v0.1.3
	github.com/unrolled/secure v1.13.0
	github.com/urfave/cli v1.22.13
	go.dedis.ch/fixbuf v1.0.3
	go.dedis.ch/kyber/v3 v3.1.0
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.24.0
	golang.org/x/crypto v0.11.0
	golang.org/x/exp v0.0.0-20230626212559-97b1e661b5df
	golang.org/x/net v0.12.0
	golang.org/x/sync v0.3.0
	golang.org/x/term v0.10.0
	golang.org/x/text v0.11.0
	golang.org/x/time v0.3.0
	golang.org/x/tools v0.11.0
	gonum.org/v1/gonum v0.13.0
	google.golang.org/protobuf v1.31.0
	gopkg.in/guregu/null.v2 v2.1.2
	gopkg.in/guregu/null.v4 v4.0.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	k8s.io/utils v0.0.0-20230308161112-d77c459e9343
)

require (
	contrib.go.opencensus.io/exporter/stackdriver v0.13.5 // indirect
	filippo.io/edwards25519 v1.0.0-rc.1 // indirect
	github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4 // indirect
	github.com/99designs/keyring v1.2.1 // indirect
	github.com/Azure/go-ansiterm v0.0.0-20230124172434-306776ec8161 // indirect
	github.com/ChainSafe/go-schnorrkel v0.0.0-20200405005733-88cbf1b4c40d // indirect
	github.com/CosmWasm/wasmvm v1.1.1 // indirect
	github.com/DataDog/zstd v1.5.2 // indirect
	github.com/MakeNowJust/heredoc v1.0.0 // indirect
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/VictoriaMetrics/fastcache v1.10.0 // indirect
	github.com/alecthomas/units v0.0.0-20211218093645-b94a6e3cc137 // indirect
	github.com/andres-erbsen/clock v0.0.0-20160526145045-9e14626cd129 // indirect
	github.com/armon/go-metrics v0.4.1 // indirect
	github.com/aws/constructs-go/constructs/v10 v10.1.255 // indirect
	github.com/aws/jsii-runtime-go v1.75.0 // indirect
	github.com/aybabtme/rgbterm v0.0.0-20170906152045-cc83f3b3ce59 // indirect
	github.com/benbjohnson/clock v1.3.4 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bgentry/speakeasy v0.1.0 // indirect
	github.com/blendle/zapdriver v1.3.1 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.3.2 // indirect
	github.com/btcsuite/btcd/chaincfg/chainhash v1.0.1 // indirect
	github.com/bytedance/sonic v1.9.1 // indirect
	github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2 v2.7.5 // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/cespare/xxhash v1.1.0 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/chai2010/gettext-go v1.0.2 // indirect
	github.com/chaos-mesh/chaos-mesh/api/v1alpha1 v0.0.0-20220226050744-799408773657 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/cockroachdb/errors v1.9.1 // indirect
	github.com/cockroachdb/logtags v0.0.0-20230118201751-21c54148d20b // indirect
	github.com/cockroachdb/pebble v0.0.0-20230209160836-829675f94811 // indirect
	github.com/cockroachdb/redact v1.1.3 // indirect
	github.com/confio/ics23/go v0.7.0 // indirect
	github.com/cosmos/btcutil v1.0.4 // indirect
	github.com/cosmos/cosmos-proto v1.0.0-alpha8 // indirect
	github.com/cosmos/go-bip39 v1.0.0 // indirect
	github.com/cosmos/gogoproto v1.4.3 // indirect
	github.com/cosmos/gorocksdb v1.2.0 // indirect
	github.com/cosmos/iavl v0.19.4 // indirect
	github.com/cosmos/ibc-go/v4 v4.2.0 // indirect
	github.com/cosmos/ledger-cosmos-go v0.11.1 // indirect
	github.com/cosmos/ledger-go v0.9.2 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/danieljoos/wincred v1.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/davidlazar/go-crypto v0.0.0-20200604182044-b73af7476f6c // indirect
	github.com/deckarep/golang-set/v2 v2.3.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.2.0 // indirect
	github.com/dfuse-io/logging v0.0.0-20210109005628-b97a57253f70 // indirect
	github.com/dgraph-io/badger/v2 v2.2007.4 // indirect
	github.com/dgraph-io/ristretto v0.1.0 // indirect
	github.com/dgryski/go-farm v0.0.0-20200201041132-a6ae2369ad13 // indirect
	github.com/docker/distribution v2.8.1+incompatible // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/dvsekhvalnov/jose2go v1.5.0 // indirect
	github.com/emicklei/go-restful/v3 v3.10.1 // indirect
	github.com/evanphx/json-patch v5.6.0+incompatible // indirect
	github.com/evanphx/json-patch/v5 v5.6.0 // indirect
	github.com/exponent-io/jsonpath v0.0.0-20210407135951-1de76d718b3f // indirect
	github.com/fatih/camelcase v1.0.0 // indirect
	github.com/flynn/noise v0.0.0-20180327030543-2492fe189ae6 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/fvbommel/sortorder v1.0.2 // indirect
	github.com/gabriel-vasile/mimetype v1.4.2 // indirect
	github.com/gagliardetto/binary v0.7.1 // indirect
	github.com/gagliardetto/treeout v0.1.4 // indirect
	github.com/gballet/go-libpcsclite v0.0.0-20191108122812-4678299bea08 // indirect
	github.com/gedex/inflector v0.0.0-20170307190818-16278e9db813 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-errors/errors v1.4.2 // indirect
	github.com/go-kit/kit v0.12.0 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-openapi/jsonpointer v0.19.6 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.14.0 // indirect
	github.com/go-resty/resty/v2 v2.7.0 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/go-webauthn/revoke v0.1.9 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/godbus/dbus v0.0.0-20190726142602-4481cbc300e2 // indirect
	github.com/gofrs/flock v0.8.1 // indirect
	github.com/gofrs/uuid v4.3.1+incompatible // indirect
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
	github.com/golang/glog v1.1.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/golang/snappy v0.0.5-0.20220116011046-fa5810519dcb // indirect
	github.com/google/btree v1.1.2 // indirect
	github.com/google/gnostic v0.6.9 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/go-tpm v0.3.3 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/gopacket v1.1.19 // indirect
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
	github.com/gorilla/context v1.1.1 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus v1.0.0-rc.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.0-rc.3 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/gsterjov/go-libsecret v0.0.0-20161001094733-a6f4afe4910c // indirect
	github.com/gtank/merlin v0.1.1 // indirect
	github.com/gtank/ristretto255 v0.1.2 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-hclog v1.4.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/golang-lru v0.6.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/yamux v0.0.0-20200609203250-aecfd211c9ce // indirect
	github.com/holiman/bloomfilter/v2 v2.0.3 // indirect
	github.com/holiman/uint256 v1.2.2 // indirect
	github.com/huandu/xstrings v1.4.0 // indirect
	github.com/huin/goupnp v1.0.3 // indirect
	github.com/imdario/mergo v0.3.15 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/ipfs/go-cid v0.0.7 // indirect
	github.com/ipfs/go-datastore v0.4.5 // indirect
	github.com/ipfs/go-ipfs-util v0.0.2 // indirect
	github.com/ipfs/go-ipns v0.0.2 // indirect
	github.com/ipfs/go-log v1.0.4 // indirect
	github.com/ipfs/go-log/v2 v2.1.1 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.2 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackpal/go-nat-pmp v1.0.2 // indirect
	github.com/jbenet/go-temp-err-catcher v0.1.0 // indirect
	github.com/jbenet/goprocess v0.1.4 // indirect
	github.com/jmhodges/levigo v1.0.0 // indirect
	github.com/jmoiron/sqlx v1.3.5 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kelseyhightower/envconfig v1.4.0 // indirect
	github.com/klauspost/compress v1.16.3 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/koron/go-ssdp v0.0.2 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/libp2p/go-addr-util v0.0.2 // indirect
	github.com/libp2p/go-buffer-pool v0.1.0 // indirect
	github.com/libp2p/go-cidranger v1.1.0 // indirect
	github.com/libp2p/go-conn-security-multistream v0.2.0 // indirect
	github.com/libp2p/go-eventbus v0.2.1 // indirect
	github.com/libp2p/go-flow-metrics v0.0.3 // indirect
	github.com/libp2p/go-libp2p v0.13.0 // indirect
	github.com/libp2p/go-libp2p-asn-util v0.0.0-20201026210036-4f868c957324 // indirect
	github.com/libp2p/go-libp2p-autonat v0.4.0 // indirect
	github.com/libp2p/go-libp2p-blankhost v0.2.0 // indirect
	github.com/libp2p/go-libp2p-circuit v0.4.0 // indirect
	github.com/libp2p/go-libp2p-discovery v0.5.0 // indirect
	github.com/libp2p/go-libp2p-kad-dht v0.11.1 // indirect
	github.com/libp2p/go-libp2p-kbucket v0.4.7 // indirect
	github.com/libp2p/go-libp2p-loggables v0.1.0 // indirect
	github.com/libp2p/go-libp2p-mplex v0.4.1 // indirect
	github.com/libp2p/go-libp2p-nat v0.0.6 // indirect
	github.com/libp2p/go-libp2p-noise v0.1.2 // indirect
	github.com/libp2p/go-libp2p-pnet v0.2.0 // indirect
	github.com/libp2p/go-libp2p-record v0.1.3 // indirect
	github.com/libp2p/go-libp2p-swarm v0.4.0 // indirect
	github.com/libp2p/go-libp2p-tls v0.1.3 // indirect
	github.com/libp2p/go-libp2p-transport-upgrader v0.4.0 // indirect
	github.com/libp2p/go-libp2p-yamux v0.5.1 // indirect
	github.com/libp2p/go-mplex v0.3.0 // indirect
	github.com/libp2p/go-msgio v0.0.6 // indirect
	github.com/libp2p/go-nat v0.0.5 // indirect
	github.com/libp2p/go-netroute v0.1.4 // indirect
	github.com/libp2p/go-openssl v0.0.7 // indirect
	github.com/libp2p/go-reuseport v0.0.2 // indirect
	github.com/libp2p/go-reuseport-transport v0.0.4 // indirect
	github.com/libp2p/go-sockaddr v0.1.0 // indirect
	github.com/libp2p/go-stream-muxer-multistream v0.3.0 // indirect
	github.com/libp2p/go-tcp-transport v0.2.1 // indirect
	github.com/libp2p/go-ws-transport v0.4.0 // indirect
	github.com/libp2p/go-yamux/v2 v2.0.0 // indirect
	github.com/liggitt/tabwriter v0.0.0-20181228230101-89fcab3d43de // indirect
	github.com/logrusorgru/aurora v2.0.3+incompatible // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/mimoo/StrobeGo v0.0.0-20210601165009-122bf33a46e0 // indirect
	github.com/minio/blake2b-simd v0.0.0-20160723061019-3f5f724cb5b1 // indirect
	github.com/minio/sha256-simd v0.1.1 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/moby/spdystream v0.2.0 // indirect
	github.com/moby/term v0.0.0-20220808134915-39b0c02b01ae // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/monochromegane/go-gitignore v0.0.0-20200626010858-205db1a8cc00 // indirect
	github.com/mostynb/zstdpool-freelist v0.0.0-20201229113212-927304c0c3b1 // indirect
	github.com/mtibben/percent v0.2.1 // indirect
	github.com/multiformats/go-base32 v0.0.3 // indirect
	github.com/multiformats/go-base36 v0.1.0 // indirect
	github.com/multiformats/go-multiaddr-dns v0.2.0 // indirect
	github.com/multiformats/go-multiaddr-fmt v0.1.0 // indirect
	github.com/multiformats/go-multiaddr-net v0.2.0 // indirect
	github.com/multiformats/go-multibase v0.0.3 // indirect
	github.com/multiformats/go-multihash v0.0.14 // indirect
	github.com/multiformats/go-multistream v0.2.2 // indirect
	github.com/multiformats/go-varint v0.0.6 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.1.0-rc3 // indirect
	github.com/opencontainers/runc v1.1.7 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/petermattis/goid v0.0.0-20180202154549-b0b1615b78e5 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c // indirect
	github.com/prometheus/procfs v0.11.0 // indirect
	github.com/pyroscope-io/godeltaprof v0.1.0 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/regen-network/cosmos-proto v0.3.1 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/russross/blackfriday v1.6.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/sasha-s/go-deadlock v0.3.1 // indirect
	github.com/shirou/gopsutil v3.21.11+incompatible // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/smartcontractkit/chainlink-env v0.32.5 // indirect
	github.com/smartcontractkit/chainlink-testing-framework v1.11.6 // indirect
	github.com/spacemonkeygo/spacelog v0.0.0-20180420211403-2296661a0572 // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/spf13/afero v1.9.2 // indirect
	github.com/spf13/cobra v1.6.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.14.0 // indirect
	github.com/status-im/keycard-go v0.2.0 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	github.com/subosito/gotenv v1.4.1 // indirect
	github.com/syndtr/goleveldb v1.0.1-0.20220614013038-64ee5596c38a // indirect
	github.com/tendermint/btcd v0.1.1 // indirect
	github.com/tendermint/crypto v0.0.0-20191022145703-50d29ede1e15 // indirect
	github.com/tendermint/go-amino v0.16.0 // indirect
	github.com/tendermint/tm-db v0.6.7 // indirect
	github.com/teris-io/shortid v0.0.0-20201117134242-e59966efd125 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/tklauser/go-sysconf v0.3.11 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/tyler-smith/go-bip39 v1.1.0 // indirect
	github.com/umbracle/fastrlp v0.0.0-20220527094140-59d5dd30e722 // indirect
	github.com/valyala/fastjson v1.4.1 // indirect
	github.com/whyrusleeping/go-keyspace v0.0.0-20160322163242-5b898ac5add1 // indirect
	github.com/whyrusleeping/multiaddr-filter v0.0.0-20160516205228-e903e4adabd7 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	github.com/xlab/treeprint v1.1.0 // indirect
	github.com/yuin/goldmark v1.4.13 // indirect
	github.com/yusufpapurcu/wmi v1.2.3 // indirect
	github.com/zondax/hid v0.9.0 // indirect
	go.dedis.ch/protobuf v1.0.11 // indirect
	go.etcd.io/bbolt v1.3.6 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.42.0 // indirect
	go.opentelemetry.io/otel v1.16.0 // indirect
	go.opentelemetry.io/otel/metric v1.16.0 // indirect
	go.opentelemetry.io/otel/sdk v1.16.0 // indirect
	go.opentelemetry.io/otel/trace v1.16.0 // indirect
	go.starlark.net v0.0.0-20220817180228-f738f5508c12 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/ratelimit v0.2.0 // indirect
	golang.org/x/arch v0.3.0 // indirect
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	golang.org/x/mod v0.12.0 // indirect
	golang.org/x/oauth2 v0.8.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	gomodules.xyz/jsonpatch/v2 v2.2.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20230320184635-7606e756e683 // indirect
	google.golang.org/grpc v1.55.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/api v0.26.2 // indirect
	k8s.io/apiextensions-apiserver v0.25.3 // indirect
	k8s.io/apimachinery v0.26.2 // indirect
	k8s.io/cli-runtime v0.25.4 // indirect
	k8s.io/client-go v0.26.2 // indirect
	k8s.io/component-base v0.25.4 // indirect
	k8s.io/klog/v2 v2.100.1 // indirect
	k8s.io/kube-openapi v0.0.0-20230501164219-8b0f38b5fd1f // indirect
	k8s.io/kubectl v0.25.4 // indirect
	sigs.k8s.io/controller-runtime v0.13.0 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/kustomize/api v0.12.1 // indirect
	sigs.k8s.io/kustomize/kyaml v0.13.9 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)

replace (
	// Fix go mod tidy issue for ambiguous imports from go-ethereum
	// See https://github.com/ugorji/go/issues/279
	github.com/btcsuite/btcd => github.com/btcsuite/btcd v0.22.1

	// replicating the replace directive on cosmos SDK
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

	// until merged upstream: https://github.com/hashicorp/go-plugin/pull/257
	github.com/hashicorp/go-plugin => github.com/jmank88/go-plugin v0.0.0-20230604120638-7bb12ec27e75

	// K8s imports are weird
	k8s.io/api => k8s.io/api v0.25.4
	k8s.io/client-go => k8s.io/client-go v0.25.4
)
