package ioc

import (
	"github.com/gin-gonic/gin"

	"github.com/crazyfrankie/douyin/app/feed/rpc/server"
)

type App struct {
	HTTPServer *gin.Engine
	RPCServer  *server.VideoServer
}
