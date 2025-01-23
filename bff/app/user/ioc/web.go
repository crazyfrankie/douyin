package ioc

import (
	"github.com/gin-gonic/gin"
	
	"github.com/crazyfrankie/douyin/bff/app/user/handler"
	"github.com/crazyfrankie/douyin/bff/mw"
)

func InitWeb(mws []gin.HandlerFunc, user *handler.Handler) *gin.Engine {
	server := gin.Default()
	server.Use(mws...)

	user.RegisterRoute(server)

	return server
}

func InitMiddlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		mw.NewAuthBuilder().
			IgnorePath("/api/user/login").
			IgnorePath("/api/user/signup").
			Auth(),
	}
}
