package pack

import (
	"errors"
	"github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/user"
	"github.com/linzijie1998/mini-tiktok/model"
	"github.com/linzijie1998/mini-tiktok/pkg/errno"
)

func BuildRegisterResp(id int64, token string, err error) *user.RegisterResponse {
	if err == nil {
		return registerResp(id, token, errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return registerResp(0, "", e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return registerResp(0, "", s)
}

func registerResp(id int64, token string, err errno.ErrNo) *user.RegisterResponse {
	return &user.RegisterResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg, UserId: id, Token: token}
}

func BuildLoginResp(id int64, token string, err error) *user.LoginResponse {
	if err == nil {
		return loginResp(id, token, errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return loginResp(0, "", e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return loginResp(0, "", s)
}

func loginResp(id int64, token string, err errno.ErrNo) *user.LoginResponse {
	return &user.LoginResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg, UserId: id, Token: token}
}

func BuildInfoResp(userInfo *user.User, err error) *user.InfoResponse {
	if err == nil {
		return infoResp(userInfo, errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return infoResp(&user.User{}, e) // User为required属性, 需要传零值不能传nil
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return infoResp(&user.User{}, s) // User为required属性, 需要传零值不能传nil
}

func infoResp(userInfo *user.User, err errno.ErrNo) *user.InfoResponse {
	return &user.InfoResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg, User: userInfo}
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
