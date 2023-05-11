// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: telem_enhanced_ea_mercury.proto

package telem

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EnhancedEAMercury struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataSource                    string  `protobuf:"bytes,1,opt,name=data_source,json=dataSource,proto3" json:"data_source,omitempty"`
	DpBenchmarkPrice              float64 `protobuf:"fixed64,2,opt,name=dp_benchmark_price,json=dpBenchmarkPrice,proto3" json:"dp_benchmark_price,omitempty"`
	DpBid                         float64 `protobuf:"fixed64,3,opt,name=dp_bid,json=dpBid,proto3" json:"dp_bid,omitempty"`
	DpAsk                         float64 `protobuf:"fixed64,4,opt,name=dp_ask,json=dpAsk,proto3" json:"dp_ask,omitempty"`
	CurrentBlockNumber            int64   `protobuf:"varint,5,opt,name=current_block_number,json=currentBlockNumber,proto3" json:"current_block_number,omitempty"`
	CurrentBlockHash              string  `protobuf:"bytes,6,opt,name=current_block_hash,json=currentBlockHash,proto3" json:"current_block_hash,omitempty"`
	CurrentBlockTimestamp         uint64  `protobuf:"varint,7,opt,name=current_block_timestamp,json=currentBlockTimestamp,proto3" json:"current_block_timestamp,omitempty"`
	BridgeTaskRunStartedTimestamp int64   `protobuf:"varint,8,opt,name=bridge_task_run_started_timestamp,json=bridgeTaskRunStartedTimestamp,proto3" json:"bridge_task_run_started_timestamp,omitempty"`
	BridgeTaskRunEndedTimestamp   int64   `protobuf:"varint,9,opt,name=bridge_task_run_ended_timestamp,json=bridgeTaskRunEndedTimestamp,proto3" json:"bridge_task_run_ended_timestamp,omitempty"`
	ProviderRequestedTimestamp    int64   `protobuf:"varint,10,opt,name=provider_requested_timestamp,json=providerRequestedTimestamp,proto3" json:"provider_requested_timestamp,omitempty"`
	ProviderReceivedTimestamp     int64   `protobuf:"varint,11,opt,name=provider_received_timestamp,json=providerReceivedTimestamp,proto3" json:"provider_received_timestamp,omitempty"`
	ProviderDataStreamEstablished int64   `protobuf:"varint,12,opt,name=provider_data_stream_established,json=providerDataStreamEstablished,proto3" json:"provider_data_stream_established,omitempty"`
	ProviderIndicatedTime         int64   `protobuf:"varint,13,opt,name=provider_indicated_time,json=providerIndicatedTime,proto3" json:"provider_indicated_time,omitempty"`
	Feed                          string  `protobuf:"bytes,14,opt,name=feed,proto3" json:"feed,omitempty"`
	ObservationBenchmarkPrice     int64   `protobuf:"varint,15,opt,name=observation_benchmark_price,json=observationBenchmarkPrice,proto3" json:"observation_benchmark_price,omitempty"`
	ObservationBid                int64   `protobuf:"varint,16,opt,name=observation_bid,json=observationBid,proto3" json:"observation_bid,omitempty"`
	ObservationAsk                int64   `protobuf:"varint,17,opt,name=observation_ask,json=observationAsk,proto3" json:"observation_ask,omitempty"`
	ConfigDigest                  string  `protobuf:"bytes,18,opt,name=config_digest,json=configDigest,proto3" json:"config_digest,omitempty"`
	Round                         int64   `protobuf:"varint,19,opt,name=round,proto3" json:"round,omitempty"`
	Epoch                         int64   `protobuf:"varint,20,opt,name=epoch,proto3" json:"epoch,omitempty"`
	AssetSymbol                   string  `protobuf:"bytes,21,opt,name=asset_symbol,json=assetSymbol,proto3" json:"asset_symbol,omitempty"`
}

