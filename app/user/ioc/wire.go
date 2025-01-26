//go:build wireinject

package ioc

import (
	"github.com/crazyfrankie/douyin/app/user/biz/repository"
	"github.com/crazyfrankie/douyin/app/user/biz/repository/dao"
	"github.com/crazyfrankie/douyin/app/user/biz/rpc"
	"github.com/crazyfrankie/douyin/app/user/biz/rpc/client"
	"github.com/crazyfrankie/douyin/app/user/biz/rpc/server"
	"github.com/crazyfrankie/douyin/app/user/biz/service"
	"github.com/google/wire"
)

func InitApp() *App {
	wire.Build(
		InitDB,
		dao.NewUserDao,
		repository.NewUserRepo,
		service.NewUserService,

		client.InitFavoriteClient,
		client.InitPublishClient,
		client.InitRelationClient,

		server.NewUserServer,
		rpc.NewUserRPCServer,

		wire.Struct(new(App), "*"),
	)
	return new(App)
}
