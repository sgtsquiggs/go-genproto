// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/containeranalysis/v1alpha1/provenance.proto

package containeranalysis // import "github.com/sgtsquiggs/go-genproto/googleapis/devtools/containeranalysis/v1alpha1"

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

// Specifies the hash algorithm, if any.
type Hash_HashType int32

const (
	// No hash requested.
	Hash_NONE Hash_HashType = 0
	// A sha256 hash.
	Hash_SHA256 Hash_HashType = 1
)

var Hash_HashType_name = map[int32]string{
	0: "NONE",
	1: "SHA256",
}
var Hash_HashType_value = map[string]int32{
	"NONE":   0,
	"SHA256": 1,
}

func (x Hash_HashType) String() string {
	return proto.EnumName(Hash_HashType_name, int32(x))
}
func (Hash_HashType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_provenance_ae080c06fe939617, []int{3, 0}
}

// Provenance of a build. Contains all information needed to verify the full
// details about the build from source to completion.
type BuildProvenance struct {
	// Unique identifier of the build.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the project.
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// Commands requested by the build.
	Commands []*Command `protobuf:"bytes,5,rep,name=commands,proto3" json:"commands,omitempty"`
	// Output of the build.
	BuiltArtifacts []*Artifact `protobuf:"bytes,6,rep,name=built_artifacts,json=builtArtifacts,proto3" json:"built_artifacts,omitempty"`
	// Time at which the build was created.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,7,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Time at which execution of the build was started.
	StartTime *timestamp.Timestamp `protobuf:"bytes,8,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Time at which execution of the build was finished.
	FinishTime *timestamp.Timestamp `protobuf:"bytes,9,opt,name=finish_time,json=finishTime,proto3" json:"finish_time,omitempty"`
	// E-mail address of the user who initiated this build. Note that this was the
	// user's e-mail address at the time the build was initiated; this address may
	// not represent the same end-user for all time.
	Creator string `protobuf:"bytes,11,opt,name=creator,proto3" json:"creator,omitempty"`
	// Google Cloud Storage bucket where logs were written.
	LogsBucket string `protobuf:"bytes,13,opt,name=logs_bucket,json=logsBucket,proto3" json:"logs_bucket,omitempty"`
	// Details of the Source input to the build.
	SourceProvenance *Source `protobuf:"bytes,14,opt,name=source_provenance,json=sourceProvenance,proto3" json:"source_provenance,omitempty"`
	// Trigger identifier if the build was triggered automatically; empty if not.
	TriggerId string `protobuf:"bytes,15,opt,name=trigger_id,json=triggerId,proto3" json:"trigger_id,omitempty"`
	// Special options applied to this build. This is a catch-all field where
	// build providers can enter any desired additional details.
	BuildOptions map[string]string `protobuf:"bytes,16,rep,name=build_options,json=buildOptions,proto3" json:"build_options,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Version string of the builder at the time this build was executed.
	BuilderVersion       string   `protobuf:"bytes,17,opt,name=builder_version,json=builderVersion,proto3" json:"builder_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildProvenance) Reset()         { *m = BuildProvenance{} }
func (m *BuildProvenance) String() string { return proto.CompactTextString(m) }
func (*BuildProvenance) ProtoMessage()    {}
func (*BuildProvenance) Descriptor() ([]byte, []int) {
	return fileDescriptor_provenance_ae080c06fe939617, []int{0}
}
func (m *BuildProvenance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildProvenance.Unmarshal(m, b)
}
func (m *BuildProvenance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildProvenance.Marshal(b, m, deterministic)
}
func (dst *BuildProvenance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildProvenance.Merge(dst, src)
}
func (m *BuildProvenance) XXX_Size() int {
	return xxx_messageInfo_BuildProvenance.Size(m)
}
func (m *BuildProvenance) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildProvenance.DiscardUnknown(m)
}

var xxx_messageInfo_BuildProvenance proto.InternalMessageInfo

func (m *BuildProvenance) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *BuildProvenance) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *BuildProvenance) GetCommands() []*Command {
	if m != nil {
		return m.Commands
	}
	return nil
}

func (m *BuildProvenance) GetBuiltArtifacts() []*Artifact {
	if m != nil {
		return m.BuiltArtifacts
	}
	return nil
}

func (m *BuildProvenance) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *BuildProvenance) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *BuildProvenance) GetFinishTime() *timestamp.Timestamp {
	if m != nil {
		return m.FinishTime
	}
	return nil
}

func (m *BuildProvenance) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *BuildProvenance) GetLogsBucket() string {
	if m != nil {
		return m.LogsBucket
	}
	return ""
}

func (m *BuildProvenance) GetSourceProvenance() *Source {
	if m != nil {
		return m.SourceProvenance
	}
	return nil
}

func (m *BuildProvenance) GetTriggerId() string {
	if m != nil {
		return m.TriggerId
	}
	return ""
}

func (m *BuildProvenance) GetBuildOptions() map[string]string {
	if m != nil {
		return m.BuildOptions
	}
	return nil
}

func (m *BuildProvenance) GetBuilderVersion() string {
	if m != nil {
		return m.BuilderVersion
	}
	return ""
}

// Source describes the location of the source used for the build.
type Source struct {
	// Source location information.
	//
	// Types that are valid to be assigned to Source:
	//	*Source_StorageSource
	//	*Source_RepoSource
	Source isSource_Source `protobuf_oneof:"source"`
	// If provided, the input binary artifacts for the build came from this
	// location.
	ArtifactStorageSource *StorageSource `protobuf:"bytes,4,opt,name=artifact_storage_source,json=artifactStorageSource,proto3" json:"artifact_storage_source,omitempty"`
	// Hash(es) of the build source, which can be used to verify that the original
	// source integrity was maintained in the build.
	//
	// The keys to this map are file paths used as build source and the values
	// contain the hash values for those files.
	//
	// If the build source came in a single package such as a gzipped tarfile
	// (.tar.gz), the FileHash will be for the single path to that file.
	FileHashes map[string]*FileHashes `protobuf:"bytes,3,rep,name=file_hashes,json=fileHashes,proto3" json:"file_hashes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// If provided, the source code used for the build came from this location.
	Context *SourceContext `protobuf:"bytes,7,opt,name=context,proto3" json:"context,omitempty"`
	// If provided, some of the source code used for the build may be found in
	// these locations, in the case where the source repository had multiple
	// remotes or submodules. This list will not include the context specified in
	// the context field.
	AdditionalContexts   []*SourceContext `protobuf:"bytes,8,rep,name=additional_contexts,json=additionalContexts,proto3" json:"additional_contexts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Source) Reset()         { *m = Source{} }
func (m *Source) String() string { return proto.CompactTextString(m) }
func (*Source) ProtoMessage()    {}
func (*Source) Descriptor() ([]byte, []int) {
	return fileDescriptor_provenance_ae080c06fe939617, []int{1}
}
func (m *Source) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Source.Unmarshal(m, b)
}
func (m *Source) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Source.Marshal(b, m, deterministic)
}
func (dst *Source) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Source.Merge(dst, src)
}
func (m *Source) XXX_Size() int {
	return xxx_messageInfo_Source.Size(m)
}
func (m *Source) XXX_DiscardUnknown() {
	xxx_messageInfo_Source.DiscardUnknown(m)
}

var xxx_messageInfo_Source proto.InternalMessageInfo

type isSource_Source interface {
	isSource_Source()
}

type Source_StorageSource struct {
	StorageSource *StorageSource `protobuf:"bytes,1,opt,name=storage_source,json=storageSource,proto3,oneof"`
}

type Source_RepoSource struct {
	RepoSource *RepoSource `protobuf:"bytes,2,opt,name=repo_source,json=repoSource,proto3,oneof"`
}

func (*Source_StorageSource) isSource_Source() {}

func (*Source_RepoSource) isSource_Source() {}

func (m *Source) GetSource() isSource_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *Source) GetStorageSource() *StorageSource {
	if x, ok := m.GetSource().(*Source_StorageSource); ok {
		return x.StorageSource
	}
	return nil
}

func (m *Source) GetRepoSource() *RepoSource {
	if x, ok := m.GetSource().(*Source_RepoSource); ok {
		return x.RepoSource
	}
	return nil
}

func (m *Source) GetArtifactStorageSource() *StorageSource {
	if m != nil {
		return m.ArtifactStorageSource
	}
	return nil
}

func (m *Source) GetFileHashes() map[string]*FileHashes {
	if m != nil {
		return m.FileHashes
	}
	return nil
}

func (m *Source) GetContext() *SourceContext {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *Source) GetAdditionalContexts() []*SourceContext {
	if m != nil {
		return m.AdditionalContexts
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Source) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Source_OneofMarshaler, _Source_OneofUnmarshaler, _Source_OneofSizer, []interface{}{
		(*Source_StorageSource)(nil),
		(*Source_RepoSource)(nil),
	}
}

func _Source_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Source)
	// source
	switch x := m.Source.(type) {
	case *Source_StorageSource:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.StorageSource); err != nil {
			return err
		}
	case *Source_RepoSource:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RepoSource); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Source.Source has unexpected type %T", x)
	}
	return nil
}

func _Source_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Source)
	switch tag {
	case 1: // source.storage_source
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StorageSource)
		err := b.DecodeMessage(msg)
		m.Source = &Source_StorageSource{msg}
		return true, err
	case 2: // source.repo_source
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RepoSource)
		err := b.DecodeMessage(msg)
		m.Source = &Source_RepoSource{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Source_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Source)
	// source
	switch x := m.Source.(type) {
	case *Source_StorageSource:
		s := proto.Size(x.StorageSource)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Source_RepoSource:
		s := proto.Size(x.RepoSource)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Container message for hashes of byte content of files, used in Source
// messages to verify integrity of source input to the build.
type FileHashes struct {
	// Collection of file hashes.
	FileHash             []*Hash  `protobuf:"bytes,1,rep,name=file_hash,json=fileHash,proto3" json:"file_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileHashes) Reset()         { *m = FileHashes{} }
func (m *FileHashes) String() string { return proto.CompactTextString(m) }
func (*FileHashes) ProtoMessage()    {}
func (*FileHashes) Descriptor() ([]byte, []int) {
	return fileDescriptor_provenance_ae080c06fe939617, []int{2}
}
func (m *FileHashes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileHashes.Unmarshal(m, b)
}
func (m *FileHashes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileHashes.Marshal(b, m, deterministic)
}
func (dst *FileHashes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileHashes.Merge(dst, src)
}
func (m *FileHashes) XXX_Size() int {
	return xxx_messageInfo_FileHashes.Size(m)
}
func (m *FileHashes) XXX_DiscardUnknown() {
	xxx_messageInfo_FileHashes.DiscardUnknown(m)
}

var xxx_messageInfo_FileHashes proto.InternalMessageInfo

func (m *FileHashes) GetFileHash() []*Hash {
	if m != nil {
		return m.FileHash
	}
	return nil
}

// Container message for hash values.
type Hash struct {
	// The type of hash that was performed.
	Type Hash_HashType `protobuf:"varint,1,opt,name=type,proto3,enum=google.devtools.containeranalysis.v1alpha1.Hash_HashType" json:"type,omitempty"`
	// The hash value.
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hash) Reset()         { *m = Hash{} }
func (m *Hash) String() string { return proto.CompactTextString(m) }
func (*Hash) ProtoMessage()    {}
func (*Hash) Descriptor() ([]byte, []int) {
	return fileDescriptor_provenance_ae080c06fe939617, []int{3}
}
func (m *Hash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hash.Unmarshal(m, b)
}
func (m *Hash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hash.Marshal(b, m, deterministic)
}
func (dst *Hash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hash.Merge(dst, src)
}
func (m *Hash) XXX_Size() int {
	return xxx_messageInfo_Hash.Size(m)
}
func (m *Hash) XXX_DiscardUnknown() {
	xxx_messageInfo_Hash.DiscardUnknown(m)
}

var xxx_messageInfo_Hash proto.InternalMessageInfo

func (m *Hash) GetType() Hash_HashType {
	if m != nil {
		return m.Type
	}
	return Hash_NONE
}

func (m *Hash) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

// StorageSource describes the location of the source in an archive file in
// Google Cloud Storage.
type StorageSource struct {
	// Google Cloud Storage bucket containing source (see [Bucket Name
	// Requirements]
	// (https://cloud.google.com/storage/docs/bucket-naming#requirements)).
	Bucket string `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	// Google Cloud Storage object containing source.
	Object string `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
	// Google Cloud Storage generation for the object.
	Generation           int64    `protobuf:"varint,3,opt,name=generation,proto3" json:"generation,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StorageSource) Reset()         { *m = StorageSource{} }
func (m *StorageSource) String() string { return proto.CompactTextString(m) }
func (*StorageSource) ProtoMessage()    {}
func (*StorageSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_provenance_ae080c06fe939617, []int{4}
}
func (m *StorageSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StorageSource.Unmarshal(m, b)
}
func (m *StorageSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StorageSource.Marshal(b, m, deterministic)
}
func (dst *StorageSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorageSource.Merge(dst, src)
}
func (m *StorageSource) XXX_Size() int {
	return xxx_messageInfo_StorageSource.Size(m)
}
func (m *StorageSource) XXX_DiscardUnknown() {
	xxx_messageInfo_StorageSource.DiscardUnknown(m)
}

var xxx_messageInfo_StorageSource proto.InternalMessageInfo

func (m *StorageSource) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *StorageSource) GetObject() string {
	if m != nil {
		return m.Object
	}
	return ""
}

func (m *StorageSource) GetGeneration() int64 {
	if m != nil {
		return m.Generation
	}
	return 0
}

// RepoSource describes the location of the source in a Google Cloud Source
// Repository.
type RepoSource struct {
	// ID of the project that owns the repo.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// Name of the repo.
	RepoName string `protobuf:"bytes,2,opt,name=repo_name,json=repoName,proto3" json:"repo_name,omitempty"`
	// A revision within the source repository must be specified in
	// one of these ways.
	//
	// Types that are valid to be assigned to Revision:
	//	*RepoSource_BranchName
	//	*RepoSource_TagName
	//	*RepoSource_CommitSha
	Revision             isRepoSource_Revision `protobuf_oneof:"revision"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *RepoSource) Reset()         { *m = RepoSource{} }
func (m *RepoSource) String() string { return proto.CompactTextString(m) }
func (*RepoSource) ProtoMessage()    {}
func (*RepoSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_provenance_ae080c06fe939617, []int{5}
}
func (m *RepoSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RepoSource.Unmarshal(m, b)
}
func (m *RepoSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RepoSource.Marshal(b, m, deterministic)
}
func (dst *RepoSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepoSource.Merge(dst, src)
}
func (m *RepoSource) XXX_Size() int {
	return xxx_messageInfo_RepoSource.Size(m)
}
func (m *RepoSource) XXX_DiscardUnknown() {
	xxx_messageInfo_RepoSource.DiscardUnknown(m)
}

var xxx_messageInfo_RepoSource proto.InternalMessageInfo

func (m *RepoSource) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *RepoSource) GetRepoName() string {
	if m != nil {
		return m.RepoName
	}
	return ""
}

