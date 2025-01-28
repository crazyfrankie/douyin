//go:build wireinject

package ioc

import (
	"github.com/crazyfrankie/douyin/app/feed/biz/repository"
	"github.com/crazyfrankie/douyin/app/feed/biz/repository/dao"
	"github.com/crazyfrankie/douyin/app/feed/biz/rpc"
	"github.com/crazyfrankie/douyin/app/feed/biz/rpc/client"
	"github.com/crazyfrankie/douyin/app/feed/biz/rpc/server"
	"github.com/crazyfrankie/douyin/app/feed/biz/service"
	"github.com/google/wire"
)

func InitApp() *App {
	wire.Build(
		InitDB,
		dao.NewFeedDao,
		repository.NewFeedRepo,
		service.NewFeedService,

		client.InitUserClient,
		client.InitFavoriteClient,
		client.InitCommentClient,
		server.NewVideoServer,

		rpc.NewFeedRPCServer,

		wire.Struct(new(App), "*"),
	)
	return new(App)
}
