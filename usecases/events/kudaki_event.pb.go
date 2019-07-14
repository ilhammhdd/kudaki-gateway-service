// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events/kudaki_event.proto

package events

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	kudaki_event "github.com/ilhammhdd/kudaki-gateway-service/entities/aggregates/kudaki_event"
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

type EventServiceCommandTopic int32

const (
	EventServiceCommandTopic_ADD_KUDAKI_EVENT       EventServiceCommandTopic = 0
	EventServiceCommandTopic_UPDATE_KUDAKI_EVENT    EventServiceCommandTopic = 1
	EventServiceCommandTopic_DELETE_KUDAKI_EVENT    EventServiceCommandTopic = 2
	EventServiceCommandTopic_RETRIEVE_KUDAKI_EVENT  EventServiceCommandTopic = 3
	EventServiceCommandTopic_RETRIEVE_KUDAKI_EVENTS EventServiceCommandTopic = 4
	EventServiceCommandTopic_ADD_PRICE              EventServiceCommandTopic = 5
	EventServiceCommandTopic_UPDATE_PRICE           EventServiceCommandTopic = 6
)

var EventServiceCommandTopic_name = map[int32]string{
	0: "ADD_KUDAKI_EVENT",
	1: "UPDATE_KUDAKI_EVENT",
	2: "DELETE_KUDAKI_EVENT",
	3: "RETRIEVE_KUDAKI_EVENT",
	4: "RETRIEVE_KUDAKI_EVENTS",
	5: "ADD_PRICE",
	6: "UPDATE_PRICE",
}

var EventServiceCommandTopic_value = map[string]int32{
	"ADD_KUDAKI_EVENT":       0,
	"UPDATE_KUDAKI_EVENT":    1,
	"DELETE_KUDAKI_EVENT":    2,
	"RETRIEVE_KUDAKI_EVENT":  3,
	"RETRIEVE_KUDAKI_EVENTS": 4,
	"ADD_PRICE":              5,
	"UPDATE_PRICE":           6,
}

func (x EventServiceCommandTopic) String() string {
	return proto.EnumName(EventServiceCommandTopic_name, int32(x))
}

func (EventServiceCommandTopic) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fa36ed941739bd5b, []int{0}
}

type EventServiceEventTopic int32

const (
	EventServiceEventTopic_KUDAKI_EVENT_ADDED      EventServiceEventTopic = 0
	EventServiceEventTopic_KUDAKI_EVENT_UPDATED    EventServiceEventTopic = 1
	EventServiceEventTopic_KUDAKI_EVENT_DELETED    EventServiceEventTopic = 2
	EventServiceEventTopic_KUDAKI_EVENT_RETRIEVED  EventServiceEventTopic = 3
	EventServiceEventTopic_KUDAKI_EVENTS_RETRIEVED EventServiceEventTopic = 4
	EventServiceEventTopic_PRICE_ADDED             EventServiceEventTopic = 5
	EventServiceEventTopic_PRICE_UPDATED           EventServiceEventTopic = 6
)

var EventServiceEventTopic_name = map[int32]string{
	0: "KUDAKI_EVENT_ADDED",
	1: "KUDAKI_EVENT_UPDATED",
	2: "KUDAKI_EVENT_DELETED",
	3: "KUDAKI_EVENT_RETRIEVED",
	4: "KUDAKI_EVENTS_RETRIEVED",
	5: "PRICE_ADDED",
	6: "PRICE_UPDATED",
}

var EventServiceEventTopic_value = map[string]int32{
	"KUDAKI_EVENT_ADDED":      0,
	"KUDAKI_EVENT_UPDATED":    1,
	"KUDAKI_EVENT_DELETED":    2,
	"KUDAKI_EVENT_RETRIEVED":  3,
	"KUDAKI_EVENTS_RETRIEVED": 4,
	"PRICE_ADDED":             5,
	"PRICE_UPDATED":           6,
}

