package service

import (
	"context"
	"errors"
	"sync/atomic"

	"github.com/crazyfrankie/douyin/app/sms/biz/repository"
)

type SendService interface {
	SendSms(ctx context.Context, biz string, args []string, numbers ...string) error
}

type SmsService struct {
	sms  []SendService
	idx  uint64
	repo *repository.CodeRepo
}

func NewSmsService(repo *repository.CodeRepo, sms ...SendService) *SmsService {
	s := make([]SendService, 0, len(sms))
	for _, v := range sms {
		s = append(s, v)
	}
	return &SmsService{
		sms:  s,
		repo: repo,
	}
}

func (s *SmsService) SendSms(ctx context.Context, biz string, args []string, numbers ...string) error {
	err := s.repo.Store(ctx, biz, numbers[0], args[0])
	if err != nil {
		return err
	}

	// 原子操作，并发安全
	idx := atomic.AddUint64(&s.idx, 0)
	length := uint64(len(s.sms))
	for i := idx; i < length+idx; i++ {
		svc := s.sms[int(i%length)]
		err := svc.SendSms(ctx, biz, args, numbers...)
		switch {
		case err == nil:
			return nil
		case errors.Is(err, context.DeadlineExceeded), errors.Is(err, context.Canceled):
			// 调用者的超时时间到了
			// 调用者主动取消了
			return err
		}
	}
	return errors.New("all sms failed")
}

func (s *SmsService) VerifySms(ctx context.Context, biz, code, number string) error {
	return s.repo.Verify(ctx, biz, number, code)
}
