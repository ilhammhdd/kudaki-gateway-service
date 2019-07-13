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
	RentalServiceCommandTopic_ADD_CART_ITEM             RentalServiceCommandTopic = 0
	RentalServiceCommandTopic_RETRIEVE_CART_ITEMS       RentalServiceCommandTopic = 1
	RentalServiceCommandTopic_DELETE_CART_ITEM          RentalServiceCommandTopic = 2
	RentalServiceCommandTopic_UPDATE_CART_ITEM          RentalServiceCommandTopic = 3
	RentalServiceCommandTopic_TENANT_CONFIRM_RETURNMENT RentalServiceCommandTopic = 4
	RentalServiceCommandTopic_OWNER_CONFIRM_RETURNMENT  RentalServiceCommandTopic = 5
)

var RentalServiceCommandTopic_name = map[int32]string{
	0: "ADD_CART_ITEM",
	1: "RETRIEVE_CART_ITEMS",
	2: "DELETE_CART_ITEM",
	3: "UPDATE_CART_ITEM",
	4: "TENANT_CONFIRM_RETURNMENT",
	5: "OWNER_CONFIRM_RETURNMENT",
}

var RentalServiceCommandTopic_value = map[string]int32{
	"ADD_CART_ITEM":             0,
	"RETRIEVE_CART_ITEMS":       1,
	"DELETE_CART_ITEM":          2,
	"UPDATE_CART_ITEM":          3,
	"TENANT_CONFIRM_RETURNMENT": 4,
	"OWNER_CONFIRM_RETURNMENT":  5,
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
	RentalServiceEventTopic_CART_ITEM_ADDED             RentalServiceEventTopic = 0
	RentalServiceEventTopic_CART_ITEMS_RETRIEVED        RentalServiceEventTopic = 1
	RentalServiceEventTopic_CART_ITEM_DELETED           RentalServiceEventTopic = 2
	RentalServiceEventTopic_CART_ITEMS_UPDATED          RentalServiceEventTopic = 3
	RentalServiceEventTopic_TENANT_CONFIRMED_RETURNMENT RentalServiceEventTopic = 4
	RentalServiceEventTopic_OWNER_CONFIRMED_RETURNMENT  RentalServiceEventTopic = 5
)

var RentalServiceEventTopic_name = map[int32]string{
	0: "CART_ITEM_ADDED",
	1: "CART_ITEMS_RETRIEVED",
	2: "CART_ITEM_DELETED",
	3: "CART_ITEMS_UPDATED",
	4: "TENANT_CONFIRMED_RETURNMENT",
	5: "OWNER_CONFIRMED_RETURNMENT",
}

