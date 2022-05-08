// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: internal/server/turbo.proto

package server

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

// TurboClient is the client API for Turbo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TurboClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	GetGlobalHash(ctx context.Context, in *GlobalHashRequest, opts ...grpc.CallOption) (*GlobalHashReply, error)
}

type turboClient struct {
	cc grpc.ClientConnInterface
}

func NewTurboClient(cc grpc.ClientConnInterface) TurboClient {
	return &turboClient{cc}
}

func (c *turboClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/server.Turbo/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *turboClient) GetGlobalHash(ctx context.Context, in *GlobalHashRequest, opts ...grpc.CallOption) (*GlobalHashReply, error) {
	out := new(GlobalHashReply)
	err := c.cc.Invoke(ctx, "/server.Turbo/GetGlobalHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TurboServer is the server API for Turbo service.
// All implementations must embed UnimplementedTurboServer
// for forward compatibility
type TurboServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
	GetGlobalHash(context.Context, *GlobalHashRequest) (*GlobalHashReply, error)
	mustEmbedUnimplementedTurboServer()
}

// UnimplementedTurboServer must be embedded to have forward compatible implementations.
type UnimplementedTurboServer struct {
}

func (UnimplementedTurboServer) Ping(context.Context, *PingRequest) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedTurboServer) GetGlobalHash(context.Context, *GlobalHashRequest) (*GlobalHashReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGlobalHash not implemented")
}
func (UnimplementedTurboServer) mustEmbedUnimplementedTurboServer() {}

// UnsafeTurboServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TurboServer will
// result in compilation errors.
type UnsafeTurboServer interface {
	mustEmbedUnimplementedTurboServer()
}

func RegisterTurboServer(s grpc.ServiceRegistrar, srv TurboServer) {
	s.RegisterService(&Turbo_ServiceDesc, srv)
}

func _Turbo_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TurboServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Turbo/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TurboServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Turbo_GetGlobalHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GlobalHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TurboServer).GetGlobalHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Turbo/GetGlobalHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TurboServer).GetGlobalHash(ctx, req.(*GlobalHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Turbo_ServiceDesc is the grpc.ServiceDesc for Turbo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Turbo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server.Turbo",
	HandlerType: (*TurboServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Turbo_Ping_Handler,
		},
		{
			MethodName: "GetGlobalHash",
			Handler:    _Turbo_GetGlobalHash_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/server/turbo.proto",
}