// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/vision/v1/text_annotation.proto

package vision // import "github.com/sgtsquiggs/go-genproto/googleapis/cloud/vision/v1"

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

// Enum to denote the type of break found. New line, space etc.
type TextAnnotation_DetectedBreak_BreakType int32

const (
	// Unknown break label type.
	TextAnnotation_DetectedBreak_UNKNOWN TextAnnotation_DetectedBreak_BreakType = 0
	// Regular space.
	TextAnnotation_DetectedBreak_SPACE TextAnnotation_DetectedBreak_BreakType = 1
	// Sure space (very wide).
	TextAnnotation_DetectedBreak_SURE_SPACE TextAnnotation_DetectedBreak_BreakType = 2
	// Line-wrapping break.
	TextAnnotation_DetectedBreak_EOL_SURE_SPACE TextAnnotation_DetectedBreak_BreakType = 3
	// End-line hyphen that is not present in text; does not co-occur with
	// `SPACE`, `LEADER_SPACE`, or `LINE_BREAK`.
	TextAnnotation_DetectedBreak_HYPHEN TextAnnotation_DetectedBreak_BreakType = 4
	// Line break that ends a paragraph.
	TextAnnotation_DetectedBreak_LINE_BREAK TextAnnotation_DetectedBreak_BreakType = 5
)

var TextAnnotation_DetectedBreak_BreakType_name = map[int32]string{
	0: "UNKNOWN",
	1: "SPACE",
	2: "SURE_SPACE",
	3: "EOL_SURE_SPACE",
	4: "HYPHEN",
	5: "LINE_BREAK",
}
var TextAnnotation_DetectedBreak_BreakType_value = map[string]int32{
	"UNKNOWN":        0,
	"SPACE":          1,
	"SURE_SPACE":     2,
	"EOL_SURE_SPACE": 3,
	"HYPHEN":         4,
	"LINE_BREAK":     5,
}

func (x TextAnnotation_DetectedBreak_BreakType) String() string {
	return proto.EnumName(TextAnnotation_DetectedBreak_BreakType_name, int32(x))
}
func (TextAnnotation_DetectedBreak_BreakType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_text_annotation_5ca1a16a28ef616a, []int{0, 1, 0}
}

// Type of a block (text, image etc) as identified by OCR.
type Block_BlockType int32

const (
	// Unknown block type.
	Block_UNKNOWN Block_BlockType = 0
	// Regular text block.
	Block_TEXT Block_BlockType = 1
	// Table block.
	Block_TABLE Block_BlockType = 2
	// Image block.
	Block_PICTURE Block_BlockType = 3
	// Horizontal/vertical line box.
	Block_RULER Block_BlockType = 4
	// Barcode block.
	Block_BARCODE Block_BlockType = 5
)

var Block_BlockType_name = map[int32]string{
	0: "UNKNOWN",
	1: "TEXT",
	2: "TABLE",
	3: "PICTURE",
	4: "RULER",
	5: "BARCODE",
}
var Block_BlockType_value = map[string]int32{
	"UNKNOWN": 0,
	"TEXT":    1,
	"TABLE":   2,
	"PICTURE": 3,
	"RULER":   4,
	"BARCODE": 5,
}

func (x Block_BlockType) String() string {
	return proto.EnumName(Block_BlockType_name, int32(x))
}
func (Block_BlockType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_text_annotation_5ca1a16a28ef616a, []int{2, 0}
}

// TextAnnotation contains a structured representation of OCR extracted text.
// The hierarchy of an OCR extracted text structure is like this:
//     TextAnnotation -> Page -> Block -> Paragraph -> Word -> Symbol
// Each structural component, starting from Page, may further have their own
// properties. Properties describe detected languages, breaks etc.. Please refer
// to the [TextAnnotation.TextProperty][google.cloud.vision.v1.TextAnnotation.TextProperty] message definition below for more
// detail.
type TextAnnotation struct {
	// List of pages detected by OCR.
	Pages []*Page `protobuf:"bytes,1,rep,name=pages,proto3" json:"pages,omitempty"`
	// UTF-8 text detected on the pages.
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextAnnotation) Reset()         { *m = TextAnnotation{} }
func (m *TextAnnotation) String() string { return proto.CompactTextString(m) }
func (*TextAnnotation) ProtoMessage()    {}
func (*TextAnnotation) Descriptor() ([]byte, []int) {
	return fileDescriptor_text_annotation_5ca1a16a28ef616a, []int{0}
}
func (m *TextAnnotation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextAnnotation.Unmarshal(m, b)
}
func (m *TextAnnotation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextAnnotation.Marshal(b, m, deterministic)
}
func (dst *TextAnnotation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextAnnotation.Merge(dst, src)
}
func (m *TextAnnotation) XXX_Size() int {
	return xxx_messageInfo_TextAnnotation.Size(m)
}
func (m *TextAnnotation) XXX_DiscardUnknown() {
	xxx_messageInfo_TextAnnotation.DiscardUnknown(m)
}

