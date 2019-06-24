// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/store.proto

package grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	store "github.com/ilhammhdd/kudaki-entities/store"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

func init() { proto.RegisterFile("grpc/store.proto", fileDescriptor_f000a2b738285d05) }

var fileDescriptor_f000a2b738285d05 = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xbf, 0x4a, 0x04, 0x31,
	0x10, 0x87, 0x39, 0x04, 0xff, 0x8c, 0x08, 0x6b, 0x10, 0x8b, 0x7d, 0x04, 0x39, 0x13, 0x38, 0x1b,
	0xc1, 0xea, 0xc4, 0xc6, 0xd6, 0xc5, 0xc6, 0x2e, 0x97, 0x19, 0x6f, 0xc3, 0x6d, 0x36, 0x21, 0x99,
	0x05, 0x9f, 0xc5, 0x97, 0xf1, 0xd5, 0x24, 0xc9, 0xa1, 0x68, 0xa5, 0x5b, 0xe6, 0x9b, 0xfc, 0xbe,
	0x99, 0x61, 0xa0, 0xd9, 0xc6, 0x60, 0x54, 0x62, 0x1f, 0x49, 0x86, 0xe8, 0xd9, 0x8b, 0x83, 0x18,
	0x4c, 0x7b, 0x59, 0x48, 0xe5, 0xaf, 0xd1, 0x8f, 0x5c, 0x8b, 0x6d, 0x53, 0xb9, 0x65, 0x72, 0x7b,
	0x72, 0xbe, 0x17, 0x68, 0x9e, 0x52, 0x45, 0xab, 0x8f, 0x05, 0x9c, 0x74, 0xf9, 0xdf, 0x13, 0x05,
	0x2f, 0x6e, 0xe1, 0x6c, 0x8d, 0xd8, 0x7d, 0x99, 0x44, 0x2b, 0x69, 0x64, 0xcb, 0x96, 0x92, 0xac,
	0x7d, 0xbf, 0x6b, 0xed, 0xa9, 0x8c, 0xc1, 0xc8, 0xae, 0xd8, 0xc4, 0x1d, 0x34, 0xcf, 0x01, 0x35,
	0xd3, 0xcc, 0xf0, 0x03, 0x0d, 0x34, 0x2b, 0xbc, 0x7a, 0x5f, 0xc0, 0xf1, 0x23, 0x93, 0x2b, 0x0b,
	0x2c, 0xe1, 0x68, 0x8d, 0x98, 0x9f, 0xe2, 0xe2, 0xb7, 0x20, 0xd3, 0x9f, 0x7d, 0x15, 0x40, 0x1d,
	0xfa, 0x1f, 0x81, 0x3a, 0xe8, 0x1f, 0x03, 0xf7, 0xcb, 0x97, 0xab, 0xad, 0xe5, 0x7e, 0xda, 0x48,
	0xe3, 0x9d, 0xb2, 0x43, 0xaf, 0x9d, 0xeb, 0x11, 0xd5, 0x6e, 0x42, 0xbd, 0xb3, 0xd7, 0xf4, 0xc6,
	0x14, 0x47, 0x3d, 0x24, 0x95, 0x2f, 0xb3, 0x39, 0x2c, 0x37, 0xb9, 0xf9, 0x0c, 0x00, 0x00, 0xff,
	0xff, 0xd9, 0xeb, 0xf2, 0x74, 0xe9, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StoreRepoClient is the client API for StoreRepo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StoreRepoClient interface {
	AddStorefront(ctx context.Context, in *store.Storefront, opts ...grpc.CallOption) (*Status, error)
	UpdateStorefront(ctx context.Context, in *store.Storefront, opts ...grpc.CallOption) (*Status, error)
	DeleteStorefront(ctx context.Context, in *store.Storefront, opts ...grpc.CallOption) (*Status, error)
}

type storeRepoClient struct {
	cc *grpc.ClientConn
}

func NewStoreRepoClient(cc *grpc.ClientConn) StoreRepoClient {
	return &storeRepoClient{cc}
}

func (c *storeRepoClient) AddStorefront(ctx context.Context, in *store.Storefront, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/rpc.StoreRepo/AddStorefront", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeRepoClient) UpdateStorefront(ctx context.Context, in *store.Storefront, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/rpc.StoreRepo/UpdateStorefront", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeRepoClient) DeleteStorefront(ctx context.Context, in *store.Storefront, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/rpc.StoreRepo/DeleteStorefront", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreRepoServer is the server API for StoreRepo service.
type StoreRepoServer interface {
	AddStorefront(context.Context, *store.Storefront) (*Status, error)
	UpdateStorefront(context.Context, *store.Storefront) (*Status, error)
	DeleteStorefront(context.Context, *store.Storefront) (*Status, error)
}

// UnimplementedStoreRepoServer can be embedded to have forward compatible implementations.
type UnimplementedStoreRepoServer struct {
}

func (*UnimplementedStoreRepoServer) AddStorefront(ctx context.Context, req *store.Storefront) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddStorefront not implemented")
}
func (*UnimplementedStoreRepoServer) UpdateStorefront(ctx context.Context, req *store.Storefront) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStorefront not implemented")
}
func (*UnimplementedStoreRepoServer) DeleteStorefront(ctx context.Context, req *store.Storefront) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStorefront not implemented")
}