var RentalServiceEventTopic_value = map[string]int32{
	"CART_ITEM_ADDED":             0,
	"CART_ITEMS_RETRIEVED":        1,
	"CART_ITEM_DELETED":           2,
	"CART_ITEMS_UPDATED":          3,
	"TENANT_CONFIRMED_RETURNMENT": 4,
	"OWNER_CONFIRMED_RETURNMENT":  5,
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

type TenantConfirmReturnment struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	KudakiToken          string   `protobuf:"bytes,2,opt,name=kudaki_token,json=kudakiToken,proto3" json:"kudaki_token,omitempty"`
	OrderUuid            string   `protobuf:"bytes,3,opt,name=order_uuid,json=orderUuid,proto3" json:"order_uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TenantConfirmReturnment) Reset()         { *m = TenantConfirmReturnment{} }
func (m *TenantConfirmReturnment) String() string { return proto.CompactTextString(m) }
func (*TenantConfirmReturnment) ProtoMessage()    {}
func (*TenantConfirmReturnment) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8d01722d54fb2c6, []int{4}
}

func (m *TenantConfirmReturnment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TenantConfirmReturnment.Unmarshal(m, b)
}
func (m *TenantConfirmReturnment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TenantConfirmReturnment.Marshal(b, m, deterministic)
}
func (m *TenantConfirmReturnment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TenantConfirmReturnment.Merge(m, src)
}
func (m *TenantConfirmReturnment) XXX_Size() int {
	return xxx_messageInfo_TenantConfirmReturnment.Size(m)
}
func (m *TenantConfirmReturnment) XXX_DiscardUnknown() {
	xxx_messageInfo_TenantConfirmReturnment.DiscardUnknown(m)
}

var xxx_messageInfo_TenantConfirmReturnment proto.InternalMessageInfo

func (m *TenantConfirmReturnment) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *TenantConfirmReturnment) GetKudakiToken() string {
	if m != nil {
		return m.KudakiToken
	}
	return ""
}

func (m *TenantConfirmReturnment) GetOrderUuid() string {
	if m != nil {
		return m.OrderUuid
	}
	return ""
}

type OwnerConfirmReturnment struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	KudakiToken          string   `protobuf:"bytes,2,opt,name=kudaki_token,json=kudakiToken,proto3" json:"kudaki_token,omitempty"`
	OrderUuid            string   `protobuf:"bytes,3,opt,name=order_uuid,json=orderUuid,proto3" json:"order_uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OwnerConfirmReturnment) Reset()         { *m = OwnerConfirmReturnment{} }
func (m *OwnerConfirmReturnment) String() string { return proto.CompactTextString(m) }
func (*OwnerConfirmReturnment) ProtoMessage()    {}
func (*OwnerConfirmReturnment) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8d01722d54fb2c6, []int{5}
}

func (m *OwnerConfirmReturnment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OwnerConfirmReturnment.Unmarshal(m, b)
}
func (m *OwnerConfirmReturnment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OwnerConfirmReturnment.Marshal(b, m, deterministic)
}
func (m *OwnerConfirmReturnment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OwnerConfirmReturnment.Merge(m, src)
}
func (m *OwnerConfirmReturnment) XXX_Size() int {
	return xxx_messageInfo_OwnerConfirmReturnment.Size(m)
}
func (m *OwnerConfirmReturnment) XXX_DiscardUnknown() {
	xxx_messageInfo_OwnerConfirmReturnment.DiscardUnknown(m)
}

var xxx_messageInfo_OwnerConfirmReturnment proto.InternalMessageInfo

func (m *OwnerConfirmReturnment) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *OwnerConfirmReturnment) GetKudakiToken() string {
	if m != nil {
		return m.KudakiToken
	}
	return ""
}

func (m *OwnerConfirmReturnment) GetOrderUuid() string {
	if m != nil {
		return m.OrderUuid
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
	return fileDescriptor_a8d01722d54fb2c6, []int{6}
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
	return fileDescriptor_a8d01722d54fb2c6, []int{7}
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
	return fileDescriptor_a8d01722d54fb2c6, []int{8}
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
	return fileDescriptor_a8d01722d54fb2c6, []int{9}
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

type TenantConfirmedReturnment struct {
	Uid                     string                         `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Tenant                  *user.User                     `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	EventStatus             *Status                        `protobuf:"bytes,3,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	ReturnmentConfirmation  *rental.ReturnmentConfirmation `protobuf:"bytes,4,opt,name=returnment_confirmation,json=returnmentConfirmation,proto3" json:"returnment_confirmation,omitempty"`
	TenantConfirmReturnment *TenantConfirmReturnment       `protobuf:"bytes,5,opt,name=tenant_confirm_returnment,json=tenantConfirmReturnment,proto3" json:"tenant_confirm_returnment,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                       `json:"-"`
	XXX_unrecognized        []byte                         `json:"-"`
	XXX_sizecache           int32                          `json:"-"`
}

func (m *TenantConfirmedReturnment) Reset()         { *m = TenantConfirmedReturnment{} }
func (m *TenantConfirmedReturnment) String() string { return proto.CompactTextString(m) }
func (*TenantConfirmedReturnment) ProtoMessage()    {}
func (*TenantConfirmedReturnment) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8d01722d54fb2c6, []int{10}
}

func (m *TenantConfirmedReturnment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TenantConfirmedReturnment.Unmarshal(m, b)
}
func (m *TenantConfirmedReturnment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TenantConfirmedReturnment.Marshal(b, m, deterministic)
}
func (m *TenantConfirmedReturnment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TenantConfirmedReturnment.Merge(m, src)
}
func (m *TenantConfirmedReturnment) XXX_Size() int {
	return xxx_messageInfo_TenantConfirmedReturnment.Size(m)
}
func (m *TenantConfirmedReturnment) XXX_DiscardUnknown() {
	xxx_messageInfo_TenantConfirmedReturnment.DiscardUnknown(m)
}

var xxx_messageInfo_TenantConfirmedReturnment proto.InternalMessageInfo

func (m *TenantConfirmedReturnment) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *TenantConfirmedReturnment) GetTenant() *user.User {
	if m != nil {
		return m.Tenant
	}
	return nil
}

