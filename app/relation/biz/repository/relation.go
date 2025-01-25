package repository

import (
	"context"
	"github.com/crazyfrankie/douyin/app/relation/biz/repository/dao"
)

type RelationRepo struct {
	dao *dao.RelationDao
}

func NewRelationRepo(dao *dao.RelationDao) *RelationRepo {
	return &RelationRepo{dao: dao}
}

func (r *RelationRepo) GetFollowExists(ctx context.Context, uid, toUid int64) (bool, error) {
	return r.dao.GetFollowExists(ctx, uid, toUid)
}

func (r *RelationRepo) AddFollow(ctx context.Context, relation dao.Relation) error {
	return r.dao.AddFollow(ctx, relation)
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