func RegisterStoreRepoServer(s *grpc.Server, srv StoreRepoServer) {
	s.RegisterService(&_StoreRepo_serviceDesc, srv)
}

func _StoreRepo_AddStorefront_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(store.Storefront)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreRepoServer).AddStorefront(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.StoreRepo/AddStorefront",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreRepoServer).AddStorefront(ctx, req.(*store.Storefront))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreRepo_UpdateStorefront_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(store.Storefront)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreRepoServer).UpdateStorefront(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.StoreRepo/UpdateStorefront",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreRepoServer).UpdateStorefront(ctx, req.(*store.Storefront))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreRepo_DeleteStorefront_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(store.Storefront)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreRepoServer).DeleteStorefront(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.StoreRepo/DeleteStorefront",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreRepoServer).DeleteStorefront(ctx, req.(*store.Storefront))
	}
	return interceptor(ctx, in, info, handler)
}

var _StoreRepo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.StoreRepo",
	HandlerType: (*StoreRepoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddStorefront",
			Handler:    _StoreRepo_AddStorefront_Handler,
		},
		{
			MethodName: "UpdateStorefront",
			Handler:    _StoreRepo_UpdateStorefront_Handler,
		},
		{
			MethodName: "DeleteStorefront",
			Handler:    _StoreRepo_DeleteStorefront_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/store.proto",
}

// ItemRepoClient is the client API for ItemRepo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ItemRepoClient interface {
	AddItem(ctx context.Context, in *store.Item, opts ...grpc.CallOption) (*Status, error)
	UpdateItem(ctx context.Context, in *store.Item, opts ...grpc.CallOption) (*Status, error)
	DeleteItem(ctx context.Context, in *store.Item, opts ...grpc.CallOption) (*Status, error)
}

type itemRepoClient struct {
	cc *grpc.ClientConn
}

func NewItemRepoClient(cc *grpc.ClientConn) ItemRepoClient {
	return &itemRepoClient{cc}
}

func (c *itemRepoClient) AddItem(ctx context.Context, in *store.Item, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/rpc.ItemRepo/AddItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemRepoClient) UpdateItem(ctx context.Context, in *store.Item, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/rpc.ItemRepo/UpdateItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemRepoClient) DeleteItem(ctx context.Context, in *store.Item, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/rpc.ItemRepo/DeleteItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ItemRepoServer is the server API for ItemRepo service.
type ItemRepoServer interface {
	AddItem(context.Context, *store.Item) (*Status, error)
	UpdateItem(context.Context, *store.Item) (*Status, error)
	DeleteItem(context.Context, *store.Item) (*Status, error)
}

// UnimplementedItemRepoServer can be embedded to have forward compatible implementations.
type UnimplementedItemRepoServer struct {
}

func (*UnimplementedItemRepoServer) AddItem(ctx context.Context, req *store.Item) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddItem not implemented")
}
func (*UnimplementedItemRepoServer) UpdateItem(ctx context.Context, req *store.Item) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateItem not implemented")
}
func (*UnimplementedItemRepoServer) DeleteItem(ctx context.Context, req *store.Item) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteItem not implemented")
}

func RegisterItemRepoServer(s *grpc.Server, srv ItemRepoServer) {
	s.RegisterService(&_ItemRepo_serviceDesc, srv)
}

func _ItemRepo_AddItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(store.Item)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemRepoServer).AddItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.ItemRepo/AddItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemRepoServer).AddItem(ctx, req.(*store.Item))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemRepo_UpdateItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(store.Item)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemRepoServer).UpdateItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.ItemRepo/UpdateItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemRepoServer).UpdateItem(ctx, req.(*store.Item))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemRepo_DeleteItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(store.Item)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemRepoServer).DeleteItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.ItemRepo/DeleteItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemRepoServer).DeleteItem(ctx, req.(*store.Item))
	}
	return interceptor(ctx, in, info, handler)
}

var _ItemRepo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.ItemRepo",
	HandlerType: (*ItemRepoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddItem",
			Handler:    _ItemRepo_AddItem_Handler,
		},
		{
			MethodName: "UpdateItem",
			Handler:    _ItemRepo_UpdateItem_Handler,
		},
		{
			MethodName: "DeleteItem",
			Handler:    _ItemRepo_DeleteItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/store.proto",
}
