// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: user.proto

package proto

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	FindOneUser(ctx context.Context, in *FindOneRequest, opts ...grpc.CallOption) (*FindOneUserResult, error)
	FindMany(ctx context.Context, in *FindManyRequest, opts ...grpc.CallOption) (*FindManyResult, error)
	SetDemoMode(ctx context.Context, in *SetDemoModeRequest, opts ...grpc.CallOption) (*SetDemoModeResult, error)
	InviteUser(ctx context.Context, in *InviteUserRequest, opts ...grpc.CallOption) (*InviteUserResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) FindOneUser(ctx context.Context, in *FindOneRequest, opts ...grpc.CallOption) (*FindOneUserResult, error) {
	out := new(FindOneUserResult)
	err := c.cc.Invoke(ctx, "/goosechase.data.user.UserService/FindOneUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) FindMany(ctx context.Context, in *FindManyRequest, opts ...grpc.CallOption) (*FindManyResult, error) {
	out := new(FindManyResult)
	err := c.cc.Invoke(ctx, "/goosechase.data.user.UserService/FindMany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SetDemoMode(ctx context.Context, in *SetDemoModeRequest, opts ...grpc.CallOption) (*SetDemoModeResult, error) {
	out := new(SetDemoModeResult)
	err := c.cc.Invoke(ctx, "/goosechase.data.user.UserService/SetDemoMode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) InviteUser(ctx context.Context, in *InviteUserRequest, opts ...grpc.CallOption) (*InviteUserResponse, error) {
	out := new(InviteUserResponse)
	err := c.cc.Invoke(ctx, "/goosechase.data.user.UserService/InviteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	FindOneUser(context.Context, *FindOneRequest) (*FindOneUserResult, error)
	FindMany(context.Context, *FindManyRequest) (*FindManyResult, error)
	SetDemoMode(context.Context, *SetDemoModeRequest) (*SetDemoModeResult, error)
	InviteUser(context.Context, *InviteUserRequest) (*InviteUserResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) FindOneUser(context.Context, *FindOneRequest) (*FindOneUserResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOneUser not implemented")
}
func (UnimplementedUserServiceServer) FindMany(context.Context, *FindManyRequest) (*FindManyResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindMany not implemented")
}
func (UnimplementedUserServiceServer) SetDemoMode(context.Context, *SetDemoModeRequest) (*SetDemoModeResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDemoMode not implemented")
}
func (UnimplementedUserServiceServer) InviteUser(context.Context, *InviteUserRequest) (*InviteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InviteUser not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_FindOneUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FindOneUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goosechase.data.user.UserService/FindOneUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FindOneUser(ctx, req.(*FindOneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_FindMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindManyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FindMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goosechase.data.user.UserService/FindMany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FindMany(ctx, req.(*FindManyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SetDemoMode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDemoModeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SetDemoMode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goosechase.data.user.UserService/SetDemoMode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SetDemoMode(ctx, req.(*SetDemoModeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_InviteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InviteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).InviteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goosechase.data.user.UserService/InviteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).InviteUser(ctx, req.(*InviteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "goosechase.data.user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindOneUser",
			Handler:    _UserService_FindOneUser_Handler,
		},
		{
			MethodName: "FindMany",
			Handler:    _UserService_FindMany_Handler,
		},
		{
			MethodName: "SetDemoMode",
			Handler:    _UserService_SetDemoMode_Handler,
		},
		{
			MethodName: "InviteUser",
			Handler:    _UserService_InviteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
