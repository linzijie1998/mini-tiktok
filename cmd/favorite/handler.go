package main

import (
	"context"

	"github.com/linzijie1998/mini-tiktok/cmd/favorite/pack"
	"github.com/linzijie1998/mini-tiktok/cmd/favorite/service"
	favorite "github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/favorite"
	"github.com/linzijie1998/mini-tiktok/pkg/errno"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *favorite.ActionRequest) (*favorite.ActionResponse, error) {
	if len(req.Token) == 0 || req.VideoId == 0 {
		return pack.BuildActionResp(errno.ParamErr), nil
	}
	err := service.NewFavoriteActionService(ctx).FavoriteAction(req)
	if err != nil {
		return pack.BuildActionResp(err), nil
	}
	return pack.BuildActionResp(nil), nil
}

// FavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, req *favorite.ListRequest) (*favorite.ListResponse, error) {
	if req.UserId == 0 {
		return pack.BuildListResp(nil, errno.ParamErr), nil
	}
	videoList, err := service.NewFavoriteListService(ctx).FavoriteList(req)
	if err != nil {
		return pack.BuildListResp(nil, err), nil
	}
	return pack.BuildListResp(videoList, nil), nil
}
