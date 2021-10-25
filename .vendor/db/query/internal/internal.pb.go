// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: internal/internal.proto

/*
Package query is a generated protocol buffer package.

It is generated from these files:
	internal/internal.proto

It has these top-level messages:
	Point
	Aux
	IteratorOptions
	Metrics
	Metric
	Interval
	IteratorStats
	VarRef
*/
package query

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Point struct {
	Name             *string        `protobuf:"bytes,1,req,name=Name" json:"Name,omitempty"`
	Tags             *string        `protobuf:"bytes,2,req,name=Tags" json:"Tags,omitempty"`
	Time             *int64         `protobuf:"varint,3,req,name=Time" json:"Time,omitempty"`
	Nil              *bool          `protobuf:"varint,4,req,name=Nil" json:"Nil,omitempty"`
	Aux              []*Aux         `protobuf:"bytes,5,rep,name=Aux" json:"Aux,omitempty"`
	Aggregated       *uint32        `protobuf:"varint,6,opt,name=Aggregated" json:"Aggregated,omitempty"`
	FloatValue       *float64       `protobuf:"fixed64,7,opt,name=FloatValue" json:"FloatValue,omitempty"`
	IntegerValue     *int64         `protobuf:"varint,8,opt,name=IntegerValue" json:"IntegerValue,omitempty"`
	StringValue      *string        `protobuf:"bytes,9,opt,name=StringValue" json:"StringValue,omitempty"`
	BooleanValue     *bool          `protobuf:"varint,10,opt,name=BooleanValue" json:"BooleanValue,omitempty"`
	UnsignedValue    *uint64        `protobuf:"varint,12,opt,name=UnsignedValue" json:"UnsignedValue,omitempty"`
	Stats            *IteratorStats `protobuf:"bytes,11,opt,name=Stats" json:"Stats,omitempty"`
	Trace            []byte         `protobuf:"bytes,13,opt,name=Trace" json:"Trace,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (m *Point) String() string            { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{0} }

func (m *Point) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Point) GetTags() string {
	if m != nil && m.Tags != nil {
		return *m.Tags
	}
	return ""
}

func (m *Point) GetTime() int64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *Point) GetNil() bool {
	if m != nil && m.Nil != nil {
		return *m.Nil
	}
	return false
}

func (m *Point) GetAux() []*Aux {
	if m != nil {
		return m.Aux
	}
	return nil
}

func (m *Point) GetAggregated() uint32 {
	if m != nil && m.Aggregated != nil {
		return *m.Aggregated
	}
	return 0
}

func (m *Point) GetFloatValue() float64 {
	if m != nil && m.FloatValue != nil {
		return *m.FloatValue
	}
	return 0
}

func (m *Point) GetIntegerValue() int64 {
	if m != nil && m.IntegerValue != nil {
		return *m.IntegerValue
	}
	return 0
}

func (m *Point) GetStringValue() string {
	if m != nil && m.StringValue != nil {
		return *m.StringValue
	}
	return ""
}

func (m *Point) GetBooleanValue() bool {
	if m != nil && m.BooleanValue != nil {
		return *m.BooleanValue
	}
	return false
}

func (m *Point) GetUnsignedValue() uint64 {
	if m != nil && m.UnsignedValue != nil {
		return *m.UnsignedValue
	}
	return 0
}

func (m *Point) GetStats() *IteratorStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *Point) GetTrace() []byte {
	if m != nil {
		return m.Trace
	}
	return nil
}

type Aux struct {
	DataType         *int32   `protobuf:"varint,1,req,name=DataType" json:"DataType,omitempty"`
	FloatValue       *float64 `protobuf:"fixed64,2,opt,name=FloatValue" json:"FloatValue,omitempty"`
	IntegerValue     *int64   `protobuf:"varint,3,opt,name=IntegerValue" json:"IntegerValue,omitempty"`
	StringValue      *string  `protobuf:"bytes,4,opt,name=StringValue" json:"StringValue,omitempty"`
	BooleanValue     *bool    `protobuf:"varint,5,opt,name=BooleanValue" json:"BooleanValue,omitempty"`
	UnsignedValue    *uint64  `protobuf:"varint,6,opt,name=UnsignedValue" json:"UnsignedValue,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Aux) Reset()                    { *m = Aux{} }
func (m *Aux) String() string            { return proto.CompactTextString(m) }
func (*Aux) ProtoMessage()               {}
func (*Aux) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{1} }

