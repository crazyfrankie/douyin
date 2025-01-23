package rpc

import (
	"net"

	"google.golang.org/grpc"

	"github.com/crazyfrankie/douyin/app/publish/biz/rpc/server"
	"github.com/crazyfrankie/douyin/app/publish/config"
)

type Server struct {
	*grpc.Server
	Addr string
}

func NewPublishRPCServer(p *server.PublishServer) *Server {
	s := grpc.NewServer()
	p.RegisterServer(s)

	return &Server{
		Server: s,
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
