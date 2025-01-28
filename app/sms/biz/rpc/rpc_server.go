package rpc

import (
	"github.com/crazyfrankie/douyin/app/sms/config"
	"google.golang.org/grpc"
	"net"

	"github.com/crazyfrankie/douyin/app/sms/biz/rpc/server"
)

type Server struct {
	*grpc.Server
	Addr string
}

func NewSmsGRPCServer(s *server.SmsServer) *Server {
	sr := grpc.NewServer()
	s.RegisterServer(sr)

	return &Server{
		Server: sr,
		Addr:   config.GetConf().RPC.Address,
	}
}

func (s *Server) Serve() error {
	conn, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}

	return s.Server.Serve(conn)
}
