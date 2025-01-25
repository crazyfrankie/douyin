package rpc

import (
	"google.golang.org/grpc"
	"net"

	"github.com/crazyfrankie/douyin/app/comment/biz/rpc/server"
	"github.com/crazyfrankie/douyin/app/comment/config"
)

type Server struct {
	*grpc.Server
	Addr string
}

func NewCommentRPCServer(c *server.CommentServer) *Server {
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
