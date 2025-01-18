package dao

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Favorite struct {
	ID      int64 `gorm:"primaryKey,autoIncrement"`
	UserID  int64 `gorm:"index:user_video_id"`
	VideoID int64 `gorm:"index:user_video_id"`
	Ctime   int64
	Utime   int64
}

type FavoriteDao struct {
	db *gorm.DB
}

func NewFavoriteDao(db *gorm.DB) *FavoriteDao {
	return &FavoriteDao{db: db}
}

func (d *FavoriteDao) AddFavorite(ctx context.Context, favorite Favorite) error {
	now := time.Now().Unix()
	favorite.Ctime = now
	favorite.Utime = now

	return d.db.WithContext(ctx).Model(&Favorite{}).Create(&favorite).Error
}

func (d *FavoriteDao) GetIsFavorite(ctx context.Context, videoId, uid int64) (bool, error) {
	var sum int64
	err := d.db.WithContext(ctx).Model(&Favorite{}).Where("video_id = ? AND user_id = ?", videoId, uid).Count(&sum).Error
	if err != nil {
		return false, err
	}
	if sum == 0 {
		return false, nil
	}

	return true, nil
}

func (d *FavoriteDao) DelFavorite(ctx context.Context, favorite Favorite) error {
	return d.db.WithContext(ctx).Model(&Favorite{}).
		Where("video_id = ? AND user_id = ?", favorite.VideoID, favorite.UserID).
		Delete(&favorite).Error
}

func (d *FavoriteDao) GetFavoriteVideosByID(ctx context.Context, uid int64) ([]int64, error) {
	var result []int64
	var favorites []Favorite
	err := d.db.WithContext(ctx).Model(&Favorite{}).Where("user_id = ?", uid).Find(&favorites).Error
	if err != nil {
		return nil, err
	}

	for _, f := range favorites {
		result = append(result, f.VideoID)
	}

	return result, nil
}

func (d *FavoriteDao) GetVideoFavoriteCount(ctx context.Context, vid int64) (int64, error) {
	var sum int64
	err := d.db.WithContext(ctx).Model(&Favorite{}).Where("video_id = ?", vid).Count(&sum).Error
	if err != nil {
		return -1, err
	}

	return sum, nil
}

func (d *FavoriteDao) GetUserFavoriteCount(ctx context.Context, uid int64) (int64, error) {
	var sum int64
	err := d.db.WithContext(ctx).Model(&Favorite{}).Where("user_id = ?", uid).Count(&sum).Error
	if err != nil {
		return -1, err
	}

	return sum, nil
}
