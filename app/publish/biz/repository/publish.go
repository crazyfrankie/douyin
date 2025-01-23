package repository

import (
	"context"

	"github.com/crazyfrankie/douyin/app/publish/biz/repository/dao"
)

type PublishRepo struct {
	dao *dao.PublishDao
}

func NewPublishRepo(dao *dao.PublishDao) *PublishRepo {
	return &PublishRepo{dao: dao}
}

func (r *PublishRepo) AddVideo(ctx context.Context, video *dao.Video) error {
	return r.dao.CreateVideo(ctx, video)
}

func (r *PublishRepo) GetUserPublishCount(ctx context.Context, uid int64) ([]int64, error) {
	return r.dao.GetUserPublishCount(ctx, uid)
}
