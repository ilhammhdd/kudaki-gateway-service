// Code generated by protoc-gen-go. DO NOT EDIT.
// source: store/storefront.proto

package store

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	user "github.com/ilhammhdd/kudaki-entities/user"
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

type Storefront struct {
	Uuid                 string     `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	User                 *user.User `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	TotalItem            int32      `protobuf:"varint,3,opt,name=total_item,json=totalItem,proto3" json:"total_item,omitempty"`
	Rating               float32    `protobuf:"fixed32,4,opt,name=rating,proto3" json:"rating,omitempty"`
	CreatedAt            uint64     `protobuf:"varint,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Storefront) Reset()         { *m = Storefront{} }
func (m *Storefront) String() string { return proto.CompactTextString(m) }
func (*Storefront) ProtoMessage()    {}
func (*Storefront) Descriptor() ([]byte, []int) {
	return fileDescriptor_57a78b9ab3cd9ee2, []int{0}
}

func (m *Storefront) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Storefront.Unmarshal(m, b)
}
func (m *Storefront) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Storefront.Marshal(b, m, deterministic)
}
func (m *Storefront) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Storefront.Merge(m, src)
}
func (m *Storefront) XXX_Size() int {
	return xxx_messageInfo_Storefront.Size(m)
}
func (m *Storefront) XXX_DiscardUnknown() {
	xxx_messageInfo_Storefront.DiscardUnknown(m)
}

var xxx_messageInfo_Storefront proto.InternalMessageInfo

func (m *Storefront) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Storefront) GetUser() *user.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Storefront) GetTotalItem() int32 {
	if m != nil {
		return m.TotalItem
	}
	return 0
}

func (m *Storefront) GetRating() float32 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func (m *Storefront) GetCreatedAt() uint64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func init() {
	proto.RegisterType((*Storefront)(nil), "entities.store.Storefront")
}

func init() { proto.RegisterFile("store/storefront.proto", fileDescriptor_57a78b9ab3cd9ee2) }

var fileDescriptor_57a78b9ab3cd9ee2 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x4f, 0x4d, 0x4b, 0x03, 0x31,
	0x10, 0x25, 0x75, 0x5b, 0xd8, 0x11, 0x14, 0x22, 0x94, 0x20, 0x08, 0xc1, 0x8b, 0x41, 0x74, 0x17,
	0xf4, 0x17, 0xe8, 0xcd, 0x6b, 0xc4, 0x8b, 0x97, 0x92, 0x36, 0x63, 0x77, 0x68, 0xb3, 0x91, 0x64,
	0xf2, 0x7b, 0xfc, 0xab, 0xb2, 0xe9, 0xc7, 0xe5, 0x31, 0xf3, 0xe6, 0xbd, 0xc7, 0x3c, 0x58, 0x66,
	0x8e, 0x09, 0xfb, 0x8a, 0x3f, 0x29, 0x8e, 0xdc, 0xfd, 0xa6, 0xc8, 0x51, 0x5e, 0xe1, 0xc8, 0xc4,
	0x84, 0xb9, 0xab, 0xa7, 0xdb, 0xeb, 0x92, 0x31, 0xf5, 0x13, 0x1c, 0x04, 0xf7, 0x7f, 0x02, 0xe0,
	0xf3, 0xec, 0x92, 0x12, 0x9a, 0x52, 0xc8, 0x2b, 0xa1, 0x85, 0x69, 0x6d, 0x9d, 0xe5, 0x03, 0x34,
	0x93, 0x41, 0xcd, 0xb4, 0x30, 0x97, 0x2f, 0x37, 0xdd, 0x39, 0xb2, 0xc6, 0x7c, 0x65, 0x4c, 0xb6,
	0x0a, 0xe4, 0x1d, 0x00, 0x47, 0x76, 0xfb, 0x15, 0x31, 0x06, 0x75, 0xa1, 0x85, 0x99, 0xdb, 0xb6,
	0x32, 0x1f, 0x8c, 0x41, 0x2e, 0x61, 0x91, 0x1c, 0xd3, 0xb8, 0x55, 0x8d, 0x16, 0x66, 0x66, 0x8f,
	0xdb, 0x64, 0xdb, 0x24, 0x74, 0x8c, 0x7e, 0xe5, 0x58, 0xcd, 0xb5, 0x30, 0x8d, 0x6d, 0x8f, 0xcc,
	0x1b, 0xbf, 0x3f, 0x7d, 0x3f, 0x6e, 0x89, 0x87, 0xb2, 0xee, 0x36, 0x31, 0xf4, 0xb4, 0x1f, 0x5c,
	0x08, 0x83, 0xf7, 0xfd, 0xae, 0x78, 0xb7, 0xa3, 0xe7, 0xd3, 0x37, 0x87, 0xee, 0xeb, 0x45, 0xad,
	0xf5, 0xfa, 0x1f, 0x00, 0x00, 0xff, 0xff, 0xb7, 0xfe, 0x0f, 0x2a, 0x11, 0x01, 0x00, 0x00,
}
