package ioc

import (
	"github.com/crazyfrankie/douyin/bff/app/feed/handler"
	"github.com/crazyfrankie/douyin/bff/mw"
	"github.com/gin-gonic/gin"
)

func InitWeb(mws []gin.HandlerFunc, feed *handler.Handler) *gin.Engine {
	server := gin.Default()
	server.Use(mws...)

	feed.RegisterRoute(server)

	return server
}

func InitMiddlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		mw.NewAuthBuilder().Auth(),
	}
}
