//go:build wireinject

package main

import (
	"github.com/crazyfrankie/douyin/app/relation/biz/repository"
	"github.com/crazyfrankie/douyin/app/relation/biz/repository/dao"
	"github.com/crazyfrankie/douyin/app/relation/biz/rpc"
	"github.com/crazyfrankie/douyin/app/relation/biz/rpc/client"
	"github.com/crazyfrankie/douyin/app/relation/biz/rpc/server"
	"github.com/crazyfrankie/douyin/app/relation/biz/service"
	"github.com/crazyfrankie/douyin/app/relation/ioc"
	"github.com/google/wire"
)

func InitApp() *ioc.App {
	wire.Build(
		ioc.InitDB,
		dao.NewRelationDao,
		repository.NewRelationRepo,
		service.NewRelationService,

		client.InitUserClient,
		server.NewRelationServer,
		rpc.NewRelationRPCServer,

		wire.Struct(new(ioc.App), "*"),
	)
	return new(ioc.App)
}
