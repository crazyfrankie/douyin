package server

import (
	"context"
	"github.com/crazyfrankie/douyin/app/favorite/biz/service"
	"github.com/crazyfrankie/douyin/rpc_gen/favorite"
)

type FavoriteServer struct {
	favorite.UnimplementedFavoriteServiceServer
	svc *service.FavoriteService
}

func NewFavoriteServer(svc *service.FavoriteService) *FavoriteServer {
	return &FavoriteServer{svc: svc}
}

func (f *FavoriteServer) FavoriteCount(ctx context.Context, request *favorite.FavoriteCountRequest) (*favorite.FavoriteCountResponse, error) {
	count, err := f.svc.GetUserFavoriteCount(ctx, request.UserId)
	if err != nil {
		return nil, err
	}

	return &favorite.FavoriteCountResponse{
		Count: count,
	}, nil
}
