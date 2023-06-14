package pack

import (
	"errors"
	"github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/feed"
	"github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/publish"
	"github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/user"
	"github.com/linzijie1998/mini-tiktok/model"
	"github.com/linzijie1998/mini-tiktok/pkg/errno"
)

func BuildActionResp(err error) *publish.ActionResponse {
	if err == nil {
		return actionResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return actionResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return actionResp(s)
}

func actionResp(err errno.ErrNo) *publish.ActionResponse {
	return &publish.ActionResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}

func BuildListResp(videoList []*feed.Video, err error) *publish.ListResponse {
	if err == nil {
		return listResp(videoList, errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return listResp(nil, e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return listResp(nil, s)
}

func listResp(videoList []*feed.Video, err errno.ErrNo) *publish.ListResponse {
	return &publish.ListResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg, VideoList: videoList}
}

func BuildRespUser(userInfo *model.User, isFollow bool) *user.User {
	if userInfo == nil {
		return nil
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
