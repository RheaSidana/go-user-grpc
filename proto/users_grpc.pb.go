// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.0
// source: proto/users.proto

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
	GetUserDetails(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*User, error)
	GetUserDetailsByIDs(ctx context.Context, opts ...grpc.CallOption) (UserService_GetUserDetailsByIDsClient, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUserDetails(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/users.UserService/GetUserDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserDetailsByIDs(ctx context.Context, opts ...grpc.CallOption) (UserService_GetUserDetailsByIDsClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserService_ServiceDesc.Streams[0], "/users.UserService/GetUserDetailsByIDs", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceGetUserDetailsByIDsClient{stream}
	return x, nil
}

type UserService_GetUserDetailsByIDsClient interface {
	Send(*UserIdListRequest) error
	Recv() (*UsersList, error)
	grpc.ClientStream
}

type userServiceGetUserDetailsByIDsClient struct {
	grpc.ClientStream
}

func (x *userServiceGetUserDetailsByIDsClient) Send(m *UserIdListRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *userServiceGetUserDetailsByIDsClient) Recv() (*UsersList, error) {
	m := new(UsersList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	GetUserDetails(context.Context, *UserIdRequest) (*User, error)
	GetUserDetailsByIDs(UserService_GetUserDetailsByIDsServer) error
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) GetUserDetails(context.Context, *UserIdRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserDetails not implemented")
}
func (UnimplementedUserServiceServer) GetUserDetailsByIDs(UserService_GetUserDetailsByIDsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUserDetailsByIDs not implemented")
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

func _UserService_GetUserDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/GetUserDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserDetails(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserDetailsByIDs_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UserServiceServer).GetUserDetailsByIDs(&userServiceGetUserDetailsByIDsServer{stream})
}

type UserService_GetUserDetailsByIDsServer interface {
	Send(*UsersList) error
	Recv() (*UserIdListRequest, error)
	grpc.ServerStream
}

type userServiceGetUserDetailsByIDsServer struct {
	grpc.ServerStream
}

func (x *userServiceGetUserDetailsByIDsServer) Send(m *UsersList) error {
	return x.ServerStream.SendMsg(m)
}

func (x *userServiceGetUserDetailsByIDsServer) Recv() (*UserIdListRequest, error) {
	m := new(UserIdListRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "users.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserDetails",
			Handler:    _UserService_GetUserDetails_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUserDetailsByIDs",
			Handler:       _UserService_GetUserDetailsByIDs_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/users.proto",
}
