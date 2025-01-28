package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	"github.com/crazyfrankie/douyin/bff/config"
	"github.com/crazyfrankie/douyin/bff/mw"
	"github.com/crazyfrankie/douyin/rpc_gen/relation"
)

func main() {
	mux := runtime.NewServeMux(runtime.WithMetadata(func(ctx context.Context, request *http.Request) metadata.MD {
		md := metadata.MD{}

		if userID, ok := request.Context().Value("user_id").(string); ok {
			md.Set("user_id", userID)
		}

		return md
	}))

	client := initClient()

	err := relation.RegisterRelationServiceHandlerClient(context.Background(), mux, client)
	if err != nil {
		panic(err)
	}

	handler := mw.NewAuthBuilder().Auth(mux)

	log.Printf("HTTP server listening on %s", config.GetConf().Server.Relation)
	if err := http.ListenAndServe(fmt.Sprintf("%s", config.GetConf().Server.Relation), handler); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func initClient() relation.RelationServiceClient {
	conn, err := grpc.NewClient("localhost:50056",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	return relation.NewRelationServiceClient(conn)
}
