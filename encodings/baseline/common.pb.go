// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package baseline

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// AttributeValueType is the enumeration of possible types that value can have.
type AttributeValueType int32

const (
	AttributeValueType_STRING AttributeValueType = 0
	AttributeValueType_INT    AttributeValueType = 1
	AttributeValueType_DOUBLE AttributeValueType = 2
	AttributeValueType_BOOL   AttributeValueType = 3
	AttributeValueType_LIST   AttributeValueType = 4
	AttributeValueType_MAP    AttributeValueType = 5
)

var AttributeValueType_name = map[int32]string{
	0: "STRING",
	1: "INT",
	2: "DOUBLE",
	3: "BOOL",
	4: "LIST",
	5: "MAP",
}

var AttributeValueType_value = map[string]int32{
	"STRING": 0,
	"INT":    1,
	"DOUBLE": 2,
	"BOOL":   3,
	"LIST":   4,
	"MAP":    5,
}

func (x AttributeValueType) String() string {
	return proto.EnumName(AttributeValueType_name, int32(x))
}

func (AttributeValueType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

// ValueType is the enumeration of possible types that value can have.
type AttributeKeyValue_ValueType int32

const (
	AttributeKeyValue_STRING AttributeKeyValue_ValueType = 0
	AttributeKeyValue_INT    AttributeKeyValue_ValueType = 1
	AttributeKeyValue_DOUBLE AttributeKeyValue_ValueType = 2
	AttributeKeyValue_BOOL   AttributeKeyValue_ValueType = 3
)

var AttributeKeyValue_ValueType_name = map[int32]string{
	0: "STRING",
	1: "INT",
	2: "DOUBLE",
	3: "BOOL",
}

var AttributeKeyValue_ValueType_value = map[string]int32{
	"STRING": 0,
	"INT":    1,
	"DOUBLE": 2,
	"BOOL":   3,
}

func (x AttributeKeyValue_ValueType) String() string {
	return proto.EnumName(AttributeKeyValue_ValueType_name, int32(x))
}

func (AttributeKeyValue_ValueType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{4, 0}
}

// InstrumentationLibrary is a message representing the instrumentation library information
// such as the fully qualified name and version.
type InstrumentationLibrary struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstrumentationLibrary) Reset()         { *m = InstrumentationLibrary{} }
func (m *InstrumentationLibrary) String() string { return proto.CompactTextString(m) }
func (*InstrumentationLibrary) ProtoMessage()    {}
func (*InstrumentationLibrary) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *InstrumentationLibrary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstrumentationLibrary.Unmarshal(m, b)
}
func (m *InstrumentationLibrary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstrumentationLibrary.Marshal(b, m, deterministic)
}
func (m *InstrumentationLibrary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstrumentationLibrary.Merge(m, src)
}
func (m *InstrumentationLibrary) XXX_Size() int {
	return xxx_messageInfo_InstrumentationLibrary.Size(m)
}
func (m *InstrumentationLibrary) XXX_DiscardUnknown() {
	xxx_messageInfo_InstrumentationLibrary.DiscardUnknown(m)
}

var xxx_messageInfo_InstrumentationLibrary proto.InternalMessageInfo

