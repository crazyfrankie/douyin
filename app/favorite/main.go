package main

import (
	"context"
	"github.com/crazyfrankie/douyin/app/favorite/ioc"
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

	err = app.RPCServer.Serve()
	if err != nil {
		panic(err)
	}

	serviceRegister(config.GetConf().RPC.Address)
}

func serviceRegister(address string) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{config.GetConf().Etcd.Address},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalf("failed to connect to etcd: %v", err)
	}
	defer cli.Close()

	// Register service
	serviceKey := "/services/favorite/" + address
	leaseResp, err := cli.Grant(context.Background(), 5)
	if err != nil {
		log.Fatalf("failed to grant lease: %v", err)
	}

	_, err = cli.Put(context.Background(), serviceKey, address, clientv3.WithLease(leaseResp.ID))
	if err != nil {
		log.Fatalf("failed to put key: %v", err)
	}

	// Keep alive lease
	ch, err := cli.KeepAlive(context.Background(), leaseResp.ID)
	if err != nil {
		log.Fatalf("failed to keep alive lease: %v", err)
	}

	for range ch {
		// Keep lease alive
	}
}