func (x EventServiceEventTopic) String() string {
	return proto.EnumName(EventServiceEventTopic_name, int32(x))
}

func (EventServiceEventTopic) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fa36ed941739bd5b, []int{1}
}

type AddKudakiEvent struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Venue                string   `protobuf:"bytes,2,opt,name=venue,proto3" json:"venue,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	DurationFrom         int64    `protobuf:"varint,4,opt,name=duration_from,json=durationFrom,proto3" json:"duration_from,omitempty"`
	DurationTo           int64    `protobuf:"varint,5,opt,name=duration_to,json=durationTo,proto3" json:"duration_to,omitempty"`
	KudakiToken          string   `protobuf:"bytes,6,opt,name=kudaki_token,json=kudakiToken,proto3" json:"kudaki_token,omitempty"`
	Name                 string   `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddKudakiEvent) Reset()         { *m = AddKudakiEvent{} }
func (m *AddKudakiEvent) String() string { return proto.CompactTextString(m) }
func (*AddKudakiEvent) ProtoMessage()    {}
func (*AddKudakiEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa36ed941739bd5b, []int{0}
}

func (m *AddKudakiEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddKudakiEvent.Unmarshal(m, b)
}
func (m *AddKudakiEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddKudakiEvent.Marshal(b, m, deterministic)
}
func (m *AddKudakiEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddKudakiEvent.Merge(m, src)
}
func (m *AddKudakiEvent) XXX_Size() int {
	return xxx_messageInfo_AddKudakiEvent.Size(m)
}
func (m *AddKudakiEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_AddKudakiEvent.DiscardUnknown(m)
}

var xxx_messageInfo_AddKudakiEvent proto.InternalMessageInfo

func (m *AddKudakiEvent) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *AddKudakiEvent) GetVenue() string {
	if m != nil {
		return m.Venue
	}
	return ""
}

func (m *AddKudakiEvent) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *AddKudakiEvent) GetDurationFrom() int64 {
	if m != nil {
		return m.DurationFrom
	}
	return 0
}

func (m *AddKudakiEvent) GetDurationTo() int64 {
	if m != nil {
		return m.DurationTo
	}
	return 0
}

func (m *AddKudakiEvent) GetKudakiToken() string {
	if m != nil {
		return m.KudakiToken
	}
	return ""
}

