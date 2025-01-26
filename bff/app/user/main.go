package main

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/proto"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/crazyfrankie/douyin/bff/config"
	"github.com/crazyfrankie/douyin/bff/mw"
	"github.com/crazyfrankie/douyin/rpc_gen/user"
)

func main() {
	mux := runtime.NewServeMux(
		runtime.WithErrorHandler(func(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
			log.Printf("gRPC-Gateway error: %v", err)
			runtime.DefaultHTTPErrorHandler(ctx, mux, marshaler, w, r, err)
		}),
		runtime.WithForwardResponseOption(func(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
			if r, ok := resp.(*user.RegisterResponse); ok {
				w.Header().Set("x-jwt-token", r.Token)
			}
			return nil
		}))
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
