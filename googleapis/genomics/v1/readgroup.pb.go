// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/genomics/v1/readgroup.proto

package genomics // import "github.com/sgtsquiggs/go-genproto/googleapis/genomics/v1"

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

// A read group is all the data that's processed the same way by the sequencer.
type ReadGroup struct {
	// The server-generated read group ID, unique for all read groups.
	// Note: This is different than the @RG ID field in the SAM spec. For that
	// value, see [name][google.genomics.v1.ReadGroup.name].
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The dataset to which this read group belongs.
	DatasetId string `protobuf:"bytes,2,opt,name=dataset_id,json=datasetId,proto3" json:"dataset_id,omitempty"`
	// The read group name. This corresponds to the @RG ID field in the SAM spec.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// A free-form text description of this read group.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// A client-supplied sample identifier for the reads in this read group.
	SampleId string `protobuf:"bytes,5,opt,name=sample_id,json=sampleId,proto3" json:"sample_id,omitempty"`
	// The experiment used to generate this read group.
	Experiment *ReadGroup_Experiment `protobuf:"bytes,6,opt,name=experiment,proto3" json:"experiment,omitempty"`
	// The predicted insert size of this read group. The insert size is the length
	// the sequenced DNA fragment from end-to-end, not including the adapters.
	PredictedInsertSize int32 `protobuf:"varint,7,opt,name=predicted_insert_size,json=predictedInsertSize,proto3" json:"predicted_insert_size,omitempty"`
	// The programs used to generate this read group. Programs are always
	// identical for all read groups within a read group set. For this reason,
	// only the first read group in a returned set will have this field
	// populated.
	Programs []*ReadGroup_Program `protobuf:"bytes,10,rep,name=programs,proto3" json:"programs,omitempty"`
	// The reference set the reads in this read group are aligned to.
	ReferenceSetId string `protobuf:"bytes,11,opt,name=reference_set_id,json=referenceSetId,proto3" json:"reference_set_id,omitempty"`
	// A map of additional read group information. This must be of the form
	// map<string, string[]> (string key mapping to a list of string values).
	Info                 map[string]*_struct.ListValue `protobuf:"bytes,12,rep,name=info,proto3" json:"info,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *ReadGroup) Reset()         { *m = ReadGroup{} }
func (m *ReadGroup) String() string { return proto.CompactTextString(m) }
func (*ReadGroup) ProtoMessage()    {}
func (*ReadGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_readgroup_d1ee37f21a1efad6, []int{0}
}
func (m *ReadGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadGroup.Unmarshal(m, b)
}
func (m *ReadGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadGroup.Marshal(b, m, deterministic)
}
func (dst *ReadGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadGroup.Merge(dst, src)
}
func (m *ReadGroup) XXX_Size() int {
	return xxx_messageInfo_ReadGroup.Size(m)
}
func (m *ReadGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadGroup.DiscardUnknown(m)
}

var xxx_messageInfo_ReadGroup proto.InternalMessageInfo

func (m *ReadGroup) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ReadGroup) GetDatasetId() string {
	if m != nil {
		return m.DatasetId
	}
	return ""
}

func (m *ReadGroup) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReadGroup) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ReadGroup) GetSampleId() string {
	if m != nil {
		return m.SampleId
	}
	return ""
}

func (m *ReadGroup) GetExperiment() *ReadGroup_Experiment {
	if m != nil {
		return m.Experiment
	}
	return nil
}

func (m *ReadGroup) GetPredictedInsertSize() int32 {
	if m != nil {
		return m.PredictedInsertSize
	}
	return 0
}

func (m *ReadGroup) GetPrograms() []*ReadGroup_Program {
	if m != nil {
		return m.Programs
	}
	return nil
}

func (m *ReadGroup) GetReferenceSetId() string {
	if m != nil {
		return m.ReferenceSetId
	}
	return ""
}

func (m *ReadGroup) GetInfo() map[string]*_struct.ListValue {
	if m != nil {
		return m.Info
	}
	return nil
}

type ReadGroup_Experiment struct {
	// A client-supplied library identifier; a library is a collection of DNA
	// fragments which have been prepared for sequencing from a sample. This
	// field is important for quality control as error or bias can be introduced
	// during sample preparation.
	LibraryId string `protobuf:"bytes,1,opt,name=library_id,json=libraryId,proto3" json:"library_id,omitempty"`
	// The platform unit used as part of this experiment, for example
	// flowcell-barcode.lane for Illumina or slide for SOLiD. Corresponds to the
	// @RG PU field in the SAM spec.
	PlatformUnit string `protobuf:"bytes,2,opt,name=platform_unit,json=platformUnit,proto3" json:"platform_unit,omitempty"`
	// The sequencing center used as part of this experiment.
	SequencingCenter string `protobuf:"bytes,3,opt,name=sequencing_center,json=sequencingCenter,proto3" json:"sequencing_center,omitempty"`
	// The instrument model used as part of this experiment. This maps to
	// sequencing technology in the SAM spec.
	InstrumentModel      string   `protobuf:"bytes,4,opt,name=instrument_model,json=instrumentModel,proto3" json:"instrument_model,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadGroup_Experiment) Reset()         { *m = ReadGroup_Experiment{} }
func (m *ReadGroup_Experiment) String() string { return proto.CompactTextString(m) }
func (*ReadGroup_Experiment) ProtoMessage()    {}
func (*ReadGroup_Experiment) Descriptor() ([]byte, []int) {
	return fileDescriptor_readgroup_d1ee37f21a1efad6, []int{0, 0}
}
func (m *ReadGroup_Experiment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadGroup_Experiment.Unmarshal(m, b)
}
func (m *ReadGroup_Experiment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadGroup_Experiment.Marshal(b, m, deterministic)
}
func (dst *ReadGroup_Experiment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadGroup_Experiment.Merge(dst, src)
}
func (m *ReadGroup_Experiment) XXX_Size() int {
	return xxx_messageInfo_ReadGroup_Experiment.Size(m)
}
func (m *ReadGroup_Experiment) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadGroup_Experiment.DiscardUnknown(m)
}