func (m *AddKudakiEvent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DeleteKudakiEvent struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	EventUuid            string   `protobuf:"bytes,2,opt,name=event_uuid,json=eventUuid,proto3" json:"event_uuid,omitempty"`
	KudakiToken          string   `protobuf:"bytes,3,opt,name=kudaki_token,json=kudakiToken,proto3" json:"kudaki_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteKudakiEvent) Reset()         { *m = DeleteKudakiEvent{} }
func (m *DeleteKudakiEvent) String() string { return proto.CompactTextString(m) }
func (*DeleteKudakiEvent) ProtoMessage()    {}
func (*DeleteKudakiEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa36ed941739bd5b, []int{1}
}

func (m *DeleteKudakiEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteKudakiEvent.Unmarshal(m, b)
}
func (m *DeleteKudakiEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteKudakiEvent.Marshal(b, m, deterministic)
}
func (m *DeleteKudakiEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteKudakiEvent.Merge(m, src)
}
func (m *DeleteKudakiEvent) XXX_Size() int {
	return xxx_messageInfo_DeleteKudakiEvent.Size(m)
}
func (m *DeleteKudakiEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteKudakiEvent.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteKudakiEvent proto.InternalMessageInfo

func (m *DeleteKudakiEvent) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *DeleteKudakiEvent) GetEventUuid() string {
	if m != nil {
		return m.EventUuid
	}
	return ""
}

func (m *DeleteKudakiEvent) GetKudakiToken() string {
	if m != nil {
		return m.KudakiToken
	}
	return ""
}

type AddPrice struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	KudakiToken          string   `protobuf:"bytes,2,opt,name=kudaki_token,json=kudakiToken,proto3" json:"kudaki_token,omitempty"`
	Duration             int64    `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
	DurationUnit         string   `protobuf:"bytes,4,opt,name=duration_unit,json=durationUnit,proto3" json:"duration_unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddPrice) Reset()         { *m = AddPrice{} }
func (m *AddPrice) String() string { return proto.CompactTextString(m) }
func (*AddPrice) ProtoMessage()    {}
func (*AddPrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa36ed941739bd5b, []int{2}
}

func (m *AddPrice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPrice.Unmarshal(m, b)
}
func (m *AddPrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPrice.Marshal(b, m, deterministic)
}
func (m *AddPrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPrice.Merge(m, src)
}
func (m *AddPrice) XXX_Size() int {
	return xxx_messageInfo_AddPrice.Size(m)
}
func (m *AddPrice) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPrice.DiscardUnknown(m)
}

var xxx_messageInfo_AddPrice proto.InternalMessageInfo

func (m *AddPrice) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *AddPrice) GetKudakiToken() string {
	if m != nil {
		return m.KudakiToken
	}
	return ""
}

func (m *AddPrice) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *AddPrice) GetDurationUnit() string {
	if m != nil {
		return m.DurationUnit
	}
	return ""
}

type UpdatePrice struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	KudakiToken          string   `protobuf:"bytes,2,opt,name=kudaki_token,json=kudakiToken,proto3" json:"kudaki_token,omitempty"`
	Duration             int64    `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
	DurationUnit         string   `protobuf:"bytes,4,opt,name=duration_unit,json=durationUnit,proto3" json:"duration_unit,omitempty"`
	PriceUuid            string   `protobuf:"bytes,5,opt,name=price_uuid,json=priceUuid,proto3" json:"price_uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePrice) Reset()         { *m = UpdatePrice{} }
func (m *UpdatePrice) String() string { return proto.CompactTextString(m) }
func (*UpdatePrice) ProtoMessage()    {}
func (*UpdatePrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa36ed941739bd5b, []int{3}
}

func (m *UpdatePrice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePrice.Unmarshal(m, b)
}
func (m *UpdatePrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePrice.Marshal(b, m, deterministic)
}
func (m *UpdatePrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePrice.Merge(m, src)
}
func (m *UpdatePrice) XXX_Size() int {
	return xxx_messageInfo_UpdatePrice.Size(m)
}
func (m *UpdatePrice) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePrice.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePrice proto.InternalMessageInfo

func (m *UpdatePrice) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *UpdatePrice) GetKudakiToken() string {
	if m != nil {
		return m.KudakiToken
	}
	return ""
}

func (m *UpdatePrice) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *UpdatePrice) GetDurationUnit() string {
	if m != nil {
		return m.DurationUnit
	}
	return ""
}

func (m *UpdatePrice) GetPriceUuid() string {
	if m != nil {
		return m.PriceUuid
	}
	return ""
}

type KudakiEventAdded struct {
	Uid                  string                    `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Organizer            *user.User                `protobuf:"bytes,2,opt,name=organizer,proto3" json:"organizer,omitempty"`
	EventStatus          *Status                   `protobuf:"bytes,3,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	KudakiEvent          *kudaki_event.KudakiEvent `protobuf:"bytes,4,opt,name=kudaki_event,json=kudakiEvent,proto3" json:"kudaki_event,omitempty"`
	AddKudakiEvent       *AddKudakiEvent           `protobuf:"bytes,5,opt,name=add_kudaki_event,json=addKudakiEvent,proto3" json:"add_kudaki_event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *KudakiEventAdded) Reset()         { *m = KudakiEventAdded{} }
func (m *KudakiEventAdded) String() string { return proto.CompactTextString(m) }
func (*KudakiEventAdded) ProtoMessage()    {}
func (*KudakiEventAdded) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa36ed941739bd5b, []int{4}
}

