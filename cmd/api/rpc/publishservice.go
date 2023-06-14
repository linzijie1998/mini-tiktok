package rpc

import (
	"context"
	"github.com/linzijie1998/mini-tiktok/cmd/api/global"
	"github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/publish"
	"github.com/linzijie1998/mini-tiktok/pkg/errno"
)

func PublishAction(ctx context.Context, req *publish.ActionRequest) (*publish.ActionResponse, error) {
	if global.PublishServiceClient == nil {
		return nil, errno.ServiceErr.WithMessage("用户微服务客户端未初始化或初始化失败")
	}
	return (*global.PublishServiceClient).PublishAction(ctx, req)
}

func PublishList(ctx context.Context, req *publish.ListRequest) (*publish.ListResponse, error) {
	if global.PublishServiceClient == nil {
		return nil, errno.ServiceErr.WithMessage("用户微服务客户端未初始化或初始化失败")
	}
	return (*global.PublishServiceClient).PublishList(ctx, req)
}
