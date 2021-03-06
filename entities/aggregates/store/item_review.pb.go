// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/store/item_review.proto

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

type ItemReview struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid                 string               `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	UserUuid             string               `protobuf:"bytes,3,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	Item                 *Item                `protobuf:"bytes,4,opt,name=item,proto3" json:"item,omitempty"`
	Review               string               `protobuf:"bytes,5,opt,name=review,proto3" json:"review,omitempty"`
	Rating               float64              `protobuf:"fixed64,6,opt,name=rating,proto3" json:"rating,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ItemReview) Reset()         { *m = ItemReview{} }
func (m *ItemReview) String() string { return proto.CompactTextString(m) }
func (*ItemReview) ProtoMessage()    {}
func (*ItemReview) Descriptor() ([]byte, []int) {
	return fileDescriptor_3806212aa403f188, []int{0}
}

func (m *ItemReview) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemReview.Unmarshal(m, b)
}
func (m *ItemReview) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemReview.Marshal(b, m, deterministic)
}
func (m *ItemReview) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemReview.Merge(m, src)
}
func (m *ItemReview) XXX_Size() int {
	return xxx_messageInfo_ItemReview.Size(m)
}
func (m *ItemReview) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemReview.DiscardUnknown(m)
}

var xxx_messageInfo_ItemReview proto.InternalMessageInfo

func (m *ItemReview) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ItemReview) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *ItemReview) GetUserUuid() string {
	if m != nil {
		return m.UserUuid
	}
	return ""
}

func (m *ItemReview) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *ItemReview) GetReview() string {
	if m != nil {
		return m.Review
	}
	return ""
}

func (m *ItemReview) GetRating() float64 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func (m *ItemReview) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*ItemReview)(nil), "aggregates.store.ItemReview")
}

func init() { proto.RegisterFile("aggregates/store/item_review.proto", fileDescriptor_3806212aa403f188) }

var fileDescriptor_3806212aa403f188 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0x3b, 0x31,
	0x10, 0xc5, 0x49, 0xdb, 0x7f, 0xff, 0x36, 0x82, 0x48, 0x0e, 0x65, 0x69, 0x0f, 0x2e, 0x3d, 0x2d,
	0x42, 0x13, 0xd0, 0x93, 0x47, 0x05, 0x11, 0xaf, 0x8b, 0x5e, 0xbc, 0x94, 0xb4, 0x19, 0xd3, 0xa1,
	0x4d, 0xb7, 0x64, 0x27, 0x2d, 0x7e, 0x5f, 0x3f, 0x88, 0xec, 0xec, 0x16, 0xa1, 0x78, 0xdb, 0x99,
	0xdf, 0xdb, 0xf7, 0xe6, 0x45, 0xce, 0xac, 0xf7, 0x11, 0xbc, 0x25, 0xa8, 0x4d, 0x4d, 0x55, 0x04,
	0x83, 0x04, 0x61, 0x11, 0xe1, 0x80, 0x70, 0xd4, 0xfb, 0x58, 0x51, 0xa5, 0xae, 0x7f, 0x35, 0x9a,
	0x35, 0x93, 0x1b, 0x5f, 0x55, 0x7e, 0x0b, 0x86, 0xf9, 0x32, 0x7d, 0x1a, 0xc2, 0x00, 0x35, 0xd9,
	0xb0, 0x6f, 0x7f, 0x99, 0x4c, 0xff, 0xb4, 0x6d, 0xe1, 0xec, 0x5b, 0x48, 0xf9, 0x4a, 0x10, 0x4a,
	0x0e, 0x51, 0x57, 0xb2, 0x87, 0x2e, 0x13, 0xb9, 0x28, 0xfa, 0x65, 0x0f, 0x9d, 0x52, 0x72, 0x90,
	0x12, 0xba, 0xac, 0x97, 0x8b, 0x62, 0x54, 0xf2, 0xb7, 0x9a, 0xca, 0x51, 0xaa, 0x21, 0x2e, 0x18,
	0xf4, 0x19, 0x5c, 0x34, 0x8b, 0xf7, 0x06, 0xde, 0xca, 0x41, 0xe3, 0x9e, 0x0d, 0x72, 0x51, 0x5c,
	0xde, 0x8d, 0xf5, 0xf9, 0xb9, 0x9a, 0xc3, 0x58, 0xa3, 0xc6, 0x72, 0xd8, 0x76, 0xcb, 0xfe, 0xb1,
	0x4b, 0x37, 0xf1, 0xde, 0x12, 0xee, 0x7c, 0x36, 0xcc, 0x45, 0x21, 0xca, 0x6e, 0x52, 0x0f, 0x52,
	0xae, 0x22, 0x58, 0x02, 0xb7, 0xb0, 0x94, 0xfd, 0xe7, 0x84, 0x89, 0x6e, 0xeb, 0xeb, 0x53, 0x7d,
	0xfd, 0x76, 0xaa, 0x5f, 0x8e, 0x3a, 0xf5, 0x23, 0x3d, 0xbd, 0x7c, 0x3c, 0x7b, 0xa4, 0x75, 0x5a,
	0xea, 0x55, 0x15, 0x0c, 0x6e, 0xd7, 0x36, 0x84, 0xb5, 0x73, 0x66, 0x93, 0x9c, 0xdd, 0xe0, 0xbc,
	0x39, 0xf0, 0x68, 0xbf, 0xe6, 0x35, 0xc4, 0x03, 0xae, 0xc0, 0xc0, 0x8e, 0x90, 0x10, 0x6a, 0x73,
	0xfe, 0x74, 0xcb, 0x21, 0xe7, 0xdc, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x3f, 0xc4, 0x1d, 0xfa,
	0xac, 0x01, 0x00, 0x00,
}
