package repository

import (
	"context"

	"github.com/crazyfrankie/douyin/app/relation/biz/repository/cache"
	"github.com/crazyfrankie/douyin/app/relation/biz/repository/dao"
)

type RelationRepo struct {
	dao   *dao.RelationDao
	cache *cache.RelationCache
}

func NewRelationRepo(dao *dao.RelationDao, cache *cache.RelationCache) *RelationRepo {
	return &RelationRepo{dao: dao, cache: cache}
}

func (r *RelationRepo) GetFollowExists(ctx context.Context, uid, toUid int64) (bool, error) {
	if r.cache.CheckFollow(ctx, uid) {
		return r.cache.ExistFollow(ctx, uid, toUid), nil
	}
	if r.cache.CheckFollower(ctx, toUid) {
		return r.cache.ExistFollower(ctx, uid, toUid), nil
	}

	return r.dao.GetFollowExists(ctx, uid, toUid)
}

func (r *RelationRepo) AddFollow(ctx context.Context, relation dao.Relation) error {
	err := r.dao.AddFollow(ctx, relation)
	if err != nil {
		return err
	}

	if !r.cache.CheckFollow(ctx, relation.UserId) {
		err = r.cache.AddFollow(ctx, relation.UserId, relation.ToUserId)
	}
	if !r.cache.CheckFollower(ctx, relation.ToUserId) {
		err = r.cache.AddFollower(ctx, relation.UserId, relation.ToUserId)
	}

	return err
}

func (r *RelationRepo) DelFollow(ctx context.Context, relation dao.Relation) error {
	return r.dao.DelFollow(ctx, relation)
}

func (r *RelationRepo) GetFollowList(ctx context.Context, uid int64) ([]dao.Relation, error) {
	return r.dao.GetFollowList(ctx, uid)
}

func (r *RelationRepo) GetFollowerList(ctx context.Context, uid int64) ([]dao.Relation, error) {
	return r.dao.GetFollowerList(ctx, uid)
}

func (r *RelationRepo) GetFriendList(ctx context.Context, uid int64) ([]int64, error) {
	return r.dao.GetFriendList(ctx, uid)
}

func (r *RelationRepo) GetIsFollow(ctx context.Context, uid, toUid int64) (bool, error) {
	return r.dao.GetIsFollow(ctx, uid, toUid)
}

func (r *RelationRepo) GetFollowCount(ctx context.Context, toUid int64) (int64, int64, error) {
	return r.dao.GetFollowCount(ctx, toUid)
}
