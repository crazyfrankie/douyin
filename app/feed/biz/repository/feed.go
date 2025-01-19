package repository

import (
	"context"
	"github.com/crazyfrankie/douyin/app/feed/biz/repository/dao"
)

type FeedRepo struct {
	dao *dao.FeedDao
}

func NewFeedRepo(dao *dao.FeedDao) *FeedRepo {
	return &FeedRepo{dao: dao}
}

func (r *FeedRepo) VideoList(ctx context.Context, vid []int64) ([]dao.Video, error) {
	return r.dao.VideoList(ctx, vid)
}

func (r *FeedRepo) QueryVideoExists(ctx context.Context, vid int64) (bool, error) {
	return r.dao.QueryVideoExistsByID(ctx, vid)
}

func (r *FeedRepo) GetVideoByLastTime(ctx context.Context, time int64) (interface{}, interface{}) {

}