var xxx_messageInfo_ReadGroup_Experiment proto.InternalMessageInfo

func (m *ReadGroup_Experiment) GetLibraryId() string {
	if m != nil {
		return m.LibraryId
	}
	return ""
}

func (m *ReadGroup_Experiment) GetPlatformUnit() string {
	if m != nil {
		return m.PlatformUnit
	}
	return ""
}

func (m *ReadGroup_Experiment) GetSequencingCenter() string {
	if m != nil {
		return m.SequencingCenter
	}
	return ""
}

func (m *ReadGroup_Experiment) GetInstrumentModel() string {
	if m != nil {
		return m.InstrumentModel
	}
	return ""
}

type ReadGroup_Program struct {
	// The command line used to run this program.
	CommandLine string `protobuf:"bytes,1,opt,name=command_line,json=commandLine,proto3" json:"command_line,omitempty"`
	// The user specified locally unique ID of the program. Used along with
	// `prevProgramId` to define an ordering between programs.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The display name of the program. This is typically the colloquial name of
	// the tool used, for example 'bwa' or 'picard'.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The ID of the program run before this one.
	PrevProgramId string `protobuf:"bytes,4,opt,name=prev_program_id,json=prevProgramId,proto3" json:"prev_program_id,omitempty"`
	// The version of the program run.
	Version              string   `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadGroup_Program) Reset()         { *m = ReadGroup_Program{} }
func (m *ReadGroup_Program) String() string { return proto.CompactTextString(m) }
func (*ReadGroup_Program) ProtoMessage()    {}
func (*ReadGroup_Program) Descriptor() ([]byte, []int) {
	return fileDescriptor_readgroup_d1ee37f21a1efad6, []int{0, 1}
}
func (m *ReadGroup_Program) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadGroup_Program.Unmarshal(m, b)
}
func (m *ReadGroup_Program) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadGroup_Program.Marshal(b, m, deterministic)
}
func (dst *ReadGroup_Program) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadGroup_Program.Merge(dst, src)
}
func (m *ReadGroup_Program) XXX_Size() int {
	return xxx_messageInfo_ReadGroup_Program.Size(m)
}
func (m *ReadGroup_Program) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadGroup_Program.DiscardUnknown(m)
}

var xxx_messageInfo_ReadGroup_Program proto.InternalMessageInfo

func (m *ReadGroup_Program) GetCommandLine() string {
	if m != nil {
		return m.CommandLine
	}
	return ""
}

func (m *ReadGroup_Program) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ReadGroup_Program) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReadGroup_Program) GetPrevProgramId() string {
	if m != nil {
		return m.PrevProgramId
	}
	return ""
}

func (m *ReadGroup_Program) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterType((*ReadGroup)(nil), "google.genomics.v1.ReadGroup")
	proto.RegisterMapType((map[string]*_struct.ListValue)(nil), "google.genomics.v1.ReadGroup.InfoEntry")
	proto.RegisterType((*ReadGroup_Experiment)(nil), "google.genomics.v1.ReadGroup.Experiment")
	proto.RegisterType((*ReadGroup_Program)(nil), "google.genomics.v1.ReadGroup.Program")
}

func init() {
	proto.RegisterFile("google/genomics/v1/readgroup.proto", fileDescriptor_readgroup_d1ee37f21a1efad6)
}

var fileDescriptor_readgroup_d1ee37f21a1efad6 = []byte{
	// 585 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xcb, 0x6e, 0xd4, 0x30,
	0x14, 0x55, 0xa6, 0xcf, 0xb9, 0xd3, 0xc7, 0x60, 0x04, 0x8a, 0x06, 0x90, 0x86, 0x22, 0x60, 0x10,
	0x52, 0x42, 0x87, 0x0d, 0x6a, 0x57, 0x14, 0x55, 0x10, 0xa9, 0x48, 0x55, 0x2a, 0x58, 0xb0, 0x89,
	0xdc, 0xf8, 0x4e, 0x64, 0x91, 0xd8, 0xc1, 0x76, 0x46, 0xb4, 0x9f, 0xc1, 0x57, 0xf0, 0x2d, 0x7c,
	0x11, 0x4b, 0x64, 0xc7, 0x49, 0x47, 0xa2, 0xea, 0xce, 0x39, 0xe7, 0x5c, 0xdf, 0xc7, 0xb9, 0x0e,
	0x1c, 0x14, 0x52, 0x16, 0x25, 0xc6, 0x05, 0x0a, 0x59, 0xf1, 0x5c, 0xc7, 0xcb, 0xc3, 0x58, 0x21,
	0x65, 0x85, 0x92, 0x4d, 0x1d, 0xd5, 0x4a, 0x1a, 0x49, 0x48, 0xab, 0x89, 0x3a, 0x4d, 0xb4, 0x3c,
	0x9c, 0x3c, 0xf6, 0x71, 0xb4, 0xe6, 0x31, 0x15, 0x42, 0x1a, 0x6a, 0xb8, 0x14, 0xba, 0x8d, 0xe8,
	0x59, 0xf7, 0x75, 0xd9, 0x2c, 0x62, 0x6d, 0x54, 0x93, 0x9b, 0x96, 0x3d, 0xf8, 0xb3, 0x09, 0xc3,
	0x14, 0x29, 0xfb, 0x68, 0x73, 0x90, 0x3d, 0x18, 0x70, 0x16, 0x06, 0xd3, 0x60, 0x36, 0x4c, 0x07,
	0x9c, 0x91, 0x27, 0x00, 0x8c, 0x1a, 0xaa, 0xd1, 0x64, 0x9c, 0x85, 0x03, 0x87, 0x0f, 0x3d, 0x92,
	0x30, 0x42, 0x60, 0x5d, 0xd0, 0x0a, 0xc3, 0x35, 0x47, 0xb8, 0x33, 0x99, 0xc2, 0x88, 0xa1, 0xce,
	0x15, 0xaf, 0x6d, 0x11, 0xe1, 0xba, 0xa3, 0x56, 0x21, 0xf2, 0x08, 0x86, 0x9a, 0x56, 0x75, 0x89,
	0xf6, 0xce, 0x0d, 0xc7, 0x6f, 0xb7, 0x40, 0xc2, 0xc8, 0x27, 0x00, 0xfc, 0x59, 0xa3, 0xe2, 0x15,
	0x0a, 0x13, 0x6e, 0x4e, 0x83, 0xd9, 0x68, 0x3e, 0x8b, 0xfe, 0x6f, 0x3a, 0xea, 0x8b, 0x8e, 0x4e,
	0x7b, 0x7d, 0xba, 0x12, 0x4b, 0xe6, 0xf0, 0xa0, 0x56, 0xc8, 0x78, 0x6e, 0x90, 0x65, 0x5c, 0x68,
	0x54, 0x26, 0xd3, 0xfc, 0x1a, 0xc3, 0xad, 0x69, 0x30, 0xdb, 0x48, 0xef, 0xf7, 0x64, 0xe2, 0xb8,
	0x0b, 0x7e, 0x8d, 0xe4, 0x3d, 0x6c, 0xd7, 0x4a, 0x16, 0x8a, 0x56, 0x3a, 0x84, 0xe9, 0xda, 0x6c,
	0x34, 0x7f, 0x7e, 0x77, 0xee, 0xf3, 0x56, 0x9d, 0xf6, 0x61, 0x64, 0x06, 0x63, 0x85, 0x0b, 0x54,
	0x28, 0x72, 0xcc, 0xfc, 0xe0, 0x46, 0xae, 0xc9, 0xbd, 0x1e, 0xbf, 0x70, 0xd3, 0x3b, 0x86, 0x75,
	0x2e, 0x16, 0x32, 0xdc, 0x71, 0x89, 0x5e, 0xde, 0x9d, 0x28, 0x11, 0x0b, 0x79, 0x2a, 0x8c, 0xba,
	0x4a, 0x5d, 0xd0, 0xe4, 0x77, 0x00, 0x70, 0xd3, 0xb8, 0x35, 0xaa, 0xe4, 0x97, 0x8a, 0xaa, 0xab,
	0xac, 0x37, 0x70, 0xe8, 0x91, 0x84, 0x91, 0x67, 0xb0, 0x5b, 0x97, 0xd4, 0x2c, 0xa4, 0xaa, 0xb2,
	0x46, 0x70, 0xe3, 0xad, 0xdc, 0xe9, 0xc0, 0x2f, 0x82, 0x1b, 0xf2, 0x1a, 0xee, 0x69, 0xfc, 0xd1,
	0xa0, 0xc8, 0xb9, 0x28, 0xb2, 0x1c, 0x85, 0x41, 0xe5, 0xad, 0x1d, 0xdf, 0x10, 0x1f, 0x1c, 0x4e,
	0x5e, 0xc1, 0x98, 0x0b, 0xbb, 0x49, 0x36, 0x7d, 0x56, 0x49, 0x86, 0xa5, 0xf7, 0x7a, 0xff, 0x06,
	0xff, 0x6c, 0xe1, 0xc9, 0xaf, 0x00, 0xb6, 0xfc, 0x9c, 0xc8, 0x53, 0xd8, 0xc9, 0x65, 0x55, 0x51,
	0xc1, 0xb2, 0x92, 0x0b, 0xf4, 0x95, 0x8e, 0x3c, 0x76, 0xc6, 0x05, 0xfa, 0x1d, 0x1c, 0xf4, 0x3b,
	0x78, 0xdb, 0x92, 0xbd, 0x80, 0xfd, 0x5a, 0xe1, 0x32, 0xf3, 0x53, 0xb7, 0x3d, 0xb7, 0xc9, 0x77,
	0x2d, 0xec, 0x93, 0x25, 0x8c, 0x84, 0xb0, 0xb5, 0x44, 0xa5, 0xed, 0x22, 0xb6, 0x8b, 0xd6, 0x7d,
	0x4e, 0x2e, 0x60, 0xd8, 0x8f, 0x94, 0x8c, 0x61, 0xed, 0x3b, 0x5e, 0xf9, 0x62, 0xec, 0x91, 0xbc,
	0x81, 0x8d, 0x25, 0x2d, 0x1b, 0x74, 0x75, 0x8c, 0xe6, 0x93, 0xce, 0x9c, 0xee, 0x11, 0x45, 0x67,
	0x5c, 0x9b, 0xaf, 0x56, 0x91, 0xb6, 0xc2, 0xa3, 0xc1, 0xbb, 0xe0, 0x84, 0xc3, 0xc3, 0x5c, 0x56,
	0xb7, 0x18, 0x79, 0xb2, 0xd7, 0x3b, 0x79, 0x6e, 0x6f, 0x38, 0x0f, 0xbe, 0x1d, 0x75, 0x2a, 0x59,
	0x52, 0x51, 0x44, 0x52, 0x15, 0xf6, 0xdd, 0xbb, 0xfb, 0xe3, 0x96, 0xa2, 0x35, 0xd7, 0xab, 0xff,
	0x82, 0xe3, 0xee, 0xfc, 0x37, 0x08, 0x2e, 0x37, 0x9d, 0xf2, 0xed, 0xbf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x37, 0xed, 0xaa, 0xaa, 0x34, 0x04, 0x00, 0x00,
}
