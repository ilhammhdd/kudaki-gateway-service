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
)

var EventServiceCommandTopic_name = map[int32]string{
	0: "ADD_KUDAKI_EVENT",
	1: "UPDATE_KUDAKI_EVENT",
	2: "DELETE_KUDAKI_EVENT",
	3: "RETRIEVE_KUDAKI_EVENT",
	4: "RETRIEVE_KUDAKI_EVENTS",
}

var EventServiceCommandTopic_value = map[string]int32{
	"ADD_KUDAKI_EVENT":       0,
	"UPDATE_KUDAKI_EVENT":    1,
	"DELETE_KUDAKI_EVENT":    2,
	"RETRIEVE_KUDAKI_EVENT":  3,
	"RETRIEVE_KUDAKI_EVENTS": 4,
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
)

var EventServiceEventTopic_name = map[int32]string{
	0: "KUDAKI_EVENT_ADDED",
	1: "KUDAKI_EVENT_UPDATED",
	2: "KUDAKI_EVENT_DELETED",
	3: "KUDAKI_EVENT_RETRIEVED",
	4: "KUDAKI_EVENTS_RETRIEVED",
}

var EventServiceEventTopic_value = map[string]int32{
	"KUDAKI_EVENT_ADDED":      0,
	"KUDAKI_EVENT_UPDATED":    1,
	"KUDAKI_EVENT_DELETED":    2,
	"KUDAKI_EVENT_RETRIEVED":  3,
	"KUDAKI_EVENTS_RETRIEVED": 4,
}

func (x EventServiceEventTopic) String() string {
	return proto.EnumName(EventServiceEventTopic_name, int32(x))
}

func (EventServiceEventTopic) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fa36ed941739bd5b, []int{1}
}

type EventPaymentServiceEventTopic int32

const (
	EventPaymentServiceEventTopic_EVENT_DOKU_INVOICE_ISSUED   EventPaymentServiceEventTopic = 0
	EventPaymentServiceEventTopic_EVENT_DOKU_PAYMENT_IDENTIFY EventPaymentServiceEventTopic = 1
	EventPaymentServiceEventTopic_EVENT_DOKU_PAYMENT_REDIRECT EventPaymentServiceEventTopic = 2
	EventPaymentServiceEventTopic_EVENT_DOKU_PAYMENT_NOTIFY   EventPaymentServiceEventTopic = 3
)

var EventPaymentServiceEventTopic_name = map[int32]string{
	0: "EVENT_DOKU_INVOICE_ISSUED",
	1: "EVENT_DOKU_PAYMENT_IDENTIFY",
	2: "EVENT_DOKU_PAYMENT_REDIRECT",
	3: "EVENT_DOKU_PAYMENT_NOTIFY",
}

var EventPaymentServiceEventTopic_value = map[string]int32{
	"EVENT_DOKU_INVOICE_ISSUED":   0,
	"EVENT_DOKU_PAYMENT_IDENTIFY": 1,
	"EVENT_DOKU_PAYMENT_REDIRECT": 2,
	"EVENT_DOKU_PAYMENT_NOTIFY":   3,
}

func (x EventPaymentServiceEventTopic) String() string {
	return proto.EnumName(EventPaymentServiceEventTopic_name, int32(x))
}

func (EventPaymentServiceEventTopic) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fa36ed941739bd5b, []int{2}
}

type AddKudakiEvent struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Venue                string   `protobuf:"bytes,2,opt,name=venue,proto3" json:"venue,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	DurationFrom         int64    `protobuf:"varint,4,opt,name=duration_from,json=durationFrom,proto3" json:"duration_from,omitempty"`
	DurationTo           int64    `protobuf:"varint,5,opt,name=duration_to,json=durationTo,proto3" json:"duration_to,omitempty"`
	AdDurationFrom       int64    `protobuf:"varint,7,opt,name=ad_duration_from,json=adDurationFrom,proto3" json:"ad_duration_from,omitempty"`
	AdDurationTo         int64    `protobuf:"varint,8,opt,name=ad_duration_to,json=adDurationTo,proto3" json:"ad_duration_to,omitempty"`
	KudakiToken          string   `protobuf:"bytes,9,opt,name=kudaki_token,json=kudakiToken,proto3" json:"kudaki_token,omitempty"`
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

func (m *AddKudakiEvent) GetName() string {
	if m != nil {
		return m.Name
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

func (m *AddKudakiEvent) GetAdDurationFrom() int64 {
	if m != nil {
		return m.AdDurationFrom
	}
	return 0
}

func (m *AddKudakiEvent) GetAdDurationTo() int64 {
	if m != nil {
		return m.AdDurationTo
	}
	return 0
}

func (m *AddKudakiEvent) GetKudakiToken() string {
	if m != nil {
		return m.KudakiToken
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

type KudakiEventAdded struct {
	Uid                  string                    `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Organizer            *user.Profile             `protobuf:"bytes,2,opt,name=organizer,proto3" json:"organizer,omitempty"`
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
	return fileDescriptor_fa36ed941739bd5b, []int{2}
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

