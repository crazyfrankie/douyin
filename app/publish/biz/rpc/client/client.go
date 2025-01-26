package client

import (
	"github.com/crazyfrankie/douyin/rpc_gen/favorite"
	"google.golang.org/grpc"
)

func InitFavoriteClient() favorite.FavoriteServiceClient {
	conn, err := grpc.NewClient("localhost:50052")
	if err != nil {
		panic(err)
	}

	return favorite.NewFavoriteServiceClient(conn)
}
