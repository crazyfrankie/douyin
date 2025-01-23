package repository

import (
	"context"

	"github.com/crazyfrankie/douyin/app/favorite/biz/repository/dao"
)

type FavoriteRepo struct {
	dao *dao.FavoriteDao
}

func NewFavoriteRepo(dao *dao.FavoriteDao) *FavoriteRepo {
	return &FavoriteRepo{dao: dao}
}

func (r *FavoriteRepo) AddFavorite(ctx context.Context, favorite dao.Favorite) error {
	return r.dao.AddFavorite(ctx, favorite)
}

func (r *FavoriteRepo) DelFavorite(ctx context.Context, favorite dao.Favorite) error {
	return r.dao.DelFavorite(ctx, favorite)
}

func (r *FavoriteRepo) GetIsFavorite(ctx context.Context, videoId, uid int64) (bool, error) {
	return r.dao.GetIsFavorite(ctx, videoId, uid)
}

func (r *FavoriteRepo) GetFavoriteVideosByID(ctx context.Context, uid int64) ([]int64, error) {
	return r.dao.GetFavoriteVideosByID(ctx, uid)
}

func (r *FavoriteRepo) GetVideoFavoriteCount(ctx context.Context, vid int64) (int64, error) {
	return r.dao.GetVideoFavoriteCount(ctx, vid)
}

func (r *FavoriteRepo) GetUserFavoriteCount(ctx context.Context, uid int64) (int64, error) {
	return r.dao.GetUserFavoriteCount(ctx, uid)
}

func (r *FavoriteRepo) GetUserFavoritedCount(ctx context.Context, vid []int64) (int64, error) {
	return r.dao.GetUserFavoritedCount(ctx, vid)
}
