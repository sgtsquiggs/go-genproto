// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/oslogin/common/common.proto

package common // import "github.com/sgtsquiggs/go-genproto/googleapis/cloud/oslogin/common"

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

// The POSIX account information associated with a Google account.
type PosixAccount struct {
	// Only one POSIX account can be marked as primary.
	Primary bool `protobuf:"varint,1,opt,name=primary,proto3" json:"primary,omitempty"`
	// The username of the POSIX account.
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// The user ID.
	Uid int64 `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	// The default group ID.
	Gid int64 `protobuf:"varint,4,opt,name=gid,proto3" json:"gid,omitempty"`
	// The path to the home directory for this account.
	HomeDirectory string `protobuf:"bytes,5,opt,name=home_directory,json=homeDirectory,proto3" json:"home_directory,omitempty"`
	// The path to the logic shell for this account.
	Shell string `protobuf:"bytes,6,opt,name=shell,proto3" json:"shell,omitempty"`
	// The GECOS (user information) entry for this account.
	Gecos string `protobuf:"bytes,7,opt,name=gecos,proto3" json:"gecos,omitempty"`
	// System identifier for which account the username or uid applies to.
	// By default, the empty value is used.
	SystemId string `protobuf:"bytes,8,opt,name=system_id,json=systemId,proto3" json:"system_id,omitempty"`
	// Output only. A POSIX account identifier.
	AccountId            string   `protobuf:"bytes,9,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PosixAccount) Reset()         { *m = PosixAccount{} }
func (m *PosixAccount) String() string { return proto.CompactTextString(m) }
func (*PosixAccount) ProtoMessage()    {}
func (*PosixAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_ebba42564a5ff09e, []int{0}
}
func (m *PosixAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PosixAccount.Unmarshal(m, b)
}
func (m *PosixAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PosixAccount.Marshal(b, m, deterministic)
}
func (dst *PosixAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PosixAccount.Merge(dst, src)
}
func (m *PosixAccount) XXX_Size() int {
	return xxx_messageInfo_PosixAccount.Size(m)
}
func (m *PosixAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_PosixAccount.DiscardUnknown(m)
}

var xxx_messageInfo_PosixAccount proto.InternalMessageInfo

func (m *PosixAccount) GetPrimary() bool {
	if m != nil {
		return m.Primary
	}
	return false
}

func (m *PosixAccount) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *PosixAccount) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *PosixAccount) GetGid() int64 {
	if m != nil {
		return m.Gid
	}
	return 0
}

func (m *PosixAccount) GetHomeDirectory() string {
	if m != nil {
		return m.HomeDirectory
	}
	return ""
}

func (m *PosixAccount) GetShell() string {
	if m != nil {
		return m.Shell
	}
	return ""
}

func (m *PosixAccount) GetGecos() string {
	if m != nil {
		return m.Gecos
	}
	return ""
}

func (m *PosixAccount) GetSystemId() string {
	if m != nil {
		return m.SystemId
	}
	return ""
}

func (m *PosixAccount) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

