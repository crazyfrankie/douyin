package cache

import "github.com/redis/go-redis/v9"

type CodeCache struct {
	cmd redis.Cmdable
}

func NewCodeCache(cmd redis.Cmdable) *CodeCache {
	return &CodeCache{cmd: cmd}
}
