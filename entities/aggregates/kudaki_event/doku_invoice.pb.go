// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/kudaki_event/doku_invoice.proto

package kudaki_event

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type DokuInvoice struct {
	Id                    int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid                  string   `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	KudakiEventUuid       string   `protobuf:"bytes,3,opt,name=kudaki_event_uuid,json=kudakiEventUuid,proto3" json:"kudaki_event_uuid,omitempty"`
	Amount                float32  `protobuf:"fixed32,4,opt,name=amount,proto3" json:"amount,omitempty"`
	PurchaseAmount        float32  `protobuf:"fixed32,5,opt,name=purchase_amount,json=purchaseAmount,proto3" json:"purchase_amount,omitempty"`
	TransactionIdMerchant string   `protobuf:"bytes,6,opt,name=transaction_id_merchant,json=transactionIdMerchant,proto3" json:"transaction_id_merchant,omitempty"`
	Words                 string   `protobuf:"bytes,7,opt,name=words,proto3" json:"words,omitempty"`
	RequestDateTime       int64    `protobuf:"varint,8,opt,name=request_date_time,json=requestDateTime,proto3" json:"request_date_time,omitempty"`
	Currency              int32    `protobuf:"varint,9,opt,name=currency,proto3" json:"currency,omitempty"`
	PurchaseCurrency      int32    `protobuf:"varint,10,opt,name=purchase_currency,json=purchaseCurrency,proto3" json:"purchase_currency,omitempty"`
	SessionId             string   `protobuf:"bytes,11,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Name                  string   `protobuf:"bytes,12,opt,name=name,proto3" json:"name,omitempty"`
	Email                 string   `protobuf:"bytes,13,opt,name=email,proto3" json:"email,omitempty"`
	Basket                string   `protobuf:"bytes,14,opt,name=basket,proto3" json:"basket,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *DokuInvoice) Reset()         { *m = DokuInvoice{} }
func (m *DokuInvoice) String() string { return proto.CompactTextString(m) }
func (*DokuInvoice) ProtoMessage()    {}
func (*DokuInvoice) Descriptor() ([]byte, []int) {
	return fileDescriptor_60136d5b6d0fd8b0, []int{0}
}

func (m *DokuInvoice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DokuInvoice.Unmarshal(m, b)
}
func (m *DokuInvoice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DokuInvoice.Marshal(b, m, deterministic)
}
func (m *DokuInvoice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DokuInvoice.Merge(m, src)
}
func (m *DokuInvoice) XXX_Size() int {
	return xxx_messageInfo_DokuInvoice.Size(m)
}
func (m *DokuInvoice) XXX_DiscardUnknown() {
	xxx_messageInfo_DokuInvoice.DiscardUnknown(m)
}

var xxx_messageInfo_DokuInvoice proto.InternalMessageInfo

func (m *DokuInvoice) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DokuInvoice) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *DokuInvoice) GetKudakiEventUuid() string {
	if m != nil {
		return m.KudakiEventUuid
	}
	return ""
}

func (m *DokuInvoice) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *DokuInvoice) GetPurchaseAmount() float32 {
	if m != nil {
		return m.PurchaseAmount
	}
	return 0
}

func (m *DokuInvoice) GetTransactionIdMerchant() string {
	if m != nil {
		return m.TransactionIdMerchant
	}
	return ""
}

func (m *DokuInvoice) GetWords() string {
	if m != nil {
		return m.Words
	}
	return ""
}

func (m *DokuInvoice) GetRequestDateTime() int64 {
	if m != nil {
		return m.RequestDateTime
	}
	return 0
}

func (m *DokuInvoice) GetCurrency() int32 {
	if m != nil {
		return m.Currency
	}
	return 0
}

func (m *DokuInvoice) GetPurchaseCurrency() int32 {
	if m != nil {
		return m.PurchaseCurrency
	}
	return 0
}

func (m *DokuInvoice) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *DokuInvoice) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DokuInvoice) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *DokuInvoice) GetBasket() string {
	if m != nil {
		return m.Basket
	}
	return ""
}

func init() {
	proto.RegisterType((*DokuInvoice)(nil), "aggregates.kudaki_event.DokuInvoice")
}

func init() {
	proto.RegisterFile("aggregates/kudaki_event/doku_invoice.proto", fileDescriptor_60136d5b6d0fd8b0)
}

var fileDescriptor_60136d5b6d0fd8b0 = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4f, 0x8f, 0xd3, 0x30,
	0x10, 0xc5, 0x95, 0x74, 0x5b, 0xb6, 0x5e, 0x68, 0x59, 0x0b, 0x58, 0x0b, 0x09, 0x29, 0xe2, 0x42,
	0xb4, 0x68, 0x9b, 0x03, 0x12, 0x77, 0x60, 0x39, 0x54, 0x02, 0x0e, 0x11, 0x5c, 0xb8, 0x58, 0x6e,
	0x3c, 0x4a, 0x46, 0x59, 0xdb, 0xc5, 0x7f, 0x5a, 0xf5, 0xeb, 0xf1, 0xc9, 0x90, 0x9d, 0xb4, 0xf4,
	0xb2, 0xb7, 0x79, 0xef, 0xfd, 0x94, 0xbc, 0x4c, 0x86, 0xdc, 0x8a, 0xb6, 0xb5, 0xd0, 0x0a, 0x0f,
	0xae, 0xea, 0x83, 0x14, 0x3d, 0x72, 0xd8, 0x81, 0xf6, 0x95, 0x34, 0x7d, 0xe0, 0xa8, 0x77, 0x06,
	0x1b, 0x58, 0x6d, 0xad, 0xf1, 0x86, 0xde, 0xfc, 0x67, 0x57, 0xe7, 0xec, 0xdb, 0xbf, 0x13, 0x72,
	0x75, 0x6f, 0xfa, 0xb0, 0x1e, 0x70, 0xba, 0x20, 0x39, 0x4a, 0x96, 0x15, 0x59, 0x39, 0xa9, 0x73,
	0x94, 0x94, 0x92, 0x8b, 0x10, 0x50, 0xb2, 0xbc, 0xc8, 0xca, 0x79, 0x9d, 0x66, 0x7a, 0x4b, 0xae,
	0xcf, 0x9f, 0xc1, 0x13, 0x30, 0x49, 0xc0, 0x72, 0x08, 0xbe, 0x46, 0xff, 0x57, 0x64, 0x5f, 0x91,
	0x99, 0x50, 0x26, 0x68, 0xcf, 0x2e, 0x8a, 0xac, 0xcc, 0xeb, 0x51, 0xd1, 0x77, 0x64, 0xb9, 0x0d,
	0xb6, 0xe9, 0x84, 0x03, 0x3e, 0x02, 0xd3, 0x04, 0x2c, 0x8e, 0xf6, 0xa7, 0x01, 0xfc, 0x48, 0x6e,
	0xbc, 0x15, 0xda, 0x89, 0xc6, 0xa3, 0xd1, 0x1c, 0x25, 0x57, 0x10, 0x01, 0xed, 0xd9, 0x2c, 0xbd,
	0xf2, 0xe5, 0x59, 0xbc, 0x96, 0xdf, 0xc7, 0x90, 0xbe, 0x20, 0xd3, 0xbd, 0xb1, 0xd2, 0xb1, 0x27,
	0x89, 0x1a, 0x44, 0xac, 0x6e, 0xe1, 0x4f, 0x00, 0xe7, 0xb9, 0x14, 0x1e, 0xb8, 0x47, 0x05, 0xec,
	0x32, 0x7d, 0xed, 0x72, 0x0c, 0xee, 0x85, 0x87, 0x9f, 0xa8, 0x80, 0xbe, 0x26, 0x97, 0x4d, 0xb0,
	0x16, 0x74, 0x73, 0x60, 0xf3, 0x22, 0x2b, 0xa7, 0xf5, 0x49, 0xd3, 0xf7, 0xe4, 0xfa, 0x54, 0xff,
	0x04, 0x91, 0x04, 0x3d, 0x3f, 0x06, 0x5f, 0x8e, 0xf0, 0x1b, 0x42, 0x1c, 0x38, 0x37, 0xd4, 0x67,
	0x57, 0xa9, 0xcf, 0x7c, 0x74, 0xd6, 0x69, 0xc5, 0x5a, 0x28, 0x60, 0x4f, 0x87, 0x15, 0xc7, 0x39,
	0xb6, 0x07, 0x25, 0xf0, 0x81, 0x3d, 0x1b, 0xda, 0x27, 0x11, 0x97, 0xb9, 0x11, 0xae, 0x07, 0xcf,
	0x16, 0xc9, 0x1e, 0xd5, 0xe7, 0x1f, 0xbf, 0xbf, 0xb5, 0xe8, 0xbb, 0xb0, 0x59, 0x35, 0x46, 0x55,
	0xf8, 0xd0, 0x09, 0xa5, 0x3a, 0x29, 0xc7, 0xab, 0xb8, 0x8b, 0xbf, 0x7d, 0x2f, 0x0e, 0x77, 0x0e,
	0xec, 0x0e, 0x1b, 0xa8, 0x40, 0x7b, 0xf4, 0x08, 0xae, 0x7a, 0xe4, 0x80, 0x36, 0xb3, 0x74, 0x34,
	0x1f, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x47, 0x91, 0x4b, 0x06, 0x62, 0x02, 0x00, 0x00,
}