// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/containeranalysis/v1beta1/attestation/attestation.proto

package attestation // import "github.com/sgtsquiggs/go-genproto/googleapis/devtools/containeranalysis/v1beta1/attestation"

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

// Type (for example schema) of the attestation payload that was signed.
type PgpSignedAttestation_ContentType int32

const (
	// `ContentType` is not set.
	PgpSignedAttestation_CONTENT_TYPE_UNSPECIFIED PgpSignedAttestation_ContentType = 0
	// Atomic format attestation signature. See
	// https://github.com/containers/image/blob/8a5d2f82a6e3263290c8e0276c3e0f64e77723e7/docs/atomic-signature.md
	// The payload extracted from `signature` is a JSON blob conforming to the
	// linked schema.
	PgpSignedAttestation_SIMPLE_SIGNING_JSON PgpSignedAttestation_ContentType = 1
)

var PgpSignedAttestation_ContentType_name = map[int32]string{
	0: "CONTENT_TYPE_UNSPECIFIED",
	1: "SIMPLE_SIGNING_JSON",
}
var PgpSignedAttestation_ContentType_value = map[string]int32{
	"CONTENT_TYPE_UNSPECIFIED": 0,
	"SIMPLE_SIGNING_JSON":      1,
}

func (x PgpSignedAttestation_ContentType) String() string {
	return proto.EnumName(PgpSignedAttestation_ContentType_name, int32(x))
}
func (PgpSignedAttestation_ContentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_attestation_c9ed0362441a7a04, []int{0, 0}
}

// An attestation wrapper with a PGP-compatible signature. This message only
// supports `ATTACHED` signatures, where the payload that is signed is included
// alongside the signature itself in the same file.
type PgpSignedAttestation struct {
	// The raw content of the signature, as output by GNU Privacy Guard (GPG) or
	// equivalent.  Since this message only supports attached signatures, the
	// payload that was signed must be attached. While the signature format
	// supported is dependent on the verification implementation, currently only
	// ASCII-armored (`--armor` to gpg), non-clearsigned (`--sign` rather than
	// `--clearsign` to gpg) are supported. Concretely, `gpg --sign --armor
	// --output=signature.gpg payload.json` will create the signature content
	// expected in this field in `signature.gpg` for the `payload.json`
	// attestation payload.
	Signature string `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	// Type (for example schema) of the attestation payload that was signed.
	// The verifier must ensure that the provided type is one that the verifier
	// supports, and that the attestation payload is a valid instantiation of that
	// type (for example by validating a JSON schema).
	ContentType PgpSignedAttestation_ContentType `protobuf:"varint,3,opt,name=content_type,json=contentType,proto3,enum=grafeas.v1beta1.attestation.PgpSignedAttestation_ContentType" json:"content_type,omitempty"`
	// This field is used by verifiers to select the public key used to validate
	// the signature.  Note that the policy of the verifier ultimately determines
	// which public keys verify a signature based on the context of the
	// verification.  There is no guarantee validation will succeed if the
	// verifier has no key matching this ID, even if it has a key under a
	// different ID that would verify the signature. Note that this ID should also
	// be present in the signature content above, but that is not expected to be
	// used by the verifier.
	//
	// Types that are valid to be assigned to KeyId:
	//	*PgpSignedAttestation_PgpKeyId
	KeyId                isPgpSignedAttestation_KeyId `protobuf_oneof:"key_id"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *PgpSignedAttestation) Reset()         { *m = PgpSignedAttestation{} }
func (m *PgpSignedAttestation) String() string { return proto.CompactTextString(m) }
func (*PgpSignedAttestation) ProtoMessage()    {}
func (*PgpSignedAttestation) Descriptor() ([]byte, []int) {
	return fileDescriptor_attestation_c9ed0362441a7a04, []int{0}
}
func (m *PgpSignedAttestation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PgpSignedAttestation.Unmarshal(m, b)
}
func (m *PgpSignedAttestation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PgpSignedAttestation.Marshal(b, m, deterministic)
}
func (dst *PgpSignedAttestation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PgpSignedAttestation.Merge(dst, src)
}
func (m *PgpSignedAttestation) XXX_Size() int {
	return xxx_messageInfo_PgpSignedAttestation.Size(m)
}
func (m *PgpSignedAttestation) XXX_DiscardUnknown() {
	xxx_messageInfo_PgpSignedAttestation.DiscardUnknown(m)
}

