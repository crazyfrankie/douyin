package rpc

import (
	"github.com/crazyfrankie/douyin/app/favorite/biz/rpc/server"
	"github.com/crazyfrankie/douyin/app/favorite/config"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	*grpc.Server
	Addr string
}

func NewFavoriteRPCServer(f *server.FavoriteServer) *Server {
	s := grpc.NewServer()
	f.RegisterServer(s)

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
