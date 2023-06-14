package service

import (
	"context"

	"github.com/linzijie1998/mini-tiktok/cmd/favorite/dal/db"
	"github.com/linzijie1998/mini-tiktok/model"

	"github.com/linzijie1998/mini-tiktok/cmd/favorite/global"

	"github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/favorite"
	"github.com/linzijie1998/mini-tiktok/pkg/constant"
	"github.com/linzijie1998/mini-tiktok/pkg/errno"
	"github.com/linzijie1998/mini-tiktok/pkg/jwt"
)

type FavoriteActionService struct {
	ctx context.Context
}

func NewFavoriteActionService(ctx context.Context) *FavoriteActionService {
	return &FavoriteActionService{ctx: ctx}
}

func (s *FavoriteActionService) FavoriteAction(req *favorite.ActionRequest) error {
	// 1.解析token
	if len(req.Token) == 0 {
		return errno.ParamErr
	}
	claims, err := jwt.NewJWT(global.Configs.JWT.SigningKey).ParseToken(req.Token)
	if err != nil {
		return err
	}
	if claims.Id == 0 || claims.Issuer != global.Configs.JWT.Issuer || claims.Subject != global.Configs.JWT.Subject {
		return errno.AuthorizationFailedErr
	}
	// 2.根据ActionType处理请求
	if req.ActionType == constant.FavoriteActionLike {
		err = db.AddFavoriteInfo(s.ctx, []*model.Favorite{
			{
				UserId:  claims.Id,
				VideoId: req.VideoId,
			},
		})
		return err
	} else if req.ActionType == constant.FavoriteActionCancel {
		return db.DeleteFavoriteInfo(s.ctx, claims.Id, req.VideoId)
	} else {
		return errno.ParamErr
	}
}
