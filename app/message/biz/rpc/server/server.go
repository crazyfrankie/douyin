package server

import (
	"context"
	"github.com/crazyfrankie/douyin/app/message/mw"

	"google.golang.org/grpc"

	"github.com/crazyfrankie/douyin/app/message/biz/service"
	"github.com/crazyfrankie/douyin/rpc_gen/message"
)

type MessageServer struct {
	message.UnimplementedMessageServiceServer
	svc *service.MessageService
}

func NewMessageServer(svc *service.MessageService) *MessageServer {
	return &MessageServer{svc: svc}
}

func (m *MessageServer) RegisterServer(server *grpc.Server) {
	message.RegisterMessageServiceServer(server, m)
}

func (m *MessageServer) MessageAction(ctx context.Context, request *message.MessageActionRequest) (*message.MessageActionResponse, error) {
	claims, err := mw.ParseToken(request.GetToken())
	if err != nil {
		return nil, err
	}

	newCtx := context.WithValue(ctx, "user_id", claims["user_id"])

	err = m.svc.MessageAction(newCtx, request)
	if err != nil {
		return nil, err
	}

	return &message.MessageActionResponse{}, nil
}

func (m *MessageServer) MessageChat(ctx context.Context, request *message.MessageChatRequest) (*message.MessageChatResponse, error) {
	claims, err := mw.ParseToken(request.GetToken())
	if err != nil {
		return nil, err
	}

	newCtx := context.WithValue(ctx, "user_id", claims["user_id"])

	resp, err := m.svc.MessageChat(newCtx, request)
	if err != nil {
		return nil, err
	}

	return &message.MessageChatResponse{
		MessageList: resp,
	}, nil
}
