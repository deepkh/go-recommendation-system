// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.0
// source: protos/user.proto

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

// RegServClient is the client API for RegServ service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegServClient interface {
	// use plain email, password to accquire AuthedUser
	Reg(ctx context.Context, in *RegReq, opts ...grpc.CallOption) (*RegRep, error)
	Confirm(ctx context.Context, in *ConfirmReq, opts ...grpc.CallOption) (*ConfirmRep, error)
}

type regServClient struct {
	cc grpc.ClientConnInterface
}

func NewRegServClient(cc grpc.ClientConnInterface) RegServClient {
	return &regServClient{cc}
}

func (c *regServClient) Reg(ctx context.Context, in *RegReq, opts ...grpc.CallOption) (*RegRep, error) {
	out := new(RegRep)
	err := c.cc.Invoke(ctx, "/RegServ/Reg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regServClient) Confirm(ctx context.Context, in *ConfirmReq, opts ...grpc.CallOption) (*ConfirmRep, error) {
	out := new(ConfirmRep)
	err := c.cc.Invoke(ctx, "/RegServ/Confirm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegServServer is the server API for RegServ service.
// All implementations must embed UnimplementedRegServServer
// for forward compatibility
type RegServServer interface {
	// use plain email, password to accquire AuthedUser
	Reg(context.Context, *RegReq) (*RegRep, error)
	Confirm(context.Context, *ConfirmReq) (*ConfirmRep, error)
	mustEmbedUnimplementedRegServServer()
}

// UnimplementedRegServServer must be embedded to have forward compatible implementations.
type UnimplementedRegServServer struct {
}

func (UnimplementedRegServServer) Reg(context.Context, *RegReq) (*RegRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reg not implemented")
}
func (UnimplementedRegServServer) Confirm(context.Context, *ConfirmReq) (*ConfirmRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Confirm not implemented")
}
func (UnimplementedRegServServer) mustEmbedUnimplementedRegServServer() {}

// UnsafeRegServServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegServServer will
// result in compilation errors.
type UnsafeRegServServer interface {
	mustEmbedUnimplementedRegServServer()
}

func RegisterRegServServer(s grpc.ServiceRegistrar, srv RegServServer) {
	s.RegisterService(&RegServ_ServiceDesc, srv)
}

func _RegServ_Reg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegServServer).Reg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RegServ/Reg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegServServer).Reg(ctx, req.(*RegReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegServ_Confirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegServServer).Confirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RegServ/Confirm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegServServer).Confirm(ctx, req.(*ConfirmReq))
	}
	return interceptor(ctx, in, info, handler)
}

// RegServ_ServiceDesc is the grpc.ServiceDesc for RegServ service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RegServ_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RegServ",
	HandlerType: (*RegServServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Reg",
			Handler:    _RegServ_Reg_Handler,
		},
		{
			MethodName: "Confirm",
			Handler:    _RegServ_Confirm_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/user.proto",
}
