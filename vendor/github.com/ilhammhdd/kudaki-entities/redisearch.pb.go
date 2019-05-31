// Code generated by protoc-gen-go. DO NOT EDIT.
// source: redisearch.proto

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

type ClientName int32

const (
	ClientName_CARTS      ClientName = 0
	ClientName_CART_ITEMS ClientName = 1
	ClientName_CHECKOUTS  ClientName = 2
)

var ClientName_name = map[int32]string{
	0: "CARTS",
	1: "CART_ITEMS",
	2: "CHECKOUTS",
}

var ClientName_value = map[string]int32{
	"CARTS":      0,
	"CART_ITEMS": 1,
	"CHECKOUTS":  2,
}

func (x ClientName) String() string {
	return proto.EnumName(ClientName_name, int32(x))
}

func (ClientName) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b3be1ea04b809d96, []int{0}
}

func init() {
	proto.RegisterEnum("entities.ClientName", ClientName_name, ClientName_value)
}

func init() { proto.RegisterFile("redisearch.proto", fileDescriptor_b3be1ea04b809d96) }

var fileDescriptor_b3be1ea04b809d96 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x4a, 0x4d, 0xc9,
	0x2c, 0x4e, 0x4d, 0x2c, 0x4a, 0xce, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x48, 0xcd,
	0x2b, 0xc9, 0x2c, 0xc9, 0x4c, 0x2d, 0xd6, 0x32, 0xe3, 0xe2, 0x72, 0xce, 0xc9, 0x4c, 0xcd, 0x2b,
	0xf1, 0x4b, 0xcc, 0x4d, 0x15, 0xe2, 0xe4, 0x62, 0x75, 0x76, 0x0c, 0x0a, 0x09, 0x16, 0x60, 0x10,
	0xe2, 0xe3, 0xe2, 0x02, 0x31, 0xe3, 0x3d, 0x43, 0x5c, 0x7d, 0x83, 0x05, 0x18, 0x85, 0x78, 0xb9,
	0x38, 0x9d, 0x3d, 0x5c, 0x9d, 0xbd, 0xfd, 0x43, 0x43, 0x82, 0x05, 0x98, 0x9c, 0xd4, 0xa2, 0x54,
	0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x33, 0x73, 0x32, 0x12, 0x73,
	0x73, 0x33, 0x52, 0x52, 0xf4, 0xb3, 0x4b, 0x53, 0x12, 0xb3, 0x33, 0x75, 0x61, 0xe6, 0x27, 0xb1,
	0x81, 0x2d, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x24, 0x26, 0xce, 0x84, 0x00, 0x00,
	0x00,
}
