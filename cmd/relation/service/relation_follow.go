package service

import (
	"context"
	"errors"

	"github.com/linzijie1998/mini-tiktok/cmd/relation/dal/db"
	"github.com/linzijie1998/mini-tiktok/cmd/relation/global"
	"github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/relation"
	"github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/user"
	"github.com/linzijie1998/mini-tiktok/pkg/errno"
	"github.com/linzijie1998/mini-tiktok/pkg/jwt"
	"gorm.io/gorm"
)

type RelationFollowListService struct {
	ctx context.Context
}

func NewRelationFollowListService(ctx context.Context) *RelationFollowListService {
	return &RelationFollowListService{ctx: ctx}
}

func (s *RelationFollowListService) FollowList(req *relation.FollowListRequest) ([]*user.User, error) {
	claims, err := jwt.NewJWT(global.Configs.JWT.SigningKey).ParseToken(req.Token)
	if err != nil {
		return nil, err
	}
	if claims.Id == 0 || claims.Issuer != global.Configs.JWT.Issuer || claims.Subject != global.Configs.JWT.Subject {
		return nil, errno.AuthorizationFailedErr
	}
	relationInfos, err := db.QueryFollowInfos(s.ctx, req.UserId, "follow_user_id")
	if err != nil {
		return nil, err
	}
	userList := make([]*user.User, len(relationInfos))
	for i, relationInfo := range relationInfos {
		userInfo, err := db.QueryFirstUserInfoByID(s.ctx, relationInfo.FollowUserId, "nickname")
		if err != nil {
			return nil, err
		}

		isFollow := false
		if claims.Id == req.UserId {
			isFollow = true
		} else {
			err = db.QueryFollowInfo(s.ctx, claims.Id, relationInfo.FollowUserId, "id")
			if err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, err
				}
				isFollow = false
			} else {
				isFollow = true
			}
		}

		userList[i] = &user.User{
			Id:       relationInfo.FollowUserId,
			Name:     userInfo.Nickname,
			IsFollow: isFollow,
		}
	}
	return userList, nil
}
