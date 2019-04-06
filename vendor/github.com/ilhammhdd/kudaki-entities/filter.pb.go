// Code generated by protoc-gen-go. DO NOT EDIT.
// source: filter.proto

package kudaki_entities

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

type Range struct {
	Min                  uint32   `protobuf:"varint,1,opt,name=min,proto3" json:"min,omitempty"`
	Max                  uint32   `protobuf:"varint,2,opt,name=max,proto3" json:"max,omitempty"`
	Unit                 string   `protobuf:"bytes,3,opt,name=unit,proto3" json:"unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Range) Reset()         { *m = Range{} }
func (m *Range) String() string { return proto.CompactTextString(m) }
func (*Range) ProtoMessage()    {}
func (*Range) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f5303cab7a20d6f, []int{0}
}

func (m *Range) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Range.Unmarshal(m, b)
}
func (m *Range) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Range.Marshal(b, m, deterministic)
}
func (m *Range) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Range.Merge(m, src)
}
func (m *Range) XXX_Size() int {
	return xxx_messageInfo_Range.Size(m)
}
func (m *Range) XXX_DiscardUnknown() {
	xxx_messageInfo_Range.DiscardUnknown(m)
}

var xxx_messageInfo_Range proto.InternalMessageInfo

func (m *Range) GetMin() uint32 {
	if m != nil {
		return m.Min
	}
	return 0
}

func (m *Range) GetMax() uint32 {
	if m != nil {
		return m.Max
	}
	return 0
}

func (m *Range) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

type Category struct {
	Name                 string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	SubCategories        []*Category `protobuf:"bytes,2,rep,name=sub_categories,json=subCategories,proto3" json:"sub_categories,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Category) Reset()         { *m = Category{} }
func (m *Category) String() string { return proto.CompactTextString(m) }
func (*Category) ProtoMessage()    {}
func (*Category) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f5303cab7a20d6f, []int{1}
}

func (m *Category) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Category.Unmarshal(m, b)
}
func (m *Category) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Category.Marshal(b, m, deterministic)
}
func (m *Category) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Category.Merge(m, src)
}
func (m *Category) XXX_Size() int {
	return xxx_messageInfo_Category.Size(m)
}
func (m *Category) XXX_DiscardUnknown() {
	xxx_messageInfo_Category.DiscardUnknown(m)
}

var xxx_messageInfo_Category proto.InternalMessageInfo

func (m *Category) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Category) GetSubCategories() []*Category {
	if m != nil {
		return m.SubCategories
	}
	return nil
}

type Filter struct {
	Range                *Range      `protobuf:"bytes,1,opt,name=range,proto3" json:"range,omitempty"`
	Categories           []*Category `protobuf:"bytes,2,rep,name=categories,proto3" json:"categories,omitempty"`
	Status               Status      `protobuf:"varint,3,opt,name=status,proto3,enum=entities.Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f5303cab7a20d6f, []int{2}
}

func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (m *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(m, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetRange() *Range {
	if m != nil {
		return m.Range
	}
	return nil
}

func (m *Filter) GetCategories() []*Category {
	if m != nil {
		return m.Categories
	}
	return nil
}

func (m *Filter) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_ON_PROGRESS
}

func init() {
	proto.RegisterType((*Range)(nil), "entities.Range")
	proto.RegisterType((*Category)(nil), "entities.Category")
	proto.RegisterType((*Filter)(nil), "entities.Filter")
}

func init() { proto.RegisterFile("filter.proto", fileDescriptor_1f5303cab7a20d6f) }

var fileDescriptor_1f5303cab7a20d6f = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x50, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0x26, 0xad, 0x0d, 0xed, 0xf4, 0xc7, 0xb2, 0xa7, 0xe0, 0x29, 0x04, 0x95, 0x5c, 0x4c, 0x21,
	0x9e, 0x3c, 0x09, 0x16, 0x7c, 0x80, 0xf5, 0xa4, 0x17, 0xd9, 0x34, 0x6b, 0x32, 0xb4, 0xbb, 0x91,
	0xec, 0x2c, 0xd4, 0x87, 0xf0, 0x9d, 0x25, 0x93, 0xc6, 0x7a, 0xf4, 0xf6, 0xf1, 0xfd, 0xcc, 0x7c,
	0x33, 0xb0, 0xf8, 0xc0, 0x03, 0xe9, 0x36, 0xfb, 0x6c, 0x1b, 0x6a, 0xc4, 0x54, 0x5b, 0x42, 0x42,
	0xed, 0xae, 0x16, 0x8e, 0x14, 0x79, 0xd7, 0xf3, 0xc9, 0x23, 0x4c, 0xa4, 0xb2, 0x95, 0x16, 0x6b,
	0x18, 0x1b, 0xb4, 0x51, 0x10, 0x07, 0xe9, 0x52, 0x76, 0x90, 0x19, 0x75, 0x8c, 0x46, 0x27, 0x46,
	0x1d, 0x85, 0x80, 0x0b, 0x6f, 0x91, 0xa2, 0x71, 0x1c, 0xa4, 0x33, 0xc9, 0x38, 0x79, 0x85, 0xe9,
	0x56, 0x91, 0xae, 0x9a, 0xf6, 0xab, 0xd3, 0xad, 0x32, 0x9a, 0x87, 0xcc, 0x24, 0x63, 0xf1, 0x00,
	0x2b, 0xe7, 0x8b, 0xf7, 0x5d, 0xef, 0x41, 0xed, 0xa2, 0x51, 0x3c, 0x4e, 0xe7, 0xb9, 0xc8, 0x86,
	0x46, 0xd9, 0x90, 0x97, 0x4b, 0xe7, 0x8b, 0xed, 0xaf, 0x31, 0xf9, 0x0e, 0x20, 0x7c, 0xe6, 0x23,
	0xc4, 0x0d, 0x4c, 0xda, 0xae, 0x26, 0x8f, 0x9e, 0xe7, 0x97, 0xe7, 0x30, 0xb7, 0x97, 0xbd, 0x2a,
	0x72, 0x80, 0x7f, 0x2d, 0xfa, 0xe3, 0x12, 0x29, 0x84, 0xfd, 0x47, 0xf8, 0xac, 0x55, 0xbe, 0x3e,
	0xfb, 0x5f, 0x98, 0x97, 0x27, 0xfd, 0xe9, 0xf6, 0xed, 0xba, 0x42, 0xaa, 0x7d, 0x91, 0xed, 0x1a,
	0xb3, 0xc1, 0x43, 0xad, 0x8c, 0xa9, 0xcb, 0x72, 0xb3, 0xf7, 0xa5, 0xda, 0xe3, 0xdd, 0x10, 0x2b,
	0x42, 0x7e, 0xed, 0xfd, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xac, 0xc0, 0x02, 0xcf, 0x82, 0x01,
	0x00, 0x00,
}