var xxx_messageInfo_PgpSignedAttestation proto.InternalMessageInfo

func (m *PgpSignedAttestation) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *PgpSignedAttestation) GetContentType() PgpSignedAttestation_ContentType {
	if m != nil {
		return m.ContentType
	}
	return PgpSignedAttestation_CONTENT_TYPE_UNSPECIFIED
}

type isPgpSignedAttestation_KeyId interface {
	isPgpSignedAttestation_KeyId()
}

type PgpSignedAttestation_PgpKeyId struct {
	PgpKeyId string `protobuf:"bytes,2,opt,name=pgp_key_id,json=pgpKeyId,proto3,oneof"`
}

func (*PgpSignedAttestation_PgpKeyId) isPgpSignedAttestation_KeyId() {}

func (m *PgpSignedAttestation) GetKeyId() isPgpSignedAttestation_KeyId {
	if m != nil {
		return m.KeyId
	}
	return nil
}

func (m *PgpSignedAttestation) GetPgpKeyId() string {
	if x, ok := m.GetKeyId().(*PgpSignedAttestation_PgpKeyId); ok {
		return x.PgpKeyId
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PgpSignedAttestation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PgpSignedAttestation_OneofMarshaler, _PgpSignedAttestation_OneofUnmarshaler, _PgpSignedAttestation_OneofSizer, []interface{}{
		(*PgpSignedAttestation_PgpKeyId)(nil),
	}
}

func _PgpSignedAttestation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PgpSignedAttestation)
	// key_id
	switch x := m.KeyId.(type) {
	case *PgpSignedAttestation_PgpKeyId:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.PgpKeyId)
	case nil:
	default:
		return fmt.Errorf("PgpSignedAttestation.KeyId has unexpected type %T", x)
	}
	return nil
}

func _PgpSignedAttestation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PgpSignedAttestation)
	switch tag {
	case 2: // key_id.pgp_key_id
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.KeyId = &PgpSignedAttestation_PgpKeyId{x}
		return true, err
	default:
		return false, nil
	}
}

func _PgpSignedAttestation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PgpSignedAttestation)
	// key_id
	switch x := m.KeyId.(type) {
	case *PgpSignedAttestation_PgpKeyId:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.PgpKeyId)))
		n += len(x.PgpKeyId)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Note kind that represents a logical attestation "role" or "authority". For
