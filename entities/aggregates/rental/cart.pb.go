// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/rental/cart.proto

package rental

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

type Cart struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid                 string               `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	UserUuid             string               `protobuf:"bytes,3,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	TotalPrice           int32                `protobuf:"varint,4,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	TotalItems           int32                `protobuf:"varint,5,opt,name=total_items,json=totalItems,proto3" json:"total_items,omitempty"`
	Open                 bool                 `protobuf:"varint,6,opt,name=open,proto3" json:"open,omitempty"`
	Delivered            bool                 `protobuf:"varint,7,opt,name=delivered,proto3" json:"delivered,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Cart) Reset()         { *m = Cart{} }
func (m *Cart) String() string { return proto.CompactTextString(m) }
func (*Cart) ProtoMessage()    {}
func (*Cart) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ef36b31966c981a, []int{0}
}

func (m *Cart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cart.Unmarshal(m, b)
}
func (m *Cart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cart.Marshal(b, m, deterministic)
}
func (m *Cart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cart.Merge(m, src)
}
func (m *Cart) XXX_Size() int {
	return xxx_messageInfo_Cart.Size(m)
}
func (m *Cart) XXX_DiscardUnknown() {
	xxx_messageInfo_Cart.DiscardUnknown(m)
}

var xxx_messageInfo_Cart proto.InternalMessageInfo

func (m *Cart) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Cart) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Cart) GetUserUuid() string {
	if m != nil {
		return m.UserUuid
	}
	return ""
}

func (m *Cart) GetTotalPrice() int32 {
	if m != nil {
		return m.TotalPrice
	}
	return 0
}

func (m *Cart) GetTotalItems() int32 {
	if m != nil {
		return m.TotalItems
	}
	return 0
}

func (m *Cart) GetOpen() bool {
	if m != nil {
		return m.Open
	}
	return false
}

func (m *Cart) GetDelivered() bool {
	if m != nil {
		return m.Delivered
	}
	return false
}

func (m *Cart) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Cart)(nil), "aggregates.rental.Cart")
}

func init() { proto.RegisterFile("aggregates/rental/cart.proto", fileDescriptor_7ef36b31966c981a) }

var fileDescriptor_7ef36b31966c981a = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xcf, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xd9, 0xfe, 0xb2, 0xd9, 0x82, 0xe0, 0x9e, 0x96, 0x5a, 0x68, 0xf0, 0x94, 0x4b, 0xb3,
	0xa0, 0x27, 0x8f, 0x2a, 0x88, 0xde, 0x24, 0xe8, 0xc5, 0x4b, 0x99, 0x66, 0xc7, 0x74, 0x68, 0xd2,
	0x0d, 0x9b, 0x49, 0xc5, 0x7f, 0xdc, 0xb3, 0x64, 0x6b, 0x89, 0xe0, 0x6d, 0xf8, 0xbe, 0xb7, 0x0c,
	0x6f, 0x47, 0x2e, 0xa0, 0x28, 0x3c, 0x16, 0xc0, 0xd8, 0x18, 0x8f, 0x7b, 0x86, 0xd2, 0xe4, 0xe0,
	0x39, 0xad, 0xbd, 0x63, 0xa7, 0x2e, 0x7a, 0x9b, 0x1e, 0xed, 0x7c, 0x59, 0x38, 0x57, 0x94, 0x68,
	0x42, 0x60, 0xd3, 0x7e, 0x18, 0xa6, 0x0a, 0x1b, 0x86, 0xaa, 0x3e, 0xbe, 0xb9, 0xfa, 0x16, 0x72,
	0xf4, 0x00, 0x9e, 0xd5, 0xb9, 0x1c, 0x90, 0xd5, 0x22, 0x16, 0xc9, 0x30, 0x1b, 0x90, 0x55, 0x4a,
	0x8e, 0xda, 0x96, 0xac, 0x1e, 0xc4, 0x22, 0x89, 0xb2, 0x30, 0xab, 0x4b, 0x19, 0xb5, 0x0d, 0xfa,
	0x75, 0x10, 0xc3, 0x20, 0xa6, 0x1d, 0x78, 0xeb, 0xe4, 0x52, 0xce, 0xd8, 0x31, 0x94, 0xeb, 0xda,
	0x53, 0x8e, 0x7a, 0x14, 0x8b, 0x64, 0x9c, 0xc9, 0x80, 0x5e, 0x3a, 0xd2, 0x07, 0x88, 0xb1, 0x6a,
	0xf4, 0xf8, 0x4f, 0xe0, 0xb9, 0x23, 0xdd, 0x4a, 0x57, 0xe3, 0x5e, 0x4f, 0x62, 0x91, 0x4c, 0xb3,
	0x30, 0xab, 0x85, 0x8c, 0x2c, 0x96, 0x74, 0x40, 0x8f, 0x56, 0x9f, 0x05, 0xd1, 0x03, 0x75, 0x2b,
	0x65, 0xee, 0x11, 0x18, 0xed, 0x1a, 0x58, 0x4f, 0x63, 0x91, 0xcc, 0xae, 0xe7, 0xe9, 0xb1, 0x73,
	0x7a, 0xea, 0x9c, 0xbe, 0x9e, 0x3a, 0x67, 0xd1, 0x6f, 0xfa, 0x8e, 0xef, 0x9f, 0xde, 0x1f, 0x0b,
	0xe2, 0x6d, 0xbb, 0x49, 0x73, 0x57, 0x19, 0x2a, 0xb7, 0x50, 0x55, 0x5b, 0x6b, 0xcd, 0xae, 0xb5,
	0xb0, 0xa3, 0x55, 0xf7, 0x8b, 0x9f, 0xf0, 0xb5, 0x6a, 0xd0, 0x1f, 0x28, 0x47, 0x83, 0x7b, 0x26,
	0x26, 0x6c, 0xcc, 0xbf, 0x0b, 0x6c, 0x26, 0x61, 0xd1, 0xcd, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x26, 0x05, 0x55, 0x9c, 0x9d, 0x01, 0x00, 0x00,
}
