// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/spanner/v1/mutation.proto

package spanner // import "github.com/sgtsquiggs/go-genproto/googleapis/spanner/v1"

import proto "github.com/sgtsquiggs/protobuf/proto"
import fmt "fmt"
import math "math"
import _struct "github.com/sgtsquiggs/protobuf/ptypes/struct"
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

// A modification to one or more Cloud Spanner rows.  Mutations can be
// applied to a Cloud Spanner database by sending them in a
// [Commit][google.spanner.v1.Spanner.Commit] call.
type Mutation struct {
	// Required. The operation to perform.
	//
	// Types that are valid to be assigned to Operation:
	//	*Mutation_Insert
	//	*Mutation_Update
	//	*Mutation_InsertOrUpdate
	//	*Mutation_Replace
	//	*Mutation_Delete_
	Operation            isMutation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Mutation) Reset()         { *m = Mutation{} }
func (m *Mutation) String() string { return proto.CompactTextString(m) }
func (*Mutation) ProtoMessage()    {}
func (*Mutation) Descriptor() ([]byte, []int) {
	return fileDescriptor_mutation_eb7ac5a738061ac4, []int{0}
}
func (m *Mutation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mutation.Unmarshal(m, b)
}
func (m *Mutation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mutation.Marshal(b, m, deterministic)
}
func (dst *Mutation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mutation.Merge(dst, src)
}
func (m *Mutation) XXX_Size() int {
	return xxx_messageInfo_Mutation.Size(m)
}
func (m *Mutation) XXX_DiscardUnknown() {
	xxx_messageInfo_Mutation.DiscardUnknown(m)
}

var xxx_messageInfo_Mutation proto.InternalMessageInfo

type isMutation_Operation interface {
	isMutation_Operation()
}

type Mutation_Insert struct {
	Insert *Mutation_Write `protobuf:"bytes,1,opt,name=insert,proto3,oneof"`
}

