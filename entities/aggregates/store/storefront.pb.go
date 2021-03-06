// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/store/storefront.proto

package store

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid                 string               `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	UserUuid             string               `protobuf:"bytes,3,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	TotalItem            int32                `protobuf:"varint,4,opt,name=total_item,json=totalItem,proto3" json:"total_item,omitempty"`
	Rating               float64              `protobuf:"fixed64,5,opt,name=rating,proto3" json:"rating,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Storefront) Reset()         { *m = Storefront{} }
func (m *Storefront) String() string { return proto.CompactTextString(m) }
func (*Storefront) ProtoMessage()    {}
func (*Storefront) Descriptor() ([]byte, []int) {
	return fileDescriptor_11e209133ba8f19a, []int{0}
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

func (m *Storefront) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Storefront) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Storefront) GetUserUuid() string {
	if m != nil {
		return m.UserUuid
	}
	return ""
}

func (m *Storefront) GetTotalItem() int32 {
	if m != nil {
		return m.TotalItem
	}
	return 0
}

func (m *Storefront) GetRating() float64 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func (m *Storefront) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Storefront)(nil), "aggregates.store.Storefront")
}

func init() { proto.RegisterFile("aggregates/store/storefront.proto", fileDescriptor_11e209133ba8f19a) }

var fileDescriptor_11e209133ba8f19a = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0x49, 0xff, 0xe1, 0x46, 0x10, 0xc9, 0x41, 0x96, 0x8a, 0xb8, 0x7a, 0xda, 0x4b, 0x13,
	0xd0, 0x93, 0x47, 0x05, 0x11, 0xaf, 0xab, 0x5e, 0xbc, 0x94, 0xb4, 0x99, 0xa6, 0x43, 0x9b, 0x4d,
	0x49, 0x26, 0x8a, 0x9f, 0xce, 0xaf, 0x26, 0x9b, 0x6d, 0x15, 0xbc, 0x84, 0xe4, 0xfd, 0x5e, 0x66,
	0x1e, 0x8f, 0x5f, 0x69, 0x6b, 0x03, 0x58, 0x4d, 0x10, 0x55, 0x24, 0x1f, 0xa0, 0x3f, 0x57, 0xc1,
	0xb7, 0x24, 0x77, 0xc1, 0x93, 0x17, 0xa7, 0x7f, 0x16, 0x99, 0xe1, 0xf4, 0xd2, 0x7a, 0x6f, 0xb7,
	0xa0, 0x32, 0x5f, 0xa4, 0x95, 0x22, 0x74, 0x10, 0x49, 0xbb, 0x5d, 0xff, 0xe5, 0xfa, 0x9b, 0x71,
	0xfe, 0xf2, 0x3b, 0x47, 0x9c, 0xf0, 0x01, 0x9a, 0x92, 0x55, 0xac, 0x1e, 0x36, 0x03, 0x34, 0x42,
	0xf0, 0x51, 0x4a, 0x68, 0xca, 0x41, 0xc5, 0xea, 0xa2, 0xc9, 0x77, 0x71, 0xce, 0x8b, 0x14, 0x21,
	0xcc, 0x33, 0x18, 0x66, 0x70, 0xd4, 0x09, 0x6f, 0x1d, 0xbc, 0xe0, 0x9c, 0x3c, 0xe9, 0xed, 0x1c,
	0x09, 0x5c, 0x39, 0xaa, 0x58, 0x3d, 0x6e, 0x8a, 0xac, 0x3c, 0x13, 0x38, 0x71, 0xc6, 0x27, 0x41,
	0x13, 0xb6, 0xb6, 0x1c, 0x57, 0xac, 0x66, 0xcd, 0xfe, 0x25, 0xee, 0x38, 0x5f, 0x06, 0xd0, 0x04,
	0x66, 0xae, 0xa9, 0x9c, 0x54, 0xac, 0x3e, 0xbe, 0x99, 0xca, 0x3e, 0xbc, 0x3c, 0x84, 0x97, 0xaf,
	0x87, 0xf0, 0x4d, 0xb1, 0x77, 0xdf, 0xd3, 0xc3, 0xd3, 0xfb, 0xa3, 0x45, 0x5a, 0xa7, 0x85, 0x5c,
	0x7a, 0xa7, 0x70, 0xbb, 0xd6, 0xce, 0xad, 0x8d, 0x51, 0x9b, 0x64, 0xf4, 0x06, 0x67, 0x5d, 0x1b,
	0x9f, 0xfa, 0x6b, 0x16, 0x21, 0x7c, 0xe0, 0x12, 0x14, 0xb4, 0x84, 0x84, 0x10, 0xd5, 0xff, 0x3a,
	0x17, 0x93, 0xbc, 0xe7, 0xf6, 0x27, 0x00, 0x00, 0xff, 0xff, 0x28, 0x37, 0x0a, 0xca, 0x69, 0x01,
	0x00, 0x00,
}
