package server

import (
	"context"

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
	err := m.svc.MessageAction(ctx, request)
	if err != nil {
		return nil, err
	}

	return &message.MessageActionResponse{}, nil
}

func (m *MessageServer) MessageChat(ctx context.Context, request *message.MessageChatRequest) (*message.MessageChatResponse, error) {
	resp, err := m.svc.MessageChat(ctx, request)
	if err != nil {
		return nil, err
	}

	return &message.MessageChatResponse{
		MessageList: resp,
	}, nil
}
