// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/type/dayofweek.proto

package dayofweek // import "github.com/sgtsquiggs/go-genproto/googleapis/type/dayofweek"

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

// Represents a day of week.
type DayOfWeek int32

const (
	// The unspecified day-of-week.
	DayOfWeek_DAY_OF_WEEK_UNSPECIFIED DayOfWeek = 0
	// The day-of-week of Monday.
	DayOfWeek_MONDAY DayOfWeek = 1
	// The day-of-week of Tuesday.
	DayOfWeek_TUESDAY DayOfWeek = 2
	// The day-of-week of Wednesday.
	DayOfWeek_WEDNESDAY DayOfWeek = 3
	// The day-of-week of Thursday.
	DayOfWeek_THURSDAY DayOfWeek = 4
	// The day-of-week of Friday.
	DayOfWeek_FRIDAY DayOfWeek = 5
	// The day-of-week of Saturday.
	DayOfWeek_SATURDAY DayOfWeek = 6
	// The day-of-week of Sunday.
	DayOfWeek_SUNDAY DayOfWeek = 7
)

var DayOfWeek_name = map[int32]string{
	0: "DAY_OF_WEEK_UNSPECIFIED",
	1: "MONDAY",
	2: "TUESDAY",
	3: "WEDNESDAY",
	4: "THURSDAY",
	5: "FRIDAY",
	6: "SATURDAY",
	7: "SUNDAY",
}
var DayOfWeek_value = map[string]int32{
	"DAY_OF_WEEK_UNSPECIFIED": 0,
	"MONDAY":                  1,
	"TUESDAY":                 2,
	"WEDNESDAY":               3,
	"THURSDAY":                4,
	"FRIDAY":                  5,
	"SATURDAY":                6,
	"SUNDAY":                  7,
}

func (x DayOfWeek) String() string {
	return proto.EnumName(DayOfWeek_name, int32(x))
}
func (DayOfWeek) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dayofweek_65556e97092f8975, []int{0}
}

func init() {
	proto.RegisterEnum("google.type.DayOfWeek", DayOfWeek_name, DayOfWeek_value)
}

func init() {
	proto.RegisterFile("google/type/dayofweek.proto", fileDescriptor_dayofweek_65556e97092f8975)
}

var fileDescriptor_dayofweek_65556e97092f8975 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0x4f, 0x49, 0xac, 0xcc, 0x4f, 0x2b, 0x4f, 0x4d,
	0xcd, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x86, 0x48, 0xea, 0x81, 0x24, 0xb5, 0x5a,
	0x18, 0xb9, 0x38, 0x5d, 0x12, 0x2b, 0xfd, 0xd3, 0xc2, 0x53, 0x53, 0xb3, 0x85, 0xa4, 0xb9, 0xc4,
	0x5d, 0x1c, 0x23, 0xe3, 0xfd, 0xdd, 0xe2, 0xc3, 0x5d, 0x5d, 0xbd, 0xe3, 0x43, 0xfd, 0x82, 0x03,
	0x5c, 0x9d, 0x3d, 0xdd, 0x3c, 0x5d, 0x5d, 0x04, 0x18, 0x84, 0xb8, 0xb8, 0xd8, 0x7c, 0xfd, 0xfd,
	0x5c, 0x1c, 0x23, 0x05, 0x18, 0x85, 0xb8, 0xb9, 0xd8, 0x43, 0x42, 0x5d, 0x83, 0x41, 0x1c, 0x26,
	0x21, 0x5e, 0x2e, 0xce, 0x70, 0x57, 0x17, 0x3f, 0x08, 0x97, 0x59, 0x88, 0x87, 0x8b, 0x23, 0xc4,
	0x23, 0x34, 0x08, 0xcc, 0x63, 0x01, 0xe9, 0x72, 0x0b, 0xf2, 0x04, 0xb1, 0x59, 0x41, 0x32, 0xc1,
	0x8e, 0x21, 0xa1, 0x41, 0x20, 0x1e, 0x1b, 0x48, 0x26, 0x38, 0x14, 0x6c, 0x1e, 0xbb, 0x53, 0x26,
	0x17, 0x7f, 0x72, 0x7e, 0xae, 0x1e, 0x92, 0xcb, 0x9c, 0xf8, 0xe0, 0xce, 0x0a, 0x00, 0x39, 0x3b,
	0x80, 0x31, 0xca, 0x0e, 0x2a, 0x9d, 0x9e, 0x9f, 0x93, 0x98, 0x97, 0xae, 0x97, 0x5f, 0x94, 0xae,
	0x9f, 0x9e, 0x9a, 0x07, 0xf6, 0x94, 0x3e, 0x44, 0x2a, 0xb1, 0x20, 0xb3, 0x18, 0xcd, 0xd3, 0xd6,
	0x70, 0xd6, 0x22, 0x26, 0x66, 0xf7, 0x90, 0x80, 0x24, 0x36, 0xb0, 0x06, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x6e, 0x23, 0xb2, 0xb3, 0x24, 0x01, 0x00, 0x00,
}
