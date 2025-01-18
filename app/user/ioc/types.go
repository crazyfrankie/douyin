package ioc

import (
	"github.com/gin-gonic/gin"

	"github.com/crazyfrankie/douyin/app/user/rpc/server"
)

type App struct {
	HTTPServer *gin.Engine
	RPCServer  *server.UserServer
}