func (m *KudakiEventAdded) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KudakiEventAdded.Unmarshal(m, b)
}
func (m *KudakiEventAdded) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KudakiEventAdded.Marshal(b, m, deterministic)
}
func (m *KudakiEventAdded) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KudakiEventAdded.Merge(m, src)
}
func (m *KudakiEventAdded) XXX_Size() int {
	return xxx_messageInfo_KudakiEventAdded.Size(m)
}
func (m *KudakiEventAdded) XXX_DiscardUnknown() {
	xxx_messageInfo_KudakiEventAdded.DiscardUnknown(m)
}

var xxx_messageInfo_KudakiEventAdded proto.InternalMessageInfo

func (m *KudakiEventAdded) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *KudakiEventAdded) GetOrganizer() *user.User {
	if m != nil {
		return m.Organizer
	}
	return nil
}

func (m *KudakiEventAdded) GetEventStatus() *Status {
	if m != nil {
		return m.EventStatus
	}
	return nil
}

func (m *KudakiEventAdded) GetKudakiEvent() *kudaki_event.KudakiEvent {
	if m != nil {
		return m.KudakiEvent
	}
	return nil
}

func (m *KudakiEventAdded) GetAddKudakiEvent() *AddKudakiEvent {
	if m != nil {
		return m.AddKudakiEvent
	}
	return nil
}

