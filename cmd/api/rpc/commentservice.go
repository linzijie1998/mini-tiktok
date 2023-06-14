package rpc

import (
	"context"

	"github.com/linzijie1998/mini-tiktok/cmd/api/global"
	"github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/comment"
	"github.com/linzijie1998/mini-tiktok/pkg/errno"
)

func CommentAction(ctx context.Context, req *comment.ActionRequest) (*comment.ActionResponse, error) {
	if global.FavoriteServiceClient == nil {
		return nil, errno.ServiceErr.WithMessage("用户微服务客户端未初始化或初始化失败")
	}
	return (*global.CommentServiceClient).CommentAction(ctx, req)
}

func CommentList(ctx context.Context, req *comment.ListRequest) (*comment.ListResponse, error) {
	if global.FavoriteServiceClient == nil {
		return nil, errno.ServiceErr.WithMessage("用户微服务客户端未初始化或初始化失败")
	}
	return (*global.CommentServiceClient).CommentList(ctx, req)
}