func (m *Aux) GetDataType() int32 {
	if m != nil && m.DataType != nil {
		return *m.DataType
	}
	return 0
}

func (m *Aux) GetFloatValue() float64 {
	if m != nil && m.FloatValue != nil {
		return *m.FloatValue
	}
	return 0
}

func (m *Aux) GetIntegerValue() int64 {
	if m != nil && m.IntegerValue != nil {
		return *m.IntegerValue
	}
	return 0
}

func (m *Aux) GetStringValue() string {
	if m != nil && m.StringValue != nil {
		return *m.StringValue
	}
	return ""
}

func (m *Aux) GetBooleanValue() bool {
	if m != nil && m.BooleanValue != nil {
		return *m.BooleanValue
	}
	return false
}

func (m *Aux) GetUnsignedValue() uint64 {
	if m != nil && m.UnsignedValue != nil {
		return *m.UnsignedValue
	}
	return 0
}

type IteratorOptions struct {
	Expr             *string        `protobuf:"bytes,1,opt,name=Expr" json:"Expr,omitempty"`
	Aux              []string       `protobuf:"bytes,2,rep,name=Aux" json:"Aux,omitempty"`
	Fields           []*VarRef      `protobuf:"bytes,17,rep,name=Fields" json:"Fields,omitempty"`
	Sources          []*Metric `protobuf:"bytes,3,rep,name=Sources" json:"Sources,omitempty"`
	Interval         *Interval      `protobuf:"bytes,4,opt,name=Interval" json:"Interval,omitempty"`
	Dimensions       []string       `protobuf:"bytes,5,rep,name=Dimensions" json:"Dimensions,omitempty"`
	GroupBy          []string       `protobuf:"bytes,19,rep,name=GroupBy" json:"GroupBy,omitempty"`
	Fill             *int32         `protobuf:"varint,6,opt,name=Fill" json:"Fill,omitempty"`
	FillValue        *float64       `protobuf:"fixed64,7,opt,name=FillValue" json:"FillValue,omitempty"`
	Condition        *string        `protobuf:"bytes,8,opt,name=Condition" json:"Condition,omitempty"`
	StartTime        *int64         `protobuf:"varint,9,opt,name=StartTime" json:"StartTime,omitempty"`
	EndTime          *int64         `protobuf:"varint,10,opt,name=EndTime" json:"EndTime,omitempty"`
	Location         *string        `protobuf:"bytes,21,opt,name=Location" json:"Location,omitempty"`
	Ascending        *bool          `protobuf:"varint,11,opt,name=Ascending" json:"Ascending,omitempty"`
	Limit            *int64         `protobuf:"varint,12,opt,name=Limit" json:"Limit,omitempty"`
	Offset           *int64         `protobuf:"varint,13,opt,name=Offset" json:"Offset,omitempty"`
	SLimit           *int64         `protobuf:"varint,14,opt,name=SLimit" json:"SLimit,omitempty"`
	SOffset          *int64         `protobuf:"varint,15,opt,name=SOffset" json:"SOffset,omitempty"`
	StripName        *bool          `protobuf:"varint,22,opt,name=StripName" json:"StripName,omitempty"`
	Dedupe           *bool          `protobuf:"varint,16,opt,name=Dedupe" json:"Dedupe,omitempty"`
	MaxSeriesN       *int64         `protobuf:"varint,18,opt,name=MaxSeriesN" json:"MaxSeriesN,omitempty"`
	Ordered          *bool          `protobuf:"varint,20,opt,name=Ordered" json:"Ordered,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *IteratorOptions) Reset()                    { *m = IteratorOptions{} }
func (m *IteratorOptions) String() string            { return proto.CompactTextString(m) }
func (*IteratorOptions) ProtoMessage()               {}
func (*IteratorOptions) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{2} }

func (m *IteratorOptions) GetExpr() string {
	if m != nil && m.Expr != nil {
		return *m.Expr
	}
	return ""
}

func (m *IteratorOptions) GetAux() []string {
	if m != nil {
		return m.Aux
	}
	return nil
}

func (m *IteratorOptions) GetFields() []*VarRef {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *IteratorOptions) GetSources() []*Metric {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *IteratorOptions) GetInterval() *Interval {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *IteratorOptions) GetDimensions() []string {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *IteratorOptions) GetGroupBy() []string {
	if m != nil {
		return m.GroupBy
	}
	return nil
}

func (m *IteratorOptions) GetFill() int32 {
	if m != nil && m.Fill != nil {
		return *m.Fill
	}
	return 0
}

func (m *IteratorOptions) GetFillValue() float64 {
	if m != nil && m.FillValue != nil {
		return *m.FillValue
	}
	return 0
}

func (m *IteratorOptions) GetCondition() string {
	if m != nil && m.Condition != nil {
		return *m.Condition
	}
	return ""
}

func (m *IteratorOptions) GetStartTime() int64 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return 0
}

func (m *IteratorOptions) GetEndTime() int64 {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return 0
}

func (m *IteratorOptions) GetLocation() string {
	if m != nil && m.Location != nil {
		return *m.Location
	}
	return ""
}

func (m *IteratorOptions) GetAscending() bool {
	if m != nil && m.Ascending != nil {
		return *m.Ascending
	}
	return false
}

func (m *IteratorOptions) GetLimit() int64 {
	if m != nil && m.Limit != nil {
		return *m.Limit
	}
	return 0
}

func (m *IteratorOptions) GetOffset() int64 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

func (m *IteratorOptions) GetSLimit() int64 {
	if m != nil && m.SLimit != nil {
		return *m.SLimit
	}
	return 0
}

func (m *IteratorOptions) GetSOffset() int64 {
	if m != nil && m.SOffset != nil {
		return *m.SOffset
	}
	return 0
}

func (m *IteratorOptions) GetStripName() bool {
	if m != nil && m.StripName != nil {
		return *m.StripName
	}
	return false
}

func (m *IteratorOptions) GetDedupe() bool {
	if m != nil && m.Dedupe != nil {
		return *m.Dedupe
	}
	return false
}

func (m *IteratorOptions) GetMaxSeriesN() int64 {
	if m != nil && m.MaxSeriesN != nil {
		return *m.MaxSeriesN
	}
	return 0
}

func (m *IteratorOptions) GetOrdered() bool {
	if m != nil && m.Ordered != nil {
		return *m.Ordered
	}
	return false
}

type Metrics struct {
	Items            []*Metric `protobuf:"bytes,1,rep,name=Items" json:"Items,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Metrics) Reset()                    { *m = Metrics{} }
func (m *Metrics) String() string            { return proto.CompactTextString(m) }
func (*Metrics) ProtoMessage()               {}
func (*Metrics) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{3} }

