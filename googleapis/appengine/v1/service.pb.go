// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/appengine/v1/service.proto

package appengine // import "github.com/sgtsquiggs/go-genproto/googleapis/appengine/v1"

import proto "github.com/sgtsquiggs/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/sgtsquiggs/go-genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Available sharding mechanisms.
type TrafficSplit_ShardBy int32

const (
	// Diversion method unspecified.
	TrafficSplit_UNSPECIFIED TrafficSplit_ShardBy = 0
	// Diversion based on a specially named cookie, "GOOGAPPUID." The cookie
	// must be set by the application itself or no diversion will occur.
	TrafficSplit_COOKIE TrafficSplit_ShardBy = 1
	// Diversion based on applying the modulus operation to a fingerprint
	// of the IP address.
	TrafficSplit_IP TrafficSplit_ShardBy = 2
)

var TrafficSplit_ShardBy_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "COOKIE",
	2: "IP",
}
var TrafficSplit_ShardBy_value = map[string]int32{
	"UNSPECIFIED": 0,
	"COOKIE":      1,
	"IP":          2,
}

func (x TrafficSplit_ShardBy) String() string {
	return proto.EnumName(TrafficSplit_ShardBy_name, int32(x))
}
func (TrafficSplit_ShardBy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_service_c3b4fd6114e9d0a5, []int{1, 0}
}

// A Service resource is a logical component of an application that can share
// state and communicate in a secure fashion with other services.
// For example, an application that handles customer requests might
// include separate services to handle tasks such as backend data
// analysis or API requests from mobile devices. Each service has a
// collection of versions that define a specific set of code used to
// implement the functionality of that service.
type Service struct {
	// Full path to the Service resource in the API.
	// Example: `apps/myapp/services/default`.
	//
	// @OutputOnly
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Relative name of the service within the application.
	// Example: `default`.
	//
	// @OutputOnly
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Mapping that defines fractional HTTP traffic diversion to
	// different versions within the service.
	Split                *TrafficSplit `protobuf:"bytes,3,opt,name=split,proto3" json:"split,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_c3b4fd6114e9d0a5, []int{0}
}
func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (dst *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(dst, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Service) GetSplit() *TrafficSplit {
	if m != nil {
		return m.Split
	}
	return nil
}

// Traffic routing configuration for versions within a single service. Traffic
// splits define how traffic directed to the service is assigned to versions.
type TrafficSplit struct {
	// Mechanism used to determine which version a request is sent to.
	// The traffic selection algorithm will
	// be stable for either type until allocations are changed.
	ShardBy TrafficSplit_ShardBy `protobuf:"varint,1,opt,name=shard_by,json=shardBy,proto3,enum=google.appengine.v1.TrafficSplit_ShardBy" json:"shard_by,omitempty"`
	// Mapping from version IDs within the service to fractional
	// (0.000, 1] allocations of traffic for that version. Each version can
	// be specified only once, but some versions in the service may not
	// have any traffic allocation. Services that have traffic allocated
	// cannot be deleted until either the service is deleted or
	// their traffic allocation is removed. Allocations must sum to 1.
	// Up to two decimal place precision is supported for IP-based splits and
	// up to three decimal places is supported for cookie-based splits.
	Allocations          map[string]float64 `protobuf:"bytes,2,rep,name=allocations,proto3" json:"allocations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TrafficSplit) Reset()         { *m = TrafficSplit{} }
func (m *TrafficSplit) String() string { return proto.CompactTextString(m) }
func (*TrafficSplit) ProtoMessage()    {}
func (*TrafficSplit) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_c3b4fd6114e9d0a5, []int{1}
}
func (m *TrafficSplit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficSplit.Unmarshal(m, b)
}
func (m *TrafficSplit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficSplit.Marshal(b, m, deterministic)
}
func (dst *TrafficSplit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficSplit.Merge(dst, src)
}
func (m *TrafficSplit) XXX_Size() int {
	return xxx_messageInfo_TrafficSplit.Size(m)
}
func (m *TrafficSplit) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficSplit.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficSplit proto.InternalMessageInfo

func (m *TrafficSplit) GetShardBy() TrafficSplit_ShardBy {
	if m != nil {
		return m.ShardBy
	}
	return TrafficSplit_UNSPECIFIED
}

func (m *TrafficSplit) GetAllocations() map[string]float64 {
	if m != nil {
		return m.Allocations
	}
	return nil
}

func init() {
	proto.RegisterType((*Service)(nil), "google.appengine.v1.Service")
	proto.RegisterType((*TrafficSplit)(nil), "google.appengine.v1.TrafficSplit")
	proto.RegisterMapType((map[string]float64)(nil), "google.appengine.v1.TrafficSplit.AllocationsEntry")
	proto.RegisterEnum("google.appengine.v1.TrafficSplit_ShardBy", TrafficSplit_ShardBy_name, TrafficSplit_ShardBy_value)
}

