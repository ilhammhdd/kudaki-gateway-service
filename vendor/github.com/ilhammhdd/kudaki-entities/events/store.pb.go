// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events/store.proto

package events

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	store "github.com/ilhammhdd/kudaki-entities/store"
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

type AddStorefrontItemRequested struct {
	Uid                  string      `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Item                 *store.Item `protobuf:"bytes,3,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AddStorefrontItemRequested) Reset()         { *m = AddStorefrontItemRequested{} }
func (m *AddStorefrontItemRequested) String() string { return proto.CompactTextString(m) }
func (*AddStorefrontItemRequested) ProtoMessage()    {}
func (*AddStorefrontItemRequested) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f52bba9433e5948, []int{0}
}

func (m *AddStorefrontItemRequested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddStorefrontItemRequested.Unmarshal(m, b)
}
func (m *AddStorefrontItemRequested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddStorefrontItemRequested.Marshal(b, m, deterministic)
}
func (m *AddStorefrontItemRequested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddStorefrontItemRequested.Merge(m, src)
}
func (m *AddStorefrontItemRequested) XXX_Size() int {
	return xxx_messageInfo_AddStorefrontItemRequested.Size(m)
}
func (m *AddStorefrontItemRequested) XXX_DiscardUnknown() {
	xxx_messageInfo_AddStorefrontItemRequested.DiscardUnknown(m)
}

var xxx_messageInfo_AddStorefrontItemRequested proto.InternalMessageInfo

func (m *AddStorefrontItemRequested) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *AddStorefrontItemRequested) GetItem() *store.Item {
	if m != nil {
		return m.Item
	}
	return nil
}

type StorefrontItemAdded struct {
	Uid                  string      `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Item                 *store.Item `protobuf:"bytes,3,opt,name=item,proto3" json:"item,omitempty"`
	EventStatus          *Status     `protobuf:"bytes,4,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *StorefrontItemAdded) Reset()         { *m = StorefrontItemAdded{} }
func (m *StorefrontItemAdded) String() string { return proto.CompactTextString(m) }
func (*StorefrontItemAdded) ProtoMessage()    {}
func (*StorefrontItemAdded) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f52bba9433e5948, []int{1}
}

func (m *StorefrontItemAdded) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StorefrontItemAdded.Unmarshal(m, b)
}
func (m *StorefrontItemAdded) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StorefrontItemAdded.Marshal(b, m, deterministic)
}
func (m *StorefrontItemAdded) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorefrontItemAdded.Merge(m, src)
}
func (m *StorefrontItemAdded) XXX_Size() int {
	return xxx_messageInfo_StorefrontItemAdded.Size(m)
}
func (m *StorefrontItemAdded) XXX_DiscardUnknown() {
	xxx_messageInfo_StorefrontItemAdded.DiscardUnknown(m)
}

var xxx_messageInfo_StorefrontItemAdded proto.InternalMessageInfo

func (m *StorefrontItemAdded) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *StorefrontItemAdded) GetItem() *store.Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *StorefrontItemAdded) GetEventStatus() *Status {
	if m != nil {
		return m.EventStatus
	}
	return nil
}

type DeleteStorefrontItemRequested struct {
	Uid                  string      `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Item                 *store.Item `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DeleteStorefrontItemRequested) Reset()         { *m = DeleteStorefrontItemRequested{} }
func (m *DeleteStorefrontItemRequested) String() string { return proto.CompactTextString(m) }
func (*DeleteStorefrontItemRequested) ProtoMessage()    {}
func (*DeleteStorefrontItemRequested) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f52bba9433e5948, []int{2}
}

func (m *DeleteStorefrontItemRequested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteStorefrontItemRequested.Unmarshal(m, b)
}
func (m *DeleteStorefrontItemRequested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteStorefrontItemRequested.Marshal(b, m, deterministic)
}
func (m *DeleteStorefrontItemRequested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteStorefrontItemRequested.Merge(m, src)
}
func (m *DeleteStorefrontItemRequested) XXX_Size() int {
	return xxx_messageInfo_DeleteStorefrontItemRequested.Size(m)
}
func (m *DeleteStorefrontItemRequested) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteStorefrontItemRequested.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteStorefrontItemRequested proto.InternalMessageInfo

func (m *DeleteStorefrontItemRequested) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *DeleteStorefrontItemRequested) GetItem() *store.Item {
	if m != nil {
		return m.Item
	}
	return nil
}

