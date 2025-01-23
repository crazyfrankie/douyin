package ioc

import (
	"google.golang.org/grpc"
	
	"github.com/crazyfrankie/douyin/rpc_gen/publish"
)

func InitClient() publish.PublishServiceClient {
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
	return publish.NewPublishServiceClient(conn)
}
