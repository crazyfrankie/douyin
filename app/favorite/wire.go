//go:build wireinject

package main

import (
	"github.com/crazyfrankie/douyin/app/favorite/biz/repository"
	"github.com/crazyfrankie/douyin/app/favorite/biz/repository/dao"
	"github.com/crazyfrankie/douyin/app/favorite/biz/rpc"
	"github.com/crazyfrankie/douyin/app/favorite/biz/rpc/client"
	"github.com/crazyfrankie/douyin/app/favorite/biz/rpc/server"
	"github.com/crazyfrankie/douyin/app/favorite/biz/service"
	"github.com/crazyfrankie/douyin/app/favorite/ioc"
	"github.com/google/wire"
)

func InitApp() *ioc.App {
	wire.Build(
		ioc.InitDB,
		dao.NewFavoriteDao,
		repository.NewFavoriteRepo,
		service.NewFavoriteService,

		client.NewFeedClient,
		server.NewFavoriteServer,

		rpc.NewFavoriteRPCServer,

		wire.Struct(new(ioc.App), "*"),
	)
	return new(ioc.App)
}
