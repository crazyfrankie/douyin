package cache

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/crazyfrankie/douyin/app/sms/common/errno"
)

var (
	//go:embed lua/set_code.lua
	luaSetCode string
	//go:embed lua/verify_code.lua
	luaVerifyCode string
)

type CodeCache struct {
	cmd redis.Cmdable
}

func NewCodeCache(cmd redis.Cmdable) *CodeCache {
	return &CodeCache{cmd: cmd}
}

func (c *CodeCache) Store(ctx context.Context, biz, phone, code string) error {
	key := c.key(biz, phone)

	res, err := c.cmd.Eval(ctx, luaSetCode, []string{key}, code).Int()

	if err != nil {
		return err
	}

	switch res {
	case 0:
		// 毫无问题
		return nil
	case -1:
		// 发送太频繁
		return errno.SendTooMany
	default:
		return errno.InternalServer
	}
}

func (c *CodeCache) Verify(ctx context.Context, biz, phone, inputCode string) error {
	key := c.key(biz, phone)

	ok, err := c.cmd.Eval(ctx, luaVerifyCode, []string{key}, inputCode).Int()
	if err != nil {
		return err
	}
	switch ok {
	case 0:
		return nil
	case -1:
		return errno.VerifyTooMany
	}
	return errno.InternalServer
}

func (c *CodeCache) key(biz, phone string) string {
	return fmt.Sprintf("phone_code:%s:%s", biz, phone)
}
