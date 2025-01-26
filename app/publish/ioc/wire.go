//go:build wireinject

package ioc

import (
	"github.com/crazyfrankie/douyin/app/publish/biz/repository"
	"github.com/crazyfrankie/douyin/app/publish/biz/repository/dao"
	"github.com/crazyfrankie/douyin/app/publish/biz/rpc"
	"github.com/crazyfrankie/douyin/app/publish/biz/rpc/client"
	"github.com/crazyfrankie/douyin/app/publish/biz/rpc/server"
	"github.com/crazyfrankie/douyin/app/publish/biz/service"
	"github.com/google/wire"
)

func InitApp() *App {
	wire.Build(
		InitDB,
		dao.NewPublishDao,
		repository.NewPublishRepo,
		service.NewPublishService,
		server.NewPublishServer,
		client.InitFavoriteClient,
		rpc.NewPublishRPCServer,

		wire.Struct(new(App), "*"),
	)
	return new(App)
}
