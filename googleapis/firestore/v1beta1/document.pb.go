// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/firestore/v1beta1/document.proto

package firestore // import "github.com/sgtsquiggs/go-genproto/googleapis/firestore/v1beta1"

import proto "github.com/sgtsquiggs/protobuf/proto"
import fmt "fmt"
import math "math"
import _struct "github.com/sgtsquiggs/protobuf/ptypes/struct"
import timestamp "github.com/sgtsquiggs/protobuf/ptypes/timestamp"
import _ "github.com/sgtsquiggs/go-genproto/googleapis/api/annotations"
import latlng "github.com/sgtsquiggs/go-genproto/googleapis/type/latlng"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A Firestore document.
//
// Must not exceed 1 MiB - 4 bytes.
type Document struct {
	// The resource name of the document, for example
	// `projects/{project_id}/databases/{database_id}/documents/{document_path}`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The document's fields.
	//
	// The map keys represent field names.
	//
	// A simple field name contains only characters `a` to `z`, `A` to `Z`,
	// `0` to `9`, or `_`, and must not start with `0` to `9`. For example,
	// `foo_bar_17`.
	//
	// Field names matching the regular expression `__.*__` are reserved. Reserved
	// field names are forbidden except in certain documented contexts. The map
	// keys, represented as UTF-8, must not exceed 1,500 bytes and cannot be
	// empty.
	//
	// Field paths may be used in other contexts to refer to structured fields
	// defined here. For `map_value`, the field path is represented by the simple
	// or quoted field names of the containing fields, delimited by `.`. For
	// example, the structured field
	// `"foo" : { map_value: { "x&y" : { string_value: "hello" }}}` would be
	// represented by the field path `foo.x&y`.
	//
	// Within a field path, a quoted field name starts and ends with `` ` `` and
	// may contain any character. Some characters, including `` ` ``, must be
	// escaped using a `\`. For example, `` `x&y` `` represents `x&y` and
	// `` `bak\`tik` `` represents `` bak`tik ``.
	Fields map[string]*Value `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Output only. The time at which the document was created.
	//
	// This value increases monotonically when a document is deleted then
	// recreated. It can also be compared to values from other documents and
	// the `read_time` of a query.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The time at which the document was last changed.
	//
	// This value is initially set to the `create_time` then increases
	// monotonically with each change to the document. It can also be
	// compared to values from other documents and the `read_time` of a query.
	UpdateTime           *timestamp.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Document) Reset()         { *m = Document{} }
func (m *Document) String() string { return proto.CompactTextString(m) }
func (*Document) ProtoMessage()    {}
func (*Document) Descriptor() ([]byte, []int) {
	return fileDescriptor_document_fe46ba4c8da20ac8, []int{0}
}
func (m *Document) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Document.Unmarshal(m, b)
}
func (m *Document) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Document.Marshal(b, m, deterministic)
}
func (dst *Document) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Document.Merge(dst, src)
}
func (m *Document) XXX_Size() int {
	return xxx_messageInfo_Document.Size(m)
}
func (m *Document) XXX_DiscardUnknown() {
	xxx_messageInfo_Document.DiscardUnknown(m)
}

var xxx_messageInfo_Document proto.InternalMessageInfo

func (m *Document) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Document) GetFields() map[string]*Value {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *Document) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Document) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

// A message that can hold any of the supported value types.
type Value struct {
	// Must have a value set.
	//
	// Types that are valid to be assigned to ValueType:
	//	*Value_NullValue
	//	*Value_BooleanValue
	//	*Value_IntegerValue
	//	*Value_DoubleValue
	//	*Value_TimestampValue
	//	*Value_StringValue
	//	*Value_BytesValue
	//	*Value_ReferenceValue
	//	*Value_GeoPointValue
	//	*Value_ArrayValue
	//	*Value_MapValue
	ValueType            isValue_ValueType `protobuf_oneof:"value_type"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_document_fe46ba4c8da20ac8, []int{1}
}
func (m *Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value.Unmarshal(m, b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value.Marshal(b, m, deterministic)
}
func (dst *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(dst, src)
}
func (m *Value) XXX_Size() int {
	return xxx_messageInfo_Value.Size(m)
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

type isValue_ValueType interface {
	isValue_ValueType()
}

type Value_NullValue struct {
	NullValue _struct.NullValue `protobuf:"varint,11,opt,name=null_value,json=nullValue,proto3,enum=google.protobuf.NullValue,oneof"`
}

type Value_BooleanValue struct {
	BooleanValue bool `protobuf:"varint,1,opt,name=boolean_value,json=booleanValue,proto3,oneof"`
}

type Value_IntegerValue struct {
	IntegerValue int64 `protobuf:"varint,2,opt,name=integer_value,json=integerValue,proto3,oneof"`
}

type Value_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,3,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type Value_TimestampValue struct {
	TimestampValue *timestamp.Timestamp `protobuf:"bytes,10,opt,name=timestamp_value,json=timestampValue,proto3,oneof"`
}

type Value_StringValue struct {
	StringValue string `protobuf:"bytes,17,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type Value_BytesValue struct {
	BytesValue []byte `protobuf:"bytes,18,opt,name=bytes_value,json=bytesValue,proto3,oneof"`
}

type Value_ReferenceValue struct {
	ReferenceValue string `protobuf:"bytes,5,opt,name=reference_value,json=referenceValue,proto3,oneof"`
}

type Value_GeoPointValue struct {
	GeoPointValue *latlng.LatLng `protobuf:"bytes,8,opt,name=geo_point_value,json=geoPointValue,proto3,oneof"`
}

type Value_ArrayValue struct {
	ArrayValue *ArrayValue `protobuf:"bytes,9,opt,name=array_value,json=arrayValue,proto3,oneof"`
}

type Value_MapValue struct {
	MapValue *MapValue `protobuf:"bytes,6,opt,name=map_value,json=mapValue,proto3,oneof"`
}

func (*Value_NullValue) isValue_ValueType() {}

func (*Value_BooleanValue) isValue_ValueType() {}

func (*Value_IntegerValue) isValue_ValueType() {}

func (*Value_DoubleValue) isValue_ValueType() {}

func (*Value_TimestampValue) isValue_ValueType() {}

func (*Value_StringValue) isValue_ValueType() {}

func (*Value_BytesValue) isValue_ValueType() {}

func (*Value_ReferenceValue) isValue_ValueType() {}

func (*Value_GeoPointValue) isValue_ValueType() {}

func (*Value_ArrayValue) isValue_ValueType() {}

func (*Value_MapValue) isValue_ValueType() {}

func (m *Value) GetValueType() isValue_ValueType {
	if m != nil {
		return m.ValueType
	}
	return nil
}

func (m *Value) GetNullValue() _struct.NullValue {
	if x, ok := m.GetValueType().(*Value_NullValue); ok {
		return x.NullValue
	}
	return _struct.NullValue_NULL_VALUE
}

func (m *Value) GetBooleanValue() bool {
	if x, ok := m.GetValueType().(*Value_BooleanValue); ok {
		return x.BooleanValue
	}
	return false
}

func (m *Value) GetIntegerValue() int64 {
	if x, ok := m.GetValueType().(*Value_IntegerValue); ok {
		return x.IntegerValue
	}
	return 0
}

func (m *Value) GetDoubleValue() float64 {
	if x, ok := m.GetValueType().(*Value_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (m *Value) GetTimestampValue() *timestamp.Timestamp {
	if x, ok := m.GetValueType().(*Value_TimestampValue); ok {
		return x.TimestampValue
	}
	return nil
}

func (m *Value) GetStringValue() string {
	if x, ok := m.GetValueType().(*Value_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *Value) GetBytesValue() []byte {
	if x, ok := m.GetValueType().(*Value_BytesValue); ok {
		return x.BytesValue
	}
	return nil
}

func (m *Value) GetReferenceValue() string {
	if x, ok := m.GetValueType().(*Value_ReferenceValue); ok {
		return x.ReferenceValue
	}
	return ""
}

func (m *Value) GetGeoPointValue() *latlng.LatLng {
	if x, ok := m.GetValueType().(*Value_GeoPointValue); ok {
		return x.GeoPointValue
	}
	return nil
}

func (m *Value) GetArrayValue() *ArrayValue {
	if x, ok := m.GetValueType().(*Value_ArrayValue); ok {
		return x.ArrayValue
	}
	return nil
}

func (m *Value) GetMapValue() *MapValue {
	if x, ok := m.GetValueType().(*Value_MapValue); ok {
		return x.MapValue
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Value) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Value_OneofMarshaler, _Value_OneofUnmarshaler, _Value_OneofSizer, []interface{}{
		(*Value_NullValue)(nil),
		(*Value_BooleanValue)(nil),
		(*Value_IntegerValue)(nil),
		(*Value_DoubleValue)(nil),
		(*Value_TimestampValue)(nil),
		(*Value_StringValue)(nil),
		(*Value_BytesValue)(nil),
		(*Value_ReferenceValue)(nil),
		(*Value_GeoPointValue)(nil),
		(*Value_ArrayValue)(nil),
		(*Value_MapValue)(nil),
	}
}

func _Value_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Value)
	// value_type
	switch x := m.ValueType.(type) {
	case *Value_NullValue:
		b.EncodeVarint(11<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.NullValue))
	case *Value_BooleanValue:
		t := uint64(0)
		if x.BooleanValue {
			t = 1
		}
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *Value_IntegerValue:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.IntegerValue))
	case *Value_DoubleValue:
		b.EncodeVarint(3<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.DoubleValue))
	case *Value_TimestampValue:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TimestampValue); err != nil {
			return err
		}
	case *Value_StringValue:
		b.EncodeVarint(17<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.StringValue)
	case *Value_BytesValue:
		b.EncodeVarint(18<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.BytesValue)
	case *Value_ReferenceValue:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.ReferenceValue)
	case *Value_GeoPointValue:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.GeoPointValue); err != nil {
			return err
		}
	case *Value_ArrayValue:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ArrayValue); err != nil {
			return err
		}
	case *Value_MapValue:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.MapValue); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Value.ValueType has unexpected type %T", x)
	}
	return nil
}

func _Value_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Value)
	switch tag {
	case 11: // value_type.null_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ValueType = &Value_NullValue{_struct.NullValue(x)}
		return true, err
	case 1: // value_type.boolean_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ValueType = &Value_BooleanValue{x != 0}
		return true, err
	case 2: // value_type.integer_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ValueType = &Value_IntegerValue{int64(x)}
		return true, err
	case 3: // value_type.double_value
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.ValueType = &Value_DoubleValue{math.Float64frombits(x)}
		return true, err
	case 10: // value_type.timestamp_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(timestamp.Timestamp)
		err := b.DecodeMessage(msg)
		m.ValueType = &Value_TimestampValue{msg}
		return true, err
	case 17: // value_type.string_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.ValueType = &Value_StringValue{x}
		return true, err
	case 18: // value_type.bytes_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.ValueType = &Value_BytesValue{x}
		return true, err
	case 5: // value_type.reference_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.ValueType = &Value_ReferenceValue{x}
		return true, err
	case 8: // value_type.geo_point_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(latlng.LatLng)
		err := b.DecodeMessage(msg)
		m.ValueType = &Value_GeoPointValue{msg}
		return true, err
	case 9: // value_type.array_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ArrayValue)
		err := b.DecodeMessage(msg)
		m.ValueType = &Value_ArrayValue{msg}
		return true, err
	case 6: // value_type.map_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(MapValue)
		err := b.DecodeMessage(msg)
		m.ValueType = &Value_MapValue{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Value_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Value)
	// value_type
	switch x := m.ValueType.(type) {
	case *Value_NullValue:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.NullValue))
	case *Value_BooleanValue:
		n += 1 // tag and wire
		n += 1
	case *Value_IntegerValue:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.IntegerValue))
	case *Value_DoubleValue:
		n += 1 // tag and wire
		n += 8
	case *Value_TimestampValue:
		s := proto.Size(x.TimestampValue)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Value_StringValue:
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(len(x.StringValue)))
		n += len(x.StringValue)
	case *Value_BytesValue:
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(len(x.BytesValue)))
		n += len(x.BytesValue)
	case *Value_ReferenceValue:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.ReferenceValue)))
		n += len(x.ReferenceValue)
	case *Value_GeoPointValue:
		s := proto.Size(x.GeoPointValue)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Value_ArrayValue:
		s := proto.Size(x.ArrayValue)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Value_MapValue:
		s := proto.Size(x.MapValue)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// An array value.
type ArrayValue struct {
	// Values in the array.
	Values               []*Value `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArrayValue) Reset()         { *m = ArrayValue{} }
func (m *ArrayValue) String() string { return proto.CompactTextString(m) }
func (*ArrayValue) ProtoMessage()    {}
func (*ArrayValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_document_fe46ba4c8da20ac8, []int{2}
}
func (m *ArrayValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArrayValue.Unmarshal(m, b)
}
func (m *ArrayValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArrayValue.Marshal(b, m, deterministic)
}
func (dst *ArrayValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArrayValue.Merge(dst, src)
}
func (m *ArrayValue) XXX_Size() int {
	return xxx_messageInfo_ArrayValue.Size(m)
}
func (m *ArrayValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ArrayValue.DiscardUnknown(m)
}

var xxx_messageInfo_ArrayValue proto.InternalMessageInfo

func (m *ArrayValue) GetValues() []*Value {
	if m != nil {
		return m.Values
	}
	return nil
}

// A map value.
type MapValue struct {
	// The map's fields.
	//
	// The map keys represent field names. Field names matching the regular
	// expression `__.*__` are reserved. Reserved field names are forbidden except
	// in certain documented contexts. The map keys, represented as UTF-8, must
	// not exceed 1,500 bytes and cannot be empty.
	Fields               map[string]*Value `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MapValue) Reset()         { *m = MapValue{} }
func (m *MapValue) String() string { return proto.CompactTextString(m) }
func (*MapValue) ProtoMessage()    {}
func (*MapValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_document_fe46ba4c8da20ac8, []int{3}
}
func (m *MapValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MapValue.Unmarshal(m, b)
}
func (m *MapValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MapValue.Marshal(b, m, deterministic)
}
func (dst *MapValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MapValue.Merge(dst, src)
}
func (m *MapValue) XXX_Size() int {
	return xxx_messageInfo_MapValue.Size(m)
}
func (m *MapValue) XXX_DiscardUnknown() {
	xxx_messageInfo_MapValue.DiscardUnknown(m)
}

var xxx_messageInfo_MapValue proto.InternalMessageInfo

func (m *MapValue) GetFields() map[string]*Value {
	if m != nil {
		return m.Fields
	}
	return nil
}

func init() {
	proto.RegisterType((*Document)(nil), "google.firestore.v1beta1.Document")
	proto.RegisterMapType((map[string]*Value)(nil), "google.firestore.v1beta1.Document.FieldsEntry")
	proto.RegisterType((*Value)(nil), "google.firestore.v1beta1.Value")
	proto.RegisterType((*ArrayValue)(nil), "google.firestore.v1beta1.ArrayValue")
	proto.RegisterType((*MapValue)(nil), "google.firestore.v1beta1.MapValue")
	proto.RegisterMapType((map[string]*Value)(nil), "google.firestore.v1beta1.MapValue.FieldsEntry")
}

func init() {
	proto.RegisterFile("google/firestore/v1beta1/document.proto", fileDescriptor_document_fe46ba4c8da20ac8)
}

var fileDescriptor_document_fe46ba4c8da20ac8 = []byte{
	// 655 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xcf, 0x6e, 0xd3, 0x4c,
	0x14, 0xc5, 0xe3, 0x24, 0x8d, 0x92, 0xeb, 0xb4, 0xfd, 0x3e, 0xb3, 0x89, 0xa2, 0x8a, 0x86, 0x00,
	0x22, 0x6c, 0x6c, 0xb5, 0x08, 0x81, 0xa8, 0x58, 0x34, 0xa5, 0x69, 0x16, 0x05, 0x55, 0x06, 0x75,
	0x51, 0x55, 0x8a, 0xc6, 0xc9, 0xc4, 0xb2, 0x18, 0xcf, 0x58, 0xe3, 0x71, 0xa5, 0xbc, 0x0e, 0x4b,
	0x16, 0xbc, 0x00, 0x3c, 0x41, 0x9f, 0x0a, 0xcd, 0xdf, 0x56, 0xd0, 0x28, 0x2b, 0x76, 0xf6, 0xbd,
	0xbf, 0x73, 0xee, 0x99, 0xf1, 0x8c, 0xe1, 0x45, 0xca, 0x58, 0x4a, 0x70, 0xb4, 0xcc, 0x38, 0x2e,
	0x05, 0xe3, 0x38, 0xba, 0x39, 0x48, 0xb0, 0x40, 0x07, 0xd1, 0x82, 0xcd, 0xab, 0x1c, 0x53, 0x11,
	0x16, 0x9c, 0x09, 0x16, 0xf4, 0x34, 0x18, 0x3a, 0x30, 0x34, 0x60, 0x7f, 0xcf, 0x58, 0xa0, 0x22,
	0x8b, 0x10, 0xa5, 0x4c, 0x20, 0x91, 0x31, 0x5a, 0x6a, 0x9d, 0xeb, 0xaa, 0xb7, 0xa4, 0x5a, 0x46,
	0xa5, 0xe0, 0xd5, 0xdc, 0xb8, 0xf6, 0xf7, 0xff, 0xec, 0x8a, 0x2c, 0xc7, 0xa5, 0x40, 0x79, 0x61,
	0x00, 0x33, 0x36, 0x12, 0xab, 0x02, 0x47, 0x04, 0x09, 0x42, 0x53, 0xdd, 0x19, 0xfe, 0xaa, 0x43,
	0xfb, 0x83, 0xc9, 0x18, 0x04, 0xd0, 0xa4, 0x28, 0xc7, 0x3d, 0x6f, 0xe0, 0x8d, 0x3a, 0xb1, 0x7a,
	0x0e, 0x26, 0xd0, 0x5a, 0x66, 0x98, 0x2c, 0xca, 0x5e, 0x7d, 0xd0, 0x18, 0xf9, 0x87, 0x61, 0xb8,
	0x6e, 0x09, 0xa1, 0xf5, 0x09, 0x27, 0x4a, 0x70, 0x4a, 0x05, 0x5f, 0xc5, 0x46, 0x1d, 0x1c, 0x81,
	0x3f, 0xe7, 0x18, 0x09, 0x3c, 0x93, 0xe1, 0x7a, 0x8d, 0x81, 0x37, 0xf2, 0x0f, 0xfb, 0xd6, 0xcc,
	0x26, 0x0f, 0xbf, 0xd8, 0xe4, 0x31, 0x68, 0x5c, 0x16, 0xa4, 0xb8, 0x2a, 0x16, 0x4e, 0xdc, 0xdc,
	0x2c, 0xd6, 0xb8, 0x2c, 0xf4, 0xaf, 0xc0, 0xbf, 0x17, 0x28, 0xf8, 0x0f, 0x1a, 0x5f, 0xf1, 0xca,
	0xac, 0x51, 0x3e, 0x06, 0xaf, 0x61, 0xeb, 0x06, 0x91, 0x0a, 0xf7, 0xea, 0xca, 0x77, 0x7f, 0xfd,
	0x0a, 0x2f, 0x25, 0x16, 0x6b, 0xfa, 0x5d, 0xfd, 0xad, 0x37, 0xbc, 0x6d, 0xc2, 0x96, 0x2a, 0x06,
	0x47, 0x00, 0xb4, 0x22, 0x64, 0xa6, 0x9d, 0xfc, 0x81, 0x37, 0xda, 0x79, 0x20, 0xe1, 0xa7, 0x8a,
	0x10, 0xc5, 0x4f, 0x6b, 0x71, 0x87, 0xda, 0x97, 0xe0, 0x39, 0x6c, 0x27, 0x8c, 0x11, 0x8c, 0xa8,
	0xd1, 0xcb, 0x74, 0xed, 0x69, 0x2d, 0xee, 0x9a, 0xb2, 0xc3, 0x32, 0x2a, 0x70, 0x8a, 0xf9, 0xec,
	0x2e, 0x70, 0x43, 0x62, 0xa6, 0xac, 0xb1, 0xa7, 0xd0, 0x5d, 0xb0, 0x2a, 0x21, 0xd8, 0x50, 0x72,
	0xaf, 0xbd, 0x69, 0x2d, 0xf6, 0x75, 0x55, 0x43, 0xa7, 0xb0, 0xeb, 0x4e, 0x89, 0xe1, 0x60, 0xd3,
	0xb6, 0x4e, 0x6b, 0xf1, 0x8e, 0x13, 0xb9, 0x59, 0xa5, 0xe0, 0x19, 0x4d, 0x8d, 0xc7, 0xff, 0x72,
	0x5b, 0xe5, 0x2c, 0x5d, 0xd5, 0xd0, 0x13, 0xf0, 0x93, 0x95, 0xc0, 0xa5, 0x61, 0x82, 0x81, 0x37,
	0xea, 0x4e, 0x6b, 0x31, 0xa8, 0xa2, 0x46, 0x5e, 0xc2, 0x2e, 0xc7, 0x4b, 0xcc, 0x31, 0x9d, 0xdb,
	0xd8, 0x5b, 0xc6, 0x6a, 0xc7, 0x35, 0x34, 0xfa, 0x1e, 0x76, 0x53, 0xcc, 0x66, 0x05, 0xcb, 0xa8,
	0x30, 0x68, 0x5b, 0x25, 0x7f, 0x64, 0x93, 0xcb, 0x63, 0x1e, 0x9e, 0x23, 0x71, 0x4e, 0xd3, 0x69,
	0x2d, 0xde, 0x4e, 0x31, 0xbb, 0x90, 0xb0, 0x96, 0x9f, 0x81, 0x8f, 0x38, 0x47, 0x2b, 0x23, 0xed,
	0x28, 0xe9, 0xb3, 0xf5, 0xdf, 0xfc, 0x58, 0xc2, 0xf6, 0x9b, 0x01, 0x72, 0x6f, 0xc1, 0x31, 0x74,
	0x72, 0x64, 0xf7, 0xae, 0xa5, 0x6c, 0x86, 0xeb, 0x6d, 0x3e, 0xa2, 0xc2, 0x9a, 0xb4, 0x73, 0xf3,
	0x3c, 0xee, 0x02, 0x28, 0xf9, 0x4c, 0x26, 0x1e, 0x9e, 0x02, 0xdc, 0x0d, 0x0b, 0xde, 0x40, 0x4b,
	0xf5, 0xca, 0x9e, 0xa7, 0x2e, 0xde, 0xc6, 0x63, 0x69, 0xf0, 0xe1, 0x0f, 0x0f, 0xda, 0x76, 0xda,
	0xbd, 0xeb, 0xeb, 0x6d, 0xba, 0xbe, 0x56, 0xf3, 0xd0, 0xf5, 0xfd, 0x97, 0x97, 0x68, 0xfc, 0xd3,
	0x83, 0xbd, 0x39, 0xcb, 0xd7, 0x2a, 0xc6, 0xdb, 0xf6, 0xcf, 0x72, 0x21, 0x8f, 0xe4, 0x85, 0x77,
	0x75, 0x6c, 0xd0, 0x94, 0x11, 0x44, 0xd3, 0x90, 0xf1, 0x34, 0x4a, 0x31, 0x55, 0x07, 0x36, 0xd2,
	0x2d, 0x54, 0x64, 0xe5, 0xdf, 0xbf, 0xe3, 0x23, 0x57, 0xf9, 0x56, 0x6f, 0x9e, 0x9d, 0x4c, 0x3e,
	0x7f, 0xaf, 0x3f, 0x3e, 0xd3, 0x56, 0x27, 0x84, 0x55, 0x8b, 0x70, 0xe2, 0x66, 0x5f, 0x1e, 0x8c,
	0xa5, 0xe2, 0xd6, 0x02, 0xd7, 0x0a, 0xb8, 0x76, 0xc0, 0xf5, 0xa5, 0xb6, 0x4c, 0x5a, 0x6a, 0xec,
	0xab, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x27, 0x82, 0xde, 0x04, 0x06, 0x00, 0x00,
}