type isRepoSource_Revision interface {
	isRepoSource_Revision()
}

type RepoSource_BranchName struct {
	BranchName string `protobuf:"bytes,3,opt,name=branch_name,json=branchName,proto3,oneof"`
}

type RepoSource_TagName struct {
	TagName string `protobuf:"bytes,4,opt,name=tag_name,json=tagName,proto3,oneof"`
}

type RepoSource_CommitSha struct {
	CommitSha string `protobuf:"bytes,5,opt,name=commit_sha,json=commitSha,proto3,oneof"`
}

func (*RepoSource_BranchName) isRepoSource_Revision() {}

func (*RepoSource_TagName) isRepoSource_Revision() {}

func (*RepoSource_CommitSha) isRepoSource_Revision() {}

func (m *RepoSource) GetRevision() isRepoSource_Revision {
	if m != nil {
		return m.Revision
	}
	return nil
}

func (m *RepoSource) GetBranchName() string {
	if x, ok := m.GetRevision().(*RepoSource_BranchName); ok {
		return x.BranchName
	}
	return ""
}

func (m *RepoSource) GetTagName() string {
	if x, ok := m.GetRevision().(*RepoSource_TagName); ok {
		return x.TagName
	}
	return ""
}

func (m *RepoSource) GetCommitSha() string {
	if x, ok := m.GetRevision().(*RepoSource_CommitSha); ok {
		return x.CommitSha
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*RepoSource) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _RepoSource_OneofMarshaler, _RepoSource_OneofUnmarshaler, _RepoSource_OneofSizer, []interface{}{
		(*RepoSource_BranchName)(nil),
		(*RepoSource_TagName)(nil),
		(*RepoSource_CommitSha)(nil),
	}
}

