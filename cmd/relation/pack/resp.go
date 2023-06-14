package pack

import (
	"errors"

	"github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/relation"
	"github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/user"
	"github.com/linzijie1998/mini-tiktok/pkg/errno"
)

func BuildActionResp(err error) *relation.ActionResponse {
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

func actionResp(err errno.ErrNo) *relation.ActionResponse {
	return &relation.ActionResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}

func BuildFollowListResp(userList []*user.User, err error) *relation.FollowListResponse {
	if err == nil {
		return listFollowResp(userList, errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return listFollowResp(nil, e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return listFollowResp(nil, s)
}

func listFollowResp(userList []*user.User, err errno.ErrNo) *relation.FollowListResponse {
	return &relation.FollowListResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg, UserList: userList}
}

func BuildFollowerListResp(userList []*user.User, err error) *relation.FollowerListResponse {
	if err == nil {
		return listFollowerResp(userList, errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return listFollowerResp(nil, e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return listFollowerResp(nil, s)
}

func listFollowerResp(userList []*user.User, err errno.ErrNo) *relation.FollowerListResponse {
	return &relation.FollowerListResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg, UserList: userList}
}

func BuildFriendListResp(userList []*relation.FriendUser, err error) *relation.FriendListResponse {
	if err == nil {
		return listFriendResp(userList, errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return listFriendResp(nil, e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return listFriendResp(nil, s)
}

func listFriendResp(userList []*relation.FriendUser, err errno.ErrNo) *relation.FriendListResponse {
	return &relation.FriendListResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg, UserList: userList}
}
