// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/server/api/errordetails/v1/message.proto

package errordetails

import (
	bytes "bytes"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type TaskAlreadyStartedFailure struct {
}

func (m *TaskAlreadyStartedFailure) Reset()      { *m = TaskAlreadyStartedFailure{} }
func (*TaskAlreadyStartedFailure) ProtoMessage() {}
func (*TaskAlreadyStartedFailure) Descriptor() ([]byte, []int) {
	return fileDescriptor_73580c2e9c4cb332, []int{0}
}
func (m *TaskAlreadyStartedFailure) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TaskAlreadyStartedFailure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TaskAlreadyStartedFailure.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TaskAlreadyStartedFailure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskAlreadyStartedFailure.Merge(m, src)
}
func (m *TaskAlreadyStartedFailure) XXX_Size() int {
	return m.Size()
}
func (m *TaskAlreadyStartedFailure) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskAlreadyStartedFailure.DiscardUnknown(m)
}

var xxx_messageInfo_TaskAlreadyStartedFailure proto.InternalMessageInfo

type CurrentBranchChangedFailure struct {
	CurrentBranchToken []byte `protobuf:"bytes,1,opt,name=current_branch_token,json=currentBranchToken,proto3" json:"current_branch_token,omitempty"`
	RequestBranchToken []byte `protobuf:"bytes,2,opt,name=request_branch_token,json=requestBranchToken,proto3" json:"request_branch_token,omitempty"`
}

func (m *CurrentBranchChangedFailure) Reset()      { *m = CurrentBranchChangedFailure{} }
func (*CurrentBranchChangedFailure) ProtoMessage() {}
func (*CurrentBranchChangedFailure) Descriptor() ([]byte, []int) {
	return fileDescriptor_73580c2e9c4cb332, []int{1}
}
func (m *CurrentBranchChangedFailure) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CurrentBranchChangedFailure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CurrentBranchChangedFailure.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CurrentBranchChangedFailure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrentBranchChangedFailure.Merge(m, src)
}
func (m *CurrentBranchChangedFailure) XXX_Size() int {
	return m.Size()
}
func (m *CurrentBranchChangedFailure) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrentBranchChangedFailure.DiscardUnknown(m)
}

var xxx_messageInfo_CurrentBranchChangedFailure proto.InternalMessageInfo

func (m *CurrentBranchChangedFailure) GetCurrentBranchToken() []byte {
	if m != nil {
		return m.CurrentBranchToken
	}
	return nil
}

func (m *CurrentBranchChangedFailure) GetRequestBranchToken() []byte {
	if m != nil {
		return m.RequestBranchToken
	}
	return nil
}

type ShardOwnershipLostFailure struct {
	OwnerHost   string `protobuf:"bytes,1,opt,name=owner_host,json=ownerHost,proto3" json:"owner_host,omitempty"`
	CurrentHost string `protobuf:"bytes,2,opt,name=current_host,json=currentHost,proto3" json:"current_host,omitempty"`
}

func (m *ShardOwnershipLostFailure) Reset()      { *m = ShardOwnershipLostFailure{} }
func (*ShardOwnershipLostFailure) ProtoMessage() {}
func (*ShardOwnershipLostFailure) Descriptor() ([]byte, []int) {
	return fileDescriptor_73580c2e9c4cb332, []int{2}
}
func (m *ShardOwnershipLostFailure) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ShardOwnershipLostFailure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ShardOwnershipLostFailure.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ShardOwnershipLostFailure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShardOwnershipLostFailure.Merge(m, src)
}
func (m *ShardOwnershipLostFailure) XXX_Size() int {
	return m.Size()
}
func (m *ShardOwnershipLostFailure) XXX_DiscardUnknown() {
	xxx_messageInfo_ShardOwnershipLostFailure.DiscardUnknown(m)
}

var xxx_messageInfo_ShardOwnershipLostFailure proto.InternalMessageInfo

func (m *ShardOwnershipLostFailure) GetOwnerHost() string {
	if m != nil {
		return m.OwnerHost
	}
	return ""
}

func (m *ShardOwnershipLostFailure) GetCurrentHost() string {
	if m != nil {
		return m.CurrentHost
	}
	return ""
}

