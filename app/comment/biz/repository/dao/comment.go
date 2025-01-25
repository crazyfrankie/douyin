package dao

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	ID      int64 `gorm:"primaryKey,autoIncrement"`
	UserID  int64
	VideoID int64
	Content string `gorm:"type:varchar(255)"`
	Ctime   int64
	Dtime   int64 `gorm:"index"`
}

type CommentDao struct {
	db *gorm.DB
}

func NewCommentDao(db *gorm.DB) *CommentDao {
	return &CommentDao{db: db}
}

func (d *CommentDao) CreateComment(ctx context.Context, comment *Comment) error {
	now := time.Now().Unix()
	comment.Ctime = now
	comment.Dtime = 0
	err := d.db.WithContext(ctx).Create(comment).Error

	return err
}

func (d *CommentDao) DeleteComment(ctx context.Context, id int64) error {
	err := d.db.WithContext(ctx).Where("id = ?", id).Delete(&Comment{}).Error

	return err
}

func (d *CommentDao) GetCommentList(ctx context.Context, vid int64) ([]Comment, error) {
	var comments []Comment
	err := d.db.WithContext(ctx).Where("video_id = ?", vid).Find(&comments).Error

	return comments, err
}
