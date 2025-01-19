package ioc

import (
	"github.com/gin-gonic/gin"

	"github.com/crazyfrankie/douyin/app/feed/rpc"
)

type App struct {
	RPCServer  *rpc.Server
	HTTPServer *gin.Engine
}
