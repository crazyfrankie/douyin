package ioc

import (
	"github.com/gin-gonic/gin"

	"github.com/crazyfrankie/douyin/bff/app/favorite/handler"
	"github.com/crazyfrankie/douyin/bff/mw"
)

func InitWeb(mws []gin.HandlerFunc, favorite *handler.Handler) *gin.Engine {
	server := gin.Default()
	server.Use(mws...)

	favorite.RegisterRoute(server)

	return server
}

func InitMiddlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		mw.NewAuthBuilder().Auth(),
	}
}
