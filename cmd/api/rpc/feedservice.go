package rpc

import (
	"context"

	"github.com/linzijie1998/mini-tiktok/cmd/api/global"
	"github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/feed"
	"github.com/linzijie1998/mini-tiktok/pkg/errno"
)

func Feed(ctx context.Context, req *feed.FeedRequest) (*feed.FeedResponse, error) {
	if global.FeedServiceClient == nil {
		return nil, errno.ServiceErr.WithMessage("用户微服务客户端未初始化或初始化失败")
	}
	return (*global.FeedServiceClient).Feed(ctx, req)
}
