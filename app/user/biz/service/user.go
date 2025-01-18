package service

import (
	"context"
	"github.com/crazyfrankie/douyin/app/user/rpc/client"
	"github.com/crazyfrankie/douyin/rpc_gen/common"
	"github.com/crazyfrankie/douyin/rpc_gen/favorite"
	"sync"

	"golang.org/x/crypto/bcrypt"

	"github.com/crazyfrankie/douyin/app/user/biz"
	"github.com/crazyfrankie/douyin/app/user/biz/repository"
	"github.com/crazyfrankie/douyin/app/user/biz/repository/dao"
	"github.com/crazyfrankie/douyin/app/user/common/errno"
)

var (
	defaultAvatar = "github.com/crazyfrankie/douyin/static/avatar/default.png"
)

type UserService struct {
	repo *repository.UserRepo
}

func NewUserService(repo *repository.UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(ctx context.Context, req biz.RegisterReq) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	u := dao.User{
		Name:     req.Name,
		Password: string(password),
		Avatar:   defaultAvatar,
	}
	var uid int64
	uid, err = s.repo.CreateUser(ctx, u)
	if err != nil {
		if uid == -1 {
			return "", errno.UserAlreadyExistErr
		}
		return "", err
	}

	var token string
	token, err = GenerateToken(uid)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UserService) Login(ctx context.Context, req biz.LoginReq) (string, error) {
	u, err := s.repo.FindByName(ctx, req.Name)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password))
	if err != nil {
		return "", errno.PasswordIsNotVerified
	}

	var token string
	token, err = GenerateToken(u.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UserService) GetUserInfo(ctx context.Context, uid int64) (*common.User, error) {
	res := &common.User{}
	errChan := make(chan error, 2)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		u, err := s.repo.FindByID(ctx, uid)
		if err != nil {
			errChan <- err
		} else {
			res.Id = u.ID
			res.Name = u.Name
			res.Avatar = u.Avatar
			res.BackgroundImage = u.BackgroundImage
			res.Signature = u.Signature
		}
		wg.Done()
	}()

	go func() {
		favCountResp, err := client.FavoriteClient.FavoriteCount(ctx, &favorite.FavoriteCountRequest{
			UserId: uid,
		})
		if err != nil {
			errChan <- err
		} else {
			res.FavoriteCount = favCountResp.Count
		}
		wg.Done()
	}()

	wg.Wait()
	select {
	case result := <-errChan:
		return &common.User{}, result
	default:
	}

	return res, nil
}