type RetryReplicationFailure struct {
	NamespaceId       string `protobuf:"bytes,1,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id,omitempty"`
	WorkflowId        string `protobuf:"bytes,2,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	RunId             string `protobuf:"bytes,3,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	StartEventId      int64  `protobuf:"varint,4,opt,name=start_event_id,json=startEventId,proto3" json:"start_event_id,omitempty"`
	StartEventVersion int64  `protobuf:"varint,5,opt,name=start_event_version,json=startEventVersion,proto3" json:"start_event_version,omitempty"`
	EndEventId        int64  `protobuf:"varint,6,opt,name=end_event_id,json=endEventId,proto3" json:"end_event_id,omitempty"`
	EndEventVersion   int64  `protobuf:"varint,7,opt,name=end_event_version,json=endEventVersion,proto3" json:"end_event_version,omitempty"`
}

func (m *RetryReplicationFailure) Reset()      { *m = RetryReplicationFailure{} }
func (*RetryReplicationFailure) ProtoMessage() {}
func (*RetryReplicationFailure) Descriptor() ([]byte, []int) {
	return fileDescriptor_73580c2e9c4cb332, []int{3}
}
func (m *RetryReplicationFailure) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RetryReplicationFailure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RetryReplicationFailure.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RetryReplicationFailure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetryReplicationFailure.Merge(m, src)
}
func (m *RetryReplicationFailure) XXX_Size() int {
	return m.Size()
}
func (m *RetryReplicationFailure) XXX_DiscardUnknown() {
	xxx_messageInfo_RetryReplicationFailure.DiscardUnknown(m)
}

var xxx_messageInfo_RetryReplicationFailure proto.InternalMessageInfo

func (m *RetryReplicationFailure) GetNamespaceId() string {
	if m != nil {
		return m.NamespaceId
	}
	return ""
}

func (m *RetryReplicationFailure) GetWorkflowId() string {
	if m != nil {
		return m.WorkflowId
	}
	return ""
}

func (m *RetryReplicationFailure) GetRunId() string {
	if m != nil {
		return m.RunId
	}
	return ""
}

func (m *RetryReplicationFailure) GetStartEventId() int64 {
	if m != nil {
		return m.StartEventId
	}
	return 0
}

func (m *RetryReplicationFailure) GetStartEventVersion() int64 {
	if m != nil {
		return m.StartEventVersion
	}
	return 0
}

func (m *RetryReplicationFailure) GetEndEventId() int64 {
	if m != nil {
		return m.EndEventId
	}
	return 0
}

func (m *RetryReplicationFailure) GetEndEventVersion() int64 {
	if m != nil {
		return m.EndEventVersion
	}
	return 0
}

func init() {
	proto.RegisterType((*TaskAlreadyStartedFailure)(nil), "temporal.server.api.errordetails.v1.TaskAlreadyStartedFailure")
	proto.RegisterType((*CurrentBranchChangedFailure)(nil), "temporal.server.api.errordetails.v1.CurrentBranchChangedFailure")
	proto.RegisterType((*ShardOwnershipLostFailure)(nil), "temporal.server.api.errordetails.v1.ShardOwnershipLostFailure")
	proto.RegisterType((*RetryReplicationFailure)(nil), "temporal.server.api.errordetails.v1.RetryReplicationFailure")
}

func init() {
	proto.RegisterFile("temporal/server/api/errordetails/v1/message.proto", fileDescriptor_73580c2e9c4cb332)
}

