//go:build wireinject

package ioc

import (
	"github.com/crazyfrankie/douyin/app/comment/biz/repository"
	"github.com/crazyfrankie/douyin/app/comment/biz/repository/dao"
	"github.com/crazyfrankie/douyin/app/comment/biz/rpc"
	"github.com/crazyfrankie/douyin/app/comment/biz/rpc/client"
	"github.com/crazyfrankie/douyin/app/comment/biz/rpc/server"
	"github.com/crazyfrankie/douyin/app/comment/biz/service"
	"github.com/google/wire"
)

func InitApp() *App {
	wire.Build(
		InitDB,
		dao.NewCommentDao,
		repository.NewCommentRepo,
		InitFilter,
		service.NewCommentService,

		client.InitUserClient,
		server.NewCommentServer,

		rpc.NewCommentRPCServer,

		wire.Struct(new(App), "*"),
	)
	return new(App)
}