// example, an organization might have one `Authority` for "QA" and one for
// "build". This Note is intended to act strictly as a grouping mechanism for
// the attached Occurrences (Attestations). This grouping mechanism also
// provides a security boundary, since IAM ACLs gate the ability for a principle
// to attach an Occurrence to a given Note. It also provides a single point of
// lookup to find all attached Attestation Occurrences, even if they don't all
// live in the same project.
type Authority struct {
	// Hint hints at the purpose of the attestation authority.
	Hint                 *Authority_Hint `protobuf:"bytes,1,opt,name=hint,proto3" json:"hint,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Authority) Reset()         { *m = Authority{} }
func (m *Authority) String() string { return proto.CompactTextString(m) }
func (*Authority) ProtoMessage()    {}
func (*Authority) Descriptor() ([]byte, []int) {
	return fileDescriptor_attestation_c9ed0362441a7a04, []int{1}
}
func (m *Authority) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Authority.Unmarshal(m, b)
}
func (m *Authority) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Authority.Marshal(b, m, deterministic)
}
func (dst *Authority) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Authority.Merge(dst, src)
}
func (m *Authority) XXX_Size() int {
	return xxx_messageInfo_Authority.Size(m)
}
func (m *Authority) XXX_DiscardUnknown() {
	xxx_messageInfo_Authority.DiscardUnknown(m)
}

var xxx_messageInfo_Authority proto.InternalMessageInfo

func (m *Authority) GetHint() *Authority_Hint {
	if m != nil {
		return m.Hint
	}
	return nil
}

// This submessage provides human-readable hints about the purpose of the
// Authority. Because the name of a Note acts as its resource reference, it is
// important to disambiguate the canonical name of the Note (which might be a
// UUID for security purposes) from "readable" names more suitable for debug
// output. Note that these hints should NOT be used to look up authorities in
// security sensitive contexts, such as when looking up Attestations to
// verify.
type Authority_Hint struct {
	// The human readable name of this Attestation Authority, for example "qa".
	HumanReadableName    string   `protobuf:"bytes,1,opt,name=human_readable_name,json=humanReadableName,proto3" json:"human_readable_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Authority_Hint) Reset()         { *m = Authority_Hint{} }
func (m *Authority_Hint) String() string { return proto.CompactTextString(m) }
func (*Authority_Hint) ProtoMessage()    {}
func (*Authority_Hint) Descriptor() ([]byte, []int) {
	return fileDescriptor_attestation_c9ed0362441a7a04, []int{1, 0}
}
func (m *Authority_Hint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Authority_Hint.Unmarshal(m, b)
}
func (m *Authority_Hint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Authority_Hint.Marshal(b, m, deterministic)
}
func (dst *Authority_Hint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Authority_Hint.Merge(dst, src)
}
func (m *Authority_Hint) XXX_Size() int {
	return xxx_messageInfo_Authority_Hint.Size(m)
}
func (m *Authority_Hint) XXX_DiscardUnknown() {
	xxx_messageInfo_Authority_Hint.DiscardUnknown(m)
}

var xxx_messageInfo_Authority_Hint proto.InternalMessageInfo

func (m *Authority_Hint) GetHumanReadableName() string {
	if m != nil {
		return m.HumanReadableName
	}
	return ""
}

// Details of an attestation occurrence.
type Details struct {
	// Attestation for the resource.
	Attestation          *Attestation `protobuf:"bytes,1,opt,name=attestation,proto3" json:"attestation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Details) Reset()         { *m = Details{} }
func (m *Details) String() string { return proto.CompactTextString(m) }
func (*Details) ProtoMessage()    {}
func (*Details) Descriptor() ([]byte, []int) {
	return fileDescriptor_attestation_c9ed0362441a7a04, []int{2}
}
func (m *Details) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Details.Unmarshal(m, b)
}
func (m *Details) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Details.Marshal(b, m, deterministic)
}
func (dst *Details) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Details.Merge(dst, src)
}
func (m *Details) XXX_Size() int {
	return xxx_messageInfo_Details.Size(m)
}
func (m *Details) XXX_DiscardUnknown() {
	xxx_messageInfo_Details.DiscardUnknown(m)
}

var xxx_messageInfo_Details proto.InternalMessageInfo

func (m *Details) GetAttestation() *Attestation {
	if m != nil {
		return m.Attestation
	}
	return nil
}

// Occurrence that represents a single "attestation". The authenticity of an
// Attestation can be verified using the attached signature. If the verifier
// trusts the public key of the signer, then verifying the signature is
// sufficient to establish trust. In this circumstance, the Authority to which
// this Attestation is attached is primarily useful for look-up (how to find
// this Attestation if you already know the Authority and artifact to be
// verified) and intent (which authority was this attestation intended to sign
// for).
type Attestation struct {
	// The signature, generally over the `resource_url`, that verifies this
	// attestation. The semantics of the signature veracity are ultimately
	// determined by the verification engine.
	//
	// Types that are valid to be assigned to Signature:
	//	*Attestation_PgpSignedAttestation
	Signature            isAttestation_Signature `protobuf_oneof:"signature"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Attestation) Reset()         { *m = Attestation{} }
func (m *Attestation) String() string { return proto.CompactTextString(m) }
func (*Attestation) ProtoMessage()    {}
func (*Attestation) Descriptor() ([]byte, []int) {
	return fileDescriptor_attestation_c9ed0362441a7a04, []int{3}
}
func (m *Attestation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Attestation.Unmarshal(m, b)
}
func (m *Attestation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Attestation.Marshal(b, m, deterministic)
}
func (dst *Attestation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attestation.Merge(dst, src)
}
func (m *Attestation) XXX_Size() int {
	return xxx_messageInfo_Attestation.Size(m)
}
func (m *Attestation) XXX_DiscardUnknown() {
	xxx_messageInfo_Attestation.DiscardUnknown(m)
}

var xxx_messageInfo_Attestation proto.InternalMessageInfo

type isAttestation_Signature interface {
	isAttestation_Signature()
}

type Attestation_PgpSignedAttestation struct {
	PgpSignedAttestation *PgpSignedAttestation `protobuf:"bytes,1,opt,name=pgp_signed_attestation,json=pgpSignedAttestation,proto3,oneof"`
}

func (*Attestation_PgpSignedAttestation) isAttestation_Signature() {}

func (m *Attestation) GetSignature() isAttestation_Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Attestation) GetPgpSignedAttestation() *PgpSignedAttestation {
	if x, ok := m.GetSignature().(*Attestation_PgpSignedAttestation); ok {
		return x.PgpSignedAttestation
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Attestation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Attestation_OneofMarshaler, _Attestation_OneofUnmarshaler, _Attestation_OneofSizer, []interface{}{
		(*Attestation_PgpSignedAttestation)(nil),
	}
}

func _Attestation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Attestation)
	// signature
	switch x := m.Signature.(type) {
	case *Attestation_PgpSignedAttestation:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PgpSignedAttestation); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Attestation.Signature has unexpected type %T", x)
	}
	return nil
}

