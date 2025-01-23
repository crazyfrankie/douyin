package ioc

import (
	"github.com/crazyfrankie/douyin/rpc_gen/feed"
	"google.golang.org/grpc"
)

func InitClient() feed.FeedServiceClient {
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
	return feed.NewFeedServiceClient(conn)
}
