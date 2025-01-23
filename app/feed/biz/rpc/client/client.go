package client

import (
	"github.com/crazyfrankie/douyin/rpc_gen/favorite"
	"google.golang.org/grpc"

	"github.com/crazyfrankie/douyin/rpc_gen/user"
)

func NewUserClient() user.UserServiceClient {
	conn, err := grpc.NewClient("localhost:50051")
	if err != nil {
		panic(err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}()
	return user.NewUserServiceClient(conn)
}

func NewFavoriteClient() favorite.FavoriteServiceClient {
	conn, err := grpc.NewClient("localhost:50052")
	if err != nil {
		panic(err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}()
	return favorite.NewFavoriteServiceClient(conn)
}