func _Attestation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Attestation)
	switch tag {
	case 1: // signature.pgp_signed_attestation
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PgpSignedAttestation)
		err := b.DecodeMessage(msg)
		m.Signature = &Attestation_PgpSignedAttestation{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Attestation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Attestation)
	// signature
	switch x := m.Signature.(type) {
	case *Attestation_PgpSignedAttestation:
		s := proto.Size(x.PgpSignedAttestation)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*PgpSignedAttestation)(nil), "grafeas.v1beta1.attestation.PgpSignedAttestation")
	proto.RegisterType((*Authority)(nil), "grafeas.v1beta1.attestation.Authority")
	proto.RegisterType((*Authority_Hint)(nil), "grafeas.v1beta1.attestation.Authority.Hint")
	proto.RegisterType((*Details)(nil), "grafeas.v1beta1.attestation.Details")
	proto.RegisterType((*Attestation)(nil), "grafeas.v1beta1.attestation.Attestation")
	proto.RegisterEnum("grafeas.v1beta1.attestation.PgpSignedAttestation_ContentType", PgpSignedAttestation_ContentType_name, PgpSignedAttestation_ContentType_value)
}

func init() {
	proto.RegisterFile("google/devtools/containeranalysis/v1beta1/attestation/attestation.proto", fileDescriptor_attestation_c9ed0362441a7a04)
}