func _RepoSource_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*RepoSource)
	// revision
	switch x := m.Revision.(type) {
	case *RepoSource_BranchName:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.BranchName)
	case *RepoSource_TagName:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.TagName)
	case *RepoSource_CommitSha:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.CommitSha)
	case nil:
	default:
		return fmt.Errorf("RepoSource.Revision has unexpected type %T", x)
	}
	return nil
}

func _RepoSource_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*RepoSource)
	switch tag {
	case 3: // revision.branch_name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Revision = &RepoSource_BranchName{x}
		return true, err
	case 4: // revision.tag_name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Revision = &RepoSource_TagName{x}
		return true, err
	case 5: // revision.commit_sha
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Revision = &RepoSource_CommitSha{x}
		return true, err
	default:
		return false, nil
	}
}

func _RepoSource_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*RepoSource)
	// revision
	switch x := m.Revision.(type) {
	case *RepoSource_BranchName:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.BranchName)))
		n += len(x.BranchName)
	case *RepoSource_TagName:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.TagName)))
		n += len(x.TagName)
	case *RepoSource_CommitSha:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.CommitSha)))
		n += len(x.CommitSha)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Command describes a step performed as part of the build pipeline.
type Command struct {
	// Name of the command, as presented on the command line, or if the command is
	// packaged as a Docker container, as presented to `docker pull`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Environment variables set before running this Command.
	Env []string `protobuf:"bytes,2,rep,name=env,proto3" json:"env,omitempty"`
	// Command-line arguments used when executing this Command.
	Args []string `protobuf:"bytes,3,rep,name=args,proto3" json:"args,omitempty"`
	// Working directory (relative to project source root) used when running
	// this Command.
	Dir string `protobuf:"bytes,4,opt,name=dir,proto3" json:"dir,omitempty"`
	// Optional unique identifier for this Command, used in wait_for to reference
	// this Command as a dependency.
	Id string `protobuf:"bytes,5,opt,name=id,proto3" json:"id,omitempty"`
	// The ID(s) of the Command(s) that this Command depends on.
	WaitFor              []string `protobuf:"bytes,6,rep,name=wait_for,json=waitFor,proto3" json:"wait_for,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}
func (*Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_provenance_ae080c06fe939617, []int{6}
}
func (m *Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command.Unmarshal(m, b)
}
func (m *Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command.Marshal(b, m, deterministic)
}
func (dst *Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command.Merge(dst, src)
}
func (m *Command) XXX_Size() int {
	return xxx_messageInfo_Command.Size(m)
}
func (m *Command) XXX_DiscardUnknown() {
	xxx_messageInfo_Command.DiscardUnknown(m)
}

var xxx_messageInfo_Command proto.InternalMessageInfo

func (m *Command) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Command) GetEnv() []string {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *Command) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Command) GetDir() string {
	if m != nil {
		return m.Dir
	}
	return ""
}

func (m *Command) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Command) GetWaitFor() []string {
	if m != nil {
		return m.WaitFor
	}
	return nil
}

// Artifact describes a build product.
type Artifact struct {
	// Name of the artifact. This may be the path to a binary or jar file, or in
	// the case of a container build, the name used to push the container image to
	// Google Container Registry, as presented to `docker push`.
	//
	// This field is deprecated in favor of the plural `names` field; it continues
	// to exist here to allow existing BuildProvenance serialized to json in
	// google.devtools.containeranalysis.v1alpha1.BuildDetails.provenance_bytes to
	// deserialize back into proto.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Hash or checksum value of a binary, or Docker Registry 2.0 digest of a
	// container.
	Checksum string `protobuf:"bytes,2,opt,name=checksum,proto3" json:"checksum,omitempty"`
	// Artifact ID, if any; for container images, this will be a URL by digest
	// like gcr.io/projectID/imagename@sha256:123456
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// Related artifact names. This may be the path to a binary or jar file, or in
	// the case of a container build, the name used to push the container image to
	// Google Container Registry, as presented to `docker push`. Note that a
	// single Artifact ID can have multiple names, for example if two tags are
	// applied to one image.
	Names                []string `protobuf:"bytes,4,rep,name=names,proto3" json:"names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Artifact) Reset()         { *m = Artifact{} }
func (m *Artifact) String() string { return proto.CompactTextString(m) }
func (*Artifact) ProtoMessage()    {}
func (*Artifact) Descriptor() ([]byte, []int) {
	return fileDescriptor_provenance_ae080c06fe939617, []int{7}
}
func (m *Artifact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Artifact.Unmarshal(m, b)
}
func (m *Artifact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Artifact.Marshal(b, m, deterministic)
}
func (dst *Artifact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Artifact.Merge(dst, src)
}
func (m *Artifact) XXX_Size() int {
	return xxx_messageInfo_Artifact.Size(m)
}
func (m *Artifact) XXX_DiscardUnknown() {
	xxx_messageInfo_Artifact.DiscardUnknown(m)
}

var xxx_messageInfo_Artifact proto.InternalMessageInfo

func (m *Artifact) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Artifact) GetChecksum() string {
	if m != nil {
		return m.Checksum
	}
	return ""
}

