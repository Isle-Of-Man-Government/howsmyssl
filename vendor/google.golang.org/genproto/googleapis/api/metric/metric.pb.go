// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/api/metric/metric.proto
// DO NOT EDIT!

/*
Package google_api is a generated protocol buffer package.

It is generated from these files:
	google.golang.org/genproto/googleapis/api/metric/metric.proto

It has these top-level messages:
	MetricDescriptor
	Metric
*/
package google_api // import "google.golang.org/genproto/googleapis/api/metric"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_api1 "google.golang.org/genproto/googleapis/api/label"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The kind of measurement. It describes how the data is reported.
type MetricDescriptor_MetricKind int32

const (
	// Do not use this default value.
	MetricDescriptor_METRIC_KIND_UNSPECIFIED MetricDescriptor_MetricKind = 0
	// Instantaneous measurements of a varying quantity.
	MetricDescriptor_GAUGE MetricDescriptor_MetricKind = 1
	// Changes over non-overlapping time intervals.
	MetricDescriptor_DELTA MetricDescriptor_MetricKind = 2
	// Cumulative value over time intervals that can overlap.
	// The overlapping intervals must have the same start time.
	MetricDescriptor_CUMULATIVE MetricDescriptor_MetricKind = 3
)

var MetricDescriptor_MetricKind_name = map[int32]string{
	0: "METRIC_KIND_UNSPECIFIED",
	1: "GAUGE",
	2: "DELTA",
	3: "CUMULATIVE",
}
var MetricDescriptor_MetricKind_value = map[string]int32{
	"METRIC_KIND_UNSPECIFIED": 0,
	"GAUGE":                   1,
	"DELTA":                   2,
	"CUMULATIVE":              3,
}

func (x MetricDescriptor_MetricKind) String() string {
	return proto.EnumName(MetricDescriptor_MetricKind_name, int32(x))
}
func (MetricDescriptor_MetricKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

// The value type of a metric.
type MetricDescriptor_ValueType int32

const (
	// Do not use this default value.
	MetricDescriptor_VALUE_TYPE_UNSPECIFIED MetricDescriptor_ValueType = 0
	// The value is a boolean.
	// This value type can be used only if the metric kind is `GAUGE`.
	MetricDescriptor_BOOL MetricDescriptor_ValueType = 1
	// The value is a signed 64-bit integer.
	MetricDescriptor_INT64 MetricDescriptor_ValueType = 2
	// The value is a double precision floating point number.
	MetricDescriptor_DOUBLE MetricDescriptor_ValueType = 3
	// The value is a text string.
	// This value type can be used only if the metric kind is `GAUGE`.
	MetricDescriptor_STRING MetricDescriptor_ValueType = 4
	// The value is a [`Distribution`][google.api.Distribution].
	MetricDescriptor_DISTRIBUTION MetricDescriptor_ValueType = 5
	// The value is money.
	MetricDescriptor_MONEY MetricDescriptor_ValueType = 6
)

var MetricDescriptor_ValueType_name = map[int32]string{
	0: "VALUE_TYPE_UNSPECIFIED",
	1: "BOOL",
	2: "INT64",
	3: "DOUBLE",
	4: "STRING",
	5: "DISTRIBUTION",
	6: "MONEY",
}
var MetricDescriptor_ValueType_value = map[string]int32{
	"VALUE_TYPE_UNSPECIFIED": 0,
	"BOOL":         1,
	"INT64":        2,
	"DOUBLE":       3,
	"STRING":       4,
	"DISTRIBUTION": 5,
	"MONEY":        6,
}

func (x MetricDescriptor_ValueType) String() string {
	return proto.EnumName(MetricDescriptor_ValueType_name, int32(x))
}
func (MetricDescriptor_ValueType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 1}
}

