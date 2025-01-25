package repository

import (
	"context"
	"github.com/crazyfrankie/douyin/app/comment/biz/repository/dao"
)

type CommentRepo struct {
	dao *dao.CommentDao
}

func NewCommentRepo(dao *dao.CommentDao) *CommentRepo {
	return &CommentRepo{dao: dao}
}

func (r *CommentRepo) CreateComment(ctx context.Context, comment *dao.Comment) error {
	return r.dao.CreateComment(ctx, comment)
}

func (r *CommentRepo) DeleteComment(ctx context.Context, id int64) error {
	return r.dao.DeleteComment(ctx, id)
}

func (r *CommentRepo) GetCommentList(ctx context.Context, vid int64) ([]dao.Comment, error) {
	return r.dao.GetCommentList(ctx, vid)
}