func (x *EnhancedEAMercury) Reset() {
	*x = EnhancedEAMercury{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telem_enhanced_ea_mercury_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnhancedEAMercury) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnhancedEAMercury) ProtoMessage() {}

func (x *EnhancedEAMercury) ProtoReflect() protoreflect.Message {
	mi := &file_telem_enhanced_ea_mercury_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnhancedEAMercury.ProtoReflect.Descriptor instead.
func (*EnhancedEAMercury) Descriptor() ([]byte, []int) {
	return file_telem_enhanced_ea_mercury_proto_rawDescGZIP(), []int{0}
}

func (x *EnhancedEAMercury) GetDataSource() string {
	if x != nil {
		return x.DataSource
	}
	return ""
}

func (x *EnhancedEAMercury) GetDpBenchmarkPrice() float64 {
	if x != nil {
		return x.DpBenchmarkPrice
	}
	return 0
}

func (x *EnhancedEAMercury) GetDpBid() float64 {
	if x != nil {
		return x.DpBid
	}
	return 0
}

func (x *EnhancedEAMercury) GetDpAsk() float64 {
	if x != nil {
		return x.DpAsk
	}
	return 0
}

func (x *EnhancedEAMercury) GetCurrentBlockNumber() int64 {
	if x != nil {
		return x.CurrentBlockNumber
	}
	return 0
}

func (x *EnhancedEAMercury) GetCurrentBlockHash() string {
	if x != nil {
		return x.CurrentBlockHash
	}
	return ""
}

func (x *EnhancedEAMercury) GetCurrentBlockTimestamp() uint64 {
	if x != nil {
		return x.CurrentBlockTimestamp
	}
	return 0
}

func (x *EnhancedEAMercury) GetBridgeTaskRunStartedTimestamp() int64 {
	if x != nil {
		return x.BridgeTaskRunStartedTimestamp
	}
	return 0
}

func (x *EnhancedEAMercury) GetBridgeTaskRunEndedTimestamp() int64 {
	if x != nil {
		return x.BridgeTaskRunEndedTimestamp
	}
	return 0
}

func (x *EnhancedEAMercury) GetProviderRequestedTimestamp() int64 {
	if x != nil {
		return x.ProviderRequestedTimestamp
	}
	return 0
}

func (x *EnhancedEAMercury) GetProviderReceivedTimestamp() int64 {
	if x != nil {
		return x.ProviderReceivedTimestamp
	}
	return 0
}

func (x *EnhancedEAMercury) GetProviderDataStreamEstablished() int64 {
	if x != nil {
		return x.ProviderDataStreamEstablished
	}
	return 0
}

func (x *EnhancedEAMercury) GetProviderIndicatedTime() int64 {
	if x != nil {
		return x.ProviderIndicatedTime
	}
	return 0
}

func (x *EnhancedEAMercury) GetFeed() string {
	if x != nil {
		return x.Feed
	}
	return ""
}

func (x *EnhancedEAMercury) GetObservationBenchmarkPrice() int64 {
	if x != nil {
		return x.ObservationBenchmarkPrice
	}
	return 0
}

func (x *EnhancedEAMercury) GetObservationBid() int64 {
	if x != nil {
		return x.ObservationBid
	}
	return 0
}

func (x *EnhancedEAMercury) GetObservationAsk() int64 {
	if x != nil {
		return x.ObservationAsk
	}
	return 0
}

func (x *EnhancedEAMercury) GetConfigDigest() string {
	if x != nil {
		return x.ConfigDigest
	}
	return ""
}

func (x *EnhancedEAMercury) GetRound() int64 {
	if x != nil {
		return x.Round
	}
	return 0
}

func (x *EnhancedEAMercury) GetEpoch() int64 {
	if x != nil {
		return x.Epoch
	}
	return 0
}

func (x *EnhancedEAMercury) GetAssetSymbol() string {
	if x != nil {
		return x.AssetSymbol
	}
	return ""
}

var File_telem_enhanced_ea_mercury_proto protoreflect.FileDescriptor

var file_telem_enhanced_ea_mercury_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x5f, 0x65, 0x6e, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x64,
	0x5f, 0x65, 0x61, 0x5f, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x22, 0xd5, 0x07, 0x0a, 0x11, 0x45, 0x6e, 0x68,
	0x61, 0x6e, 0x63, 0x65, 0x64, 0x45, 0x41, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x12, 0x1f,
	0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x2c, 0x0a, 0x12, 0x64, 0x70, 0x5f, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x64, 0x70, 0x42,
	0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x15, 0x0a,
	0x06, 0x64, 0x70, 0x5f, 0x62, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x64,
	0x70, 0x42, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x64, 0x70, 0x5f, 0x61, 0x73, 0x6b, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x64, 0x70, 0x41, 0x73, 0x6b, 0x12, 0x30, 0x0a, 0x14, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2c, 0x0a,
	0x12, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x36, 0x0a, 0x17, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x15, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x48, 0x0a, 0x21, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x5f, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1d,
	0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x44, 0x0a,
	0x1f, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x72, 0x75, 0x6e,
	0x5f, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1b, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x75, 0x6e, 0x45, 0x6e, 0x64, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x40, 0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1a, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x3e, 0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x19, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x47, 0x0a, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x65, 0x73,
	0x74, 0x61, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x1d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x45, 0x73, 0x74, 0x61, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x36,
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x64, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x15, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x65, 0x65, 0x64, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x65, 0x65, 0x64, 0x12, 0x3e, 0x0a, 0x1b, 0x6f, 0x62,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d,
	0x61, 0x72, 0x6b, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x19, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x65, 0x6e, 0x63,
	0x68, 0x6d, 0x61, 0x72, 0x6b, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x62,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x69, 0x64, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0e, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x61, 0x73, 0x6b, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6f, 0x62,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x73, 0x6b, 0x12, 0x23, 0x0a, 0x0d,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x12, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x69, 0x67, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x21, 0x0a,
	0x0c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x42, 0x4e, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6b, 0x69, 0x74, 0x2f,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x68,
	0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x6d,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_telem_enhanced_ea_mercury_proto_rawDescOnce sync.Once
	file_telem_enhanced_ea_mercury_proto_rawDescData = file_telem_enhanced_ea_mercury_proto_rawDesc
)