// Defines a metric type and its schema.
type MetricDescriptor struct {
	// Resource name. The format of the name may vary between different
	// implementations. For examples:
	//
	//     projects/{project_id}/metricDescriptors/{type=**}
	//     metricDescriptors/{type=**}
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The metric type including a DNS name prefix, for example
	// `"compute.googleapis.com/instance/cpu/utilization"`. Metric types
	// should use a natural hierarchical grouping such as the following:
	//
	//     compute.googleapis.com/instance/cpu/utilization
	//     compute.googleapis.com/instance/disk/read_ops_count
	//     compute.googleapis.com/instance/network/received_bytes_count
	//
	// Note that if the metric type changes, the monitoring data will be
	// discontinued, and anything depends on it will break, such as monitoring
	// dashboards, alerting rules and quota limits. Therefore, once a metric has
	// been published, its type should be immutable.
	Type string `protobuf:"bytes,8,opt,name=type" json:"type,omitempty"`
	// The set of labels that can be used to describe a specific instance of this
	// metric type. For example, the
	// `compute.googleapis.com/instance/network/received_bytes_count` metric type
	// has a label, `loadbalanced`, that specifies whether the traffic was
	// received through a load balanced IP address.
	Labels []*google_api1.LabelDescriptor `protobuf:"bytes,2,rep,name=labels" json:"labels,omitempty"`
	// Whether the metric records instantaneous values, changes to a value, etc.
	MetricKind MetricDescriptor_MetricKind `protobuf:"varint,3,opt,name=metric_kind,json=metricKind,enum=google.api.MetricDescriptor_MetricKind" json:"metric_kind,omitempty"`
	// Whether the measurement is an integer, a floating-point number, etc.
	ValueType MetricDescriptor_ValueType `protobuf:"varint,4,opt,name=value_type,json=valueType,enum=google.api.MetricDescriptor_ValueType" json:"value_type,omitempty"`
	// The unit in which the metric value is reported. It is only applicable
	// if the `value_type` is `INT64`, `DOUBLE`, or `DISTRIBUTION`. The
	// supported units are a subset of [The Unified Code for Units of
	// Measure](http://unitsofmeasure.org/ucum.html) standard:
	//
	// **Basic units (UNIT)**
	//
	// * `bit`   bit
	// * `By`    byte
	// * `s`     second
	// * `min`   minute
	// * `h`     hour
	// * `d`     day
	//
	// **Prefixes (PREFIX)**
	//
	// * `k`     kilo    (10**3)
	// * `M`     mega    (10**6)
	// * `G`     giga    (10**9)
	// * `T`     tera    (10**12)
	// * `P`     peta    (10**15)
	// * `E`     exa     (10**18)
	// * `Z`     zetta   (10**21)
	// * `Y`     yotta   (10**24)
	// * `m`     milli   (10**-3)
	// * `u`     micro   (10**-6)
	// * `n`     nano    (10**-9)
	// * `p`     pico    (10**-12)
	// * `f`     femto   (10**-15)
	// * `a`     atto    (10**-18)
	// * `z`     zepto   (10**-21)
	// * `y`     yocto   (10**-24)
	// * `Ki`    kibi    (2**10)
	// * `Mi`    mebi    (2**20)
	// * `Gi`    gibi    (2**30)
	// * `Ti`    tebi    (2**40)
	//
	// **Grammar**
	//
	// The grammar includes the dimensionless unit `1`, such as `1/s`.
	//
	// The grammar also includes these connectors:
	//
	// * `/`    division (as an infix operator, e.g. `1/s`).
	// * `.`    multiplication (as an infix operator, e.g. `GBy.d`)
	//
	// The grammar for a unit is as follows:
	//
	//     Expression = Component { "." Component } { "/" Component } ;
	//
	//     Component = [ PREFIX ] UNIT [ Annotation ]
	//               | Annotation
	//               | "1"
	//               ;
	//
	//     Annotation = "{" NAME "}" ;
	//
	// Notes:
	//
	// * `Annotation` is just a comment if it follows a `UNIT` and is
	//    equivalent to `1` if it is used alone. For examples,
	//    `{requests}/s == 1/s`, `By{transmitted}/s == By/s`.
	// * `NAME` is a sequence of non-blank printable ASCII characters not
	//    containing '{' or '}'.
	Unit string `protobuf:"bytes,5,opt,name=unit" json:"unit,omitempty"`
	// A detailed description of the metric, which can be used in documentation.
	Description string `protobuf:"bytes,6,opt,name=description" json:"description,omitempty"`
	// A concise name for the metric, which can be displayed in user interfaces.
	// Use sentence case without an ending period, for example "Request count".
	DisplayName string `protobuf:"bytes,7,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
}

func (m *MetricDescriptor) Reset()                    { *m = MetricDescriptor{} }
func (m *MetricDescriptor) String() string            { return proto.CompactTextString(m) }
func (*MetricDescriptor) ProtoMessage()               {}
func (*MetricDescriptor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MetricDescriptor) GetLabels() []*google_api1.LabelDescriptor {
	if m != nil {
		return m.Labels
	}
	return nil
}

// A specific metric identified by specifying values for all of the
// labels of a [`MetricDescriptor`][google.api.MetricDescriptor].
type Metric struct {
	// An existing metric type, see [google.api.MetricDescriptor][google.api.MetricDescriptor].
	// For example, `compute.googleapis.com/instance/cpu/usage_time`.
	Type string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	// The set of labels that uniquely identify a metric. To specify a
	// metric, all labels enumerated in the `MetricDescriptor` must be
	// assigned values.
	Labels map[string]string `protobuf:"bytes,2,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Metric) Reset()                    { *m = Metric{} }
func (m *Metric) String() string            { return proto.CompactTextString(m) }
func (*Metric) ProtoMessage()               {}
func (*Metric) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Metric) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterType((*MetricDescriptor)(nil), "google.api.MetricDescriptor")
	proto.RegisterType((*Metric)(nil), "google.api.Metric")
	proto.RegisterEnum("google.api.MetricDescriptor_MetricKind", MetricDescriptor_MetricKind_name, MetricDescriptor_MetricKind_value)
	proto.RegisterEnum("google.api.MetricDescriptor_ValueType", MetricDescriptor_ValueType_name, MetricDescriptor_ValueType_value)
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/api/metric/metric.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 491 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x52, 0xdf, 0x6b, 0x9b, 0x50,
	0x14, 0x9e, 0x89, 0x71, 0xcd, 0xb1, 0x04, 0xb9, 0x8c, 0x4d, 0x52, 0xd8, 0xb2, 0x3c, 0x6c, 0x7d,
	0x32, 0xd0, 0x8e, 0xb2, 0x1f, 0xec, 0x21, 0xc6, 0xbb, 0x4c, 0x6a, 0x34, 0x58, 0x0d, 0xf4, 0x49,
	0x6c, 0x22, 0x22, 0x35, 0xea, 0xd4, 0x16, 0xf2, 0x57, 0xec, 0x2f, 0xd8, 0xff, 0xba, 0xfb, 0xc3,
	0x26, 0x92, 0xc1, 0xe8, 0x8b, 0xf9, 0xee, 0x77, 0xce, 0xf9, 0xee, 0x77, 0x6e, 0x3e, 0xf8, 0x1e,
	0xe7, 0x79, 0x9c, 0x46, 0x5a, 0x9c, 0xa7, 0x61, 0x16, 0x6b, 0x79, 0x19, 0x4f, 0xe2, 0x28, 0x2b,
	0xca, 0xbc, 0xce, 0x27, 0xbc, 0x14, 0x16, 0x49, 0x35, 0x21, 0x9f, 0xc9, 0x36, 0xaa, 0xcb, 0x64,
	0xdd, 0xfc, 0x68, 0xac, 0x05, 0x41, 0x33, 0x4e, 0xea, 0xc3, 0x6f, 0xcf, 0x97, 0x4a, 0xc3, 0xbb,
	0x28, 0xe5, 0x5f, 0x2e, 0x34, 0xfe, 0x23, 0x82, 0xb2, 0x60, 0xca, 0x46, 0x54, 0xad, 0xcb, 0xa4,
	0xa8, 0xf3, 0x12, 0x21, 0x10, 0xb3, 0x70, 0x1b, 0xa9, 0xc2, 0x48, 0x38, 0xef, 0xbb, 0x0c, 0x53,
	0xae, 0xde, 0x15, 0x91, 0x7a, 0xc2, 0x39, 0x8a, 0xd1, 0x25, 0x48, 0x4c, 0xab, 0x52, 0x3b, 0xa3,
	0xee, 0xb9, 0x7c, 0x71, 0xa6, 0x1d, 0x6c, 0x69, 0x16, 0xad, 0x1c, 0x44, 0xdd, 0xa6, 0x15, 0xfd,
	0x04, 0x99, 0xaf, 0x12, 0xdc, 0x27, 0xd9, 0x46, 0xed, 0x12, 0xbd, 0xc1, 0xc5, 0xc7, 0xf6, 0xe4,
	0xb1, 0x9f, 0x86, 0xb8, 0x26, 0xed, 0x2e, 0x6c, 0xf7, 0x18, 0x61, 0x80, 0xc7, 0x30, 0x7d, 0x88,
	0x02, 0x66, 0x4c, 0x64, 0x42, 0x1f, 0xfe, 0x2b, 0xb4, 0xa2, 0xed, 0x1e, 0xe9, 0x76, 0xfb, 0x8f,
	0x4f, 0x90, 0x6e, 0xf6, 0x90, 0x25, 0xb5, 0xda, 0xe3, 0x9b, 0x51, 0x8c, 0x46, 0x20, 0x6f, 0x9a,
	0xb1, 0x24, 0xcf, 0x54, 0x89, 0x95, 0xda, 0x14, 0x7a, 0x0f, 0xa7, 0x9b, 0xa4, 0x2a, 0xd2, 0x70,
	0x17, 0xb0, 0xb7, 0x7a, 0xd9, 0xb4, 0x70, 0xce, 0x26, 0xd4, 0xd8, 0x01, 0x38, 0x38, 0x47, 0x67,
	0xf0, 0x66, 0x81, 0x3d, 0xd7, 0x9c, 0x05, 0xd7, 0xa6, 0x6d, 0x04, 0xbe, 0x7d, 0xb3, 0xc4, 0x33,
	0xf3, 0x87, 0x89, 0x0d, 0xe5, 0x05, 0xea, 0x43, 0x6f, 0x3e, 0xf5, 0xe7, 0x58, 0x11, 0x28, 0x34,
	0xb0, 0xe5, 0x4d, 0x95, 0x0e, 0x1a, 0x00, 0xcc, 0xfc, 0x85, 0x6f, 0x4d, 0x3d, 0x73, 0x85, 0x95,
	0xee, 0xf8, 0x17, 0xf4, 0xf7, 0x1b, 0xa0, 0x21, 0xbc, 0x5e, 0x4d, 0x2d, 0x1f, 0x07, 0xde, 0xed,
	0x12, 0x1f, 0xc9, 0x9d, 0x80, 0xa8, 0x3b, 0x8e, 0xc5, 0xd5, 0x4c, 0xdb, 0xbb, 0xfa, 0x44, 0xd4,
	0x00, 0x24, 0xc3, 0xf1, 0x75, 0x8b, 0x28, 0x51, 0x7c, 0x43, 0xbc, 0xd8, 0x73, 0x45, 0x44, 0x0a,
	0x9c, 0x1a, 0x26, 0x3d, 0xe9, 0xbe, 0x67, 0x3a, 0xb6, 0xd2, 0xa3, 0x43, 0x0b, 0xc7, 0xc6, 0xb7,
	0x8a, 0x34, 0xfe, 0x2d, 0x80, 0xc4, 0x97, 0xd8, 0x27, 0xa0, 0xdb, 0x4a, 0xc0, 0xd5, 0x51, 0x02,
	0xde, 0xfe, 0xfb, 0xfc, 0x3c, 0x08, 0x15, 0xce, 0xea, 0x72, 0xf7, 0x14, 0x82, 0xe1, 0x17, 0x90,
	0x5b, 0x34, 0xb1, 0xd0, 0xbd, 0x8f, 0x76, 0x4d, 0xde, 0x28, 0x44, 0xaf, 0xa0, 0xc7, 0xfe, 0x21,
	0xa2, 0x4b, 0x39, 0x7e, 0xf8, 0xda, 0xf9, 0x2c, 0xe8, 0xef, 0x60, 0xb0, 0xce, 0xb7, 0xad, 0x7b,
	0x74, 0x99, 0x5f, 0xb4, 0xa4, 0x81, 0x5e, 0x0a, 0x77, 0x12, 0x4b, 0xf6, 0xe5, 0xdf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x56, 0x02, 0x32, 0xa4, 0x63, 0x03, 0x00, 0x00,
}
