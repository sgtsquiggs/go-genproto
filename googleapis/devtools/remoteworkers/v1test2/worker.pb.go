// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/remoteworkers/v1test2/worker.proto

package remoteworkers // import "github.com/sgtsquiggs/go-genproto/googleapis/devtools/remoteworkers/v1test2"

import proto "github.com/sgtsquiggs/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Describes a worker, which is a list of one or more devices and the
// connections between them. A device could be a computer, a phone, or even an
// accelerator like a GPU; it's up to the farm administrator to decide how to
// model their farm. For example, if a farm only has one type of GPU, the GPU
// could be modelled as a "has_gpu" property on its host computer; if it has
// many subproperties itself, it might be better to model it as a separate
// device.
//
// The first device in the worker is the "primary device" - that is, the device
// running a bot and which is responsible for actually executing commands. All
// other devices are considered to be attached devices, and must be controllable
// by the primary device.
//
// This message (and all its submessages) can be used in two contexts:
//
// * Status: sent by the bot to report the current capabilities of the device to
// allow reservation matching.
// * Request: sent by a client to request a device with certain capabilities in
// a reservation.
//
// Several of the fields in this message have different semantics depending on
// which of which of these contexts it is used. These semantics are described
// below.
//
// Several messages in Worker and its submessages have the concept of keys and
// values, such as `Worker.Property` and `Device.Property`. All keys are simple
// strings, but certain keys are "standard" keys and should be broadly supported
// across farms and implementations; these are listed below each relevant
// message. Bot implementations or farm admins may add *additional* keys, but
// these SHOULD all begin with an underscore so they do not conflict with
// standard keys that may be added in the future.
//
// Keys are not context sensitive.
//
// See http://goo.gl/NurY8g for more information on the Worker message.
type Worker struct {
	// A list of devices; the first device is the primary device. See the `Device`
	// message for more information.
	Devices []*Device `protobuf:"bytes,1,rep,name=devices,proto3" json:"devices,omitempty"`
	// A worker may contain "global" properties. For example, certain machines
	// might be reserved for certain types of jobs, like short-running compilation
	// versus long-running integration tests. This property is known as a "pool"
	// and is not related to any one device within the worker; rather, it applies
	// to the worker as a whole.
	//
	// The behaviour of repeated keys is identical to that of Device.Property.
	Properties []*Worker_Property `protobuf:"bytes,2,rep,name=properties,proto3" json:"properties,omitempty"`
	// Bots can be configured in certain ways when accepting leases. For example,
	// many leases are executed inside a Docker container. To support this, the
	// bot needs to be able to report that it has Docker installed (and knows how
	// to execute something inside a container), and the task submitter needs to
	// specify which image should be used to start the container. Similarly, a
	// lease may be able to run as one of several users on the worker; in such
	// cases, the bot needs to report what users are available, and the submitter
	// needs to choose one.
	//
	// Therefore, when this message is reported by the bot to the service, each
	// key represents a *type* of configuration that the bot knows how to set,
	// while each *value* represents a legal value for that configuration (the
	// empty string is interpretted as a wildcard, such as for Docker images).
	// When this message is sent by the server to the bot in the context of a
	// lease, it represents a command to the bot to apply the setting. Keys may
	// be repeated during reporting but not in a lease.
	Configs              []*Worker_Config `protobuf:"bytes,3,rep,name=configs,proto3" json:"configs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Worker) Reset()         { *m = Worker{} }
func (m *Worker) String() string { return proto.CompactTextString(m) }
func (*Worker) ProtoMessage()    {}
func (*Worker) Descriptor() ([]byte, []int) {
	return fileDescriptor_worker_52520c784d3d1f3d, []int{0}
}
func (m *Worker) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Worker.Unmarshal(m, b)
}
func (m *Worker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Worker.Marshal(b, m, deterministic)
}
func (dst *Worker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Worker.Merge(dst, src)
}
func (m *Worker) XXX_Size() int {
	return xxx_messageInfo_Worker.Size(m)
}
func (m *Worker) XXX_DiscardUnknown() {
	xxx_messageInfo_Worker.DiscardUnknown(m)
}

var xxx_messageInfo_Worker proto.InternalMessageInfo

func (m *Worker) GetDevices() []*Device {
	if m != nil {
		return m.Devices
	}
	return nil
}

func (m *Worker) GetProperties() []*Worker_Property {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *Worker) GetConfigs() []*Worker_Config {
	if m != nil {
		return m.Configs
	}
	return nil
}

// A global property; see the `properties` field for more information.
type Worker_Property struct {
	// For general information on keys, see the documentation to `Worker`.
	//
	// The current set of standard keys are:
	//
	// * pool: different workers can be reserved for different purposes. For
	// example, an admin might want to segregate long-running integration tests
	// from short-running unit tests, so unit tests will always get some
	// throughput. To support this, the server can assign different values for
	// `pool` (such as "itest" and "utest") to different workers, and then have
	// jobs request workers from those pools.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The property's value.
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Worker_Property) Reset()         { *m = Worker_Property{} }
func (m *Worker_Property) String() string { return proto.CompactTextString(m) }
func (*Worker_Property) ProtoMessage()    {}
func (*Worker_Property) Descriptor() ([]byte, []int) {
	return fileDescriptor_worker_52520c784d3d1f3d, []int{0, 0}
}
func (m *Worker_Property) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Worker_Property.Unmarshal(m, b)
}
func (m *Worker_Property) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Worker_Property.Marshal(b, m, deterministic)
}
func (dst *Worker_Property) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Worker_Property.Merge(dst, src)
}
func (m *Worker_Property) XXX_Size() int {
	return xxx_messageInfo_Worker_Property.Size(m)
}
func (m *Worker_Property) XXX_DiscardUnknown() {
	xxx_messageInfo_Worker_Property.DiscardUnknown(m)
}

var xxx_messageInfo_Worker_Property proto.InternalMessageInfo

func (m *Worker_Property) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Worker_Property) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// A configuration request or report; see the `configs` field for more
// information.
type Worker_Config struct {
	// For general information on keys, see the documentation to `Worker`.
	//
	// The current set of standard keys are:
	//
	// * DockerImage: the image of the container. When being reported by the
	// bot, the empty value should always be included if the bot is able to pull
	// its own images; the bot may optionally *also* report images that are
	// present in its cache. When being requested in a lease, the value is the
	// URI of the image (eg `gcr.io/user/image@sha256:hash`).
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The configuration's value.
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Worker_Config) Reset()         { *m = Worker_Config{} }
func (m *Worker_Config) String() string { return proto.CompactTextString(m) }
func (*Worker_Config) ProtoMessage()    {}
func (*Worker_Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_worker_52520c784d3d1f3d, []int{0, 1}
}
func (m *Worker_Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Worker_Config.Unmarshal(m, b)
}
func (m *Worker_Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Worker_Config.Marshal(b, m, deterministic)
}
func (dst *Worker_Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Worker_Config.Merge(dst, src)
}
func (m *Worker_Config) XXX_Size() int {
	return xxx_messageInfo_Worker_Config.Size(m)
}
func (m *Worker_Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Worker_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Worker_Config proto.InternalMessageInfo

func (m *Worker_Config) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Worker_Config) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Any device, including computers, phones, accelerators (e.g. GPUs), etc. All
// names must be unique.
type Device struct {
	// The handle can be thought of as the "name" of the device, and must be
	// unique within a Worker.
	//
	// In the Status context, the handle should be some human-understandable name,
	// perhaps corresponding to a label physically written on the device to make
	// it easy to locate. In the Request context, the name should be the
	// *logical* name expected by the task. The bot is responsible for mapping the
	// logical name expected by the task to a machine-readable name that the task
	// can actually use, such as a USB address. The method by which this mapping
	// is communicated to the task is not covered in this API.
	Handle string `protobuf:"bytes,1,opt,name=handle,proto3" json:"handle,omitempty"`
	// Properties of this device that don't change based on the tasks that are
	// running on it, e.g. OS, CPU architecture, etc.
	//
	// Keys may be repeated, and have the following interpretation:
	//
	//    * Status context: the device can support *any* the listed values. For
	//    example, an "ISA" property might include "x86", "x86-64" and "sse4".
	//
	//    * Request context: the device *must* support *all* of the listed values.
	Properties           []*Device_Property `protobuf:"bytes,2,rep,name=properties,proto3" json:"properties,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Device) Reset()         { *m = Device{} }
func (m *Device) String() string { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()    {}
func (*Device) Descriptor() ([]byte, []int) {
	return fileDescriptor_worker_52520c784d3d1f3d, []int{1}
}
func (m *Device) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Device.Unmarshal(m, b)
}
func (m *Device) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Device.Marshal(b, m, deterministic)
}
func (dst *Device) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Device.Merge(dst, src)
}
func (m *Device) XXX_Size() int {
	return xxx_messageInfo_Device.Size(m)
}
func (m *Device) XXX_DiscardUnknown() {
	xxx_messageInfo_Device.DiscardUnknown(m)
}

