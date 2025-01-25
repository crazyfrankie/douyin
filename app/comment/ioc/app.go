package ioc

import "github.com/crazyfrankie/douyin/app/comment/biz/rpc"

type App struct {
	RPCServer *rpc.Server
}