func init() {
	proto.RegisterFile("google/appengine/v1/service.proto", fileDescriptor_service_c3b4fd6114e9d0a5)
}

var fileDescriptor_service_c3b4fd6114e9d0a5 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x51, 0x8b, 0xda, 0x40,
	0x10, 0xc7, 0x9b, 0xb5, 0x6a, 0x3b, 0x11, 0x1b, 0xb6, 0x85, 0x06, 0xe9, 0x83, 0xfa, 0x64, 0x5f,
	0x36, 0x98, 0x3e, 0xb4, 0xb4, 0xa5, 0x50, 0x35, 0x85, 0x50, 0x38, 0x43, 0xe2, 0xbd, 0xdc, 0xcb,
	0xb1, 0xc6, 0x75, 0x6f, 0xb9, 0xb8, 0x1b, 0x92, 0x5c, 0x20, 0x5f, 0xe5, 0x3e, 0xed, 0x91, 0x6c,
	0xf0, 0xbc, 0x43, 0xf0, 0x6d, 0x66, 0xf2, 0xfb, 0xff, 0x67, 0x32, 0xb3, 0x30, 0xe1, 0x4a, 0xf1,
	0x84, 0x39, 0x34, 0x4d, 0x99, 0xe4, 0x42, 0x32, 0xa7, 0x9c, 0x3b, 0x39, 0xcb, 0x4a, 0x11, 0x33,
	0x92, 0x66, 0xaa, 0x50, 0xf8, 0xa3, 0x46, 0xc8, 0x11, 0x21, 0xe5, 0x7c, 0xf4, 0xe5, 0xa8, 0x13,
	0x0e, 0x95, 0x52, 0x15, 0xb4, 0x10, 0x4a, 0xe6, 0x5a, 0x32, 0xdd, 0x43, 0x3f, 0xd2, 0x1e, 0x18,
	0xc3, 0x5b, 0x49, 0x0f, 0xcc, 0x36, 0xc6, 0xc6, 0xec, 0x7d, 0xd8, 0xc4, 0x78, 0x08, 0x48, 0xec,
	0x6c, 0xd4, 0x54, 0x90, 0xd8, 0xe1, 0xef, 0xd0, 0xcd, 0xd3, 0x44, 0x14, 0x76, 0x67, 0x6c, 0xcc,
	0x4c, 0x77, 0x42, 0xce, 0x74, 0x24, 0x9b, 0x8c, 0xee, 0xf7, 0x22, 0x8e, 0x6a, 0x30, 0xd4, 0xfc,
	0xf4, 0x11, 0xc1, 0xe0, 0xb4, 0x8e, 0x57, 0xf0, 0x2e, 0xbf, 0xa3, 0xd9, 0xee, 0x76, 0x5b, 0x35,
	0x1d, 0x87, 0xee, 0xd7, 0x8b, 0x66, 0x24, 0xaa, 0x15, 0x8b, 0x2a, 0xec, 0xe7, 0x3a, 0xc0, 0x1b,
	0x30, 0x69, 0x92, 0xa8, 0x58, 0xff, 0x93, 0x8d, 0xc6, 0x9d, 0x99, 0xe9, 0xba, 0x97, 0x8d, 0xfe,
	0x3e, 0x8b, 0x3c, 0x59, 0x64, 0x55, 0x78, 0x6a, 0x33, 0xfa, 0x03, 0xd6, 0x6b, 0x00, 0x5b, 0xd0,
	0xb9, 0x67, 0x55, 0xbb, 0x9c, 0x3a, 0xc4, 0x9f, 0xa0, 0x5b, 0xd2, 0xe4, 0x81, 0x35, 0xeb, 0x31,
	0x42, 0x9d, 0xfc, 0x44, 0x3f, 0x8c, 0x29, 0x81, 0x7e, 0x3b, 0x29, 0xfe, 0x00, 0xe6, 0xf5, 0x55,
	0x14, 0x78, 0x4b, 0xff, 0x9f, 0xef, 0xad, 0xac, 0x37, 0x18, 0xa0, 0xb7, 0x5c, 0xaf, 0xff, 0xfb,
	0x9e, 0x65, 0xe0, 0x1e, 0x20, 0x3f, 0xb0, 0xd0, 0x82, 0xc3, 0xe7, 0x58, 0x1d, 0xce, 0x4d, 0xbd,
	0x18, 0xb4, 0xd7, 0x09, 0xea, 0x6b, 0x05, 0xc6, 0xcd, 0xef, 0x16, 0xe2, 0x2a, 0xa1, 0x92, 0x13,
	0x95, 0x71, 0x87, 0x33, 0xd9, 0xdc, 0xd2, 0xd1, 0x9f, 0x68, 0x2a, 0xf2, 0x17, 0x8f, 0xe4, 0xd7,
	0x31, 0xd9, 0xf6, 0x1a, 0xf0, 0xdb, 0x53, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5f, 0xbc, 0x13, 0xf4,
	0x4c, 0x02, 0x00, 0x00,
}
