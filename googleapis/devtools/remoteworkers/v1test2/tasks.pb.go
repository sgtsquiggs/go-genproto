// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/remoteworkers/v1test2/tasks.proto

package remoteworkers // import "github.com/sgtsquiggs/go-genproto/googleapis/devtools/remoteworkers/v1test2"

import proto "github.com/sgtsquiggs/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/sgtsquiggs/protobuf/ptypes/any"
import _ "github.com/sgtsquiggs/go-genproto/googleapis/api/annotations"
import status "github.com/sgtsquiggs/go-genproto/googleapis/rpc/status"
import field_mask "github.com/sgtsquiggs/go-genproto/protobuf/field_mask"

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

// DEPRECATED - use Lease.payload instead.
// A Task represents a unit of work. Its result and logs are defined as
// subresources.
//
// If all the `Any` fields are populated, this can be a very large message, and
// clients may not want the entire message returned on every call to every
// method. Such clients should request partial responses
// (https://cloud.google.com/apis/design/design_patterns#partial_response) and
// servers should implement partial responses in order to reduce unnecessry
// overhead.
type Task struct {
	// The name of this task. Output only.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The actual task to perform. For example, this could be CommandTask to run a
	// command line.
	Description *any.Any `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Handles to logs. The key is a human-readable name like `stdout`, and the
	// handle is a resource name that can be passed to ByteStream or other
	// accessors.
	//
	// An implementation may define some logs by default (like `stdout`), and may
	// allow clients to add new logs via AddTaskLog.
	Logs                 map[string]string `protobuf:"bytes,3,rep,name=logs,proto3" json:"logs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_tasks_667ae2bf8892e2d4, []int{0}
}
func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (dst *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(dst, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Task) GetDescription() *any.Any {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *Task) GetLogs() map[string]string {
	if m != nil {
		return m.Logs
	}
	return nil
}

// DEPRECATED - use Lease.assignment_result instead.
// The result and metadata of the task.
type TaskResult struct {
	// The name of the task result; must be a name of a `Task` followed by
	// `/result`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The result may be updated several times; the client must only set
	// `complete` to true to indicate that no further updates are allowed.
	// If this is not true, the `status` field must not be examined since its zero
	// value is equivalent to `OK`.
	//
	// Once a task is completed, it must not be updated with further results,
	// though the implementation may choose to continue to receive logs.
	Complete bool `protobuf:"varint,2,opt,name=complete,proto3" json:"complete,omitempty"`
	// The final status of the task itself. For example, if task.description
	// included a timeout which was violated, status.code may be
	// DEADLINE_EXCEEDED. This field can only be read if `complete` is true.
	Status *status.Status `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	// Any non-log output, such as output files and exit codes. See
	// CommandResult as an example.
	Output *any.Any `protobuf:"bytes,4,opt,name=output,proto3" json:"output,omitempty"`
	// Any information about how the command was executed, eg runtime. See
	// CommandOverhead as an example.
	Meta                 *any.Any `protobuf:"bytes,5,opt,name=meta,proto3" json:"meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskResult) Reset()         { *m = TaskResult{} }
func (m *TaskResult) String() string { return proto.CompactTextString(m) }
func (*TaskResult) ProtoMessage()    {}
func (*TaskResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_tasks_667ae2bf8892e2d4, []int{1}
}
func (m *TaskResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskResult.Unmarshal(m, b)
}
func (m *TaskResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskResult.Marshal(b, m, deterministic)
}
func (dst *TaskResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskResult.Merge(dst, src)
}
func (m *TaskResult) XXX_Size() int {
	return xxx_messageInfo_TaskResult.Size(m)
}
func (m *TaskResult) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskResult.DiscardUnknown(m)
}

var xxx_messageInfo_TaskResult proto.InternalMessageInfo

func (m *TaskResult) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TaskResult) GetComplete() bool {
	if m != nil {
		return m.Complete
	}
	return false
}

