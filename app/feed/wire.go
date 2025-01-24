//go:build wireinject

package main

import (
	"github.com/crazyfrankie/douyin/app/feed/biz/repository"
	"github.com/crazyfrankie/douyin/app/feed/biz/repository/dao"
	"github.com/crazyfrankie/douyin/app/feed/biz/rpc"
	"github.com/crazyfrankie/douyin/app/feed/biz/rpc/client"
	"github.com/crazyfrankie/douyin/app/feed/biz/rpc/server"
	"github.com/crazyfrankie/douyin/app/feed/biz/service"
	"github.com/crazyfrankie/douyin/app/feed/ioc"
	"github.com/google/wire"
)

func InitApp() *ioc.App {
	wire.Build(
		ioc.InitDB,
		dao.NewFeedDao,
		repository.NewFeedRepo,
		service.NewFeedService,

		client.InitUserClient,
		client.InitFavoriteClient,
		server.NewVideoServer,

		rpc.NewFeedRPCServer,

		wire.Struct(new(ioc.App), "*"),
	)
	return new(ioc.App)
}
