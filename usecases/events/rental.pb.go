// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events/rental.proto

package events

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	rental "github.com/ilhammhdd/kudaki-gateway-service/entities/aggregates/rental"
	user "github.com/ilhammhdd/kudaki-gateway-service/entities/aggregates/user"
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

//
//rental service command
type RentalServiceCommandTopic int32

const (
	RentalServiceCommandTopic_ADD_CART_ITEM       RentalServiceCommandTopic = 0
	RentalServiceCommandTopic_RETRIEVE_CART_ITEMS RentalServiceCommandTopic = 1
	RentalServiceCommandTopic_DELETE_CART_ITEM    RentalServiceCommandTopic = 2
	RentalServiceCommandTopic_UPDATE_CART_ITEM    RentalServiceCommandTopic = 3
)

var RentalServiceCommandTopic_name = map[int32]string{
	0: "ADD_CART_ITEM",
	1: "RETRIEVE_CART_ITEMS",
	2: "DELETE_CART_ITEM",
	3: "UPDATE_CART_ITEM",
}

var RentalServiceCommandTopic_value = map[string]int32{
	"ADD_CART_ITEM":       0,
	"RETRIEVE_CART_ITEMS": 1,
	"DELETE_CART_ITEM":    2,
	"UPDATE_CART_ITEM":    3,
}

func (x RentalServiceCommandTopic) String() string {
	return proto.EnumName(RentalServiceCommandTopic_name, int32(x))
}

func (RentalServiceCommandTopic) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a8d01722d54fb2c6, []int{0}
}

//
//rental service events
type RentalServiceEventTopic int32

const (
	RentalServiceEventTopic_CART_ITEM_ADDED      RentalServiceEventTopic = 0
	RentalServiceEventTopic_CART_ITEMS_RETRIEVED RentalServiceEventTopic = 1
	RentalServiceEventTopic_CART_ITEM_DELETED    RentalServiceEventTopic = 2
	RentalServiceEventTopic_CART_ITEMS_UPDATED   RentalServiceEventTopic = 3
)

var RentalServiceEventTopic_name = map[int32]string{
	0: "CART_ITEM_ADDED",
	1: "CART_ITEMS_RETRIEVED",
	2: "CART_ITEM_DELETED",
	3: "CART_ITEMS_UPDATED",
}

var RentalServiceEventTopic_value = map[string]int32{
	"CART_ITEM_ADDED":      0,
	"CART_ITEMS_RETRIEVED": 1,
	"CART_ITEM_DELETED":    2,
	"CART_ITEMS_UPDATED":   3,
}

func (x RentalServiceEventTopic) String() string {
	return proto.EnumName(RentalServiceEventTopic_name, int32(x))
}

func (RentalServiceEventTopic) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a8d01722d54fb2c6, []int{1}
}

