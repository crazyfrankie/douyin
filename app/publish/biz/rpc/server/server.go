package server

import (
	"context"

	"google.golang.org/grpc"

	"github.com/crazyfrankie/douyin/app/publish/biz/service"
	"github.com/crazyfrankie/douyin/rpc_gen/publish"
)

type PublishServer struct {
	publish.UnimplementedPublishServiceServer
	svc *service.PublishService
}

func NewPublishServer(svc *service.PublishService) *PublishServer {
	return &PublishServer{svc: svc}
}

func (p *PublishServer) RegisterServer(server *grpc.Server) {
	publish.RegisterPublishServiceServer(server, p)
}
func (p *PublishServer) PublishAction(ctx context.Context, request *publish.PublishActionRequest) (*publish.PublishActionResponse, error) {
	err := p.svc.PublishAction(ctx, request)
	if err != nil {
		return nil, err
	}

	return &publish.PublishActionResponse{}, nil
}

func (p *PublishServer) PublishList(ctx context.Context, request *publish.PublishListRequest) (*publish.PublishListResponse, error) {
	resp, err := p.svc.PublishList(ctx, request)
	if err != nil {
		return nil, err
	}

	return &publish.PublishListResponse{
		Videos: resp,
	}, nil
}

func (p *PublishServer) PublishCount(ctx context.Context, request *publish.PublishCountRequest) (*publish.PublishCountResponse, error) {
	videoIds, err := p.svc.PublishCount(ctx, request.GetUserId())
	if err != nil {
		return nil, err
	}

	return &publish.PublishCountResponse{
		VideoId: videoIds,
	}, nil
}
