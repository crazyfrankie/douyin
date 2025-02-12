//go:build wireinject

package ioc

import (
	"github.com/crazyfrankie/douyin/app/relation/biz/repository"
	"github.com/crazyfrankie/douyin/app/relation/biz/repository/cache"
	"github.com/crazyfrankie/douyin/app/relation/biz/repository/dao"
	"github.com/crazyfrankie/douyin/app/relation/biz/rpc"
	"github.com/crazyfrankie/douyin/app/relation/biz/rpc/client"
	"github.com/crazyfrankie/douyin/app/relation/biz/rpc/server"
	"github.com/crazyfrankie/douyin/app/relation/biz/service"
	"github.com/google/wire"
)

func InitApp() *App {
	wire.Build(
		InitDB,
		InitRedis,
		dao.NewRelationDao,
		cache.NewRelationCache,
		repository.NewRelationRepo,
		service.NewRelationService,

		client.InitUserClient,
		server.NewRelationServer,
		rpc.NewRelationRPCServer,

		wire.Struct(new(App), "*"),
	)
	return new(App)
}
