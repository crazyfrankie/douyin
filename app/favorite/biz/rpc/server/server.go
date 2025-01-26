package server

import (
	"context"
	"github.com/crazyfrankie/douyin/app/favorite/mw"
	"google.golang.org/grpc"

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

func (f *FavoriteServer) RegisterServer(server *grpc.Server) {
	favorite.RegisterFavoriteServiceServer(server, f)
}

func (f *FavoriteServer) FavoriteAction(ctx context.Context, request *favorite.FavoriteActionRequest) (*favorite.FavoriteActionResponse, error) {
	claims, err := mw.ParseToken(request.GetToken())
	if err != nil {
		return nil, err
	}

	newCtx := context.WithValue(ctx, "user_id", claims["user_id"])
	err = f.svc.FavoriteAction(newCtx, request)
	if err != nil {
		return nil, err
	}

	return &favorite.FavoriteActionResponse{}, nil
}

func (f *FavoriteServer) FavoriteList(ctx context.Context, request *favorite.FavoriteListRequest) (*favorite.FavoriteListResponse, error) {
	claims, err := mw.ParseToken(request.GetToken())
	if err != nil {
		return nil, err
	}

	newCtx := context.WithValue(ctx, "user_id", claims["user_id"])

	videos, err := f.svc.FavoriteList(newCtx, request)
	if err != nil {
		return nil, err
	}

	return &favorite.FavoriteListResponse{
		Videos: videos,
	}, nil
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

func (f *FavoriteServer) UserFavorited(ctx context.Context, request *favorite.UserFavoritedRequest) (*favorite.UserFavoritedResponse, error) {
	count, err := f.svc.GetUserFavoritedCount(ctx, request.VideoId)
	if err != nil {
		return nil, err
	}

	return &favorite.UserFavoritedResponse{
		Count: count,
	}, nil
}

func (f *FavoriteServer) VideoFavoriteCount(ctx context.Context, request *favorite.VideoFavoriteCountRequest) (*favorite.VideoFavoriteCountResponse, error) {
	count, err := f.svc.GetVideoFavoriteCount(ctx, request.VideoId)
	if err != nil {
		return nil, err
	}

	return &favorite.VideoFavoriteCountResponse{
		Count: count,
	}, nil
}

func (f *FavoriteServer) IsFavorite(ctx context.Context, request *favorite.IsFavoriteRequest) (*favorite.IsFavoriteResponse, error) {
	isFavorite, err := f.svc.GetIsFavorite(ctx, request.VideoId, request.UserId)
	if err != nil {
		return nil, err
	}

	return &favorite.IsFavoriteResponse{
		IsFavorite: isFavorite,
	}, nil
}
