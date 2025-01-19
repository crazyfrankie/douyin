package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	app := InitApp()

	err = app.RPCServer.Serve()

	serviceRegister(config.GetConf().RPC.Address)

	server := &http.Server{
		Addr:    config.GetConf().Server.Address,
		Handler: app.HTTPServer,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Printf("Server is running at http://%s", config.GetConf().Server.Address)

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Print("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced shutting down: %s", err)
	}

	log.Printf("Server exited gracefully")
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
	serviceKey := "/services/user/" + address
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
