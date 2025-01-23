package dao

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Video struct {
	ID       int64 `gorm:"primaryKey,autoIncrement"`
	AuthorID int64 `gorm:"unique"`
	PlayURL  string
	CoverURL string
	Title    string
	Ctime    int64
	Utime    int64 `gorm:"index:utime_index"`
}

type PublishDao struct {
	db *gorm.DB
}

func (d *PublishDao) CreateVideo(ctx context.Context, video *Video) error {
	now := time.Now().Unix()
	video.Ctime = now
	video.Utime = now
	return d.db.WithContext(ctx).Create(video).Error
}

func (d *PublishDao) GetUserPublishCount(ctx context.Context, uid int64) ([]int64, error) {
	var videos []*Video
	err := d.db.WithContext(ctx).Where("author_id = ?", uid).Find(&videos).Error
	if err != nil {
		return nil, err
	}

	var res []int64
	for _, v := range videos {
		res = append(res, v.ID)
	}

	return res, nil
}
