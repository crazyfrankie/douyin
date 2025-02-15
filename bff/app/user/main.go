package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"net/http"
	"time"

	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/crazyfrankie/douyin/bff/config"
	"github.com/crazyfrankie/douyin/bff/mw"
	"github.com/crazyfrankie/douyin/rpc_gen/user"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func main() {
	mux := runtime.NewServeMux()

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{config.GetConf().ETCD.Addr},
		DialTimeout: time.Second * 5,
	})

	client := initClient(cli)

	err = user.RegisterUserServiceHandlerClient(context.Background(), mux, client)
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

func initClient(cli *clientv3.Client) user.UserServiceClient {
	builder, err := resolver.NewBuilder(cli)
	conn, err := grpc.Dial("etcd:///service/user", grpc.WithResolvers(builder),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	return user.NewUserServiceClient(conn)
}
