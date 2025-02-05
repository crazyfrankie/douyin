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
	"github.com/crazyfrankie/douyin/rpc_gen/publish"
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

	err := publish.RegisterPublishServiceHandlerClient(context.Background(), mux, client)
	if err != nil {
		panic(err)
	}

	handler := mw.NewAuthBuilder().Auth(mux)

	log.Printf("HTTP server listening on %s", config.GetConf().Server.Publish)
	if err := http.ListenAndServe(fmt.Sprintf("%s", config.GetConf().Server.Publish), handler); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func initClient() publish.PublishServiceClient {
	conn, err := grpc.NewClient("localhost:50054",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	return publish.NewPublishServiceClient(conn)
}