var fileDescriptor_attestation_c9ed0362441a7a04 = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x6d, 0xba, 0x69, 0xac, 0x37, 0x08, 0x8d, 0x6c, 0x82, 0x0a, 0xa6, 0x69, 0xca, 0x53, 0x25,
	0x24, 0x47, 0x1d, 0x12, 0x2f, 0x08, 0xa1, 0x7e, 0x84, 0x36, 0x03, 0xb2, 0x28, 0xe9, 0x1e, 0xe0,
	0xc5, 0x73, 0x97, 0x8b, 0x6b, 0x91, 0xda, 0x56, 0xe2, 0x4e, 0xca, 0x3b, 0xf0, 0xc0, 0xcf, 0xe0,
	0x97, 0xa2, 0xa6, 0x85, 0x46, 0xda, 0x54, 0xa9, 0x6f, 0xd7, 0x3e, 0x3e, 0xf7, 0x9c, 0xfb, 0x61,
	0x18, 0x71, 0xa5, 0x78, 0x86, 0x5e, 0x8a, 0x77, 0x46, 0xa9, 0xac, 0xf0, 0x6e, 0x95, 0x34, 0x4c,
	0x48, 0xcc, 0x99, 0x64, 0x59, 0x59, 0x88, 0xc2, 0xbb, 0xeb, 0x4e, 0xd1, 0xb0, 0xae, 0xc7, 0x8c,
	0xc1, 0xc2, 0x30, 0x23, 0x94, 0xac, 0xc7, 0x44, 0xe7, 0xca, 0x28, 0xe7, 0x25, 0xcf, 0xd9, 0x37,
	0x64, 0x05, 0x59, 0x3f, 0x27, 0xb5, 0x27, 0xee, 0xaf, 0x26, 0x9c, 0x44, 0x5c, 0x27, 0x82, 0x4b,
	0x4c, 0x7b, 0x1b, 0xc0, 0x39, 0x85, 0x56, 0x21, 0xb8, 0x64, 0x66, 0x91, 0x63, 0xdb, 0x3a, 0xb7,
	0x3a, 0xad, 0x78, 0x73, 0xe1, 0xdc, 0xc0, 0xe3, 0xa5, 0x1d, 0x94, 0x86, 0x9a, 0x52, 0x63, 0x7b,
	0xef, 0xdc, 0xea, 0x3c, 0xb9, 0x78, 0x47, 0xb6, 0x48, 0x91, 0x87, 0x64, 0xc8, 0x60, 0x95, 0x65,
	0x52, 0x6a, 0x8c, 0xed, 0xdb, 0xcd, 0xc1, 0x39, 0x03, 0xd0, 0x5c, 0xd3, 0xef, 0x58, 0x52, 0x91,
	0xb6, 0x9b, 0x4b, 0x03, 0xe3, 0x46, 0x7c, 0xa8, 0xb9, 0xfe, 0x88, 0x65, 0x90, 0xba, 0x43, 0xb0,
	0x6b, 0x5c, 0xe7, 0x14, 0xda, 0x83, 0xab, 0x70, 0xe2, 0x87, 0x13, 0x3a, 0xf9, 0x12, 0xf9, 0xf4,
	0x3a, 0x4c, 0x22, 0x7f, 0x10, 0x7c, 0x08, 0xfc, 0xe1, 0x51, 0xc3, 0x79, 0x0e, 0xc7, 0x49, 0xf0,
	0x39, 0xfa, 0xe4, 0xd3, 0x24, 0x18, 0x85, 0x41, 0x38, 0xa2, 0x97, 0xc9, 0x55, 0x78, 0x64, 0xf5,
	0x0f, 0xe1, 0x60, 0xa5, 0xe0, 0xfe, 0xb0, 0xa0, 0xd5, 0x5b, 0x98, 0x99, 0xca, 0x85, 0x29, 0x9d,
	0xf7, 0xb0, 0x3f, 0x13, 0xd2, 0x54, 0x85, 0xdb, 0x17, 0xaf, 0xb6, 0xd6, 0xf5, 0x9f, 0x45, 0xc6,
	0x42, 0x9a, 0xb8, 0x22, 0xbe, 0x78, 0x03, 0xfb, 0xcb, 0x93, 0x43, 0xe0, 0x78, 0xb6, 0x98, 0x33,
	0x49, 0x73, 0x64, 0x29, 0x9b, 0x66, 0x48, 0x25, 0x9b, 0xff, 0x6b, 0xe8, 0xd3, 0x0a, 0x8a, 0xd7,
	0x48, 0xc8, 0xe6, 0xe8, 0x5e, 0xc3, 0xa3, 0x21, 0x1a, 0x26, 0xb2, 0xc2, 0xb9, 0x04, 0xbb, 0x26,
	0xb3, 0xb6, 0xd2, 0xd9, 0x6e, 0x65, 0x13, 0xc7, 0x75, 0xb2, 0xfb, 0xd3, 0x02, 0xbb, 0x3e, 0x5d,
	0x01, 0xcf, 0x96, 0xdd, 0x2d, 0xaa, 0x79, 0xd0, 0xfb, 0x32, 0xdd, 0x9d, 0x27, 0x39, 0x6e, 0xc4,
	0x27, 0xfa, 0x81, 0xfb, 0xbe, 0x5d, 0x5b, 0xa4, 0xfe, 0x6f, 0x0b, 0xce, 0x84, 0xda, 0x96, 0x3c,
	0xb2, 0xbe, 0xde, 0xac, 0x36, 0x9f, 0x70, 0x95, 0x31, 0xc9, 0x89, 0xca, 0xb9, 0xc7, 0x51, 0x56,
	0xcb, 0xec, 0xad, 0x20, 0xa6, 0x45, 0xb1, 0xe3, 0xc7, 0x78, 0x5b, 0x8b, 0xff, 0x34, 0xf7, 0x46,
	0x71, 0x6f, 0x7a, 0x50, 0xa5, 0x7c, 0xfd, 0x37, 0x00, 0x00, 0xff, 0xff, 0x69, 0x0c, 0x36, 0xf8,
	0x6a, 0x03, 0x00, 0x00,
}
