package server

import (
	"context"
	"github.com/crazyfrankie/douyin/app/user/biz/service"
	"github.com/crazyfrankie/douyin/rpc_gen/user"
)

type UserServer struct {
	user.UnimplementedUserServiceServer
	svc *service.UserService
}

func NewUserServer(svc *service.UserService) *UserServer {
	return &UserServer{svc: svc}
}

func (u *UserServer) GetUserInfo(ctx context.Context, req *user.GetUserInfoRequest) (*user.GetUserInfoResponse, error) {
	res, err := u.svc.GetUserInfo(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
 	
	return &user.GetUserInfoResponse{
		User: res,
	}, nil
}