// The SSH public key information associated with a Google account.
type SshPublicKey struct {
	// Public key text in SSH format, defined by
	// <a href="https://www.ietf.org/rfc/rfc4253.txt" target="_blank">RFC4253</a>
	// section 6.6.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// An expiration time in microseconds since epoch.
	ExpirationTimeUsec int64 `protobuf:"varint,2,opt,name=expiration_time_usec,json=expirationTimeUsec,proto3" json:"expiration_time_usec,omitempty"`
	// Output only. The SHA-256 fingerprint of the SSH public key.
	Fingerprint          string   `protobuf:"bytes,3,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SshPublicKey) Reset()         { *m = SshPublicKey{} }
func (m *SshPublicKey) String() string { return proto.CompactTextString(m) }
func (*SshPublicKey) ProtoMessage()    {}
func (*SshPublicKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_ebba42564a5ff09e, []int{1}
}
func (m *SshPublicKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SshPublicKey.Unmarshal(m, b)
}
func (m *SshPublicKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SshPublicKey.Marshal(b, m, deterministic)
}
func (dst *SshPublicKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SshPublicKey.Merge(dst, src)
}
func (m *SshPublicKey) XXX_Size() int {
	return xxx_messageInfo_SshPublicKey.Size(m)
}
func (m *SshPublicKey) XXX_DiscardUnknown() {
	xxx_messageInfo_SshPublicKey.DiscardUnknown(m)
}

var xxx_messageInfo_SshPublicKey proto.InternalMessageInfo

func (m *SshPublicKey) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SshPublicKey) GetExpirationTimeUsec() int64 {
	if m != nil {
		return m.ExpirationTimeUsec
	}
	return 0
}

func (m *SshPublicKey) GetFingerprint() string {
	if m != nil {
		return m.Fingerprint
	}
	return ""
}

func init() {
	proto.RegisterType((*PosixAccount)(nil), "google.cloud.oslogin.common.PosixAccount")
	proto.RegisterType((*SshPublicKey)(nil), "google.cloud.oslogin.common.SshPublicKey")
}

func init() {
	proto.RegisterFile("google/cloud/oslogin/common/common.proto", fileDescriptor_common_ebba42564a5ff09e)
}

var fileDescriptor_common_ebba42564a5ff09e = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x41, 0x6b, 0x14, 0x31,
	0x14, 0xc7, 0x99, 0xae, 0x6d, 0x77, 0xe2, 0x2a, 0x12, 0x7a, 0x08, 0x5d, 0xc5, 0xa5, 0x20, 0xec,
	0x69, 0x46, 0xf0, 0xe8, 0xa9, 0xad, 0x20, 0x45, 0xc1, 0x65, 0xd4, 0x8b, 0x2c, 0x2c, 0xd3, 0xe4,
	0x99, 0x7d, 0x38, 0xc9, 0x1b, 0x92, 0x0c, 0x74, 0xbf, 0x92, 0x07, 0x3f, 0x88, 0x5f, 0xc8, 0xab,
	0x24, 0x99, 0x51, 0x0f, 0xd2, 0x53, 0xf2, 0xff, 0xff, 0xfe, 0xf3, 0x92, 0x97, 0x37, 0x6c, 0xad,
	0x89, 0x74, 0x07, 0xb5, 0xec, 0x68, 0x50, 0x35, 0xf9, 0x8e, 0x34, 0xda, 0x5a, 0x92, 0x31, 0x34,
	0x2d, 0x55, 0xef, 0x28, 0x10, 0x5f, 0xe6, 0x64, 0x95, 0x92, 0xd5, 0x98, 0xac, 0x72, 0xe4, 0xfc,
	0xe9, 0x58, 0xa6, 0xed, 0xb1, 0x6e, 0xad, 0xa5, 0xd0, 0x06, 0x24, 0xeb, 0xf3, 0xa7, 0x17, 0xbf,
	0x0a, 0xb6, 0xd8, 0x90, 0xc7, 0xbb, 0x4b, 0x29, 0x69, 0xb0, 0x81, 0x0b, 0x76, 0xda, 0x3b, 0x34,
	0xad, 0x3b, 0x88, 0x62, 0x55, 0xac, 0xe7, 0xcd, 0x24, 0xf9, 0x39, 0x9b, 0x0f, 0x1e, 0x9c, 0x6d,
	0x0d, 0x88, 0xa3, 0x55, 0xb1, 0x2e, 0x9b, 0x3f, 0x9a, 0x3f, 0x61, 0xb3, 0x01, 0x95, 0x98, 0xad,
	0x8a, 0xf5, 0xac, 0x89, 0xdb, 0xe8, 0x68, 0x54, 0xe2, 0x41, 0x76, 0x34, 0x2a, 0xfe, 0x82, 0x3d,
	0xde, 0x93, 0x81, 0x9d, 0x42, 0x07, 0x32, 0x90, 0x3b, 0x88, 0xe3, 0x54, 0xe5, 0x51, 0x74, 0xdf,
	0x4c, 0x26, 0x3f, 0x63, 0xc7, 0x7e, 0x0f, 0x5d, 0x27, 0x4e, 0x12, 0xcd, 0x22, 0xba, 0x1a, 0x24,
	0x79, 0x71, 0x9a, 0xdd, 0x24, 0xf8, 0x92, 0x95, 0xfe, 0xe0, 0x03, 0x98, 0x1d, 0x2a, 0x31, 0xcf,
	0x77, 0xca, 0xc6, 0x8d, 0xe2, 0xcf, 0x18, 0x6b, 0x73, 0x53, 0x91, 0x96, 0x89, 0x96, 0xa3, 0x73,
	0xa3, 0x2e, 0x02, 0x5b, 0x7c, 0xf4, 0xfb, 0xcd, 0x70, 0xdb, 0xa1, 0x7c, 0x07, 0x87, 0x78, 0xe1,
	0x6f, 0x90, 0x9b, 0x2e, 0x9b, 0xb8, 0xe5, 0x2f, 0xd9, 0x19, 0xdc, 0xf5, 0xe8, 0xd2, 0x83, 0xed,
	0x02, 0x1a, 0xd8, 0x0d, 0x1e, 0x64, 0x6a, 0x7e, 0xd6, 0xf0, 0xbf, 0xec, 0x13, 0x1a, 0xf8, 0xec,
	0x41, 0xf2, 0x15, 0x7b, 0xf8, 0x15, 0xad, 0x06, 0xd7, 0x3b, 0xb4, 0x21, 0x3d, 0x47, 0xd9, 0xfc,
	0x6b, 0x5d, 0xfd, 0x28, 0xd8, 0x73, 0x49, 0xa6, 0xba, 0x67, 0x62, 0x57, 0x8b, 0x0f, 0xfe, 0x7d,
	0xd4, 0x9b, 0x38, 0xa1, 0x2f, 0x97, 0x63, 0x54, 0x53, 0xd7, 0x5a, 0x5d, 0x91, 0xd3, 0xb5, 0x06,
	0x9b, 0xa6, 0x57, 0x67, 0xd4, 0xf6, 0xe8, 0xff, 0xfb, 0x97, 0xbc, 0xce, 0xcb, 0xf7, 0xa3, 0xe5,
	0xdb, 0x5c, 0xe3, 0x3a, 0x1d, 0x37, 0x96, 0xaf, 0xae, 0x13, 0xfd, 0x39, 0xd1, 0x6d, 0xa2, 0xdb,
	0x91, 0x6e, 0x33, 0xbd, 0x3d, 0x49, 0x27, 0xbd, 0xfa, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x30,
	0xf7, 0x5b, 0x8e, 0x02, 0x00, 0x00,
}