func (m *TenantConfirmedReturnment) GetEventStatus() *Status {
	if m != nil {
		return m.EventStatus
	}
	return nil
}

func (m *TenantConfirmedReturnment) GetReturnmentConfirmation() *rental.ReturnmentConfirmation {
	if m != nil {
		return m.ReturnmentConfirmation
	}
	return nil
}

func (m *TenantConfirmedReturnment) GetTenantConfirmReturnment() *TenantConfirmReturnment {
	if m != nil {
		return m.TenantConfirmReturnment
	}
	return nil
}

type OwnerConfirmedReturnment struct {
	Uid                    string                         `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Owner                  *user.User                     `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	EventStatus            *Status                        `protobuf:"bytes,3,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	ReturnmentConfirmation *rental.ReturnmentConfirmation `protobuf:"bytes,4,opt,name=returnment_confirmation,json=returnmentConfirmation,proto3" json:"returnment_confirmation,omitempty"`
	OwnerConfirmReturnment *OwnerConfirmReturnment        `protobuf:"bytes,5,opt,name=owner_confirm_returnment,json=ownerConfirmReturnment,proto3" json:"owner_confirm_returnment,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                       `json:"-"`
	XXX_unrecognized       []byte                         `json:"-"`
	XXX_sizecache          int32                          `json:"-"`
}

func (m *OwnerConfirmedReturnment) Reset()         { *m = OwnerConfirmedReturnment{} }
func (m *OwnerConfirmedReturnment) String() string { return proto.CompactTextString(m) }
func (*OwnerConfirmedReturnment) ProtoMessage()    {}
func (*OwnerConfirmedReturnment) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8d01722d54fb2c6, []int{11}
}

func (m *OwnerConfirmedReturnment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OwnerConfirmedReturnment.Unmarshal(m, b)
}
func (m *OwnerConfirmedReturnment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OwnerConfirmedReturnment.Marshal(b, m, deterministic)
}
func (m *OwnerConfirmedReturnment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OwnerConfirmedReturnment.Merge(m, src)
}
func (m *OwnerConfirmedReturnment) XXX_Size() int {
	return xxx_messageInfo_OwnerConfirmedReturnment.Size(m)
}
func (m *OwnerConfirmedReturnment) XXX_DiscardUnknown() {
	xxx_messageInfo_OwnerConfirmedReturnment.DiscardUnknown(m)
}

var xxx_messageInfo_OwnerConfirmedReturnment proto.InternalMessageInfo

func (m *OwnerConfirmedReturnment) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *OwnerConfirmedReturnment) GetOwner() *user.User {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *OwnerConfirmedReturnment) GetEventStatus() *Status {
	if m != nil {
		return m.EventStatus
	}
	return nil
}

func (m *OwnerConfirmedReturnment) GetReturnmentConfirmation() *rental.ReturnmentConfirmation {
	if m != nil {
		return m.ReturnmentConfirmation
	}
	return nil
}

