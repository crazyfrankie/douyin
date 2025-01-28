package memory

import (
	"context"
	"fmt"

	"github.com/crazyfrankie/douyin/app/sms/biz/service"
)

type MemorySms struct {
}

func NewMemorySms() service.SendService {
	return &MemorySms{}
}

func (m *MemorySms) SendSms(ctx context.Context, biz string, args []string, numbers ...string) error {
	fmt.Println(args)

	return nil
}
