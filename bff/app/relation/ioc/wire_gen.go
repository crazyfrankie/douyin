// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package ioc

import (
	"github.com/crazyfrankie/douyin/bff/app/relation/handler"
	"github.com/gin-gonic/gin"
)

// Injectors from wire.go:

func InitGin() *gin.Engine {
	v := InitMiddlewares()
	relationServiceClient := InitClient()
	handlerHandler := handler.NewHandler(relationServiceClient)
	engine := InitWeb(v, handlerHandler)
	return engine
}
