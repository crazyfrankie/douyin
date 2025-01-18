package service

import (
	"context"
	"github.com/crazyfrankie/douyin/rpc_gen/feed"
	"log"
	"sync"

	"github.com/crazyfrankie/douyin/app/favorite/biz"
	"github.com/crazyfrankie/douyin/app/favorite/biz/repository"
	"github.com/crazyfrankie/douyin/app/favorite/biz/repository/dao"
	"github.com/crazyfrankie/douyin/app/favorite/common/constants"
	"github.com/crazyfrankie/douyin/app/favorite/common/errno"
	"github.com/crazyfrankie/douyin/app/favorite/rpc/client"
	"github.com/crazyfrankie/douyin/rpc_gen/common"
)

type FavoriteService struct {
	repo *repository.FavoriteRepo
}

func NewFavoriteService(repo *repository.FavoriteRepo) *FavoriteService {
	return &FavoriteService{
		repo: repo,
	}
}

func (s *FavoriteService) FavoriteAction(ctx context.Context, req biz.FavoriteActionReq, uid int64) error {
	// 查询视频是否存在
	_, err := client.FeedClient.VideoExists(ctx, &feed.VideoExistsRequest{
		VideoId: req.VideoID,
	})
	if err != nil {
		return err
	}

	if req.ActionType != constants.FavoriteActionType && req.ActionType != constants.UnFavoriteActionType {
		return errno.ParamErr
	}

	favorite := dao.Favorite{
		VideoID: req.VideoID,
		UserID:  uid,
	}

	exists, err := s.repo.GetIsFavorite(ctx, favorite.VideoID, favorite.UserID)
	if err != nil {
		return err
	}
	if req.ActionType == constants.FavoriteActionType {
		if exists {
			return errno.FavoriteRelationAlreadyExistErr
		}
		err = s.repo.AddFavorite(ctx, favorite)
	} else {
		if !exists {
			return errno.FavoriteRelationNotExistErr
		}
		err = s.repo.DelFavorite(ctx, favorite)
	}

	return err
}

func (s *FavoriteService) FavoriteList(ctx context.Context, uid int64) ([]*common.Video, error) {
	favorsID, err := s.repo.GetFavoriteVideosByID(ctx, uid)
	if err != nil {
		return []*common.Video{}, err
	}

	videosResp, err := client.FeedClient.VideoList(ctx, &feed.FeedListRequest{
		VideoIds: favorsID,
	})
	var videos []*common.Video
	if err != nil {
		return videos, err
	}

	for _, item := range videosResp.Videos {
		// Get user info and comment count
		resp, err := client.FeedClient.VideoInfo(ctx, &feed.FeedInfoRequest{
			VideoId: item.Id,
			UserId:  uid,
		})
		if err != nil {
			return []*common.Video{}, err
		}

		var wg sync.WaitGroup
		wg.Add(2)

		// Get the number of video likes
		go func() {
			var err error
			resp.Video.FavoriteCount, err = s.repo.GetVideoFavoriteCount(ctx, item.Id)
			if err != nil {
				log.Printf("GetVideoFavoriteCount func error:" + err.Error())
			}
			wg.Done()
		}()

		// Get favorite exist
		go func() {
			var err error
			resp.Video.IsFavorite, err = s.repo.GetIsFavorite(ctx, item.Id, uid)
			if err != nil {
				log.Printf("GetIsFavorite func error:" + err.Error())
			}
			wg.Done()
		}()

		wg.Wait()
		resp.Video.Id = item.Id
		resp.Video.PlayUrl = item.PlayUrl
		resp.Video.CoverUrl = item.CoverUrl
		resp.Video.Title = item.Title

		videos = append(videos, resp.Video)
	}

	return videos, nil
}

func (s *FavoriteService) GetUserFavoriteCount(ctx context.Context, vid int64) (int64, error) {
	return s.repo.GetUserFavoriteCount(ctx, vid)
}