func (m *Metrics) GetItems() []*Metric {
	if m != nil {
		return m.Items
	}
	return nil
}

type Metric struct {
	Database         *string `protobuf:"bytes,1,opt,name=Database" json:"Database,omitempty"`
	TimeToLive  *string `protobuf:"bytes,2,opt,name=TimeToLive" json:"TimeToLive,omitempty"`
	Name             *string `protobuf:"bytes,3,opt,name=Name" json:"Name,omitempty"`
	Regex            *string `protobuf:"bytes,4,opt,name=Regex" json:"Regex,omitempty"`
	IsTarget         *bool   `protobuf:"varint,5,opt,name=IsTarget" json:"IsTarget,omitempty"`
	SystemIterator   *string `protobuf:"bytes,6,opt,name=SystemIterator" json:"SystemIterator,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Metric) Reset()                    { *m = Metric{} }
func (m *Metric) String() string            { return proto.CompactTextString(m) }
func (*Metric) ProtoMessage()               {}
func (*Metric) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{4} }

func (m *Metric) GetDatabase() string {
	if m != nil && m.Database != nil {
		return *m.Database
	}
	return ""
}

func (m *Metric) GetTimeToLive() string {
	if m != nil && m.TimeToLive != nil {
		return *m.TimeToLive
	}
	return ""
}

func (m *Metric) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Metric) GetRegex() string {
	if m != nil && m.Regex != nil {
		return *m.Regex
	}
	return ""
}

func (m *Metric) GetIsTarget() bool {
	if m != nil && m.IsTarget != nil {
		return *m.IsTarget
	}
	return false
}

func (m *Metric) GetSystemIterator() string {
	if m != nil && m.SystemIterator != nil {
		return *m.SystemIterator
	}
	return ""
}

type Interval struct {
	Duration         *int64 `protobuf:"varint,1,opt,name=Duration" json:"Duration,omitempty"`
	Offset           *int64 `protobuf:"varint,2,opt,name=Offset" json:"Offset,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Interval) Reset()                    { *m = Interval{} }
func (m *Interval) String() string            { return proto.CompactTextString(m) }
func (*Interval) ProtoMessage()               {}
func (*Interval) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{5} }