type Mutation_Update struct {
	Update *Mutation_Write `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type Mutation_InsertOrUpdate struct {
	InsertOrUpdate *Mutation_Write `protobuf:"bytes,3,opt,name=insert_or_update,json=insertOrUpdate,proto3,oneof"`
}

type Mutation_Replace struct {
	Replace *Mutation_Write `protobuf:"bytes,4,opt,name=replace,proto3,oneof"`
}

type Mutation_Delete_ struct {
	Delete *Mutation_Delete `protobuf:"bytes,5,opt,name=delete,proto3,oneof"`
}

func (*Mutation_Insert) isMutation_Operation() {}

func (*Mutation_Update) isMutation_Operation() {}

func (*Mutation_InsertOrUpdate) isMutation_Operation() {}

func (*Mutation_Replace) isMutation_Operation() {}

func (*Mutation_Delete_) isMutation_Operation() {}

func (m *Mutation) GetOperation() isMutation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *Mutation) GetInsert() *Mutation_Write {
	if x, ok := m.GetOperation().(*Mutation_Insert); ok {
		return x.Insert
	}
	return nil
}

func (m *Mutation) GetUpdate() *Mutation_Write {
	if x, ok := m.GetOperation().(*Mutation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *Mutation) GetInsertOrUpdate() *Mutation_Write {
	if x, ok := m.GetOperation().(*Mutation_InsertOrUpdate); ok {
		return x.InsertOrUpdate
	}
	return nil
}

func (m *Mutation) GetReplace() *Mutation_Write {
	if x, ok := m.GetOperation().(*Mutation_Replace); ok {
		return x.Replace
	}
	return nil
}

func (m *Mutation) GetDelete() *Mutation_Delete {
	if x, ok := m.GetOperation().(*Mutation_Delete_); ok {
		return x.Delete
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Mutation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Mutation_OneofMarshaler, _Mutation_OneofUnmarshaler, _Mutation_OneofSizer, []interface{}{
		(*Mutation_Insert)(nil),
		(*Mutation_Update)(nil),
		(*Mutation_InsertOrUpdate)(nil),
		(*Mutation_Replace)(nil),
		(*Mutation_Delete_)(nil),
	}
}

func _Mutation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Mutation)
	// operation
	switch x := m.Operation.(type) {
	case *Mutation_Insert:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Insert); err != nil {
			return err
		}
	case *Mutation_Update:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Update); err != nil {
			return err
		}
	case *Mutation_InsertOrUpdate:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.InsertOrUpdate); err != nil {
			return err
		}
	case *Mutation_Replace:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Replace); err != nil {
			return err
		}
	case *Mutation_Delete_:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Delete); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Mutation.Operation has unexpected type %T", x)
	}
	return nil
}

func _Mutation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Mutation)
	switch tag {
	case 1: // operation.insert
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Mutation_Write)
		err := b.DecodeMessage(msg)
		m.Operation = &Mutation_Insert{msg}
		return true, err
	case 2: // operation.update
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Mutation_Write)
		err := b.DecodeMessage(msg)
		m.Operation = &Mutation_Update{msg}
		return true, err
	case 3: // operation.insert_or_update
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Mutation_Write)
		err := b.DecodeMessage(msg)
		m.Operation = &Mutation_InsertOrUpdate{msg}
		return true, err
	case 4: // operation.replace
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Mutation_Write)
		err := b.DecodeMessage(msg)
		m.Operation = &Mutation_Replace{msg}
		return true, err
	case 5: // operation.delete
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Mutation_Delete)
		err := b.DecodeMessage(msg)
		m.Operation = &Mutation_Delete_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Mutation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Mutation)
	// operation
	switch x := m.Operation.(type) {
	case *Mutation_Insert:
		s := proto.Size(x.Insert)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Mutation_Update:
		s := proto.Size(x.Update)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Mutation_InsertOrUpdate:
		s := proto.Size(x.InsertOrUpdate)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Mutation_Replace:
		s := proto.Size(x.Replace)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Mutation_Delete_:
		s := proto.Size(x.Delete)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Arguments to [insert][google.spanner.v1.Mutation.insert], [update][google.spanner.v1.Mutation.update], [insert_or_update][google.spanner.v1.Mutation.insert_or_update], and
// [replace][google.spanner.v1.Mutation.replace] operations.
type Mutation_Write struct {
	// Required. The table whose rows will be written.
	Table string `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	// The names of the columns in [table][google.spanner.v1.Mutation.Write.table] to be written.
	//
	// The list of columns must contain enough columns to allow
	// Cloud Spanner to derive values for all primary key columns in the
	// row(s) to be modified.
	Columns []string `protobuf:"bytes,2,rep,name=columns,proto3" json:"columns,omitempty"`
	// The values to be written. `values` can contain more than one
	// list of values. If it does, then multiple rows are written, one
	// for each entry in `values`. Each list in `values` must have
	// exactly as many entries as there are entries in [columns][google.spanner.v1.Mutation.Write.columns]
	// above. Sending multiple lists is equivalent to sending multiple
	// `Mutation`s, each containing one `values` entry and repeating
	// [table][google.spanner.v1.Mutation.Write.table] and [columns][google.spanner.v1.Mutation.Write.columns]. Individual values in each list are
	// encoded as described [here][google.spanner.v1.TypeCode].
	Values               []*_struct.ListValue `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Mutation_Write) Reset()         { *m = Mutation_Write{} }
func (m *Mutation_Write) String() string { return proto.CompactTextString(m) }
func (*Mutation_Write) ProtoMessage()    {}
func (*Mutation_Write) Descriptor() ([]byte, []int) {
	return fileDescriptor_mutation_eb7ac5a738061ac4, []int{0, 0}
}
func (m *Mutation_Write) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mutation_Write.Unmarshal(m, b)
}
func (m *Mutation_Write) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mutation_Write.Marshal(b, m, deterministic)
}
func (dst *Mutation_Write) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mutation_Write.Merge(dst, src)
}
func (m *Mutation_Write) XXX_Size() int {
	return xxx_messageInfo_Mutation_Write.Size(m)
}
func (m *Mutation_Write) XXX_DiscardUnknown() {
	xxx_messageInfo_Mutation_Write.DiscardUnknown(m)
}

var xxx_messageInfo_Mutation_Write proto.InternalMessageInfo

func (m *Mutation_Write) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *Mutation_Write) GetColumns() []string {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *Mutation_Write) GetValues() []*_struct.ListValue {
	if m != nil {
		return m.Values
	}
	return nil
}

// Arguments to [delete][google.spanner.v1.Mutation.delete] operations.
type Mutation_Delete struct {
	// Required. The table whose rows will be deleted.
	Table string `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	// Required. The primary keys of the rows within [table][google.spanner.v1.Mutation.Delete.table] to delete.
	// Delete is idempotent. The transaction will succeed even if some or all
	// rows do not exist.
	KeySet               *KeySet  `protobuf:"bytes,2,opt,name=key_set,json=keySet,proto3" json:"key_set,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Mutation_Delete) Reset()         { *m = Mutation_Delete{} }
func (m *Mutation_Delete) String() string { return proto.CompactTextString(m) }
func (*Mutation_Delete) ProtoMessage()    {}
func (*Mutation_Delete) Descriptor() ([]byte, []int) {
	return fileDescriptor_mutation_eb7ac5a738061ac4, []int{0, 1}
}
func (m *Mutation_Delete) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mutation_Delete.Unmarshal(m, b)
}
func (m *Mutation_Delete) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mutation_Delete.Marshal(b, m, deterministic)
}
func (dst *Mutation_Delete) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mutation_Delete.Merge(dst, src)
}
func (m *Mutation_Delete) XXX_Size() int {
	return xxx_messageInfo_Mutation_Delete.Size(m)
}
func (m *Mutation_Delete) XXX_DiscardUnknown() {
	xxx_messageInfo_Mutation_Delete.DiscardUnknown(m)
}

var xxx_messageInfo_Mutation_Delete proto.InternalMessageInfo

func (m *Mutation_Delete) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *Mutation_Delete) GetKeySet() *KeySet {
	if m != nil {
		return m.KeySet
	}
	return nil
}

func init() {
	proto.RegisterType((*Mutation)(nil), "google.spanner.v1.Mutation")
	proto.RegisterType((*Mutation_Write)(nil), "google.spanner.v1.Mutation.Write")
	proto.RegisterType((*Mutation_Delete)(nil), "google.spanner.v1.Mutation.Delete")
}

func init() {
	proto.RegisterFile("google/spanner/v1/mutation.proto", fileDescriptor_mutation_eb7ac5a738061ac4)
}

var fileDescriptor_mutation_eb7ac5a738061ac4 = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xd1, 0xea, 0xd3, 0x30,
	0x14, 0xc6, 0xed, 0xba, 0x75, 0x2e, 0x43, 0xd1, 0xa2, 0x58, 0x8b, 0x17, 0x75, 0x57, 0xbb, 0x4a,
	0x69, 0xbd, 0x11, 0xa6, 0x37, 0x53, 0x50, 0xd0, 0xe1, 0xe8, 0x70, 0x82, 0x0c, 0x46, 0xd6, 0x1d,
	0x4b, 0x69, 0x96, 0x94, 0x24, 0x1d, 0xec, 0x45, 0xbc, 0xf4, 0x01, 0x7c, 0x14, 0x9f, 0x4a, 0x9a,
	0xa4, 0x32, 0x9c, 0xfe, 0xd9, 0xff, 0xaa, 0x3d, 0x7c, 0xdf, 0xef, 0x3b, 0xe7, 0x24, 0x41, 0x51,
	0xc1, 0x79, 0x41, 0x21, 0x96, 0x35, 0x61, 0x0c, 0x44, 0x7c, 0x4c, 0xe2, 0x43, 0xa3, 0x88, 0x2a,
	0x39, 0xc3, 0xb5, 0xe0, 0x8a, 0xfb, 0x0f, 0x8d, 0x03, 0x5b, 0x07, 0x3e, 0x26, 0xe1, 0x33, 0x0b,
	0x91, 0xba, 0x8c, 0x09, 0x63, 0xdc, 0xf8, 0xa5, 0x01, 0xfe, 0xa8, 0xba, 0xda, 0x35, 0xdf, 0x62,
	0xa9, 0x44, 0x93, 0xab, 0xbf, 0xd4, 0xb3, 0x86, 0x15, 0x9c, 0x2c, 0x3b, 0xf9, 0xd1, 0x47, 0x77,
	0x17, 0xb6, 0xbf, 0x3f, 0x43, 0x5e, 0xc9, 0x24, 0x08, 0x15, 0x38, 0x91, 0x33, 0x1d, 0xa7, 0xcf,
	0xf1, 0xc5, 0x28, 0xb8, 0x33, 0xe3, 0x2f, 0xa2, 0x54, 0xf0, 0xfe, 0x4e, 0x66, 0x91, 0x16, 0x6e,
	0xea, 0x3d, 0x51, 0x10, 0xf4, 0x6e, 0x01, 0x1b, 0xc4, 0x5f, 0xa0, 0x07, 0x26, 0x66, 0xcb, 0xc5,
	0xd6, 0xc6, 0xb8, 0xd7, 0xc7, 0xdc, 0x37, 0xf0, 0x27, 0xf1, 0xd9, 0xc4, 0xbd, 0x46, 0x43, 0x01,
	0x35, 0x25, 0x39, 0x04, 0xfd, 0xeb, 0x53, 0x3a, 0xc6, 0x7f, 0x85, 0xbc, 0x3d, 0x50, 0x50, 0x10,
	0x0c, 0x34, 0x3d, 0xb9, 0x89, 0x7e, 0xab, 0x9d, 0xed, 0x2e, 0x86, 0x09, 0x2b, 0x34, 0xd0, 0x89,
	0xfe, 0x23, 0x34, 0x50, 0x64, 0x47, 0x41, 0x9f, 0xe6, 0x28, 0x33, 0x85, 0x1f, 0xa0, 0x61, 0xce,
	0x69, 0x73, 0x60, 0x32, 0xe8, 0x45, 0xee, 0x74, 0x94, 0x75, 0xa5, 0x9f, 0x22, 0xef, 0x48, 0x68,
	0x03, 0x32, 0x70, 0x23, 0x77, 0x3a, 0x4e, 0xc3, 0xae, 0x6d, 0x77, 0xb1, 0xf8, 0x63, 0x29, 0xd5,
	0xba, 0xb5, 0x64, 0xd6, 0x19, 0x66, 0xc8, 0x33, 0x03, 0xfc, 0xa7, 0x5b, 0x8a, 0x86, 0x15, 0x9c,
	0xb6, 0x12, 0x94, 0xbd, 0x96, 0xa7, 0xff, 0xd8, 0xe5, 0x03, 0x9c, 0x56, 0xa0, 0x32, 0xaf, 0xd2,
	0xdf, 0xf9, 0x18, 0x8d, 0x78, 0x0d, 0x42, 0xaf, 0x37, 0xff, 0xee, 0xa0, 0xc7, 0x39, 0x3f, 0x5c,
	0x52, 0xf3, 0x7b, 0xdd, 0x11, 0x2c, 0xdb, 0xf1, 0x96, 0xce, 0xd7, 0x97, 0xd6, 0x53, 0x70, 0x4a,
	0x58, 0x81, 0xb9, 0x28, 0xe2, 0x02, 0x98, 0x1e, 0x3e, 0x36, 0x12, 0xa9, 0x4b, 0x79, 0xf6, 0x10,
	0x67, 0xf6, 0xf7, 0x67, 0xef, 0xc9, 0x3b, 0x83, 0xbe, 0xa1, 0xbc, 0xd9, 0xe3, 0x95, 0x6d, 0xb2,
	0x4e, 0x7e, 0x75, 0xca, 0x46, 0x2b, 0x1b, 0xab, 0x6c, 0xd6, 0xc9, 0xce, 0xd3, 0xc1, 0x2f, 0x7e,
	0x07, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x69, 0x1c, 0xbc, 0x51, 0x03, 0x00, 0x00,
}
