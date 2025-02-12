package repository

import (
	"context"

	"github.com/crazyfrankie/douyin/app/user/biz/repository/dao"
)

type UserRepo struct {
	dao *dao.UserDao
}

func NewUserRepo(dao *dao.UserDao) *UserRepo {
	return &UserRepo{dao: dao}
}

func (repo *UserRepo) CreateUser(ctx context.Context, user dao.User) (int64, error) {
	return repo.dao.Create(ctx, user)
}

func (repo *UserRepo) FindByName(ctx context.Context, name string) (dao.User, error) {
	return repo.dao.FindUserByName(ctx, name)
}

func (repo *UserRepo) FindByID(ctx context.Context, uid int64) (dao.User, error) {
	return repo.dao.FindUserByID(ctx, uid)
}

func (repo *UserRepo) FindByPhone(ctx context.Context, phone string) (dao.User, error) {
	return repo.dao.FindUserByPhone(ctx, phone)
}
