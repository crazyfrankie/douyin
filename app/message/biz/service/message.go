package service

import (
	"context"

	"github.com/crazyfrankie/douyin/app/message/biz/repository"
	"github.com/crazyfrankie/douyin/app/message/biz/repository/dao"
	"github.com/crazyfrankie/douyin/rpc_gen/message"
)

type MessageService struct {
	repo *repository.MessageRepo
}

func NewMessageService(repo *repository.MessageRepo) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) MessageAction(ctx context.Context, req *message.MessageActionRequest) error {
	userId := ctx.Value("user_id").(float64)
	uid := int64(userId)

	msg := dao.Message{
		FromUserId: uid,
		ToUserId:   req.ToUserId,
		Content:    req.Content,
	}

	return s.repo.AddMessage(ctx, msg)
}

func (s *MessageService) MessageChat(ctx context.Context, req *message.MessageChatRequest) ([]*message.Message, error) {
	userId := ctx.Value("user_id").(float64)
	uid := int64(userId)
	toUid, preMsgTime := req.GetToUserId(), req.GetPreMsgTime()

	var messages []*message.Message
	dbMessages, err := s.repo.GetMessageList(ctx, uid, toUid, preMsgTime)
	if err != nil {
		return nil, err
	}

	for _, dbMsg := range dbMessages {
		msg := &message.Message{
			Id:         dbMsg.Id,
			ToUserId:   dbMsg.ToUserId,
			FromUserId: dbMsg.FromUserId,
			Content:    dbMsg.Content,
			CreateTime: dbMsg.Ctime,
		}
		messages = append(messages, msg)
	}

	return messages, nil
}
