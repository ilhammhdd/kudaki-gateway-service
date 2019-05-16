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

type UpdateStorefrontItemRequested struct {
	Uid                  string      `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Item                 *store.Item `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	User                 *user.User  `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UpdateStorefrontItemRequested) Reset()         { *m = UpdateStorefrontItemRequested{} }
func (m *UpdateStorefrontItemRequested) String() string { return proto.CompactTextString(m) }
func (*UpdateStorefrontItemRequested) ProtoMessage()    {}
func (*UpdateStorefrontItemRequested) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f52bba9433e5948, []int{4}
}

func (m *UpdateStorefrontItemRequested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateStorefrontItemRequested.Unmarshal(m, b)
}
func (m *UpdateStorefrontItemRequested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateStorefrontItemRequested.Marshal(b, m, deterministic)
}
func (m *UpdateStorefrontItemRequested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateStorefrontItemRequested.Merge(m, src)
}
func (m *UpdateStorefrontItemRequested) XXX_Size() int {
	return xxx_messageInfo_UpdateStorefrontItemRequested.Size(m)
}
func (m *UpdateStorefrontItemRequested) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateStorefrontItemRequested.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateStorefrontItemRequested proto.InternalMessageInfo

func (m *UpdateStorefrontItemRequested) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *UpdateStorefrontItemRequested) GetItem() *store.Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *UpdateStorefrontItemRequested) GetUser() *user.User {
	if m != nil {
		return m.User
	}
	return nil
}

type StorefrontItemUpdated struct {
	Uid                  string      `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Item                 *store.Item `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	User                 *user.User  `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	EventStatus          *Status     `protobuf:"bytes,4,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *StorefrontItemUpdated) Reset()         { *m = StorefrontItemUpdated{} }
func (m *StorefrontItemUpdated) String() string { return proto.CompactTextString(m) }
func (*StorefrontItemUpdated) ProtoMessage()    {}
func (*StorefrontItemUpdated) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f52bba9433e5948, []int{5}
}

func (m *StorefrontItemUpdated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StorefrontItemUpdated.Unmarshal(m, b)
}
func (m *StorefrontItemUpdated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StorefrontItemUpdated.Marshal(b, m, deterministic)
}
func (m *StorefrontItemUpdated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorefrontItemUpdated.Merge(m, src)
}
func (m *StorefrontItemUpdated) XXX_Size() int {
	return xxx_messageInfo_StorefrontItemUpdated.Size(m)
}
func (m *StorefrontItemUpdated) XXX_DiscardUnknown() {
	xxx_messageInfo_StorefrontItemUpdated.DiscardUnknown(m)
}

var xxx_messageInfo_StorefrontItemUpdated proto.InternalMessageInfo

func (m *StorefrontItemUpdated) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *StorefrontItemUpdated) GetItem() *store.Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *StorefrontItemUpdated) GetUser() *user.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *StorefrontItemUpdated) GetEventStatus() *Status {
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
	return fileDescriptor_4f52bba9433e5948, []int{6}
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
	First                int32        `protobuf:"varint,4,opt,name=first,proto3" json:"first,omitempty"`
	Limit                int32        `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	Last                 int32        `protobuf:"varint,6,opt,name=last,proto3" json:"last,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *StorefrontItemsRetrieved) Reset()         { *m = StorefrontItemsRetrieved{} }
func (m *StorefrontItemsRetrieved) String() string { return proto.CompactTextString(m) }
func (*StorefrontItemsRetrieved) ProtoMessage()    {}
func (*StorefrontItemsRetrieved) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f52bba9433e5948, []int{7}
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

func (m *StorefrontItemsRetrieved) GetFirst() int32 {
	if m != nil {
		return m.First
	}
	return 0
}

func (m *StorefrontItemsRetrieved) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *StorefrontItemsRetrieved) GetLast() int32 {
	if m != nil {
		return m.Last
	}
	return 0
}

func init() {
	proto.RegisterType((*AddStorefrontItemRequested)(nil), "event.AddStorefrontItemRequested")
	proto.RegisterType((*StorefrontItemAdded)(nil), "event.StorefrontItemAdded")
	proto.RegisterType((*DeleteStorefrontItemRequested)(nil), "event.DeleteStorefrontItemRequested")
	proto.RegisterType((*StorefrontItemDeleted)(nil), "event.StorefrontItemDeleted")
	proto.RegisterType((*UpdateStorefrontItemRequested)(nil), "event.UpdateStorefrontItemRequested")
	proto.RegisterType((*StorefrontItemUpdated)(nil), "event.StorefrontItemUpdated")
	proto.RegisterType((*RetrieveStorefrontItemsRequested)(nil), "event.RetrieveStorefrontItemsRequested")
	proto.RegisterType((*StorefrontItemsRetrieved)(nil), "event.StorefrontItemsRetrieved")
}

func init() { proto.RegisterFile("events/store.proto", fileDescriptor_4f52bba9433e5948) }

var fileDescriptor_4f52bba9433e5948 = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0xdf, 0x6a, 0x1a, 0x41,
	0x14, 0xc6, 0x19, 0xdd, 0x15, 0x3a, 0xb6, 0x54, 0x46, 0x85, 0x45, 0x10, 0x64, 0x6f, 0x2a, 0x88,
	0xbb, 0xa5, 0x7d, 0x02, 0x4b, 0x6f, 0x7a, 0x3b, 0x22, 0x94, 0xf6, 0x22, 0xac, 0x99, 0x63, 0x1c,
	0xdc, 0x75, 0xcd, 0xcc, 0xac, 0x2f, 0x90, 0x84, 0xbc, 0x4d, 0x5e, 0x24, 0x2f, 0x15, 0xe6, 0x8c,
	0x7f, 0x62, 0x94, 0x88, 0x91, 0xe4, 0x66, 0x99, 0xf3, 0xed, 0x99, 0x73, 0x7e, 0xdf, 0xe1, 0x30,
	0x94, 0xc1, 0x12, 0xe6, 0x46, 0xc7, 0xda, 0xe4, 0x0a, 0xa2, 0x85, 0xca, 0x4d, 0xce, 0x7c, 0xd4,
	0x5a, 0x35, 0xd4, 0x62, 0x69, 0x20, 0x73, 0x3f, 0x5a, 0xf5, 0x4d, 0x72, 0x62, 0x0a, 0xbd, 0x12,
	0xbf, 0x16, 0x1a, 0x54, 0x6c, 0x3f, 0x4e, 0x08, 0xff, 0xd2, 0xd6, 0x40, 0x88, 0xa1, 0xbd, 0x3c,
	0x51, 0xf9, 0xdc, 0xfc, 0x31, 0x90, 0x71, 0xb8, 0x2e, 0x40, 0x1b, 0x10, 0xac, 0x46, 0xcb, 0x85,
	0x14, 0x01, 0xe9, 0x90, 0xee, 0x27, 0x6e, 0x8f, 0xac, 0x4b, 0x3d, 0xdb, 0x23, 0x28, 0x77, 0x48,
	0xb7, 0xfa, 0xa3, 0x11, 0xc1, 0xdc, 0x48, 0x23, 0x41, 0x47, 0x8e, 0x09, 0xaf, 0x63, 0x46, 0x78,
	0x43, 0x68, 0x7d, 0xb7, 0xee, 0x40, 0x88, 0xf3, 0x6a, 0xb2, 0xef, 0xf4, 0x33, 0xba, 0xba, 0x70,
	0xa6, 0x02, 0x0f, 0x6f, 0x7c, 0x89, 0x50, 0x8c, 0x86, 0x28, 0xf2, 0x2a, 0x46, 0x2e, 0x08, 0xff,
	0xd3, 0xf6, 0x6f, 0x48, 0xc1, 0xc0, 0xe9, 0x16, 0x4b, 0x47, 0x2d, 0xde, 0x11, 0xda, 0xdc, 0xad,
	0xeb, 0x7a, 0x9d, 0x55, 0x75, 0xcf, 0x64, 0xf9, 0xa8, 0xc9, 0x5b, 0x42, 0xdb, 0xa3, 0x85, 0x48,
	0xde, 0xc5, 0x25, 0xfb, 0x46, 0x3d, 0xbb, 0x30, 0x2b, 0x8e, 0xfa, 0x36, 0x13, 0xd7, 0x68, 0xa4,
	0x41, 0x71, 0x4c, 0x08, 0x1f, 0xf6, 0xc6, 0xe1, 0xa0, 0x3e, 0xa6, 0xfd, 0x1b, 0x96, 0xe3, 0x9e,
	0xd0, 0x0e, 0x07, 0xa3, 0x24, 0x2c, 0x5f, 0x4c, 0x4e, 0xbf, 0x36, 0xba, 0x35, 0x51, 0xe9, 0x18,
	0x11, 0xa3, 0xde, 0x44, 0xe5, 0x6e, 0xb1, 0x7d, 0x8e, 0x67, 0xd6, 0xa0, 0x7e, 0x2a, 0x33, 0x69,
	0x10, 0xcf, 0xe7, 0x2e, 0x08, 0x1f, 0x09, 0x0d, 0xf6, 0x08, 0x1c, 0xd8, 0x21, 0x82, 0x1e, 0xf5,
	0xed, 0x6c, 0xf4, 0x0a, 0xa1, 0x79, 0x68, 0x7c, 0x9a, 0xbb, 0x9c, 0xd3, 0xf7, 0xc9, 0x32, 0x4e,
	0xa4, 0xd2, 0x1b, 0x46, 0x0c, 0xb6, 0xe4, 0xfe, 0x33, 0x72, 0xeb, 0x31, 0x4d, 0xb4, 0x09, 0x2a,
	0xce, 0xa3, 0x3d, 0xff, 0xea, 0xff, 0xeb, 0x5d, 0x49, 0x33, 0x2d, 0xc6, 0xd1, 0x65, 0x9e, 0xc5,
	0x32, 0x9d, 0x26, 0x59, 0x36, 0x15, 0x22, 0x9e, 0x15, 0x22, 0x99, 0xc9, 0xfe, 0x1a, 0x36, 0x76,
	0x2f, 0xd4, 0xb8, 0x82, 0x4f, 0xd1, 0xcf, 0xa7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x75, 0x56, 0x12,
	0x4b, 0xdf, 0x04, 0x00, 0x00,
}
