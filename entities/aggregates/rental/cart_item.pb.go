// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/rental/cart_item.proto

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

type CartItem struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid                 string               `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Cart                 *Cart                `protobuf:"bytes,3,opt,name=cart,proto3" json:"cart,omitempty"`
	ItemUuid             string               `protobuf:"bytes,4,opt,name=item_uuid,json=itemUuid,proto3" json:"item_uuid,omitempty"`
	TotalItem            int32                `protobuf:"varint,5,opt,name=total_item,json=totalItem,proto3" json:"total_item,omitempty"`
	TotalPrice           int32                `protobuf:"varint,6,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	UnitPrice            int32                `protobuf:"varint,7,opt,name=unit_price,json=unitPrice,proto3" json:"unit_price,omitempty"`
	DurationFrom         *timestamp.Timestamp `protobuf:"bytes,8,opt,name=duration_from,json=durationFrom,proto3" json:"duration_from,omitempty"`
	DurationTo           *timestamp.Timestamp `protobuf:"bytes,9,opt,name=duration_to,json=durationTo,proto3" json:"duration_to,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CartItem) Reset()         { *m = CartItem{} }
func (m *CartItem) String() string { return proto.CompactTextString(m) }
func (*CartItem) ProtoMessage()    {}
func (*CartItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_8026613f915ecff3, []int{0}
}

func (m *CartItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CartItem.Unmarshal(m, b)
}
func (m *CartItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CartItem.Marshal(b, m, deterministic)
}
func (m *CartItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CartItem.Merge(m, src)
}
func (m *CartItem) XXX_Size() int {
	return xxx_messageInfo_CartItem.Size(m)
}
func (m *CartItem) XXX_DiscardUnknown() {
	xxx_messageInfo_CartItem.DiscardUnknown(m)
}

var xxx_messageInfo_CartItem proto.InternalMessageInfo

func (m *CartItem) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CartItem) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *CartItem) GetCart() *Cart {
	if m != nil {
		return m.Cart
	}
	return nil
}

func (m *CartItem) GetItemUuid() string {
	if m != nil {
		return m.ItemUuid
	}
	return ""
}

func (m *CartItem) GetTotalItem() int32 {
	if m != nil {
		return m.TotalItem
	}
	return 0
}

func (m *CartItem) GetTotalPrice() int32 {
	if m != nil {
		return m.TotalPrice
	}
	return 0
}

func (m *CartItem) GetUnitPrice() int32 {
	if m != nil {
		return m.UnitPrice
	}
	return 0
}

func (m *CartItem) GetDurationFrom() *timestamp.Timestamp {
	if m != nil {
		return m.DurationFrom
	}
	return nil
}

func (m *CartItem) GetDurationTo() *timestamp.Timestamp {
	if m != nil {
		return m.DurationTo
	}
	return nil
}

func (m *CartItem) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*CartItem)(nil), "aggregates.rental.CartItem")
}

func init() { proto.RegisterFile("aggregates/rental/cart_item.proto", fileDescriptor_8026613f915ecff3) }

var fileDescriptor_8026613f915ecff3 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcf, 0x6a, 0xe3, 0x30,
	0x18, 0xc4, 0x71, 0xfe, 0x6d, 0xac, 0xec, 0x2e, 0xac, 0x2e, 0x2b, 0xb2, 0xbb, 0xc4, 0xdb, 0x93,
	0xa1, 0x44, 0x82, 0xf6, 0x54, 0x7a, 0x28, 0x6d, 0x21, 0xb4, 0xb7, 0x62, 0xd2, 0x4b, 0x2f, 0x46,
	0xb1, 0x14, 0x47, 0xc4, 0xb2, 0x82, 0xfc, 0xb9, 0xa5, 0x6f, 0xd1, 0x47, 0x2e, 0x92, 0xe2, 0xf6,
	0x90, 0x42, 0x6e, 0xe6, 0x9b, 0xf9, 0x8d, 0xc7, 0x63, 0xf4, 0x9f, 0x97, 0xa5, 0x95, 0x25, 0x07,
	0xd9, 0x30, 0x2b, 0x6b, 0xe0, 0x15, 0x2b, 0xb8, 0x85, 0x5c, 0x81, 0xd4, 0x74, 0x67, 0x0d, 0x18,
	0xfc, 0xeb, 0xd3, 0x42, 0x83, 0x65, 0x3a, 0x2b, 0x8d, 0x29, 0x2b, 0xc9, 0xbc, 0x61, 0xd5, 0xae,
	0x19, 0x28, 0x2d, 0x1b, 0xe0, 0x7a, 0x17, 0x98, 0xe9, 0xdf, 0xaf, 0x63, 0x83, 0x7a, 0xf2, 0xd6,
	0x47, 0xe3, 0x5b, 0x6e, 0xe1, 0x1e, 0xa4, 0xc6, 0x3f, 0x51, 0x4f, 0x09, 0x12, 0x25, 0x51, 0xda,
	0xcf, 0x7a, 0x4a, 0x60, 0x8c, 0x06, 0x6d, 0xab, 0x04, 0xe9, 0x25, 0x51, 0x1a, 0x67, 0xfe, 0x19,
	0x9f, 0xa2, 0x81, 0xc3, 0x49, 0x3f, 0x89, 0xd2, 0xc9, 0xd9, 0x6f, 0x7a, 0xd0, 0x88, 0xba, 0xb8,
	0xcc, 0x9b, 0xf0, 0x1f, 0x14, 0xbb, 0xf6, 0xb9, 0x4f, 0x19, 0xf8, 0x94, 0xb1, 0x3b, 0x3c, 0xba,
	0xa4, 0x7f, 0x08, 0x81, 0x01, 0x5e, 0xf9, 0x0f, 0x24, 0xc3, 0x24, 0x4a, 0x87, 0x59, 0xec, 0x2f,
	0xbe, 0xcc, 0x0c, 0x4d, 0x82, 0xbc, 0xb3, 0xaa, 0x90, 0x64, 0xe4, 0xf5, 0x40, 0x3c, 0xb8, 0x8b,
	0xe3, 0xdb, 0x5a, 0xc1, 0x5e, 0xff, 0x16, 0x78, 0x77, 0x09, 0xf2, 0x15, 0xfa, 0x21, 0x5a, 0xcb,
	0x41, 0x99, 0x3a, 0x5f, 0x5b, 0xa3, 0xc9, 0xd8, 0x37, 0x9e, 0xd2, 0x30, 0x18, 0xed, 0x06, 0xa3,
	0xcb, 0x6e, 0xb0, 0xec, 0x7b, 0x07, 0x2c, 0xac, 0xd1, 0xf8, 0x12, 0x4d, 0x3e, 0x02, 0xc0, 0x90,
	0xf8, 0x28, 0x8e, 0x3a, 0xfb, 0xd2, 0xe0, 0x0b, 0x84, 0x0a, 0x2b, 0x39, 0x48, 0x91, 0x73, 0x20,
	0xe8, 0x28, 0x1b, 0xef, 0xdd, 0xd7, 0x70, 0x73, 0xf7, 0xb4, 0x28, 0x15, 0x6c, 0xda, 0x15, 0x2d,
	0x8c, 0x66, 0xaa, 0xda, 0x70, 0xad, 0x37, 0x42, 0xb0, 0x6d, 0x2b, 0xf8, 0x56, 0xcd, 0xdd, 0xd6,
	0x2f, 0xfc, 0x75, 0xde, 0x48, 0xfb, 0xac, 0x0a, 0xc9, 0x64, 0x0d, 0x0a, 0x94, 0x6c, 0xd8, 0xc1,
	0x7f, 0x5e, 0x8d, 0xfc, 0x8b, 0xce, 0xdf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xb3, 0xd6, 0x01, 0x6c,
	0x5a, 0x02, 0x00, 0x00,
}
