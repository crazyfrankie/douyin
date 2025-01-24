package ioc

import (
	"github.com/crazyfrankie/douyin/app/publish/biz/rpc"
)

type App struct {
	RPCServer *rpc.Server
}
