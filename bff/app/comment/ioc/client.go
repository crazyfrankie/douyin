package ioc

import (
	"google.golang.org/grpc"

	"github.com/crazyfrankie/douyin/rpc_gen/comment"
)

func InitClient() comment.CommentServiceClient {
	conn, err := grpc.NewClient("localhost:50055")
	if err != nil {
		panic(err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}()
	return comment.NewCommentServiceClient(conn)
}