func (m *Artifact) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Artifact) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func init() {
	proto.RegisterType((*BuildProvenance)(nil), "google.devtools.containeranalysis.v1alpha1.BuildProvenance")
	proto.RegisterMapType((map[string]string)(nil), "google.devtools.containeranalysis.v1alpha1.BuildProvenance.BuildOptionsEntry")
	proto.RegisterType((*Source)(nil), "google.devtools.containeranalysis.v1alpha1.Source")
	proto.RegisterMapType((map[string]*FileHashes)(nil), "google.devtools.containeranalysis.v1alpha1.Source.FileHashesEntry")
	proto.RegisterType((*FileHashes)(nil), "google.devtools.containeranalysis.v1alpha1.FileHashes")
	proto.RegisterType((*Hash)(nil), "google.devtools.containeranalysis.v1alpha1.Hash")
	proto.RegisterType((*StorageSource)(nil), "google.devtools.containeranalysis.v1alpha1.StorageSource")
	proto.RegisterType((*RepoSource)(nil), "google.devtools.containeranalysis.v1alpha1.RepoSource")
	proto.RegisterType((*Command)(nil), "google.devtools.containeranalysis.v1alpha1.Command")
	proto.RegisterType((*Artifact)(nil), "google.devtools.containeranalysis.v1alpha1.Artifact")
	proto.RegisterEnum("google.devtools.containeranalysis.v1alpha1.Hash_HashType", Hash_HashType_name, Hash_HashType_value)
}

