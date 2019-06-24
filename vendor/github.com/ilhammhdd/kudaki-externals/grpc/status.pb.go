// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/status.proto

package grpc

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

type Status struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Errs                 []string `protobuf:"bytes,2,rep,name=errs,proto3" json:"errs,omitempty"`
	LastInsertedId       int64    `protobuf:"varint,3,opt,name=last_inserted_id,json=lastInsertedId,proto3" json:"last_inserted_id,omitempty"`
	RowsAffected         int64    `protobuf:"varint,4,opt,name=rows_affected,json=rowsAffected,proto3" json:"rows_affected,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_1053741527bacce7, []int{0}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *Status) GetErrs() []string {
	if m != nil {
		return m.Errs
	}
	return nil
}

func (m *Status) GetLastInsertedId() int64 {
	if m != nil {
		return m.LastInsertedId
	}
	return 0
}

func (m *Status) GetRowsAffected() int64 {
	if m != nil {
		return m.RowsAffected
	}
	return 0
}

func init() {
	proto.RegisterType((*Status)(nil), "rpc.Status")
}

func init() { proto.RegisterFile("grpc/status.proto", fileDescriptor_1053741527bacce7) }

var fileDescriptor_1053741527bacce7 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x2c, 0xce, 0x41, 0x4b, 0x80, 0x30,
	0x18, 0xc6, 0x71, 0x74, 0x22, 0x35, 0x4a, 0x6a, 0xa7, 0x1d, 0xa5, 0x2e, 0x23, 0xca, 0x1d, 0xfa,
	0x04, 0x75, 0xf3, 0x6a, 0xb7, 0x2e, 0x32, 0xb7, 0xa9, 0x63, 0xea, 0xe4, 0xdd, 0x2b, 0x05, 0x7d,
	0xf9, 0x70, 0x78, 0x7b, 0xf8, 0xf1, 0x1c, 0xfe, 0xf4, 0x71, 0x82, 0x5d, 0xcb, 0x88, 0x0a, 0x8f,
	0xd8, 0xec, 0x10, 0x30, 0x30, 0x02, 0xbb, 0x7e, 0xfa, 0xa3, 0xe5, 0x57, 0x42, 0x56, 0xd1, 0x3c,
	0x78, 0x9e, 0xd5, 0x99, 0xb8, 0xe9, 0xf2, 0xe0, 0x19, 0xa3, 0x85, 0x05, 0x88, 0x3c, 0xaf, 0x89,
	0xb8, 0xed, 0xd2, 0x66, 0x82, 0x3e, 0x2c, 0x2a, 0x62, 0xef, 0xb6, 0x68, 0x01, 0xad, 0xe9, 0x9d,
	0xe1, 0xa4, 0xce, 0x04, 0xe9, 0xaa, 0xd3, 0xdb, 0x8b, 0x5b, 0xc3, 0x9e, 0xe9, 0x3d, 0x84, 0x9f,
	0xd8, 0xab, 0x71, 0xb4, 0x1a, 0xad, 0xe1, 0x45, 0xba, 0xdd, 0x9d, 0xf8, 0x71, 0xd9, 0xe7, 0xeb,
	0xf7, 0xcb, 0xe4, 0x70, 0x3e, 0x86, 0x46, 0x87, 0x55, 0xba, 0x65, 0x56, 0xeb, 0x3a, 0x1b, 0x23,
	0xfd, 0x61, 0x94, 0x77, 0x6f, 0xf6, 0x17, 0x2d, 0x6c, 0x6a, 0x89, 0xf2, 0x8c, 0x1f, 0xca, 0x94,
	0xfd, 0xfe, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x19, 0x6e, 0x15, 0xcb, 0x00, 0x00, 0x00,
}
