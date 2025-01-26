package server

import (
	"context"
	"fmt"
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
	token, err := u.svc.Register(ctx, req)
	if err != nil {
		return nil, err
	}

	return &user.RegisterResponse{
		Token: token,
	}, nil
}

func (u *UserServer) Login(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	token, err := u.svc.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	fmt.Println(token)
	return &user.LoginResponse{
		Token: token,
	}, nil
}

func (u *UserServer) GetUserInfo(ctx context.Context, req *user.GetUserInfoRequest) (*user.GetUserInfoResponse, error) {
	res, err := u.svc.GetUserInfo(ctx, req)
	if err != nil {
		return nil, err
	}

	return &user.GetUserInfoResponse{
		User: res,
	}, nil
}

func (u *UserServer) GetUserExists(ctx context.Context, req *user.GetUserExistsRequest) (*user.GetUserExistsResponse, error) {
	res, err := u.svc.GetUserExists(ctx, req)
	if err != nil {
		return nil, err
	}

	return &user.GetUserExistsResponse{
		Exists: res,
	}, nil
}
