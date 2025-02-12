package repository

import (
	"context"

	"github.com/crazyfrankie/douyin/app/sms/biz/repository/cache"
)

type CodeRepo struct {
	cache *cache.CodeCache
}

func NewCodeRepo(cache *cache.CodeCache) *CodeRepo {
	return &CodeRepo{cache: cache}
}

func (r *CodeRepo) Store(ctx context.Context, biz, phone, code string) error {
	return r.cache.Store(ctx, biz, phone, code)
}

func (r *CodeRepo) Verify(ctx context.Context, biz, phone, code string) error {
	return r.cache.Verify(ctx, biz, phone, code)
}
