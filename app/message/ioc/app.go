package ioc

import "github.com/crazyfrankie/douyin/app/message/biz/rpc"

type App struct {
	RPCServer *rpc.Server
}
