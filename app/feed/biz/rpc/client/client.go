package client

import (
	"github.com/crazyfrankie/douyin/rpc_gen/comment"
	"github.com/crazyfrankie/douyin/rpc_gen/favorite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/crazyfrankie/douyin/rpc_gen/user"
)

func InitUserClient() user.UserServiceClient {
	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}

	return user.NewUserServiceClient(conn)
}

func InitFavoriteClient() favorite.FavoriteServiceClient {
	conn, err := grpc.NewClient("localhost:50052",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	return favorite.NewFavoriteServiceClient(conn)
}

func InitCommentClient() comment.CommentServiceClient {
	conn, err := grpc.NewClient("localhost:50055",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	return comment.NewCommentServiceClient(conn)
}