type AddCartItem struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	ItemUuid             string   `protobuf:"bytes,3,opt,name=item_uuid,json=itemUuid,proto3" json:"item_uuid,omitempty"`
	ItemAmount           int32    `protobuf:"varint,4,opt,name=item_amount,json=itemAmount,proto3" json:"item_amount,omitempty"`
	KudakiToken          string   `protobuf:"bytes,5,opt,name=kudaki_token,json=kudakiToken,proto3" json:"kudaki_token,omitempty"`
	DurationFrom         int64    `protobuf:"varint,6,opt,name=duration_from,json=durationFrom,proto3" json:"duration_from,omitempty"`
	Duration             int32    `protobuf:"varint,7,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddCartItem) Reset()         { *m = AddCartItem{} }
func (m *AddCartItem) String() string { return proto.CompactTextString(m) }
func (*AddCartItem) ProtoMessage()    {}
func (*AddCartItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8d01722d54fb2c6, []int{0}
}

func (m *AddCartItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddCartItem.Unmarshal(m, b)
}
func (m *AddCartItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddCartItem.Marshal(b, m, deterministic)
}
func (m *AddCartItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddCartItem.Merge(m, src)
}
func (m *AddCartItem) XXX_Size() int {
	return xxx_messageInfo_AddCartItem.Size(m)
}
func (m *AddCartItem) XXX_DiscardUnknown() {
	xxx_messageInfo_AddCartItem.DiscardUnknown(m)
}

var xxx_messageInfo_AddCartItem proto.InternalMessageInfo

func (m *AddCartItem) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *AddCartItem) GetItemUuid() string {
	if m != nil {
		return m.ItemUuid
	}
	return ""
}

func (m *AddCartItem) GetItemAmount() int32 {
	if m != nil {
		return m.ItemAmount
	}
	return 0
}

func (m *AddCartItem) GetKudakiToken() string {
	if m != nil {
		return m.KudakiToken
	}
	return ""
}

func (m *AddCartItem) GetDurationFrom() int64 {
	if m != nil {
		return m.DurationFrom
	}
	return 0
}

func (m *AddCartItem) GetDuration() int32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

type RetrieveCartItems struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	KudakiToken          string   `protobuf:"bytes,2,opt,name=kudaki_token,json=kudakiToken,proto3" json:"kudaki_token,omitempty"`
	CartUuid             string   `protobuf:"bytes,3,opt,name=cart_uuid,json=cartUuid,proto3" json:"cart_uuid,omitempty"`
	Offset               int32    `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int32    `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	ResultSchema         []byte   `protobuf:"bytes,6,opt,name=result_schema,json=resultSchema,proto3" json:"result_schema,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetrieveCartItems) Reset()         { *m = RetrieveCartItems{} }
func (m *RetrieveCartItems) String() string { return proto.CompactTextString(m) }
func (*RetrieveCartItems) ProtoMessage()    {}
func (*RetrieveCartItems) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8d01722d54fb2c6, []int{1}
}

func (m *RetrieveCartItems) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrieveCartItems.Unmarshal(m, b)
}
func (m *RetrieveCartItems) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrieveCartItems.Marshal(b, m, deterministic)
}
func (m *RetrieveCartItems) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveCartItems.Merge(m, src)
}
func (m *RetrieveCartItems) XXX_Size() int {
	return xxx_messageInfo_RetrieveCartItems.Size(m)
}
func (m *RetrieveCartItems) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveCartItems.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveCartItems proto.InternalMessageInfo

func (m *RetrieveCartItems) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RetrieveCartItems) GetKudakiToken() string {
	if m != nil {
		return m.KudakiToken
	}
	return ""
}

func (m *RetrieveCartItems) GetCartUuid() string {
	if m != nil {
		return m.CartUuid
	}
	return ""
}

func (m *RetrieveCartItems) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *RetrieveCartItems) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *RetrieveCartItems) GetResultSchema() []byte {
	if m != nil {
		return m.ResultSchema
	}
	return nil
}