var xxx_messageInfo_Device proto.InternalMessageInfo

func (m *Device) GetHandle() string {
	if m != nil {
		return m.Handle
	}
	return ""
}

func (m *Device) GetProperties() []*Device_Property {
	if m != nil {
		return m.Properties
	}
	return nil
}

// A device property; see `properties` for more information.
type Device_Property struct {
	// For general information on keys, see the documentation to `Worker`.
	//
	// The current set of standard keys are:
	//
	// * os: a human-readable description of the OS. Examples include `linux`,
	// `ubuntu` and `ubuntu 14.04` (note that a bot may advertise itself as more
	// than one). This will be replaced in the future by more well-structured
	// keys and values to represent OS variants.
	//
	// * has-docker: "true" if the bot has Docker installed. This will be
	// replaced in the future by a more structured message for Docker support.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The property's value.
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Device_Property) Reset()         { *m = Device_Property{} }
func (m *Device_Property) String() string { return proto.CompactTextString(m) }
func (*Device_Property) ProtoMessage()    {}
func (*Device_Property) Descriptor() ([]byte, []int) {
	return fileDescriptor_worker_52520c784d3d1f3d, []int{1, 0}
}
func (m *Device_Property) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Device_Property.Unmarshal(m, b)
}
func (m *Device_Property) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Device_Property.Marshal(b, m, deterministic)
}
func (dst *Device_Property) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Device_Property.Merge(dst, src)
}
func (m *Device_Property) XXX_Size() int {
	return xxx_messageInfo_Device_Property.Size(m)
}
func (m *Device_Property) XXX_DiscardUnknown() {
	xxx_messageInfo_Device_Property.DiscardUnknown(m)
}