func (m *Interval) GetDuration() int64 {
	if m != nil && m.Duration != nil {
		return *m.Duration
	}
	return 0
}

func (m *Interval) GetOffset() int64 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

type IteratorStats struct {
	SeriesN          *int64 `protobuf:"varint,1,opt,name=SeriesN" json:"SeriesN,omitempty"`
	PointN           *int64 `protobuf:"varint,2,opt,name=PointN" json:"PointN,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *IteratorStats) Reset()                    { *m = IteratorStats{} }
func (m *IteratorStats) String() string            { return proto.CompactTextString(m) }
func (*IteratorStats) ProtoMessage()               {}
func (*IteratorStats) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{6} }

func (m *IteratorStats) GetSeriesN() int64 {
	if m != nil && m.SeriesN != nil {
		return *m.SeriesN
	}
	return 0
}

func (m *IteratorStats) GetPointN() int64 {
	if m != nil && m.PointN != nil {
		return *m.PointN
	}
	return 0
}

type VarRef struct {
	Val              *string `protobuf:"bytes,1,req,name=Val" json:"Val,omitempty"`
	Type             *int32  `protobuf:"varint,2,opt,name=Type" json:"Type,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *VarRef) Reset()                    { *m = VarRef{} }
func (m *VarRef) String() string            { return proto.CompactTextString(m) }
func (*VarRef) ProtoMessage()               {}
func (*VarRef) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{7} }

func (m *VarRef) GetVal() string {
	if m != nil && m.Val != nil {
		return *m.Val
	}
	return ""
}

func (m *VarRef) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func init() {
	proto.RegisterType((*Point)(nil), "query.Point")
	proto.RegisterType((*Aux)(nil), "query.Aux")
	proto.RegisterType((*IteratorOptions)(nil), "query.IteratorOptions")
	proto.RegisterType((*Metrics)(nil), "query.Metrics")
	proto.RegisterType((*Metric)(nil), "query.Metric")
	proto.RegisterType((*Interval)(nil), "query.Interval")
	proto.RegisterType((*IteratorStats)(nil), "query.IteratorStats")
	proto.RegisterType((*VarRef)(nil), "query.VarRef")
}

func init() { proto.RegisterFile("internal/internal.proto", fileDescriptorInternal) }

