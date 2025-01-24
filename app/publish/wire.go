//go:build wireinject

package main

import (
	"github.com/crazyfrankie/douyin/app/publish/biz/repository"
	"github.com/crazyfrankie/douyin/app/publish/biz/repository/dao"
	"github.com/crazyfrankie/douyin/app/publish/biz/rpc"
	"github.com/crazyfrankie/douyin/app/publish/biz/rpc/client"
	"github.com/crazyfrankie/douyin/app/publish/biz/rpc/server"
	"github.com/crazyfrankie/douyin/app/publish/biz/service"
	"github.com/crazyfrankie/douyin/app/publish/ioc"
	"github.com/google/wire"
)

func InitApp() *ioc.App {
	wire.Build(
		ioc.InitDB,
		dao.NewPublishDao,
		repository.NewPublishRepo,
		service.NewPublishService,
		server.NewPublishServer,
		client.InitFavoriteClient,
		rpc.NewPublishRPCServer,

		wire.Struct(new(ioc.App), "*"),
	)
	return new(ioc.App)
}
