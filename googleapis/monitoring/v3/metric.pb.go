// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/monitoring/v3/metric.proto

package monitoring // import "github.com/sgtsquiggs/go-genproto/googleapis/monitoring/v3"

import proto "github.com/sgtsquiggs/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/sgtsquiggs/go-genproto/googleapis/api/annotations"
import _ "github.com/sgtsquiggs/go-genproto/googleapis/api/label"
import metric "github.com/sgtsquiggs/go-genproto/googleapis/api/metric"
import monitoredres "github.com/sgtsquiggs/go-genproto/googleapis/api/monitoredres"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A single data point in a time series.
type Point struct {
	// The time interval to which the data point applies.  For `GAUGE` metrics,
	// only the end time of the interval is used.  For `DELTA` metrics, the start
	// and end time should specify a non-zero interval, with subsequent points
	// specifying contiguous and non-overlapping intervals.  For `CUMULATIVE`
	// metrics, the start and end time should specify a non-zero interval, with
	// subsequent points specifying the same start time and increasing end times,
	// until an event resets the cumulative value to zero and sets a new start
	// time for the following points.
	Interval *TimeInterval `protobuf:"bytes,1,opt,name=interval,proto3" json:"interval,omitempty"`
	// The value of the data point.
	Value                *TypedValue `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Point) Reset()         { *m = Point{} }
func (m *Point) String() string { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()    {}
func (*Point) Descriptor() ([]byte, []int) {
	return fileDescriptor_metric_360b9781a7cf8bf5, []int{0}
}
func (m *Point) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Point.Unmarshal(m, b)
}
func (m *Point) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Point.Marshal(b, m, deterministic)
}
func (dst *Point) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Point.Merge(dst, src)
}
func (m *Point) XXX_Size() int {
	return xxx_messageInfo_Point.Size(m)
}
func (m *Point) XXX_DiscardUnknown() {
	xxx_messageInfo_Point.DiscardUnknown(m)
}

var xxx_messageInfo_Point proto.InternalMessageInfo

func (m *Point) GetInterval() *TimeInterval {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *Point) GetValue() *TypedValue {
	if m != nil {
		return m.Value
	}
	return nil
}

// A collection of data points that describes the time-varying values
// of a metric. A time series is identified by a combination of a
// fully-specified monitored resource and a fully-specified metric.
// This type is used for both listing and creating time series.
type TimeSeries struct {
	// The associated metric. A fully-specified metric used to identify the time
	// series.
	Metric *metric.Metric `protobuf:"bytes,1,opt,name=metric,proto3" json:"metric,omitempty"`
	// The associated monitored resource.  Custom metrics can use only certain
	// monitored resource types in their time series data.
	Resource *monitoredres.MonitoredResource `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	// Output only. The associated monitored resource metadata. When reading a
	// a timeseries, this field will include metadata labels that are explicitly
	// named in the reduction. When creating a timeseries, this field is ignored.
	Metadata *monitoredres.MonitoredResourceMetadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// The metric kind of the time series. When listing time series, this metric
	// kind might be different from the metric kind of the associated metric if
	// this time series is an alignment or reduction of other time series.
	//
	// When creating a time series, this field is optional. If present, it must be
	// the same as the metric kind of the associated metric. If the associated
	// metric's descriptor must be auto-created, then this field specifies the
	// metric kind of the new descriptor and must be either `GAUGE` (the default)
	// or `CUMULATIVE`.
	MetricKind metric.MetricDescriptor_MetricKind `protobuf:"varint,3,opt,name=metric_kind,json=metricKind,proto3,enum=google.api.MetricDescriptor_MetricKind" json:"metric_kind,omitempty"`
	// The value type of the time series. When listing time series, this value
	// type might be different from the value type of the associated metric if
	// this time series is an alignment or reduction of other time series.
	//
	// When creating a time series, this field is optional. If present, it must be
	// the same as the type of the data in the `points` field.
	ValueType metric.MetricDescriptor_ValueType `protobuf:"varint,4,opt,name=value_type,json=valueType,proto3,enum=google.api.MetricDescriptor_ValueType" json:"value_type,omitempty"`
	// The data points of this time series. When listing time series, points are
	// returned in reverse time order.
	//
	// When creating a time series, this field must contain exactly one point and
	// the point's type must be the same as the value type of the associated
	// metric. If the associated metric's descriptor must be auto-created, then
	// the value type of the descriptor is determined by the point's type, which
	// must be `BOOL`, `INT64`, `DOUBLE`, or `DISTRIBUTION`.
	Points               []*Point `protobuf:"bytes,5,rep,name=points,proto3" json:"points,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeSeries) Reset()         { *m = TimeSeries{} }
func (m *TimeSeries) String() string { return proto.CompactTextString(m) }
func (*TimeSeries) ProtoMessage()    {}
func (*TimeSeries) Descriptor() ([]byte, []int) {
	return fileDescriptor_metric_360b9781a7cf8bf5, []int{1}
}
func (m *TimeSeries) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeSeries.Unmarshal(m, b)
}
func (m *TimeSeries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeSeries.Marshal(b, m, deterministic)
}
func (dst *TimeSeries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeSeries.Merge(dst, src)
}
func (m *TimeSeries) XXX_Size() int {
	return xxx_messageInfo_TimeSeries.Size(m)
}
func (m *TimeSeries) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeSeries.DiscardUnknown(m)
}

var xxx_messageInfo_TimeSeries proto.InternalMessageInfo

func (m *TimeSeries) GetMetric() *metric.Metric {
	if m != nil {
		return m.Metric
	}
	return nil
}

func (m *TimeSeries) GetResource() *monitoredres.MonitoredResource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *TimeSeries) GetMetadata() *monitoredres.MonitoredResourceMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *TimeSeries) GetMetricKind() metric.MetricDescriptor_MetricKind {
	if m != nil {
		return m.MetricKind
	}
	return metric.MetricDescriptor_METRIC_KIND_UNSPECIFIED
}

func (m *TimeSeries) GetValueType() metric.MetricDescriptor_ValueType {
	if m != nil {
		return m.ValueType
	}
	return metric.MetricDescriptor_VALUE_TYPE_UNSPECIFIED
}

func (m *TimeSeries) GetPoints() []*Point {
	if m != nil {
		return m.Points
	}
	return nil
}

func init() {
	proto.RegisterType((*Point)(nil), "google.monitoring.v3.Point")
	proto.RegisterType((*TimeSeries)(nil), "google.monitoring.v3.TimeSeries")
}

func init() {
	proto.RegisterFile("google/monitoring/v3/metric.proto", fileDescriptor_metric_360b9781a7cf8bf5)
}

var fileDescriptor_metric_360b9781a7cf8bf5 = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x51, 0xab, 0xd3, 0x30,
	0x14, 0xc7, 0xe9, 0xae, 0x9b, 0x33, 0x03, 0x1f, 0x82, 0x68, 0x99, 0x0a, 0x73, 0xa2, 0x0e, 0x1f,
	0x5a, 0x58, 0x41, 0x10, 0xe1, 0x82, 0x57, 0x45, 0x45, 0x2e, 0x8c, 0x28, 0x7b, 0x90, 0xc1, 0xc8,
	0x6d, 0x0f, 0x25, 0xd8, 0xe4, 0x84, 0x34, 0x2b, 0xdc, 0x27, 0x3f, 0x8c, 0x6f, 0x7e, 0x14, 0x3f,
	0x93, 0x0f, 0xd2, 0x24, 0xdd, 0x76, 0xb1, 0xf7, 0xbe, 0xb5, 0xf9, 0xff, 0xfe, 0xe7, 0x7f, 0x72,
	0x72, 0xc8, 0x93, 0x12, 0xb1, 0xac, 0x20, 0x95, 0xa8, 0x84, 0x45, 0x23, 0x54, 0x99, 0x36, 0x59,
	0x2a, 0xc1, 0x1a, 0x91, 0x27, 0xda, 0xa0, 0x45, 0x7a, 0xcf, 0x23, 0xc9, 0x01, 0x49, 0x9a, 0x6c,
	0xfa, 0x28, 0x18, 0xb9, 0x16, 0x29, 0x57, 0x0a, 0x2d, 0xb7, 0x02, 0x55, 0xed, 0x3d, 0xd3, 0xfb,
	0x47, 0x6a, 0xc5, 0x2f, 0xa0, 0x0a, 0xe7, 0x0f, 0x8e, 0xce, 0x8f, 0x43, 0xa6, 0x4f, 0x8f, 0x05,
	0x1f, 0x04, 0xc5, 0xd6, 0x40, 0x8d, 0x3b, 0x93, 0x43, 0x80, 0xfa, 0x9b, 0xcd, 0x51, 0x4a, 0x54,
	0x1e, 0x99, 0xff, 0x24, 0xc3, 0x15, 0x0a, 0x65, 0xe9, 0x29, 0x19, 0x0b, 0x65, 0xc1, 0x34, 0xbc,
	0x8a, 0xa3, 0x59, 0xb4, 0x98, 0x2c, 0xe7, 0x49, 0xdf, 0x45, 0x92, 0x6f, 0x42, 0xc2, 0xe7, 0x40,
	0xb2, 0xbd, 0x87, 0xbe, 0x22, 0xc3, 0x86, 0x57, 0x3b, 0x88, 0x07, 0xce, 0x3c, 0xbb, 0xc6, 0x7c,
	0xa9, 0xa1, 0x58, 0xb7, 0x1c, 0xf3, 0xf8, 0xfc, 0xef, 0x80, 0x90, 0xb6, 0xe4, 0x57, 0x30, 0x02,
	0x6a, 0xfa, 0x92, 0x8c, 0xfc, 0x3d, 0x43, 0x13, 0xb4, 0xab, 0xc3, 0xb5, 0x48, 0xce, 0x9d, 0xc2,
	0x02, 0x41, 0x5f, 0x93, 0x71, 0x77, 0xe1, 0x90, 0xfa, 0xf8, 0x0a, 0xdd, 0x8d, 0x85, 0x05, 0x88,
	0xed, 0x71, 0xfa, 0x96, 0x8c, 0x25, 0x58, 0x5e, 0x70, 0xcb, 0xe3, 0xdb, 0xce, 0xfa, 0xec, 0x46,
	0xeb, 0x79, 0x80, 0xd9, 0xde, 0x46, 0x3f, 0x91, 0x89, 0xef, 0x63, 0xfb, 0x43, 0xa8, 0x22, 0x3e,
	0x99, 0x45, 0x8b, 0xbb, 0xcb, 0x17, 0xff, 0xb7, 0xfb, 0x1e, 0xea, 0xdc, 0x08, 0x6d, 0xd1, 0x84,
	0x83, 0x2f, 0x42, 0x15, 0x8c, 0xc8, 0xfd, 0x37, 0xfd, 0x40, 0x88, 0x9b, 0xc5, 0xd6, 0x5e, 0x6a,
	0x88, 0x6f, 0xb9, 0x42, 0xcf, 0x6f, 0x2c, 0xe4, 0x26, 0xd8, 0xce, 0x92, 0xdd, 0x69, 0xba, 0x4f,
	0x9a, 0x91, 0x91, 0x6e, 0x9f, 0xb2, 0x8e, 0x87, 0xb3, 0x93, 0xc5, 0x64, 0xf9, 0xb0, 0xff, 0x09,
	0xdc, 0x73, 0xb3, 0x80, 0x9e, 0xfd, 0x8a, 0x48, 0x9c, 0xa3, 0xec, 0x45, 0xcf, 0x26, 0x3e, 0x78,
	0xd5, 0x6e, 0xca, 0x2a, 0xfa, 0x7e, 0x1a, 0xa0, 0x12, 0x2b, 0xae, 0xca, 0x04, 0x4d, 0x99, 0x96,
	0xa0, 0xdc, 0x1e, 0xa5, 0x5e, 0xe2, 0x5a, 0xd4, 0x57, 0xb7, 0xed, 0xcd, 0xe1, 0xef, 0xf7, 0x60,
	0xfa, 0xd1, 0x17, 0x78, 0x57, 0xe1, 0xae, 0xe8, 0x86, 0xdc, 0x66, 0xad, 0xb3, 0x3f, 0x9d, 0xb8,
	0x71, 0xe2, 0xe6, 0x20, 0x6e, 0xd6, 0xd9, 0xc5, 0xc8, 0x85, 0x64, 0xff, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x5a, 0x88, 0xc9, 0x0b, 0x7e, 0x03, 0x00, 0x00,
}
