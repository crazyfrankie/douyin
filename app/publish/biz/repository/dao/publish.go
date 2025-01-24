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
	Utime    int64 `gorm:"index:utime_index"`
}

type PublishDao struct {
	db *gorm.DB
}

func NewPublishDao(db *gorm.DB) *PublishDao {
	return &PublishDao{db: db}
}

func (d *PublishDao) CreateVideo(ctx context.Context, video *Video) error {
	return d.db.WithContext(ctx).Create(video).Error
}

func (d *PublishDao) GetPublishVideos(ctx context.Context, uid int64) ([]int64, error) {
	var videos []Video
	err := d.db.WithContext(ctx).Where("author_id = ?", uid).Find(&videos).Error
	if err != nil {
		return nil, err
	}

	res := make([]int64, 0, len(videos))
	for _, v := range videos {
		res = append(res, v.ID)
	}

	return res, nil
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

func (d *PublishDao) GetPublishVideoInfo(ctx context.Context, vid int64) (Video, error) {
	var video Video
	err := d.db.WithContext(ctx).Where("id = ?", vid).Find(&video).Error
	if err != nil {
		return Video{}, err
	}

	return video, nil
}
