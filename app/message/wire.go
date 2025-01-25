//go:build wireinject

package main

import (
	"github.com/crazyfrankie/douyin/app/message/biz/repository"
	"github.com/crazyfrankie/douyin/app/message/biz/repository/dao"
	"github.com/crazyfrankie/douyin/app/message/biz/rpc"
	"github.com/crazyfrankie/douyin/app/message/biz/rpc/server"
	"github.com/crazyfrankie/douyin/app/message/biz/service"
	"github.com/crazyfrankie/douyin/app/message/ioc"
	"github.com/google/wire"
)

func InitApp() *ioc.App {
	wire.Build(
		ioc.InitDB,
		dao.NewMessageDao,
		repository.NewMessageRepo,
		service.NewMessageService,
		server.NewMessageServer,
		rpc.NewMessageRPCServer,

		wire.Struct(new(ioc.App), "*"),
	)
	return new(ioc.App)
}