var xxx_messageInfo_Device_Property proto.InternalMessageInfo

func (m *Device_Property) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Device_Property) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Worker)(nil), "google.devtools.remoteworkers.v1test2.Worker")
	proto.RegisterType((*Worker_Property)(nil), "google.devtools.remoteworkers.v1test2.Worker.Property")
	proto.RegisterType((*Worker_Config)(nil), "google.devtools.remoteworkers.v1test2.Worker.Config")
	proto.RegisterType((*Device)(nil), "google.devtools.remoteworkers.v1test2.Device")
	proto.RegisterType((*Device_Property)(nil), "google.devtools.remoteworkers.v1test2.Device.Property")
}

func init() {
	proto.RegisterFile("google/devtools/remoteworkers/v1test2/worker.proto", fileDescriptor_worker_52520c784d3d1f3d)
}

var fileDescriptor_worker_52520c784d3d1f3d = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xc1, 0x4a, 0xf3, 0x40,
	0x10, 0xc7, 0x49, 0xca, 0x97, 0x7e, 0x8e, 0x17, 0x59, 0x45, 0x42, 0x4f, 0xa5, 0x50, 0xa8, 0x07,
	0x37, 0x36, 0x8a, 0x17, 0x6f, 0xb5, 0xd0, 0x9b, 0x94, 0xa5, 0xb4, 0xe0, 0x2d, 0xb6, 0xe3, 0x1a,
	0x9a, 0x66, 0xc2, 0x6e, 0x8c, 0xf4, 0x75, 0x3c, 0x8a, 0x6f, 0xe1, 0x83, 0xf8, 0x2a, 0xd2, 0xdd,
	0x04, 0x0c, 0x08, 0x46, 0x3d, 0x25, 0x33, 0xe1, 0xf7, 0x9b, 0xc9, 0x9f, 0x81, 0x50, 0x12, 0xc9,
	0x04, 0x83, 0x15, 0x16, 0x39, 0x51, 0xa2, 0x03, 0x85, 0x1b, 0xca, 0xf1, 0x89, 0xd4, 0x1a, 0x95,
	0x0e, 0x8a, 0x61, 0x8e, 0x3a, 0x0f, 0x03, 0x5b, 0xf3, 0x4c, 0x51, 0x4e, 0xac, 0x6f, 0x19, 0x5e,
	0x31, 0xbc, 0xc6, 0xf0, 0x92, 0xe9, 0xbd, 0xbb, 0xe0, 0x2d, 0x4c, 0x8f, 0x4d, 0xa0, 0xbd, 0xc2,
	0x22, 0x5e, 0xa2, 0xf6, 0x9d, 0x6e, 0x6b, 0xb0, 0x1f, 0x9e, 0xf2, 0x46, 0x0e, 0x3e, 0x36, 0x94,
	0xa8, 0x68, 0x36, 0x07, 0xc8, 0x14, 0x65, 0xa8, 0xf2, 0x18, 0xb5, 0xef, 0x1a, 0xd7, 0x65, 0x43,
	0x97, 0xdd, 0x85, 0x4f, 0x2d, 0xbf, 0x15, 0x9f, 0x4c, 0xec, 0x06, 0xda, 0x4b, 0x4a, 0xef, 0x63,
	0xa9, 0xfd, 0x96, 0x91, 0x5e, 0xfc, 0x4c, 0x7a, 0x6d, 0x60, 0x51, 0x49, 0x3a, 0x21, 0xfc, 0xaf,
	0xe6, 0xb0, 0x03, 0x68, 0xad, 0x71, 0xeb, 0x3b, 0x5d, 0x67, 0xb0, 0x27, 0x76, 0xaf, 0xec, 0x08,
	0xfe, 0x15, 0x51, 0xf2, 0x88, 0xbe, 0x6b, 0x7a, 0xb6, 0xe8, 0x9c, 0x81, 0x67, 0x35, 0x4d, 0x89,
	0xde, 0xab, 0x03, 0x9e, 0x4d, 0x88, 0x1d, 0x83, 0xf7, 0x10, 0xa5, 0xab, 0x04, 0x4b, 0xaa, 0xac,
	0xfe, 0x14, 0x98, 0x55, 0x7f, 0x19, 0xd8, 0x6f, 0x7e, 0x70, 0xf4, 0xe6, 0xc0, 0xc9, 0x92, 0x36,
	0xcd, 0xa6, 0x8f, 0x0e, 0x85, 0x69, 0xdb, 0x80, 0xb5, 0x7d, 0x4c, 0x9d, 0x5b, 0x51, 0xd2, 0x92,
	0x92, 0x28, 0x95, 0x9c, 0x94, 0x0c, 0x24, 0xa6, 0xe6, 0x30, 0x03, 0xfb, 0x29, 0xca, 0x62, 0xfd,
	0xcd, 0x3d, 0x5f, 0xd5, 0xba, 0xcf, 0xae, 0x2b, 0x16, 0x2f, 0x6e, 0x7f, 0x62, 0xcd, 0x63, 0x2c,
	0x66, 0x66, 0xaf, 0xda, 0x02, 0x7c, 0x3e, 0x9c, 0xed, 0xd0, 0x3b, 0xcf, 0xcc, 0x3a, 0xff, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x4e, 0x67, 0x96, 0xee, 0x3a, 0x03, 0x00, 0x00,
}
