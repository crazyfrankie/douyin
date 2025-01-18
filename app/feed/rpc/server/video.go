package server

import (
	"context"
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

func (v *VideoServer) VideoList(ctx context.Context, request *feed.FeedListRequest) (*feed.FeedListResponse, error) {
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

	return &feed.FeedListResponse{
		Videos: videos,
	}, nil
}

func (v *VideoServer) VideoInfo(ctx context.Context, request *feed.FeedInfoRequest) (*feed.FeedInfoResponse, error) {
	resp, err := v.svc.VideoInfo(ctx, request.VideoId, request.UserId)
	if err != nil {
		return nil, err
	}

	return &feed.FeedInfoResponse{
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