type DeleteCartItem struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	KudakiToken          string   `protobuf:"bytes,2,opt,name=kudaki_token,json=kudakiToken,proto3" json:"kudaki_token,omitempty"`
	CartItemUuid         string   `protobuf:"bytes,3,opt,name=cart_item_uuid,json=cartItemUuid,proto3" json:"cart_item_uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCartItem) Reset()         { *m = DeleteCartItem{} }
func (m *DeleteCartItem) String() string { return proto.CompactTextString(m) }
func (*DeleteCartItem) ProtoMessage()    {}
func (*DeleteCartItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8d01722d54fb2c6, []int{2}
}

func (m *DeleteCartItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCartItem.Unmarshal(m, b)
}
func (m *DeleteCartItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCartItem.Marshal(b, m, deterministic)
}
func (m *DeleteCartItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCartItem.Merge(m, src)
}
func (m *DeleteCartItem) XXX_Size() int {
	return xxx_messageInfo_DeleteCartItem.Size(m)
}
func (m *DeleteCartItem) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCartItem.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCartItem proto.InternalMessageInfo

func (m *DeleteCartItem) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *DeleteCartItem) GetKudakiToken() string {
	if m != nil {
		return m.KudakiToken
	}
	return ""
}

func (m *DeleteCartItem) GetCartItemUuid() string {
	if m != nil {
		return m.CartItemUuid
	}
	return ""
}

type UpdateCartItem struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	KudakiToken          string   `protobuf:"bytes,2,opt,name=kudaki_token,json=kudakiToken,proto3" json:"kudaki_token,omitempty"`
	TotalItem            int32    `protobuf:"varint,3,opt,name=total_item,json=totalItem,proto3" json:"total_item,omitempty"`
	CartItemUuid         string   `protobuf:"bytes,4,opt,name=cart_item_uuid,json=cartItemUuid,proto3" json:"cart_item_uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateCartItem) Reset()         { *m = UpdateCartItem{} }
func (m *UpdateCartItem) String() string { return proto.CompactTextString(m) }
func (*UpdateCartItem) ProtoMessage()    {}
func (*UpdateCartItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8d01722d54fb2c6, []int{3}
}

func (m *UpdateCartItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCartItem.Unmarshal(m, b)
}
func (m *UpdateCartItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCartItem.Marshal(b, m, deterministic)
}
func (m *UpdateCartItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCartItem.Merge(m, src)
}
func (m *UpdateCartItem) XXX_Size() int {
	return xxx_messageInfo_UpdateCartItem.Size(m)
}
func (m *UpdateCartItem) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCartItem.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCartItem proto.InternalMessageInfo

func (m *UpdateCartItem) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *UpdateCartItem) GetKudakiToken() string {
	if m != nil {
		return m.KudakiToken
	}
	return ""
}

func (m *UpdateCartItem) GetTotalItem() int32 {
	if m != nil {
		return m.TotalItem
	}
	return 0
}

func (m *UpdateCartItem) GetCartItemUuid() string {
	if m != nil {
		return m.CartItemUuid
	}
	return ""
}

type CartItemAdded struct {
	Uid                  string           `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	CartItem             *rental.CartItem `protobuf:"bytes,2,opt,name=cart_item,json=cartItem,proto3" json:"cart_item,omitempty"`
	EventStatus          *Status          `protobuf:"bytes,4,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	AddCartItem          *AddCartItem     `protobuf:"bytes,5,opt,name=add_cart_item,json=addCartItem,proto3" json:"add_cart_item,omitempty"`
	Requester            *user.User       `protobuf:"bytes,6,opt,name=requester,proto3" json:"requester,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CartItemAdded) Reset()         { *m = CartItemAdded{} }
func (m *CartItemAdded) String() string { return proto.CompactTextString(m) }
func (*CartItemAdded) ProtoMessage()    {}
func (*CartItemAdded) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8d01722d54fb2c6, []int{4}
}

func (m *CartItemAdded) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CartItemAdded.Unmarshal(m, b)
}
func (m *CartItemAdded) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CartItemAdded.Marshal(b, m, deterministic)
}
func (m *CartItemAdded) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CartItemAdded.Merge(m, src)
}
func (m *CartItemAdded) XXX_Size() int {
	return xxx_messageInfo_CartItemAdded.Size(m)
}
func (m *CartItemAdded) XXX_DiscardUnknown() {
	xxx_messageInfo_CartItemAdded.DiscardUnknown(m)
}

var xxx_messageInfo_CartItemAdded proto.InternalMessageInfo

func (m *CartItemAdded) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *CartItemAdded) GetCartItem() *rental.CartItem {
	if m != nil {
		return m.CartItem
	}
	return nil
}

func (m *CartItemAdded) GetEventStatus() *Status {
	if m != nil {
		return m.EventStatus
	}
	return nil
}

func (m *CartItemAdded) GetAddCartItem() *AddCartItem {
	if m != nil {
		return m.AddCartItem
	}
	return nil
}

func (m *CartItemAdded) GetRequester() *user.User {
	if m != nil {
		return m.Requester
	}
	return nil
}

type CartItemsRetrieved struct {
	Uid         string     `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Requester   *user.User `protobuf:"bytes,2,opt,name=requester,proto3" json:"requester,omitempty"`
	EventStatus *Status    `protobuf:"bytes,3,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	// aggregates.rental.Cart cart = 4;
	RetrieveCartItems    *RetrieveCartItems `protobuf:"bytes,6,opt,name=retrieve_cart_items,json=retrieveCartItems,proto3" json:"retrieve_cart_items,omitempty"`
	Result               []byte             `protobuf:"bytes,7,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CartItemsRetrieved) Reset()         { *m = CartItemsRetrieved{} }
func (m *CartItemsRetrieved) String() string { return proto.CompactTextString(m) }
func (*CartItemsRetrieved) ProtoMessage()    {}
func (*CartItemsRetrieved) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8d01722d54fb2c6, []int{5}
}

