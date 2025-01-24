package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"path"
	"strconv"
	"sync"
	"time"

	"github.com/crazyfrankie/douyin/app/publish/biz/repository"
	"github.com/crazyfrankie/douyin/app/publish/biz/repository/dao"
	"github.com/crazyfrankie/douyin/app/publish/common/constants"
	"github.com/crazyfrankie/douyin/app/publish/mw"
	"github.com/crazyfrankie/douyin/rpc_gen/common"
	"github.com/crazyfrankie/douyin/rpc_gen/favorite"
	"github.com/crazyfrankie/douyin/rpc_gen/publish"
)

type PublishService struct {
	repo       *repository.PublishRepo
	favoClient favorite.FavoriteServiceClient
}

func NewPublishService(repo *repository.PublishRepo, favoClient favorite.FavoriteServiceClient) *PublishService {
	return &PublishService{
		repo:       repo,
		favoClient: favoClient,
	}
}

func (s *PublishService) PublishAction(ctx context.Context, req *publish.PublishActionRequest) error {
	var fileHeader multipart.FileHeader
	err := json.Unmarshal(req.Data, &fileHeader)
	if err != nil {
		return err
	}

	now := time.Now().Unix()
	fileName := fmt.Sprintf("%d.%d", req.GetUserId(), now)
	fileHeader.Filename = fileName + path.Ext(fileHeader.Filename)

	uploadInfo, err := mw.PutToBucket(ctx, constants.MinioVideoBucketName, &fileHeader)
	if err != nil {
		return err
	}

	playUrl := constants.MinioImgBucketName + "/" + fileName

	buf, err := mw.GetSnapshot(mw.URLconvert(ctx, playUrl))
	log.Printf("video upload size:" + strconv.FormatInt(uploadInfo.Size, 10))

	uploadInfo, err = mw.PutToBucketByBuf(ctx, constants.MinioImgBucketName, fileName+".png", buf)
	log.Printf("image upload size:" + strconv.FormatInt(uploadInfo.Size, 10))
	if err != nil {
		log.Printf("minio上传封面失败: %v", err)
		return err
	}

	video := &dao.Video{
		AuthorID: req.GetUserId(),
		Title:    req.GetTitle(),
		PlayURL:  playUrl,
		CoverURL: constants.MinioImgBucketName + "/" + fileName + ".png",
		Ctime:    now,
		Utime:    now,
	}

	return s.repo.AddVideo(ctx, video)
}

func (s *PublishService) PublishList(ctx context.Context, uid int64) ([]*common.Video, error) {
	videoIds, err := s.repo.GetPublishVideos(ctx, uid)
	if err != nil {
		return nil, err
	}

	var videos []*common.Video
	for _, id := range videoIds {
		video, err := s.videoInfo(ctx, id, uid)
		if err != nil {
			return nil, err
		}

		videos = append(videos, video)
	}

	return videos, nil
}

func (s *PublishService) videoInfo(ctx context.Context, vid, uid int64) (*common.Video, error) {
	var video *common.Video
	var wg sync.WaitGroup

	errChan := make(chan error, 4)
	defer close(errChan)
	wg.Add(4)

	go func() {
		v, err := s.repo.GetPublishVideoInfo(ctx, vid)
		if err != nil {
			errChan <- err
		} else {
			video.Id = v.ID
			video.Title = v.Title
			video.PlayUrl = v.PlayURL
			video.CoverUrl = v.CoverURL
		}

		wg.Done()
	}()

	// Get video favorited count
	go func() {
		resp, err := s.favoClient.VideoFavoriteCount(ctx, &favorite.VideoFavoriteCountRequest{
			VideoId: vid,
		})
		if err != nil {
			errChan <- err
		} else {
			video.FavoriteCount = resp.Count
		}

		wg.Done()
	}()

	// Get is_favorite
	go func() {
		resp, err := s.favoClient.IsFavorite(ctx, &favorite.IsFavoriteRequest{
			VideoId: vid,
			UserId:  uid,
		})
		if err != nil {
			errChan <- err
		} else {
			video.IsFavorite = resp.IsFavorite
		}

		wg.Done()
	}()

	// Get comment count
	go func() {

	}()

	wg.Wait()

	select {
	case err := <-errChan:
		return &common.Video{}, err
	default:
	}

	return video, nil
}

func (s *PublishService) PublishCount(ctx context.Context, uid int64) ([]int64, error) {
	return s.repo.GetUserPublishCount(ctx, uid)
}
