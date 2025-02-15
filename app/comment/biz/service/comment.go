package service

import (
	"context"
	"github.com/crazyfrankie/douyin/app/comment/common/errno"
	"google.golang.org/grpc/metadata"
	"strconv"
	"time"

	"github.com/importcjj/sensitive"

	"github.com/crazyfrankie/douyin/app/comment/biz/repository"
	"github.com/crazyfrankie/douyin/app/comment/biz/repository/dao"
	"github.com/crazyfrankie/douyin/rpc_gen/comment"
	"github.com/crazyfrankie/douyin/rpc_gen/user"
)

type CommentService struct {
	repo       *repository.CommentRepo
	filter     *sensitive.Filter
	userClient user.UserServiceClient
}

func NewCommentService(repo *repository.CommentRepo, filter *sensitive.Filter, userClient user.UserServiceClient) *CommentService {
	return &CommentService{repo: repo, filter: filter, userClient: userClient}
}

func (s *CommentService) CommentAction(ctx context.Context, req *comment.CommentActionRequest) (*comment.Comment, error) {
	var com *comment.Comment
	if req.ActionType == 1 {
		// 敏感词过滤
		content := s.filter.Replace(req.GetCommentText(), '*')

		cm := &dao.Comment{
			UserID:  req.GetUserId(),
			VideoID: req.GetVideoId(),
			Content: content,
		}
		err := s.repo.CreateComment(ctx, cm)
		if err != nil {
			return com, err
		}

		com.Id = cm.ID
		com.Content = cm.Content
		com.CreateDate = time.Unix(cm.Ctime, 0).Format("2006-01-02")
		resp, err := s.userClient.GetUserInfo(ctx, &user.GetUserInfoRequest{
			UserId: cm.UserID,
		})
		if err != nil {
			return com, err
		}
		com.User = resp.GetUser()

		return com, nil
	} else {
		err := s.repo.DeleteComment(ctx, req.CommentId)
		if err != nil {
			return com, err
		}

		return com, nil
	}
}

func (s *CommentService) CommentList(ctx context.Context, req *comment.CommentListRequest) ([]*comment.Comment, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errno.ParamErr
	}
	userId := md["user_id"][0]
	uId, _ := strconv.Atoi(userId)
	uid := int64(uId)

	var commens []*comment.Comment

	dbComments, err := s.repo.GetCommentList(ctx, req.GetVideoId())
	if err != nil {
		return commens, err
	}

	for _, c := range dbComments {
		var com *comment.Comment
		err := s.commentInfo(ctx, &c, com, uid)
		if err != nil {
			return commens, err
		}

		commens = append(commens, com)
	}

	return commens, nil
}

func (s *CommentService) commentInfo(ctx context.Context, dbComment *dao.Comment, comment *comment.Comment, currUserId int64) error {
	comment.Id = dbComment.ID
	comment.Content = dbComment.Content
	comment.CreateDate = time.Unix(dbComment.Ctime, 0).Format("2006-01-02")
	resp, err := s.userClient.GetUserInfo(ctx, &user.GetUserInfoRequest{
		UserId:        currUserId,
		UserIdToQuery: dbComment.UserID,
	})
	if err != nil {
		return err
	}
	comment.User = resp.GetUser()

	return nil
}

func (s *CommentService) CommentCount(ctx context.Context, req *comment.CommentCountRequest) (int64, error) {
	count, err := s.repo.GetCommentCount(ctx, req.VideoId)
	if err != nil {
		return -1, err
	}

	return count, nil
}
