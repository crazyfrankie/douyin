package service

import (
	"context"
	"github.com/crazyfrankie/douyin/rpc_gen/comment"
	"log"
	"sync"
	"time"

	"github.com/crazyfrankie/douyin/app/feed/biz/repository"
	"github.com/crazyfrankie/douyin/app/feed/biz/repository/dao"
	"github.com/crazyfrankie/douyin/rpc_gen/common"
	"github.com/crazyfrankie/douyin/rpc_gen/favorite"
	"github.com/crazyfrankie/douyin/rpc_gen/feed"
	"github.com/crazyfrankie/douyin/rpc_gen/user"
)

type FeedService struct {
	repo          *repository.FeedRepo
	userClient    user.UserServiceClient
	favorClient   favorite.FavoriteServiceClient
	commentClient comment.CommentServiceClient
}

func NewFeedService(repo *repository.FeedRepo, userClient user.UserServiceClient, favorClient favorite.FavoriteServiceClient, commentClient comment.CommentServiceClient) *FeedService {
	return &FeedService{repo: repo, userClient: userClient, favorClient: favorClient, commentClient: commentClient}
}

// Feed returns a list of recommended videos for logged-in user
func (s *FeedService) Feed(ctx context.Context, req *feed.FeedRequest) (*feed.FeedResponse, error) {
	var lastTime int64
	if req.LatestTime == 0 {
		lastTime = time.Now().Unix()
	} else {
		lastTime = req.LatestTime
	}
	dbVideos, err := s.repo.GetVideoByLastTime(ctx, lastTime)
	if err != nil {
		return &feed.FeedResponse{}, err
	}

	var videos []*common.Video
	for _, v := range dbVideos {
		resp, err := s.VideoInfo(ctx, &feed.VideoInfoRequest{
			VideoId: v.ID,
			UserId:  v.AuthorID,
		})
		if err != nil {
			return &feed.FeedResponse{}, err
		}

		video := &common.Video{
			Id: v.ID,
			Author: &common.User{
				Id:              v.AuthorID,
				Name:            resp.Author.Name,
				Avatar:          resp.Author.Avatar,
				Signature:       resp.Author.Signature,
				BackgroundImage: resp.Author.BackgroundImage,
				FollowCount:     resp.Author.FollowCount,
				FollowerCount:   resp.Author.FollowerCount,
				TotalFavorited:  resp.Author.TotalFavorited,
				WorkCount:       resp.Author.WorkCount,
				FavoriteCount:   resp.Author.FavoriteCount,
			},
			PlayUrl:       v.PlayURL,
			CoverUrl:      v.CoverURL,
			Title:         v.Title,
			IsFavorite:    false,
			CommentCount:  resp.CommentCount,
			FavoriteCount: resp.FavoriteCount,
		}

		videos = append(videos, video)
	}

	return &feed.FeedResponse{
		Videos:   videos,
		NextTime: dbVideos[len(dbVideos)-1].Utime,
	}, nil
}

func (s *FeedService) VideoList(ctx context.Context, vid []int64) ([]dao.Video, error) {
	return s.repo.VideoList(ctx, vid)
}

func (s *FeedService) VideoInfo(ctx context.Context, req *feed.VideoInfoRequest) (*common.Video, error) {
	uid, vid, idToQuery := req.GetUserId(), req.GetVideoId(), req.GetUserIdToQuery()

	var wg sync.WaitGroup
	wg.Add(4)

	var video *common.Video

	// Get user information
	go func() {
		resp, err := s.userClient.GetUserInfo(ctx, &user.GetUserInfoRequest{
			UserId:        uid,
			UserIdToQuery: idToQuery,
		})
		if err != nil {
			log.Printf("GetUserInfo func error:" + err.Error())
		}
		video.Author = &common.User{
			Id:              resp.User.Id,
			Name:            resp.User.Name,
			Avatar:          resp.User.Name,
			Signature:       resp.User.Signature,
			BackgroundImage: resp.User.BackgroundImage,
			FollowCount:     resp.User.FollowCount,
			FollowerCount:   resp.User.FollowerCount,
			TotalFavorited:  resp.User.TotalFavorited,
			WorkCount:       resp.User.WorkCount,
			FavoriteCount:   resp.User.FavoriteCount,
			IsFollow:        resp.User.IsFollow,
		}

		wg.Done()
	}()

	// Get the number of video likes
	go func() {
		resp, err := s.favorClient.VideoFavoriteCount(ctx, &favorite.VideoFavoriteCountRequest{
			VideoId: vid,
		})
		if err != nil {
			log.Printf("VideoFavoriteCount func error:" + err.Error())
		}
		video.FavoriteCount = resp.Count
	}()

	// Get is_favorite
	go func() {
		resp, err := s.favorClient.IsFavorite(ctx, &favorite.IsFavoriteRequest{
			VideoId: vid,
			UserId:  uid,
		})
		if err != nil {
			log.Printf("IsFavorite func error:" + err.Error())
		}
		video.IsFavorite = resp.IsFavorite
	}()

	// Get comment count
	go func() {
		resp, err := s.commentClient.CommentCount(ctx, &comment.CommentCountRequest{
			VideoId: vid,
		})
		if err != nil {
			log.Printf("CommentCount func error:" + err.Error())
		}
		video.CommentCount = resp.Count
	}()

	wg.Wait()

	return video, nil
}

func (s *FeedService) GetVideoExists(ctx context.Context, vid int64) (bool, error) {
	return s.repo.QueryVideoExists(ctx, vid)
}
