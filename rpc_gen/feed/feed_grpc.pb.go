// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: idl/douyin/feed.proto

package feed

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
	FeedService_Feed_FullMethodName        = "/feed.FeedService/Feed"
	FeedService_VideoList_FullMethodName   = "/feed.FeedService/VideoList"
	FeedService_VideoInfo_FullMethodName   = "/feed.FeedService/VideoInfo"
	FeedService_VideoExists_FullMethodName = "/feed.FeedService/VideoExists"
)

// FeedServiceClient is the client API for FeedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeedServiceClient interface {
	Feed(ctx context.Context, in *FeedRequest, opts ...grpc.CallOption) (*FeedResponse, error)
	VideoList(ctx context.Context, in *VideoListRequest, opts ...grpc.CallOption) (*VideoListResponse, error)
	VideoInfo(ctx context.Context, in *VideoInfoRequest, opts ...grpc.CallOption) (*VideoInfoResponse, error)
	VideoExists(ctx context.Context, in *VideoExistsRequest, opts ...grpc.CallOption) (*VideoExistsResponse, error)
}

type feedServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFeedServiceClient(cc grpc.ClientConnInterface) FeedServiceClient {
	return &feedServiceClient{cc}
}

func (c *feedServiceClient) Feed(ctx context.Context, in *FeedRequest, opts ...grpc.CallOption) (*FeedResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FeedResponse)
	err := c.cc.Invoke(ctx, FeedService_Feed_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedServiceClient) VideoList(ctx context.Context, in *VideoListRequest, opts ...grpc.CallOption) (*VideoListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VideoListResponse)
	err := c.cc.Invoke(ctx, FeedService_VideoList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedServiceClient) VideoInfo(ctx context.Context, in *VideoInfoRequest, opts ...grpc.CallOption) (*VideoInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VideoInfoResponse)
	err := c.cc.Invoke(ctx, FeedService_VideoInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedServiceClient) VideoExists(ctx context.Context, in *VideoExistsRequest, opts ...grpc.CallOption) (*VideoExistsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VideoExistsResponse)
	err := c.cc.Invoke(ctx, FeedService_VideoExists_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedServiceServer is the server API for FeedService service.
// All implementations must embed UnimplementedFeedServiceServer
// for forward compatibility.
type FeedServiceServer interface {
	Feed(context.Context, *FeedRequest) (*FeedResponse, error)
	VideoList(context.Context, *VideoListRequest) (*VideoListResponse, error)
	VideoInfo(context.Context, *VideoInfoRequest) (*VideoInfoResponse, error)
	VideoExists(context.Context, *VideoExistsRequest) (*VideoExistsResponse, error)
	mustEmbedUnimplementedFeedServiceServer()
}

// UnimplementedFeedServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFeedServiceServer struct{}

func (UnimplementedFeedServiceServer) Feed(context.Context, *FeedRequest) (*FeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Feed not implemented")
}
func (UnimplementedFeedServiceServer) VideoList(context.Context, *VideoListRequest) (*VideoListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VideoList not implemented")
}
func (UnimplementedFeedServiceServer) VideoInfo(context.Context, *VideoInfoRequest) (*VideoInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VideoInfo not implemented")
}
func (UnimplementedFeedServiceServer) VideoExists(context.Context, *VideoExistsRequest) (*VideoExistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VideoExists not implemented")
}
func (UnimplementedFeedServiceServer) mustEmbedUnimplementedFeedServiceServer() {}
func (UnimplementedFeedServiceServer) testEmbeddedByValue()                     {}

// UnsafeFeedServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeedServiceServer will
// result in compilation errors.
type UnsafeFeedServiceServer interface {
	mustEmbedUnimplementedFeedServiceServer()
}

func RegisterFeedServiceServer(s grpc.ServiceRegistrar, srv FeedServiceServer) {
	// If the following call pancis, it indicates UnimplementedFeedServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FeedService_ServiceDesc, srv)
}

func _FeedService_Feed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServiceServer).Feed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeedService_Feed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServiceServer).Feed(ctx, req.(*FeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedService_VideoList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VideoListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServiceServer).VideoList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeedService_VideoList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServiceServer).VideoList(ctx, req.(*VideoListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedService_VideoInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VideoInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServiceServer).VideoInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeedService_VideoInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServiceServer).VideoInfo(ctx, req.(*VideoInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedService_VideoExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VideoExistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServiceServer).VideoExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeedService_VideoExists_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServiceServer).VideoExists(ctx, req.(*VideoExistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FeedService_ServiceDesc is the grpc.ServiceDesc for FeedService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FeedService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "feed.FeedService",
	HandlerType: (*FeedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Feed",
			Handler:    _FeedService_Feed_Handler,
		},
		{
			MethodName: "VideoList",
			Handler:    _FeedService_VideoList_Handler,
		},
		{
			MethodName: "VideoInfo",
			Handler:    _FeedService_VideoInfo_Handler,
		},
		{
			MethodName: "VideoExists",
			Handler:    _FeedService_VideoExists_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "idl/douyin/feed.proto",
}
