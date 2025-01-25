package ioc

import (
	"google.golang.org/grpc"

	"github.com/crazyfrankie/douyin/rpc_gen/favorite"
)

func InitClient() favorite.FavoriteServiceClient {
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