type KudakiEventDeleted struct {
	Uid                  string                    `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Organizer            *user.User                `protobuf:"bytes,2,opt,name=organizer,proto3" json:"organizer,omitempty"`
	EventStatus          *Status                   `protobuf:"bytes,3,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	KudakiEvent          *kudaki_event.KudakiEvent `protobuf:"bytes,4,opt,name=kudaki_event,json=kudakiEvent,proto3" json:"kudaki_event,omitempty"`
	DeleteKudakiEvent    *DeleteKudakiEvent        `protobuf:"bytes,5,opt,name=delete_kudaki_event,json=deleteKudakiEvent,proto3" json:"delete_kudaki_event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *KudakiEventDeleted) Reset()         { *m = KudakiEventDeleted{} }
func (m *KudakiEventDeleted) String() string { return proto.CompactTextString(m) }
func (*KudakiEventDeleted) ProtoMessage()    {}
func (*KudakiEventDeleted) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa36ed941739bd5b, []int{5}
}

func (m *KudakiEventDeleted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KudakiEventDeleted.Unmarshal(m, b)
}
func (m *KudakiEventDeleted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KudakiEventDeleted.Marshal(b, m, deterministic)
}
func (m *KudakiEventDeleted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KudakiEventDeleted.Merge(m, src)
}
func (m *KudakiEventDeleted) XXX_Size() int {
	return xxx_messageInfo_KudakiEventDeleted.Size(m)
}
func (m *KudakiEventDeleted) XXX_DiscardUnknown() {
	xxx_messageInfo_KudakiEventDeleted.DiscardUnknown(m)
}

var xxx_messageInfo_KudakiEventDeleted proto.InternalMessageInfo

func (m *KudakiEventDeleted) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *KudakiEventDeleted) GetOrganizer() *user.User {
	if m != nil {
		return m.Organizer
	}
	return nil
}

func (m *KudakiEventDeleted) GetEventStatus() *Status {
	if m != nil {
		return m.EventStatus
	}
	return nil
}

func (m *KudakiEventDeleted) GetKudakiEvent() *kudaki_event.KudakiEvent {
	if m != nil {
		return m.KudakiEvent
	}
	return nil
}

func (m *KudakiEventDeleted) GetDeleteKudakiEvent() *DeleteKudakiEvent {
	if m != nil {
		return m.DeleteKudakiEvent
	}
	return nil
}

type PriceAdded struct {
	Uid                  string              `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Creator              *user.User          `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	EventStatus          *Status             `protobuf:"bytes,3,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	AddPrice             *AddPrice           `protobuf:"bytes,4,opt,name=add_price,json=addPrice,proto3" json:"add_price,omitempty"`
	Price                *kudaki_event.Price `protobuf:"bytes,5,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PriceAdded) Reset()         { *m = PriceAdded{} }
func (m *PriceAdded) String() string { return proto.CompactTextString(m) }
func (*PriceAdded) ProtoMessage()    {}
func (*PriceAdded) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa36ed941739bd5b, []int{6}
}

func (m *PriceAdded) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PriceAdded.Unmarshal(m, b)
}
func (m *PriceAdded) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PriceAdded.Marshal(b, m, deterministic)
}
func (m *PriceAdded) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PriceAdded.Merge(m, src)
}
func (m *PriceAdded) XXX_Size() int {
	return xxx_messageInfo_PriceAdded.Size(m)
}
func (m *PriceAdded) XXX_DiscardUnknown() {
	xxx_messageInfo_PriceAdded.DiscardUnknown(m)
}

var xxx_messageInfo_PriceAdded proto.InternalMessageInfo

func (m *PriceAdded) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *PriceAdded) GetCreator() *user.User {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *PriceAdded) GetEventStatus() *Status {
	if m != nil {
		return m.EventStatus
	}
	return nil
}

func (m *PriceAdded) GetAddPrice() *AddPrice {
	if m != nil {
		return m.AddPrice
	}
	return nil
}

func (m *PriceAdded) GetPrice() *kudaki_event.Price {
	if m != nil {
		return m.Price
	}
	return nil
}

type PriceUpdated struct {
	Uid                  string              `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Creator              *user.User          `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	EventStatus          *Status             `protobuf:"bytes,3,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	UpdatePrice          *UpdatePrice        `protobuf:"bytes,4,opt,name=update_price,json=updatePrice,proto3" json:"update_price,omitempty"`
	Price                *kudaki_event.Price `protobuf:"bytes,5,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PriceUpdated) Reset()         { *m = PriceUpdated{} }
func (m *PriceUpdated) String() string { return proto.CompactTextString(m) }
func (*PriceUpdated) ProtoMessage()    {}
func (*PriceUpdated) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa36ed941739bd5b, []int{7}
}

func (m *PriceUpdated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PriceUpdated.Unmarshal(m, b)
}
func (m *PriceUpdated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PriceUpdated.Marshal(b, m, deterministic)
}
func (m *PriceUpdated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PriceUpdated.Merge(m, src)
}
func (m *PriceUpdated) XXX_Size() int {
	return xxx_messageInfo_PriceUpdated.Size(m)
}
func (m *PriceUpdated) XXX_DiscardUnknown() {
	xxx_messageInfo_PriceUpdated.DiscardUnknown(m)
}

var xxx_messageInfo_PriceUpdated proto.InternalMessageInfo

func (m *PriceUpdated) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *PriceUpdated) GetCreator() *user.User {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *PriceUpdated) GetEventStatus() *Status {
	if m != nil {
		return m.EventStatus
	}
	return nil
}

func (m *PriceUpdated) GetUpdatePrice() *UpdatePrice {
	if m != nil {
		return m.UpdatePrice
	}
	return nil
}

func (m *PriceUpdated) GetPrice() *kudaki_event.Price {
	if m != nil {
		return m.Price
	}
	return nil
}

func init() {
	proto.RegisterEnum("aggregates.event.EventServiceCommandTopic", EventServiceCommandTopic_name, EventServiceCommandTopic_value)
	proto.RegisterEnum("aggregates.event.EventServiceEventTopic", EventServiceEventTopic_name, EventServiceEventTopic_value)
	proto.RegisterType((*AddKudakiEvent)(nil), "aggregates.event.AddKudakiEvent")
	proto.RegisterType((*DeleteKudakiEvent)(nil), "aggregates.event.DeleteKudakiEvent")
	proto.RegisterType((*AddPrice)(nil), "aggregates.event.AddPrice")
	proto.RegisterType((*UpdatePrice)(nil), "aggregates.event.UpdatePrice")
	proto.RegisterType((*KudakiEventAdded)(nil), "aggregates.event.KudakiEventAdded")
	proto.RegisterType((*KudakiEventDeleted)(nil), "aggregates.event.KudakiEventDeleted")
	proto.RegisterType((*PriceAdded)(nil), "aggregates.event.PriceAdded")
	proto.RegisterType((*PriceUpdated)(nil), "aggregates.event.PriceUpdated")
}

func init() { proto.RegisterFile("events/kudaki_event.proto", fileDescriptor_fa36ed941739bd5b) }

var fileDescriptor_fa36ed941739bd5b = []byte{
	// 797 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x56, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0xae, 0xf3, 0xd5, 0x66, 0x9c, 0xf4, 0x75, 0xb7, 0x5f, 0x6e, 0x5e, 0xf5, 0x7d, 0x83, 0xcb,
	0xa1, 0xaa, 0xd4, 0x04, 0xb5, 0x48, 0x1c, 0x10, 0x12, 0xa1, 0x36, 0xa8, 0x14, 0xa1, 0xca, 0x49,
	0x7a, 0xe0, 0x12, 0x6d, 0xb3, 0x4b, 0x6a, 0xb5, 0xb6, 0x23, 0x7f, 0x14, 0xc1, 0x89, 0x9f, 0x82,
	0xb8, 0xf2, 0x0f, 0xb8, 0xf1, 0x2b, 0xf8, 0x1b, 0xdc, 0xb9, 0xa0, 0x9d, 0xb5, 0x5b, 0x3b, 0x49,
	0x41, 0xea, 0xa5, 0x12, 0x97, 0x68, 0xe7, 0x99, 0xc9, 0xcc, 0x33, 0xcf, 0xce, 0x58, 0x0b, 0x1b,
	0xfc, 0x92, 0x7b, 0x51, 0xd8, 0x3e, 0x8f, 0x19, 0x3d, 0x77, 0x06, 0x68, 0xb5, 0xc6, 0x81, 0x1f,
	0xf9, 0x44, 0xa3, 0xa3, 0x51, 0xc0, 0x47, 0x34, 0xe2, 0x61, 0x0b, 0xf1, 0x46, 0xe3, 0x1a, 0x69,
	0xc7, 0x21, 0x0f, 0xf0, 0x47, 0x46, 0x37, 0x76, 0x32, 0xbe, 0x6c, 0xb2, 0x19, 0x99, 0x1b, 0x5b,
	0x37, 0xc5, 0x8e, 0x03, 0x67, 0xc8, 0x93, 0xa0, 0xe5, 0x84, 0x59, 0x18, 0xd1, 0x28, 0x0e, 0x25,
	0x68, 0x7c, 0x57, 0x60, 0xb1, 0xc3, 0xd8, 0x11, 0xfe, 0xc9, 0x12, 0x01, 0x44, 0x83, 0x62, 0xec,
	0x30, 0x5d, 0x69, 0x2a, 0xdb, 0x55, 0x5b, 0x1c, 0xc9, 0x0a, 0x94, 0x2f, 0xb9, 0x17, 0x73, 0xbd,
	0x80, 0x98, 0x34, 0x48, 0x13, 0x54, 0xc6, 0xc3, 0x61, 0xe0, 0x8c, 0x23, 0xc7, 0xf7, 0xf4, 0x22,
	0xfa, 0xb2, 0x10, 0xd9, 0x82, 0x3a, 0x8b, 0x03, 0x2a, 0xce, 0x83, 0xb7, 0x81, 0xef, 0xea, 0xa5,
	0xa6, 0xb2, 0x5d, 0xb4, 0x6b, 0x29, 0xf8, 0x3c, 0xf0, 0x5d, 0xf2, 0x3f, 0xa8, 0x57, 0x41, 0x91,
	0xaf, 0x97, 0x31, 0x04, 0x52, 0xa8, 0xe7, 0x93, 0x7b, 0x50, 0x4b, 0x7a, 0x8a, 0xfc, 0x73, 0xee,
	0xe9, 0x15, 0x59, 0x48, 0x62, 0x3d, 0x01, 0x11, 0x02, 0x25, 0x8f, 0xba, 0x5c, 0x9f, 0x47, 0x17,
	0x9e, 0x8d, 0x11, 0x2c, 0x99, 0xfc, 0x82, 0x47, 0xfc, 0xf7, 0xbd, 0x6d, 0x02, 0xa0, 0x2e, 0x83,
	0x58, 0x38, 0x64, 0x83, 0x55, 0x44, 0xfa, 0xc2, 0x3d, 0x59, 0xbc, 0x38, 0x55, 0xdc, 0xf8, 0xa8,
	0xc0, 0x42, 0x87, 0xb1, 0x63, 0x21, 0xf5, 0x8c, 0x02, 0x93, 0x19, 0x0a, 0xd3, 0xf4, 0x1b, 0xb0,
	0x90, 0xf6, 0x8b, 0x05, 0x8a, 0xf6, 0x95, 0x9d, 0xd3, 0x30, 0xf6, 0x9c, 0x08, 0x35, 0xac, 0x5e,
	0x6b, 0xd8, 0xf7, 0x9c, 0xc8, 0xf8, 0xac, 0x80, 0xda, 0x1f, 0x33, 0x1a, 0xf1, 0x3b, 0x64, 0x21,
	0xa4, 0xc4, 0x79, 0x93, 0x52, 0x96, 0xa5, 0x94, 0x88, 0x08, 0x29, 0x8d, 0x4f, 0x05, 0xd0, 0x32,
	0x77, 0xd1, 0x61, 0x8c, 0xb3, 0x19, 0x4c, 0xf7, 0xa1, 0xea, 0x07, 0x23, 0xea, 0x39, 0x1f, 0x78,
	0x80, 0x34, 0xd5, 0xbd, 0xd5, 0x56, 0x66, 0x73, 0x70, 0x45, 0xfa, 0x21, 0x0f, 0xec, 0xeb, 0x38,
	0xf2, 0x00, 0x6a, 0xf2, 0x16, 0xe5, 0x70, 0x23, 0x7f, 0x75, 0xaf, 0x2e, 0xd7, 0xac, 0xd5, 0x45,
	0xd0, 0x56, 0xd1, 0x92, 0x06, 0x79, 0x71, 0x25, 0x08, 0xa2, 0xd8, 0x90, 0xba, 0x77, 0x3f, 0x5b,
	0x29, 0xb7, 0x68, 0x19, 0xe6, 0xa9, 0x6c, 0x72, 0xa4, 0x5e, 0x82, 0x46, 0x19, 0x1b, 0xe4, 0x92,
	0x95, 0x31, 0x59, 0xb3, 0x35, 0xb9, 0xf0, 0xad, 0xfc, 0xaa, 0xd9, 0x8b, 0x34, 0x67, 0x1b, 0x5f,
	0x0a, 0x40, 0x32, 0xb6, 0x9c, 0xdf, 0xbf, 0x41, 0xa4, 0x2e, 0x2c, 0x33, 0x6c, 0x66, 0x96, 0x4e,
	0x5b, 0xd3, 0x3a, 0x4d, 0x6d, 0xae, 0xbd, 0xc4, 0x26, 0x21, 0xe3, 0x87, 0x02, 0x80, 0xf3, 0x7e,
	0xd3, 0x28, 0xb5, 0x61, 0x7e, 0x18, 0x70, 0x1a, 0xf9, 0x7f, 0xd0, 0x28, 0x8d, 0xba, 0x85, 0x42,
	0x8f, 0xa0, 0x2a, 0x6e, 0x1f, 0xa7, 0x3c, 0x91, 0xa7, 0x31, 0xf3, 0xda, 0x91, 0xa8, 0xbd, 0x40,
	0xd3, 0x0f, 0xc5, 0x43, 0x28, 0xcb, 0x3f, 0x49, 0x0d, 0xfe, 0xbb, 0x51, 0x53, 0xf9, 0x47, 0x19,
	0x6c, 0xfc, 0x54, 0xa0, 0x86, 0x80, 0xdc, 0xf6, 0x3b, 0x6a, 0xfa, 0x29, 0xd4, 0x62, 0xac, 0x9f,
	0xeb, 0x7b, 0x73, 0xba, 0xef, 0xcc, 0x37, 0xc9, 0x56, 0xe3, 0xcc, 0x07, 0xea, 0x56, 0xdd, 0xef,
	0x7c, 0x55, 0x40, 0xc7, 0xab, 0xef, 0xf2, 0xe0, 0xd2, 0x19, 0xf2, 0x03, 0xdf, 0x75, 0xa9, 0xc7,
	0x7a, 0xfe, 0xd8, 0x19, 0x92, 0x15, 0xd0, 0x3a, 0xa6, 0x39, 0x38, 0xea, 0x9b, 0x9d, 0xa3, 0xc3,
	0x81, 0x75, 0x62, 0xbd, 0xee, 0x69, 0x73, 0x64, 0x1d, 0x96, 0xfb, 0xc7, 0x66, 0xa7, 0x67, 0xe5,
	0x1d, 0x8a, 0x70, 0x98, 0xd6, 0x2b, 0x6b, 0xd2, 0x51, 0x20, 0x1b, 0xb0, 0x6a, 0x5b, 0x3d, 0xfb,
	0xd0, 0x3a, 0x99, 0x70, 0x15, 0x49, 0x03, 0xd6, 0x66, 0xba, 0xba, 0x5a, 0x89, 0xd4, 0xa1, 0x2a,
	0xca, 0x1f, 0xdb, 0x87, 0x07, 0x96, 0x56, 0x26, 0x1a, 0xd4, 0x92, 0xba, 0x12, 0xa9, 0xec, 0x7c,
	0x53, 0x60, 0x2d, 0x4b, 0x1e, 0xcf, 0x92, 0xfa, 0x1a, 0x90, 0x6c, 0xba, 0x41, 0xc7, 0x34, 0x2d,
	0x53, 0x9b, 0x23, 0x3a, 0xac, 0xe4, 0x70, 0x99, 0xd1, 0xd4, 0x94, 0x29, 0x8f, 0x6c, 0xc5, 0xd4,
	0x0a, 0x82, 0x63, 0xce, 0x93, 0x12, 0x36, 0xb5, 0x22, 0xf9, 0x17, 0xd6, 0x73, 0xb4, 0x33, 0xce,
	0x12, 0xf9, 0x07, 0x54, 0xa4, 0x9a, 0x54, 0x2f, 0x93, 0x25, 0xa8, 0x4b, 0x20, 0x2d, 0x5b, 0x79,
	0xf6, 0xe4, 0xcd, 0xe3, 0x91, 0x13, 0x9d, 0xc5, 0xa7, 0xad, 0xa1, 0xef, 0xb6, 0x9d, 0x8b, 0x33,
	0xea, 0xba, 0x67, 0x8c, 0x25, 0x6f, 0x8e, 0x5d, 0x71, 0x7f, 0xef, 0xe8, 0xfb, 0xdd, 0x50, 0x76,
	0x28, 0x5e, 0x33, 0x43, 0x1a, 0xf2, 0xb0, 0x2d, 0x5f, 0x1e, 0xa7, 0x15, 0x7c, 0x73, 0xec, 0xff,
	0x0a, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x71, 0xaf, 0xed, 0x24, 0x09, 0x00, 0x00,
}
