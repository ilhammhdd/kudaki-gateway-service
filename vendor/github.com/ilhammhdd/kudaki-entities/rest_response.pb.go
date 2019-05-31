// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rest_response.proto

package kudaki_entities

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

type Error struct {
	Details              []string `protobuf:"bytes,1,rep,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_343c0b34fe72ecf7, []int{0}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetDetails() []string {
	if m != nil {
		return m.Details
	}
	return nil
}

type Response struct {
	Success              bool              `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Errors               map[string]*Error `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Data                 []byte            `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_343c0b34fe72ecf7, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Response) GetErrors() map[string]*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *Response) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Error)(nil), "entities.Error")
	proto.RegisterType((*Response)(nil), "entities.Response")
	proto.RegisterMapType((map[string]*Error)(nil), "entities.Response.ErrorsEntry")
}

func init() { proto.RegisterFile("rest_response.proto", fileDescriptor_343c0b34fe72ecf7) }

var fileDescriptor_343c0b34fe72ecf7 = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x4d, 0x4b, 0xc4, 0x30,
	0x10, 0x40, 0x49, 0xeb, 0xae, 0xdd, 0x54, 0x50, 0xe2, 0x25, 0x78, 0x90, 0xba, 0xa8, 0xe4, 0x62,
	0x16, 0x56, 0x10, 0xf1, 0x28, 0xec, 0xc5, 0x63, 0x8e, 0x5e, 0x24, 0xdb, 0x0c, 0x36, 0xf4, 0x23,
	0x4b, 0x26, 0x15, 0xf6, 0xbf, 0xf9, 0xe3, 0xa4, 0xb1, 0x41, 0x6f, 0x33, 0xe4, 0xbd, 0x3c, 0x18,
	0x7a, 0xe9, 0x01, 0xc3, 0x87, 0x07, 0x3c, 0xb8, 0x01, 0x41, 0x1e, 0xbc, 0x0b, 0x8e, 0x15, 0x30,
	0x04, 0x1b, 0x2c, 0xe0, 0xfa, 0x86, 0x2e, 0x76, 0xde, 0x3b, 0xcf, 0x38, 0x3d, 0x35, 0x10, 0xb4,
	0xed, 0x90, 0x93, 0x2a, 0x17, 0x2b, 0x95, 0xd6, 0xf5, 0x37, 0xa1, 0x85, 0x9a, 0xfd, 0x09, 0xc3,
	0xb1, 0xae, 0x01, 0x27, 0x8c, 0x88, 0x42, 0xa5, 0x95, 0x3d, 0xd1, 0x25, 0x4c, 0x3f, 0x21, 0xcf,
	0xaa, 0x5c, 0x94, 0xdb, 0x6b, 0x99, 0x22, 0x32, 0xd9, 0x32, 0xa6, 0x70, 0x37, 0x04, 0x7f, 0x54,
	0x33, 0xcd, 0x18, 0x3d, 0x31, 0x3a, 0x68, 0x9e, 0x57, 0x44, 0x9c, 0xa9, 0x38, 0x5f, 0xbd, 0xd1,
	0xf2, 0x1f, 0xca, 0x2e, 0x68, 0xde, 0xc2, 0x31, 0x06, 0x57, 0x6a, 0x1a, 0xd9, 0x1d, 0x5d, 0x7c,
	0xe9, 0x6e, 0x04, 0x9e, 0x55, 0x44, 0x94, 0xdb, 0xf3, 0xbf, 0x56, 0xf4, 0xd4, 0xef, 0xeb, 0x4b,
	0xf6, 0x4c, 0x5e, 0xef, 0xdf, 0x6f, 0x3f, 0x6d, 0x68, 0xc6, 0xbd, 0xac, 0x5d, 0xbf, 0xb1, 0x5d,
	0xa3, 0xfb, 0xbe, 0x31, 0x66, 0xd3, 0x8e, 0x46, 0xb7, 0xf6, 0x21, 0x89, 0xfb, 0x65, 0x3c, 0xcd,
	0xe3, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x96, 0xe9, 0x07, 0x13, 0x31, 0x01, 0x00, 0x00,
}