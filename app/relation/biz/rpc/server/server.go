package server

import (
	"context"
	"github.com/crazyfrankie/douyin/app/relation/mw"

	"google.golang.org/grpc"

	"github.com/crazyfrankie/douyin/app/relation/biz/service"
	"github.com/crazyfrankie/douyin/rpc_gen/relation"
)

type RelationServer struct {
	relation.UnimplementedRelationServiceServer
	svc *service.RelationService
}

func NewRelationServer(svc *service.RelationService) *RelationServer {
	return &RelationServer{svc: svc}
}

func (r *RelationServer) RegisterServer(server *grpc.Server) {
	relation.RegisterRelationServiceServer(server, r)
}

func (r *RelationServer) RelationAction(ctx context.Context, request *relation.RelationActionRequest) (*relation.RelationActionResponse, error) {
	claims, err := mw.ParseToken(request.GetToken())
	if err != nil {
		return nil, err
	}

	newCtx := context.WithValue(ctx, "user_id", claims["user_id"])

	err = r.svc.FollowAction(newCtx, request)
	if err != nil {
		return nil, err
	}

	return &relation.RelationActionResponse{}, nil
}

func (r *RelationServer) RelationFollowList(ctx context.Context, request *relation.RelationFollowListRequest) (*relation.RelationFollowListResponse, error) {
	claims, err := mw.ParseToken(request.GetToken())
	if err != nil {
		return nil, err
	}

	newCtx := context.WithValue(ctx, "user_id", claims["user_id"])

	resp, err := r.svc.FollowList(newCtx, request)
	if err != nil {
		return nil, err
	}

	return &relation.RelationFollowListResponse{
		UserList: resp,
	}, nil
}

func (r *RelationServer) RelationFollowerList(ctx context.Context, request *relation.RelationFollowerListRequest) (*relation.RelationFollowerListResponse, error) {
	claims, err := mw.ParseToken(request.GetToken())
	if err != nil {
		return nil, err
	}

	newCtx := context.WithValue(ctx, "user_id", claims["user_id"])

	resp, err := r.svc.FollowerList(newCtx, request)
	if err != nil {
		return nil, err
	}

	return &relation.RelationFollowerListResponse{
		UserList: resp,
	}, nil
}

func (r *RelationServer) RelationFriendList(ctx context.Context, request *relation.RelationFriendListRequest) (*relation.RelationFriendListResponse, error) {
	claims, err := mw.ParseToken(request.GetToken())
	if err != nil {
		return nil, err
	}

	newCtx := context.WithValue(ctx, "user_id", claims["user_id"])

	resp, err := r.svc.FriendList(newCtx, request)
	if err != nil {
		return nil, err
	}

	return &relation.RelationFriendListResponse{
		UserList: resp,
	}, nil
}

func (r *RelationServer) RelationFollowCount(ctx context.Context, request *relation.RelationFollowCountRequest) (*relation.RelationFollowCountResponse, error) {
	followCount, followerCount, err := r.svc.GetFollowCount(ctx, request)
	if err != nil {
		return nil, err
	}

	return &relation.RelationFollowCountResponse{
		FollowCount:   followCount,
		FollowerCount: followerCount,
	}, nil
}

func (r *RelationServer) RelationIsFollow(ctx context.Context, request *relation.RelationIsFollowRequest) (*relation.RelationIsFollowResponse, error) {
	resp, err := r.svc.IsFollow(ctx, request)
	if err != nil {
		return nil, err
	}

	return &relation.RelationIsFollowResponse{
		IsFollow: resp,
	}, nil
}
