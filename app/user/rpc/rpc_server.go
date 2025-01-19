package rpc

import (
	"net"

	"google.golang.org/grpc"
	
	"github.com/crazyfrankie/douyin/app/user/config"
	"github.com/crazyfrankie/douyin/app/user/rpc/server"
)

type Server struct {
	*grpc.Server
	Addr string
}

func NewUserRPCServer(u *server.UserServer) *Server {
	s := grpc.NewServer()
	u.RegisterServer(s)

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