func (m *InstrumentationLibrary) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InstrumentationLibrary) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// Resource information. This describes the source of telemetry data.
type Resource struct {
	// labels is a collection of attributes that describe the resource. See OpenTelemetry
	// specification semantic conventions for standardized label names:
	// https://github.com/open-telemetry/opentelemetry-specification/blob/master/specification/data-semantic-conventions.md
	Labels []*AttributeKeyValue `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty"`
	// dropped_labels_count is the number of dropped labels. If the value is 0, then
	// no labels were dropped.
	DroppedLabelsCount   int32    `protobuf:"varint,2,opt,name=dropped_labels_count,json=droppedLabelsCount,proto3" json:"dropped_labels_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetLabels() []*AttributeKeyValue {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Resource) GetDroppedLabelsCount() int32 {
	if m != nil {
		return m.DroppedLabelsCount
	}
	return 0
}

// AttributeValue is a value that is used to store Log body.
type AttributeValue struct {
	// type of the value.
	Type                 AttributeValueType       `protobuf:"varint,1,opt,name=type,proto3,enum=baseline.AttributeValueType" json:"type,omitempty"`
	StringValue          string                   `protobuf:"bytes,2,opt,name=string_value,json=stringValue,proto3" json:"string_value,omitempty"`
	IntValue             int64                    `protobuf:"varint,3,opt,name=int_value,json=intValue,proto3" json:"int_value,omitempty"`
	DoubleValue          float64                  `protobuf:"fixed64,4,opt,name=double_value,json=doubleValue,proto3" json:"double_value,omitempty"`
	BoolValue            bool                     `protobuf:"varint,5,opt,name=bool_value,json=boolValue,proto3" json:"bool_value,omitempty"`
	ListOrMap            *AttributeValueListOrMap `protobuf:"bytes,6,opt,name=listOrMap,proto3" json:"listOrMap,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *AttributeValue) Reset()         { *m = AttributeValue{} }
func (m *AttributeValue) String() string { return proto.CompactTextString(m) }
func (*AttributeValue) ProtoMessage()    {}
func (*AttributeValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

func (m *AttributeValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttributeValue.Unmarshal(m, b)
}
func (m *AttributeValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttributeValue.Marshal(b, m, deterministic)
}
func (m *AttributeValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttributeValue.Merge(m, src)
}
func (m *AttributeValue) XXX_Size() int {
	return xxx_messageInfo_AttributeValue.Size(m)
}
func (m *AttributeValue) XXX_DiscardUnknown() {
	xxx_messageInfo_AttributeValue.DiscardUnknown(m)
}

var xxx_messageInfo_AttributeValue proto.InternalMessageInfo

func (m *AttributeValue) GetType() AttributeValueType {
	if m != nil {
		return m.Type
	}
	return AttributeValueType_STRING
}

func (m *AttributeValue) GetStringValue() string {
	if m != nil {
		return m.StringValue
	}
	return ""
}

func (m *AttributeValue) GetIntValue() int64 {
	if m != nil {
		return m.IntValue
	}
	return 0
}

func (m *AttributeValue) GetDoubleValue() float64 {
	if m != nil {
		return m.DoubleValue
	}
	return 0
}

func (m *AttributeValue) GetBoolValue() bool {
	if m != nil {
		return m.BoolValue
	}
	return false
}

func (m *AttributeValue) GetListOrMap() *AttributeValueListOrMap {
	if m != nil {
		return m.ListOrMap
	}
	return nil
}

type AttributeValueListOrMap struct {
	List                 []*AttributeValue    `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Map                  []*AttributeKeyValue `protobuf:"bytes,2,rep,name=map,proto3" json:"map,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AttributeValueListOrMap) Reset()         { *m = AttributeValueListOrMap{} }
func (m *AttributeValueListOrMap) String() string { return proto.CompactTextString(m) }
func (*AttributeValueListOrMap) ProtoMessage()    {}
func (*AttributeValueListOrMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{3}
}

func (m *AttributeValueListOrMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttributeValueListOrMap.Unmarshal(m, b)
}
func (m *AttributeValueListOrMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttributeValueListOrMap.Marshal(b, m, deterministic)
}
func (m *AttributeValueListOrMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttributeValueListOrMap.Merge(m, src)
}
func (m *AttributeValueListOrMap) XXX_Size() int {
	return xxx_messageInfo_AttributeValueListOrMap.Size(m)
}
func (m *AttributeValueListOrMap) XXX_DiscardUnknown() {
	xxx_messageInfo_AttributeValueListOrMap.DiscardUnknown(m)
}

var xxx_messageInfo_AttributeValueListOrMap proto.InternalMessageInfo

func (m *AttributeValueListOrMap) GetList() []*AttributeValue {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *AttributeValueListOrMap) GetMap() []*AttributeKeyValue {
	if m != nil {
		return m.Map
	}
	return nil
}

// AttributeKeyValue is a key-value pair that is used to store Span attributes, Resource
// labels, etc.
type AttributeKeyValue struct {
	// key part of the key-value pair.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// type of the value.
	Type                 AttributeKeyValue_ValueType `protobuf:"varint,2,opt,name=type,proto3,enum=baseline.AttributeKeyValue_ValueType" json:"type,omitempty"`
	StringValue          string                      `protobuf:"bytes,3,opt,name=string_value,json=stringValue,proto3" json:"string_value,omitempty"`
	IntValue             int64                       `protobuf:"varint,4,opt,name=int_value,json=intValue,proto3" json:"int_value,omitempty"`
	DoubleValue          float64                     `protobuf:"fixed64,5,opt,name=double_value,json=doubleValue,proto3" json:"double_value,omitempty"`
	BoolValue            bool                        `protobuf:"varint,6,opt,name=bool_value,json=boolValue,proto3" json:"bool_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *AttributeKeyValue) Reset()         { *m = AttributeKeyValue{} }
func (m *AttributeKeyValue) String() string { return proto.CompactTextString(m) }
func (*AttributeKeyValue) ProtoMessage()    {}
func (*AttributeKeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{4}
}

func (m *AttributeKeyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttributeKeyValue.Unmarshal(m, b)
}
func (m *AttributeKeyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttributeKeyValue.Marshal(b, m, deterministic)
}
func (m *AttributeKeyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttributeKeyValue.Merge(m, src)
}
func (m *AttributeKeyValue) XXX_Size() int {
	return xxx_messageInfo_AttributeKeyValue.Size(m)
}
func (m *AttributeKeyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_AttributeKeyValue.DiscardUnknown(m)
}

var xxx_messageInfo_AttributeKeyValue proto.InternalMessageInfo

func (m *AttributeKeyValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *AttributeKeyValue) GetType() AttributeKeyValue_ValueType {
	if m != nil {
		return m.Type
	}
	return AttributeKeyValue_STRING
}

func (m *AttributeKeyValue) GetStringValue() string {
	if m != nil {
		return m.StringValue
	}
	return ""
}

func (m *AttributeKeyValue) GetIntValue() int64 {
	if m != nil {
		return m.IntValue
	}
	return 0
}

func (m *AttributeKeyValue) GetDoubleValue() float64 {
	if m != nil {
		return m.DoubleValue
	}
	return 0
}

func (m *AttributeKeyValue) GetBoolValue() bool {
	if m != nil {
		return m.BoolValue
	}
	return false
}

func init() {
	proto.RegisterEnum("baseline.AttributeValueType", AttributeValueType_name, AttributeValueType_value)
	proto.RegisterEnum("baseline.AttributeKeyValue_ValueType", AttributeKeyValue_ValueType_name, AttributeKeyValue_ValueType_value)
	proto.RegisterType((*InstrumentationLibrary)(nil), "baseline.InstrumentationLibrary")
	proto.RegisterType((*Resource)(nil), "baseline.Resource")
	proto.RegisterType((*AttributeValue)(nil), "baseline.AttributeValue")
	proto.RegisterType((*AttributeValueListOrMap)(nil), "baseline.AttributeValueListOrMap")
	proto.RegisterType((*AttributeKeyValue)(nil), "baseline.AttributeKeyValue")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xdf, 0x6a, 0xdb, 0x30,
	0x18, 0xc5, 0x27, 0xdb, 0x49, 0x93, 0x2f, 0xa1, 0x78, 0x62, 0x6c, 0x1e, 0xed, 0xc0, 0x31, 0x0c,
	0xcc, 0xd8, 0x4c, 0x9b, 0xc2, 0x60, 0x57, 0xa3, 0xe9, 0xfe, 0x85, 0xb9, 0x4d, 0x50, 0xb3, 0xde,
	0x06, 0x3b, 0x11, 0x43, 0xcc, 0x91, 0x3c, 0x59, 0x0e, 0xe4, 0x01, 0xf6, 0x12, 0x7b, 0xc4, 0x3d,
	0xc5, 0x90, 0xe4, 0x74, 0xb0, 0x2e, 0x0b, 0xbd, 0xfb, 0x72, 0xce, 0x4f, 0x47, 0x5f, 0x38, 0x32,
	0xf4, 0x17, 0x62, 0xb5, 0x12, 0x3c, 0x29, 0xa5, 0x50, 0x02, 0x77, 0xf2, 0xac, 0xa2, 0x05, 0xe3,
	0x34, 0xfa, 0x00, 0x8f, 0xc7, 0xbc, 0x52, 0xb2, 0x5e, 0x51, 0xae, 0x32, 0xc5, 0x04, 0x4f, 0x59,
	0x2e, 0x33, 0xb9, 0xc1, 0x18, 0x3c, 0x9e, 0xad, 0x68, 0x80, 0x42, 0x14, 0x77, 0x89, 0x99, 0x71,
	0x00, 0x07, 0x6b, 0x2a, 0x2b, 0x26, 0x78, 0xe0, 0x18, 0x79, 0xfb, 0x33, 0xfa, 0x0e, 0x1d, 0x42,
	0x2b, 0x51, 0xcb, 0x05, 0xc5, 0x67, 0xd0, 0x2e, 0xb2, 0x9c, 0x16, 0x55, 0x80, 0x42, 0x37, 0xee,
	0x0d, 0x8f, 0x92, 0xed, 0x75, 0xc9, 0xb9, 0x52, 0x92, 0xe5, 0xb5, 0xa2, 0x9f, 0xe9, 0xe6, 0x26,
	0x2b, 0x6a, 0x4a, 0x1a, 0x14, 0x9f, 0xc0, 0xa3, 0xa5, 0x14, 0x65, 0x49, 0x97, 0x73, 0xab, 0xcc,
	0x17, 0xa2, 0xe6, 0xca, 0xdc, 0xd3, 0x22, 0xb8, 0xf1, 0x52, 0x63, 0x5d, 0x68, 0x27, 0xfa, 0xe1,
	0xc0, 0xe1, 0x6d, 0x9e, 0x09, 0xc3, 0x27, 0xe0, 0xa9, 0x4d, 0x69, 0x77, 0x3e, 0x1c, 0x1e, 0xff,
	0xe3, 0x5e, 0xc3, 0xcd, 0x36, 0x25, 0x25, 0x86, 0xc4, 0x03, 0xe8, 0x57, 0x4a, 0x32, 0xfe, 0x75,
	0xbe, 0xd6, 0x4e, 0xf3, 0xb7, 0x7a, 0x56, 0xb3, 0xa1, 0x47, 0xd0, 0x65, 0x5c, 0x35, 0xbe, 0x1b,
	0xa2, 0xd8, 0x25, 0x1d, 0xc6, 0x95, 0x35, 0x07, 0xd0, 0x5f, 0x8a, 0x3a, 0x2f, 0x68, 0xe3, 0x7b,
	0x21, 0x8a, 0x11, 0xe9, 0x59, 0xcd, 0x22, 0xcf, 0x00, 0x72, 0x21, 0x8a, 0x06, 0x68, 0x85, 0x28,
	0xee, 0x90, 0xae, 0x56, 0xac, 0xfd, 0x16, 0xba, 0x05, 0xab, 0xd4, 0x44, 0x5e, 0x66, 0x65, 0xd0,
	0x0e, 0x51, 0xdc, 0x1b, 0x0e, 0x76, 0x2d, 0x9e, 0x6e, 0x41, 0xf2, 0xe7, 0x4c, 0xb4, 0x86, 0x27,
	0x3b, 0x28, 0xfc, 0x12, 0x3c, 0xcd, 0x35, 0x3d, 0x04, 0xbb, 0x62, 0x89, 0xa1, 0xf0, 0x2b, 0x70,
	0x57, 0x59, 0x19, 0x38, 0xfb, 0x4b, 0xd3, 0x5c, 0xf4, 0xd3, 0x81, 0x87, 0x77, 0x2c, 0xec, 0x83,
	0xfb, 0x8d, 0x6e, 0x9a, 0x57, 0xa3, 0x47, 0xfc, 0xa6, 0x29, 0xc5, 0x31, 0xa5, 0x3c, 0xff, 0x4f,
	0x6e, 0xb2, 0xaf, 0x1d, 0x77, 0x4f, 0x3b, 0xde, 0x9e, 0x76, 0x5a, 0xfb, 0xda, 0x69, 0xff, 0xd5,
	0x4e, 0xf4, 0x1a, 0xba, 0xb7, 0x4b, 0x61, 0x80, 0xf6, 0xf5, 0x8c, 0x8c, 0xaf, 0x3e, 0xfa, 0x0f,
	0xf0, 0x01, 0xb8, 0xe3, 0xab, 0x99, 0x8f, 0xb4, 0xf8, 0x6e, 0xf2, 0x65, 0x94, 0xbe, 0xf7, 0x1d,
	0xdc, 0x01, 0x6f, 0x34, 0x99, 0xa4, 0xbe, 0xfb, 0x82, 0x00, 0xbe, 0xfb, 0xe6, 0xee, 0x11, 0xa0,
	0xa7, 0x74, 0x7c, 0x3d, 0xf3, 0x3d, 0x0d, 0x5e, 0x9e, 0x4f, 0xfd, 0xd6, 0xe8, 0x13, 0x1c, 0x33,
	0x91, 0x88, 0x92, 0xf2, 0x05, 0xe5, 0x55, 0x5d, 0xd9, 0x8f, 0x39, 0x51, 0x32, 0x5b, 0xd0, 0x64,
	0x7d, 0x3a, 0x82, 0x99, 0x9e, 0xa6, 0x5a, 0x9c, 0xa2, 0x5f, 0xce, 0xd3, 0x49, 0x49, 0xf9, 0x85,
	0x25, 0x8d, 0x98, 0x18, 0x3f, 0xb9, 0x39, 0xcd, 0xdb, 0xe6, 0xe4, 0xd9, 0xef, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x36, 0xf2, 0x4d, 0xff, 0x16, 0x04, 0x00, 0x00,
}
