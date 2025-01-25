package repository

import (
	"context"

	"github.com/crazyfrankie/douyin/app/message/biz/repository/dao"
)

type MessageRepo struct {
	dao *dao.MessageDao
}

func NewMessageRepo(dao *dao.MessageDao) *MessageRepo {
	return &MessageRepo{dao: dao}
}

func (r *MessageRepo) AddMessage(ctx context.Context, msg dao.Message) error {
	return r.dao.AddMessage(ctx, msg)
}

func (r *MessageRepo) GetMessageList(ctx context.Context, fromUserId, toUserId int64, preMsgTime int64) ([]dao.Message, error) {
	return r.dao.GetMessageList(ctx, fromUserId, toUserId, preMsgTime)
}
