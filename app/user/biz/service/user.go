package service

import (
	"context"
	"sync"

	"golang.org/x/crypto/bcrypt"

	"github.com/crazyfrankie/douyin/app/user/biz/repository"
	"github.com/crazyfrankie/douyin/app/user/biz/repository/dao"
	"github.com/crazyfrankie/douyin/app/user/common/errno"
	"github.com/crazyfrankie/douyin/rpc_gen/common"
	"github.com/crazyfrankie/douyin/rpc_gen/favorite"
	"github.com/crazyfrankie/douyin/rpc_gen/publish"
	"github.com/crazyfrankie/douyin/rpc_gen/user"
)

var (
	defaultAvatar = "github.com/crazyfrankie/douyin/static/avatar/default.png"
)

type UserService struct {
	repo          *repository.UserRepo
	favorClient   favorite.FavoriteServiceClient
	publishClient publish.PublishServiceClient
}

func NewUserService(repo *repository.UserRepo, favorClient favorite.FavoriteServiceClient, publishClient publish.PublishServiceClient) *UserService {
	return &UserService{repo: repo, favorClient: favorClient, publishClient: publishClient}
}

func (s *UserService) Register(ctx context.Context, req *user.RegisterRequest) (string, error) {
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

func (s *UserService) Login(ctx context.Context, req *user.LoginRequest) (string, error) {
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

func (s *UserService) GetUserInfo(ctx context.Context, req *user.GetUserInfoRequest) (*common.User, error) {
	uid := req.GetUserId()

	res := &common.User{}
	errChan := make(chan error, 4)

	var wg sync.WaitGroup
	wg.Add(4)

	// Get user dao info
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

	// Get user favorite count
	go func() {
		favCountResp, err := s.favorClient.FavoriteCount(ctx, &favorite.FavoriteCountRequest{
			UserId: uid,
		})
		if err != nil {
			errChan <- err
		} else {
			res.FavoriteCount = favCountResp.Count
		}
		wg.Done()
	}()

	var videoIds []int64
	// Get user published videos count (work count)
	go func() {
		videosCountResp, err := s.publishClient.PublishCount(ctx, &publish.PublishCountRequest{
			UserId: uid,
		})
		if err != nil {
			errChan <- err
		} else {
			res.WorkCount = int64(len(videosCountResp.VideoId))
		}

		copy(videoIds, videosCountResp.VideoId)

		wg.Done()
	}()

	// Get user total favorited count
	go func() {
		favorited, err := s.favorClient.UserFavorited(ctx, &favorite.UserFavoritedRequest{
			VideoId: videoIds,
		})
		if err != nil {
			errChan <- err
		} else {
			res.TotalFavorited = favorited.Count
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
