package feed

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/crazyfrankie/douyin/app/feed/config"
	"github.com/crazyfrankie/douyin/app/feed/ioc"
	"github.com/crazyfrankie/douyin/app/feed/rpc/client"
	"github.com/crazyfrankie/douyin/rpc_gen/feed"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	app := ioc.InitApp()

	rpcInit(app, config.GetConf().RPC.Address)

	client.InitClient()

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

type healthImpl struct {
	grpc_health_v1.UnimplementedHealthServer
}

func (h *healthImpl) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

func rpcInit(app *ioc.App, address string) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", address))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	feed.RegisterFeedServiceServer(s, app.RPCServer)
	// 健康检查
	grpc_health_v1.RegisterHealthServer(s, &healthImpl{})

	go func() {
		if err := s.Serve(lis); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Printf("gRPC Server is running at %s", address)

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
