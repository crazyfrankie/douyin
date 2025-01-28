package ioc

import "github.com/crazyfrankie/douyin/app/sms/biz/rpc"

type App struct {
	RPCServer *rpc.Server
}
