// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/ml/v1/operation_metadata.proto

package ml // import "github.com/sgtsquiggs/go-genproto/googleapis/cloud/ml/v1"

import proto "github.com/sgtsquiggs/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/sgtsquiggs/protobuf/ptypes/timestamp"
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

// The operation type.
type OperationMetadata_OperationType int32

const (
	// Unspecified operation type.
	OperationMetadata_OPERATION_TYPE_UNSPECIFIED OperationMetadata_OperationType = 0
	// An operation to create a new version.
	OperationMetadata_CREATE_VERSION OperationMetadata_OperationType = 1
	// An operation to delete an existing version.
	OperationMetadata_DELETE_VERSION OperationMetadata_OperationType = 2
	// An operation to delete an existing model.
	OperationMetadata_DELETE_MODEL OperationMetadata_OperationType = 3
)

var OperationMetadata_OperationType_name = map[int32]string{
	0: "OPERATION_TYPE_UNSPECIFIED",
	1: "CREATE_VERSION",
	2: "DELETE_VERSION",
	3: "DELETE_MODEL",
}
var OperationMetadata_OperationType_value = map[string]int32{
	"OPERATION_TYPE_UNSPECIFIED": 0,
	"CREATE_VERSION":             1,
	"DELETE_VERSION":             2,
	"DELETE_MODEL":               3,
}

func (x OperationMetadata_OperationType) String() string {
	return proto.EnumName(OperationMetadata_OperationType_name, int32(x))
}
func (OperationMetadata_OperationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_operation_metadata_67f21a738daf74c6, []int{0, 0}
}