type StorefrontItemDeleted struct {
	Uid                  string      `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Item                 *store.Item `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	EventStatus          *Status     `protobuf:"bytes,3,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *StorefrontItemDeleted) Reset()         { *m = StorefrontItemDeleted{} }
func (m *StorefrontItemDeleted) String() string { return proto.CompactTextString(m) }
func (*StorefrontItemDeleted) ProtoMessage()    {}
func (*StorefrontItemDeleted) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f52bba9433e5948, []int{3}
}

func (m *StorefrontItemDeleted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StorefrontItemDeleted.Unmarshal(m, b)
}
func (m *StorefrontItemDeleted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StorefrontItemDeleted.Marshal(b, m, deterministic)
}
func (m *StorefrontItemDeleted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorefrontItemDeleted.Merge(m, src)
}
func (m *StorefrontItemDeleted) XXX_Size() int {
	return xxx_messageInfo_StorefrontItemDeleted.Size(m)
}
func (m *StorefrontItemDeleted) XXX_DiscardUnknown() {
	xxx_messageInfo_StorefrontItemDeleted.DiscardUnknown(m)
}

var xxx_messageInfo_StorefrontItemDeleted proto.InternalMessageInfo

func (m *StorefrontItemDeleted) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *StorefrontItemDeleted) GetItem() *store.Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *StorefrontItemDeleted) GetEventStatus() *Status {
	if m != nil {
		return m.EventStatus
	}
	return nil
}

type RetrieveStorefrontItemsRequested struct {
	Uid                  string     `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	User                 *user.User `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	From                 int32      `protobuf:"varint,3,opt,name=from,proto3" json:"from,omitempty"`
	Limit                int32      `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RetrieveStorefrontItemsRequested) Reset()         { *m = RetrieveStorefrontItemsRequested{} }
func (m *RetrieveStorefrontItemsRequested) String() string { return proto.CompactTextString(m) }
func (*RetrieveStorefrontItemsRequested) ProtoMessage()    {}
func (*RetrieveStorefrontItemsRequested) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f52bba9433e5948, []int{4}
}

func (m *RetrieveStorefrontItemsRequested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrieveStorefrontItemsRequested.Unmarshal(m, b)
}
func (m *RetrieveStorefrontItemsRequested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrieveStorefrontItemsRequested.Marshal(b, m, deterministic)
}
func (m *RetrieveStorefrontItemsRequested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveStorefrontItemsRequested.Merge(m, src)
}
func (m *RetrieveStorefrontItemsRequested) XXX_Size() int {
	return xxx_messageInfo_RetrieveStorefrontItemsRequested.Size(m)
}
func (m *RetrieveStorefrontItemsRequested) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveStorefrontItemsRequested.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveStorefrontItemsRequested proto.InternalMessageInfo

func (m *RetrieveStorefrontItemsRequested) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RetrieveStorefrontItemsRequested) GetUser() *user.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *RetrieveStorefrontItemsRequested) GetFrom() int32 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *RetrieveStorefrontItemsRequested) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type StorefrontItemsRetrieved struct {
	Uid                  string       `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Items                *store.Items `protobuf:"bytes,2,opt,name=items,proto3" json:"items,omitempty"`
	EventStatus          *Status      `protobuf:"bytes,3,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	From                 int32        `protobuf:"varint,4,opt,name=from,proto3" json:"from,omitempty"`
	Limit                int32        `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	To                   int32        `protobuf:"varint,6,opt,name=to,proto3" json:"to,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *StorefrontItemsRetrieved) Reset()         { *m = StorefrontItemsRetrieved{} }
func (m *StorefrontItemsRetrieved) String() string { return proto.CompactTextString(m) }
func (*StorefrontItemsRetrieved) ProtoMessage()    {}
func (*StorefrontItemsRetrieved) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f52bba9433e5948, []int{5}
}

func (m *StorefrontItemsRetrieved) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StorefrontItemsRetrieved.Unmarshal(m, b)
}
func (m *StorefrontItemsRetrieved) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StorefrontItemsRetrieved.Marshal(b, m, deterministic)
}
func (m *StorefrontItemsRetrieved) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorefrontItemsRetrieved.Merge(m, src)
}
func (m *StorefrontItemsRetrieved) XXX_Size() int {
	return xxx_messageInfo_StorefrontItemsRetrieved.Size(m)
}
func (m *StorefrontItemsRetrieved) XXX_DiscardUnknown() {
	xxx_messageInfo_StorefrontItemsRetrieved.DiscardUnknown(m)
}

var xxx_messageInfo_StorefrontItemsRetrieved proto.InternalMessageInfo

func (m *StorefrontItemsRetrieved) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *StorefrontItemsRetrieved) GetItems() *store.Items {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *StorefrontItemsRetrieved) GetEventStatus() *Status {
	if m != nil {
		return m.EventStatus
	}
	return nil
}

func (m *StorefrontItemsRetrieved) GetFrom() int32 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *StorefrontItemsRetrieved) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *StorefrontItemsRetrieved) GetTo() int32 {
	if m != nil {
		return m.To
	}
	return 0
}

func init() {
	proto.RegisterType((*AddStorefrontItemRequested)(nil), "event.AddStorefrontItemRequested")
	proto.RegisterType((*StorefrontItemAdded)(nil), "event.StorefrontItemAdded")
	proto.RegisterType((*DeleteStorefrontItemRequested)(nil), "event.DeleteStorefrontItemRequested")
	proto.RegisterType((*StorefrontItemDeleted)(nil), "event.StorefrontItemDeleted")
	proto.RegisterType((*RetrieveStorefrontItemsRequested)(nil), "event.RetrieveStorefrontItemsRequested")
	proto.RegisterType((*StorefrontItemsRetrieved)(nil), "event.StorefrontItemsRetrieved")
}

