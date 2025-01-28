package service

import (
	"context"
	"strconv"

	"google.golang.org/grpc/metadata"
	
	"github.com/crazyfrankie/douyin/app/relation/biz/repository"
	"github.com/crazyfrankie/douyin/app/relation/biz/repository/dao"
	"github.com/crazyfrankie/douyin/app/relation/common/errno"
	"github.com/crazyfrankie/douyin/rpc_gen/common"
	"github.com/crazyfrankie/douyin/rpc_gen/relation"
	"github.com/crazyfrankie/douyin/rpc_gen/user"
)

type RelationService struct {
	repo       *repository.RelationRepo
	userClient user.UserServiceClient
}

func NewRelationService(repo *repository.RelationRepo, userClient user.UserServiceClient) *RelationService {
	return &RelationService{repo: repo, userClient: userClient}
}

func (s *RelationService) FollowAction(ctx context.Context, req *relation.RelationActionRequest) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errno.ParamErr
	}
	userId := md["user_id"][0]
	uId, _ := strconv.Atoi(userId)
	uid := int64(uId)
	_, err := s.userClient.GetUserExists(ctx, &user.GetUserExistsRequest{
		UserId: uid,
	})
	if err != nil {
		return err
	}

	follow := dao.Relation{
		UserId:   uid,
		ToUserId: req.GetToUserId(),
	}
	exists, _ := s.repo.GetFollowExists(ctx, uid, req.GetToUserId())
	if req.GetActionType() == 1 {
		if exists {
			return errno.FollowRelationAlreadyExistErr
		}
		err = s.repo.AddFollow(ctx, follow)
	} else {
		if !exists {
			return errno.FollowRelationNotExistErr
		}
		err = s.repo.DelFollow(ctx, follow)
	}

	return err
}

func (s *RelationService) FollowList(ctx context.Context, req *relation.RelationFollowListRequest) ([]*common.User, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errno.ParamErr
	}
	userId := md["user_id"][0]
	uId, _ := strconv.Atoi(userId)
	uid := int64(uId)
	_, err := s.userClient.GetUserExists(ctx, &user.GetUserExistsRequest{
		UserId: uid,
	})
	if err != nil {
		return nil, err
	}

	var follows []*common.User
	dbFollows, err := s.repo.GetFollowList(ctx, uid)
	if err != nil {
		return follows, err
	}

	for _, f := range dbFollows {
		resp, err := s.userClient.GetUserInfo(ctx, &user.GetUserInfoRequest{
			UserIdToQuery: f.ToUserId,
			UserId:        uid,
		})
		if err != nil {
			continue
		}
		follows = append(follows, resp.GetUser())
	}

	return follows, nil
}

func (s *RelationService) FollowerList(ctx context.Context, req *relation.RelationFollowerListRequest) ([]*common.User, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errno.ParamErr
	}
	userId := md["user_id"][0]
	uId, _ := strconv.Atoi(userId)
	uid := int64(uId)
	_, err := s.userClient.GetUserExists(ctx, &user.GetUserExistsRequest{
		UserId: uid,
	})
	if err != nil {
		return nil, err
	}

	var followers []*common.User
	dbFollowers, err := s.repo.GetFollowerList(ctx, uid)
	if err != nil {
		return followers, err
	}

	for _, f := range dbFollowers {
		resp, err := s.userClient.GetUserInfo(ctx, &user.GetUserInfoRequest{
			UserIdToQuery: f.UserId,
			UserId:        uid,
		})
		if err != nil {
			continue
		}
		followers = append(followers, resp.GetUser())
	}

	return followers, nil
}

func (s *RelationService) FriendList(ctx context.Context, req *relation.RelationFriendListRequest) ([]*relation.FriendUser, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errno.ParamErr
	}
	userId := md["user_id"][0]
	uId, _ := strconv.Atoi(userId)
	uid := int64(uId)
	_, err := s.userClient.GetUserExists(ctx, &user.GetUserExistsRequest{
		UserId: uid,
	})
	if err != nil {
		return nil, err
	}

	var friends []*relation.FriendUser
	dbFriends, err := s.repo.GetFriendList(ctx, uid)

	for _, f := range dbFriends {
		resp, err := s.userClient.GetUserInfo(ctx, &user.GetUserInfoRequest{
			UserIdToQuery: f,
			UserId:        uid,
		})
		if err != nil {
			continue
		}
		friends = append(friends, &relation.FriendUser{
			User: resp.GetUser(),
		})
	}

	return friends, nil
}

func (s *RelationService) GetFollowCount(ctx context.Context, req *relation.RelationFollowCountRequest) (int64, int64, error) {
	return s.repo.GetFollowCount(ctx, req.GetUserId())
}

func (s *RelationService) IsFollow(ctx context.Context, req *relation.RelationIsFollowRequest) (bool, error) {
	return s.repo.GetIsFollow(ctx, req.GetUserId(), req.GetToUserId())
}
