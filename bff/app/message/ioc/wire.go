//go:build wireinject

package ioc

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"github.com/crazyfrankie/douyin/bff/app/message/handler"
)

func InitGin() *gin.Engine {
	wire.Build(
		InitClient,
		handler.NewHandler,
		InitMiddlewares,
		InitWeb,
	)
	return new(gin.Engine)
}
