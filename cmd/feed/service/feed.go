package service

import (
	"context"
	"errors"
	"time"

	"github.com/linzijie1998/mini-tiktok/cmd/feed/dal/db"
	"github.com/linzijie1998/mini-tiktok/cmd/feed/global"
	"github.com/linzijie1998/mini-tiktok/cmd/feed/pack"
	"github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/feed"
	"github.com/linzijie1998/mini-tiktok/pkg/constant"
	"github.com/linzijie1998/mini-tiktok/pkg/errno"
	"github.com/linzijie1998/mini-tiktok/pkg/jwt"
	"gorm.io/gorm"
)

type FeedService struct {
	ctx context.Context
}

func NewFeedService(ctx context.Context) *FeedService {
	return &FeedService{ctx: ctx}
}

func (s *FeedService) Feed(req *feed.FeedRequest) ([]*feed.Video, int64, error) {
	// 1. 判断是否属于登录状态
	var userId int64
	if req.Token != nil && len(*req.Token) != 0 {
		claims, err := jwt.NewJWT(global.Configs.JWT.SigningKey).ParseToken(*req.Token)
		if err != nil {
			return nil, 0, err
		}
		if claims.Id == 0 || claims.Issuer != global.Configs.JWT.Issuer || claims.Subject != global.Configs.JWT.Subject {
			return nil, 0, errno.AuthorizationFailedErr
		}
		userId = claims.Id
	}
	// 2. 处理latestTime
	if req.LatestTime == nil || *req.LatestTime == 0 {
		latestTime := time.Now().UnixNano() / 1e6
		req.LatestTime = &latestTime
	}
	// 3. 根据latestTime查询视频信息
	videoInfos, err := db.QueryVideoInfoWithLimit(
		s.ctx,
		constant.MaxQueryVideoNum,
		time.Unix(*req.LatestTime/1e3, *req.LatestTime/1e3),
		"id, author_id, title, video_path, cover_path, favorite_count, comment_count",
	)
	if err != nil {
		return nil, 0, err
	}
	// 4. 查询视频作者用户信息和关注状态以及视频是否点赞
	videoList := make([]*feed.Video, len(videoInfos))
	for i, videoInfo := range videoInfos {
		userInfo, err := db.QueryFirstUserInfoByID(
			s.ctx,
			videoInfo.AuthorId,
			"id, nickname, avatar, background_image, signature, follow_count, follower_count, total_favorited, favorite_count, work_count",
		)
		if err != nil {
			return nil, 0, err
		}
		var isFavorite, isFollow bool
		if userId != 0 {
			err = db.QueryFavoriteInfo(s.ctx, userId, videoInfo.Id)
			if err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, 0, err
				}
				isFavorite = false
			} else {
				isFavorite = true
			}

			err = db.QueryFollowInfo(s.ctx, userId, videoInfo.AuthorId, "id")
			if err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, 0, err
				}
				isFollow = false
			} else {
				isFollow = true
			}
		}

		videoList[i] = pack.BuildRespVideo(videoInfo, userInfo, isFollow, isFavorite)
	}
	nextTime := time.Now().UnixNano() / 1e6
	return videoList, nextTime, nil
}
