//go:build wireinject

package main

import (
	"github.com/crazyfrankie/douyin/app/comment/biz/repository"
	"github.com/crazyfrankie/douyin/app/comment/biz/repository/dao"
	"github.com/crazyfrankie/douyin/app/comment/biz/rpc"
	"github.com/crazyfrankie/douyin/app/comment/biz/rpc/client"
	"github.com/crazyfrankie/douyin/app/comment/biz/rpc/server"
	"github.com/crazyfrankie/douyin/app/comment/biz/service"
	"github.com/crazyfrankie/douyin/app/comment/ioc"
	"github.com/google/wire"
)

func InitApp() *ioc.App {
	wire.Build(
		ioc.InitDB,
		dao.NewCommentDao,
		repository.NewCommentRepo,
		ioc.InitFilter,
		service.NewCommentService,

		client.InitUserClient,
		server.NewCommentServer,

		rpc.NewCommentRPCServer,

		wire.Struct(new(ioc.App), "*"),
	)
	return new(ioc.App)
}
