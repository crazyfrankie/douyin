package server

import (
	"context"

	"google.golang.org/grpc"

	"github.com/crazyfrankie/douyin/app/feed/biz/service"
	"github.com/crazyfrankie/douyin/rpc_gen/common"
	"github.com/crazyfrankie/douyin/rpc_gen/feed"
)

type VideoServer struct {
	feed.UnimplementedFeedServiceServer
	svc *service.FeedService
}

func NewVideoServer(svc *service.FeedService) *VideoServer {
	return &VideoServer{svc: svc}
}

func (v *VideoServer) RegisterServer(server *grpc.Server) {
	feed.RegisterFeedServiceServer(server, v)
}

func (v *VideoServer) Feed(ctx context.Context, request *feed.FeedRequest) (*feed.FeedResponse, error) {
	resp, err := v.svc.Feed(ctx, request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VideoServer) VideoList(ctx context.Context, request *feed.VideoListRequest) (*feed.VideoListResponse, error) {
	resp, err := v.svc.VideoList(ctx, request.VideoIds)
	if err != nil {
		return nil, err
	}

	var videos []*common.Video
	for _, v := range resp {
		video := &common.Video{
			Id: v.ID,
			Author: &common.User{
				Id: v.AuthorID,
			},
			PlayUrl:  v.PlayURL,
			CoverUrl: v.CoverURL,
			Title:    v.Title,
		}

		videos = append(videos, video)
	}

	return &feed.VideoListResponse{
		Videos: videos,
	}, nil
}

func (v *VideoServer) VideoInfo(ctx context.Context, request *feed.VideoInfoRequest) (*feed.VideoInfoResponse, error) {
	resp, err := v.svc.VideoInfo(ctx, request)
	if err != nil {
		return nil, err
	}

	return &feed.VideoInfoResponse{
		Video: resp,
	}, nil
}

func (v *VideoServer) VideoExists(ctx context.Context, request *feed.VideoExistsRequest) (*feed.VideoExistsResponse, error) {
	exists, err := v.svc.GetVideoExists(ctx, request.VideoId)
	if err != nil {
		return nil, err
	}

	return &feed.VideoExistsResponse{
		Exists: exists,
	}, nil
}
