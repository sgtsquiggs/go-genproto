// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/httpbody.proto

package httpbody // import "github.com/sgtsquiggs/go-genproto/googleapis/api/httpbody"

import proto "github.com/sgtsquiggs/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/sgtsquiggs/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Message that represents an arbitrary HTTP body. It should only be used for
// payload formats that can't be represented as JSON, such as raw binary or
// an HTML page.
//
//
// This message can be used both in streaming and non-streaming API methods in
// the request as well as the response.
//
// It can be used as a top-level request field, which is convenient if one
// wants to extract parameters from either the URL or HTTP template into the
// request fields and also want access to the raw HTTP body.
//
// Example:
//
//     message GetResourceRequest {
//       // A unique request id.
//       string request_id = 1;
//
//       // The raw HTTP body is bound to this field.
//       google.api.HttpBody http_body = 2;
//     }
//
//     service ResourceService {
//       rpc GetResource(GetResourceRequest) returns (google.api.HttpBody);
//       rpc UpdateResource(google.api.HttpBody) returns (google.protobuf.Empty);
//     }
//
// Example with streaming methods:
//
//     service CaldavService {
//       rpc GetCalendar(stream google.api.HttpBody)
//         returns (stream google.api.HttpBody);
//       rpc UpdateCalendar(stream google.api.HttpBody)
//         returns (stream google.api.HttpBody);
//     }
//
// Use of this type only changes how the request and response bodies are
// handled, all other features will continue to work unchanged.
type HttpBody struct {
	// The HTTP Content-Type string representing the content type of the body.
	ContentType string `protobuf:"bytes,1,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	// HTTP body binary data.
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	// Application specific response metadata. Must be set in the first response
	// for streaming APIs.
	Extensions           []*any.Any `protobuf:"bytes,3,rep,name=extensions,proto3" json:"extensions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *HttpBody) Reset()         { *m = HttpBody{} }
func (m *HttpBody) String() string { return proto.CompactTextString(m) }
func (*HttpBody) ProtoMessage()    {}
func (*HttpBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_httpbody_d3ca4c70025a4501, []int{0}
}
func (m *HttpBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpBody.Unmarshal(m, b)
}
func (m *HttpBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpBody.Marshal(b, m, deterministic)
}
func (dst *HttpBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpBody.Merge(dst, src)
}
func (m *HttpBody) XXX_Size() int {
	return xxx_messageInfo_HttpBody.Size(m)
}
func (m *HttpBody) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpBody.DiscardUnknown(m)
}

var xxx_messageInfo_HttpBody proto.InternalMessageInfo

func (m *HttpBody) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *HttpBody) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *HttpBody) GetExtensions() []*any.Any {
	if m != nil {
		return m.Extensions
	}
	return nil
}

func init() {
	proto.RegisterType((*HttpBody)(nil), "google.api.HttpBody")
}

func init() { proto.RegisterFile("google/api/httpbody.proto", fileDescriptor_httpbody_d3ca4c70025a4501) }

var fileDescriptor_httpbody_d3ca4c70025a4501 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x95, 0xb6, 0x42, 0x70, 0x2d, 0x0c, 0x16, 0x43, 0x60, 0x0a, 0x4c, 0x99, 0x6c, 0x09,
	0xd8, 0x3a, 0x35, 0x0b, 0xb0, 0x45, 0x11, 0x13, 0x0b, 0x72, 0x9a, 0xc3, 0x44, 0x2a, 0x77, 0xa7,
	0xe6, 0x10, 0xf8, 0xeb, 0xf0, 0x49, 0x11, 0xf9, 0x23, 0xe8, 0xf6, 0xe4, 0xdf, 0xcf, 0x7e, 0xcf,
	0x70, 0x11, 0x98, 0xc3, 0x0e, 0x9d, 0x97, 0xd6, 0xbd, 0xa9, 0x4a, 0xcd, 0x4d, 0xb4, 0xb2, 0x67,
	0x65, 0x03, 0x03, 0xb2, 0x5e, 0xda, 0xcb, 0x49, 0xeb, 0x49, 0xfd, 0xf1, 0xea, 0x3c, 0x8d, 0xda,
	0xf5, 0x27, 0x1c, 0x3f, 0xa8, 0x4a, 0xc1, 0x4d, 0x34, 0x57, 0xb0, 0xda, 0x32, 0x29, 0x92, 0xbe,
	0x68, 0x14, 0x4c, 0x93, 0x2c, 0xc9, 0x4f, 0xaa, 0xe5, 0x78, 0xf6, 0x14, 0x05, 0x8d, 0x81, 0x45,
	0xe3, 0xd5, 0xa7, 0xb3, 0x2c, 0xc9, 0x57, 0x55, 0x9f, 0xcd, 0x1d, 0x00, 0x7e, 0x29, 0x52, 0xd7,
	0x32, 0x75, 0xe9, 0x3c, 0x9b, 0xe7, 0xcb, 0x9b, 0x73, 0x3b, 0xd6, 0x4f, 0x95, 0x76, 0x43, 0xb1,
	0xfa, 0xe7, 0x15, 0x08, 0x67, 0x5b, 0x7e, 0xb7, 0x7f, 0x2b, 0x8b, 0xd3, 0x69, 0x48, 0xf9, 0x7b,
	0xa7, 0x4c, 0x9e, 0xd7, 0x23, 0x0c, 0xbc, 0xf3, 0x14, 0x2c, 0xef, 0x83, 0x0b, 0x48, 0xfd, 0x8b,
	0x6e, 0x40, 0x5e, 0xda, 0xee, 0xe0, 0xf3, 0xeb, 0x29, 0x7c, 0xcf, 0x16, 0xf7, 0x9b, 0xf2, 0xb1,
	0x3e, 0xea, 0xf5, 0xdb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x11, 0x70, 0x02, 0x2a, 0x01,
	0x00, 0x00,
}
