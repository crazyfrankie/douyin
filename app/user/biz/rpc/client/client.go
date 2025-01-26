package client

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/crazyfrankie/douyin/rpc_gen/favorite"
	"github.com/crazyfrankie/douyin/rpc_gen/publish"
	"github.com/crazyfrankie/douyin/rpc_gen/relation"
)

func InitFavoriteClient() favorite.FavoriteServiceClient {
	conn, err := grpc.NewClient("localhost:50052",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	return favorite.NewFavoriteServiceClient(conn)
}

func InitPublishClient() publish.PublishServiceClient {
	conn, err := grpc.NewClient("localhost:50054",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	return publish.NewPublishServiceClient(conn)
}

func InitRelationClient() relation.RelationServiceClient {
	conn, err := grpc.NewClient("localhost:50056",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	return relation.NewRelationServiceClient(conn)
}
