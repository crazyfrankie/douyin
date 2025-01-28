package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/crazyfrankie/douyin/bff/config"
	"github.com/crazyfrankie/douyin/bff/mw"
	"github.com/crazyfrankie/douyin/rpc_gen/user"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func main() {
	mux := runtime.NewServeMux()
	client := initClient()

	err := user.RegisterUserServiceHandlerClient(context.Background(), mux, client)
	if err != nil {
		panic(err)
	}
	handler := mw.NewAuthBuilder().
		IgnorePath("/api/user/login").
		IgnorePath("/api/user/register").
		Auth(mux)

	log.Printf("HTTP server listening on %s", config.GetConf().Server.User)
	if err := http.ListenAndServe(fmt.Sprintf("%s", config.GetConf().Server.User), handler); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func initClient() user.UserServiceClient {
	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	return user.NewUserServiceClient(conn)
}
