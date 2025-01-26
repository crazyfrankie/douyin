package client

import (
	"github.com/crazyfrankie/douyin/rpc_gen/favorite"
	"google.golang.org/grpc"

	"github.com/crazyfrankie/douyin/rpc_gen/user"
)

func InitUserClient() user.UserServiceClient {
	conn, err := grpc.NewClient("localhost:50051")
	if err != nil {
		panic(err)
	}

	return user.NewUserServiceClient(conn)
}

func InitFavoriteClient() favorite.FavoriteServiceClient {
	conn, err := grpc.NewClient("localhost:50052")
	if err != nil {
		panic(err)
	}

	return favorite.NewFavoriteServiceClient(conn)
}