func (m *CartItemsRetrieved) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CartItemsRetrieved.Unmarshal(m, b)
}
func (m *CartItemsRetrieved) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CartItemsRetrieved.Marshal(b, m, deterministic)
}
func (m *CartItemsRetrieved) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CartItemsRetrieved.Merge(m, src)
}
func (m *CartItemsRetrieved) XXX_Size() int {
	return xxx_messageInfo_CartItemsRetrieved.Size(m)
}
func (m *CartItemsRetrieved) XXX_DiscardUnknown() {
	xxx_messageInfo_CartItemsRetrieved.DiscardUnknown(m)
}

var xxx_messageInfo_CartItemsRetrieved proto.InternalMessageInfo

func (m *CartItemsRetrieved) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *CartItemsRetrieved) GetRequester() *user.User {
	if m != nil {
		return m.Requester
	}
	return nil
}

func (m *CartItemsRetrieved) GetEventStatus() *Status {
	if m != nil {
		return m.EventStatus
	}
	return nil
}

func (m *CartItemsRetrieved) GetRetrieveCartItems() *RetrieveCartItems {
	if m != nil {
		return m.RetrieveCartItems
	}
	return nil
}

func (m *CartItemsRetrieved) GetResult() []byte {
	if m != nil {
		return m.Result
	}
	return nil
}

type CartItemDeleted struct {
	Uid                  string           `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	User                 *user.User       `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	EventStatus          *Status          `protobuf:"bytes,3,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	DeleteCartItem       *DeleteCartItem  `protobuf:"bytes,4,opt,name=delete_cart_item,json=deleteCartItem,proto3" json:"delete_cart_item,omitempty"`
	CartItem             *rental.CartItem `protobuf:"bytes,5,opt,name=cart_item,json=cartItem,proto3" json:"cart_item,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CartItemDeleted) Reset()         { *m = CartItemDeleted{} }
func (m *CartItemDeleted) String() string { return proto.CompactTextString(m) }
func (*CartItemDeleted) ProtoMessage()    {}
func (*CartItemDeleted) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8d01722d54fb2c6, []int{6}
}

func (m *CartItemDeleted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CartItemDeleted.Unmarshal(m, b)
}
func (m *CartItemDeleted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CartItemDeleted.Marshal(b, m, deterministic)
}
func (m *CartItemDeleted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CartItemDeleted.Merge(m, src)
}
func (m *CartItemDeleted) XXX_Size() int {
	return xxx_messageInfo_CartItemDeleted.Size(m)
}
func (m *CartItemDeleted) XXX_DiscardUnknown() {
	xxx_messageInfo_CartItemDeleted.DiscardUnknown(m)
}

var xxx_messageInfo_CartItemDeleted proto.InternalMessageInfo

