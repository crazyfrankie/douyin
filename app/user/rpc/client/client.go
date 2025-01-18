package client

import (
	"sync"

	"google.golang.org/grpc"

	"github.com/crazyfrankie/douyin/rpc_gen/favorite"
)

var (
	FavoriteClient favorite.FavoriteServiceClient
	once           sync.Once
)

func InitClient() {
	once.Do(func() {
		initFavoriteClient()
	})
}

func initFavoriteClient() {
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
	FavoriteClient = favorite.NewFavoriteServiceClient(conn)
}
