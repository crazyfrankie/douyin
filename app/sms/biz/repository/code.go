package repository

import (
	"github.com/crazyfrankie/douyin/app/sms/biz/repository/cache"
)

type CodeRepo struct {
	cache *cache.CodeCache
}

func NewCodeRepo(cache *cache.CodeCache) *CodeRepo {
	return &CodeRepo{cache: cache}
}
