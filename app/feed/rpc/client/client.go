package client

import (
	"sync"

	"google.golang.org/grpc"

	"github.com/crazyfrankie/douyin/rpc_gen/user"
)

var (
	UserClient user.UserServiceClient
	once       sync.Once
)

func InitClient() {
	once.Do(func() {
		initUserClient()
	})
}

func initUserClient() {
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
	UserClient = user.NewUserServiceClient(conn)
}
