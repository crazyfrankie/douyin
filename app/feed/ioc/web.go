package ioc

import (
	"github.com/crazyfrankie/douyin/app/feed/biz/handler"
	"github.com/crazyfrankie/douyin/app/feed/biz/service"

	"github.com/gin-gonic/gin"
)

func InitWeb(user *handler.FeedHandler) *gin.Engine {
	server := gin.Default()
	server.Use(service.NewAuthBuilder().Auth())
	user.RegisterRoute(server)

	return server
}
