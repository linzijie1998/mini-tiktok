package service

import (
	"context"
	"fmt"
	"github.com/linzijie1998/mini-tiktok/cmd/favorite/dal"
	"github.com/linzijie1998/mini-tiktok/cmd/favorite/dal/mongodb"
	"github.com/linzijie1998/mini-tiktok/cmd/favorite/global"
	"github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/favorite"
	"github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/feed"
	"github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/user"
	"github.com/linzijie1998/mini-tiktok/pkg/errno"
	"github.com/linzijie1998/mini-tiktok/pkg/jwt"
)

type FavoriteListService struct {
	ctx context.Context
}

func NewFavoriteListService(ctx context.Context) *FavoriteListService {
	return &FavoriteListService{ctx: ctx}
}

func (s *FavoriteListService) FavoriteList(req *favorite.ListRequest) ([]*feed.Video, error) {
	// 判断登录状态
	var userId int64
	if len(req.Token) != 0 {
		claims, err := jwt.NewJWT(global.Configs.JWT.SigningKey).ParseToken(req.Token)
		if err != nil {
			return nil, err
		}
		if claims.Id == 0 || claims.Issuer != global.Configs.JWT.Issuer || claims.Subject != global.Configs.JWT.Subject {
			return nil, errno.AuthorizationFailedErr
		}
		userId = claims.Id
	}

	favoriteInfos, err := mongodb.GetFavoriteList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	videoList := make([]*feed.Video, len(favoriteInfos))
	for i, vid := range favoriteInfos {
		videoInfo, err := dal.QueryVideoInfoById(s.ctx, vid)
		if err != nil {
			return nil, err
		}
		isFavorite := false
		if userId != 0 {
			if userId == req.UserId {
				isFavorite = true
			} else {
				isFavorite, err = mongodb.GetFavoriteInfo(s.ctx, userId, vid)
				if err != nil {
					return nil, err
				}
			}
		}
		videoList[i] = &feed.Video{
			Id:            vid,
			Author:        &user.User{Id: videoInfo.AuthorId},
			CoverUrl:      fmt.Sprintf("%s/%s", global.Configs.FileAccess.NginxUrl, videoInfo.CoverPath),
			FavoriteCount: videoInfo.FavoriteCount,
			IsFavorite:    isFavorite,
		}
	}
	return videoList, nil
}
