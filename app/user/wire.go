//go:build wireinject

package main

import (
	"github.com/crazyfrankie/douyin/app/user/biz/handler"
	"github.com/crazyfrankie/douyin/app/user/biz/repository"
	"github.com/crazyfrankie/douyin/app/user/biz/repository/dao"
	"github.com/crazyfrankie/douyin/app/user/biz/service"
	"github.com/crazyfrankie/douyin/app/user/ioc"
	"github.com/crazyfrankie/douyin/app/user/rpc"
	"github.com/crazyfrankie/douyin/app/user/rpc/client"
	"github.com/crazyfrankie/douyin/app/user/rpc/server"
	"github.com/google/wire"
)

func InitApp() *ioc.App {
	wire.Build(
		ioc.InitDB,
		dao.NewUserDao,
		repository.NewUserRepo,
		service.NewUserService,
		handler.NewHandler,

		client.NewFavoriteClient,

		server.NewUserServer,
		rpc.NewUserRPCServer,
		ioc.InitWeb,

		wire.Struct(new(ioc.App), "*"),
	)
	return new(ioc.App)
}
