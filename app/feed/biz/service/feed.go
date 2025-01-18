package service

import (
	"context"
	"github.com/crazyfrankie/douyin/rpc_gen/common"
	"log"
	"sync"

	"github.com/crazyfrankie/douyin/app/feed/biz"
	"github.com/crazyfrankie/douyin/app/feed/biz/repository"
	"github.com/crazyfrankie/douyin/app/feed/biz/repository/dao"
	"github.com/crazyfrankie/douyin/app/feed/rpc/client"
	"github.com/crazyfrankie/douyin/rpc_gen/user"
)

type FeedService struct {
	repo *repository.FeedRepo
}

func NewFeedService(repo *repository.FeedRepo) *FeedService {
	return &FeedService{
		repo: repo,
	}
}

// Feed returns a list of recommended videos for logged-in user
func (s *FeedService) Feed(ctx context.Context, req biz.FeedReq) (biz.FeedResp, error) {
	return biz.FeedResp{}, nil
}

func (s *FeedService) VideoList(ctx context.Context, vid []int64) ([]dao.Video, error) {
	return s.repo.VideoList(ctx, vid)
}

func (s *FeedService) VideoInfo(ctx context.Context, vid, uid int64) (*common.Video, error) {
	var wg sync.WaitGroup
	wg.Add(2)

	var video *common.Video

	// Get user information
	go func() {
		resp, err := client.UserClient.GetUserInfo(ctx, &user.GetUserInfoRequest{
			UserId: uid,
		})
		if err != nil {
			log.Printf("GetUserInfo func error:" + err.Error())
		}
		video.Author = &common.User{
			Id:              resp.User.Id,
			Name:            resp.User.Name,
			Avatar:          resp.User.Name,
			Signature:       resp.User.Signature,
			FollowCount:     resp.User.FollowCount,
			FollowerCount:   resp.User.FollowerCount,
			WorkCount:       resp.User.WorkCount,
			BackgroundImage: resp.User.BackgroundImage,
			TotalFavorited:  resp.User.TotalFavorited,
			FavoriteCount:   resp.User.FavoriteCount,
			IsFollow:        resp.User.IsFollow,
		}

		wg.Done()
	}()

	// Get comment count
	go func() {
	
	}()

	wg.Wait()

	return video, nil
}

func (s *FeedService) GetVideoExists(ctx context.Context, vid int64) (bool, error) {
	return s.repo.QueryVideoExists(ctx, vid)
}