var fileDescriptor_73580c2e9c4cb332 = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0x77, 0x52, 0x1b, 0xe9, 0x24, 0x28, 0x5d, 0x15, 0x53, 0x8a, 0x63, 0x1a, 0x3d, 0x14,
	0x0f, 0x1b, 0x83, 0xe0, 0xc5, 0x93, 0x2d, 0x8a, 0x01, 0x41, 0x48, 0x8b, 0x07, 0x41, 0xc2, 0x74,
	0xe7, 0x99, 0x1d, 0xb2, 0x99, 0x59, 0xdf, 0xcc, 0x6e, 0xe8, 0x4d, 0xbf, 0x81, 0x47, 0x3f, 0x82,
	0x1f, 0xc5, 0x63, 0x8e, 0x3d, 0x9a, 0xcd, 0xc5, 0x63, 0x3f, 0x82, 0xec, 0x6c, 0x36, 0x59, 0x3c,
	0xf4, 0x38, 0xff, 0xff, 0xef, 0xff, 0x7f, 0x8f, 0xe1, 0xd1, 0x81, 0x85, 0x59, 0xa2, 0x91, 0xc7,
	0x7d, 0x03, 0x98, 0x01, 0xf6, 0x79, 0x22, 0xfb, 0x80, 0xa8, 0x51, 0x80, 0xe5, 0x32, 0x36, 0xfd,
	0x6c, 0xd0, 0x9f, 0x81, 0x31, 0x7c, 0x02, 0x41, 0x82, 0xda, 0x6a, 0xff, 0x49, 0x15, 0x09, 0xca,
	0x48, 0xc0, 0x13, 0x19, 0xd4, 0x23, 0x41, 0x36, 0xe8, 0x1d, 0xd2, 0x83, 0x73, 0x6e, 0xa6, 0xaf,
	0x63, 0x04, 0x2e, 0x2e, 0xcf, 0x2c, 0x47, 0x0b, 0xe2, 0x2d, 0x97, 0x71, 0x8a, 0xd0, 0xfb, 0x4e,
	0xe8, 0xe1, 0x69, 0x8a, 0x08, 0xca, 0x9e, 0x20, 0x57, 0x61, 0x74, 0x1a, 0x71, 0x35, 0xd9, 0xf8,
	0xfe, 0x73, 0x7a, 0x3f, 0x2c, 0xed, 0xf1, 0x85, 0xf3, 0xc7, 0x56, 0x4f, 0x41, 0x75, 0x48, 0x97,
	0x1c, 0xb7, 0x47, 0x7e, 0x58, 0x8f, 0x9e, 0x17, 0x4e, 0x91, 0x40, 0xf8, 0x9a, 0x82, 0xf9, 0x2f,
	0xd1, 0x28, 0x13, 0x6b, 0xaf, 0x96, 0xe8, 0x7d, 0xa6, 0x07, 0x67, 0x11, 0x47, 0xf1, 0x61, 0xae,
	0x00, 0x4d, 0x24, 0x93, 0xf7, 0xda, 0xd8, 0x6a, 0x81, 0x47, 0x94, 0xea, 0x42, 0x1f, 0x47, 0xda,
	0x58, 0x37, 0x76, 0x6f, 0xb4, 0xe7, 0x94, 0x77, 0xda, 0x58, 0xff, 0x88, 0xb6, 0xab, 0xfd, 0x1c,
	0xd0, 0x70, 0x40, 0x6b, 0xad, 0x15, 0x48, 0xef, 0x67, 0x83, 0x3e, 0x1c, 0x81, 0xc5, 0xcb, 0x11,
	0x24, 0xb1, 0x0c, 0xb9, 0x95, 0x5a, 0x55, 0xed, 0x47, 0xb4, 0xad, 0xf8, 0x0c, 0x4c, 0xc2, 0x43,
	0x18, 0x4b, 0xb1, 0xee, 0x6f, 0x6d, 0xb4, 0xa1, 0xf0, 0x1f, 0xd3, 0xd6, 0x5c, 0xe3, 0xf4, 0x4b,
	0xac, 0xe7, 0x05, 0x51, 0x0e, 0xa0, 0x95, 0x34, 0x14, 0xfe, 0x03, 0xda, 0xc4, 0x54, 0x15, 0xde,
	0x8e, 0xf3, 0x76, 0x31, 0x55, 0x43, 0xe1, 0x3f, 0xa5, 0x77, 0x4c, 0xf1, 0xd7, 0x63, 0xc8, 0x8a,
	0xed, 0xa4, 0xe8, 0xdc, 0xea, 0x92, 0xe3, 0x9d, 0x51, 0xdb, 0xa9, 0x6f, 0x0a, 0x71, 0x28, 0xfc,
	0x80, 0xde, 0xab, 0x53, 0x19, 0xa0, 0x91, 0x5a, 0x75, 0x76, 0x1d, 0xba, 0xbf, 0x45, 0x3f, 0x96,
	0x86, 0xdf, 0xa5, 0x6d, 0x50, 0x62, 0xdb, 0xd9, 0x74, 0x20, 0x05, 0x25, 0xaa, 0xc6, 0x67, 0x74,
	0x7f, 0x4b, 0x54, 0x7d, 0xb7, 0x1d, 0x76, 0xb7, 0xc2, 0xd6, 0x6d, 0x27, 0xf1, 0x62, 0xc9, 0xbc,
	0xab, 0x25, 0xf3, 0xae, 0x97, 0x8c, 0x7c, 0xcb, 0x19, 0xf9, 0x95, 0x33, 0xf2, 0x3b, 0x67, 0x64,
	0x91, 0x33, 0xf2, 0x27, 0x67, 0xe4, 0x6f, 0xce, 0xbc, 0xeb, 0x9c, 0x91, 0x1f, 0x2b, 0xe6, 0x2d,
	0x56, 0xcc, 0xbb, 0x5a, 0x31, 0xef, 0xd3, 0xcb, 0x89, 0x0e, 0x36, 0x87, 0x27, 0xf5, 0x0d, 0xe7,
	0xfa, 0xaa, 0xfe, 0xbe, 0x68, 0xba, 0xa3, 0x7d, 0xf1, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xa9, 0xaa,
	0xa0, 0x9f, 0xe9, 0x02, 0x00, 0x00,
}