func file_telem_enhanced_ea_mercury_proto_rawDescGZIP() []byte {
	file_telem_enhanced_ea_mercury_proto_rawDescOnce.Do(func() {
		file_telem_enhanced_ea_mercury_proto_rawDescData = protoimpl.X.CompressGZIP(file_telem_enhanced_ea_mercury_proto_rawDescData)
	})
	return file_telem_enhanced_ea_mercury_proto_rawDescData
}

var file_telem_enhanced_ea_mercury_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_telem_enhanced_ea_mercury_proto_goTypes = []interface{}{
	(*EnhancedEAMercury)(nil), // 0: telem.EnhancedEAMercury
}
var file_telem_enhanced_ea_mercury_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_telem_enhanced_ea_mercury_proto_init() }
func file_telem_enhanced_ea_mercury_proto_init() {
	if File_telem_enhanced_ea_mercury_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_telem_enhanced_ea_mercury_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnhancedEAMercury); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_telem_enhanced_ea_mercury_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_telem_enhanced_ea_mercury_proto_goTypes,
		DependencyIndexes: file_telem_enhanced_ea_mercury_proto_depIdxs,
		MessageInfos:      file_telem_enhanced_ea_mercury_proto_msgTypes,
	}.Build()
	File_telem_enhanced_ea_mercury_proto = out.File
	file_telem_enhanced_ea_mercury_proto_rawDesc = nil
	file_telem_enhanced_ea_mercury_proto_goTypes = nil
	file_telem_enhanced_ea_mercury_proto_depIdxs = nil
}
