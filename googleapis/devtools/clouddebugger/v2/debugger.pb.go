// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/clouddebugger/v2/debugger.proto

package clouddebugger // import "github.com/sgtsquiggs/go-genproto/googleapis/devtools/clouddebugger/v2"

import proto "github.com/sgtsquiggs/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/sgtsquiggs/protobuf/ptypes/empty"
import _ "github.com/sgtsquiggs/go-genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "github.com/sgtsquiggs/grpc-go"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Request to set a breakpoint
type SetBreakpointRequest struct {
	// ID of the debuggee where the breakpoint is to be set.
	DebuggeeId string `protobuf:"bytes,1,opt,name=debuggee_id,json=debuggeeId,proto3" json:"debuggee_id,omitempty"`
	// Breakpoint specification to set.
	// The field `location` of the breakpoint must be set.
	Breakpoint *Breakpoint `protobuf:"bytes,2,opt,name=breakpoint,proto3" json:"breakpoint,omitempty"`
	// The client version making the call.
	// Schema: `domain/type/version` (e.g., `google.com/intellij/v1`).
	ClientVersion        string   `protobuf:"bytes,4,opt,name=client_version,json=clientVersion,proto3" json:"client_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetBreakpointRequest) Reset()         { *m = SetBreakpointRequest{} }
func (m *SetBreakpointRequest) String() string { return proto.CompactTextString(m) }
func (*SetBreakpointRequest) ProtoMessage()    {}
func (*SetBreakpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_debugger_04096ebac35efe72, []int{0}
}
func (m *SetBreakpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetBreakpointRequest.Unmarshal(m, b)
}
func (m *SetBreakpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetBreakpointRequest.Marshal(b, m, deterministic)
}
func (dst *SetBreakpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetBreakpointRequest.Merge(dst, src)
}
func (m *SetBreakpointRequest) XXX_Size() int {
	return xxx_messageInfo_SetBreakpointRequest.Size(m)
}
func (m *SetBreakpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetBreakpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetBreakpointRequest proto.InternalMessageInfo

func (m *SetBreakpointRequest) GetDebuggeeId() string {
	if m != nil {
		return m.DebuggeeId
	}
	return ""
}

func (m *SetBreakpointRequest) GetBreakpoint() *Breakpoint {
	if m != nil {
		return m.Breakpoint
	}
	return nil
}

func (m *SetBreakpointRequest) GetClientVersion() string {
	if m != nil {
		return m.ClientVersion
	}
	return ""
}

// Response for setting a breakpoint.
type SetBreakpointResponse struct {
	// Breakpoint resource.
	// The field `id` is guaranteed to be set (in addition to the echoed fileds).
	Breakpoint           *Breakpoint `protobuf:"bytes,1,opt,name=breakpoint,proto3" json:"breakpoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SetBreakpointResponse) Reset()         { *m = SetBreakpointResponse{} }
func (m *SetBreakpointResponse) String() string { return proto.CompactTextString(m) }
func (*SetBreakpointResponse) ProtoMessage()    {}
func (*SetBreakpointResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_debugger_04096ebac35efe72, []int{1}
}
func (m *SetBreakpointResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetBreakpointResponse.Unmarshal(m, b)
}
func (m *SetBreakpointResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetBreakpointResponse.Marshal(b, m, deterministic)
}
func (dst *SetBreakpointResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetBreakpointResponse.Merge(dst, src)
}
func (m *SetBreakpointResponse) XXX_Size() int {
	return xxx_messageInfo_SetBreakpointResponse.Size(m)
}
func (m *SetBreakpointResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetBreakpointResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetBreakpointResponse proto.InternalMessageInfo

func (m *SetBreakpointResponse) GetBreakpoint() *Breakpoint {
	if m != nil {
		return m.Breakpoint
	}
	return nil
}

// Request to get breakpoint information.
type GetBreakpointRequest struct {
	// ID of the debuggee whose breakpoint to get.
	DebuggeeId string `protobuf:"bytes,1,opt,name=debuggee_id,json=debuggeeId,proto3" json:"debuggee_id,omitempty"`
	// ID of the breakpoint to get.
	BreakpointId string `protobuf:"bytes,2,opt,name=breakpoint_id,json=breakpointId,proto3" json:"breakpoint_id,omitempty"`
	// The client version making the call.
	// Schema: `domain/type/version` (e.g., `google.com/intellij/v1`).
	ClientVersion        string   `protobuf:"bytes,4,opt,name=client_version,json=clientVersion,proto3" json:"client_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBreakpointRequest) Reset()         { *m = GetBreakpointRequest{} }
func (m *GetBreakpointRequest) String() string { return proto.CompactTextString(m) }
func (*GetBreakpointRequest) ProtoMessage()    {}
func (*GetBreakpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_debugger_04096ebac35efe72, []int{2}
}
func (m *GetBreakpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBreakpointRequest.Unmarshal(m, b)
}
func (m *GetBreakpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBreakpointRequest.Marshal(b, m, deterministic)
}
func (dst *GetBreakpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBreakpointRequest.Merge(dst, src)
}
func (m *GetBreakpointRequest) XXX_Size() int {
	return xxx_messageInfo_GetBreakpointRequest.Size(m)
}
func (m *GetBreakpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBreakpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBreakpointRequest proto.InternalMessageInfo

func (m *GetBreakpointRequest) GetDebuggeeId() string {
	if m != nil {
		return m.DebuggeeId
	}
	return ""
}

func (m *GetBreakpointRequest) GetBreakpointId() string {
	if m != nil {
		return m.BreakpointId
	}
	return ""
}

func (m *GetBreakpointRequest) GetClientVersion() string {
	if m != nil {
		return m.ClientVersion
	}
	return ""
}

// Response for getting breakpoint information.
type GetBreakpointResponse struct {
	// Complete breakpoint state.
	// The fields `id` and `location` are guaranteed to be set.
	Breakpoint           *Breakpoint `protobuf:"bytes,1,opt,name=breakpoint,proto3" json:"breakpoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetBreakpointResponse) Reset()         { *m = GetBreakpointResponse{} }
func (m *GetBreakpointResponse) String() string { return proto.CompactTextString(m) }
func (*GetBreakpointResponse) ProtoMessage()    {}
func (*GetBreakpointResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_debugger_04096ebac35efe72, []int{3}
}
func (m *GetBreakpointResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBreakpointResponse.Unmarshal(m, b)
}
func (m *GetBreakpointResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBreakpointResponse.Marshal(b, m, deterministic)
}
func (dst *GetBreakpointResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBreakpointResponse.Merge(dst, src)
}
func (m *GetBreakpointResponse) XXX_Size() int {
	return xxx_messageInfo_GetBreakpointResponse.Size(m)
}
func (m *GetBreakpointResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBreakpointResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBreakpointResponse proto.InternalMessageInfo

func (m *GetBreakpointResponse) GetBreakpoint() *Breakpoint {
	if m != nil {
		return m.Breakpoint
	}
	return nil
}

// Request to delete a breakpoint.
type DeleteBreakpointRequest struct {
	// ID of the debuggee whose breakpoint to delete.
	DebuggeeId string `protobuf:"bytes,1,opt,name=debuggee_id,json=debuggeeId,proto3" json:"debuggee_id,omitempty"`
	// ID of the breakpoint to delete.
	BreakpointId string `protobuf:"bytes,2,opt,name=breakpoint_id,json=breakpointId,proto3" json:"breakpoint_id,omitempty"`
	// The client version making the call.
	// Schema: `domain/type/version` (e.g., `google.com/intellij/v1`).
	ClientVersion        string   `protobuf:"bytes,3,opt,name=client_version,json=clientVersion,proto3" json:"client_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBreakpointRequest) Reset()         { *m = DeleteBreakpointRequest{} }
func (m *DeleteBreakpointRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteBreakpointRequest) ProtoMessage()    {}
func (*DeleteBreakpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_debugger_04096ebac35efe72, []int{4}
}
func (m *DeleteBreakpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBreakpointRequest.Unmarshal(m, b)
}
func (m *DeleteBreakpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBreakpointRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteBreakpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBreakpointRequest.Merge(dst, src)
}
func (m *DeleteBreakpointRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteBreakpointRequest.Size(m)
}
func (m *DeleteBreakpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBreakpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBreakpointRequest proto.InternalMessageInfo

func (m *DeleteBreakpointRequest) GetDebuggeeId() string {
	if m != nil {
		return m.DebuggeeId
	}
	return ""
}

func (m *DeleteBreakpointRequest) GetBreakpointId() string {
	if m != nil {
		return m.BreakpointId
	}
	return ""
}

func (m *DeleteBreakpointRequest) GetClientVersion() string {
	if m != nil {
		return m.ClientVersion
	}
	return ""
}

// Request to list breakpoints.
type ListBreakpointsRequest struct {
	// ID of the debuggee whose breakpoints to list.
	DebuggeeId string `protobuf:"bytes,1,opt,name=debuggee_id,json=debuggeeId,proto3" json:"debuggee_id,omitempty"`
	// When set to `true`, the response includes the list of breakpoints set by
	// any user. Otherwise, it includes only breakpoints set by the caller.
	IncludeAllUsers bool `protobuf:"varint,2,opt,name=include_all_users,json=includeAllUsers,proto3" json:"include_all_users,omitempty"`
	// When set to `true`, the response includes active and inactive
	// breakpoints. Otherwise, it includes only active breakpoints.
	IncludeInactive bool `protobuf:"varint,3,opt,name=include_inactive,json=includeInactive,proto3" json:"include_inactive,omitempty"`
	// When set, the response includes only breakpoints with the specified action.
	Action *ListBreakpointsRequest_BreakpointActionValue `protobuf:"bytes,4,opt,name=action,proto3" json:"action,omitempty"`
	// This field is deprecated. The following fields are always stripped out of
	// the result: `stack_frames`, `evaluated_expressions` and `variable_table`.
	StripResults bool `protobuf:"varint,5,opt,name=strip_results,json=stripResults,proto3" json:"strip_results,omitempty"`
	// A wait token that, if specified, blocks the call until the breakpoints
	// list has changed, or a server selected timeout has expired.  The value
	// should be set from the last response. The error code
	// `google.rpc.Code.ABORTED` (RPC) is returned on wait timeout, which
	// should be called again with the same `wait_token`.
	WaitToken string `protobuf:"bytes,6,opt,name=wait_token,json=waitToken,proto3" json:"wait_token,omitempty"`
	// The client version making the call.
	// Schema: `domain/type/version` (e.g., `google.com/intellij/v1`).
	ClientVersion        string   `protobuf:"bytes,8,opt,name=client_version,json=clientVersion,proto3" json:"client_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBreakpointsRequest) Reset()         { *m = ListBreakpointsRequest{} }
func (m *ListBreakpointsRequest) String() string { return proto.CompactTextString(m) }
func (*ListBreakpointsRequest) ProtoMessage()    {}
func (*ListBreakpointsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_debugger_04096ebac35efe72, []int{5}
}
func (m *ListBreakpointsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBreakpointsRequest.Unmarshal(m, b)
}
func (m *ListBreakpointsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBreakpointsRequest.Marshal(b, m, deterministic)
}
func (dst *ListBreakpointsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBreakpointsRequest.Merge(dst, src)
}
func (m *ListBreakpointsRequest) XXX_Size() int {
	return xxx_messageInfo_ListBreakpointsRequest.Size(m)
}
func (m *ListBreakpointsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBreakpointsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListBreakpointsRequest proto.InternalMessageInfo

func (m *ListBreakpointsRequest) GetDebuggeeId() string {
	if m != nil {
		return m.DebuggeeId
	}
	return ""
}

func (m *ListBreakpointsRequest) GetIncludeAllUsers() bool {
	if m != nil {
		return m.IncludeAllUsers
	}
	return false
}

func (m *ListBreakpointsRequest) GetIncludeInactive() bool {
	if m != nil {
		return m.IncludeInactive
	}
	return false
}

func (m *ListBreakpointsRequest) GetAction() *ListBreakpointsRequest_BreakpointActionValue {
	if m != nil {
		return m.Action
	}
	return nil
}

func (m *ListBreakpointsRequest) GetStripResults() bool {
	if m != nil {
		return m.StripResults
	}
	return false
}

func (m *ListBreakpointsRequest) GetWaitToken() string {
	if m != nil {
		return m.WaitToken
	}
	return ""
}

func (m *ListBreakpointsRequest) GetClientVersion() string {
	if m != nil {
		return m.ClientVersion
	}
	return ""
}

// Wrapper message for `Breakpoint.Action`. Defines a filter on the action
// field of breakpoints.
type ListBreakpointsRequest_BreakpointActionValue struct {
	// Only breakpoints with the specified action will pass the filter.
	Value                Breakpoint_Action `protobuf:"varint,1,opt,name=value,proto3,enum=google.devtools.clouddebugger.v2.Breakpoint_Action" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListBreakpointsRequest_BreakpointActionValue) Reset() {
	*m = ListBreakpointsRequest_BreakpointActionValue{}
}
func (m *ListBreakpointsRequest_BreakpointActionValue) String() string {
	return proto.CompactTextString(m)
}
func (*ListBreakpointsRequest_BreakpointActionValue) ProtoMessage() {}
func (*ListBreakpointsRequest_BreakpointActionValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_debugger_04096ebac35efe72, []int{5, 0}
}
func (m *ListBreakpointsRequest_BreakpointActionValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBreakpointsRequest_BreakpointActionValue.Unmarshal(m, b)
}
func (m *ListBreakpointsRequest_BreakpointActionValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBreakpointsRequest_BreakpointActionValue.Marshal(b, m, deterministic)
}
func (dst *ListBreakpointsRequest_BreakpointActionValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBreakpointsRequest_BreakpointActionValue.Merge(dst, src)
}
func (m *ListBreakpointsRequest_BreakpointActionValue) XXX_Size() int {
	return xxx_messageInfo_ListBreakpointsRequest_BreakpointActionValue.Size(m)
}
func (m *ListBreakpointsRequest_BreakpointActionValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBreakpointsRequest_BreakpointActionValue.DiscardUnknown(m)
}

var xxx_messageInfo_ListBreakpointsRequest_BreakpointActionValue proto.InternalMessageInfo

func (m *ListBreakpointsRequest_BreakpointActionValue) GetValue() Breakpoint_Action {
	if m != nil {
		return m.Value
	}
	return Breakpoint_CAPTURE
}

// Response for listing breakpoints.
type ListBreakpointsResponse struct {
	// List of breakpoints matching the request.
	// The fields `id` and `location` are guaranteed to be set on each breakpoint.
	// The fields: `stack_frames`, `evaluated_expressions` and `variable_table`
	// are cleared on each breakpoint regardless of its status.
	Breakpoints []*Breakpoint `protobuf:"bytes,1,rep,name=breakpoints,proto3" json:"breakpoints,omitempty"`
	// A wait token that can be used in the next call to `list` (REST) or
	// `ListBreakpoints` (RPC) to block until the list of breakpoints has changes.
	NextWaitToken        string   `protobuf:"bytes,2,opt,name=next_wait_token,json=nextWaitToken,proto3" json:"next_wait_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBreakpointsResponse) Reset()         { *m = ListBreakpointsResponse{} }
func (m *ListBreakpointsResponse) String() string { return proto.CompactTextString(m) }
func (*ListBreakpointsResponse) ProtoMessage()    {}
func (*ListBreakpointsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_debugger_04096ebac35efe72, []int{6}
}
func (m *ListBreakpointsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBreakpointsResponse.Unmarshal(m, b)
}
func (m *ListBreakpointsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBreakpointsResponse.Marshal(b, m, deterministic)
}
func (dst *ListBreakpointsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBreakpointsResponse.Merge(dst, src)
}
func (m *ListBreakpointsResponse) XXX_Size() int {
	return xxx_messageInfo_ListBreakpointsResponse.Size(m)
}
func (m *ListBreakpointsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBreakpointsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListBreakpointsResponse proto.InternalMessageInfo

func (m *ListBreakpointsResponse) GetBreakpoints() []*Breakpoint {
	if m != nil {
		return m.Breakpoints
	}
	return nil
}

func (m *ListBreakpointsResponse) GetNextWaitToken() string {
	if m != nil {
		return m.NextWaitToken
	}
	return ""
}

// Request to list debuggees.
type ListDebuggeesRequest struct {
	// Project number of a Google Cloud project whose debuggees to list.
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// When set to `true`, the result includes all debuggees. Otherwise, the
	// result includes only debuggees that are active.
	IncludeInactive bool `protobuf:"varint,3,opt,name=include_inactive,json=includeInactive,proto3" json:"include_inactive,omitempty"`
	// The client version making the call.
	// Schema: `domain/type/version` (e.g., `google.com/intellij/v1`).
	ClientVersion        string   `protobuf:"bytes,4,opt,name=client_version,json=clientVersion,proto3" json:"client_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListDebuggeesRequest) Reset()         { *m = ListDebuggeesRequest{} }
func (m *ListDebuggeesRequest) String() string { return proto.CompactTextString(m) }
func (*ListDebuggeesRequest) ProtoMessage()    {}
func (*ListDebuggeesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_debugger_04096ebac35efe72, []int{7}
}
func (m *ListDebuggeesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDebuggeesRequest.Unmarshal(m, b)
}
func (m *ListDebuggeesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDebuggeesRequest.Marshal(b, m, deterministic)
}
func (dst *ListDebuggeesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDebuggeesRequest.Merge(dst, src)
}
func (m *ListDebuggeesRequest) XXX_Size() int {
	return xxx_messageInfo_ListDebuggeesRequest.Size(m)
}
func (m *ListDebuggeesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDebuggeesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListDebuggeesRequest proto.InternalMessageInfo

func (m *ListDebuggeesRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *ListDebuggeesRequest) GetIncludeInactive() bool {
	if m != nil {
		return m.IncludeInactive
	}
	return false
}

func (m *ListDebuggeesRequest) GetClientVersion() string {
	if m != nil {
		return m.ClientVersion
	}
	return ""
}

// Response for listing debuggees.
type ListDebuggeesResponse struct {
	// List of debuggees accessible to the calling user.
	// The fields `debuggee.id` and `description` are guaranteed to be set.
	// The `description` field is a human readable field provided by agents and
	// can be displayed to users.
	Debuggees            []*Debuggee `protobuf:"bytes,1,rep,name=debuggees,proto3" json:"debuggees,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListDebuggeesResponse) Reset()         { *m = ListDebuggeesResponse{} }
func (m *ListDebuggeesResponse) String() string { return proto.CompactTextString(m) }
func (*ListDebuggeesResponse) ProtoMessage()    {}
func (*ListDebuggeesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_debugger_04096ebac35efe72, []int{8}
}
func (m *ListDebuggeesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDebuggeesResponse.Unmarshal(m, b)
}
func (m *ListDebuggeesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDebuggeesResponse.Marshal(b, m, deterministic)
}
func (dst *ListDebuggeesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDebuggeesResponse.Merge(dst, src)
}
func (m *ListDebuggeesResponse) XXX_Size() int {
	return xxx_messageInfo_ListDebuggeesResponse.Size(m)
}
func (m *ListDebuggeesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDebuggeesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListDebuggeesResponse proto.InternalMessageInfo

func (m *ListDebuggeesResponse) GetDebuggees() []*Debuggee {
	if m != nil {
		return m.Debuggees
	}
	return nil
}

func init() {
	proto.RegisterType((*SetBreakpointRequest)(nil), "google.devtools.clouddebugger.v2.SetBreakpointRequest")
	proto.RegisterType((*SetBreakpointResponse)(nil), "google.devtools.clouddebugger.v2.SetBreakpointResponse")
	proto.RegisterType((*GetBreakpointRequest)(nil), "google.devtools.clouddebugger.v2.GetBreakpointRequest")
	proto.RegisterType((*GetBreakpointResponse)(nil), "google.devtools.clouddebugger.v2.GetBreakpointResponse")
	proto.RegisterType((*DeleteBreakpointRequest)(nil), "google.devtools.clouddebugger.v2.DeleteBreakpointRequest")
	proto.RegisterType((*ListBreakpointsRequest)(nil), "google.devtools.clouddebugger.v2.ListBreakpointsRequest")
	proto.RegisterType((*ListBreakpointsRequest_BreakpointActionValue)(nil), "google.devtools.clouddebugger.v2.ListBreakpointsRequest.BreakpointActionValue")
	proto.RegisterType((*ListBreakpointsResponse)(nil), "google.devtools.clouddebugger.v2.ListBreakpointsResponse")
	proto.RegisterType((*ListDebuggeesRequest)(nil), "google.devtools.clouddebugger.v2.ListDebuggeesRequest")
	proto.RegisterType((*ListDebuggeesResponse)(nil), "google.devtools.clouddebugger.v2.ListDebuggeesResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Debugger2Client is the client API for Debugger2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/github.com/sgtsquiggs/grpc-go#ClientConn.NewStream.
type Debugger2Client interface {
	// Sets the breakpoint to the debuggee.
	SetBreakpoint(ctx context.Context, in *SetBreakpointRequest, opts ...grpc.CallOption) (*SetBreakpointResponse, error)
	// Gets breakpoint information.
	GetBreakpoint(ctx context.Context, in *GetBreakpointRequest, opts ...grpc.CallOption) (*GetBreakpointResponse, error)
	// Deletes the breakpoint from the debuggee.
	DeleteBreakpoint(ctx context.Context, in *DeleteBreakpointRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Lists all breakpoints for the debuggee.
	ListBreakpoints(ctx context.Context, in *ListBreakpointsRequest, opts ...grpc.CallOption) (*ListBreakpointsResponse, error)
	// Lists all the debuggees that the user has access to.
	ListDebuggees(ctx context.Context, in *ListDebuggeesRequest, opts ...grpc.CallOption) (*ListDebuggeesResponse, error)
}

type debugger2Client struct {
	cc *grpc.ClientConn
}

func NewDebugger2Client(cc *grpc.ClientConn) Debugger2Client {
	return &debugger2Client{cc}
}

func (c *debugger2Client) SetBreakpoint(ctx context.Context, in *SetBreakpointRequest, opts ...grpc.CallOption) (*SetBreakpointResponse, error) {
	out := new(SetBreakpointResponse)
	err := c.cc.Invoke(ctx, "/google.devtools.clouddebugger.v2.Debugger2/SetBreakpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugger2Client) GetBreakpoint(ctx context.Context, in *GetBreakpointRequest, opts ...grpc.CallOption) (*GetBreakpointResponse, error) {
	out := new(GetBreakpointResponse)
	err := c.cc.Invoke(ctx, "/google.devtools.clouddebugger.v2.Debugger2/GetBreakpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugger2Client) DeleteBreakpoint(ctx context.Context, in *DeleteBreakpointRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.devtools.clouddebugger.v2.Debugger2/DeleteBreakpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugger2Client) ListBreakpoints(ctx context.Context, in *ListBreakpointsRequest, opts ...grpc.CallOption) (*ListBreakpointsResponse, error) {
	out := new(ListBreakpointsResponse)
	err := c.cc.Invoke(ctx, "/google.devtools.clouddebugger.v2.Debugger2/ListBreakpoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugger2Client) ListDebuggees(ctx context.Context, in *ListDebuggeesRequest, opts ...grpc.CallOption) (*ListDebuggeesResponse, error) {
	out := new(ListDebuggeesResponse)
	err := c.cc.Invoke(ctx, "/google.devtools.clouddebugger.v2.Debugger2/ListDebuggees", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Debugger2Server is the server API for Debugger2 service.
type Debugger2Server interface {
	// Sets the breakpoint to the debuggee.
	SetBreakpoint(context.Context, *SetBreakpointRequest) (*SetBreakpointResponse, error)
	// Gets breakpoint information.
	GetBreakpoint(context.Context, *GetBreakpointRequest) (*GetBreakpointResponse, error)
	// Deletes the breakpoint from the debuggee.
	DeleteBreakpoint(context.Context, *DeleteBreakpointRequest) (*empty.Empty, error)
	// Lists all breakpoints for the debuggee.
	ListBreakpoints(context.Context, *ListBreakpointsRequest) (*ListBreakpointsResponse, error)
	// Lists all the debuggees that the user has access to.
	ListDebuggees(context.Context, *ListDebuggeesRequest) (*ListDebuggeesResponse, error)
}

func RegisterDebugger2Server(s *grpc.Server, srv Debugger2Server) {
	s.RegisterService(&_Debugger2_serviceDesc, srv)
}

func _Debugger2_SetBreakpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetBreakpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Debugger2Server).SetBreakpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.clouddebugger.v2.Debugger2/SetBreakpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Debugger2Server).SetBreakpoint(ctx, req.(*SetBreakpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Debugger2_GetBreakpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBreakpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Debugger2Server).GetBreakpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.clouddebugger.v2.Debugger2/GetBreakpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Debugger2Server).GetBreakpoint(ctx, req.(*GetBreakpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Debugger2_DeleteBreakpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBreakpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Debugger2Server).DeleteBreakpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.clouddebugger.v2.Debugger2/DeleteBreakpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Debugger2Server).DeleteBreakpoint(ctx, req.(*DeleteBreakpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Debugger2_ListBreakpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBreakpointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Debugger2Server).ListBreakpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.clouddebugger.v2.Debugger2/ListBreakpoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Debugger2Server).ListBreakpoints(ctx, req.(*ListBreakpointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Debugger2_ListDebuggees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDebuggeesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Debugger2Server).ListDebuggees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.clouddebugger.v2.Debugger2/ListDebuggees",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Debugger2Server).ListDebuggees(ctx, req.(*ListDebuggeesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Debugger2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.devtools.clouddebugger.v2.Debugger2",
	HandlerType: (*Debugger2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetBreakpoint",
			Handler:    _Debugger2_SetBreakpoint_Handler,
		},
		{
			MethodName: "GetBreakpoint",
			Handler:    _Debugger2_GetBreakpoint_Handler,
		},
		{
			MethodName: "DeleteBreakpoint",
			Handler:    _Debugger2_DeleteBreakpoint_Handler,
		},
		{
			MethodName: "ListBreakpoints",
			Handler:    _Debugger2_ListBreakpoints_Handler,
		},
		{
			MethodName: "ListDebuggees",
			Handler:    _Debugger2_ListDebuggees_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/devtools/clouddebugger/v2/debugger.proto",
}

func init() {
	proto.RegisterFile("google/devtools/clouddebugger/v2/debugger.proto", fileDescriptor_debugger_04096ebac35efe72)
}

var fileDescriptor_debugger_04096ebac35efe72 = []byte{
	// 797 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xdd, 0x6a, 0xdb, 0x48,
	0x14, 0x66, 0x9c, 0xcd, 0x8f, 0x8f, 0xe3, 0x24, 0x3b, 0xe4, 0x47, 0x78, 0xff, 0x8c, 0xf6, 0x87,
	0x6c, 0x76, 0x91, 0x16, 0x65, 0xd9, 0x4d, 0x76, 0x6f, 0x1a, 0x37, 0xc5, 0x31, 0xa4, 0x21, 0xb8,
	0xad, 0x0b, 0x25, 0x60, 0x64, 0x7b, 0x22, 0xd4, 0x28, 0x1a, 0x55, 0x33, 0x72, 0x5b, 0x42, 0x6e,
	0x52, 0xe8, 0x7d, 0xe9, 0x0b, 0xf4, 0xba, 0x14, 0xfa, 0x02, 0x2d, 0xf4, 0xae, 0x90, 0xdb, 0xbe,
	0x42, 0x1f, 0xa4, 0x48, 0x9a, 0x89, 0x65, 0x57, 0xad, 0x2d, 0x07, 0x72, 0x37, 0xfa, 0x66, 0xce,
	0x99, 0xef, 0xfb, 0xe6, 0xcc, 0x19, 0x81, 0x6e, 0x51, 0x6a, 0x39, 0x44, 0xef, 0x90, 0x2e, 0xa7,
	0xd4, 0x61, 0x7a, 0xdb, 0xa1, 0x41, 0xa7, 0x43, 0x5a, 0x81, 0x65, 0x11, 0x5f, 0xef, 0x1a, 0xba,
	0x1c, 0x6b, 0x9e, 0x4f, 0x39, 0xc5, 0xe5, 0x38, 0x40, 0x93, 0x01, 0x5a, 0x5f, 0x80, 0xd6, 0x35,
	0x4a, 0xdf, 0x8b, 0x94, 0xa6, 0x67, 0xeb, 0xa6, 0xeb, 0x52, 0x6e, 0x72, 0x9b, 0xba, 0x2c, 0x8e,
	0x2f, 0xfd, 0x31, 0x7c, 0x43, 0x93, 0x9b, 0x62, 0xf1, 0x77, 0x62, 0x71, 0xf4, 0xd5, 0x0a, 0x0e,
	0x75, 0x72, 0xec, 0xf1, 0xc7, 0xf1, 0xa4, 0xfa, 0x0a, 0xc1, 0xe2, 0x2d, 0xc2, 0x2b, 0x3e, 0x31,
	0x8f, 0x3c, 0x6a, 0xbb, 0xbc, 0x4e, 0x1e, 0x04, 0x84, 0x71, 0xfc, 0x13, 0x14, 0x44, 0x3e, 0xd2,
	0xb4, 0x3b, 0x0a, 0x2a, 0xa3, 0xd5, 0x7c, 0x1d, 0x24, 0x54, 0xeb, 0xe0, 0x5d, 0x80, 0xd6, 0x45,
	0x94, 0x92, 0x2b, 0xa3, 0xd5, 0x82, 0xf1, 0xa7, 0x36, 0x4c, 0x98, 0x96, 0xd8, 0x29, 0x11, 0x8f,
	0x7f, 0x85, 0xb9, 0xb6, 0x63, 0x13, 0x97, 0x37, 0xbb, 0xc4, 0x67, 0x36, 0x75, 0x95, 0x6f, 0xa2,
	0x1d, 0x8b, 0x31, 0xda, 0x88, 0x41, 0x95, 0xc0, 0xd2, 0x00, 0x5b, 0xe6, 0x51, 0x97, 0x91, 0x01,
	0x36, 0xe8, 0x72, 0x6c, 0xd4, 0x27, 0x08, 0x16, 0xab, 0x63, 0xb9, 0xf2, 0x33, 0x14, 0x7b, 0x79,
	0xc2, 0x25, 0xb9, 0x68, 0xc9, 0x6c, 0x0f, 0xac, 0x75, 0x32, 0x88, 0xad, 0x5e, 0x81, 0xd8, 0xa7,
	0x08, 0x56, 0xb6, 0x89, 0x43, 0x38, 0xb9, 0x3a, 0xbd, 0x13, 0x69, 0x7a, 0xdf, 0x4f, 0xc0, 0xf2,
	0xae, 0xcd, 0x12, 0x8a, 0xd9, 0xc8, 0x3c, 0xd6, 0xe0, 0x5b, 0xdb, 0x6d, 0x3b, 0x41, 0x87, 0x34,
	0x4d, 0xc7, 0x69, 0x06, 0x8c, 0xf8, 0x2c, 0xe2, 0x32, 0x53, 0x9f, 0x17, 0x13, 0x5b, 0x8e, 0x73,
	0x27, 0x84, 0xf1, 0xef, 0xb0, 0x20, 0xd7, 0xda, 0xae, 0xd9, 0xe6, 0x76, 0x97, 0x44, 0x84, 0x7a,
	0x4b, 0x6b, 0x02, 0xc6, 0x87, 0x30, 0x15, 0x8e, 0xc4, 0x09, 0x15, 0x8c, 0xbd, 0xe1, 0x2e, 0xa7,
	0x2b, 0x48, 0x98, 0xbf, 0x15, 0x25, 0x6c, 0x98, 0x4e, 0x40, 0xea, 0x22, 0x7b, 0x68, 0x23, 0xe3,
	0xbe, 0xed, 0x35, 0x7d, 0xc2, 0x02, 0x87, 0x33, 0x65, 0x32, 0xe2, 0x33, 0x1b, 0x81, 0xf5, 0x18,
	0xc3, 0x3f, 0x00, 0x3c, 0x34, 0x6d, 0xde, 0xe4, 0xf4, 0x88, 0xb8, 0xca, 0x54, 0xe4, 0x41, 0x3e,
	0x44, 0x6e, 0x87, 0x40, 0x8a, 0xcb, 0x33, 0x29, 0x2e, 0x97, 0x5a, 0xb0, 0x94, 0xca, 0x05, 0xd7,
	0x60, 0xb2, 0x1b, 0x0e, 0x22, 0x77, 0xe7, 0x8c, 0xf5, 0x2c, 0x05, 0xa5, 0xc5, 0x89, 0xea, 0x71,
	0x06, 0xf5, 0x19, 0x82, 0x95, 0xcf, 0x7c, 0x10, 0xc5, 0xbb, 0x07, 0x85, 0x5e, 0x71, 0x30, 0x05,
	0x95, 0x27, 0x32, 0x57, 0x6f, 0x32, 0x01, 0xfe, 0x0d, 0xe6, 0x5d, 0xf2, 0x88, 0x37, 0x13, 0xd6,
	0xc4, 0x35, 0x58, 0x0c, 0xe1, 0xbb, 0xd2, 0x1e, 0xf5, 0x0c, 0xc1, 0x62, 0xc8, 0x69, 0x5b, 0x14,
	0xcd, 0x45, 0x6d, 0x29, 0x30, 0xed, 0xf9, 0xf4, 0x3e, 0x69, 0x73, 0x11, 0x28, 0x3f, 0xb3, 0x14,
	0xca, 0x88, 0x57, 0xda, 0x84, 0xa5, 0x01, 0x0e, 0xc2, 0x95, 0x1d, 0xc8, 0xcb, 0x6a, 0x96, 0x9e,
	0xac, 0x0d, 0xf7, 0x44, 0xe6, 0xa9, 0xf7, 0x82, 0x8d, 0xb7, 0xd3, 0x90, 0x17, 0xb8, 0x6f, 0xe0,
	0x73, 0x04, 0xc5, 0xbe, 0x8e, 0x89, 0xff, 0x19, 0x9e, 0x36, 0xed, 0x41, 0x28, 0xfd, 0x9b, 0x39,
	0x2e, 0x96, 0xa6, 0xee, 0x9c, 0x7d, 0xf8, 0xf8, 0x3c, 0x57, 0x51, 0xff, 0x4e, 0x3e, 0x84, 0xfa,
	0x05, 0x61, 0xfd, 0x24, 0x71, 0xb3, 0x4f, 0xf5, 0xc4, 0xd1, 0xea, 0x8c, 0xf0, 0xff, 0x92, 0x8f,
	0x44, 0x28, 0xa6, 0x9a, 0x55, 0x4c, 0x75, 0x4c, 0x31, 0xd5, 0xaf, 0x89, 0xc1, 0xd7, 0x32, 0x8b,
	0x39, 0xe9, 0xeb, 0x93, 0xa7, 0xf8, 0x35, 0x82, 0x85, 0xc1, 0xb6, 0x8b, 0x37, 0x47, 0x39, 0xf3,
	0xd4, 0x56, 0x5d, 0x5a, 0x96, 0xa1, 0xf2, 0x9d, 0xd7, 0x6e, 0x84, 0xef, 0xbc, 0x64, 0xbc, 0x76,
	0x79, 0xc6, 0xef, 0x10, 0xcc, 0x0f, 0xdc, 0x6a, 0xbc, 0x31, 0x6e, 0x43, 0x2c, 0x6d, 0x8e, 0x11,
	0x29, 0x0e, 0x61, 0x23, 0x92, 0x64, 0xe0, 0xbf, 0xb2, 0x4a, 0xc2, 0x2f, 0x10, 0x14, 0xfb, 0x2e,
	0xe0, 0x28, 0x15, 0x94, 0xd6, 0x35, 0x46, 0xa9, 0xa0, 0xd4, 0x9b, 0xae, 0xfe, 0x18, 0x91, 0x57,
	0xf0, 0x72, 0x3a, 0xf9, 0xca, 0x1b, 0x04, 0xbf, 0xb4, 0xe9, 0xf1, 0xd0, 0xf4, 0x95, 0xa2, 0xbc,
	0xe5, 0xfb, 0xe1, 0x81, 0xef, 0xa3, 0x7b, 0x37, 0x45, 0x88, 0x45, 0x1d, 0xd3, 0xb5, 0x34, 0xea,
	0x5b, 0xba, 0x45, 0xdc, 0xa8, 0x1c, 0xc4, 0x1f, 0xaa, 0xe9, 0xd9, 0xec, 0xcb, 0x3f, 0x8d, 0xff,
	0xf7, 0x01, 0x2f, 0x73, 0x4a, 0x35, 0xce, 0x77, 0x3d, 0x84, 0x65, 0xaf, 0xf1, 0xb5, 0x86, 0x71,
	0x2e, 0xa7, 0x0e, 0xa2, 0xa9, 0x03, 0x39, 0x75, 0xd0, 0x30, 0x5a, 0x53, 0xd1, 0x7e, 0xeb, 0x9f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x52, 0x23, 0xb7, 0x95, 0x14, 0x0b, 0x00, 0x00,
}