func init() {
	proto.RegisterFile("google/devtools/containeranalysis/v1alpha1/provenance.proto", fileDescriptor_provenance_ae080c06fe939617)
}

var fileDescriptor_provenance_ae080c06fe939617 = []byte{
	// 1026 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0xee, 0xfa, 0x77, 0xf7, 0xb8, 0x71, 0x92, 0xa1, 0xc0, 0xe2, 0x52, 0x62, 0x2c, 0x21, 0x22,
	0x2e, 0x6c, 0x9a, 0x42, 0x45, 0xc8, 0x45, 0x15, 0x47, 0x2d, 0xa9, 0x44, 0x92, 0x6a, 0x53, 0x21,
	0x41, 0x85, 0x96, 0xf1, 0xee, 0x78, 0x3d, 0xcd, 0x7a, 0x67, 0x99, 0x19, 0x1b, 0x7c, 0xc7, 0x03,
	0xc0, 0x4b, 0xf0, 0x0e, 0xdc, 0xf2, 0x2e, 0xbc, 0x09, 0x9a, 0x9f, 0xb5, 0x1d, 0xa7, 0x55, 0xbb,
	0xea, 0xcd, 0x6a, 0xce, 0x37, 0xe7, 0x7c, 0x67, 0xe6, 0xcc, 0x77, 0x66, 0x16, 0x8e, 0x12, 0xc6,
	0x92, 0x94, 0x0c, 0x62, 0x32, 0x97, 0x8c, 0xa5, 0x62, 0x10, 0xb1, 0x4c, 0x62, 0x9a, 0x11, 0x8e,
	0x33, 0x9c, 0x2e, 0x04, 0x15, 0x83, 0xf9, 0x7d, 0x9c, 0xe6, 0x13, 0x7c, 0x7f, 0x90, 0x73, 0x36,
	0x27, 0x19, 0xce, 0x22, 0xd2, 0xcf, 0x39, 0x93, 0x0c, 0x7d, 0x61, 0x82, 0xfb, 0x45, 0x70, 0xff,
	0x46, 0x70, 0xbf, 0x08, 0xee, 0x7c, 0x6c, 0x13, 0xe1, 0x9c, 0x0e, 0x70, 0x96, 0x31, 0x89, 0x25,
	0x65, 0x99, 0x30, 0x4c, 0x9d, 0x47, 0x25, 0x96, 0x21, 0xd8, 0x8c, 0x47, 0x24, 0x54, 0x1e, 0xe4,
	0x77, 0x69, 0x09, 0xf6, 0x2c, 0x81, 0xb6, 0x46, 0xb3, 0xf1, 0x40, 0xd2, 0x29, 0x11, 0x12, 0x4f,
	0x73, 0xe3, 0xd0, 0xfb, 0xb7, 0x01, 0xdb, 0xc3, 0x19, 0x4d, 0xe3, 0x67, 0xcb, 0x5d, 0xa0, 0x36,
	0x54, 0x68, 0xec, 0x3b, 0x5d, 0x67, 0xdf, 0x0b, 0x2a, 0x34, 0x46, 0xf7, 0x00, 0x72, 0xce, 0x5e,
	0x92, 0x48, 0x86, 0x34, 0xf6, 0x2b, 0x1a, 0xf7, 0x2c, 0xf2, 0x34, 0x46, 0x17, 0xe0, 0x46, 0x6c,
	0x3a, 0xc5, 0x59, 0x2c, 0xfc, 0x7a, 0xb7, 0xba, 0xdf, 0x3a, 0x78, 0xd0, 0x7f, 0xfb, 0x0a, 0xf4,
	0x4f, 0x4c, 0x6c, 0xb0, 0x24, 0x41, 0x3f, 0xc3, 0xf6, 0x68, 0x46, 0x53, 0x19, 0x62, 0x2e, 0xe9,
	0x18, 0x47, 0x52, 0xf8, 0x0d, 0xcd, 0xfb, 0x55, 0x19, 0xde, 0x63, 0x1b, 0x1c, 0xb4, 0x35, 0x59,
	0x61, 0x0a, 0x74, 0x04, 0xad, 0x88, 0x13, 0x2c, 0x49, 0xa8, 0x8a, 0xe1, 0x37, 0xbb, 0xce, 0x7e,
	0xeb, 0xa0, 0x53, 0x50, 0x17, 0x95, 0xea, 0x3f, 0x2f, 0x2a, 0x15, 0x80, 0x71, 0x57, 0x00, 0x3a,
	0x04, 0x10, 0x12, 0x73, 0x69, 0x62, 0xdd, 0x37, 0xc6, 0x7a, 0xda, 0x5b, 0x87, 0x1e, 0x41, 0x6b,
	0x4c, 0x33, 0x2a, 0x26, 0x26, 0xd6, 0x7b, 0x73, 0x5e, 0xe3, 0xae, 0x83, 0x7d, 0x68, 0xea, 0x55,
	0x30, 0xee, 0xb7, 0xf4, 0x01, 0x14, 0x26, 0xda, 0x83, 0x56, 0xca, 0x12, 0x11, 0x8e, 0x66, 0xd1,
	0x15, 0x91, 0xfe, 0x96, 0x9e, 0x05, 0x05, 0x0d, 0x35, 0x82, 0x42, 0xd8, 0xb5, 0xda, 0x58, 0x29,
	0xd5, 0x6f, 0xeb, 0xec, 0x07, 0x65, 0x0a, 0x7a, 0xa9, 0x49, 0x82, 0x1d, 0x43, 0xb6, 0xa6, 0x97,
	0x7b, 0x00, 0x92, 0xd3, 0x24, 0x21, 0x5c, 0xe9, 0x63, 0xdb, 0xe8, 0xc3, 0x22, 0x4f, 0x63, 0xc4,
	0x61, 0x4b, 0x9d, 0x40, 0x1c, 0xb2, 0x5c, 0x6b, 0xdb, 0xdf, 0xd1, 0x87, 0x79, 0x56, 0x26, 0xf7,
	0x86, 0x44, 0x8d, 0x7d, 0x61, 0xf8, 0x1e, 0x67, 0x92, 0x2f, 0x82, 0xdb, 0xa3, 0x35, 0x08, 0x7d,
	0x6e, 0x24, 0x14, 0x13, 0x1e, 0xce, 0x09, 0x17, 0x94, 0x65, 0xfe, 0xae, 0x5e, 0x57, 0xdb, 0xc2,
	0x3f, 0x18, 0xb4, 0xf3, 0x08, 0x76, 0x6f, 0x70, 0xa1, 0x1d, 0xa8, 0x5e, 0x91, 0x85, 0xed, 0x00,
	0x35, 0x44, 0x77, 0xa0, 0x3e, 0xc7, 0xe9, 0x8c, 0x58, 0xf5, 0x1b, 0xe3, 0xdb, 0xca, 0x37, 0x4e,
	0xef, 0xbf, 0x3a, 0x34, 0x4c, 0x65, 0xd0, 0x08, 0xda, 0x42, 0x32, 0x8e, 0x13, 0x12, 0x9a, 0x1a,
	0x69, 0x86, 0xd6, 0xc1, 0x61, 0xa9, 0x2a, 0x1b, 0x06, 0x43, 0x79, 0x7a, 0x2b, 0xd8, 0x12, 0xeb,
	0x00, 0xfa, 0x11, 0x5a, 0x9c, 0xe4, 0xac, 0x48, 0x50, 0xd1, 0x09, 0x1e, 0x96, 0x49, 0x10, 0x90,
	0x9c, 0x2d, 0xd9, 0x81, 0x2f, 0x2d, 0xf4, 0x2b, 0x7c, 0x58, 0x34, 0x5c, 0xb8, 0xb1, 0x8f, 0xda,
	0x3b, 0xee, 0x23, 0x78, 0xbf, 0x60, 0xbe, 0x06, 0xa3, 0x48, 0xb5, 0x44, 0x4a, 0xc2, 0x09, 0x16,
	0x13, 0x22, 0xfc, 0xaa, 0x16, 0xc6, 0xb0, 0xbc, 0x28, 0xfb, 0x4f, 0x68, 0x4a, 0x4e, 0x35, 0x89,
	0x51, 0x03, 0x8c, 0x97, 0x00, 0xba, 0x84, 0xa6, 0xbd, 0x14, 0x6d, 0xaf, 0x1f, 0x96, 0x4f, 0x70,
	0x62, 0x08, 0x82, 0x82, 0x09, 0xbd, 0x84, 0xf7, 0x70, 0x1c, 0x53, 0x25, 0x1a, 0x9c, 0x16, 0x97,
	0xae, 0xf0, 0x5d, 0xbd, 0x83, 0x77, 0x48, 0x80, 0x56, 0xac, 0x16, 0x12, 0x9d, 0x19, 0x6c, 0x6f,
	0xec, 0xef, 0x15, 0x0a, 0xfd, 0x7e, 0x5d, 0xa1, 0x25, 0x25, 0xb1, 0x62, 0x5f, 0x53, 0xf6, 0xd0,
	0x85, 0x86, 0x39, 0xfe, 0xde, 0x0b, 0x80, 0x95, 0x0b, 0x3a, 0x03, 0x6f, 0x79, 0x68, 0xbe, 0xa3,
	0x37, 0xfc, 0x65, 0x99, 0x6c, 0x8a, 0x26, 0x70, 0x8b, 0x03, 0xea, 0xfd, 0xe5, 0x40, 0x4d, 0x0d,
	0xd0, 0x19, 0xd4, 0xe4, 0x22, 0x37, 0x4d, 0xd3, 0x2e, 0x57, 0x43, 0x15, 0xaf, 0x3f, 0xcf, 0x17,
	0x39, 0x09, 0x34, 0xcd, 0xf5, 0x96, 0xbd, 0x6d, 0x37, 0xd6, 0xeb, 0x82, 0x5b, 0xf8, 0x21, 0x17,
	0x6a, 0xe7, 0x17, 0xe7, 0x8f, 0x77, 0x6e, 0x21, 0x80, 0xc6, 0xe5, 0xe9, 0xf1, 0xc1, 0xd7, 0x0f,
	0x77, 0x9c, 0x5e, 0x08, 0x5b, 0xd7, 0x45, 0xfa, 0x01, 0x34, 0xec, 0xdd, 0x6a, 0xca, 0x6d, 0x2d,
	0x85, 0xb3, 0x91, 0x7a, 0x03, 0xed, 0xa5, 0x60, 0x2d, 0xf4, 0x09, 0x40, 0x42, 0xd4, 0x32, 0xd5,
	0x31, 0xfa, 0xd5, 0xae, 0xb3, 0x5f, 0x0d, 0xd6, 0x90, 0xde, 0x3f, 0x0e, 0xc0, 0xaa, 0x09, 0x37,
	0x5e, 0x57, 0x67, 0xf3, 0x75, 0xbd, 0x0b, 0x9e, 0x6e, 0xf8, 0x0c, 0x4f, 0x8b, 0xdb, 0xc7, 0x55,
	0xc0, 0x39, 0x9e, 0x12, 0xf4, 0x29, 0xb4, 0x46, 0x1c, 0x67, 0xd1, 0xc4, 0x4c, 0xab, 0x5c, 0x9e,
	0xea, 0x6a, 0x03, 0x6a, 0x97, 0xbb, 0xe0, 0x4a, 0x9c, 0x98, 0xf9, 0x9a, 0x9d, 0x6f, 0x4a, 0x9c,
	0xe8, 0xc9, 0x3d, 0x00, 0xf5, 0xea, 0x52, 0x19, 0x8a, 0x09, 0xf6, 0xeb, 0x76, 0xda, 0x33, 0xd8,
	0xe5, 0x04, 0x0f, 0x01, 0x5c, 0x4e, 0xe6, 0x54, 0x5d, 0x95, 0xbd, 0x3f, 0x1c, 0x68, 0xda, 0xc7,
	0x1a, 0x21, 0xa8, 0x69, 0x46, 0xb3, 0x5c, 0x3d, 0x56, 0x9a, 0x24, 0xd9, 0xdc, 0xaf, 0x74, 0xab,
	0x4a, 0x93, 0x24, 0x9b, 0x2b, 0x2f, 0xcc, 0x13, 0xd3, 0xd7, 0x5e, 0xa0, 0xc7, 0xca, 0x2b, 0xa6,
	0xdc, 0x2c, 0x25, 0x50, 0x43, 0xfb, 0xbb, 0x51, 0x5f, 0xfe, 0x6e, 0x7c, 0x04, 0xee, 0x6f, 0x98,
	0xca, 0x70, 0xcc, 0xb8, 0x7e, 0xf7, 0xbd, 0xa0, 0xa9, 0xec, 0x27, 0x8c, 0xf7, 0x7e, 0x01, 0xb7,
	0x78, 0xc7, 0x5f, 0xb9, 0x84, 0x0e, 0xb8, 0xd1, 0x84, 0x44, 0x57, 0x62, 0x36, 0x2d, 0x6a, 0x55,
	0xd8, 0x36, 0x4d, 0x75, 0x99, 0xe6, 0x0e, 0xd4, 0x55, 0x8c, 0xf0, 0x6b, 0x3a, 0x87, 0x31, 0x86,
	0x7f, 0x3a, 0xf0, 0x59, 0xc4, 0xa6, 0x85, 0xf8, 0x5e, 0xaf, 0xb9, 0x67, 0xce, 0x4f, 0x2f, 0xac,
	0x53, 0xc2, 0x52, 0x9c, 0x25, 0x7d, 0xc6, 0x93, 0x41, 0x42, 0x32, 0xfd, 0x90, 0x0f, 0xcc, 0x14,
	0xce, 0xa9, 0x78, 0x9b, 0x9f, 0xb7, 0xa3, 0x1b, 0x53, 0x7f, 0x57, 0xaa, 0xdf, 0x9d, 0x1c, 0x8f,
	0x1a, 0x9a, 0xed, 0xc1, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x91, 0xec, 0x1b, 0x85, 0x90, 0x0a,
	0x00, 0x00,
}