func (m *CartItemDeleted) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *CartItemDeleted) GetUser() *user.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *CartItemDeleted) GetEventStatus() *Status {
	if m != nil {
		return m.EventStatus
	}
	return nil
}

func (m *CartItemDeleted) GetDeleteCartItem() *DeleteCartItem {
	if m != nil {
		return m.DeleteCartItem
	}
	return nil
}

func (m *CartItemDeleted) GetCartItem() *rental.CartItem {
	if m != nil {
		return m.CartItem
	}
	return nil
}

type CartItemsUpdated struct {
	Uid                  string             `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	User                 *user.User         `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	InitialCartItem      []*rental.CartItem `protobuf:"bytes,3,rep,name=initial_cart_item,json=initialCartItem,proto3" json:"initial_cart_item,omitempty"`
	UpdateCartItem       *UpdateCartItem    `protobuf:"bytes,4,opt,name=update_cart_item,json=updateCartItem,proto3" json:"update_cart_item,omitempty"`
	EventStatus          *Status            `protobuf:"bytes,5,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	UpdatedCartItem      []*rental.CartItem `protobuf:"bytes,6,rep,name=updated_cart_item,json=updatedCartItem,proto3" json:"updated_cart_item,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CartItemsUpdated) Reset()         { *m = CartItemsUpdated{} }
func (m *CartItemsUpdated) String() string { return proto.CompactTextString(m) }
func (*CartItemsUpdated) ProtoMessage()    {}
func (*CartItemsUpdated) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8d01722d54fb2c6, []int{7}
}

func (m *CartItemsUpdated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CartItemsUpdated.Unmarshal(m, b)
}
func (m *CartItemsUpdated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CartItemsUpdated.Marshal(b, m, deterministic)
}
func (m *CartItemsUpdated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CartItemsUpdated.Merge(m, src)
}
func (m *CartItemsUpdated) XXX_Size() int {
	return xxx_messageInfo_CartItemsUpdated.Size(m)
}
func (m *CartItemsUpdated) XXX_DiscardUnknown() {
	xxx_messageInfo_CartItemsUpdated.DiscardUnknown(m)
}

var xxx_messageInfo_CartItemsUpdated proto.InternalMessageInfo

func (m *CartItemsUpdated) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *CartItemsUpdated) GetUser() *user.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *CartItemsUpdated) GetInitialCartItem() []*rental.CartItem {
	if m != nil {
		return m.InitialCartItem
	}
	return nil
}

func (m *CartItemsUpdated) GetUpdateCartItem() *UpdateCartItem {
	if m != nil {
		return m.UpdateCartItem
	}
	return nil
}

func (m *CartItemsUpdated) GetEventStatus() *Status {
	if m != nil {
		return m.EventStatus
	}
	return nil
}

func (m *CartItemsUpdated) GetUpdatedCartItem() []*rental.CartItem {
	if m != nil {
		return m.UpdatedCartItem
	}
	return nil
}

func init() {
	proto.RegisterEnum("event.RentalServiceCommandTopic", RentalServiceCommandTopic_name, RentalServiceCommandTopic_value)
	proto.RegisterEnum("event.RentalServiceEventTopic", RentalServiceEventTopic_name, RentalServiceEventTopic_value)
	proto.RegisterType((*AddCartItem)(nil), "event.AddCartItem")
	proto.RegisterType((*RetrieveCartItems)(nil), "event.RetrieveCartItems")
	proto.RegisterType((*DeleteCartItem)(nil), "event.DeleteCartItem")
	proto.RegisterType((*UpdateCartItem)(nil), "event.UpdateCartItem")
	proto.RegisterType((*CartItemAdded)(nil), "event.CartItemAdded")
	proto.RegisterType((*CartItemsRetrieved)(nil), "event.CartItemsRetrieved")
	proto.RegisterType((*CartItemDeleted)(nil), "event.CartItemDeleted")
	proto.RegisterType((*CartItemsUpdated)(nil), "event.CartItemsUpdated")
}

