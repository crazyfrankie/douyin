//go:build wireinject

package main

import (
	"github.com/crazyfrankie/douyin/app/feed/biz/handler"
	"github.com/crazyfrankie/douyin/app/feed/biz/repository"
	"github.com/crazyfrankie/douyin/app/feed/biz/repository/dao"
	"github.com/crazyfrankie/douyin/app/feed/biz/service"
	"github.com/crazyfrankie/douyin/app/feed/ioc"
	"github.com/crazyfrankie/douyin/app/feed/rpc"
	"github.com/crazyfrankie/douyin/app/feed/rpc/client"
	"github.com/crazyfrankie/douyin/app/feed/rpc/server"
	"github.com/google/wire"
)

func InitApp() *ioc.App {
	wire.Build(
		ioc.InitDB,
		dao.NewFeedDao,
		repository.NewFeedRepo,
		service.NewFeedService,
		handler.NewFeedHandler,

		client.NewUserClient,
		server.NewVideoServer,

		rpc.NewFeedRPCServer,
		ioc.InitWeb,

		wire.Struct(new(ioc.App), "*"),
	)
	return new(ioc.App)
}
