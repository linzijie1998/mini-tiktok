package pack

import (
	"errors"
	"fmt"

	"github.com/linzijie1998/mini-tiktok/cmd/feed/global"
	"github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/feed"
	"github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/user"
	"github.com/linzijie1998/mini-tiktok/model"
	"github.com/linzijie1998/mini-tiktok/pkg/errno"
)

func BuildFeedResp(videoList []*feed.Video, nextTime int64, err error) *feed.FeedResponse {
	if err == nil {
		return feedResp(videoList, nextTime, errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return feedResp(nil, 0, e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return feedResp(nil, 0, s)
}

func feedResp(videoList []*feed.Video, nextTime int64, err errno.ErrNo) *feed.FeedResponse {
	return &feed.FeedResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg, VideoList: videoList, NextTime: &nextTime}
}

func BuildRespVideo(videoInfo *model.Video, userInfo *model.User, isFollow bool, isFavorite bool) *feed.Video {
	if videoInfo == nil || userInfo == nil {
		return nil
	}
	return &feed.Video{
		Id:            videoInfo.Id,
		Author:        buildRespUser(userInfo, isFollow),
		Title:         videoInfo.Title,
		PlayUrl:       fmt.Sprintf("%s/%s", global.Configs.FileAccess.NginxUrl, videoInfo.VideoPath),
		CoverUrl:      fmt.Sprintf("%s/%s", global.Configs.FileAccess.NginxUrl, videoInfo.CoverPath),
		FavoriteCount: videoInfo.FavoriteCount,
		CommentCount:  videoInfo.CommentCount,
		IsFavorite:    isFavorite,
	}
}

func buildRespUser(userInfo *model.User, isFollow bool) *user.User {
	if userInfo == nil {
		return nil
	}
	if len(userInfo.Avatar) == 0 {
		userInfo.Avatar = global.Configs.StaticResource.DefaultAvatar
	}
	if len(userInfo.BackgroundImage) == 0 {
		userInfo.BackgroundImage = global.Configs.StaticResource.DefaultBackgroundImage
	}
	return &user.User{
		Id:              userInfo.Id,
		Name:            userInfo.Nickname,
		IsFollow:        isFollow,
		Avatar:          &userInfo.Avatar,
		BackgroundImage: &userInfo.BackgroundImage,
		Signature:       &userInfo.Signature,
		FollowCount:     &userInfo.FollowCount,
		FollowerCount:   &userInfo.FollowerCount,
		TotalFavorited:  &userInfo.TotalFavorited,
		FavoriteCount:   &userInfo.FavoriteCount,
		WorkCount:       &userInfo.WorkCount,
	}
}
