// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/kudaki_event/invoice.proto

package kudaki_event

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type InvoiceStatus int32

const (
	InvoiceStatus_UNPAID InvoiceStatus = 0
	InvoiceStatus_PAID   InvoiceStatus = 1
)

var InvoiceStatus_name = map[int32]string{
	0: "UNPAID",
	1: "PAID",
}

var InvoiceStatus_value = map[string]int32{
	"UNPAID": 0,
	"PAID":   1,
}

func (x InvoiceStatus) String() string {
	return proto.EnumName(InvoiceStatus_name, int32(x))
}

func (InvoiceStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_70f0e5555df8a552, []int{0}
}

type Invoice struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid                 string               `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	KudakiEvent          *KudakiEvent         `protobuf:"bytes,3,opt,name=kudaki_event,json=kudakiEvent,proto3" json:"kudaki_event,omitempty"`
	TotalPrice           int32                `protobuf:"varint,4,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	Status               InvoiceStatus        `protobuf:"varint,5,opt,name=status,proto3,enum=aggregates.kudaki_event.InvoiceStatus" json:"status,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Invoice) Reset()         { *m = Invoice{} }
func (m *Invoice) String() string { return proto.CompactTextString(m) }
func (*Invoice) ProtoMessage()    {}
func (*Invoice) Descriptor() ([]byte, []int) {
	return fileDescriptor_70f0e5555df8a552, []int{0}
}

func (m *Invoice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Invoice.Unmarshal(m, b)
}
func (m *Invoice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Invoice.Marshal(b, m, deterministic)
}
func (m *Invoice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Invoice.Merge(m, src)
}
func (m *Invoice) XXX_Size() int {
	return xxx_messageInfo_Invoice.Size(m)
}
func (m *Invoice) XXX_DiscardUnknown() {
	xxx_messageInfo_Invoice.DiscardUnknown(m)
}

var xxx_messageInfo_Invoice proto.InternalMessageInfo

func (m *Invoice) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Invoice) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Invoice) GetKudakiEvent() *KudakiEvent {
	if m != nil {
		return m.KudakiEvent
	}
	return nil
}

func (m *Invoice) GetTotalPrice() int32 {
	if m != nil {
		return m.TotalPrice
	}
	return 0
}

func (m *Invoice) GetStatus() InvoiceStatus {
	if m != nil {
		return m.Status
	}
	return InvoiceStatus_UNPAID
}

func (m *Invoice) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterEnum("aggregates.kudaki_event.InvoiceStatus", InvoiceStatus_name, InvoiceStatus_value)
	proto.RegisterType((*Invoice)(nil), "aggregates.kudaki_event.Invoice")
}

func init() {
	proto.RegisterFile("aggregates/kudaki_event/invoice.proto", fileDescriptor_70f0e5555df8a552)
}

var fileDescriptor_70f0e5555df8a552 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x41, 0x6b, 0xfa, 0x40,
	0x10, 0xc5, 0xff, 0x89, 0x9a, 0x7f, 0x1d, 0x5b, 0x29, 0x7b, 0x69, 0xf0, 0x62, 0x28, 0xb5, 0x04,
	0xc1, 0x0d, 0xd8, 0x53, 0x2f, 0x05, 0x4b, 0x4b, 0x91, 0x16, 0x91, 0xb4, 0xbd, 0xf4, 0x22, 0x6b,
	0x76, 0x1a, 0x17, 0x8d, 0x2b, 0xc9, 0xc4, 0xd2, 0xaf, 0xd1, 0x4f, 0x5c, 0xdc, 0x44, 0xb4, 0x87,
	0xdc, 0xe6, 0x2d, 0xbf, 0x99, 0xf7, 0x78, 0x0b, 0x3d, 0x11, 0xc7, 0x29, 0xc6, 0x82, 0x30, 0x0b,
	0x96, 0xb9, 0x14, 0x4b, 0x35, 0xc3, 0x2d, 0xae, 0x29, 0x50, 0xeb, 0xad, 0x56, 0x11, 0xf2, 0x4d,
	0xaa, 0x49, 0xb3, 0x8b, 0x03, 0xc6, 0x8f, 0xb1, 0x4e, 0x37, 0xd6, 0x3a, 0x5e, 0x61, 0x60, 0xb0,
	0x79, 0xfe, 0x19, 0x90, 0x4a, 0x30, 0x23, 0x91, 0x6c, 0x8a, 0xcd, 0x4e, 0xbf, 0xca, 0xe0, 0x58,
	0x14, 0xec, 0xe5, 0x8f, 0x0d, 0xff, 0xc7, 0x85, 0x2f, 0x6b, 0x83, 0xad, 0xa4, 0x6b, 0x79, 0x96,
	0x5f, 0x0b, 0x6d, 0x25, 0x19, 0x83, 0x7a, 0x9e, 0x2b, 0xe9, 0xda, 0x9e, 0xe5, 0x37, 0x43, 0x33,
	0xb3, 0x27, 0x38, 0x3d, 0xbe, 0xe2, 0xd6, 0x3c, 0xcb, 0x6f, 0x0d, 0xaf, 0x78, 0x45, 0x58, 0xfe,
	0x6c, 0xc4, 0xe3, 0x6e, 0x0e, 0x5b, 0xcb, 0x83, 0x60, 0x5d, 0x68, 0x91, 0x26, 0xb1, 0x9a, 0x6d,
	0x52, 0x15, 0xa1, 0x5b, 0xf7, 0x2c, 0xbf, 0x11, 0x82, 0x79, 0x9a, 0xee, 0x5e, 0xd8, 0x1d, 0x38,
	0x19, 0x09, 0xca, 0x33, 0xb7, 0xe1, 0x59, 0x7e, 0x7b, 0x78, 0x5d, 0xe9, 0x51, 0xe6, 0x7f, 0x35,
	0x74, 0x58, 0x6e, 0xb1, 0x5b, 0x80, 0x28, 0x45, 0x41, 0x28, 0x67, 0x82, 0x5c, 0xc7, 0xe4, 0xec,
	0xf0, 0xa2, 0x3b, 0xbe, 0xef, 0x8e, 0xbf, 0xed, 0xbb, 0x0b, 0x9b, 0x25, 0x3d, 0xa2, 0x7e, 0x0f,
	0xce, 0xfe, 0xdc, 0x64, 0x00, 0xce, 0xfb, 0x64, 0x3a, 0x1a, 0x3f, 0x9c, 0xff, 0x63, 0x27, 0x50,
	0x37, 0x93, 0x75, 0x3f, 0xf9, 0x78, 0x89, 0x15, 0x2d, 0xf2, 0x39, 0x8f, 0x74, 0x12, 0xa8, 0xd5,
	0x42, 0x24, 0xc9, 0x42, 0xca, 0xb2, 0xe6, 0xc1, 0x2e, 0xe9, 0x97, 0xf8, 0x1e, 0x64, 0x98, 0x6e,
	0x55, 0x84, 0x01, 0xae, 0x49, 0x91, 0xc2, 0x2c, 0xa8, 0xf8, 0x9e, 0xb9, 0x63, 0x52, 0xdd, 0xfc,
	0x06, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x50, 0x3a, 0x8b, 0x21, 0x02, 0x00, 0x00,
}
