package rpc

import (
	"net"

	"google.golang.org/grpc"

	"github.com/crazyfrankie/douyin/app/feed/config"
	"github.com/crazyfrankie/douyin/app/feed/rpc/server"
)

type Server struct {
	*grpc.Server
	Addr string
}

func NewFeedRPCServer(v *server.VideoServer) *Server {
	s := grpc.NewServer()
	v.RegisterServer(s)

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
