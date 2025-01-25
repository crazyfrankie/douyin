package ioc

import (
	"google.golang.org/grpc"
	
	"github.com/crazyfrankie/douyin/rpc_gen/message"
)

func InitClient() message.MessageServiceClient {
	conn, err := grpc.NewClient("localhost:50057")
	if err != nil {
		panic(err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}()
	return message.NewMessageServiceClient(conn)
}