// Represents the metadata of the long-running operation.
type OperationMetadata struct {
	// The time the operation was submitted.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The time operation processing started.
	StartTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// The time operation processing completed.
	EndTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Indicates whether a request to cancel this operation has been made.
	IsCancellationRequested bool `protobuf:"varint,4,opt,name=is_cancellation_requested,json=isCancellationRequested,proto3" json:"is_cancellation_requested,omitempty"`
	// The operation type.
	OperationType OperationMetadata_OperationType `protobuf:"varint,5,opt,name=operation_type,json=operationType,proto3,enum=google.cloud.ml.v1.OperationMetadata_OperationType" json:"operation_type,omitempty"`
	// Contains the name of the model associated with the operation.
	ModelName string `protobuf:"bytes,6,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"`
	// Contains the version associated with the operation.
	Version              *Version `protobuf:"bytes,7,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperationMetadata) Reset()         { *m = OperationMetadata{} }
func (m *OperationMetadata) String() string { return proto.CompactTextString(m) }
func (*OperationMetadata) ProtoMessage()    {}
func (*OperationMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_operation_metadata_67f21a738daf74c6, []int{0}
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

func (m *OperationMetadata) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *OperationMetadata) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *OperationMetadata) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *OperationMetadata) GetIsCancellationRequested() bool {
	if m != nil {
		return m.IsCancellationRequested
	}
	return false
}

func (m *OperationMetadata) GetOperationType() OperationMetadata_OperationType {
	if m != nil {
		return m.OperationType
	}
	return OperationMetadata_OPERATION_TYPE_UNSPECIFIED
}

func (m *OperationMetadata) GetModelName() string {
	if m != nil {
		return m.ModelName
	}
	return ""
}

func (m *OperationMetadata) GetVersion() *Version {
	if m != nil {
		return m.Version
	}
	return nil
}

func init() {
	proto.RegisterType((*OperationMetadata)(nil), "google.cloud.ml.v1.OperationMetadata")
	proto.RegisterEnum("google.cloud.ml.v1.OperationMetadata_OperationType", OperationMetadata_OperationType_name, OperationMetadata_OperationType_value)
}

func init() {
	proto.RegisterFile("google/cloud/ml/v1/operation_metadata.proto", fileDescriptor_operation_metadata_67f21a738daf74c6)
}

var fileDescriptor_operation_metadata_67f21a738daf74c6 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x5f, 0x6b, 0xdb, 0x30,
	0x14, 0xc5, 0xe7, 0xb6, 0x6b, 0x1a, 0x75, 0x0d, 0x99, 0x1e, 0xb6, 0xcc, 0xfb, 0x17, 0xfa, 0x30,
	0x02, 0x03, 0x99, 0xb4, 0xdb, 0xc3, 0xd6, 0xa7, 0x36, 0xd1, 0x20, 0xd0, 0xc6, 0xc6, 0xf5, 0x0a,
	0xeb, 0x8b, 0x51, 0xed, 0x3b, 0x23, 0x90, 0x25, 0xcf, 0x52, 0x0c, 0xfd, 0x2c, 0xfb, 0xb2, 0x23,
	0x92, 0x4d, 0x33, 0x52, 0xe8, 0xa3, 0xce, 0xfd, 0x9d, 0xab, 0xab, 0x7b, 0x84, 0x3e, 0x17, 0x4a,
	0x15, 0x02, 0x82, 0x4c, 0xa8, 0x55, 0x1e, 0x94, 0x22, 0x68, 0xa6, 0x81, 0xaa, 0xa0, 0x66, 0x86,
	0x2b, 0x99, 0x96, 0x60, 0x58, 0xce, 0x0c, 0x23, 0x55, 0xad, 0x8c, 0xc2, 0xd8, 0xc1, 0xc4, 0xc2,
	0xa4, 0x14, 0xa4, 0x99, 0xfa, 0xef, 0xda, 0x06, 0xac, 0xe2, 0x01, 0x93, 0x52, 0x19, 0xeb, 0xd4,
	0xce, 0xe1, 0x7f, 0x7a, 0xa4, 0x7d, 0xa9, 0x72, 0x10, 0xa9, 0x86, 0xba, 0xe1, 0x19, 0xb4, 0xdc,
	0xc7, 0x96, 0xb3, 0xa7, 0xbb, 0xd5, 0xef, 0xc0, 0xf0, 0x12, 0xb4, 0x61, 0x65, 0xe5, 0x80, 0xe3,
	0xbf, 0x7b, 0xe8, 0x65, 0xd8, 0xcd, 0x75, 0xd5, 0x8e, 0x85, 0xcf, 0xd0, 0x61, 0x56, 0x03, 0x33,
	0x90, 0xae, 0xf9, 0x91, 0x37, 0xf6, 0x26, 0x87, 0x27, 0x3e, 0x69, 0xc7, 0xec, 0x9a, 0x91, 0xa4,
	0x6b, 0x16, 0x23, 0x87, 0xaf, 0x05, 0xfc, 0x0d, 0x21, 0x6d, 0x58, 0x6d, 0x9c, 0x77, 0xe7, 0x49,
	0x6f, 0xdf, 0xd2, 0xd6, 0xfa, 0x15, 0x1d, 0x80, 0xcc, 0x9d, 0x71, 0xf7, 0x49, 0x63, 0x0f, 0x64,
	0x6e, 0x6d, 0xdf, 0xd1, 0x1b, 0xae, 0xd3, 0x8c, 0xc9, 0x0c, 0x84, 0x70, 0x1b, 0xae, 0xe1, 0xcf,
	0x0a, 0xb4, 0x81, 0x7c, 0xb4, 0x37, 0xf6, 0x26, 0x07, 0xf1, 0x6b, 0xae, 0x67, 0x1b, 0xf5, 0xb8,
	0x2b, 0xe3, 0x5b, 0x34, 0x78, 0xc8, 0xc5, 0xdc, 0x57, 0x30, 0x7a, 0x3e, 0xf6, 0x26, 0x83, 0x93,
	0x53, 0xb2, 0x1d, 0x0a, 0xd9, 0xda, 0xd4, 0x83, 0x92, 0xdc, 0x57, 0x10, 0x1f, 0xa9, 0xcd, 0x23,
	0x7e, 0x8f, 0x90, 0x0b, 0x45, 0xb2, 0x12, 0x46, 0xfb, 0x63, 0x6f, 0xd2, 0x8f, 0xfb, 0x56, 0x59,
	0x32, 0xfb, 0xda, 0x5e, 0x03, 0xb5, 0xe6, 0x4a, 0x8e, 0x7a, 0xf6, 0xb1, 0x6f, 0x1f, 0xbb, 0xf3,
	0xc6, 0x21, 0x71, 0xc7, 0x1e, 0x73, 0x74, 0xf4, 0xdf, 0xad, 0xf8, 0x03, 0xf2, 0xc3, 0x88, 0xc6,
	0xe7, 0xc9, 0x22, 0x5c, 0xa6, 0xc9, 0xaf, 0x88, 0xa6, 0x3f, 0x97, 0xd7, 0x11, 0x9d, 0x2d, 0x7e,
	0x2c, 0xe8, 0x7c, 0xf8, 0x0c, 0x63, 0x34, 0x98, 0xc5, 0xf4, 0x3c, 0xa1, 0xe9, 0x0d, 0x8d, 0xaf,
	0x17, 0xe1, 0x72, 0xe8, 0xad, 0xb5, 0x39, 0xbd, 0xa4, 0x1b, 0xda, 0x0e, 0x1e, 0xa2, 0x17, 0xad,
	0x76, 0x15, 0xce, 0xe9, 0xe5, 0x70, 0xf7, 0x42, 0x20, 0x3f, 0x53, 0xe5, 0xd6, 0x54, 0xac, 0xe2,
	0xa4, 0x99, 0x5e, 0xbc, 0xda, 0x5a, 0x47, 0xb4, 0x0e, 0x29, 0xf2, 0x6e, 0xbf, 0xb4, 0x8e, 0x42,
	0x09, 0x26, 0x0b, 0xa2, 0xea, 0x22, 0x28, 0x40, 0xda, 0x08, 0x03, 0x57, 0x62, 0x15, 0xd7, 0x9b,
	0xbf, 0xf7, 0xac, 0x14, 0x77, 0xfb, 0x16, 0x38, 0xfd, 0x17, 0x00, 0x00, 0xff, 0xff, 0x03, 0xf9,
	0xcc, 0xf1, 0x3c, 0x03, 0x00, 0x00,
}