func init() { proto.RegisterFile("events/rental.proto", fileDescriptor_a8d01722d54fb2c6) }

var fileDescriptor_a8d01722d54fb2c6 = []byte{
	// 820 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x5e, 0xc7, 0x75, 0xd8, 0x1c, 0x27, 0xa9, 0x33, 0xed, 0x76, 0x43, 0x2b, 0x44, 0x36, 0x70,
	0x51, 0x56, 0xda, 0x04, 0x75, 0x25, 0x84, 0x84, 0xd0, 0x2a, 0xd4, 0x66, 0xa9, 0x04, 0x12, 0x9a,
	0x24, 0x5c, 0x70, 0x63, 0xcd, 0x66, 0xa6, 0x89, 0x95, 0x38, 0x2e, 0xe3, 0x71, 0x57, 0xbc, 0x02,
	0x4f, 0xc4, 0x2d, 0xef, 0xc0, 0x7b, 0x70, 0x0f, 0x37, 0x68, 0xce, 0xd8, 0x89, 0xbd, 0x09, 0xdb,
	0xb2, 0x7b, 0x53, 0xf5, 0x7c, 0x73, 0x7e, 0xbe, 0x6f, 0xce, 0x39, 0xe3, 0xc0, 0x91, 0xb8, 0x15,
	0x6b, 0x95, 0x0e, 0xa5, 0x58, 0x2b, 0xb6, 0x1a, 0xdc, 0xc8, 0x44, 0x25, 0xc4, 0x41, 0xf0, 0xb4,
	0x38, 0x4b, 0x15, 0x53, 0x59, 0x6a, 0xce, 0x4e, 0x9f, 0xb0, 0xf9, 0x5c, 0x8a, 0x39, 0x53, 0xa2,
	0x08, 0x1a, 0xce, 0x98, 0x54, 0x61, 0xa4, 0x44, 0x9c, 0xbb, 0x9c, 0x96, 0x5c, 0xb2, 0x54, 0x48,
	0xfc, 0x63, 0xce, 0xfa, 0x7f, 0x58, 0xe0, 0x8e, 0x38, 0xbf, 0x64, 0x52, 0x5d, 0x29, 0x11, 0x13,
	0x0f, 0xec, 0x2c, 0xe2, 0x5d, 0xab, 0x67, 0x9d, 0x37, 0xa8, 0xfe, 0x97, 0x9c, 0x41, 0x43, 0xe7,
	0x0a, 0x33, 0x8d, 0xdb, 0x88, 0x3f, 0xd4, 0xc0, 0x54, 0x1f, 0x7e, 0x0c, 0x2e, 0x1e, 0xb2, 0x38,
	0xc9, 0xd6, 0xaa, 0x7b, 0xd0, 0xb3, 0xce, 0x1d, 0x0a, 0x1a, 0x1a, 0x21, 0x42, 0x9e, 0x40, 0x73,
	0x99, 0x71, 0xb6, 0x8c, 0x42, 0x95, 0x2c, 0xc5, 0xba, 0xeb, 0x60, 0x02, 0xd7, 0x60, 0x13, 0x0d,
	0x91, 0x4f, 0xa0, 0xc5, 0x33, 0xc9, 0x54, 0x94, 0xac, 0xc3, 0x6b, 0x99, 0xc4, 0xdd, 0x7a, 0xcf,
	0x3a, 0xb7, 0x69, 0xb3, 0x00, 0xbf, 0x95, 0x49, 0x4c, 0x4e, 0xe1, 0x61, 0x61, 0x77, 0x3f, 0xc0,
	0x2a, 0x1b, 0xbb, 0xff, 0xbb, 0x05, 0x1d, 0x2a, 0x94, 0x8c, 0xc4, 0xad, 0x28, 0x84, 0xa4, 0x7b,
	0x94, 0xbc, 0xc9, 0xa5, 0xb6, 0xcb, 0xe5, 0x0c, 0x1a, 0x78, 0x7b, 0x65, 0xb1, 0x1a, 0x40, 0xb1,
	0x27, 0x50, 0x4f, 0xae, 0xaf, 0x53, 0x51, 0xe8, 0xcc, 0x2d, 0x72, 0x0c, 0xce, 0x2a, 0x8a, 0x23,
	0x85, 0xe2, 0x1c, 0x6a, 0x0c, 0x2d, 0x4b, 0x8a, 0x34, 0x5b, 0xa9, 0x30, 0x9d, 0x2d, 0x44, 0xcc,
	0x50, 0x56, 0x93, 0x36, 0x0d, 0x38, 0x46, 0xac, 0xbf, 0x84, 0xb6, 0x2f, 0x56, 0x42, 0x89, 0xb7,
	0x34, 0xe0, 0x1e, 0xb4, 0x3f, 0x85, 0xf6, 0xa6, 0xe9, 0x65, 0xee, 0xcd, 0x59, 0x9e, 0x56, 0xf3,
	0xef, 0xff, 0x66, 0x41, 0x7b, 0x7a, 0xc3, 0xd9, 0xfb, 0x56, 0xfb, 0x08, 0x40, 0x25, 0x8a, 0xad,
	0xb0, 0x1c, 0x56, 0x72, 0x68, 0x03, 0x11, 0xcc, 0xb9, 0x4b, 0xe6, 0x60, 0x0f, 0x99, 0xbf, 0x2d,
	0x68, 0x15, 0x34, 0x46, 0x9c, 0x0b, 0xbe, 0x87, 0xcb, 0x97, 0x79, 0x37, 0xb0, 0x8e, 0x26, 0xe2,
	0x5e, 0x9c, 0x0d, 0xb6, 0xc3, 0x3c, 0xc8, 0x97, 0xa4, 0x48, 0x63, 0x5a, 0x85, 0x1c, 0x3e, 0x87,
	0x26, 0x2e, 0x4b, 0x68, 0x76, 0x05, 0x19, 0xb8, 0x17, 0xad, 0x01, 0x82, 0x83, 0x31, 0x82, 0xd4,
	0x45, 0xcb, 0x18, 0xe4, 0x0b, 0x68, 0x31, 0xce, 0xc3, 0x6d, 0x3d, 0x07, 0x43, 0x48, 0x1e, 0x52,
	0xda, 0x11, 0xea, 0xb2, 0xd2, 0xc2, 0x3c, 0x87, 0x86, 0x14, 0xbf, 0x64, 0x22, 0x55, 0x42, 0x62,
	0x8b, 0xdd, 0x8b, 0x47, 0x65, 0x8e, 0xb8, 0x6b, 0xd3, 0x54, 0x48, 0xba, 0xf5, 0xeb, 0xff, 0x65,
	0x01, 0xd9, 0x4c, 0x6a, 0x31, 0xba, 0xfb, 0x6e, 0xa0, 0x92, 0xbd, 0x76, 0xbf, 0xec, 0x3b, 0xe2,
	0xed, 0x3b, 0xc5, 0x7f, 0x07, 0x47, 0x32, 0x67, 0xb1, 0xbd, 0x81, 0x34, 0x97, 0xd3, 0xcd, 0x03,
	0x77, 0x56, 0x8c, 0x76, 0xe4, 0xce, 0xd6, 0x9d, 0x40, 0xdd, 0x0c, 0x38, 0x6e, 0x69, 0x93, 0xe6,
	0x56, 0xff, 0x1f, 0x0b, 0x0e, 0x0b, 0x2f, 0x33, 0xf1, 0xfb, 0xe4, 0x7e, 0x06, 0x07, 0x5a, 0xd1,
	0xdb, 0x95, 0xa2, 0xcb, 0x3b, 0x88, 0x7c, 0x01, 0x1e, 0xc7, 0xca, 0xa5, 0x26, 0x1f, 0xe4, 0x85,
	0x4c, 0x54, 0x75, 0x15, 0x69, 0x9b, 0x57, 0x57, 0xb3, 0x32, 0x8e, 0xce, 0xff, 0x18, 0xc7, 0xfe,
	0x9f, 0x35, 0xf0, 0x36, 0x77, 0x64, 0x56, 0xf0, 0x3d, 0xe5, 0xbf, 0x84, 0x4e, 0xb4, 0x8e, 0x54,
	0xc4, 0x56, 0x25, 0x35, 0x76, 0xcf, 0xbe, 0x8b, 0xd3, 0x61, 0x1e, 0xb5, 0x11, 0xf5, 0x02, 0xbc,
	0x0c, 0x09, 0xfd, 0xe7, 0xad, 0x54, 0x9f, 0x0c, 0xda, 0xce, 0xaa, 0x4f, 0xc8, 0x9b, 0x8d, 0x70,
	0xee, 0x6c, 0xc4, 0x4b, 0xe8, 0x98, 0x1c, 0xe5, 0x75, 0xab, 0xdf, 0x83, 0x7b, 0x1e, 0x55, 0x00,
	0x4f, 0x53, 0xf8, 0x90, 0xa2, 0xcf, 0x58, 0xc8, 0xdb, 0x68, 0x26, 0x2e, 0x93, 0x38, 0x66, 0x6b,
	0x3e, 0x49, 0x6e, 0xa2, 0x19, 0xe9, 0x40, 0x6b, 0xe4, 0xfb, 0xe1, 0xe5, 0x88, 0x4e, 0xc2, 0xab,
	0x49, 0xf0, 0x83, 0xf7, 0x80, 0x3c, 0x86, 0x23, 0x1a, 0x4c, 0xe8, 0x55, 0xf0, 0x53, 0xb0, 0xc5,
	0xc7, 0x9e, 0x45, 0x8e, 0xc1, 0xf3, 0x83, 0xef, 0x83, 0x49, 0x09, 0xf6, 0x6a, 0x1a, 0x9d, 0xfe,
	0xe8, 0x8f, 0x2a, 0xa8, 0xfd, 0xf4, 0x35, 0x3c, 0xae, 0x14, 0x0d, 0xb4, 0x32, 0x53, 0xf2, 0x08,
	0x0e, 0x37, 0x9e, 0xe1, 0xc8, 0xf7, 0x03, 0xdf, 0x7b, 0x40, 0xba, 0x70, 0xbc, 0xad, 0x15, 0x16,
	0xf5, 0x7d, 0xcf, 0x22, 0x8f, 0xa0, 0xb3, 0x75, 0x37, 0xf5, 0x7d, 0xaf, 0x46, 0x4e, 0x80, 0x94,
	0x02, 0x0c, 0x03, 0xdf, 0xb3, 0xbf, 0xf9, 0xfa, 0xe7, 0xaf, 0xe6, 0x91, 0x5a, 0x64, 0xaf, 0x06,
	0xb3, 0x24, 0x1e, 0x46, 0xab, 0x05, 0x8b, 0xe3, 0x05, 0xe7, 0x43, 0xf3, 0x34, 0x3f, 0xd3, 0x77,
	0xf6, 0x9a, 0xfd, 0xfa, 0x2c, 0x35, 0xb4, 0xf4, 0x47, 0x7e, 0xc6, 0x52, 0x91, 0x0e, 0xcd, 0xaf,
	0x86, 0x57, 0x75, 0xfc, 0xe0, 0x3f, 0xff, 0x37, 0x00, 0x00, 0xff, 0xff, 0x3c, 0xee, 0xc5, 0x1d,
	0x62, 0x08, 0x00, 0x00,
}