func (this *TaskAlreadyStartedFailure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TaskAlreadyStartedFailure)
	if !ok {
		that2, ok := that.(TaskAlreadyStartedFailure)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *CurrentBranchChangedFailure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CurrentBranchChangedFailure)
	if !ok {
		that2, ok := that.(CurrentBranchChangedFailure)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.CurrentBranchToken, that1.CurrentBranchToken) {
		return false
	}
	if !bytes.Equal(this.RequestBranchToken, that1.RequestBranchToken) {
		return false
	}
	return true
}
func (this *ShardOwnershipLostFailure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ShardOwnershipLostFailure)
	if !ok {
		that2, ok := that.(ShardOwnershipLostFailure)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.OwnerHost != that1.OwnerHost {
		return false
	}
	if this.CurrentHost != that1.CurrentHost {
		return false
	}
	return true
}
func (this *RetryReplicationFailure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RetryReplicationFailure)
	if !ok {
		that2, ok := that.(RetryReplicationFailure)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.NamespaceId != that1.NamespaceId {
		return false
	}
	if this.WorkflowId != that1.WorkflowId {
		return false
	}
	if this.RunId != that1.RunId {
		return false
	}
	if this.StartEventId != that1.StartEventId {
		return false
	}
	if this.StartEventVersion != that1.StartEventVersion {
		return false
	}
	if this.EndEventId != that1.EndEventId {
		return false
	}
	if this.EndEventVersion != that1.EndEventVersion {
		return false
	}
	return true
}
func (this *TaskAlreadyStartedFailure) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&errordetails.TaskAlreadyStartedFailure{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *CurrentBranchChangedFailure) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&errordetails.CurrentBranchChangedFailure{")
	s = append(s, "CurrentBranchToken: "+fmt.Sprintf("%#v", this.CurrentBranchToken)+",\n")
	s = append(s, "RequestBranchToken: "+fmt.Sprintf("%#v", this.RequestBranchToken)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ShardOwnershipLostFailure) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&errordetails.ShardOwnershipLostFailure{")
	s = append(s, "OwnerHost: "+fmt.Sprintf("%#v", this.OwnerHost)+",\n")
	s = append(s, "CurrentHost: "+fmt.Sprintf("%#v", this.CurrentHost)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *RetryReplicationFailure) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 11)
	s = append(s, "&errordetails.RetryReplicationFailure{")
	s = append(s, "NamespaceId: "+fmt.Sprintf("%#v", this.NamespaceId)+",\n")
	s = append(s, "WorkflowId: "+fmt.Sprintf("%#v", this.WorkflowId)+",\n")
	s = append(s, "RunId: "+fmt.Sprintf("%#v", this.RunId)+",\n")
	s = append(s, "StartEventId: "+fmt.Sprintf("%#v", this.StartEventId)+",\n")
	s = append(s, "StartEventVersion: "+fmt.Sprintf("%#v", this.StartEventVersion)+",\n")
	s = append(s, "EndEventId: "+fmt.Sprintf("%#v", this.EndEventId)+",\n")
	s = append(s, "EndEventVersion: "+fmt.Sprintf("%#v", this.EndEventVersion)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringMessage(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *TaskAlreadyStartedFailure) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TaskAlreadyStartedFailure) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TaskAlreadyStartedFailure) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *CurrentBranchChangedFailure) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CurrentBranchChangedFailure) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CurrentBranchChangedFailure) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RequestBranchToken) > 0 {
		i -= len(m.RequestBranchToken)
		copy(dAtA[i:], m.RequestBranchToken)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.RequestBranchToken)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.CurrentBranchToken) > 0 {
		i -= len(m.CurrentBranchToken)
		copy(dAtA[i:], m.CurrentBranchToken)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.CurrentBranchToken)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ShardOwnershipLostFailure) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ShardOwnershipLostFailure) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ShardOwnershipLostFailure) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CurrentHost) > 0 {
		i -= len(m.CurrentHost)
		copy(dAtA[i:], m.CurrentHost)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.CurrentHost)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.OwnerHost) > 0 {
		i -= len(m.OwnerHost)
		copy(dAtA[i:], m.OwnerHost)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.OwnerHost)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RetryReplicationFailure) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RetryReplicationFailure) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RetryReplicationFailure) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EndEventVersion != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.EndEventVersion))
		i--
		dAtA[i] = 0x38
	}
	if m.EndEventId != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.EndEventId))
		i--
		dAtA[i] = 0x30
	}
	if m.StartEventVersion != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.StartEventVersion))
		i--
		dAtA[i] = 0x28
	}
	if m.StartEventId != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.StartEventId))
		i--
		dAtA[i] = 0x20
	}
	if len(m.RunId) > 0 {
		i -= len(m.RunId)
		copy(dAtA[i:], m.RunId)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.RunId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.WorkflowId) > 0 {
		i -= len(m.WorkflowId)
		copy(dAtA[i:], m.WorkflowId)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.WorkflowId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.NamespaceId) > 0 {
		i -= len(m.NamespaceId)
		copy(dAtA[i:], m.NamespaceId)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.NamespaceId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TaskAlreadyStartedFailure) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *CurrentBranchChangedFailure) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CurrentBranchToken)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.RequestBranchToken)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	return n
}

