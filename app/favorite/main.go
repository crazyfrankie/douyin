package main

import (
	"context"
	"fmt"
	"github.com/crazyfrankie/douyin/app/favorite/ioc"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"log"
	"time"

	"github.com/joho/godotenv"
	clientv3 "go.etcd.io/etcd/client/v3"

	"github.com/crazyfrankie/douyin/app/favorite/config"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	app := ioc.InitApp()

	serviceRegister(config.GetConf().RPC.Address)

	err = app.RPCServer.Serve()
	if err != nil {
		panic(err)
	}
}

func serviceRegister(port string) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{config.GetConf().Etcd.Address},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalf("failed to connect to etcd: %v", err)
	}

	em, err := endpoints.NewManager(cli, "service/favorite")
	if err != nil {
		panic(err)
	}

	addr := "localhost" + port
	serviceKey := "service/favorite/" + addr

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	leaseResp, err := cli.Grant(ctx, 15)
	if err != nil {
		log.Fatalf("failed to grant lease: %v", err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err = em.AddEndpoint(ctx, serviceKey, endpoints.Endpoint{Addr: addr}, clientv3.WithLease(leaseResp.ID))

	go func() {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ch, err := cli.KeepAlive(ctx, leaseResp.ID)
		if err != nil {
			log.Fatalf("KeepAlive failed: %v", err)
		}

		for {
			select {
			case _, ok := <-ch:
				if !ok { // 通道关闭，租约停止
					log.Println("KeepAlive channel closed")
					return
				}
				fmt.Println("Lease renewed")
			case <-ctx.Done():
				return
			}
		}
	}()
}
