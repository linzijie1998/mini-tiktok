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

type RelationFollowerListService struct {
	ctx context.Context
}

func NewRelationFollowerListService(ctx context.Context) *RelationFollowerListService {
	return &RelationFollowerListService{ctx: ctx}
}

func (s *RelationFollowerListService) FollowerList(req *relation.FollowerListRequest) ([]*user.User, error) {
	claims, err := jwt.NewJWT(global.Configs.JWT.SigningKey).ParseToken(req.Token)
	if err != nil {
		return nil, err
	}
	if claims.Id == 0 || claims.Issuer != global.Configs.JWT.Issuer || claims.Subject != global.Configs.JWT.Subject {
		return nil, errno.AuthorizationFailedErr
	}
	relationInfos, err := db.QueryFollowerInfos(s.ctx, req.UserId, "user_id")
	if err != nil {
		return nil, err
	}
	userList := make([]*user.User, len(relationInfos))
	for i, relationInfo := range relationInfos {
		userInfo, err := db.QueryFirstUserInfoByID(s.ctx, relationInfo.UserId, "nickname")
		if err != nil {
			return nil, err
		}

		isFollow := false
		err = db.QueryFollowInfo(s.ctx, claims.Id, relationInfo.UserId, "id")
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, err
			}
			isFollow = false
		} else {
			isFollow = true
		}

		userList[i] = &user.User{
			Id:       relationInfo.UserId,
			Name:     userInfo.Nickname,
			IsFollow: isFollow,
		}
	}
	return userList, nil
}
