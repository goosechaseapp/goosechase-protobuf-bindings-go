// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: client.proto

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

// ClientServiceClient is the client API for ClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientServiceClient interface {
	CreateClient(ctx context.Context, in *CreateClientRequest, opts ...grpc.CallOption) (*CreateClientResult, error)
}

type clientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientServiceClient(cc grpc.ClientConnInterface) ClientServiceClient {
	return &clientServiceClient{cc}
}

func (c *clientServiceClient) CreateClient(ctx context.Context, in *CreateClientRequest, opts ...grpc.CallOption) (*CreateClientResult, error) {
	out := new(CreateClientResult)
	err := c.cc.Invoke(ctx, "/goosechase.data.client.ClientService/CreateClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientServiceServer is the server API for ClientService service.
// All implementations must embed UnimplementedClientServiceServer
// for forward compatibility
type ClientServiceServer interface {
	CreateClient(context.Context, *CreateClientRequest) (*CreateClientResult, error)
	mustEmbedUnimplementedClientServiceServer()
}

// UnimplementedClientServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClientServiceServer struct {
}

func (UnimplementedClientServiceServer) CreateClient(context.Context, *CreateClientRequest) (*CreateClientResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClient not implemented")
}
func (UnimplementedClientServiceServer) mustEmbedUnimplementedClientServiceServer() {}

// UnsafeClientServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientServiceServer will
// result in compilation errors.
type UnsafeClientServiceServer interface {
	mustEmbedUnimplementedClientServiceServer()
}

func RegisterClientServiceServer(s grpc.ServiceRegistrar, srv ClientServiceServer) {
	s.RegisterService(&ClientService_ServiceDesc, srv)
}

func _ClientService_CreateClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).CreateClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goosechase.data.client.ClientService/CreateClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).CreateClient(ctx, req.(*CreateClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientService_ServiceDesc is the grpc.ServiceDesc for ClientService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "goosechase.data.client.ClientService",
	HandlerType: (*ClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateClient",
			Handler:    _ClientService_CreateClient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client.proto",
}
