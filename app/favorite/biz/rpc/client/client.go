package client

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/crazyfrankie/douyin/rpc_gen/feed"
)

func InitFeedClient() feed.FeedServiceClient {
	conn, err := grpc.NewClient("localhost:50053",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	return feed.NewFeedServiceClient(conn)
}
