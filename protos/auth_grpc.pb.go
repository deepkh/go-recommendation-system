// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.0
// source: protos/auth.proto

package protos

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

// AuthServClient is the client API for AuthServ service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServClient interface {
	// use plain username, password to accquire AuthedUser
	Auth(ctx context.Context, in *AuthReq, opts ...grpc.CallOption) (*AuthRep, error)
	// use token to accquire AuthedUser
	CheckToken(ctx context.Context, in *CheckTokenReq, opts ...grpc.CallOption) (*CheckTokenRep, error)
}

type authServClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServClient(cc grpc.ClientConnInterface) AuthServClient {
	return &authServClient{cc}
}

func (c *authServClient) Auth(ctx context.Context, in *AuthReq, opts ...grpc.CallOption) (*AuthRep, error) {
	out := new(AuthRep)
	err := c.cc.Invoke(ctx, "/AuthServ/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServClient) CheckToken(ctx context.Context, in *CheckTokenReq, opts ...grpc.CallOption) (*CheckTokenRep, error) {
	out := new(CheckTokenRep)
	err := c.cc.Invoke(ctx, "/AuthServ/CheckToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServServer is the server API for AuthServ service.
// All implementations must embed UnimplementedAuthServServer
// for forward compatibility
type AuthServServer interface {
	// use plain username, password to accquire AuthedUser
	Auth(context.Context, *AuthReq) (*AuthRep, error)
	// use token to accquire AuthedUser
	CheckToken(context.Context, *CheckTokenReq) (*CheckTokenRep, error)
	mustEmbedUnimplementedAuthServServer()
}

// UnimplementedAuthServServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServServer struct {
}

func (UnimplementedAuthServServer) Auth(context.Context, *AuthReq) (*AuthRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (UnimplementedAuthServServer) CheckToken(context.Context, *CheckTokenReq) (*CheckTokenRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckToken not implemented")
}
func (UnimplementedAuthServServer) mustEmbedUnimplementedAuthServServer() {}

// UnsafeAuthServServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServServer will
// result in compilation errors.
type UnsafeAuthServServer interface {
	mustEmbedUnimplementedAuthServServer()
}

func RegisterAuthServServer(s grpc.ServiceRegistrar, srv AuthServServer) {
	s.RegisterService(&AuthServ_ServiceDesc, srv)
}

func _AuthServ_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthServ/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServServer).Auth(ctx, req.(*AuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthServ_CheckToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServServer).CheckToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthServ/CheckToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServServer).CheckToken(ctx, req.(*CheckTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthServ_ServiceDesc is the grpc.ServiceDesc for AuthServ service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthServ_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AuthServ",
	HandlerType: (*AuthServServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Auth",
			Handler:    _AuthServ_Auth_Handler,
		},
		{
			MethodName: "CheckToken",
			Handler:    _AuthServ_CheckToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/auth.proto",
}
