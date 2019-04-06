// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

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

type Services int32

const (
	Services_USER         Services = 0
	Services_MOUNTAIN     Services = 1
	Services_KUDAKI_EVENT Services = 2
	Services_RENTAL       Services = 3
)

var Services_name = map[int32]string{
	0: "USER",
	1: "MOUNTAIN",
	2: "KUDAKI_EVENT",
	3: "RENTAL",
}

var Services_value = map[string]int32{
	"USER":         0,
	"MOUNTAIN":     1,
	"KUDAKI_EVENT": 2,
	"RENTAL":       3,
}

func (x Services) String() string {
	return proto.EnumName(Services_name, int32(x))
}

func (Services) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func init() {
	proto.RegisterEnum("entities.Services", Services_name, Services_value)
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x48, 0xcd, 0x2b, 0xc9, 0x2c,
	0xc9, 0x4c, 0x2d, 0xd6, 0x72, 0xe0, 0xe2, 0x08, 0x86, 0x48, 0x15, 0x0b, 0x71, 0x70, 0xb1, 0x84,
	0x06, 0xbb, 0x06, 0x09, 0x30, 0x08, 0xf1, 0x70, 0x71, 0xf8, 0xfa, 0x87, 0xfa, 0x85, 0x38, 0x7a,
	0xfa, 0x09, 0x30, 0x0a, 0x09, 0x70, 0xf1, 0x78, 0x87, 0xba, 0x38, 0x7a, 0x7b, 0xc6, 0xbb, 0x86,
	0xb9, 0xfa, 0x85, 0x08, 0x30, 0x09, 0x71, 0x71, 0xb1, 0x05, 0xb9, 0xfa, 0x85, 0x38, 0xfa, 0x08,
	0x30, 0x3b, 0xa9, 0x45, 0xa9, 0xa4, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea,
	0x67, 0xe6, 0x64, 0x24, 0xe6, 0xe6, 0x66, 0xa4, 0xa4, 0xe8, 0x67, 0x97, 0xa6, 0x24, 0x66, 0x67,
	0xea, 0xc2, 0x6c, 0x4a, 0x62, 0x03, 0x5b, 0x6d, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xa2, 0xda,
	0x89, 0xc5, 0x8b, 0x00, 0x00, 0x00,
}