var fileDescriptorInternal = []byte{
	// 796 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x6d, 0x6f, 0xe3, 0x44,
	0x10, 0x96, 0xe3, 0x3a, 0x8d, 0x27, 0xcd, 0xf5, 0x58, 0x4a, 0x59, 0xa1, 0x13, 0xb2, 0x2c, 0x40,
	0x16, 0xa0, 0x22, 0xf5, 0x13, 0x9f, 0x90, 0x72, 0xf4, 0x8a, 0x2a, 0xdd, 0xb5, 0xa7, 0x4d, 0xe9,
	0xf7, 0x25, 0x9e, 0x5a, 0x2b, 0x39, 0xeb, 0xb0, 0x5e, 0xa3, 0xe4, 0x07, 0xf4, 0x87, 0xf1, 0x13,
	0xf8, 0x47, 0x68, 0x67, 0xd7, 0x89, 0x53, 0x81, 0x7a, 0x9f, 0x32, 0xcf, 0x33, 0x93, 0x7d, 0x79,
	0xe6, 0x99, 0x35, 0x7c, 0xa9, 0xb4, 0x45, 0xa3, 0x65, 0xfd, 0x53, 0x1f, 0x5c, 0xac, 0x4d, 0x63,
	0x1b, 0x96, 0xfc, 0xd9, 0xa1, 0xd9, 0xe6, 0x4f, 0x31, 0x24, 0x1f, 0x1b, 0xa5, 0x2d, 0x63, 0x70,
	0x74, 0x2b, 0x57, 0xc8, 0xa3, 0x6c, 0x54, 0xa4, 0x82, 0x62, 0xc7, 0xdd, 0xcb, 0xaa, 0xe5, 0x23,
	0xcf, 0xb9, 0x98, 0x38, 0xb5, 0x42, 0x1e, 0x67, 0xa3, 0x22, 0x16, 0x14, 0xb3, 0xd7, 0x10, 0xdf,
	0xaa, 0x9a, 0x1f, 0x65, 0xa3, 0x62, 0x22, 0x5c, 0xc8, 0xde, 0x40, 0x3c, 0xef, 0x36, 0x3c, 0xc9,
	0xe2, 0x62, 0x7a, 0x09, 0x17, 0xb4, 0xd9, 0xc5, 0xbc, 0xdb, 0x08, 0x47, 0xb3, 0xaf, 0x01, 0xe6,
	0x55, 0x65, 0xb0, 0x92, 0x16, 0x4b, 0x3e, 0xce, 0xa2, 0x62, 0x26, 0x06, 0x8c, 0xcb, 0x5f, 0xd7,
	0x8d, 0xb4, 0x0f, 0xb2, 0xee, 0x90, 0x1f, 0x67, 0x51, 0x11, 0x89, 0x01, 0xc3, 0x72, 0x38, 0xb9,
	0xd1, 0x16, 0x2b, 0x34, 0xbe, 0x62, 0x92, 0x45, 0x45, 0x2c, 0x0e, 0x38, 0x96, 0xc1, 0x74, 0x61,
	0x8d, 0xd2, 0x95, 0x2f, 0x49, 0xb3, 0xa8, 0x48, 0xc5, 0x90, 0x72, 0xab, 0xbc, 0x6d, 0x9a, 0x1a,
	0xa5, 0xf6, 0x25, 0x90, 0x45, 0xc5, 0x44, 0x1c, 0x70, 0xec, 0x1b, 0x98, 0xfd, 0xae, 0x5b, 0x55,
	0x69, 0x2c, 0x7d, 0xd1, 0x49, 0x16, 0x15, 0x47, 0xe2, 0x90, 0x64, 0xdf, 0x43, 0xb2, 0xb0, 0xd2,
	0xb6, 0x7c, 0x9a, 0x45, 0xc5, 0xf4, 0xf2, 0x2c, 0xdc, 0xf7, 0xc6, 0xa2, 0x91, 0xb6, 0x31, 0x94,
	0x13, 0xbe, 0x84, 0x9d, 0x41, 0x72, 0x6f, 0xe4, 0x12, 0xf9, 0x2c, 0x8b, 0x8a, 0x13, 0xe1, 0x41,
	0xfe, 0x4f, 0x44, 0x82, 0xb1, 0xaf, 0x60, 0x72, 0x25, 0xad, 0xbc, 0xdf, 0xae, 0x7d, 0x27, 0x12,
	0xb1, 0xc3, 0xcf, 0x54, 0x19, 0xbd, 0xa8, 0x4a, 0xfc, 0xb2, 0x2a, 0x47, 0x2f, 0xab, 0x92, 0x7c,
	0x8a, 0x2a, 0xe3, 0xff, 0x50, 0x25, 0x7f, 0x4a, 0xe0, 0xb4, 0x97, 0xe0, 0x6e, 0x6d, 0x55, 0xa3,
	0xc9, 0x3d, 0xef, 0x36, 0x6b, 0xc3, 0x23, 0xda, 0x98, 0x62, 0xe7, 0x1e, 0xe7, 0x95, 0x51, 0x16,
	0x17, 0xa9, 0xf7, 0xc7, 0xb7, 0x30, 0xbe, 0x56, 0x58, 0x97, 0x2d, 0xff, 0x8c, 0x0c, 0x34, 0x0b,
	0x82, 0x3e, 0x48, 0x23, 0xf0, 0x51, 0x84, 0x24, 0xfb, 0x11, 0x8e, 0x17, 0x4d, 0x67, 0x96, 0xd8,
	0xf2, 0x98, 0xea, 0x58, 0xa8, 0xfb, 0x80, 0xb2, 0xed, 0x0c, 0xae, 0x50, 0x5b, 0xd1, 0x97, 0xb0,
	0x1f, 0x60, 0xe2, 0xa4, 0x30, 0x7f, 0xc9, 0x9a, 0xee, 0x3d, 0xbd, 0x3c, 0xed, 0xfb, 0x14, 0x68,
	0xb1, 0x2b, 0x70, 0x5a, 0x5f, 0xa9, 0x15, 0xea, 0xd6, 0x9d, 0x9a, 0x6c, 0x9c, 0x8a, 0x01, 0xc3,
	0x38, 0x1c, 0xff, 0x66, 0x9a, 0x6e, 0xfd, 0x76, 0xcb, 0x3f, 0xa7, 0x64, 0x0f, 0xdd, 0x0d, 0xaf,
	0x55, 0x5d, 0x93, 0x24, 0x89, 0xa0, 0x98, 0xbd, 0x81, 0xd4, 0xfd, 0x0e, 0xed, 0xbc, 0x27, 0x5c,
	0xf6, 0xd7, 0x46, 0x97, 0xca, 0x29, 0x44, 0x56, 0x4e, 0xc5, 0x9e, 0x70, 0xd9, 0x85, 0x95, 0xc6,
	0xd2, 0xd0, 0xa5, 0xd4, 0xd2, 0x3d, 0xe1, 0xce, 0xf1, 0x4e, 0x97, 0x94, 0x03, 0xca, 0xf5, 0xd0,
	0x39, 0xe9, 0x7d, 0xb3, 0x94, 0xb4, 0xe8, 0x17, 0xb4, 0xe8, 0x0e, 0xbb, 0x35, 0xe7, 0xed, 0x12,
	0x75, 0xa9, 0x74, 0x45, 0x9e, 0x9d, 0x88, 0x3d, 0xe1, 0x1c, 0xfa, 0x5e, 0xad, 0x94, 0x25, 0xaf,
	0xc7, 0xc2, 0x03, 0x76, 0x0e, 0xe3, 0xbb, 0xc7, 0xc7, 0x16, 0x2d, 0x19, 0x37, 0x16, 0x01, 0x39,
	0x7e, 0xe1, 0xcb, 0x5f, 0x79, 0xde, 0x23, 0x77, 0xb2, 0x45, 0xf8, 0xc3, 0xa9, 0x3f, 0x59, 0x80,
	0xfe, 0x46, 0x46, 0xad, 0xe9, 0xb9, 0x39, 0xf7, 0xbb, 0xef, 0x08, 0xb7, 0xde, 0x15, 0x96, 0xdd,
	0x1a, 0xf9, 0x6b, 0x4a, 0x05, 0xe4, 0x3a, 0xf2, 0x41, 0x6e, 0x16, 0x68, 0x14, 0xb6, 0xb7, 0x9c,
	0xd1, 0x92, 0x03, 0xc6, 0xed, 0x77, 0x67, 0x4a, 0x34, 0x58, 0xf2, 0x33, 0xfa, 0x63, 0x0f, 0xf3,
	0x9f, 0xe1, 0x64, 0x60, 0x88, 0x96, 0x15, 0x90, 0xdc, 0x58, 0x5c, 0xb5, 0x3c, 0xfa, 0x5f, 0xd3,
	0xf8, 0x82, 0xfc, 0xef, 0x08, 0xa6, 0x03, 0xba, 0x9f, 0xce, 0x3f, 0x64, 0x8b, 0xc1, 0xc1, 0x3b,
	0xcc, 0x0a, 0x38, 0x15, 0x68, 0x51, 0x3b, 0x81, 0x3f, 0x36, 0xb5, 0x5a, 0x6e, 0x69, 0x44, 0x53,
	0xf1, 0x9c, 0xde, 0xbd, 0xb4, 0xb1, 0x9f, 0x01, 0xba, 0xf5, 0x19, 0x24, 0x02, 0x2b, 0xdc, 0x84,
	0x89, 0xf4, 0xc0, 0xed, 0x77, 0xd3, 0xde, 0x4b, 0x53, 0xa1, 0x0d, 0x73, 0xb8, 0xc3, 0xec, 0x3b,
	0x78, 0xb5, 0xd8, 0xb6, 0x16, 0x57, 0xfd, 0x88, 0x91, 0xe3, 0x52, 0xf1, 0x8c, 0xcd, 0x7f, 0xd9,
	0xdb, 0x9e, 0xce, 0xdf, 0x19, 0xef, 0x89, 0x88, 0x14, 0xdc, 0xe1, 0x41, 0x7f, 0x47, 0xc3, 0xfe,
	0xe6, 0x73, 0x98, 0x1d, 0xbc, 0x63, 0xd4, 0xd8, 0xd0, 0x85, 0x28, 0x34, 0x36, 0xb4, 0xe0, 0x1c,
	0xc6, 0xf4, 0x2d, 0xb9, 0xed, 0x97, 0xf0, 0x28, 0xbf, 0x80, 0xb1, 0x9f, 0x5c, 0x37, 0xea, 0x0f,
	0xb2, 0x0e, 0xdf, 0x18, 0x17, 0xd2, 0xe7, 0xc4, 0x3d, 0x76, 0x23, 0x3f, 0x2e, 0x2e, 0xfe, 0x37,
	0x00, 0x00, 0xff, 0xff, 0x07, 0x98, 0x54, 0xa1, 0xb5, 0x06, 0x00, 0x00,
}
