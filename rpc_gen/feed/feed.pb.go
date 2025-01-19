// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.29.3
// source: idl/feed.proto

package feed

import (
	common "github.com/crazyfrankie/douyin/rpc_gen/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type VideoListRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	VideoIds      []int64                `protobuf:"varint,1,rep,packed,name=video_ids,json=videoIds,proto3" json:"video_ids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VideoListRequest) Reset() {
	*x = VideoListRequest{}
	mi := &file_idl_feed_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VideoListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoListRequest) ProtoMessage() {}

func (x *VideoListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_feed_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoListRequest.ProtoReflect.Descriptor instead.
func (*VideoListRequest) Descriptor() ([]byte, []int) {
	return file_idl_feed_proto_rawDescGZIP(), []int{0}
}

func (x *VideoListRequest) GetVideoIds() []int64 {
	if x != nil {
		return x.VideoIds
	}
	return nil
}

type VideoListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Videos        []*common.Video        `protobuf:"bytes,1,rep,name=videos,proto3" json:"videos,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VideoListResponse) Reset() {
	*x = VideoListResponse{}
	mi := &file_idl_feed_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VideoListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoListResponse) ProtoMessage() {}

func (x *VideoListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_feed_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoListResponse.ProtoReflect.Descriptor instead.
func (*VideoListResponse) Descriptor() ([]byte, []int) {
	return file_idl_feed_proto_rawDescGZIP(), []int{1}
}

func (x *VideoListResponse) GetVideos() []*common.Video {
	if x != nil {
		return x.Videos
	}
	return nil
}

type VideoInfoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	VideoId       int64                  `protobuf:"varint,2,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VideoInfoRequest) Reset() {
	*x = VideoInfoRequest{}
	mi := &file_idl_feed_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VideoInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoInfoRequest) ProtoMessage() {}

func (x *VideoInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_feed_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoInfoRequest.ProtoReflect.Descriptor instead.
func (*VideoInfoRequest) Descriptor() ([]byte, []int) {
	return file_idl_feed_proto_rawDescGZIP(), []int{2}
}

func (x *VideoInfoRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *VideoInfoRequest) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

type VideoInfoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Video         *common.Video          `protobuf:"bytes,1,opt,name=video,proto3" json:"video,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VideoInfoResponse) Reset() {
	*x = VideoInfoResponse{}
	mi := &file_idl_feed_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VideoInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoInfoResponse) ProtoMessage() {}

func (x *VideoInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_feed_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoInfoResponse.ProtoReflect.Descriptor instead.
func (*VideoInfoResponse) Descriptor() ([]byte, []int) {
	return file_idl_feed_proto_rawDescGZIP(), []int{3}
}

func (x *VideoInfoResponse) GetVideo() *common.Video {
	if x != nil {
		return x.Video
	}
	return nil
}

type VideoExistsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	VideoId       int64                  `protobuf:"varint,1,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VideoExistsRequest) Reset() {
	*x = VideoExistsRequest{}
	mi := &file_idl_feed_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VideoExistsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoExistsRequest) ProtoMessage() {}

func (x *VideoExistsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_feed_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoExistsRequest.ProtoReflect.Descriptor instead.
func (*VideoExistsRequest) Descriptor() ([]byte, []int) {
	return file_idl_feed_proto_rawDescGZIP(), []int{4}
}

func (x *VideoExistsRequest) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

type VideoExistsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Exists        bool                   `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VideoExistsResponse) Reset() {
	*x = VideoExistsResponse{}
	mi := &file_idl_feed_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VideoExistsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoExistsResponse) ProtoMessage() {}

func (x *VideoExistsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_feed_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoExistsResponse.ProtoReflect.Descriptor instead.
func (*VideoExistsResponse) Descriptor() ([]byte, []int) {
	return file_idl_feed_proto_rawDescGZIP(), []int{5}
}

func (x *VideoExistsResponse) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

var File_idl_feed_proto protoreflect.FileDescriptor

var file_idl_feed_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x69, 0x64, 0x6c, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x66, 0x65, 0x65, 0x64, 0x1a, 0x10, 0x69, 0x64, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x10, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x73, 0x22, 0x3a, 0x0a, 0x11, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25,
	0x0a, 0x06, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x06, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x73, 0x22, 0x46, 0x0a, 0x10, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x22, 0x38, 0x0a,
	0x11, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x52, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x22, 0x2f, 0x0a, 0x12, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x13, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x32, 0xcd, 0x01, 0x0a, 0x0b, 0x46, 0x65, 0x65, 0x64,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x66,
	0x65, 0x65, 0x64, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x16, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x66, 0x65, 0x65,
	0x64, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0b, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x45, 0x78, 0x69, 0x73,
	0x74, 0x73, 0x12, 0x18, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x45,
	0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x66,
	0x65, 0x65, 0x64, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2f, 0x66, 0x65, 0x65, 0x64,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idl_feed_proto_rawDescOnce sync.Once
	file_idl_feed_proto_rawDescData = file_idl_feed_proto_rawDesc
)

func file_idl_feed_proto_rawDescGZIP() []byte {
	file_idl_feed_proto_rawDescOnce.Do(func() {
		file_idl_feed_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_feed_proto_rawDescData)
	})
	return file_idl_feed_proto_rawDescData
}

var file_idl_feed_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_idl_feed_proto_goTypes = []any{
	(*VideoListRequest)(nil),    // 0: feed.VideoListRequest
	(*VideoListResponse)(nil),   // 1: feed.VideoListResponse
	(*VideoInfoRequest)(nil),    // 2: feed.VideoInfoRequest
	(*VideoInfoResponse)(nil),   // 3: feed.VideoInfoResponse
	(*VideoExistsRequest)(nil),  // 4: feed.VideoExistsRequest
	(*VideoExistsResponse)(nil), // 5: feed.VideoExistsResponse
	(*common.Video)(nil),        // 6: common.Video
}
var file_idl_feed_proto_depIdxs = []int32{
	6, // 0: feed.VideoListResponse.videos:type_name -> common.Video
	6, // 1: feed.VideoInfoResponse.video:type_name -> common.Video
	0, // 2: feed.FeedService.VideoList:input_type -> feed.VideoListRequest
	2, // 3: feed.FeedService.VideoInfo:input_type -> feed.VideoInfoRequest
	4, // 4: feed.FeedService.VideoExists:input_type -> feed.VideoExistsRequest
	1, // 5: feed.FeedService.VideoList:output_type -> feed.VideoListResponse
	3, // 6: feed.FeedService.VideoInfo:output_type -> feed.VideoInfoResponse
	5, // 7: feed.FeedService.VideoExists:output_type -> feed.VideoExistsResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_idl_feed_proto_init() }
func file_idl_feed_proto_init() {
	if File_idl_feed_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_idl_feed_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_feed_proto_goTypes,
		DependencyIndexes: file_idl_feed_proto_depIdxs,
		MessageInfos:      file_idl_feed_proto_msgTypes,
	}.Build()
	File_idl_feed_proto = out.File
	file_idl_feed_proto_rawDesc = nil
	file_idl_feed_proto_goTypes = nil
	file_idl_feed_proto_depIdxs = nil
}
