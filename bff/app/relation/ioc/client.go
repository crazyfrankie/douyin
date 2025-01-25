package ioc

import (
	"google.golang.org/grpc"

	"github.com/crazyfrankie/douyin/rpc_gen/relation"
)

func InitClient() relation.RelationServiceClient {
	conn, err := grpc.NewClient("localhost:50056")
	if err != nil {
		panic(err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}()
	return relation.NewRelationServiceClient(conn)
}
