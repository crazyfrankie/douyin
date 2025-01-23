package ioc

import (
	"github.com/crazyfrankie/douyin/app/user/biz/rpc"
)

type App struct {
	RPCServer *rpc.Server
}
