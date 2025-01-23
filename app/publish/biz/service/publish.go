package service

import (
	"context"
	"encoding/json"
	"mime/multipart"

	"github.com/crazyfrankie/douyin/app/publish/biz/repository"
	"github.com/crazyfrankie/douyin/app/publish/biz/repository/dao"
	"github.com/crazyfrankie/douyin/rpc_gen/publish"
)

type PublishService struct {
	repo *repository.PublishRepo
}

func NewPublishService(repo *repository.PublishRepo) *PublishService {
	return &PublishService{
		repo: repo,
	}
}

func (s *PublishService) PublishAction(ctx context.Context, req *publish.PublishActionRequest) error {
	var fileHeader multipart.FileHeader
	err := json.Unmarshal(req.Data, &fileHeader)
	if err != nil {
		return err
	}

	video := &dao.Video{}
	return s.repo.AddVideo(ctx, video)
}

func (s *PublishService) PublishList(ctx context.Context, uid int64) ([]dao.Video, error) {
	return nil, nil
}

func (s *PublishService) PublishCount(ctx context.Context, uid int64) ([]int64, error) {
	return s.repo.GetUserPublishCount(ctx, uid)
}
