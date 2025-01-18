package client

import (
	"github.com/crazyfrankie/douyin/rpc_gen/user"
	"google.golang.org/grpc"
	"sync"

	"github.com/crazyfrankie/douyin/rpc_gen/feed"
)

var (
	FeedClient feed.FeedServiceClient
	UserClient user.UserServiceClient
	once       sync.Once
)

func InitClient() {
	once.Do(func() {
		initFeedClient()
	})
}

func initFeedClient() {
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
	FeedClient = feed.NewFeedServiceClient(conn)
}