func init() { proto.RegisterFile("events/store.proto", fileDescriptor_4f52bba9433e5948) }

var fileDescriptor_4f52bba9433e5948 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xd1, 0x6a, 0xea, 0x40,
	0x10, 0x86, 0x89, 0x26, 0xc2, 0x19, 0xcf, 0x39, 0x95, 0x55, 0x21, 0x08, 0x05, 0xc9, 0x4d, 0x05,
	0x31, 0x29, 0xed, 0x13, 0x58, 0x7a, 0xd3, 0xdb, 0x95, 0x42, 0x69, 0x2f, 0x8a, 0x76, 0xc7, 0xba,
	0x68, 0xdc, 0x36, 0x3b, 0xf1, 0x05, 0x4a, 0xe9, 0x93, 0xf5, 0xbd, 0x4a, 0x66, 0xad, 0xad, 0x55,
	0x2a, 0xe2, 0x4d, 0xd8, 0xf9, 0x33, 0x33, 0xfb, 0xfd, 0xbb, 0x3b, 0x20, 0x70, 0x81, 0x73, 0xb2,
	0x89, 0x25, 0x93, 0x61, 0xfc, 0x94, 0x19, 0x32, 0x22, 0x60, 0xad, 0x55, 0x63, 0x2d, 0xd1, 0x84,
	0xa9, 0xfb, 0xd1, 0xaa, 0xaf, 0x92, 0x87, 0x94, 0xdb, 0xa5, 0x78, 0x94, 0x5b, 0xcc, 0x92, 0xe2,
	0xe3, 0x84, 0xe8, 0x06, 0x5a, 0x7d, 0xa5, 0x06, 0x45, 0xf1, 0x38, 0x33, 0x73, 0xba, 0x22, 0x4c,
	0x25, 0x3e, 0xe7, 0x68, 0x09, 0x95, 0xa8, 0x41, 0x39, 0xd7, 0x2a, 0xf4, 0xda, 0x5e, 0xe7, 0x8f,
	0x2c, 0x96, 0xa2, 0x03, 0x7e, 0xb1, 0x47, 0x58, 0x6e, 0x7b, 0x9d, 0xea, 0x59, 0x23, 0xc6, 0x39,
	0x69, 0xd2, 0x68, 0x63, 0xc7, 0xc4, 0xe5, 0x9c, 0x11, 0xbd, 0x78, 0x50, 0x5f, 0xef, 0xdb, 0x57,
	0xea, 0xb0, 0x9e, 0xe2, 0x14, 0xfe, 0xb2, 0xab, 0x7b, 0x67, 0x2a, 0xf4, 0xb9, 0xe2, 0x5f, 0xcc,
	0x62, 0x3c, 0x60, 0x51, 0x56, 0x39, 0x72, 0x41, 0x74, 0x07, 0xc7, 0x97, 0x38, 0x43, 0xc2, 0xfd,
	0x2d, 0x96, 0x76, 0x5a, 0x7c, 0xf5, 0xa0, 0xb9, 0xde, 0xd7, 0xed, 0x75, 0x50, 0xd7, 0x0d, 0x93,
	0xe5, 0x9d, 0x26, 0xdf, 0x3c, 0x68, 0x4b, 0xa4, 0x4c, 0xe3, 0xe2, 0x87, 0x4f, 0xfb, 0x9b, 0xd1,
	0x13, 0xf0, 0x8b, 0x97, 0xb0, 0x44, 0xaa, 0x7f, 0x21, 0xf1, 0xfb, 0xb8, 0xb6, 0x98, 0x49, 0x4e,
	0x10, 0x02, 0xfc, 0x71, 0x66, 0xdc, 0x05, 0x05, 0x92, 0xd7, 0xa2, 0x01, 0xc1, 0x4c, 0xa7, 0x9a,
	0xf8, 0x0e, 0x02, 0xe9, 0x82, 0xe8, 0xdd, 0x83, 0x70, 0x83, 0xc0, 0x81, 0x6d, 0x23, 0xe8, 0x42,
	0x50, 0x58, 0xb6, 0x4b, 0x84, 0xe6, 0xb6, 0x53, 0xb1, 0xd2, 0xe5, 0xec, 0x7f, 0x2e, 0x2b, 0x6e,
	0x7f, 0x1b, 0x77, 0xf0, 0x8d, 0x5b, 0xfc, 0x87, 0x12, 0x99, 0xb0, 0xc2, 0x52, 0x89, 0xcc, 0x45,
	0xef, 0xb6, 0xfb, 0xa8, 0x69, 0x92, 0x8f, 0xe2, 0x07, 0x93, 0x26, 0x7a, 0x36, 0x19, 0xa6, 0xe9,
	0x44, 0xa9, 0x64, 0x9a, 0xab, 0xe1, 0x54, 0xf7, 0x3e, 0x31, 0x13, 0x37, 0x63, 0xa3, 0x0a, 0x0f,
	0xd3, 0xf9, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0x38, 0x25, 0x18, 0x6d, 0xa1, 0x03, 0x00, 0x00,
}
