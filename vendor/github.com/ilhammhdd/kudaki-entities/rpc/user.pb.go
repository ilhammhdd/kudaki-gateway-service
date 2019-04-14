// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/user.proto

package rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	events "github.com/ilhammhdd/kudaki-entities/events"
	grpc "google.golang.org/grpc"
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

func init() { proto.RegisterFile("rpc/user.proto", fileDescriptor_27d7a9c2ccec3127) }

var fileDescriptor_27d7a9c2ccec3127 = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x0f, 0xd5, 0x1e, 0x16, 0x54, 0x3a, 0x50, 0x91, 0x80, 0xa0, 0x9e, 0x44, 0x30, 0x01,
	0x3d, 0x79, 0xb4, 0xa0, 0xa7, 0x0a, 0x92, 0xa2, 0x07, 0x6f, 0x69, 0x76, 0x4c, 0x86, 0x36, 0xbb,
	0xeb, 0xee, 0xac, 0xa2, 0xff, 0xc6, 0x7f, 0x2a, 0x49, 0xb7, 0x6e, 0xaa, 0x3d, 0xce, 0x7b, 0x6f,
	0xbf, 0x79, 0xec, 0x88, 0x7d, 0x6b, 0xca, 0xcc, 0x3b, 0xb4, 0xa9, 0xb1, 0x9a, 0x35, 0x0c, 0xac,
	0x29, 0x93, 0x11, 0xbe, 0xa3, 0x62, 0xd7, 0xd3, 0xaf, 0xbe, 0x07, 0x62, 0xe7, 0xc9, 0xa1, 0x85,
	0x89, 0x18, 0xce, 0xa8, 0x52, 0xde, 0xc0, 0x61, 0xda, 0xc5, 0xd2, 0xd5, 0x98, 0xe3, 0x9b, 0x47,
	0xc7, 0x28, 0x93, 0x93, 0xa0, 0xb7, 0xf1, 0x67, 0xb4, 0xf4, 0x4a, 0x65, 0xc1, 0xa4, 0xd5, 0x5d,
	0x53, 0xd0, 0x72, 0x86, 0x8a, 0xe1, 0x46, 0x88, 0xce, 0xf8, 0xec, 0x88, 0x49, 0xc8, 0x47, 0x29,
	0xb2, 0x0e, 0x7a, 0x3b, 0x50, 0x7a, 0x03, 0x99, 0xd8, 0x9d, 0xea, 0x8a, 0x14, 0x8c, 0x83, 0xd3,
	0x4d, 0xff, 0x1f, 0x4c, 0x75, 0x55, 0xa1, 0x24, 0x05, 0xb9, 0x80, 0x16, 0x79, 0xeb, 0xb9, 0x46,
	0xc5, 0xa1, 0x0a, 0x9c, 0xf5, 0x3a, 0x6e, 0x5a, 0x11, 0x75, 0xb4, 0x3d, 0x83, 0x12, 0xee, 0xc5,
	0x5e, 0x8e, 0x0e, 0xf9, 0xb1, 0x70, 0xee, 0x43, 0x5b, 0x09, 0xc7, 0x21, 0xba, 0xa1, 0x46, 0xd2,
	0xfa, 0xa7, 0xa2, 0xe3, 0xb0, 0xe5, 0x3c, 0x88, 0xd1, 0x1a, 0xae, 0x2d, 0x7d, 0xad, 0xaa, 0x9d,
	0xfe, 0x59, 0xfb, 0xeb, 0x44, 0xde, 0x78, 0x4b, 0x04, 0xe5, 0xe4, 0xe2, 0xe5, 0xbc, 0x22, 0xae,
	0xfd, 0x3c, 0x2d, 0x75, 0x93, 0xd1, 0xb2, 0x2e, 0x9a, 0xa6, 0x96, 0x32, 0x5b, 0x78, 0x59, 0x2c,
	0xe8, 0xb2, 0xed, 0xcf, 0x84, 0x2e, 0xb3, 0xa6, 0x9c, 0x0f, 0xbb, 0xb3, 0x5e, 0xff, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x7e, 0xfd, 0x07, 0xd5, 0x00, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	Signup(ctx context.Context, in *events.SignupRequested, opts ...grpc.CallOption) (*events.UserVerificationEmailSent, error)
	VerifyUser(ctx context.Context, in *events.VerifyUserRequested, opts ...grpc.CallOption) (*events.Signedup, error)
	Login(ctx context.Context, in *events.LoginRequested, opts ...grpc.CallOption) (*events.Loggedin, error)
	UserAuthentication(ctx context.Context, in *events.UserAuthenticationRequested, opts ...grpc.CallOption) (*events.UserAuthenticated, error)
	ResetPassword(ctx context.Context, in *events.ResetPasswordRequested, opts ...grpc.CallOption) (*events.PasswordReseted, error)
	UserAuthorization(ctx context.Context, in *events.UserAuthorizationRequested, opts ...grpc.CallOption) (*events.UserAuthorized, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) Signup(ctx context.Context, in *events.SignupRequested, opts ...grpc.CallOption) (*events.UserVerificationEmailSent, error) {
	out := new(events.UserVerificationEmailSent)
	err := c.cc.Invoke(ctx, "/rpc.User/Signup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) VerifyUser(ctx context.Context, in *events.VerifyUserRequested, opts ...grpc.CallOption) (*events.Signedup, error) {
	out := new(events.Signedup)
	err := c.cc.Invoke(ctx, "/rpc.User/VerifyUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Login(ctx context.Context, in *events.LoginRequested, opts ...grpc.CallOption) (*events.Loggedin, error) {
	out := new(events.Loggedin)
	err := c.cc.Invoke(ctx, "/rpc.User/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserAuthentication(ctx context.Context, in *events.UserAuthenticationRequested, opts ...grpc.CallOption) (*events.UserAuthenticated, error) {
	out := new(events.UserAuthenticated)
	err := c.cc.Invoke(ctx, "/rpc.User/UserAuthentication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ResetPassword(ctx context.Context, in *events.ResetPasswordRequested, opts ...grpc.CallOption) (*events.PasswordReseted, error) {
	out := new(events.PasswordReseted)
	err := c.cc.Invoke(ctx, "/rpc.User/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserAuthorization(ctx context.Context, in *events.UserAuthorizationRequested, opts ...grpc.CallOption) (*events.UserAuthorized, error) {
	out := new(events.UserAuthorized)
	err := c.cc.Invoke(ctx, "/rpc.User/UserAuthorization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	Signup(context.Context, *events.SignupRequested) (*events.UserVerificationEmailSent, error)
	VerifyUser(context.Context, *events.VerifyUserRequested) (*events.Signedup, error)
	Login(context.Context, *events.LoginRequested) (*events.Loggedin, error)
	UserAuthentication(context.Context, *events.UserAuthenticationRequested) (*events.UserAuthenticated, error)
	ResetPassword(context.Context, *events.ResetPasswordRequested) (*events.PasswordReseted, error)
	UserAuthorization(context.Context, *events.UserAuthorizationRequested) (*events.UserAuthorized, error)
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_Signup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(events.SignupRequested)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Signup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.User/Signup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Signup(ctx, req.(*events.SignupRequested))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_VerifyUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(events.VerifyUserRequested)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).VerifyUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.User/VerifyUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).VerifyUser(ctx, req.(*events.VerifyUserRequested))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(events.LoginRequested)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.User/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Login(ctx, req.(*events.LoginRequested))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserAuthentication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(events.UserAuthenticationRequested)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserAuthentication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.User/UserAuthentication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserAuthentication(ctx, req.(*events.UserAuthenticationRequested))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(events.ResetPasswordRequested)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.User/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ResetPassword(ctx, req.(*events.ResetPasswordRequested))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserAuthorization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(events.UserAuthorizationRequested)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserAuthorization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.User/UserAuthorization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserAuthorization(ctx, req.(*events.UserAuthorizationRequested))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Signup",
			Handler:    _User_Signup_Handler,
		},
		{
			MethodName: "VerifyUser",
			Handler:    _User_VerifyUser_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _User_Login_Handler,
		},
		{
			MethodName: "UserAuthentication",
			Handler:    _User_UserAuthentication_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _User_ResetPassword_Handler,
		},
		{
			MethodName: "UserAuthorization",
			Handler:    _User_UserAuthorization_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/user.proto",
}
