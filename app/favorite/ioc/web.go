package ioc

import (
	"github.com/gin-gonic/gin"

	"github.com/crazyfrankie/douyin/app/favorite/biz/handler"
	"github.com/crazyfrankie/douyin/app/favorite/biz/service"
)

func InitWeb(favorite *handler.FavoriteHandler) *gin.Engine {
	server := gin.Default()
	server.Use(service.NewAuthBuilder().Auth())

	favorite.RegisterRoute(server)

	return server
}
