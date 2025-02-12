package cache

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RelationCache struct {
	cmd redis.Cmdable
}

func NewRelationCache(cmd redis.Cmdable) *RelationCache {
	return &RelationCache{cmd: cmd}
}

func (c *RelationCache) AddFollow(ctx context.Context, uid, toUid int64) error {
	key := c.followKey(uid)

	return c.cmd.SAdd(ctx, key, toUid).Err()
}

func (c *RelationCache) AddFollower(ctx context.Context, uid, toUid int64) error {
	key := c.followerKey(toUid)

	return c.cmd.SAdd(ctx, key, uid).Err()
}

func (c *RelationCache) CheckFollow(ctx context.Context, uid int64) bool {
	key := c.followKey(uid)

	e, err := c.cmd.Exists(ctx, key).Result()
	if err != nil {
		return false
	}

	if e > 0 {
		return true
	}

	return false
}

func (c *RelationCache) CheckFollower(ctx context.Context, uid int64) bool {
	key := c.followerKey(uid)

	e, err := c.cmd.Exists(ctx, key).Result()
	if err != nil {
		return false
	}

	if e > 0 {
		return true
	}

	return false
}

func (c *RelationCache) ExistFollow(ctx context.Context, uid, toUid int64) bool {
	key := c.followKey(uid)

	exists, err := c.cmd.SIsMember(ctx, key, toUid).Result()
	if err != nil {
		return false
	}

	return exists
}

func (c *RelationCache) ExistFollower(ctx context.Context, uid, toUid int64) bool {
	key := c.followerKey(toUid)

	exists, err := c.cmd.SIsMember(ctx, key, uid).Result()
	if err != nil {
		return false
	}

	return exists
}

func (c *RelationCache) followKey(uid int64) string {
	return fmt.Sprintf("%d:follow", uid)
}

func (c *RelationCache) followerKey(uid int64) string {
	return fmt.Sprintf("%d:follower", uid)
}
