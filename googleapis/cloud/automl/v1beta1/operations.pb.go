// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/operations.proto

package automl // import "github.com/sgtsquiggs/go-genproto/googleapis/cloud/automl/v1beta1"

import proto "github.com/sgtsquiggs/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/sgtsquiggs/protobuf/ptypes/empty"
import timestamp "github.com/sgtsquiggs/protobuf/ptypes/timestamp"
import _ "github.com/sgtsquiggs/go-genproto/googleapis/api/annotations"
import status "github.com/sgtsquiggs/go-genproto/googleapis/rpc/status"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Metadata used across all long running operations returned by AutoML API.
type OperationMetadata struct {
	// Ouptut only. Details of specific operation. Even if this field is empty,
	// the presence allows to distinguish different types of operations.
	//
	// Types that are valid to be assigned to Details:
	//	*OperationMetadata_CreateModelDetails
	Details isOperationMetadata_Details `protobuf_oneof:"details"`
	// Output only. Progress of operation. Range: [0, 100].
	ProgressPercent int32 `protobuf:"varint,13,opt,name=progress_percent,json=progressPercent,proto3" json:"progress_percent,omitempty"`
	// Output only. Partial failures encountered.
	// E.g. single files that couldn't be read.
	// This field should never exceed 20 entries.
	// Status details field will contain standard GCP error details.
	PartialFailures []*status.Status `protobuf:"bytes,2,rep,name=partial_failures,json=partialFailures,proto3" json:"partial_failures,omitempty"`
	// Output only. Time when the operation was created.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Time when the operation was updated for the last time.
	UpdateTime           *timestamp.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *OperationMetadata) Reset()         { *m = OperationMetadata{} }
func (m *OperationMetadata) String() string { return proto.CompactTextString(m) }
func (*OperationMetadata) ProtoMessage()    {}
func (*OperationMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_operations_8f7a4667c678e484, []int{0}
}
func (m *OperationMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperationMetadata.Unmarshal(m, b)
}
func (m *OperationMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperationMetadata.Marshal(b, m, deterministic)
}
func (dst *OperationMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperationMetadata.Merge(dst, src)
}
func (m *OperationMetadata) XXX_Size() int {
	return xxx_messageInfo_OperationMetadata.Size(m)
}
func (m *OperationMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_OperationMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_OperationMetadata proto.InternalMessageInfo

type isOperationMetadata_Details interface {
	isOperationMetadata_Details()
}

type OperationMetadata_CreateModelDetails struct {
	CreateModelDetails *CreateModelOperationMetadata `protobuf:"bytes,10,opt,name=create_model_details,json=createModelDetails,proto3,oneof"`
}

func (*OperationMetadata_CreateModelDetails) isOperationMetadata_Details() {}

func (m *OperationMetadata) GetDetails() isOperationMetadata_Details {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *OperationMetadata) GetCreateModelDetails() *CreateModelOperationMetadata {
	if x, ok := m.GetDetails().(*OperationMetadata_CreateModelDetails); ok {
		return x.CreateModelDetails
	}
	return nil
}

func (m *OperationMetadata) GetProgressPercent() int32 {
	if m != nil {
		return m.ProgressPercent
	}
	return 0
}

func (m *OperationMetadata) GetPartialFailures() []*status.Status {
	if m != nil {
		return m.PartialFailures
	}
	return nil
}

func (m *OperationMetadata) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *OperationMetadata) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*OperationMetadata) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _OperationMetadata_OneofMarshaler, _OperationMetadata_OneofUnmarshaler, _OperationMetadata_OneofSizer, []interface{}{
		(*OperationMetadata_CreateModelDetails)(nil),
	}
}

func _OperationMetadata_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*OperationMetadata)
	// details
	switch x := m.Details.(type) {
	case *OperationMetadata_CreateModelDetails:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CreateModelDetails); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("OperationMetadata.Details has unexpected type %T", x)
	}
	return nil
}

func _OperationMetadata_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*OperationMetadata)
	switch tag {
	case 10: // details.create_model_details
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CreateModelOperationMetadata)
		err := b.DecodeMessage(msg)
		m.Details = &OperationMetadata_CreateModelDetails{msg}
		return true, err
	default:
		return false, nil
	}
}

