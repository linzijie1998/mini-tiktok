package rpc

import (
	"context"

	"github.com/linzijie1998/mini-tiktok/cmd/api/global"
	"github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/relation"
	"github.com/linzijie1998/mini-tiktok/pkg/errno"
)

func RelationAction(ctx context.Context, req *relation.ActionRequest) (*relation.ActionResponse, error) {
	if global.RelationServiceClient == nil {
		return nil, errno.ServiceErr.WithMessage("用户微服务客户端未初始化或初始化失败")
	}
	return (*global.RelationServiceClient).RelationAction(ctx, req)
}

func RelationFollowList(ctx context.Context, req *relation.FollowListRequest) (*relation.FollowListResponse, error) {
	if global.RelationServiceClient == nil {
		return nil, errno.ServiceErr.WithMessage("用户微服务客户端未初始化或初始化失败")
	}
	return (*global.RelationServiceClient).RelationFollowList(ctx, req)
}

func RelationFollowerList(ctx context.Context, req *relation.FollowerListRequest) (*relation.FollowerListResponse, error) {
	if global.RelationServiceClient == nil {
		return nil, errno.ServiceErr.WithMessage("用户微服务客户端未初始化或初始化失败")
	}
	return (*global.RelationServiceClient).RelationFollowerList(ctx, req)
}

func RelationFriendList(ctx context.Context, req *relation.FriendListRequest) (*relation.FriendListResponse, error) {
	if global.RelationServiceClient == nil {
		return nil, errno.ServiceErr.WithMessage("用户微服务客户端未初始化或初始化失败")
	}
	return (*global.RelationServiceClient).RelationFriendList(ctx, req)
}
