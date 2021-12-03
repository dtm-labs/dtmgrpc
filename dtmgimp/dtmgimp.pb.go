/*
 * Copyright (c) 2021 yedf. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: dtmgrpc/dtmgimp/dtmgimp.proto

package dtmgimp

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DtmTransOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WaitResult    bool  `protobuf:"varint,1,opt,name=WaitResult,proto3" json:"WaitResult,omitempty"`
	TimeoutToFail int64 `protobuf:"varint,2,opt,name=TimeoutToFail,proto3" json:"TimeoutToFail,omitempty"`
	RetryInterval int64 `protobuf:"varint,3,opt,name=RetryInterval,proto3" json:"RetryInterval,omitempty"`
}

func (x *DtmTransOptions) Reset() {
	*x = DtmTransOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dtmgrpc_dtmgimp_dtmgimp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DtmTransOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DtmTransOptions) ProtoMessage() {}

func (x *DtmTransOptions) ProtoReflect() protoreflect.Message {
	mi := &file_dtmgrpc_dtmgimp_dtmgimp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DtmTransOptions.ProtoReflect.Descriptor instead.
func (*DtmTransOptions) Descriptor() ([]byte, []int) {
	return file_dtmgrpc_dtmgimp_dtmgimp_proto_rawDescGZIP(), []int{0}
}

func (x *DtmTransOptions) GetWaitResult() bool {
	if x != nil {
		return x.WaitResult
	}
	return false
}

func (x *DtmTransOptions) GetTimeoutToFail() int64 {
	if x != nil {
		return x.TimeoutToFail
	}
	return 0
}

func (x *DtmTransOptions) GetRetryInterval() int64 {
	if x != nil {
		return x.RetryInterval
	}
	return 0
}

// DtmRequest request sent to dtm server
type DtmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gid           string           `protobuf:"bytes,1,opt,name=Gid,proto3" json:"Gid,omitempty"`
	TransType     string           `protobuf:"bytes,2,opt,name=TransType,proto3" json:"TransType,omitempty"`
	TransOptions  *DtmTransOptions `protobuf:"bytes,3,opt,name=TransOptions,proto3" json:"TransOptions,omitempty"`
	CustomedData  string           `protobuf:"bytes,4,opt,name=CustomedData,proto3" json:"CustomedData,omitempty"`
	BinPayloads   [][]byte         `protobuf:"bytes,5,rep,name=BinPayloads,proto3" json:"BinPayloads,omitempty"`     // for MSG/SAGA branch payloads
	QueryPrepared string           `protobuf:"bytes,6,opt,name=QueryPrepared,proto3" json:"QueryPrepared,omitempty"` // for MSG
	Steps         string           `protobuf:"bytes,7,opt,name=Steps,proto3" json:"Steps,omitempty"`
}

func (x *DtmRequest) Reset() {
	*x = DtmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dtmgrpc_dtmgimp_dtmgimp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DtmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DtmRequest) ProtoMessage() {}

func (x *DtmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dtmgrpc_dtmgimp_dtmgimp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DtmRequest.ProtoReflect.Descriptor instead.
func (*DtmRequest) Descriptor() ([]byte, []int) {
	return file_dtmgrpc_dtmgimp_dtmgimp_proto_rawDescGZIP(), []int{1}
}

func (x *DtmRequest) GetGid() string {
	if x != nil {
		return x.Gid
	}
	return ""
}

func (x *DtmRequest) GetTransType() string {
	if x != nil {
		return x.TransType
	}
	return ""
}

func (x *DtmRequest) GetTransOptions() *DtmTransOptions {
	if x != nil {
		return x.TransOptions
	}
	return nil
}

func (x *DtmRequest) GetCustomedData() string {
	if x != nil {
		return x.CustomedData
	}
	return ""
}

func (x *DtmRequest) GetBinPayloads() [][]byte {
	if x != nil {
		return x.BinPayloads
	}
	return nil
}

func (x *DtmRequest) GetQueryPrepared() string {
	if x != nil {
		return x.QueryPrepared
	}
	return ""
}

func (x *DtmRequest) GetSteps() string {
	if x != nil {
		return x.Steps
	}
	return ""
}

type DtmGidReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gid string `protobuf:"bytes,1,opt,name=Gid,proto3" json:"Gid,omitempty"`
}

func (x *DtmGidReply) Reset() {
	*x = DtmGidReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dtmgrpc_dtmgimp_dtmgimp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DtmGidReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DtmGidReply) ProtoMessage() {}

func (x *DtmGidReply) ProtoReflect() protoreflect.Message {
	mi := &file_dtmgrpc_dtmgimp_dtmgimp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DtmGidReply.ProtoReflect.Descriptor instead.
func (*DtmGidReply) Descriptor() ([]byte, []int) {
	return file_dtmgrpc_dtmgimp_dtmgimp_proto_rawDescGZIP(), []int{2}
}

func (x *DtmGidReply) GetGid() string {
	if x != nil {
		return x.Gid
	}
	return ""
}

type DtmBranchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gid         string            `protobuf:"bytes,1,opt,name=Gid,proto3" json:"Gid,omitempty"`
	TransType   string            `protobuf:"bytes,2,opt,name=TransType,proto3" json:"TransType,omitempty"`
	BranchID    string            `protobuf:"bytes,3,opt,name=BranchID,proto3" json:"BranchID,omitempty"`
	Op          string            `protobuf:"bytes,4,opt,name=Op,proto3" json:"Op,omitempty"`
	Data        map[string]string `protobuf:"bytes,5,rep,name=Data,proto3" json:"Data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	BusiPayload []byte            `protobuf:"bytes,6,opt,name=BusiPayload,proto3" json:"BusiPayload,omitempty"`
}

func (x *DtmBranchRequest) Reset() {
	*x = DtmBranchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dtmgrpc_dtmgimp_dtmgimp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DtmBranchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DtmBranchRequest) ProtoMessage() {}

func (x *DtmBranchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dtmgrpc_dtmgimp_dtmgimp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DtmBranchRequest.ProtoReflect.Descriptor instead.
func (*DtmBranchRequest) Descriptor() ([]byte, []int) {
	return file_dtmgrpc_dtmgimp_dtmgimp_proto_rawDescGZIP(), []int{3}
}

func (x *DtmBranchRequest) GetGid() string {
	if x != nil {
		return x.Gid
	}
	return ""
}

func (x *DtmBranchRequest) GetTransType() string {
	if x != nil {
		return x.TransType
	}
	return ""
}

func (x *DtmBranchRequest) GetBranchID() string {
	if x != nil {
		return x.BranchID
	}
	return ""
}

func (x *DtmBranchRequest) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

func (x *DtmBranchRequest) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *DtmBranchRequest) GetBusiPayload() []byte {
	if x != nil {
		return x.BusiPayload
	}
	return nil
}

var File_dtmgrpc_dtmgimp_dtmgimp_proto protoreflect.FileDescriptor

var file_dtmgrpc_dtmgimp_dtmgimp_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x64, 0x74, 0x6d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x64, 0x74, 0x6d, 0x67, 0x69, 0x6d,
	0x70, 0x2f, 0x64, 0x74, 0x6d, 0x67, 0x69, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x64, 0x74, 0x6d, 0x67, 0x69, 0x6d, 0x70, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x0f, 0x44, 0x74, 0x6d, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x57, 0x61, 0x69, 0x74,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x57, 0x61,
	0x69, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x54, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x54, 0x6f, 0x46, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0d, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x54, 0x6f, 0x46, 0x61, 0x69, 0x6c, 0x12, 0x24,
	0x0a, 0x0d, 0x52, 0x65, 0x74, 0x72, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x52, 0x65, 0x74, 0x72, 0x79, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x22, 0xfc, 0x01, 0x0a, 0x0a, 0x44, 0x74, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x47, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x47, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64, 0x74, 0x6d, 0x67,
	0x69, 0x6d, 0x70, 0x2e, 0x44, 0x74, 0x6d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x64, 0x44, 0x61, 0x74,
	0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x69, 0x6e, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0b, 0x42, 0x69, 0x6e, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x53, 0x74, 0x65, 0x70, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x74,
	0x65, 0x70, 0x73, 0x22, 0x1f, 0x0a, 0x0b, 0x44, 0x74, 0x6d, 0x47, 0x69, 0x64, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x47, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x47, 0x69, 0x64, 0x22, 0x82, 0x02, 0x0a, 0x10, 0x44, 0x74, 0x6d, 0x42, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x47, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x47, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x42, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x42, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x4f, 0x70, 0x12, 0x37, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x64, 0x74, 0x6d, 0x67, 0x69, 0x6d, 0x70, 0x2e, 0x44, 0x74,
	0x6d, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x20,
	0x0a, 0x0b, 0x42, 0x75, 0x73, 0x69, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0b, 0x42, 0x75, 0x73, 0x69, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xb1, 0x02, 0x0a, 0x03, 0x44, 0x74,
	0x6d, 0x12, 0x38, 0x0a, 0x06, 0x4e, 0x65, 0x77, 0x47, 0x69, 0x64, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x64, 0x74, 0x6d, 0x67, 0x69, 0x6d, 0x70, 0x2e, 0x44, 0x74,
	0x6d, 0x47, 0x69, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x06, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x12, 0x13, 0x2e, 0x64, 0x74, 0x6d, 0x67, 0x69, 0x6d, 0x70, 0x2e,
	0x44, 0x74, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x12,
	0x13, 0x2e, 0x64, 0x74, 0x6d, 0x67, 0x69, 0x6d, 0x70, 0x2e, 0x44, 0x74, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x36,
	0x0a, 0x05, 0x41, 0x62, 0x6f, 0x72, 0x74, 0x12, 0x13, 0x2e, 0x64, 0x74, 0x6d, 0x67, 0x69, 0x6d,
	0x70, 0x2e, 0x44, 0x74, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x19, 0x2e, 0x64, 0x74, 0x6d, 0x67, 0x69,
	0x6d, 0x70, 0x2e, 0x44, 0x74, 0x6d, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x0b, 0x5a,
	0x09, 0x2e, 0x2f, 0x64, 0x74, 0x6d, 0x67, 0x69, 0x6d, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_dtmgrpc_dtmgimp_dtmgimp_proto_rawDescOnce sync.Once
	file_dtmgrpc_dtmgimp_dtmgimp_proto_rawDescData = file_dtmgrpc_dtmgimp_dtmgimp_proto_rawDesc
)

func file_dtmgrpc_dtmgimp_dtmgimp_proto_rawDescGZIP() []byte {
	file_dtmgrpc_dtmgimp_dtmgimp_proto_rawDescOnce.Do(func() {
		file_dtmgrpc_dtmgimp_dtmgimp_proto_rawDescData = protoimpl.X.CompressGZIP(file_dtmgrpc_dtmgimp_dtmgimp_proto_rawDescData)
	})
	return file_dtmgrpc_dtmgimp_dtmgimp_proto_rawDescData
}

var file_dtmgrpc_dtmgimp_dtmgimp_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_dtmgrpc_dtmgimp_dtmgimp_proto_goTypes = []interface{}{
	(*DtmTransOptions)(nil),  // 0: dtmgimp.DtmTransOptions
	(*DtmRequest)(nil),       // 1: dtmgimp.DtmRequest
	(*DtmGidReply)(nil),      // 2: dtmgimp.DtmGidReply
	(*DtmBranchRequest)(nil), // 3: dtmgimp.DtmBranchRequest
	nil,                      // 4: dtmgimp.DtmBranchRequest.DataEntry
	(*emptypb.Empty)(nil),    // 5: google.protobuf.Empty
}
var file_dtmgrpc_dtmgimp_dtmgimp_proto_depIdxs = []int32{
	0, // 0: dtmgimp.DtmRequest.TransOptions:type_name -> dtmgimp.DtmTransOptions
	4, // 1: dtmgimp.DtmBranchRequest.Data:type_name -> dtmgimp.DtmBranchRequest.DataEntry
	5, // 2: dtmgimp.Dtm.NewGid:input_type -> google.protobuf.Empty
	1, // 3: dtmgimp.Dtm.Submit:input_type -> dtmgimp.DtmRequest
	1, // 4: dtmgimp.Dtm.Prepare:input_type -> dtmgimp.DtmRequest
	1, // 5: dtmgimp.Dtm.Abort:input_type -> dtmgimp.DtmRequest
	3, // 6: dtmgimp.Dtm.RegisterBranch:input_type -> dtmgimp.DtmBranchRequest
	2, // 7: dtmgimp.Dtm.NewGid:output_type -> dtmgimp.DtmGidReply
	5, // 8: dtmgimp.Dtm.Submit:output_type -> google.protobuf.Empty
	5, // 9: dtmgimp.Dtm.Prepare:output_type -> google.protobuf.Empty
	5, // 10: dtmgimp.Dtm.Abort:output_type -> google.protobuf.Empty
	5, // 11: dtmgimp.Dtm.RegisterBranch:output_type -> google.protobuf.Empty
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_dtmgrpc_dtmgimp_dtmgimp_proto_init() }
func file_dtmgrpc_dtmgimp_dtmgimp_proto_init() {
	if File_dtmgrpc_dtmgimp_dtmgimp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dtmgrpc_dtmgimp_dtmgimp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DtmTransOptions); i {
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
		file_dtmgrpc_dtmgimp_dtmgimp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DtmRequest); i {
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
		file_dtmgrpc_dtmgimp_dtmgimp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DtmGidReply); i {
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
		file_dtmgrpc_dtmgimp_dtmgimp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DtmBranchRequest); i {
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
			RawDescriptor: file_dtmgrpc_dtmgimp_dtmgimp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dtmgrpc_dtmgimp_dtmgimp_proto_goTypes,
		DependencyIndexes: file_dtmgrpc_dtmgimp_dtmgimp_proto_depIdxs,
		MessageInfos:      file_dtmgrpc_dtmgimp_dtmgimp_proto_msgTypes,
	}.Build()
	File_dtmgrpc_dtmgimp_dtmgimp_proto = out.File
	file_dtmgrpc_dtmgimp_dtmgimp_proto_rawDesc = nil
	file_dtmgrpc_dtmgimp_dtmgimp_proto_goTypes = nil
	file_dtmgrpc_dtmgimp_dtmgimp_proto_depIdxs = nil
}
