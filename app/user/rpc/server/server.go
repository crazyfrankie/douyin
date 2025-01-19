package server

import (
	"context"
	"github.com/crazyfrankie/douyin/app/user/biz"
	"github.com/crazyfrankie/douyin/app/user/biz/service"
	"github.com/crazyfrankie/douyin/rpc_gen/user"
	"google.golang.org/grpc"
)

type UserServer struct {
	user.UnimplementedUserServiceServer
	svc *service.UserService
}

func NewUserServer(svc *service.UserService) *UserServer {
	return &UserServer{svc: svc}
}

func (u *UserServer) RegisterServer(server *grpc.Server) {
	user.RegisterUserServiceServer(server, u)
}

func (u *UserServer) Register(ctx context.Context, req *user.RegisterRequest) (*user.RegisterResponse, error) {
	token, err := u.svc.Register(ctx, biz.RegisterReq{
		Name:     req.Name,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &user.RegisterResponse{
		Token: token,
	}, nil
}

func (u *UserServer) Login(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	token, err := u.svc.Login(ctx, biz.LoginReq{
		Name:     req.Name,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &user.LoginResponse{
		Token: token,
	}, nil
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