func (m *KudakiEventAdded) GetOrganizer() *user.Profile {
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
	return fileDescriptor_fa36ed941739bd5b, []int{3}
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

type KudakiEventDokuInvoiceIssued struct {
	Uid                  string                    `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Organizer            *user.Profile             `protobuf:"bytes,2,opt,name=organizer,proto3" json:"organizer,omitempty"`
	EventStatus          *Status                   `protobuf:"bytes,3,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	DokuInvoice          *kudaki_event.DokuInvoice `protobuf:"bytes,4,opt,name=doku_invoice,json=dokuInvoice,proto3" json:"doku_invoice,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *KudakiEventDokuInvoiceIssued) Reset()         { *m = KudakiEventDokuInvoiceIssued{} }
func (m *KudakiEventDokuInvoiceIssued) String() string { return proto.CompactTextString(m) }
func (*KudakiEventDokuInvoiceIssued) ProtoMessage()    {}
func (*KudakiEventDokuInvoiceIssued) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa36ed941739bd5b, []int{4}
}

func (m *KudakiEventDokuInvoiceIssued) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KudakiEventDokuInvoiceIssued.Unmarshal(m, b)
}
func (m *KudakiEventDokuInvoiceIssued) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KudakiEventDokuInvoiceIssued.Marshal(b, m, deterministic)
}
func (m *KudakiEventDokuInvoiceIssued) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KudakiEventDokuInvoiceIssued.Merge(m, src)
}
func (m *KudakiEventDokuInvoiceIssued) XXX_Size() int {
	return xxx_messageInfo_KudakiEventDokuInvoiceIssued.Size(m)
}
func (m *KudakiEventDokuInvoiceIssued) XXX_DiscardUnknown() {
	xxx_messageInfo_KudakiEventDokuInvoiceIssued.DiscardUnknown(m)
}

var xxx_messageInfo_KudakiEventDokuInvoiceIssued proto.InternalMessageInfo

func (m *KudakiEventDokuInvoiceIssued) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *KudakiEventDokuInvoiceIssued) GetOrganizer() *user.Profile {
	if m != nil {
		return m.Organizer
	}
	return nil
}

func (m *KudakiEventDokuInvoiceIssued) GetEventStatus() *Status {
	if m != nil {
		return m.EventStatus
	}
	return nil
}

func (m *KudakiEventDokuInvoiceIssued) GetDokuInvoice() *kudaki_event.DokuInvoice {
	if m != nil {
		return m.DokuInvoice
	}
	return nil
}

func init() {
	proto.RegisterEnum("aggregates.event.EventServiceCommandTopic", EventServiceCommandTopic_name, EventServiceCommandTopic_value)
	proto.RegisterEnum("aggregates.event.EventServiceEventTopic", EventServiceEventTopic_name, EventServiceEventTopic_value)
	proto.RegisterEnum("aggregates.event.EventPaymentServiceEventTopic", EventPaymentServiceEventTopic_name, EventPaymentServiceEventTopic_value)
	proto.RegisterType((*AddKudakiEvent)(nil), "aggregates.event.AddKudakiEvent")
	proto.RegisterType((*DeleteKudakiEvent)(nil), "aggregates.event.DeleteKudakiEvent")
	proto.RegisterType((*KudakiEventAdded)(nil), "aggregates.event.KudakiEventAdded")
	proto.RegisterType((*KudakiEventDeleted)(nil), "aggregates.event.KudakiEventDeleted")
	proto.RegisterType((*KudakiEventDokuInvoiceIssued)(nil), "aggregates.event.KudakiEventDokuInvoiceIssued")
}

func init() { proto.RegisterFile("events/kudaki_event.proto", fileDescriptor_fa36ed941739bd5b) }

