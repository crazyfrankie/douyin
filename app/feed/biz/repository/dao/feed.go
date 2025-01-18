package dao

import (
	"context"
	
	"gorm.io/gorm"
)

type Video struct {
	ID       int64 `gorm:"primaryKey,autoIncrement"`
	AuthorID int64 `gorm:"unique"`
	PlayURL  string
	CoverURL string
	Title    string
	Ctime    int64
	Utime    int64
}

type FeedDao struct {
	db *gorm.DB
}

func NewFeedDao(db *gorm.DB) *FeedDao {
	return &FeedDao{db: db}
}

//func (d *VideoDao) CreateVideo(ctx context.Context, video model.Video) error {
//
//}

func (d *FeedDao) VideoList(ctx context.Context, vid []int64) ([]Video, error) {
	var videos []Video
	for _, id := range vid {
		var video Video
		err := d.db.WithContext(ctx).Model(&Video{}).Where("id = ?", id).First(&video).Error
		if err != nil {
			return videos, err
		}
		videos = append(videos, video)
	}

	return videos, nil
}

func (d *FeedDao) QueryVideoExistsByID(ctx context.Context, vid int64) (bool, error) {
	var video Video
	err := d.db.WithContext(ctx).Model(&Video{}).Where("id = ?", vid).Find(&video).Error
	if err != nil {
		return false, err
	}
	if video == (Video{}) {
		return false, nil
	}

	return true, nil
}
