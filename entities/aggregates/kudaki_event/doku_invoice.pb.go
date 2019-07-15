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
	MallId                int32    `protobuf:"varint,15,opt,name=mall_id,json=mallId,proto3" json:"mall_id,omitempty"`
	ChainMerchant         int32    `protobuf:"varint,16,opt,name=chain_merchant,json=chainMerchant,proto3" json:"chain_merchant,omitempty"`
	Amount                int64    `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	PurchaseAmount        int64    `protobuf:"varint,5,opt,name=purchase_amount,json=purchaseAmount,proto3" json:"purchase_amount,omitempty"`
	TransactionIdMerchant string   `protobuf:"bytes,6,opt,name=transaction_id_merchant,json=transactionIdMerchant,proto3" json:"transaction_id_merchant,omitempty"`
	Words                 []byte   `protobuf:"bytes,7,opt,name=words,proto3" json:"words,omitempty"`
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

func (m *DokuInvoice) GetMallId() int32 {
	if m != nil {
		return m.MallId
	}
	return 0
}

func (m *DokuInvoice) GetChainMerchant() int32 {
	if m != nil {
		return m.ChainMerchant
	}
	return 0
}

func (m *DokuInvoice) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *DokuInvoice) GetPurchaseAmount() int64 {
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

func (m *DokuInvoice) GetWords() []byte {
	if m != nil {
		return m.Words
	}
	return nil
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
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4f, 0x6f, 0x13, 0x31,
	0x10, 0xc5, 0xb5, 0x69, 0x92, 0x36, 0x6e, 0x9b, 0xb4, 0x16, 0x10, 0x0b, 0x09, 0x29, 0x42, 0x42,
	0x44, 0x45, 0xcd, 0x1e, 0x90, 0xb8, 0x03, 0xe5, 0x10, 0x09, 0x38, 0x44, 0x70, 0xe1, 0x62, 0x39,
	0xeb, 0x51, 0x76, 0xb4, 0xb1, 0x5d, 0xfc, 0x27, 0x55, 0xaf, 0x7c, 0x72, 0xe4, 0xd9, 0x4d, 0x9a,
	0x4b, 0x6f, 0x7e, 0xef, 0xfd, 0xb4, 0xf3, 0xbc, 0x1e, 0x76, 0xa3, 0x36, 0x1b, 0x0f, 0x1b, 0x15,
	0x21, 0x94, 0x4d, 0xd2, 0xaa, 0x41, 0x09, 0x3b, 0xb0, 0xb1, 0xd4, 0xae, 0x49, 0x12, 0xed, 0xce,
	0x61, 0x05, 0x8b, 0x7b, 0xef, 0xa2, 0xe3, 0xd3, 0x27, 0x76, 0x71, 0xcc, 0xbe, 0xfd, 0xd7, 0x67,
	0xe7, 0x77, 0xae, 0x49, 0xcb, 0x16, 0xe7, 0x63, 0xd6, 0x43, 0x2d, 0x8a, 0x59, 0x31, 0x3f, 0x59,
	0xf5, 0x50, 0x73, 0xce, 0xfa, 0x29, 0xa1, 0x16, 0xbd, 0x59, 0x31, 0x1f, 0xad, 0xe8, 0xcc, 0x6f,
	0xd8, 0xf5, 0xf1, 0x37, 0x24, 0x01, 0x27, 0x04, 0x4c, 0xda, 0xe0, 0x5b, 0xf6, 0x7f, 0x67, 0x76,
	0xca, 0x4e, 0x8d, 0xda, 0x6e, 0x25, 0x6a, 0x31, 0x99, 0x15, 0xf3, 0xc1, 0x6a, 0x98, 0xe5, 0x52,
	0xf3, 0x77, 0x6c, 0x5c, 0xd5, 0x0a, 0xad, 0x34, 0xe0, 0xab, 0x5a, 0xd9, 0x28, 0xae, 0x28, 0xbf,
	0x24, 0xf7, 0x47, 0x67, 0xf2, 0x57, 0x6c, 0xa8, 0x8c, 0x4b, 0x36, 0x8a, 0x3e, 0x75, 0xea, 0x14,
	0x7f, 0xcf, 0x26, 0xf7, 0x29, 0x33, 0x01, 0x64, 0x07, 0x0c, 0x08, 0x18, 0xef, 0xed, 0xcf, 0x2d,
	0xf8, 0x89, 0x4d, 0xa3, 0x57, 0x36, 0xa8, 0x2a, 0xa2, 0xb3, 0x12, 0xf5, 0xd3, 0xc0, 0x21, 0x55,
	0x7e, 0x79, 0x14, 0x2f, 0xf5, 0x61, 0xf0, 0x0b, 0x36, 0x78, 0x70, 0x5e, 0x07, 0x71, 0x3a, 0x2b,
	0xe6, 0x17, 0xab, 0x56, 0xe4, 0xab, 0x7b, 0xf8, 0x9b, 0x20, 0x44, 0xa9, 0x55, 0x04, 0x19, 0xd1,
	0x80, 0x38, 0xa3, 0xc1, 0x93, 0x2e, 0xb8, 0x53, 0x11, 0x7e, 0xa1, 0x01, 0xfe, 0x9a, 0x9d, 0x55,
	0xc9, 0x7b, 0xb0, 0xd5, 0xa3, 0x18, 0xd1, 0xdd, 0x0e, 0x9a, 0x7f, 0x60, 0xd7, 0x87, 0xfa, 0x07,
	0x88, 0x11, 0x74, 0xb5, 0x0f, 0xbe, 0xee, 0xe1, 0x37, 0x8c, 0x05, 0x08, 0xa1, 0xad, 0x2f, 0xce,
	0xa9, 0xf5, 0xa8, 0x73, 0x96, 0xf4, 0x44, 0x56, 0x19, 0x10, 0x17, 0xed, 0x13, 0xe5, 0x73, 0x6e,
	0x0f, 0x46, 0xe1, 0x56, 0x5c, 0x92, 0xd9, 0x8a, 0xfc, 0x33, 0xd7, 0x2a, 0x34, 0x10, 0xc5, 0x98,
	0xec, 0x4e, 0x7d, 0xf9, 0xf9, 0xe7, 0xfb, 0x06, 0x63, 0x9d, 0xd6, 0x8b, 0xca, 0x99, 0x12, 0xb7,
	0xb5, 0x32, 0xa6, 0xd6, 0xba, 0xdb, 0xaa, 0xdb, 0xbc, 0x36, 0x0f, 0xea, 0xf1, 0x36, 0x80, 0xdf,
	0x61, 0x05, 0x25, 0xd8, 0x88, 0x11, 0x21, 0x94, 0xcf, 0x2c, 0xe0, 0x7a, 0x48, 0x4b, 0xf7, 0xf1,
	0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x83, 0x83, 0x19, 0x67, 0xa2, 0x02, 0x00, 0x00,
}
