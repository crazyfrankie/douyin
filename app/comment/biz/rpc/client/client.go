package client

import (
	"google.golang.org/grpc"
	
	"github.com/crazyfrankie/douyin/rpc_gen/user"
)

func InitUserClient() user.UserServiceClient {
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
	return user.NewUserServiceClient(conn)
}
