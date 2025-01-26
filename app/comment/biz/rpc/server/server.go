package server

import (
	"context"
	"github.com/crazyfrankie/douyin/app/comment/mw"

	"google.golang.org/grpc"

	"github.com/crazyfrankie/douyin/app/comment/biz/service"
	"github.com/crazyfrankie/douyin/rpc_gen/comment"
)

type CommentServer struct {
	comment.UnimplementedCommentServiceServer
	svc *service.CommentService
}

func NewCommentServer(svc *service.CommentService) *CommentServer {
	return &CommentServer{svc: svc}
}

func (s *CommentServer) RegisterServer(server *grpc.Server) {
	comment.RegisterCommentServiceServer(server, s)
}
func (s *CommentServer) CommentAction(ctx context.Context, request *comment.CommentActionRequest) (*comment.CommentActionResponse, error) {
	resp, err := s.svc.CommentAction(ctx, request)
	if err != nil {
		return nil, err
	}

	return &comment.CommentActionResponse{
		Comment: resp,
	}, nil
}

func (s *CommentServer) CommentList(ctx context.Context, request *comment.CommentListRequest) (*comment.CommentListResponse, error) {
	claims, err := mw.ParseToken(request.GetToken())
	if err != nil {
		return nil, err
	}

	newCtx := context.WithValue(ctx, "user_id", claims["user_id"])
	resp, err := s.svc.CommentList(newCtx, request)
	if err != nil {
		return nil, err
	}

	return &comment.CommentListResponse{
		CommentList: resp,
	}, nil
}
