// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.5
// source: message.proto

package proto_go

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Hello_Say_FullMethodName = "/Hello/Say"
)

// HelloClient is the client API for Hello service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloClient interface {
	Say(ctx context.Context, in *SayRequest, opts ...grpc.CallOption) (*SayResponse, error)
}

type helloClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloClient(cc grpc.ClientConnInterface) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) Say(ctx context.Context, in *SayRequest, opts ...grpc.CallOption) (*SayResponse, error) {
	out := new(SayResponse)
	err := c.cc.Invoke(ctx, Hello_Say_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServer is the server API for Hello service.
// All implementations must embed UnimplementedHelloServer
// for forward compatibility
type HelloServer interface {
	Say(context.Context, *SayRequest) (*SayResponse, error)
	mustEmbedUnimplementedHelloServer()
}

// UnimplementedHelloServer must be embedded to have forward compatible implementations.
type UnimplementedHelloServer struct {
}

func (UnimplementedHelloServer) Say(context.Context, *SayRequest) (*SayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Say not implemented")
}
func (UnimplementedHelloServer) mustEmbedUnimplementedHelloServer() {}

// UnsafeHelloServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloServer will
// result in compilation errors.
type UnsafeHelloServer interface {
	mustEmbedUnimplementedHelloServer()
}

func RegisterHelloServer(s grpc.ServiceRegistrar, srv HelloServer) {
	s.RegisterService(&Hello_ServiceDesc, srv)
}

func _Hello_Say_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).Say(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hello_Say_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).Say(ctx, req.(*SayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Hello_ServiceDesc is the grpc.ServiceDesc for Hello service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Hello_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Say",
			Handler:    _Hello_Say_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}

const (
	Gogrpc_UserLogincode_FullMethodName = "/Gogrpc/UserLogincode"
	Gogrpc_UserLoginauth_FullMethodName = "/Gogrpc/UserLoginauth"
	Gogrpc_CommMsg_FullMethodName       = "/Gogrpc/CommMsg"
)

// GogrpcClient is the client API for Gogrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GogrpcClient interface {
	UserLogincode(ctx context.Context, in *C2S_UserLogincode, opts ...grpc.CallOption) (*S2C_UserLogincode, error)
	UserLoginauth(ctx context.Context, in *C2S_UserLoginauth, opts ...grpc.CallOption) (*S2C_UserLoginauth, error)
	CommMsg(ctx context.Context, in *C2S_Map, opts ...grpc.CallOption) (*S2C_Map, error)
}

type gogrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewGogrpcClient(cc grpc.ClientConnInterface) GogrpcClient {
	return &gogrpcClient{cc}
}

func (c *gogrpcClient) UserLogincode(ctx context.Context, in *C2S_UserLogincode, opts ...grpc.CallOption) (*S2C_UserLogincode, error) {
	out := new(S2C_UserLogincode)
	err := c.cc.Invoke(ctx, Gogrpc_UserLogincode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gogrpcClient) UserLoginauth(ctx context.Context, in *C2S_UserLoginauth, opts ...grpc.CallOption) (*S2C_UserLoginauth, error) {
	out := new(S2C_UserLoginauth)
	err := c.cc.Invoke(ctx, Gogrpc_UserLoginauth_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gogrpcClient) CommMsg(ctx context.Context, in *C2S_Map, opts ...grpc.CallOption) (*S2C_Map, error) {
	out := new(S2C_Map)
	err := c.cc.Invoke(ctx, Gogrpc_CommMsg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GogrpcServer is the server API for Gogrpc service.
// All implementations must embed UnimplementedGogrpcServer
// for forward compatibility
type GogrpcServer interface {
	UserLogincode(context.Context, *C2S_UserLogincode) (*S2C_UserLogincode, error)
	UserLoginauth(context.Context, *C2S_UserLoginauth) (*S2C_UserLoginauth, error)
	CommMsg(context.Context, *C2S_Map) (*S2C_Map, error)
	mustEmbedUnimplementedGogrpcServer()
}

// UnimplementedGogrpcServer must be embedded to have forward compatible implementations.
type UnimplementedGogrpcServer struct {
}

func (UnimplementedGogrpcServer) UserLogincode(context.Context, *C2S_UserLogincode) (*S2C_UserLogincode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogincode not implemented")
}
func (UnimplementedGogrpcServer) UserLoginauth(context.Context, *C2S_UserLoginauth) (*S2C_UserLoginauth, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLoginauth not implemented")
}
func (UnimplementedGogrpcServer) CommMsg(context.Context, *C2S_Map) (*S2C_Map, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommMsg not implemented")
}
func (UnimplementedGogrpcServer) mustEmbedUnimplementedGogrpcServer() {}

// UnsafeGogrpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GogrpcServer will
// result in compilation errors.
type UnsafeGogrpcServer interface {
	mustEmbedUnimplementedGogrpcServer()
}

func RegisterGogrpcServer(s grpc.ServiceRegistrar, srv GogrpcServer) {
	s.RegisterService(&Gogrpc_ServiceDesc, srv)
}

func _Gogrpc_UserLogincode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(C2S_UserLogincode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GogrpcServer).UserLogincode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gogrpc_UserLogincode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GogrpcServer).UserLogincode(ctx, req.(*C2S_UserLogincode))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gogrpc_UserLoginauth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(C2S_UserLoginauth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GogrpcServer).UserLoginauth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gogrpc_UserLoginauth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GogrpcServer).UserLoginauth(ctx, req.(*C2S_UserLoginauth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gogrpc_CommMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(C2S_Map)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GogrpcServer).CommMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gogrpc_CommMsg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GogrpcServer).CommMsg(ctx, req.(*C2S_Map))
	}
	return interceptor(ctx, in, info, handler)
}

// Gogrpc_ServiceDesc is the grpc.ServiceDesc for Gogrpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gogrpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Gogrpc",
	HandlerType: (*GogrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserLogincode",
			Handler:    _Gogrpc_UserLogincode_Handler,
		},
		{
			MethodName: "UserLoginauth",
			Handler:    _Gogrpc_UserLoginauth_Handler,
		},
		{
			MethodName: "CommMsg",
			Handler:    _Gogrpc_CommMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}
