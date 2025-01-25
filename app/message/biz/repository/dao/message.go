package dao

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Message struct {
	Id         int64 `gorm:"primarykey,autoIncrement"`
	FromUserId int64 `gorm:"index_id_time"`
	ToUserId   int64 `gorm:"index_id_time"`
	Content    string
	Ctime      int64 `gorm:"index_id_time"`
}

type MessageDao struct {
	db *gorm.DB
}

func NewMessageDao(db *gorm.DB) *MessageDao {
	return &MessageDao{db: db}
}

func (d *MessageDao) AddMessage(ctx context.Context, msg Message) error {
	now := time.Now().Unix()
	msg.Ctime = now

	return d.db.WithContext(ctx).Create(&msg).Error
}

func (d *MessageDao) GetMessageList(ctx context.Context, fromUserId, toUserId int64, preMsgTime int64) ([]Message, error) {
	var msgs []Message

	err := d.db.WithContext(ctx).Where("from_user_id = ? and to_user_id = ? and ctime > ?", fromUserId, toUserId, preMsgTime).Find(&msgs).Error
	if err != nil {
		return nil, err
	}

	return msgs, nil
}