func _OperationMetadata_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*OperationMetadata)
	// details
	switch x := m.Details.(type) {
	case *OperationMetadata_CreateModelDetails:
		s := proto.Size(x.CreateModelDetails)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Details of CreateModel operation.
type CreateModelOperationMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateModelOperationMetadata) Reset()         { *m = CreateModelOperationMetadata{} }
func (m *CreateModelOperationMetadata) String() string { return proto.CompactTextString(m) }
func (*CreateModelOperationMetadata) ProtoMessage()    {}
func (*CreateModelOperationMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_operations_8f7a4667c678e484, []int{1}
}
func (m *CreateModelOperationMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateModelOperationMetadata.Unmarshal(m, b)
}
func (m *CreateModelOperationMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateModelOperationMetadata.Marshal(b, m, deterministic)
}
func (dst *CreateModelOperationMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateModelOperationMetadata.Merge(dst, src)
}
func (m *CreateModelOperationMetadata) XXX_Size() int {
	return xxx_messageInfo_CreateModelOperationMetadata.Size(m)
}
func (m *CreateModelOperationMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateModelOperationMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_CreateModelOperationMetadata proto.InternalMessageInfo

func init() {
	proto.RegisterType((*OperationMetadata)(nil), "google.cloud.automl.v1beta1.OperationMetadata")
	proto.RegisterType((*CreateModelOperationMetadata)(nil), "google.cloud.automl.v1beta1.CreateModelOperationMetadata")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/operations.proto", fileDescriptor_operations_8f7a4667c678e484)
}

var fileDescriptor_operations_8f7a4667c678e484 = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4f, 0xab, 0xd3, 0x40,
	0x14, 0xc5, 0x6d, 0xeb, 0x1f, 0x9c, 0x22, 0xad, 0x41, 0x30, 0xd4, 0x62, 0x4b, 0x37, 0x56, 0x90,
	0x19, 0x5a, 0x57, 0x52, 0x5c, 0x58, 0x45, 0xdc, 0x14, 0x4b, 0x74, 0xe5, 0x26, 0xdc, 0x26, 0xb7,
	0x21, 0x30, 0xc9, 0x0c, 0x33, 0x37, 0x05, 0x3f, 0xb6, 0xdf, 0x40, 0x32, 0x33, 0xd1, 0xc7, 0x7b,
	0x8f, 0x74, 0x39, 0xf7, 0xfc, 0xce, 0xcd, 0xb9, 0x87, 0xb0, 0x77, 0x85, 0x52, 0x85, 0x44, 0x91,
	0x49, 0xd5, 0xe4, 0x02, 0x1a, 0x52, 0x95, 0x14, 0x97, 0xcd, 0x09, 0x09, 0x36, 0x42, 0x69, 0x34,
	0x40, 0xa5, 0xaa, 0x2d, 0xd7, 0x46, 0x91, 0x8a, 0x5e, 0x79, 0x9a, 0x3b, 0x9a, 0x7b, 0x9a, 0x07,
	0x7a, 0x36, 0x0f, 0xab, 0x40, 0x97, 0x02, 0xea, 0x5a, 0xd1, 0x4d, 0xeb, 0xec, 0x4d, 0xdf, 0x87,
	0x2a, 0x95, 0xa3, 0x0c, 0xe0, 0xf6, 0x2a, 0x98, 0xe2, 0x05, 0x64, 0xe3, 0xb6, 0x07, 0x4f, 0xc8,
	0x25, 0xdc, 0xeb, 0xd4, 0x9c, 0x05, 0x56, 0x9a, 0x7e, 0x07, 0x71, 0x71, 0x5b, 0xa4, 0xb2, 0x42,
	0x4b, 0x50, 0xe9, 0x00, 0xbc, 0x0c, 0x80, 0xd1, 0x99, 0xb0, 0x04, 0xd4, 0x84, 0xcc, 0xab, 0x3f,
	0x43, 0xf6, 0xfc, 0x7b, 0xd7, 0xc1, 0x01, 0x09, 0x72, 0x20, 0x88, 0x2a, 0xf6, 0x22, 0x33, 0x08,
	0x84, 0xa9, 0x4f, 0x93, 0x23, 0x41, 0x29, 0x6d, 0xcc, 0x96, 0x83, 0xf5, 0x78, 0xfb, 0x81, 0xf7,
	0x74, 0xc4, 0x3f, 0x3b, 0xe3, 0xa1, 0xf5, 0xdd, 0x59, 0xfc, 0xed, 0x41, 0x12, 0x65, 0xff, 0xf5,
	0x2f, 0x7e, 0x6d, 0xf4, 0x96, 0x4d, 0xb5, 0x51, 0x85, 0x41, 0x6b, 0x53, 0x8d, 0x26, 0xc3, 0x9a,
	0xe2, 0x67, 0xcb, 0xc1, 0xfa, 0x51, 0x32, 0xe9, 0xe6, 0x47, 0x3f, 0x8e, 0x3e, 0xb2, 0xa9, 0x06,
	0x43, 0x25, 0xc8, 0xf4, 0x0c, 0xa5, 0x6c, 0x0c, 0xda, 0x78, 0xb8, 0x1c, 0xad, 0xc7, 0xdb, 0xa8,
	0x4b, 0x65, 0x74, 0xc6, 0x7f, 0xb8, 0x1b, 0x93, 0x49, 0x60, 0xbf, 0x06, 0x34, 0xda, 0xb1, 0x71,
	0x38, 0xac, 0x6d, 0x28, 0x1e, 0xb9, 0x7b, 0x66, 0x9d, 0xb3, 0xab, 0x8f, 0xff, 0xec, 0xea, 0x4b,
	0x98, 0xc7, 0xdb, 0x41, 0x6b, 0x6e, 0x74, 0xfe, 0xcf, 0xfc, 0xf0, 0xba, 0xd9, 0xe3, 0xed, 0x60,
	0xff, 0x94, 0x3d, 0x09, 0x2d, 0xae, 0x5e, 0xb3, 0x79, 0x5f, 0x49, 0xfb, 0x33, 0x5b, 0x64, 0xaa,
	0xea, 0x2b, 0xf9, 0x38, 0xf8, 0xf5, 0x29, 0xc8, 0x85, 0x92, 0x50, 0x17, 0x5c, 0x99, 0x42, 0x14,
	0x58, 0xbb, 0x10, 0xc2, 0x4b, 0xa0, 0x4b, 0x7b, 0xef, 0x2f, 0xb6, 0xf3, 0xcf, 0xd3, 0x63, 0x47,
	0xbf, 0xff, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x7b, 0xa7, 0x51, 0x21, 0x03, 0x00, 0x00,
}