func (m *OwnerConfirmedReturnment) GetOwnerConfirmReturnment() *OwnerConfirmReturnment {
	if m != nil {
		return m.OwnerConfirmReturnment
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
	proto.RegisterType((*TenantConfirmReturnment)(nil), "event.TenantConfirmReturnment")
	proto.RegisterType((*OwnerConfirmReturnment)(nil), "event.OwnerConfirmReturnment")
	proto.RegisterType((*CartItemAdded)(nil), "event.CartItemAdded")
	proto.RegisterType((*CartItemsRetrieved)(nil), "event.CartItemsRetrieved")
	proto.RegisterType((*CartItemDeleted)(nil), "event.CartItemDeleted")
	proto.RegisterType((*CartItemsUpdated)(nil), "event.CartItemsUpdated")
	proto.RegisterType((*TenantConfirmedReturnment)(nil), "event.TenantConfirmedReturnment")
	proto.RegisterType((*OwnerConfirmedReturnment)(nil), "event.OwnerConfirmedReturnment")
}

func init() { proto.RegisterFile("events/rental.proto", fileDescriptor_a8d01722d54fb2c6) }

var fileDescriptor_a8d01722d54fb2c6 = []byte{
	// 1046 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x57, 0x4d, 0x6f, 0xdb, 0x46,
	0x10, 0x0d, 0x25, 0x53, 0x8d, 0x46, 0x1f, 0xa1, 0xd6, 0x8e, 0x44, 0xcb, 0x75, 0xe2, 0xa8, 0x3d,
	0x38, 0x29, 0x2c, 0x15, 0x0e, 0x50, 0x14, 0x28, 0x8a, 0x40, 0x15, 0x99, 0xd4, 0x40, 0x2d, 0x17,
	0x6b, 0xaa, 0x01, 0x72, 0x21, 0xd6, 0xe2, 0xda, 0x26, 0x2c, 0x92, 0xee, 0x72, 0xe9, 0xa0, 0x7f,
	0xa1, 0xbf, 0xa6, 0xc7, 0x5c, 0x7b, 0xeb, 0xa9, 0xa7, 0xfe, 0x8f, 0xde, 0xdb, 0x4b, 0xc1, 0x5d,
	0x52, 0x22, 0x23, 0x5a, 0x76, 0x1d, 0xa0, 0xe8, 0xc5, 0xf0, 0xbe, 0x9d, 0xd9, 0x79, 0x6f, 0x38,
	0xfb, 0x48, 0xc1, 0x3a, 0xbd, 0xa2, 0x3e, 0x0f, 0x07, 0x8c, 0xfa, 0x9c, 0xcc, 0xfa, 0x97, 0x2c,
	0xe0, 0x01, 0x52, 0x05, 0xd8, 0x4d, 0xf7, 0x42, 0x4e, 0x78, 0x14, 0xca, 0xbd, 0xee, 0x13, 0x72,
	0x76, 0xc6, 0xe8, 0x19, 0xe1, 0x34, 0x4d, 0x1a, 0x4c, 0x09, 0xe3, 0xb6, 0xcb, 0xa9, 0x97, 0x84,
	0x74, 0x33, 0x21, 0x51, 0x48, 0x99, 0xf8, 0x93, 0xec, 0x0d, 0x96, 0xd3, 0x19, 0xe5, 0x11, 0xf3,
	0x3d, 0xea, 0x73, 0x7b, 0x1a, 0xf8, 0xa7, 0x2e, 0xf3, 0x08, 0x77, 0x03, 0x5f, 0x26, 0xf4, 0x7e,
	0x55, 0xa0, 0x36, 0x74, 0x9c, 0x11, 0x61, 0xfc, 0x80, 0x53, 0x0f, 0x69, 0x50, 0x8e, 0x5c, 0x47,
	0x57, 0x76, 0x94, 0xdd, 0x2a, 0x8e, 0xff, 0x45, 0x5b, 0x50, 0x8d, 0x8b, 0xdb, 0x51, 0x8c, 0x97,
	0x05, 0x7e, 0x3f, 0x06, 0x26, 0xf1, 0xe6, 0x63, 0xa8, 0x89, 0x4d, 0xe2, 0x05, 0x91, 0xcf, 0xf5,
	0xb5, 0x1d, 0x65, 0x57, 0xc5, 0x10, 0x43, 0x43, 0x81, 0xa0, 0x27, 0x50, 0xbf, 0x88, 0x1c, 0x72,
	0xe1, 0xda, 0x3c, 0xb8, 0xa0, 0xbe, 0xae, 0x8a, 0x03, 0x6a, 0x12, 0xb3, 0x62, 0x08, 0x7d, 0x02,
	0x0d, 0x27, 0x62, 0x82, 0x94, 0x7d, 0xca, 0x02, 0x4f, 0xaf, 0xec, 0x28, 0xbb, 0x65, 0x5c, 0x4f,
	0xc1, 0x97, 0x2c, 0xf0, 0x50, 0x17, 0xee, 0xa7, 0x6b, 0xfd, 0x23, 0x51, 0x65, 0xbe, 0xee, 0xbd,
	0x53, 0xa0, 0x85, 0x29, 0x67, 0x2e, 0xbd, 0xa2, 0xa9, 0x90, 0xb0, 0x40, 0xc9, 0xfb, 0x5c, 0x4a,
	0xcb, 0x5c, 0xb6, 0xa0, 0x2a, 0xda, 0x9d, 0x15, 0x1b, 0x03, 0x42, 0x6c, 0x1b, 0x2a, 0xc1, 0xe9,
	0x69, 0x48, 0x53, 0x9d, 0xc9, 0x0a, 0x6d, 0x80, 0x3a, 0x73, 0x3d, 0x97, 0x0b, 0x71, 0x2a, 0x96,
	0x8b, 0x58, 0x16, 0xa3, 0x61, 0x34, 0xe3, 0x76, 0x38, 0x3d, 0xa7, 0x1e, 0x11, 0xb2, 0xea, 0xb8,
	0x2e, 0xc1, 0x63, 0x81, 0xf5, 0x2e, 0xa0, 0x69, 0xd0, 0x19, 0xe5, 0x74, 0xc5, 0x03, 0xb8, 0x05,
	0xed, 0x4f, 0xa1, 0x39, 0x9f, 0x92, 0x2c, 0xf7, 0xfa, 0x34, 0x39, 0x36, 0xe6, 0xdf, 0xfb, 0x59,
	0x81, 0xe6, 0xe4, 0xd2, 0x21, 0x1f, 0x5a, 0x6d, 0x1b, 0x80, 0x07, 0x9c, 0xcc, 0x44, 0x39, 0x51,
	0x49, 0xc5, 0x55, 0x81, 0x88, 0x33, 0x97, 0xc9, 0xac, 0x15, 0x90, 0xf1, 0xa0, 0x63, 0x51, 0x9f,
	0xf8, 0x7c, 0x24, 0x87, 0x12, 0xcf, 0xc7, 0xf4, 0xce, 0xa4, 0x02, 0xe6, 0x50, 0x96, 0x95, 0x5f,
	0x15, 0x88, 0x28, 0x37, 0x83, 0xf6, 0xd1, 0x5b, 0x9f, 0xb2, 0xff, 0xa6, 0xda, 0x5f, 0x0a, 0x34,
	0xd2, 0x1e, 0x0f, 0x1d, 0x87, 0x3a, 0x05, 0x55, 0xbe, 0x4c, 0x46, 0x4d, 0x34, 0x31, 0x2e, 0x51,
	0xdb, 0xdf, 0xea, 0x2f, 0xae, 0x6f, 0x3f, 0xb1, 0x8c, 0xf4, 0x18, 0x39, 0x87, 0xa2, 0xc1, 0x9f,
	0x43, 0x5d, 0x58, 0x87, 0x2d, 0x9d, 0x43, 0xb4, 0xb7, 0xb6, 0xdf, 0xe8, 0x0b, 0xb0, 0x7f, 0x2c,
	0x40, 0x5c, 0x13, 0x2b, 0xb9, 0x40, 0x5f, 0x40, 0x83, 0x38, 0x8e, 0xbd, 0xa8, 0xa7, 0x8a, 0x14,
	0x94, 0xa4, 0x64, 0x0c, 0x00, 0xd7, 0x48, 0xc6, 0x0d, 0x9e, 0x43, 0x95, 0xd1, 0x1f, 0x23, 0x1a,
	0x72, 0xca, 0xc4, 0xfc, 0xd6, 0xf6, 0x1f, 0x66, 0x39, 0x0a, 0xe7, 0x99, 0x84, 0x94, 0xe1, 0x45,
	0x5c, 0xef, 0x4f, 0x05, 0xd0, 0xfc, 0x1a, 0xa6, 0xf7, 0xb2, 0xa8, 0x03, 0xb9, 0xd3, 0x4b, 0xb7,
	0x3b, 0x7d, 0x49, 0x7c, 0xf9, 0x46, 0xf1, 0xdf, 0xc2, 0x3a, 0x4b, 0x58, 0x2c, 0x3a, 0x10, 0x26,
	0x72, 0xf4, 0x24, 0x71, 0xc9, 0x3f, 0x70, 0x8b, 0x2d, 0x59, 0x4a, 0x1b, 0x2a, 0xf2, 0xf6, 0x0a,
	0x0b, 0xaa, 0xe3, 0x64, 0xd5, 0xfb, 0x5b, 0x81, 0x07, 0x69, 0x94, 0xbc, 0xce, 0x45, 0x72, 0x9f,
	0xc2, 0x5a, 0xac, 0x68, 0xb5, 0x52, 0x11, 0x72, 0x07, 0x91, 0x2f, 0x40, 0x73, 0x44, 0xe5, 0xcc,
	0x43, 0x5e, 0x4b, 0x0a, 0xc9, 0xac, 0xbc, 0xcf, 0xe0, 0xa6, 0x93, 0xf7, 0x9d, 0xdc, 0x38, 0xaa,
	0xff, 0x62, 0x1c, 0x7b, 0x7f, 0x94, 0x40, 0x9b, 0xf7, 0x48, 0xfa, 0xcb, 0x07, 0xca, 0x7f, 0x05,
	0x2d, 0xd7, 0x77, 0xb9, 0x4b, 0x66, 0x19, 0x35, 0xe5, 0x9d, 0xf2, 0x4d, 0x9c, 0x1e, 0x24, 0x59,
	0x73, 0x51, 0x2f, 0x40, 0x8b, 0x04, 0xa1, 0x6b, 0xbb, 0x92, 0xf7, 0x43, 0xdc, 0x8c, 0xf2, 0xfe,
	0xf8, 0xfe, 0x83, 0x50, 0x6f, 0x7c, 0x10, 0xaf, 0xa0, 0x25, 0xcf, 0xc8, 0x5e, 0xb7, 0xca, 0x2d,
	0xb8, 0x27, 0x59, 0x29, 0xd0, 0xfb, 0xbd, 0x04, 0x9b, 0x39, 0x87, 0xa4, 0xce, 0x4a, 0xd7, 0xda,
	0x83, 0x0a, 0x17, 0xe1, 0xab, 0x3b, 0x9c, 0x04, 0xdd, 0x61, 0xc4, 0x4e, 0xa0, 0x73, 0xcd, 0xb7,
	0x44, 0xd2, 0xd3, 0xa7, 0x05, 0xfa, 0x16, 0x94, 0x47, 0x99, 0x04, 0xdc, 0x66, 0x85, 0x38, 0x7a,
	0x03, 0x9b, 0x92, 0x5f, 0x7a, 0xbe, 0xbd, 0x08, 0x4c, 0x9a, 0xff, 0x28, 0xa1, 0x78, 0xcd, 0xdb,
	0x03, 0x77, 0x78, 0xf1, 0x46, 0xef, 0xb7, 0x12, 0xe8, 0xd9, 0x77, 0xc0, 0x0d, 0xfd, 0xfc, 0x0c,
	0xd4, 0x20, 0x8e, 0x5e, 0xdd, 0x4e, 0x19, 0xf3, 0x3f, 0xed, 0xe6, 0x6b, 0xd0, 0x05, 0xbd, 0xeb,
	0x9b, 0xb9, 0x9d, 0x30, 0x2c, 0x7e, 0x37, 0xe2, 0x76, 0x50, 0x88, 0x3f, 0xfb, 0x45, 0x81, 0x4d,
	0x2c, 0x28, 0x1d, 0x53, 0x76, 0xe5, 0x4e, 0xe9, 0x28, 0xf0, 0x3c, 0xe2, 0x3b, 0x56, 0x70, 0xe9,
	0x4e, 0x51, 0x0b, 0x1a, 0x43, 0xc3, 0xb0, 0x47, 0x43, 0x6c, 0xd9, 0x07, 0x96, 0x79, 0xa8, 0xdd,
	0x43, 0x1d, 0x58, 0xc7, 0xa6, 0x85, 0x0f, 0xcc, 0x1f, 0xcc, 0x05, 0x7e, 0xac, 0x29, 0x68, 0x03,
	0x34, 0xc3, 0xfc, 0xce, 0xb4, 0x32, 0xb0, 0x56, 0x8a, 0xd1, 0xc9, 0xf7, 0xc6, 0x30, 0x87, 0x96,
	0xd1, 0x36, 0x6c, 0x5a, 0xe6, 0x78, 0x38, 0xb6, 0xec, 0xd1, 0xd1, 0xf8, 0xe5, 0x01, 0x3e, 0xb4,
	0xb1, 0x69, 0x4d, 0xf0, 0xf8, 0xd0, 0x1c, 0x5b, 0xda, 0x1a, 0xfa, 0x18, 0xf4, 0xa3, 0xd7, 0x63,
	0x13, 0x17, 0xed, 0xaa, 0xcf, 0xde, 0x29, 0xd0, 0xc9, 0x51, 0x36, 0x63, 0xe1, 0x92, 0xf0, 0x3a,
	0x3c, 0x98, 0xd7, 0xb1, 0x87, 0x86, 0x61, 0x1a, 0xda, 0x3d, 0xa4, 0xc3, 0xc6, 0x82, 0xa9, 0x9d,
	0xb2, 0x37, 0x34, 0x05, 0x3d, 0x84, 0xd6, 0x22, 0x5c, 0xb2, 0x37, 0xb4, 0x12, 0x6a, 0x03, 0xca,
	0x24, 0x48, 0xfe, 0x86, 0x56, 0x46, 0x8f, 0x61, 0x2b, 0x4f, 0xdb, 0x34, 0xf2, 0xc4, 0x1f, 0x41,
	0x37, 0x47, 0x3c, 0xbf, 0xaf, 0x7e, 0xf3, 0xf5, 0x9b, 0xaf, 0xce, 0x5c, 0x7e, 0x1e, 0x9d, 0xf4,
	0xa7, 0x81, 0x37, 0x70, 0x67, 0xe7, 0xc4, 0xf3, 0xce, 0x1d, 0x67, 0x20, 0x3f, 0x48, 0xf6, 0xe2,
	0x09, 0x79, 0x4b, 0x7e, 0xda, 0x0b, 0xa5, 0xae, 0xf8, 0xe7, 0xc0, 0x94, 0x84, 0x34, 0x1c, 0xc8,
	0xdf, 0x17, 0x27, 0x15, 0xf1, 0xa5, 0xff, 0xfc, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x85, 0x9e,
	0x50, 0x3d, 0x8c, 0x0c, 0x00, 0x00,
}
