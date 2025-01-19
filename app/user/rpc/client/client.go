package client

import (
	"google.golang.org/grpc"

	"github.com/crazyfrankie/douyin/rpc_gen/favorite"
)

func NewFavoriteClient() favorite.FavoriteServiceClient {
	conn, err := grpc.NewClient("localhost:50051")
	if err != nil {
		panic(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)
	return favorite.NewFavoriteServiceClient(conn)
}
