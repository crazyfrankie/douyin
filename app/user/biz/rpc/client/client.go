package client

import (
	"github.com/crazyfrankie/douyin/rpc_gen/publish"
	"google.golang.org/grpc"

	"github.com/crazyfrankie/douyin/rpc_gen/favorite"
)

func NewFavoriteClient() favorite.FavoriteServiceClient {
	conn, err := grpc.NewClient("localhost:50052")
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

func NewPublishClient() publish.PublishServiceClient {
	conn, err := grpc.NewClient("localhost:50054")
	if err != nil {
		panic(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)
	return publish.NewPublishServiceClient(conn)
}
