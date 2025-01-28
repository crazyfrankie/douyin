package tencent

import (
	"context"

	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

type TencentSms struct {
	client   *sms.Client
	appId    *string
	signName *string
}

func (t *TencentSms) SendSms(ctx context.Context, biz string, args []string, numbers ...string) error {
	//TODO implement me
	panic("implement me")
}
