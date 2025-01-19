package ioc

import (
	"github.com/gin-gonic/gin"

	"github.com/crazyfrankie/douyin/app/user/biz/handler"
	"github.com/crazyfrankie/douyin/app/user/biz/service"
)

func InitWeb(user *handler.Handler) *gin.Engine {
	server := gin.Default()
	server.Use(service.NewAuthBuilder().IgnorePath("/api/user/register").IgnorePath("/api/user/login").Auth())

	user.RegisterRoute(server)

	return server
}
