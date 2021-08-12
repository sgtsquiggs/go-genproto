// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/clouderrorreporting/v1beta1/report_errors_service.proto

package clouderrorreporting // import "github.com/sgtsquiggs/go-genproto/googleapis/devtools/clouderrorreporting/v1beta1"

import proto "github.com/sgtsquiggs/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/sgtsquiggs/protobuf/ptypes/timestamp"
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

// A request for reporting an individual error event.
type ReportErrorEventRequest struct {
	// [Required] The resource name of the Google Cloud Platform project. Written
	// as `projects/` plus the
	// [Google Cloud Platform project ID](https://support.google.com/cloud/answer/6158840).
	// Example: `projects/my-project-123`.
	ProjectName string `protobuf:"bytes,1,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	// [Required] The error event to be reported.
	Event                *ReportedErrorEvent `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ReportErrorEventRequest) Reset()         { *m = ReportErrorEventRequest{} }
func (m *ReportErrorEventRequest) String() string { return proto.CompactTextString(m) }
func (*ReportErrorEventRequest) ProtoMessage()    {}
func (*ReportErrorEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_report_errors_service_7b4a77afec65ac92, []int{0}
}
func (m *ReportErrorEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportErrorEventRequest.Unmarshal(m, b)
}
func (m *ReportErrorEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportErrorEventRequest.Marshal(b, m, deterministic)
}
func (dst *ReportErrorEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportErrorEventRequest.Merge(dst, src)
}
func (m *ReportErrorEventRequest) XXX_Size() int {
	return xxx_messageInfo_ReportErrorEventRequest.Size(m)
}
func (m *ReportErrorEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportErrorEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReportErrorEventRequest proto.InternalMessageInfo

func (m *ReportErrorEventRequest) GetProjectName() string {
	if m != nil {
		return m.ProjectName
	}
	return ""
}

func (m *ReportErrorEventRequest) GetEvent() *ReportedErrorEvent {
	if m != nil {
		return m.Event
	}
	return nil
}

// Response for reporting an individual error event.
// Data may be added to this message in the future.
type ReportErrorEventResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReportErrorEventResponse) Reset()         { *m = ReportErrorEventResponse{} }
func (m *ReportErrorEventResponse) String() string { return proto.CompactTextString(m) }
func (*ReportErrorEventResponse) ProtoMessage()    {}
func (*ReportErrorEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_report_errors_service_7b4a77afec65ac92, []int{1}
}
func (m *ReportErrorEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportErrorEventResponse.Unmarshal(m, b)
}
func (m *ReportErrorEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportErrorEventResponse.Marshal(b, m, deterministic)
}
func (dst *ReportErrorEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportErrorEventResponse.Merge(dst, src)
}
func (m *ReportErrorEventResponse) XXX_Size() int {
	return xxx_messageInfo_ReportErrorEventResponse.Size(m)
}
func (m *ReportErrorEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportErrorEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReportErrorEventResponse proto.InternalMessageInfo

// An error event which is reported to the Error Reporting system.
type ReportedErrorEvent struct {
	// [Optional] Time when the event occurred.
	// If not provided, the time when the event was received by the
	// Error Reporting system will be used.
	EventTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	// [Required] The service context in which this error has occurred.
	ServiceContext *ServiceContext `protobuf:"bytes,2,opt,name=service_context,json=serviceContext,proto3" json:"service_context,omitempty"`
	// [Required] A message describing the error. The message can contain an
	// exception stack in one of the supported programming languages and formats.
	// In that case, the message is parsed and detailed exception information
	// is returned when retrieving the error event again.
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	// [Optional] A description of the context in which the error occurred.
	Context              *ErrorContext `protobuf:"bytes,4,opt,name=context,proto3" json:"context,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ReportedErrorEvent) Reset()         { *m = ReportedErrorEvent{} }
func (m *ReportedErrorEvent) String() string { return proto.CompactTextString(m) }
func (*ReportedErrorEvent) ProtoMessage()    {}
func (*ReportedErrorEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_report_errors_service_7b4a77afec65ac92, []int{2}
}
func (m *ReportedErrorEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportedErrorEvent.Unmarshal(m, b)
}
func (m *ReportedErrorEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportedErrorEvent.Marshal(b, m, deterministic)
}
func (dst *ReportedErrorEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportedErrorEvent.Merge(dst, src)
}
func (m *ReportedErrorEvent) XXX_Size() int {
	return xxx_messageInfo_ReportedErrorEvent.Size(m)
}
func (m *ReportedErrorEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportedErrorEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ReportedErrorEvent proto.InternalMessageInfo

func (m *ReportedErrorEvent) GetEventTime() *timestamp.Timestamp {
	if m != nil {
		return m.EventTime
	}
	return nil
}

func (m *ReportedErrorEvent) GetServiceContext() *ServiceContext {
	if m != nil {
		return m.ServiceContext
	}
	return nil
}

func (m *ReportedErrorEvent) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ReportedErrorEvent) GetContext() *ErrorContext {
	if m != nil {
		return m.Context
	}
	return nil
}

func init() {
	proto.RegisterType((*ReportErrorEventRequest)(nil), "google.devtools.clouderrorreporting.v1beta1.ReportErrorEventRequest")
	proto.RegisterType((*ReportErrorEventResponse)(nil), "google.devtools.clouderrorreporting.v1beta1.ReportErrorEventResponse")
	proto.RegisterType((*ReportedErrorEvent)(nil), "google.devtools.clouderrorreporting.v1beta1.ReportedErrorEvent")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ReportErrorsServiceClient is the client API for ReportErrorsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/github.com/sgtsquiggs/grpc-go#ClientConn.NewStream.
type ReportErrorsServiceClient interface {
	// Report an individual error event.
	//
	// This endpoint accepts <strong>either</strong> an OAuth token,
	// <strong>or</strong> an
	// <a href="https://support.google.com/cloud/answer/6158862">API key</a>
	// for authentication. To use an API key, append it to the URL as the value of
	// a `key` parameter. For example:
	// <pre>POST https://clouderrorreporting.googleapis.com/v1beta1/projects/example-project/events:report?key=123ABC456</pre>
	ReportErrorEvent(ctx context.Context, in *ReportErrorEventRequest, opts ...grpc.CallOption) (*ReportErrorEventResponse, error)
}

type reportErrorsServiceClient struct {
	cc *grpc.ClientConn
}

func NewReportErrorsServiceClient(cc *grpc.ClientConn) ReportErrorsServiceClient {
	return &reportErrorsServiceClient{cc}
}

func (c *reportErrorsServiceClient) ReportErrorEvent(ctx context.Context, in *ReportErrorEventRequest, opts ...grpc.CallOption) (*ReportErrorEventResponse, error) {
	out := new(ReportErrorEventResponse)
	err := c.cc.Invoke(ctx, "/google.devtools.clouderrorreporting.v1beta1.ReportErrorsService/ReportErrorEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReportErrorsServiceServer is the server API for ReportErrorsService service.
type ReportErrorsServiceServer interface {
	// Report an individual error event.
	//
	// This endpoint accepts <strong>either</strong> an OAuth token,
	// <strong>or</strong> an
	// <a href="https://support.google.com/cloud/answer/6158862">API key</a>
	// for authentication. To use an API key, append it to the URL as the value of
	// a `key` parameter. For example:
	// <pre>POST https://clouderrorreporting.googleapis.com/v1beta1/projects/example-project/events:report?key=123ABC456</pre>
	ReportErrorEvent(context.Context, *ReportErrorEventRequest) (*ReportErrorEventResponse, error)
}

func RegisterReportErrorsServiceServer(s *grpc.Server, srv ReportErrorsServiceServer) {
	s.RegisterService(&_ReportErrorsService_serviceDesc, srv)
}

func _ReportErrorsService_ReportErrorEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportErrorEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportErrorsServiceServer).ReportErrorEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.clouderrorreporting.v1beta1.ReportErrorsService/ReportErrorEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportErrorsServiceServer).ReportErrorEvent(ctx, req.(*ReportErrorEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReportErrorsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.devtools.clouderrorreporting.v1beta1.ReportErrorsService",
	HandlerType: (*ReportErrorsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportErrorEvent",
			Handler:    _ReportErrorsService_ReportErrorEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/devtools/clouderrorreporting/v1beta1/report_errors_service.proto",
}

func init() {
	proto.RegisterFile("google/devtools/clouderrorreporting/v1beta1/report_errors_service.proto", fileDescriptor_report_errors_service_7b4a77afec65ac92)
}

var fileDescriptor_report_errors_service_7b4a77afec65ac92 = []byte{
	// 490 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xcd, 0x8a, 0x13, 0x41,
	0x10, 0xc7, 0x99, 0xf8, 0xb1, 0x6c, 0x47, 0x54, 0xda, 0x83, 0xc3, 0x20, 0xb8, 0xc6, 0xcb, 0xa2,
	0x30, 0x6d, 0xe2, 0xc5, 0xec, 0x22, 0x0b, 0x59, 0xc3, 0xde, 0x64, 0x99, 0xd5, 0x3d, 0x48, 0x70,
	0xe8, 0x4c, 0xca, 0x61, 0x24, 0xd3, 0x35, 0x76, 0x77, 0x82, 0x20, 0x5e, 0x7c, 0x85, 0x7d, 0x05,
	0x4f, 0x3e, 0x8a, 0x57, 0x5f, 0xc0, 0x83, 0x0f, 0xa1, 0x37, 0xe9, 0xaf, 0x25, 0x6b, 0x72, 0x70,
	0xf4, 0x58, 0xd3, 0x55, 0xbf, 0xff, 0xbf, 0x3e, 0x86, 0x1c, 0x95, 0x88, 0xe5, 0x1c, 0xd8, 0x0c,
	0x96, 0x1a, 0x71, 0xae, 0x58, 0x31, 0xc7, 0xc5, 0x0c, 0xa4, 0x44, 0x29, 0xa1, 0x41, 0xa9, 0x2b,
	0x51, 0xb2, 0x65, 0x7f, 0x0a, 0x9a, 0xf7, 0x99, 0xfb, 0x92, 0xdb, 0x57, 0x95, 0x2b, 0x90, 0xcb,
	0xaa, 0x80, 0xb4, 0x91, 0xa8, 0x91, 0x3e, 0x74, 0xa0, 0x34, 0x80, 0xd2, 0x0d, 0xa0, 0xd4, 0x83,
	0x92, 0x3b, 0x5e, 0x95, 0x37, 0x15, 0xe3, 0x42, 0xa0, 0xe6, 0xba, 0x42, 0xa1, 0x1c, 0x2a, 0x79,
	0xd2, 0xc6, 0x53, 0x81, 0x75, 0x8d, 0xc2, 0x57, 0xde, 0xf5, 0x95, 0x36, 0x9a, 0x2e, 0xde, 0x30,
	0x5d, 0xd5, 0xa0, 0x34, 0xaf, 0x1b, 0x97, 0xd0, 0x3b, 0x8b, 0xc8, 0xed, 0xcc, 0x32, 0xc6, 0x06,
	0x37, 0x5e, 0x82, 0xd0, 0x19, 0xbc, 0x5b, 0x80, 0xd2, 0xf4, 0x1e, 0xb9, 0xd6, 0x48, 0x7c, 0x0b,
	0x85, 0xce, 0x05, 0xaf, 0x21, 0x8e, 0x76, 0xa2, 0xdd, 0xed, 0xac, 0xeb, 0xbf, 0x3d, 0xe7, 0x35,
	0xd0, 0x97, 0xe4, 0x0a, 0x98, 0x92, 0xb8, 0xb3, 0x13, 0xed, 0x76, 0x07, 0x07, 0x69, 0x8b, 0xa6,
	0x53, 0xa7, 0x0b, 0xb3, 0x15, 0x65, 0x47, 0xeb, 0x25, 0x24, 0x5e, 0x37, 0xa5, 0x1a, 0x14, 0x0a,
	0x7a, 0x9f, 0x3b, 0x84, 0xae, 0x57, 0xd2, 0x21, 0x21, 0xb6, 0x36, 0x37, 0x1d, 0x5a, 0xab, 0xdd,
	0x41, 0x12, 0xec, 0x84, 0xf6, 0xd3, 0x17, 0xa1, 0xfd, 0x6c, 0xdb, 0x66, 0x9b, 0x98, 0xce, 0xc8,
	0x0d, 0xbf, 0xba, 0xbc, 0x40, 0xa1, 0xe1, 0x7d, 0x68, 0x67, 0xbf, 0x55, 0x3b, 0x27, 0x8e, 0x71,
	0xe8, 0x10, 0xd9, 0x75, 0x75, 0x21, 0xa6, 0x31, 0xd9, 0xaa, 0x41, 0x29, 0x5e, 0x42, 0x7c, 0xc9,
	0x0e, 0x32, 0x84, 0xf4, 0x84, 0x6c, 0x05, 0xdd, 0xcb, 0x56, 0x77, 0xd8, 0x4a, 0xd7, 0x0e, 0x21,
	0xa8, 0x06, 0xd2, 0xe0, 0x67, 0x44, 0x6e, 0xad, 0xcc, 0x50, 0x79, 0x77, 0xf4, 0x7b, 0x44, 0x6e,
	0xfe, 0x39, 0x5b, 0xfa, 0xec, 0x1f, 0xf6, 0xb6, 0x76, 0x2f, 0xc9, 0xf8, 0x3f, 0x29, 0x7e, 0xc1,
	0x07, 0x9f, 0xbe, 0xfd, 0x38, 0xeb, 0x0c, 0x7b, 0x8f, 0xce, 0x4f, 0xfa, 0xc3, 0xea, 0x19, 0x3e,
	0xf5, 0x81, 0x62, 0x0f, 0x3e, 0x32, 0xbb, 0x44, 0xb5, 0xe7, 0xe8, 0x7b, 0xee, 0x7a, 0x46, 0xbf,
	0x22, 0x62, 0xfe, 0x82, 0x36, 0x6e, 0x46, 0xf1, 0x86, 0x59, 0x1d, 0x9b, 0xab, 0x39, 0x8e, 0x5e,
	0xbd, 0xf6, 0xa0, 0x12, 0xe7, 0x5c, 0x94, 0x29, 0xca, 0x92, 0x95, 0x20, 0xec, 0x4d, 0x31, 0xf7,
	0xc4, 0x9b, 0x4a, 0xfd, 0xd5, 0xdf, 0xb9, 0xbf, 0xe1, 0xed, 0x4b, 0xe7, 0xfe, 0x91, 0x13, 0x38,
	0x34, 0x8f, 0x6e, 0x9f, 0xd9, 0xb9, 0xc3, 0xd3, 0xfe, 0xc8, 0x54, 0x7e, 0x0d, 0x59, 0x13, 0x9b,
	0x35, 0xb9, 0x98, 0x35, 0x39, 0x75, 0xfc, 0xe9, 0x55, 0x6b, 0xeb, 0xf1, 0xef, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x2c, 0xd1, 0x8e, 0x76, 0xc7, 0x04, 0x00, 0x00,
}
