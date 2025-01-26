//go:build wireinject

package ioc

import (
	"github.com/crazyfrankie/douyin/app/favorite/biz/repository"
	"github.com/crazyfrankie/douyin/app/favorite/biz/repository/dao"
	"github.com/crazyfrankie/douyin/app/favorite/biz/rpc"
	"github.com/crazyfrankie/douyin/app/favorite/biz/rpc/client"
	"github.com/crazyfrankie/douyin/app/favorite/biz/rpc/server"
	"github.com/crazyfrankie/douyin/app/favorite/biz/service"
	"github.com/google/wire"
)

func InitApp() *App {
	wire.Build(
		InitDB,
		dao.NewFavoriteDao,
		repository.NewFavoriteRepo,
		service.NewFavoriteService,

		client.InitFeedClient,
		server.NewFavoriteServer,

		rpc.NewFavoriteRPCServer,

		wire.Struct(new(App), "*"),
	)
	return new(App)
}
