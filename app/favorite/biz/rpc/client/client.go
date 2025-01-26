package client

import (
	"google.golang.org/grpc"

	"github.com/crazyfrankie/douyin/rpc_gen/feed"
)

func InitFeedClient() feed.FeedServiceClient {
	conn, err := grpc.NewClient("localhost:50053")
	if err != nil {
		panic(err)
	}

	return feed.NewFeedServiceClient(conn)
}