var fileDescriptor_fa36ed941739bd5b = []byte{
	// 756 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0x4d, 0x4f, 0xdb, 0x4a,
	0x14, 0xc5, 0x49, 0xe0, 0xbd, 0xdc, 0x00, 0x32, 0xc3, 0x97, 0x09, 0x2f, 0x22, 0x2f, 0xb0, 0x40,
	0x91, 0x48, 0x2a, 0x90, 0xba, 0xa9, 0xba, 0x48, 0x19, 0x53, 0xb9, 0x69, 0x43, 0xe4, 0x38, 0x48,
	0x74, 0x63, 0x99, 0xcc, 0x10, 0xac, 0x60, 0x4f, 0xe4, 0x8f, 0x54, 0xf4, 0x9f, 0x74, 0x51, 0x75,
	0xd1, 0x45, 0x17, 0xfd, 0x59, 0xfd, 0x23, 0x95, 0x67, 0x12, 0x32, 0x4e, 0xd2, 0xb2, 0x6c, 0x37,
	0xc8, 0xf7, 0x9c, 0xc3, 0xbd, 0xc7, 0x67, 0x6e, 0x3c, 0xb0, 0x47, 0x47, 0xd4, 0x8f, 0xc2, 0xfa,
	0x20, 0x26, 0xce, 0xc0, 0xb5, 0x79, 0x55, 0x1b, 0x06, 0x2c, 0x62, 0x48, 0x75, 0xfa, 0xfd, 0x80,
	0xf6, 0x9d, 0x88, 0x86, 0x35, 0x8e, 0x17, 0x8b, 0x53, 0xa4, 0x1e, 0x87, 0x34, 0xe0, 0x7f, 0x84,
	0xba, 0x58, 0x95, 0x38, 0xb9, 0x59, 0x9d, 0xb0, 0x41, 0x6c, 0xbb, 0xfe, 0x88, 0xb9, 0x3d, 0x3a,
	0xd6, 0x96, 0x66, 0xfb, 0x0c, 0x03, 0x76, 0xeb, 0xde, 0xd3, 0xa7, 0x5a, 0xcd, 0x9b, 0x2c, 0x6e,
	0x8e, 0xfd, 0x87, 0x91, 0x13, 0xc5, 0xa1, 0x00, 0x2b, 0xdf, 0x32, 0xb0, 0xde, 0x20, 0xa4, 0xc9,
	0xe5, 0x7a, 0x22, 0x40, 0x2a, 0x64, 0x63, 0x97, 0x68, 0x4a, 0x59, 0x39, 0xce, 0x9b, 0xc9, 0x23,
	0x42, 0x90, 0xf3, 0x1d, 0x8f, 0x6a, 0x2b, 0x1c, 0xe2, 0xcf, 0x68, 0x0b, 0x96, 0x47, 0xd4, 0x8f,
	0xa9, 0x96, 0xe1, 0xa0, 0x28, 0x50, 0x19, 0x0a, 0x84, 0x86, 0xbd, 0xc0, 0x1d, 0x46, 0x2e, 0xf3,
	0xb5, 0x2c, 0xe7, 0x64, 0x08, 0x1d, 0xc2, 0x1a, 0x89, 0x03, 0x27, 0x79, 0xb6, 0x6f, 0x03, 0xe6,
	0x69, 0xb9, 0xb2, 0x72, 0x9c, 0x35, 0x57, 0x27, 0xe0, 0x45, 0xc0, 0x3c, 0x74, 0x00, 0x85, 0x47,
	0x51, 0xc4, 0xb4, 0x65, 0x2e, 0x81, 0x09, 0x64, 0x31, 0x74, 0x0c, 0xaa, 0x43, 0xec, 0x74, 0xa3,
	0x7f, 0xb8, 0x6a, 0xdd, 0x21, 0x58, 0x6e, 0x75, 0x04, 0xeb, 0xb2, 0x32, 0x62, 0xda, 0xbf, 0x62,
	0xe0, 0x54, 0x67, 0x31, 0xf4, 0x3f, 0xac, 0x8e, 0x13, 0x8b, 0xd8, 0x80, 0xfa, 0x5a, 0x5e, 0x18,
	0x17, 0x98, 0x95, 0x40, 0x95, 0x3e, 0x6c, 0x60, 0x7a, 0x4f, 0x23, 0xfa, 0xfb, 0xac, 0x4a, 0x00,
	0x3c, 0x67, 0x3b, 0x4e, 0x08, 0x11, 0x4e, 0x9e, 0x23, 0xdd, 0x84, 0x9e, 0x1d, 0x94, 0x9d, 0x1f,
	0xf4, 0x35, 0x03, 0xaa, 0x34, 0xa3, 0x41, 0x08, 0x25, 0x0b, 0x06, 0x3d, 0x87, 0x3c, 0x0b, 0xfa,
	0x8e, 0xef, 0x7e, 0xa4, 0x01, 0x9f, 0x53, 0x38, 0xd5, 0x6a, 0xd2, 0x1e, 0xf2, 0x85, 0x6b, 0x8b,
	0x6d, 0x31, 0xa7, 0x52, 0xf4, 0x0c, 0x56, 0x85, 0x41, 0xb1, 0x07, 0xdc, 0x41, 0xe1, 0x74, 0x4d,
	0xec, 0x6d, 0xad, 0xc3, 0x41, 0xb3, 0xc0, 0x2b, 0x51, 0xa0, 0xd7, 0x8f, 0x9e, 0x39, 0xca, 0x4f,
	0xac, 0x70, 0x7a, 0x24, 0x0f, 0x4b, 0xad, 0x9b, 0x64, 0x7e, 0xf2, 0x66, 0x22, 0xad, 0x37, 0xc9,
	0xa9, 0x11, 0x3b, 0xd5, 0x6c, 0x99, 0x37, 0x2b, 0xd7, 0x66, 0x7f, 0x41, 0xb5, 0xf4, 0x56, 0x26,
	0xe7, 0x2a, 0xd7, 0x95, 0xef, 0x19, 0x40, 0x52, 0x2d, 0x8e, 0x66, 0x51, 0x4e, 0x67, 0xf3, 0x39,
	0x6d, 0xcf, 0xe5, 0xd4, 0x0d, 0x69, 0xf0, 0x97, 0x84, 0xd4, 0x81, 0x4d, 0xc2, 0x5f, 0x66, 0x51,
	0x4e, 0x87, 0xf3, 0x39, 0xcd, 0x2d, 0xa5, 0xb9, 0x41, 0x66, 0xa1, 0xca, 0x0f, 0x05, 0xfe, 0x93,
	0xd3, 0x62, 0x83, 0xd8, 0x10, 0xdf, 0x19, 0x23, 0x0c, 0xe3, 0x3f, 0xbf, 0x5f, 0xf2, 0x97, 0xef,
	0xc9, 0xe8, 0x24, 0xf7, 0x66, 0x81, 0x4c, 0x8b, 0xea, 0x27, 0x05, 0x34, 0xfe, 0x7e, 0x1d, 0x1a,
	0x8c, 0xdc, 0x1e, 0x3d, 0x67, 0x9e, 0xe7, 0xf8, 0xc4, 0x62, 0x43, 0xb7, 0x87, 0xb6, 0x40, 0x6d,
	0x60, 0x6c, 0x37, 0xbb, 0xb8, 0xd1, 0x34, 0x6c, 0xfd, 0x4a, 0x6f, 0x59, 0xea, 0x12, 0xda, 0x85,
	0xcd, 0x6e, 0x1b, 0x37, 0x2c, 0x3d, 0x4d, 0x28, 0x09, 0x81, 0xf5, 0xb7, 0xfa, 0x2c, 0x91, 0x41,
	0x7b, 0xb0, 0x6d, 0xea, 0x96, 0x69, 0xe8, 0x57, 0x33, 0x54, 0x16, 0x15, 0x61, 0x67, 0x21, 0xd5,
	0x51, 0x73, 0xd5, 0xcf, 0x0a, 0xec, 0xc8, 0xde, 0xf8, 0xb3, 0x70, 0xb6, 0x03, 0x48, 0x56, 0xdb,
	0x0d, 0x8c, 0x75, 0xac, 0x2e, 0x21, 0x0d, 0xb6, 0x52, 0xb8, 0x30, 0x8a, 0x55, 0x65, 0x8e, 0x11,
	0x4e, 0xb1, 0x9a, 0x49, 0x2c, 0xa4, 0x98, 0x89, 0x1f, 0xac, 0x66, 0xd1, 0x3e, 0xec, 0xa6, 0x5c,
	0x49, 0x64, 0xae, 0xfa, 0x45, 0x81, 0x12, 0xf7, 0xd4, 0x76, 0x1e, 0xbc, 0x85, 0x36, 0x4b, 0xb0,
	0x37, 0x9e, 0x76, 0xd9, 0xec, 0xda, 0x46, 0xeb, 0xea, 0xd2, 0x38, 0xd7, 0x6d, 0xa3, 0xd3, 0xe9,
	0x72, 0xb7, 0x07, 0xb0, 0x2f, 0xd1, 0xed, 0xc6, 0xf5, 0xbb, 0xa4, 0x30, 0xb0, 0xde, 0xb2, 0x8c,
	0x8b, 0x6b, 0x55, 0xf9, 0x85, 0xc0, 0xd4, 0xb1, 0x61, 0xea, 0xe7, 0x49, 0xb2, 0xe9, 0x01, 0x13,
	0x41, 0xeb, 0x92, 0xff, 0x7f, 0xf6, 0xd5, 0xcb, 0xf7, 0x2f, 0xfa, 0x6e, 0x74, 0x17, 0xdf, 0xd4,
	0x7a, 0xcc, 0xab, 0xbb, 0xf7, 0x77, 0x8e, 0xe7, 0xdd, 0x11, 0x32, 0xbe, 0xea, 0x4e, 0x92, 0x45,
	0xf9, 0xe0, 0x3c, 0x9c, 0x84, 0xc2, 0x78, 0x72, 0x5b, 0xf6, 0x9c, 0x90, 0x86, 0x75, 0x71, 0xed,
	0xdd, 0xac, 0xf0, 0x0b, 0xef, 0xec, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb4, 0xb4, 0x47, 0xe9,
	0xc7, 0x07, 0x00, 0x00,
}
