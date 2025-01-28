//go:build wireinject

package ioc

import (
	"github.com/crazyfrankie/douyin/app/sms/biz/repository/cache"
	"github.com/google/wire"

	"github.com/crazyfrankie/douyin/app/sms/biz/repository"
	"github.com/crazyfrankie/douyin/app/sms/biz/rpc"
	"github.com/crazyfrankie/douyin/app/sms/biz/rpc/server"
	"github.com/crazyfrankie/douyin/app/sms/biz/service"
	"github.com/crazyfrankie/douyin/app/sms/biz/service/memory"
)

func InitSms() []service.SendService {
	my := memory.NewMemorySms()

	return []service.SendService{my}
}

func InitApp() *App {
	wire.Build(
		InitRedis,
		cache.NewCodeCache,
		repository.NewCodeRepo,
		InitSms,
		service.NewSmsService,

		server.NewSmsServer,
		rpc.NewSmsGRPCServer,

		wire.Struct(new(App), "*"),
	)

	return new(App)
}
