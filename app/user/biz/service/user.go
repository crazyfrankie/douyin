package service

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/crazyfrankie/douyin/app/user/biz/repository"
	"github.com/crazyfrankie/douyin/app/user/biz/repository/dao"
	"github.com/crazyfrankie/douyin/app/user/common/constants"
	"github.com/crazyfrankie/douyin/app/user/common/errno"
	"github.com/crazyfrankie/douyin/app/user/config"
	"github.com/crazyfrankie/douyin/rpc_gen/common"
	"github.com/crazyfrankie/douyin/rpc_gen/favorite"
	"github.com/crazyfrankie/douyin/rpc_gen/publish"
	"github.com/crazyfrankie/douyin/rpc_gen/relation"
	"github.com/crazyfrankie/douyin/rpc_gen/sms"
	"github.com/crazyfrankie/douyin/rpc_gen/user"
)

var (
	defaultAvatar = "github.com/crazyfrankie/douyin/static/avatar/default.png"
)

type UserService struct {
	repo           *repository.UserRepo
	favorClient    favorite.FavoriteServiceClient
	publishClient  publish.PublishServiceClient
	relationClient relation.RelationServiceClient
	smsClient      sms.SmsServiceClient
}

func NewUserService(repo *repository.UserRepo, favorClient favorite.FavoriteServiceClient, publishClient publish.PublishServiceClient, relationClient relation.RelationServiceClient, smsClient sms.SmsServiceClient) *UserService {
	return &UserService{repo: repo, favorClient: favorClient, publishClient: publishClient, relationClient: relationClient, smsClient: smsClient}
}

func (s *UserService) Register(ctx context.Context, req *user.RegisterRequest) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	u := dao.User{
		Name:     req.GetName(),
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

func (s *UserService) SendCode(ctx context.Context, phone string) (string, error) {
	u, err := s.repo.FindByPhone(ctx, phone)
	if err != nil {
		return "", err
	}
	var biz string
	if u.ID == 0 {
		biz = constants.Register
	} else {
		biz = constants.Login
	}

	code := generateCode()
	hashCode := generateHMAC(code, config.GetConf().JWT.SecretKey)

	_, err = s.smsClient.SendSms(ctx, &sms.SendSmsRequest{
		Biz:     biz,
		Args:    []string{hashCode},
		Numbers: []string{phone},
	})
	if err != nil {
		return "", err
	}

	return biz, nil
}

func (s *UserService) VerifyCode(ctx context.Context, biz, phone, code string) (string, error) {
	hashCode := generateHMAC(code, config.GetConf().JWT.SecretKey)
	_, err := s.smsClient.VerifySms(ctx, &sms.VerifySmsRequest{
		Biz:    biz,
		Number: phone,
		Code:   hashCode,
	})
	if err != nil {
		return "", err
	}

	var uid int64
	if biz == constants.Register {
		u := dao.User{
			Phone:    phone,
			Name:     phone,
			Avatar:   defaultAvatar,
			Password: phone,
		}
		uid, err = s.repo.CreateUser(ctx, u)
		if err != nil {
			return "", err
		}
	}

	var token string
	token, err = GenerateToken(uid)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UserService) GetUserInfo(ctx context.Context, req *user.GetUserInfoRequest) (*common.User, error) {
	toUserId, currentId := req.GetUserIdToQuery(), req.GetUserId()

	res := &common.User{}
	errChan := make(chan error, 4)

	var wg sync.WaitGroup
	wg.Add(6)

	// Get user dao info
	go func() {
		u, err := s.repo.FindByID(ctx, toUserId)
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
			UserId: toUserId,
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
			UserId: toUserId,
		})
		if err != nil {
			errChan <- err
		} else {
			res.WorkCount = int64(len(videosCountResp.VideoId))
		}

		copy(videoIds, videosCountResp.VideoId)

		wg.Done()
	}()

	// Get follow and follower count
	go func() {
		resp, err := s.relationClient.RelationFollowCount(ctx, &relation.RelationFollowCountRequest{
			UserId: toUserId,
		})
		if err != nil {
			errChan <- err
		} else {
			res.FollowCount = resp.FollowCount
			res.FollowerCount = resp.FollowerCount
		}

		wg.Done()
	}()

	// Get is follow
	go func() {
		if currentId != 0 {
			resp, err := s.relationClient.RelationIsFollow(ctx, &relation.RelationIsFollowRequest{
				UserId:   currentId,
				ToUserId: toUserId,
			})
			if err != nil {
				errChan <- err
			} else {
				res.IsFollow = resp.IsFollow
			}
		} else {
			res.IsFollow = false
		}

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

func (s *UserService) GetUserExists(ctx context.Context, req *user.GetUserExistsRequest) (bool, error) {
	u, err := s.repo.FindByID(ctx, req.GetUserId())
	if err != nil {
		return false, err
	}

	return u.ID != 0, nil
}

func generateCode() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	var code strings.Builder
	for i := 0; i < 6; i++ {
		digit := rand.Intn(10)
		code.WriteString(strconv.Itoa(digit))
	}
	return code.String()
}

func generateHMAC(code, key string) string {
	// 创建一个新的 HMAC 哈希对象，使用 SHA-256 哈希算法，并以 key 作为密钥。
	h := hmac.New(sha256.New, []byte(key))

	// 将输入的 code 数据写入到 HMAC 哈希对象中，进行哈希计算。
	h.Write([]byte(code))

	// 计算哈希值并返回其十六进制表示形式。
	return hex.EncodeToString(h.Sum(nil))
}
