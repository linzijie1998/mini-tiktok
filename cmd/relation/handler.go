package main

import (
	"context"

	"github.com/linzijie1998/mini-tiktok/cmd/relation/pack"
	"github.com/linzijie1998/mini-tiktok/cmd/relation/service"
	relation "github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/relation"
	"github.com/linzijie1998/mini-tiktok/pkg/errno"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// RelationAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *relation.ActionRequest) (*relation.ActionResponse, error) {
	if len(req.Token) == 0 || req.ToUserId == 0 {
		return pack.BuildActionResp(errno.ParamErr), nil
	}
	if err := service.NewRelationActionService(ctx).RelationAction(req); err != nil {
		return pack.BuildActionResp(err), nil
	}
	return pack.BuildActionResp(nil), nil
}

// RelationFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowList(ctx context.Context, req *relation.FollowListRequest) (*relation.FollowListResponse, error) {
	if len(req.Token) == 0 || req.UserId == 0 {
		return pack.BuildFollowListResp(nil, errno.ParamErr), nil
	}
	userList, err := service.NewRelationFollowListService(ctx).FollowList(req)
	if err != nil {
		return pack.BuildFollowListResp(nil, err), nil
	}
	return pack.BuildFollowListResp(userList, nil), nil
}

// RelationFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowerList(ctx context.Context, req *relation.FollowerListRequest) (*relation.FollowerListResponse, error) {
	if len(req.Token) == 0 || req.UserId == 0 {
		return pack.BuildFollowerListResp(nil, errno.ParamErr), nil
	}
	userList, err := service.NewRelationFollowerListService(ctx).FollowerList(req)
	if err != nil {
		return pack.BuildFollowerListResp(nil, err), nil
	}
	return pack.BuildFollowerListResp(userList, nil), nil
}

// RelationFriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFriendList(ctx context.Context, req *relation.FriendListRequest) (*relation.FriendListResponse, error) {
	if len(req.Token) == 0 || req.UserId == 0 {
		return pack.BuildFriendListResp(nil, errno.ParamErr), nil
	}
	userList, err := service.NewRelationFriendListService(ctx).FriendList(req)
	if err != nil {
		return pack.BuildFriendListResp(nil, err), nil
	}
	return pack.BuildFriendListResp(userList, nil), nil
}
