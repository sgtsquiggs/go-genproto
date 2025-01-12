// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/text.proto

package automl // import "github.com/sgtsquiggs/go-genproto/googleapis/cloud/automl/v1beta1"

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

// Dataset metadata for classification.
type TextClassificationDatasetMetadata struct {
	// Required.
	// Type of the classification problem.
	ClassificationType   ClassificationType `protobuf:"varint,1,opt,name=classification_type,json=classificationType,proto3,enum=google.cloud.automl.v1beta1.ClassificationType" json:"classification_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TextClassificationDatasetMetadata) Reset()         { *m = TextClassificationDatasetMetadata{} }
func (m *TextClassificationDatasetMetadata) String() string { return proto.CompactTextString(m) }
func (*TextClassificationDatasetMetadata) ProtoMessage()    {}
func (*TextClassificationDatasetMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_text_1fb9d74955be1951, []int{0}
}
func (m *TextClassificationDatasetMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextClassificationDatasetMetadata.Unmarshal(m, b)
}
func (m *TextClassificationDatasetMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextClassificationDatasetMetadata.Marshal(b, m, deterministic)
}
func (dst *TextClassificationDatasetMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextClassificationDatasetMetadata.Merge(dst, src)
}
func (m *TextClassificationDatasetMetadata) XXX_Size() int {
	return xxx_messageInfo_TextClassificationDatasetMetadata.Size(m)
}
func (m *TextClassificationDatasetMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TextClassificationDatasetMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TextClassificationDatasetMetadata proto.InternalMessageInfo

func (m *TextClassificationDatasetMetadata) GetClassificationType() ClassificationType {
	if m != nil {
		return m.ClassificationType
	}
	return ClassificationType_CLASSIFICATION_TYPE_UNSPECIFIED
}

// Model metadata that is specific to text classification.
type TextClassificationModelMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextClassificationModelMetadata) Reset()         { *m = TextClassificationModelMetadata{} }
func (m *TextClassificationModelMetadata) String() string { return proto.CompactTextString(m) }
func (*TextClassificationModelMetadata) ProtoMessage()    {}
func (*TextClassificationModelMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_text_1fb9d74955be1951, []int{1}
}
func (m *TextClassificationModelMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextClassificationModelMetadata.Unmarshal(m, b)
}
func (m *TextClassificationModelMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextClassificationModelMetadata.Marshal(b, m, deterministic)
}
func (dst *TextClassificationModelMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextClassificationModelMetadata.Merge(dst, src)
}
func (m *TextClassificationModelMetadata) XXX_Size() int {
	return xxx_messageInfo_TextClassificationModelMetadata.Size(m)
}
func (m *TextClassificationModelMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TextClassificationModelMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TextClassificationModelMetadata proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TextClassificationDatasetMetadata)(nil), "google.cloud.automl.v1beta1.TextClassificationDatasetMetadata")
	proto.RegisterType((*TextClassificationModelMetadata)(nil), "google.cloud.automl.v1beta1.TextClassificationModelMetadata")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/text.proto", fileDescriptor_text_1fb9d74955be1951)
}

var fileDescriptor_text_1fb9d74955be1951 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x31, 0x4f, 0x03, 0x31,
	0x0c, 0x85, 0x75, 0x0b, 0x12, 0x19, 0x18, 0x8e, 0x05, 0x15, 0xa4, 0xd2, 0x0e, 0x88, 0x29, 0xa6,
	0x30, 0x32, 0x51, 0x58, 0x2b, 0x55, 0xa8, 0x13, 0x0b, 0xb8, 0x39, 0x13, 0x45, 0x4a, 0xe3, 0xd0,
	0xb8, 0xa8, 0xfd, 0x01, 0xfc, 0x6f, 0xd4, 0xe4, 0x18, 0x4e, 0x87, 0x6e, 0x4c, 0xfc, 0xbd, 0xe7,
	0xf7, 0xac, 0x6e, 0x2c, 0xb3, 0xf5, 0x04, 0xc6, 0xf3, 0xae, 0x01, 0xdc, 0x09, 0x6f, 0x3c, 0x7c,
	0xcf, 0xd6, 0x24, 0x38, 0x03, 0xa1, 0xbd, 0xe8, 0xb8, 0x65, 0xe1, 0xfa, 0xb2, 0x70, 0x3a, 0x73,
	0xba, 0x70, 0xba, 0xe5, 0x46, 0x57, 0xad, 0x09, 0x46, 0x07, 0x18, 0x02, 0x0b, 0x8a, 0xe3, 0x90,
	0x8a, 0x74, 0x74, 0x37, 0xb4, 0xc2, 0x78, 0x4c, 0xc9, 0x7d, 0x3a, 0x93, 0x25, 0x45, 0x31, 0xfd,
	0xa9, 0xd4, 0x64, 0x45, 0x7b, 0x79, 0xee, 0x0c, 0x5f, 0x50, 0x30, 0x91, 0x2c, 0x48, 0xb0, 0x41,
	0xc1, 0xfa, 0x43, 0x9d, 0x77, 0xd5, 0xef, 0x72, 0x88, 0x74, 0x51, 0x5d, 0x57, 0xb7, 0x67, 0xf7,
	0xa0, 0x07, 0x02, 0xeb, 0xae, 0xf1, 0xea, 0x10, 0xe9, 0xb5, 0x36, 0xbd, 0xbf, 0xe9, 0x44, 0x8d,
	0xfb, 0x31, 0x16, 0xdc, 0x90, 0xff, 0x0b, 0x31, 0xff, 0x52, 0x63, 0xc3, 0x9b, 0xa1, 0x65, 0xf3,
	0xd3, 0xa3, 0xc7, 0xf2, 0x58, 0x6c, 0x59, 0xbd, 0x3d, 0xb5, 0xa4, 0x65, 0x8f, 0xc1, 0x6a, 0xde,
	0x5a, 0xb0, 0x14, 0x72, 0x6d, 0x28, 0x23, 0x8c, 0x2e, 0xfd, 0x7b, 0xab, 0xc7, 0xf2, 0x5c, 0x9f,
	0x64, 0xfa, 0xe1, 0x37, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x8d, 0x0f, 0x20, 0xbb, 0x01, 0x00, 0x00,
}
