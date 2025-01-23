//go:build wireinject

package ioc

import (
	"github.com/crazyfrankie/douyin/bff/app/favorite/handler"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
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
