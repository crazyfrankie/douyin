//go:build wireinject

package main

import (
	"github.com/crazyfrankie/douyin/app/user/biz/repository"
	"github.com/crazyfrankie/douyin/app/user/biz/repository/dao"
	"github.com/crazyfrankie/douyin/app/user/biz/rpc"
	"github.com/crazyfrankie/douyin/app/user/biz/rpc/client"
	"github.com/crazyfrankie/douyin/app/user/biz/rpc/server"
	"github.com/crazyfrankie/douyin/app/user/biz/service"
	"github.com/crazyfrankie/douyin/app/user/ioc"
	"github.com/google/wire"
)

func InitApp() *ioc.App {
	wire.Build(
		ioc.InitDB,
		dao.NewUserDao,
		repository.NewUserRepo,
		service.NewUserService,

		client.NewFavoriteClient,
		client.NewPublishClient,

		server.NewUserServer,
		rpc.NewUserRPCServer,

		wire.Struct(new(ioc.App), "*"),
	)
	return new(ioc.App)
}
