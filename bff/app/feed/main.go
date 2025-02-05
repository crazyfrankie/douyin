package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/crazyfrankie/douyin/bff/config"
	"github.com/crazyfrankie/douyin/bff/mw"
	"github.com/crazyfrankie/douyin/rpc_gen/feed"
)

func main() {
	mux := runtime.NewServeMux()

	client := initClient()

	err := feed.RegisterFeedServiceHandlerClient(context.Background(), mux, client)
	if err != nil {
		panic(err)
	}

	handler := mw.NewAuthBuilder().Auth(mux)

	log.Printf("HTTP server listening on %s", config.GetConf().Server.Feed)
	if err := http.ListenAndServe(fmt.Sprintf("%s", config.GetConf().Server.Feed), handler); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func initClient() feed.FeedServiceClient {
	conn, err := grpc.NewClient("localhost:50053",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	return feed.NewFeedServiceClient(conn)
}
