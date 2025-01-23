package ioc

import (
	"github.com/crazyfrankie/douyin/app/favorite/biz/rpc"
)

type App struct {
	RPCServer *rpc.Server
}
