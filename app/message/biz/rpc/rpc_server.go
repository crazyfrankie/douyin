package rpc

import (
	"google.golang.org/grpc"
	"net"

	"github.com/crazyfrankie/douyin/app/message/biz/rpc/server"
	"github.com/crazyfrankie/douyin/app/message/config"
)

type Server struct {
	*grpc.Server
	Addr string
}

func NewMessageRPCServer(c *server.MessageServer) *Server {
	s := grpc.NewServer()
	c.RegisterServer(s)

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
