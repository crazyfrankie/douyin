package service

import (
	"context"
	"fmt"
	"strconv"

	"google.golang.org/grpc/metadata"
	
	"github.com/crazyfrankie/douyin/app/favorite/biz/repository"
	"github.com/crazyfrankie/douyin/app/favorite/biz/repository/dao"
	"github.com/crazyfrankie/douyin/app/favorite/common/constants"
	"github.com/crazyfrankie/douyin/app/favorite/common/errno"
	"github.com/crazyfrankie/douyin/rpc_gen/common"
	"github.com/crazyfrankie/douyin/rpc_gen/favorite"
	"github.com/crazyfrankie/douyin/rpc_gen/feed"
)

type FavoriteService struct {
	repo       *repository.FavoriteRepo
	feedClient feed.FeedServiceClient
}

func NewFavoriteService(repo *repository.FavoriteRepo, feedClient feed.FeedServiceClient) *FavoriteService {
	return &FavoriteService{repo: repo, feedClient: feedClient}
}

// FavoriteAction adds or deletes a like relationship between the current user and the current video.
func (s *FavoriteService) FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errno.ParamErr
	}
	fmt.Println("md:", md)
	userId := md["user_id"][0]
	uId, _ := strconv.Atoi(userId)
	uid := int64(uId)
	vid := req.GetVideoId()

	// 查询视频是否存在
	_, err := s.feedClient.VideoExists(ctx, &feed.VideoExistsRequest{
		VideoId: req.GetVideoId(),
	})
	if err != nil {
		return err
	}

	if req.ActionType != constants.FavoriteActionType && req.ActionType != constants.UnFavoriteActionType {
		return errno.ParamErr
	}

	fav := dao.Favorite{
		VideoID: vid,
		UserID:  uid,
	}

	exists, err := s.repo.GetIsFavorite(ctx, fav.VideoID, fav.UserID)
	if err != nil {
		return err
	}
	if req.ActionType == constants.FavoriteActionType {
		if exists {
			return errno.FavoriteRelationAlreadyExistErr
		}
		err = s.repo.AddFavorite(ctx, fav)
	} else {
		if !exists {
			return errno.FavoriteRelationNotExistErr
		}
		err = s.repo.DelFavorite(ctx, fav)
	}

	return err
}

// FavoriteList returns the list of videos liked by the current user.
func (s *FavoriteService) FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) ([]*common.Video, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errno.ParamErr
	}
	userId := md["user_id"][0]
	uId, _ := strconv.Atoi(userId)
	uid := int64(uId)

	favorsID, err := s.repo.GetFavoriteVideosByID(ctx, uid)
	if err != nil {
		return []*common.Video{}, err
	}

	videosResp, err := s.feedClient.VideoList(ctx, &feed.VideoListRequest{
		VideoIds: favorsID,
	})
	var videos []*common.Video
	if err != nil {
		return videos, err
	}

	for _, item := range videosResp.Videos {
		// Get user info and comment count
		resp, err := s.feedClient.VideoInfo(ctx, &feed.VideoInfoRequest{
			VideoId:       item.Id,
			UserId:        uid,
			UserIdToQuery: item.Author.Id,
		})
		if err != nil {
			return []*common.Video{}, err
		}

		resp.Video.Id = item.Id
		resp.Video.PlayUrl = item.PlayUrl
		resp.Video.CoverUrl = item.CoverUrl
		resp.Video.Title = item.Title

		videos = append(videos, resp.Video)
	}

	return videos, nil
}

// GetUserFavoriteCount returns the total number of likes for the current user.
func (s *FavoriteService) GetUserFavoriteCount(ctx context.Context, vid int64) (int64, error) {
	return s.repo.GetUserFavoriteCount(ctx, vid)
}

// GetUserFavoritedCount returns the total number of likes for the current user's posts.
func (s *FavoriteService) GetUserFavoritedCount(ctx context.Context, vid []int64) (int64, error) {
	return s.repo.GetUserFavoritedCount(ctx, vid)
}

// GetVideoFavoriteCount returns the total number of likes for the current video.
func (s *FavoriteService) GetVideoFavoriteCount(ctx context.Context, vid int64) (int64, error) {
	return s.repo.GetVideoFavoriteCount(ctx, vid)
}

// GetIsFavorite returns whether the current user has liked the current video.
func (s *FavoriteService) GetIsFavorite(ctx context.Context, videoId, uid int64) (bool, error) {
	return s.repo.GetIsFavorite(ctx, videoId, uid)
}