var xxx_messageInfo_TextAnnotation proto.InternalMessageInfo

func (m *TextAnnotation) GetPages() []*Page {
	if m != nil {
		return m.Pages
	}
	return nil
}

func (m *TextAnnotation) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

// Detected language for a structural component.
type TextAnnotation_DetectedLanguage struct {
	// The BCP-47 language code, such as "en-US" or "sr-Latn". For more
	// information, see
	// http://www.unicode.org/reports/tr35/#Unicode_locale_identifier.
	LanguageCode string `protobuf:"bytes,1,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	// Confidence of detected language. Range [0, 1].
	Confidence           float32  `protobuf:"fixed32,2,opt,name=confidence,proto3" json:"confidence,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextAnnotation_DetectedLanguage) Reset()         { *m = TextAnnotation_DetectedLanguage{} }
func (m *TextAnnotation_DetectedLanguage) String() string { return proto.CompactTextString(m) }
func (*TextAnnotation_DetectedLanguage) ProtoMessage()    {}
func (*TextAnnotation_DetectedLanguage) Descriptor() ([]byte, []int) {
	return fileDescriptor_text_annotation_5ca1a16a28ef616a, []int{0, 0}
}
func (m *TextAnnotation_DetectedLanguage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextAnnotation_DetectedLanguage.Unmarshal(m, b)
}
func (m *TextAnnotation_DetectedLanguage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextAnnotation_DetectedLanguage.Marshal(b, m, deterministic)
}
func (dst *TextAnnotation_DetectedLanguage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextAnnotation_DetectedLanguage.Merge(dst, src)
}
func (m *TextAnnotation_DetectedLanguage) XXX_Size() int {
	return xxx_messageInfo_TextAnnotation_DetectedLanguage.Size(m)
}
func (m *TextAnnotation_DetectedLanguage) XXX_DiscardUnknown() {
	xxx_messageInfo_TextAnnotation_DetectedLanguage.DiscardUnknown(m)
}

var xxx_messageInfo_TextAnnotation_DetectedLanguage proto.InternalMessageInfo

func (m *TextAnnotation_DetectedLanguage) GetLanguageCode() string {
	if m != nil {
		return m.LanguageCode
	}
	return ""
}

func (m *TextAnnotation_DetectedLanguage) GetConfidence() float32 {
	if m != nil {
		return m.Confidence
	}
	return 0
}

// Detected start or end of a structural component.
type TextAnnotation_DetectedBreak struct {
	// Detected break type.
	Type TextAnnotation_DetectedBreak_BreakType `protobuf:"varint,1,opt,name=type,proto3,enum=google.cloud.vision.v1.TextAnnotation_DetectedBreak_BreakType" json:"type,omitempty"`
	// True if break prepends the element.
	IsPrefix             bool     `protobuf:"varint,2,opt,name=is_prefix,json=isPrefix,proto3" json:"is_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextAnnotation_DetectedBreak) Reset()         { *m = TextAnnotation_DetectedBreak{} }
func (m *TextAnnotation_DetectedBreak) String() string { return proto.CompactTextString(m) }
func (*TextAnnotation_DetectedBreak) ProtoMessage()    {}
func (*TextAnnotation_DetectedBreak) Descriptor() ([]byte, []int) {
	return fileDescriptor_text_annotation_5ca1a16a28ef616a, []int{0, 1}
}
func (m *TextAnnotation_DetectedBreak) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextAnnotation_DetectedBreak.Unmarshal(m, b)
}
func (m *TextAnnotation_DetectedBreak) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextAnnotation_DetectedBreak.Marshal(b, m, deterministic)
}
func (dst *TextAnnotation_DetectedBreak) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextAnnotation_DetectedBreak.Merge(dst, src)
}
func (m *TextAnnotation_DetectedBreak) XXX_Size() int {
	return xxx_messageInfo_TextAnnotation_DetectedBreak.Size(m)
}
func (m *TextAnnotation_DetectedBreak) XXX_DiscardUnknown() {
	xxx_messageInfo_TextAnnotation_DetectedBreak.DiscardUnknown(m)
}

var xxx_messageInfo_TextAnnotation_DetectedBreak proto.InternalMessageInfo

func (m *TextAnnotation_DetectedBreak) GetType() TextAnnotation_DetectedBreak_BreakType {
	if m != nil {
		return m.Type
	}
	return TextAnnotation_DetectedBreak_UNKNOWN
}

func (m *TextAnnotation_DetectedBreak) GetIsPrefix() bool {
	if m != nil {
		return m.IsPrefix
	}
	return false
}

// Additional information detected on the structural component.
type TextAnnotation_TextProperty struct {
	// A list of detected languages together with confidence.
	DetectedLanguages []*TextAnnotation_DetectedLanguage `protobuf:"bytes,1,rep,name=detected_languages,json=detectedLanguages,proto3" json:"detected_languages,omitempty"`
	// Detected start or end of a text segment.
	DetectedBreak        *TextAnnotation_DetectedBreak `protobuf:"bytes,2,opt,name=detected_break,json=detectedBreak,proto3" json:"detected_break,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *TextAnnotation_TextProperty) Reset()         { *m = TextAnnotation_TextProperty{} }
func (m *TextAnnotation_TextProperty) String() string { return proto.CompactTextString(m) }
func (*TextAnnotation_TextProperty) ProtoMessage()    {}
func (*TextAnnotation_TextProperty) Descriptor() ([]byte, []int) {
	return fileDescriptor_text_annotation_5ca1a16a28ef616a, []int{0, 2}
}
func (m *TextAnnotation_TextProperty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextAnnotation_TextProperty.Unmarshal(m, b)
}
func (m *TextAnnotation_TextProperty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextAnnotation_TextProperty.Marshal(b, m, deterministic)
}
func (dst *TextAnnotation_TextProperty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextAnnotation_TextProperty.Merge(dst, src)
}
func (m *TextAnnotation_TextProperty) XXX_Size() int {
	return xxx_messageInfo_TextAnnotation_TextProperty.Size(m)
}
func (m *TextAnnotation_TextProperty) XXX_DiscardUnknown() {
	xxx_messageInfo_TextAnnotation_TextProperty.DiscardUnknown(m)
}

var xxx_messageInfo_TextAnnotation_TextProperty proto.InternalMessageInfo

func (m *TextAnnotation_TextProperty) GetDetectedLanguages() []*TextAnnotation_DetectedLanguage {
	if m != nil {
		return m.DetectedLanguages
	}
	return nil
}

func (m *TextAnnotation_TextProperty) GetDetectedBreak() *TextAnnotation_DetectedBreak {
	if m != nil {
		return m.DetectedBreak
	}
	return nil
}

// Detected page from OCR.
type Page struct {
	// Additional information detected on the page.
	Property *TextAnnotation_TextProperty `protobuf:"bytes,1,opt,name=property,proto3" json:"property,omitempty"`
	// Page width. For PDFs the unit is points. For images (including
	// TIFFs) the unit is pixels.
	Width int32 `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	// Page height. For PDFs the unit is points. For images (including
	// TIFFs) the unit is pixels.
	Height int32 `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	// List of blocks of text, images etc on this page.
	Blocks []*Block `protobuf:"bytes,4,rep,name=blocks,proto3" json:"blocks,omitempty"`
	// Confidence of the OCR results on the page. Range [0, 1].
	Confidence           float32  `protobuf:"fixed32,5,opt,name=confidence,proto3" json:"confidence,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Page) Reset()         { *m = Page{} }
func (m *Page) String() string { return proto.CompactTextString(m) }
func (*Page) ProtoMessage()    {}
func (*Page) Descriptor() ([]byte, []int) {
	return fileDescriptor_text_annotation_5ca1a16a28ef616a, []int{1}
}
func (m *Page) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Page.Unmarshal(m, b)
}
func (m *Page) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Page.Marshal(b, m, deterministic)
}
func (dst *Page) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Page.Merge(dst, src)
}
func (m *Page) XXX_Size() int {
	return xxx_messageInfo_Page.Size(m)
}
func (m *Page) XXX_DiscardUnknown() {
	xxx_messageInfo_Page.DiscardUnknown(m)
}

var xxx_messageInfo_Page proto.InternalMessageInfo

func (m *Page) GetProperty() *TextAnnotation_TextProperty {
	if m != nil {
		return m.Property
	}
	return nil
}

func (m *Page) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Page) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Page) GetBlocks() []*Block {
	if m != nil {
		return m.Blocks
	}
	return nil
}

func (m *Page) GetConfidence() float32 {
	if m != nil {
		return m.Confidence
	}
	return 0
}

// Logical element on the page.
type Block struct {
	// Additional information detected for the block.
	Property *TextAnnotation_TextProperty `protobuf:"bytes,1,opt,name=property,proto3" json:"property,omitempty"`
	// The bounding box for the block.
	// The vertices are in the order of top-left, top-right, bottom-right,
	// bottom-left. When a rotation of the bounding box is detected the rotation
	// is represented as around the top-left corner as defined when the text is
	// read in the 'natural' orientation.
	// For example:
	//
	// * when the text is horizontal it might look like:
	//
	//         0----1
	//         |    |
	//         3----2
	//
	// * when it's rotated 180 degrees around the top-left corner it becomes:
	//
	//         2----3
	//         |    |
	//         1----0
	//
	//   and the vertice order will still be (0, 1, 2, 3).
	BoundingBox *BoundingPoly `protobuf:"bytes,2,opt,name=bounding_box,json=boundingBox,proto3" json:"bounding_box,omitempty"`
	// List of paragraphs in this block (if this blocks is of type text).
	Paragraphs []*Paragraph `protobuf:"bytes,3,rep,name=paragraphs,proto3" json:"paragraphs,omitempty"`
	// Detected block type (text, image etc) for this block.
	BlockType Block_BlockType `protobuf:"varint,4,opt,name=block_type,json=blockType,proto3,enum=google.cloud.vision.v1.Block_BlockType" json:"block_type,omitempty"`
	// Confidence of the OCR results on the block. Range [0, 1].
	Confidence           float32  `protobuf:"fixed32,5,opt,name=confidence,proto3" json:"confidence,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_text_annotation_5ca1a16a28ef616a, []int{2}
}
func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (dst *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(dst, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetProperty() *TextAnnotation_TextProperty {
	if m != nil {
		return m.Property
	}
	return nil
}

func (m *Block) GetBoundingBox() *BoundingPoly {
	if m != nil {
		return m.BoundingBox
	}
	return nil
}

func (m *Block) GetParagraphs() []*Paragraph {
	if m != nil {
		return m.Paragraphs
	}
	return nil
}

func (m *Block) GetBlockType() Block_BlockType {
	if m != nil {
		return m.BlockType
	}
	return Block_UNKNOWN
}

func (m *Block) GetConfidence() float32 {
	if m != nil {
		return m.Confidence
	}
	return 0
}

// Structural unit of text representing a number of words in certain order.
type Paragraph struct {
	// Additional information detected for the paragraph.
	Property *TextAnnotation_TextProperty `protobuf:"bytes,1,opt,name=property,proto3" json:"property,omitempty"`
	// The bounding box for the paragraph.
	// The vertices are in the order of top-left, top-right, bottom-right,
	// bottom-left. When a rotation of the bounding box is detected the rotation
	// is represented as around the top-left corner as defined when the text is
	// read in the 'natural' orientation.
	// For example:
	//   * when the text is horizontal it might look like:
	//      0----1
	//      |    |
	//      3----2
	//   * when it's rotated 180 degrees around the top-left corner it becomes:
	//      2----3
	//      |    |
	//      1----0
	//   and the vertice order will still be (0, 1, 2, 3).
	BoundingBox *BoundingPoly `protobuf:"bytes,2,opt,name=bounding_box,json=boundingBox,proto3" json:"bounding_box,omitempty"`
	// List of words in this paragraph.
	Words []*Word `protobuf:"bytes,3,rep,name=words,proto3" json:"words,omitempty"`
	// Confidence of the OCR results for the paragraph. Range [0, 1].
	Confidence           float32  `protobuf:"fixed32,4,opt,name=confidence,proto3" json:"confidence,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Paragraph) Reset()         { *m = Paragraph{} }
func (m *Paragraph) String() string { return proto.CompactTextString(m) }
func (*Paragraph) ProtoMessage()    {}
func (*Paragraph) Descriptor() ([]byte, []int) {
	return fileDescriptor_text_annotation_5ca1a16a28ef616a, []int{3}
}
func (m *Paragraph) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Paragraph.Unmarshal(m, b)
}
func (m *Paragraph) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Paragraph.Marshal(b, m, deterministic)
}
func (dst *Paragraph) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Paragraph.Merge(dst, src)
}
func (m *Paragraph) XXX_Size() int {
	return xxx_messageInfo_Paragraph.Size(m)
}
func (m *Paragraph) XXX_DiscardUnknown() {
	xxx_messageInfo_Paragraph.DiscardUnknown(m)
}

var xxx_messageInfo_Paragraph proto.InternalMessageInfo

func (m *Paragraph) GetProperty() *TextAnnotation_TextProperty {
	if m != nil {
		return m.Property
	}
	return nil
}

func (m *Paragraph) GetBoundingBox() *BoundingPoly {
	if m != nil {
		return m.BoundingBox
	}
	return nil
}

func (m *Paragraph) GetWords() []*Word {
	if m != nil {
		return m.Words
	}
	return nil
}

func (m *Paragraph) GetConfidence() float32 {
	if m != nil {
		return m.Confidence
	}
	return 0
}

// A word representation.
type Word struct {
	// Additional information detected for the word.
	Property *TextAnnotation_TextProperty `protobuf:"bytes,1,opt,name=property,proto3" json:"property,omitempty"`
	// The bounding box for the word.
	// The vertices are in the order of top-left, top-right, bottom-right,
	// bottom-left. When a rotation of the bounding box is detected the rotation
	// is represented as around the top-left corner as defined when the text is
	// read in the 'natural' orientation.
	// For example:
	//   * when the text is horizontal it might look like:
	//      0----1
	//      |    |
	//      3----2
	//   * when it's rotated 180 degrees around the top-left corner it becomes:
	//      2----3
	//      |    |
	//      1----0
	//   and the vertice order will still be (0, 1, 2, 3).
	BoundingBox *BoundingPoly `protobuf:"bytes,2,opt,name=bounding_box,json=boundingBox,proto3" json:"bounding_box,omitempty"`
	// List of symbols in the word.
	// The order of the symbols follows the natural reading order.
	Symbols []*Symbol `protobuf:"bytes,3,rep,name=symbols,proto3" json:"symbols,omitempty"`
	// Confidence of the OCR results for the word. Range [0, 1].
	Confidence           float32  `protobuf:"fixed32,4,opt,name=confidence,proto3" json:"confidence,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Word) Reset()         { *m = Word{} }
func (m *Word) String() string { return proto.CompactTextString(m) }
func (*Word) ProtoMessage()    {}
func (*Word) Descriptor() ([]byte, []int) {
	return fileDescriptor_text_annotation_5ca1a16a28ef616a, []int{4}
}
func (m *Word) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Word.Unmarshal(m, b)
}
func (m *Word) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Word.Marshal(b, m, deterministic)
}
func (dst *Word) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Word.Merge(dst, src)
}
func (m *Word) XXX_Size() int {
	return xxx_messageInfo_Word.Size(m)
}
func (m *Word) XXX_DiscardUnknown() {
	xxx_messageInfo_Word.DiscardUnknown(m)
}

var xxx_messageInfo_Word proto.InternalMessageInfo

func (m *Word) GetProperty() *TextAnnotation_TextProperty {
	if m != nil {
		return m.Property
	}
	return nil
}

func (m *Word) GetBoundingBox() *BoundingPoly {
	if m != nil {
		return m.BoundingBox
	}
	return nil
}

func (m *Word) GetSymbols() []*Symbol {
	if m != nil {
		return m.Symbols
	}
	return nil
}

func (m *Word) GetConfidence() float32 {
	if m != nil {
		return m.Confidence
	}
	return 0
}

// A single symbol representation.
type Symbol struct {
	// Additional information detected for the symbol.
	Property *TextAnnotation_TextProperty `protobuf:"bytes,1,opt,name=property,proto3" json:"property,omitempty"`
	// The bounding box for the symbol.
	// The vertices are in the order of top-left, top-right, bottom-right,
	// bottom-left. When a rotation of the bounding box is detected the rotation
	// is represented as around the top-left corner as defined when the text is
	// read in the 'natural' orientation.
	// For example:
	//   * when the text is horizontal it might look like:
	//      0----1
	//      |    |
	//      3----2
	//   * when it's rotated 180 degrees around the top-left corner it becomes:
	//      2----3
	//      |    |
	//      1----0
	//   and the vertice order will still be (0, 1, 2, 3).
	BoundingBox *BoundingPoly `protobuf:"bytes,2,opt,name=bounding_box,json=boundingBox,proto3" json:"bounding_box,omitempty"`
	// The actual UTF-8 representation of the symbol.
	Text string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	// Confidence of the OCR results for the symbol. Range [0, 1].
	Confidence           float32  `protobuf:"fixed32,4,opt,name=confidence,proto3" json:"confidence,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Symbol) Reset()         { *m = Symbol{} }
func (m *Symbol) String() string { return proto.CompactTextString(m) }
func (*Symbol) ProtoMessage()    {}
func (*Symbol) Descriptor() ([]byte, []int) {
	return fileDescriptor_text_annotation_5ca1a16a28ef616a, []int{5}
}
func (m *Symbol) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Symbol.Unmarshal(m, b)
}
func (m *Symbol) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Symbol.Marshal(b, m, deterministic)
}
func (dst *Symbol) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Symbol.Merge(dst, src)
}
func (m *Symbol) XXX_Size() int {
	return xxx_messageInfo_Symbol.Size(m)
}
func (m *Symbol) XXX_DiscardUnknown() {
	xxx_messageInfo_Symbol.DiscardUnknown(m)
}

var xxx_messageInfo_Symbol proto.InternalMessageInfo

func (m *Symbol) GetProperty() *TextAnnotation_TextProperty {
	if m != nil {
		return m.Property
	}
	return nil
}

func (m *Symbol) GetBoundingBox() *BoundingPoly {
	if m != nil {
		return m.BoundingBox
	}
	return nil
}

func (m *Symbol) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Symbol) GetConfidence() float32 {
	if m != nil {
		return m.Confidence
	}
	return 0
}

func init() {
	proto.RegisterType((*TextAnnotation)(nil), "google.cloud.vision.v1.TextAnnotation")
	proto.RegisterType((*TextAnnotation_DetectedLanguage)(nil), "google.cloud.vision.v1.TextAnnotation.DetectedLanguage")
	proto.RegisterType((*TextAnnotation_DetectedBreak)(nil), "google.cloud.vision.v1.TextAnnotation.DetectedBreak")
	proto.RegisterType((*TextAnnotation_TextProperty)(nil), "google.cloud.vision.v1.TextAnnotation.TextProperty")
	proto.RegisterType((*Page)(nil), "google.cloud.vision.v1.Page")
	proto.RegisterType((*Block)(nil), "google.cloud.vision.v1.Block")
	proto.RegisterType((*Paragraph)(nil), "google.cloud.vision.v1.Paragraph")
	proto.RegisterType((*Word)(nil), "google.cloud.vision.v1.Word")
	proto.RegisterType((*Symbol)(nil), "google.cloud.vision.v1.Symbol")
	proto.RegisterEnum("google.cloud.vision.v1.TextAnnotation_DetectedBreak_BreakType", TextAnnotation_DetectedBreak_BreakType_name, TextAnnotation_DetectedBreak_BreakType_value)
	proto.RegisterEnum("google.cloud.vision.v1.Block_BlockType", Block_BlockType_name, Block_BlockType_value)
}

func init() {
	proto.RegisterFile("google/cloud/vision/v1/text_annotation.proto", fileDescriptor_text_annotation_5ca1a16a28ef616a)
}

var fileDescriptor_text_annotation_5ca1a16a28ef616a = []byte{
	// 774 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0xcf, 0x6f, 0xd3, 0x58,
	0x10, 0x5e, 0x27, 0x76, 0x1a, 0x4f, 0xda, 0xc8, 0xfb, 0x76, 0x55, 0x45, 0xd9, 0x6e, 0xd5, 0xcd,
	0x82, 0xe8, 0x01, 0x39, 0x6a, 0x0a, 0x02, 0x09, 0x84, 0x14, 0xa7, 0xa6, 0xad, 0x1a, 0x25, 0xd6,
	0x6b, 0x42, 0xf9, 0x71, 0xb0, 0x1c, 0xfb, 0xd5, 0xb1, 0x9a, 0xfa, 0x59, 0xb6, 0xdb, 0x26, 0xe2,
	0x9f, 0x41, 0xfc, 0x4b, 0x48, 0x9c, 0xb8, 0x72, 0xe6, 0x0a, 0x9c, 0x90, 0x9f, 0xed, 0x34, 0x09,
	0x98, 0x02, 0xe2, 0xd0, 0x8b, 0xf5, 0x66, 0xf2, 0xcd, 0x37, 0xf3, 0xcd, 0x78, 0xf2, 0x0c, 0xb7,
	0x6d, 0x4a, 0xed, 0x11, 0xa9, 0x9b, 0x23, 0x7a, 0x66, 0xd5, 0xcf, 0x9d, 0xc0, 0xa1, 0x6e, 0xfd,
	0x7c, 0xab, 0x1e, 0x92, 0x71, 0xa8, 0x1b, 0xae, 0x4b, 0x43, 0x23, 0x74, 0xa8, 0x2b, 0x7b, 0x3e,
	0x0d, 0x29, 0x5a, 0x8d, 0xd1, 0x32, 0x43, 0xcb, 0x31, 0x5a, 0x3e, 0xdf, 0xaa, 0xae, 0x25, 0x2c,
	0x86, 0xe7, 0xd4, 0x2f, 0x83, 0x82, 0x38, 0xaa, 0x7a, 0x33, 0x23, 0x87, 0x4d, 0xe8, 0x29, 0x09,
	0xfd, 0x49, 0x0c, 0xab, 0x7d, 0xe0, 0xa1, 0xdc, 0x23, 0xe3, 0xb0, 0x39, 0x25, 0x40, 0x0d, 0x10,
	0x3c, 0xc3, 0x26, 0x41, 0x85, 0xdb, 0xc8, 0x6f, 0x96, 0x1a, 0x6b, 0xf2, 0xb7, 0xf3, 0xcb, 0x9a,
	0x61, 0x13, 0x1c, 0x43, 0x11, 0x02, 0x3e, 0x2a, 0xbe, 0x92, 0xdb, 0xe0, 0x36, 0x45, 0xcc, 0xce,
	0xd5, 0x23, 0x90, 0x76, 0x48, 0x48, 0xcc, 0x90, 0x58, 0x6d, 0xc3, 0xb5, 0xcf, 0x0c, 0x9b, 0xa0,
	0xff, 0x61, 0x65, 0x94, 0x9c, 0x75, 0x93, 0x5a, 0xa4, 0xc2, 0xb1, 0x80, 0xe5, 0xd4, 0xd9, 0xa2,
	0x16, 0x41, 0xeb, 0x00, 0x26, 0x75, 0x8f, 0x1d, 0x8b, 0xb8, 0x26, 0x61, 0x94, 0x39, 0x3c, 0xe3,
	0xa9, 0xbe, 0xe7, 0x60, 0x25, 0x65, 0x56, 0x7c, 0x62, 0x9c, 0x20, 0x0c, 0x7c, 0x38, 0xf1, 0x62,
	0xb6, 0x72, 0xe3, 0x51, 0x56, 0xc5, 0xf3, 0x42, 0xe5, 0x39, 0x0e, 0x99, 0x3d, 0x7b, 0x13, 0x8f,
	0x60, 0xc6, 0x85, 0xfe, 0x01, 0xd1, 0x09, 0x74, 0xcf, 0x27, 0xc7, 0xce, 0x98, 0x15, 0x51, 0xc4,
	0x45, 0x27, 0xd0, 0x98, 0x5d, 0x33, 0x41, 0x9c, 0xe2, 0x51, 0x09, 0x96, 0xfa, 0x9d, 0x83, 0x4e,
	0xf7, 0xa8, 0x23, 0xfd, 0x81, 0x44, 0x10, 0x0e, 0xb5, 0x66, 0x4b, 0x95, 0x38, 0x54, 0x06, 0x38,
	0xec, 0x63, 0x55, 0x8f, 0xed, 0x1c, 0x42, 0x50, 0x56, 0xbb, 0x6d, 0x7d, 0xc6, 0x97, 0x47, 0x00,
	0x85, 0xbd, 0x67, 0xda, 0x9e, 0xda, 0x91, 0xf8, 0x08, 0xdf, 0xde, 0xef, 0xa8, 0xba, 0x82, 0xd5,
	0xe6, 0x81, 0x24, 0x54, 0xdf, 0x70, 0xb0, 0x1c, 0x95, 0xac, 0xf9, 0xd4, 0x23, 0x7e, 0x38, 0x41,
	0xc7, 0x80, 0xac, 0xa4, 0x66, 0x3d, 0xed, 0x58, 0x3a, 0xa6, 0x7b, 0x3f, 0x29, 0x3a, 0x1d, 0x09,
	0xfe, 0xd3, 0x5a, 0xf0, 0x04, 0xe8, 0x05, 0x94, 0xa7, 0x79, 0x06, 0x91, 0x4c, 0xa6, 0xbf, 0xd4,
	0xb8, 0xf3, 0x2b, 0x8d, 0xc5, 0x2b, 0xd6, 0xac, 0x59, 0x7b, 0xc7, 0x01, 0x1f, 0xbd, 0x3a, 0xa8,
	0x0b, 0x45, 0x2f, 0x51, 0xc6, 0x06, 0x57, 0x6a, 0x6c, 0xff, 0x20, 0xff, 0x6c, 0x53, 0xf0, 0x94,
	0x04, 0xfd, 0x0d, 0xc2, 0x85, 0x63, 0x85, 0x43, 0x56, 0xad, 0x80, 0x63, 0x03, 0xad, 0x42, 0x61,
	0x48, 0x1c, 0x7b, 0x18, 0x56, 0xf2, 0xcc, 0x9d, 0x58, 0xe8, 0x2e, 0x14, 0x06, 0x23, 0x6a, 0x9e,
	0x04, 0x15, 0x9e, 0x35, 0xf0, 0xdf, 0xac, 0xe4, 0x4a, 0x84, 0xc2, 0x09, 0x78, 0xe1, 0xe5, 0x14,
	0x16, 0x5f, 0xce, 0xda, 0xab, 0x3c, 0x08, 0x2c, 0xe2, 0xf7, 0xeb, 0xdb, 0x85, 0xe5, 0x01, 0x3d,
	0x73, 0x2d, 0xc7, 0xb5, 0xf5, 0x01, 0x1d, 0x27, 0x43, 0xb9, 0x91, 0x59, 0x77, 0x82, 0xd5, 0xe8,
	0x68, 0x82, 0x4b, 0x69, 0xa4, 0x42, 0xc7, 0xa8, 0x09, 0xe0, 0x19, 0xbe, 0x61, 0xfb, 0x86, 0x37,
	0x0c, 0x2a, 0x79, 0x26, 0xff, 0xbf, 0xec, 0x35, 0x4f, 0x90, 0x78, 0x26, 0x08, 0x3d, 0x06, 0x60,
	0x0d, 0xd1, 0xd9, 0xde, 0xf1, 0x6c, 0xef, 0x6e, 0x7d, 0xb7, 0x83, 0xf1, 0x93, 0x2d, 0x98, 0x38,
	0x48, 0x8f, 0x57, 0xb6, 0x13, 0x83, 0x38, 0x8d, 0x9b, 0x5f, 0xb4, 0x22, 0xf0, 0x3d, 0xf5, 0x69,
	0x4f, 0xe2, 0xa2, 0x95, 0xeb, 0x35, 0x95, 0x76, 0xb4, 0x62, 0x25, 0x58, 0xd2, 0xf6, 0x5b, 0xbd,
	0x3e, 0x8e, 0x76, 0x4b, 0x04, 0x01, 0xf7, 0xdb, 0x2a, 0x96, 0xf8, 0xc8, 0xaf, 0x34, 0x71, 0xab,
	0xbb, 0xa3, 0x4a, 0x42, 0xed, 0x13, 0x07, 0xe2, 0x54, 0xd5, 0x35, 0x1e, 0x53, 0x03, 0x84, 0x0b,
	0xea, 0x5b, 0xe9, 0x84, 0x32, 0xff, 0x88, 0x8f, 0xa8, 0x6f, 0xe1, 0x18, 0xba, 0xd0, 0x4f, 0xfe,
	0xab, 0x7e, 0x7e, 0xe6, 0x80, 0x8f, 0xf0, 0xd7, 0x58, 0xf6, 0x7d, 0x58, 0x0a, 0x26, 0xa7, 0x03,
	0x3a, 0x4a, 0x85, 0xaf, 0x67, 0x71, 0x1c, 0x32, 0x18, 0x4e, 0xe1, 0x57, 0x8a, 0x7f, 0xcb, 0x41,
	0x21, 0x8e, 0xb9, 0xc6, 0xf2, 0xd3, 0xab, 0x34, 0x7f, 0x79, 0x95, 0x5e, 0x25, 0x4c, 0x79, 0x09,
	0x55, 0x93, 0x9e, 0x66, 0xe4, 0x52, 0xfe, 0x9a, 0x57, 0xa0, 0x45, 0x17, 0xbf, 0xc6, 0x3d, 0x7f,
	0x98, 0xc0, 0x6d, 0x1a, 0xdd, 0x25, 0x32, 0xf5, 0xed, 0xba, 0x4d, 0x5c, 0xf6, 0x59, 0x50, 0x8f,
	0x7f, 0x32, 0x3c, 0x27, 0x58, 0xfc, 0x80, 0x78, 0x10, 0x9f, 0x3e, 0x72, 0xdc, 0xeb, 0x1c, 0xbf,
	0xdb, 0x7a, 0xd2, 0x19, 0x14, 0x58, 0xc8, 0xf6, 0x97, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x88,
	0x57, 0x38, 0xd6, 0x08, 0x00, 0x00,
}