func (m *ShardOwnershipLostFailure) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OwnerHost)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.CurrentHost)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	return n
}

func (m *RetryReplicationFailure) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NamespaceId)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.WorkflowId)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.RunId)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.StartEventId != 0 {
		n += 1 + sovMessage(uint64(m.StartEventId))
	}
	if m.StartEventVersion != 0 {
		n += 1 + sovMessage(uint64(m.StartEventVersion))
	}
	if m.EndEventId != 0 {
		n += 1 + sovMessage(uint64(m.EndEventId))
	}
	if m.EndEventVersion != 0 {
		n += 1 + sovMessage(uint64(m.EndEventVersion))
	}
	return n
}

func sovMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *TaskAlreadyStartedFailure) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TaskAlreadyStartedFailure{`,
		`}`,
	}, "")
	return s
}
func (this *CurrentBranchChangedFailure) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CurrentBranchChangedFailure{`,
		`CurrentBranchToken:` + fmt.Sprintf("%v", this.CurrentBranchToken) + `,`,
		`RequestBranchToken:` + fmt.Sprintf("%v", this.RequestBranchToken) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ShardOwnershipLostFailure) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ShardOwnershipLostFailure{`,
		`OwnerHost:` + fmt.Sprintf("%v", this.OwnerHost) + `,`,
		`CurrentHost:` + fmt.Sprintf("%v", this.CurrentHost) + `,`,
		`}`,
	}, "")
	return s
}
func (this *RetryReplicationFailure) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RetryReplicationFailure{`,
		`NamespaceId:` + fmt.Sprintf("%v", this.NamespaceId) + `,`,
		`WorkflowId:` + fmt.Sprintf("%v", this.WorkflowId) + `,`,
		`RunId:` + fmt.Sprintf("%v", this.RunId) + `,`,
		`StartEventId:` + fmt.Sprintf("%v", this.StartEventId) + `,`,
		`StartEventVersion:` + fmt.Sprintf("%v", this.StartEventVersion) + `,`,
		`EndEventId:` + fmt.Sprintf("%v", this.EndEventId) + `,`,
		`EndEventVersion:` + fmt.Sprintf("%v", this.EndEventVersion) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringMessage(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *TaskAlreadyStartedFailure) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TaskAlreadyStartedFailure: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TaskAlreadyStartedFailure: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CurrentBranchChangedFailure) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CurrentBranchChangedFailure: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CurrentBranchChangedFailure: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentBranchToken", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CurrentBranchToken = append(m.CurrentBranchToken[:0], dAtA[iNdEx:postIndex]...)
			if m.CurrentBranchToken == nil {
				m.CurrentBranchToken = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestBranchToken", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestBranchToken = append(m.RequestBranchToken[:0], dAtA[iNdEx:postIndex]...)
			if m.RequestBranchToken == nil {
				m.RequestBranchToken = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ShardOwnershipLostFailure) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ShardOwnershipLostFailure: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ShardOwnershipLostFailure: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerHost", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OwnerHost = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentHost", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CurrentHost = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RetryReplicationFailure) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RetryReplicationFailure: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RetryReplicationFailure: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NamespaceId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NamespaceId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WorkflowId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WorkflowId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RunId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RunId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartEventId", wireType)
			}
			m.StartEventId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartEventId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartEventVersion", wireType)
			}
			m.StartEventVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartEventVersion |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndEventId", wireType)
			}
			m.EndEventId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndEventId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndEventVersion", wireType)
			}
			m.EndEventVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndEventVersion |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessage = fmt.Errorf("proto: unexpected end of group")
)
