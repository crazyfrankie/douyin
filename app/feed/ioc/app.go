package ioc

import (
	"github.com/crazyfrankie/douyin/app/feed/biz/rpc"
)

type App struct {
	RPCServer *rpc.Server
}
