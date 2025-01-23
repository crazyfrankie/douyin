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

	"github.com/crazyfrankie/douyin/bff/app/feed/ioc"
	"github.com/crazyfrankie/douyin/bff/config"
)

func main() {
	engine := ioc.InitGin()

	server := &http.Server{
		Addr:    config.GetConf().Server.Feed,
		Handler: engine,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Printf("Server is running at http://%s", config.GetConf().Server.Feed)

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
