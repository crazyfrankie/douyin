package ioc

import (
	"github.com/gin-gonic/gin"

	"github.com/crazyfrankie/douyin/bff/app/comment/handler"
	"github.com/crazyfrankie/douyin/bff/mw"
)

func InitWeb(mws []gin.HandlerFunc, comment *handler.Handler) *gin.Engine {
	server := gin.Default()
	server.Use(mws...)

	comment.RegisterRoute(server)

	return server
}

func InitMiddlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		mw.NewAuthBuilder().Auth(),
	}
}
