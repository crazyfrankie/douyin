// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: idl/relation.proto

package relation

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	RelationService_RelationAction_FullMethodName       = "/relation.RelationService/RelationAction"
	RelationService_RelationFollowList_FullMethodName   = "/relation.RelationService/RelationFollowList"
	RelationService_RelationFollowerList_FullMethodName = "/relation.RelationService/RelationFollowerList"
	RelationService_RelationFriendList_FullMethodName   = "/relation.RelationService/RelationFriendList"
	RelationService_RelationFollowCount_FullMethodName  = "/relation.RelationService/RelationFollowCount"
	RelationService_RelationIsFollow_FullMethodName     = "/relation.RelationService/RelationIsFollow"
)

// RelationServiceClient is the client API for RelationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RelationServiceClient interface {
	RelationAction(ctx context.Context, in *RelationActionRequest, opts ...grpc.CallOption) (*RelationActionResponse, error)
	RelationFollowList(ctx context.Context, in *RelationFollowListRequest, opts ...grpc.CallOption) (*RelationFollowListResponse, error)
	RelationFollowerList(ctx context.Context, in *RelationFollowerListRequest, opts ...grpc.CallOption) (*RelationFollowerListResponse, error)
	RelationFriendList(ctx context.Context, in *RelationFriendListRequest, opts ...grpc.CallOption) (*RelationFriendListResponse, error)
	RelationFollowCount(ctx context.Context, in *RelationFollowCountRequest, opts ...grpc.CallOption) (*RelationFollowCountResponse, error)
	RelationIsFollow(ctx context.Context, in *RelationIsFollowRequest, opts ...grpc.CallOption) (*RelationIsFollowResponse, error)
}

type relationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRelationServiceClient(cc grpc.ClientConnInterface) RelationServiceClient {
	return &relationServiceClient{cc}
}

func (c *relationServiceClient) RelationAction(ctx context.Context, in *RelationActionRequest, opts ...grpc.CallOption) (*RelationActionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RelationActionResponse)
	err := c.cc.Invoke(ctx, RelationService_RelationAction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) RelationFollowList(ctx context.Context, in *RelationFollowListRequest, opts ...grpc.CallOption) (*RelationFollowListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RelationFollowListResponse)
	err := c.cc.Invoke(ctx, RelationService_RelationFollowList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) RelationFollowerList(ctx context.Context, in *RelationFollowerListRequest, opts ...grpc.CallOption) (*RelationFollowerListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RelationFollowerListResponse)
	err := c.cc.Invoke(ctx, RelationService_RelationFollowerList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) RelationFriendList(ctx context.Context, in *RelationFriendListRequest, opts ...grpc.CallOption) (*RelationFriendListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RelationFriendListResponse)
	err := c.cc.Invoke(ctx, RelationService_RelationFriendList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) RelationFollowCount(ctx context.Context, in *RelationFollowCountRequest, opts ...grpc.CallOption) (*RelationFollowCountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RelationFollowCountResponse)
	err := c.cc.Invoke(ctx, RelationService_RelationFollowCount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) RelationIsFollow(ctx context.Context, in *RelationIsFollowRequest, opts ...grpc.CallOption) (*RelationIsFollowResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RelationIsFollowResponse)
	err := c.cc.Invoke(ctx, RelationService_RelationIsFollow_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RelationServiceServer is the server API for RelationService service.
// All implementations must embed UnimplementedRelationServiceServer
// for forward compatibility.
type RelationServiceServer interface {
	RelationAction(context.Context, *RelationActionRequest) (*RelationActionResponse, error)
	RelationFollowList(context.Context, *RelationFollowListRequest) (*RelationFollowListResponse, error)
	RelationFollowerList(context.Context, *RelationFollowerListRequest) (*RelationFollowerListResponse, error)
	RelationFriendList(context.Context, *RelationFriendListRequest) (*RelationFriendListResponse, error)
	RelationFollowCount(context.Context, *RelationFollowCountRequest) (*RelationFollowCountResponse, error)
	RelationIsFollow(context.Context, *RelationIsFollowRequest) (*RelationIsFollowResponse, error)
	mustEmbedUnimplementedRelationServiceServer()
}

// UnimplementedRelationServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRelationServiceServer struct{}

func (UnimplementedRelationServiceServer) RelationAction(context.Context, *RelationActionRequest) (*RelationActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RelationAction not implemented")
}
func (UnimplementedRelationServiceServer) RelationFollowList(context.Context, *RelationFollowListRequest) (*RelationFollowListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RelationFollowList not implemented")
}
func (UnimplementedRelationServiceServer) RelationFollowerList(context.Context, *RelationFollowerListRequest) (*RelationFollowerListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RelationFollowerList not implemented")
}
func (UnimplementedRelationServiceServer) RelationFriendList(context.Context, *RelationFriendListRequest) (*RelationFriendListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RelationFriendList not implemented")
}
func (UnimplementedRelationServiceServer) RelationFollowCount(context.Context, *RelationFollowCountRequest) (*RelationFollowCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RelationFollowCount not implemented")
}
func (UnimplementedRelationServiceServer) RelationIsFollow(context.Context, *RelationIsFollowRequest) (*RelationIsFollowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RelationIsFollow not implemented")
}
func (UnimplementedRelationServiceServer) mustEmbedUnimplementedRelationServiceServer() {}
func (UnimplementedRelationServiceServer) testEmbeddedByValue()                         {}

// UnsafeRelationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RelationServiceServer will
// result in compilation errors.
type UnsafeRelationServiceServer interface {
	mustEmbedUnimplementedRelationServiceServer()
}

func RegisterRelationServiceServer(s grpc.ServiceRegistrar, srv RelationServiceServer) {
	// If the following call pancis, it indicates UnimplementedRelationServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RelationService_ServiceDesc, srv)
}

func _RelationService_RelationAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelationActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).RelationAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationService_RelationAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).RelationAction(ctx, req.(*RelationActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_RelationFollowList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelationFollowListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).RelationFollowList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationService_RelationFollowList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).RelationFollowList(ctx, req.(*RelationFollowListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_RelationFollowerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelationFollowerListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).RelationFollowerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationService_RelationFollowerList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).RelationFollowerList(ctx, req.(*RelationFollowerListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_RelationFriendList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelationFriendListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).RelationFriendList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationService_RelationFriendList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).RelationFriendList(ctx, req.(*RelationFriendListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_RelationFollowCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelationFollowCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).RelationFollowCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationService_RelationFollowCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).RelationFollowCount(ctx, req.(*RelationFollowCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_RelationIsFollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelationIsFollowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).RelationIsFollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationService_RelationIsFollow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).RelationIsFollow(ctx, req.(*RelationIsFollowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RelationService_ServiceDesc is the grpc.ServiceDesc for RelationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RelationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "relation.RelationService",
	HandlerType: (*RelationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RelationAction",
			Handler:    _RelationService_RelationAction_Handler,
		},
		{
			MethodName: "RelationFollowList",
			Handler:    _RelationService_RelationFollowList_Handler,
		},
		{
			MethodName: "RelationFollowerList",
			Handler:    _RelationService_RelationFollowerList_Handler,
		},
		{
			MethodName: "RelationFriendList",
			Handler:    _RelationService_RelationFriendList_Handler,
		},
		{
			MethodName: "RelationFollowCount",
			Handler:    _RelationService_RelationFollowCount_Handler,
		},
		{
			MethodName: "RelationIsFollow",
			Handler:    _RelationService_RelationIsFollow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "idl/relation.proto",
}