func (m *TaskResult) GetStatus() *status.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *TaskResult) GetOutput() *any.Any {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *TaskResult) GetMeta() *any.Any {
	if m != nil {
		return m.Meta
	}
	return nil
}

// Request message for `GetTask`.
type GetTaskRequest struct {
	// The task name.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTaskRequest) Reset()         { *m = GetTaskRequest{} }
func (m *GetTaskRequest) String() string { return proto.CompactTextString(m) }
func (*GetTaskRequest) ProtoMessage()    {}
func (*GetTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_tasks_667ae2bf8892e2d4, []int{2}
}
func (m *GetTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTaskRequest.Unmarshal(m, b)
}
func (m *GetTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTaskRequest.Marshal(b, m, deterministic)
}
func (dst *GetTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTaskRequest.Merge(dst, src)
}
func (m *GetTaskRequest) XXX_Size() int {
	return xxx_messageInfo_GetTaskRequest.Size(m)
}
func (m *GetTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTaskRequest proto.InternalMessageInfo

func (m *GetTaskRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Request message for `UpdateTaskResult`.
type UpdateTaskResultRequest struct {
	// The task result name; must match `result.name`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The result being updated.
	Result *TaskResult `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	// The fields within `result` that are specified.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// If this is being updated by a bot from BotManager, the source should be
	// bot.session_id. That way, if two bots accidentally get the same name, we'll
	// know to reject updates from the older one.
	Source               string   `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTaskResultRequest) Reset()         { *m = UpdateTaskResultRequest{} }
func (m *UpdateTaskResultRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTaskResultRequest) ProtoMessage()    {}
func (*UpdateTaskResultRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_tasks_667ae2bf8892e2d4, []int{3}
}
func (m *UpdateTaskResultRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTaskResultRequest.Unmarshal(m, b)
}
func (m *UpdateTaskResultRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTaskResultRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateTaskResultRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTaskResultRequest.Merge(dst, src)
}
func (m *UpdateTaskResultRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateTaskResultRequest.Size(m)
}
func (m *UpdateTaskResultRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTaskResultRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTaskResultRequest proto.InternalMessageInfo

func (m *UpdateTaskResultRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateTaskResultRequest) GetResult() *TaskResult {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *UpdateTaskResultRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

func (m *UpdateTaskResultRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

// Request message for `AddTaskLog`.
type AddTaskLogRequest struct {
	// The name of the task that will own the new log.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The human-readable name of the log, like `stdout` or a relative file path.
	LogId                string   `protobuf:"bytes,2,opt,name=log_id,json=logId,proto3" json:"log_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddTaskLogRequest) Reset()         { *m = AddTaskLogRequest{} }
func (m *AddTaskLogRequest) String() string { return proto.CompactTextString(m) }
func (*AddTaskLogRequest) ProtoMessage()    {}
func (*AddTaskLogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_tasks_667ae2bf8892e2d4, []int{4}
}
func (m *AddTaskLogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddTaskLogRequest.Unmarshal(m, b)
}
func (m *AddTaskLogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddTaskLogRequest.Marshal(b, m, deterministic)
}
func (dst *AddTaskLogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddTaskLogRequest.Merge(dst, src)
}
func (m *AddTaskLogRequest) XXX_Size() int {
	return xxx_messageInfo_AddTaskLogRequest.Size(m)
}
func (m *AddTaskLogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddTaskLogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddTaskLogRequest proto.InternalMessageInfo

func (m *AddTaskLogRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddTaskLogRequest) GetLogId() string {
	if m != nil {
		return m.LogId
	}
	return ""
}

// Response message for `AddTaskLog`.
type AddTaskLogResponse struct {
	// The handle for the new log, as would be returned in Task.logs.
	Handle               string   `protobuf:"bytes,1,opt,name=handle,proto3" json:"handle,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddTaskLogResponse) Reset()         { *m = AddTaskLogResponse{} }
func (m *AddTaskLogResponse) String() string { return proto.CompactTextString(m) }
func (*AddTaskLogResponse) ProtoMessage()    {}
func (*AddTaskLogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_tasks_667ae2bf8892e2d4, []int{5}
}
func (m *AddTaskLogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddTaskLogResponse.Unmarshal(m, b)
}
func (m *AddTaskLogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddTaskLogResponse.Marshal(b, m, deterministic)
}
func (dst *AddTaskLogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddTaskLogResponse.Merge(dst, src)
}
func (m *AddTaskLogResponse) XXX_Size() int {
	return xxx_messageInfo_AddTaskLogResponse.Size(m)
}
func (m *AddTaskLogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddTaskLogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddTaskLogResponse proto.InternalMessageInfo

func (m *AddTaskLogResponse) GetHandle() string {
	if m != nil {
		return m.Handle
	}
	return ""
}

func init() {
	proto.RegisterType((*Task)(nil), "google.devtools.remoteworkers.v1test2.Task")
	proto.RegisterMapType((map[string]string)(nil), "google.devtools.remoteworkers.v1test2.Task.LogsEntry")
	proto.RegisterType((*TaskResult)(nil), "google.devtools.remoteworkers.v1test2.TaskResult")
	proto.RegisterType((*GetTaskRequest)(nil), "google.devtools.remoteworkers.v1test2.GetTaskRequest")
	proto.RegisterType((*UpdateTaskResultRequest)(nil), "google.devtools.remoteworkers.v1test2.UpdateTaskResultRequest")
	proto.RegisterType((*AddTaskLogRequest)(nil), "google.devtools.remoteworkers.v1test2.AddTaskLogRequest")
	proto.RegisterType((*AddTaskLogResponse)(nil), "google.devtools.remoteworkers.v1test2.AddTaskLogResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TasksClient is the client API for Tasks service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/github.com/sgtsquiggs/grpc-go#ClientConn.NewStream.
type TasksClient interface {
	// DEPRECATED - use Lease.payload instead.
	// GetTask reads the current state of the task. Tasks must be created through
	// some other interface, and should be immutable once created and exposed to
	// the bots.
	GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*Task, error)
	// DEPRECATED - use Lease.result instead.
	// UpdateTaskResult updates the result.
	UpdateTaskResult(ctx context.Context, in *UpdateTaskResultRequest, opts ...grpc.CallOption) (*TaskResult, error)
	// DEPRECATED - precreate logs prior to sending to bot.
	// AddTaskLog creates a new streaming log. The log is streamed and marked as
	// completed through other interfaces (i.e., ByteStream). This can be called
	// by the bot if it wants to create a new log; the server can also predefine
	// logs that do not need to be created (e.g. `stdout`).
	AddTaskLog(ctx context.Context, in *AddTaskLogRequest, opts ...grpc.CallOption) (*AddTaskLogResponse, error)
}

type tasksClient struct {
	cc *grpc.ClientConn
}

func NewTasksClient(cc *grpc.ClientConn) TasksClient {
	return &tasksClient{cc}
}

func (c *tasksClient) GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/google.devtools.remoteworkers.v1test2.Tasks/GetTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksClient) UpdateTaskResult(ctx context.Context, in *UpdateTaskResultRequest, opts ...grpc.CallOption) (*TaskResult, error) {
	out := new(TaskResult)
	err := c.cc.Invoke(ctx, "/google.devtools.remoteworkers.v1test2.Tasks/UpdateTaskResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksClient) AddTaskLog(ctx context.Context, in *AddTaskLogRequest, opts ...grpc.CallOption) (*AddTaskLogResponse, error) {
	out := new(AddTaskLogResponse)
	err := c.cc.Invoke(ctx, "/google.devtools.remoteworkers.v1test2.Tasks/AddTaskLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TasksServer is the server API for Tasks service.
type TasksServer interface {
	// DEPRECATED - use Lease.payload instead.
	// GetTask reads the current state of the task. Tasks must be created through
	// some other interface, and should be immutable once created and exposed to
	// the bots.
	GetTask(context.Context, *GetTaskRequest) (*Task, error)
	// DEPRECATED - use Lease.result instead.
	// UpdateTaskResult updates the result.
	UpdateTaskResult(context.Context, *UpdateTaskResultRequest) (*TaskResult, error)
	// DEPRECATED - precreate logs prior to sending to bot.
	// AddTaskLog creates a new streaming log. The log is streamed and marked as
	// completed through other interfaces (i.e., ByteStream). This can be called
	// by the bot if it wants to create a new log; the server can also predefine
	// logs that do not need to be created (e.g. `stdout`).
	AddTaskLog(context.Context, *AddTaskLogRequest) (*AddTaskLogResponse, error)
}

func RegisterTasksServer(s *grpc.Server, srv TasksServer) {
	s.RegisterService(&_Tasks_serviceDesc, srv)
}

func _Tasks_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.remoteworkers.v1test2.Tasks/GetTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).GetTask(ctx, req.(*GetTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tasks_UpdateTaskResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTaskResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).UpdateTaskResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.remoteworkers.v1test2.Tasks/UpdateTaskResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).UpdateTaskResult(ctx, req.(*UpdateTaskResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tasks_AddTaskLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTaskLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).AddTaskLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.remoteworkers.v1test2.Tasks/AddTaskLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).AddTaskLog(ctx, req.(*AddTaskLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tasks_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.devtools.remoteworkers.v1test2.Tasks",
	HandlerType: (*TasksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTask",
			Handler:    _Tasks_GetTask_Handler,
		},
		{
			MethodName: "UpdateTaskResult",
			Handler:    _Tasks_UpdateTaskResult_Handler,
		},
		{
			MethodName: "AddTaskLog",
			Handler:    _Tasks_AddTaskLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/devtools/remoteworkers/v1test2/tasks.proto",
}

func init() {
	proto.RegisterFile("google/devtools/remoteworkers/v1test2/tasks.proto", fileDescriptor_tasks_667ae2bf8892e2d4)
}

var fileDescriptor_tasks_667ae2bf8892e2d4 = []byte{
	// 682 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xcf, 0x6b, 0x13, 0x41,
	0x14, 0x66, 0x92, 0x34, 0x36, 0x2f, 0x20, 0x75, 0xa8, 0x36, 0x2e, 0x3d, 0xc4, 0xc5, 0x4a, 0x4c,
	0xcb, 0x2e, 0x89, 0xa8, 0x35, 0xc5, 0x42, 0x8b, 0x5a, 0x0a, 0x15, 0x64, 0xad, 0x16, 0xbc, 0x94,
	0x69, 0x76, 0x3a, 0x86, 0x6c, 0x76, 0xd6, 0x9d, 0xd9, 0x48, 0x90, 0x5e, 0x3c, 0x79, 0x15, 0xff,
	0x03, 0x8f, 0x5e, 0xbc, 0x7a, 0x14, 0xef, 0x9e, 0x04, 0xff, 0x02, 0xff, 0x10, 0x99, 0x1f, 0x69,
	0x9b, 0xfe, 0x88, 0xa9, 0xb7, 0x99, 0x79, 0xdf, 0xf7, 0xde, 0xfb, 0xde, 0xf7, 0xb2, 0x81, 0x06,
	0xe3, 0x9c, 0x45, 0xd4, 0x0f, 0x69, 0x5f, 0x72, 0x1e, 0x09, 0x3f, 0xa5, 0x3d, 0x2e, 0xe9, 0x5b,
	0x9e, 0x76, 0x69, 0x2a, 0xfc, 0x7e, 0x43, 0x52, 0x21, 0x9b, 0xbe, 0x24, 0xa2, 0x2b, 0xbc, 0x24,
	0xe5, 0x92, 0xe3, 0x05, 0x43, 0xf1, 0x86, 0x14, 0x6f, 0x84, 0xe2, 0x59, 0x8a, 0x33, 0x6f, 0x33,
	0x93, 0xa4, 0xe3, 0x93, 0x38, 0xe6, 0x92, 0xc8, 0x0e, 0x8f, 0x6d, 0x12, 0xe7, 0xba, 0x8d, 0xea,
	0xdb, 0x5e, 0xb6, 0xef, 0x93, 0x78, 0x60, 0x43, 0xd5, 0x93, 0xa1, 0xfd, 0x0e, 0x8d, 0xc2, 0xdd,
	0x1e, 0x11, 0x5d, 0x8b, 0x98, 0xb3, 0x88, 0x34, 0x69, 0xfb, 0x42, 0x12, 0x99, 0xd9, 0xac, 0xee,
	0x6f, 0x04, 0x85, 0x6d, 0x22, 0xba, 0x18, 0x43, 0x21, 0x26, 0x3d, 0x5a, 0x41, 0x55, 0x54, 0x2b,
	0x05, 0xfa, 0x8c, 0xef, 0x41, 0x39, 0xa4, 0xa2, 0x9d, 0x76, 0x12, 0xd5, 0x48, 0x25, 0x57, 0x45,
	0xb5, 0x72, 0x73, 0xd6, 0xb3, 0x6a, 0x86, 0xd5, 0xbc, 0xb5, 0x78, 0x10, 0x1c, 0x07, 0xe2, 0x4d,
	0x28, 0x44, 0x9c, 0x89, 0x4a, 0xbe, 0x9a, 0xaf, 0x95, 0x9b, 0x77, 0xbd, 0x89, 0xe4, 0x7b, 0xaa,
	0x0d, 0x6f, 0x8b, 0x33, 0xf1, 0x38, 0x96, 0xe9, 0x20, 0xd0, 0x29, 0x9c, 0xfb, 0x50, 0x3a, 0x7c,
	0xc2, 0x33, 0x90, 0xef, 0xd2, 0x81, 0x6d, 0x51, 0x1d, 0xf1, 0x2c, 0x4c, 0xf5, 0x49, 0x94, 0x51,
	0xdd, 0x5b, 0x29, 0x30, 0x97, 0x56, 0x6e, 0x19, 0xb9, 0xdf, 0x11, 0x80, 0xca, 0x18, 0x50, 0x91,
	0x45, 0xf2, 0x4c, 0x79, 0x0e, 0x4c, 0xb7, 0x79, 0x2f, 0x89, 0xa8, 0x34, 0xfc, 0xe9, 0xe0, 0xf0,
	0x8e, 0xeb, 0x50, 0x34, 0x73, 0xaa, 0xe4, 0xb5, 0x6a, 0x3c, 0x14, 0x91, 0x26, 0x6d, 0xef, 0xb9,
	0x8e, 0x04, 0x16, 0x81, 0x97, 0xa0, 0xc8, 0x33, 0x99, 0x64, 0xb2, 0x52, 0x18, 0x33, 0x21, 0x8b,
	0xc1, 0x35, 0x28, 0xf4, 0xa8, 0x24, 0x95, 0xa9, 0x31, 0x58, 0x8d, 0x70, 0x6f, 0xc2, 0xe5, 0x0d,
	0x2a, 0x8d, 0x88, 0x37, 0x19, 0x15, 0x67, 0xaa, 0x70, 0x7f, 0x22, 0x98, 0x7b, 0x91, 0x84, 0x44,
	0xd2, 0x23, 0xb9, 0x63, 0xf0, 0x78, 0x13, 0x8a, 0xa9, 0x06, 0x59, 0x3f, 0x1b, 0x17, 0xb0, 0xc7,
	0x66, 0xb7, 0x09, 0xf0, 0x0a, 0x94, 0x33, 0x5d, 0x59, 0xaf, 0x9a, 0x9d, 0x94, 0x73, 0x4a, 0xd1,
	0x13, 0xb5, 0x8d, 0x4f, 0x15, 0x1d, 0x0c, 0x5c, 0x9d, 0xf1, 0x35, 0x28, 0x0a, 0x9e, 0xa5, 0x6d,
	0xaa, 0xa7, 0x56, 0x0a, 0xec, 0xcd, 0x5d, 0x85, 0x2b, 0x6b, 0x61, 0xa8, 0xaa, 0x6d, 0x71, 0x36,
	0x4e, 0xc8, 0x55, 0x28, 0x46, 0x9c, 0xed, 0x76, 0xc2, 0xa1, 0xf9, 0x11, 0x67, 0x9b, 0xa1, 0xbb,
	0x04, 0xf8, 0x38, 0x5f, 0x24, 0x3c, 0x16, 0x54, 0x55, 0x7b, 0x4d, 0xe2, 0x30, 0x1a, 0xa6, 0xb0,
	0xb7, 0xe6, 0x87, 0x02, 0x4c, 0x29, 0xac, 0xc0, 0x1f, 0x11, 0x5c, 0xb2, 0xe3, 0xc6, 0x93, 0xae,
	0xec, 0xa8, 0x3d, 0xce, 0xe2, 0x05, 0x46, 0xe9, 0xba, 0xef, 0x7f, 0xfd, 0xf9, 0x94, 0x9b, 0xc7,
	0xce, 0xe1, 0x27, 0xe3, 0x9d, 0x92, 0xf5, 0xb0, 0x5e, 0x37, 0xdf, 0x0e, 0xbf, 0x7e, 0x80, 0xbf,
	0x21, 0x98, 0x39, 0xe9, 0x2d, 0x5e, 0x9d, 0xb0, 0xca, 0x39, 0x4b, 0xe1, 0x5c, 0xdc, 0x70, 0xb7,
	0xa1, 0x7b, 0x5d, 0x6c, 0xde, 0x38, 0xb7, 0x57, 0xdf, 0xac, 0xc4, 0x41, 0x6b, 0xb8, 0x1b, 0x5f,
	0x11, 0xc0, 0x91, 0x0f, 0x78, 0x79, 0xc2, 0xa2, 0xa7, 0xac, 0x77, 0x1e, 0xfc, 0x07, 0xd3, 0x98,
	0xee, 0x2e, 0xe9, 0xb6, 0x6f, 0xb9, 0xe7, 0xb7, 0x7d, 0xd0, 0x22, 0x61, 0xb8, 0xc5, 0x59, 0x0b,
	0xd5, 0xd7, 0x7f, 0x20, 0xb8, 0xdd, 0xe6, 0xbd, 0xc9, 0xca, 0xad, 0xe3, 0x40, 0x3f, 0xef, 0x98,
	0x67, 0xbd, 0x42, 0xcf, 0xd0, 0xab, 0xc0, 0x92, 0x19, 0x8f, 0x48, 0xcc, 0x3c, 0x9e, 0x32, 0x9f,
	0xd1, 0x58, 0xff, 0x12, 0x7c, 0x13, 0x22, 0x49, 0x47, 0xfc, 0xe3, 0xbf, 0x63, 0x65, 0xe4, 0xf5,
	0x73, 0x2e, 0x17, 0xec, 0x7c, 0xc9, 0x2d, 0x6c, 0x98, 0xcc, 0x8f, 0x68, 0x7f, 0x5b, 0xb7, 0x35,
	0x52, 0xdf, 0x7b, 0xd9, 0xd8, 0x56, 0xd4, 0xbd, 0xa2, 0xae, 0x75, 0xe7, 0x6f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xbb, 0x4a, 0x2c, 0x76, 0xa6, 0x06, 0x00, 0x00,
}
