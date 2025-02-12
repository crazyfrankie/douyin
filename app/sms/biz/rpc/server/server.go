package server

import (
	"context"

	"google.golang.org/grpc"

	"github.com/crazyfrankie/douyin/app/sms/biz/service"
	"github.com/crazyfrankie/douyin/rpc_gen/sms"
)

type SmsServer struct {
	sms.UnimplementedSmsServiceServer
	svc *service.SmsService
}

func NewSmsServer(svc *service.SmsService) *SmsServer {
	return &SmsServer{svc: svc}
}

func (s *SmsServer) RegisterServer(server *grpc.Server) {
	sms.RegisterSmsServiceServer(server, s)
}

func (s *SmsServer) SendSms(ctx context.Context, request *sms.SendSmsRequest) (*sms.SendSmsResponse, error) {
	err := s.svc.SendSms(ctx, request.GetBiz(), request.GetArgs(), request.GetNumbers()...)
	if err != nil {
		return nil, err
	}

	return &sms.SendSmsResponse{}, nil
}

func (s *SmsServer) VerifySms(ctx context.Context, request *sms.VerifySmsRequest) (*sms.VerifySmsResponse, error) {
	err := s.svc.VerifySms(ctx, request.GetBiz(), request.GetNumber(), request.GetCode())
	if err != nil {
		return nil, err
	}

	return &sms.VerifySmsResponse{}, nil
}
