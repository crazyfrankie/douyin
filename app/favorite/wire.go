//go:build wireinject

package main

import (
	"github.com/crazyfrankie/douyin/app/favorite/biz/handler"
	"github.com/crazyfrankie/douyin/app/favorite/biz/repository"
	"github.com/crazyfrankie/douyin/app/favorite/biz/repository/dao"
	"github.com/crazyfrankie/douyin/app/favorite/biz/service"
	"github.com/crazyfrankie/douyin/app/favorite/ioc"
	"github.com/crazyfrankie/douyin/app/favorite/rpc"
	"github.com/crazyfrankie/douyin/app/favorite/rpc/client"
	"github.com/crazyfrankie/douyin/app/favorite/rpc/server"

	"github.com/google/wire"
)

func InitApp() *ioc.App {
	wire.Build(
		ioc.InitDB,
		dao.NewFavoriteDao,
		repository.NewFavoriteRepo,
		service.NewFavoriteService,
		handler.NewFavoriteHandler,

		client.NewFeedClient,
		server.NewFavoriteServer,

		rpc.NewFavoriteRPCServer,
		ioc.InitWeb,
		
		wire.Struct(new(ioc.App), "*"),
	)
	return new(ioc.App)
}
