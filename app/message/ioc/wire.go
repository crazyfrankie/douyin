//go:build wireinject

package ioc

import (
	"github.com/crazyfrankie/douyin/app/message/biz/repository"
	"github.com/crazyfrankie/douyin/app/message/biz/repository/dao"
	"github.com/crazyfrankie/douyin/app/message/biz/rpc"
	"github.com/crazyfrankie/douyin/app/message/biz/rpc/server"
	"github.com/crazyfrankie/douyin/app/message/biz/service"
	"github.com/google/wire"
)

func InitApp() *App {
	wire.Build(
		InitDB,
		dao.NewMessageDao,
		repository.NewMessageRepo,
		service.NewMessageService,
		server.NewMessageServer,
		rpc.NewMessageRPCServer,

		wire.Struct(new(App), "*"),
	)
	return new(App)
}
