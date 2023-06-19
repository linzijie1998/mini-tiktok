package service

import (
	"context"
	"github.com/linzijie1998/mini-tiktok/cmd/user/dal"
	"github.com/linzijie1998/mini-tiktok/cmd/user/dal/cache"
	"github.com/linzijie1998/mini-tiktok/cmd/user/global"
	"github.com/linzijie1998/mini-tiktok/cmd/user/pack"
	"github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/user"
	"github.com/linzijie1998/mini-tiktok/pkg/errno"
	"github.com/linzijie1998/mini-tiktok/pkg/jwt"
)

type UserInfoService struct {
	ctx context.Context
}

func NewUserInfoService(ctx context.Context) *UserInfoService {
	return &UserInfoService{ctx: ctx}
}

func (s *UserInfoService) UserInfo(req *user.InfoRequest) (*user.User, error) {
	// 根据uid查询用户信息
	userInfo, err := dal.QueryUserInfoById(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	// 解析token, 判断关注状态
	isFollow := false
	if len(req.Token) != 0 {
		// 解析token
		claims, err := jwt.NewJWT(global.Configs.JWT.SigningKey).ParseToken(req.Token)
		if err != nil {
			return nil, err
		}
		// 校验信息
		if claims.Id == 0 || claims.Issuer != global.Configs.JWT.Issuer || claims.Subject != global.Configs.JWT.Subject {
			return nil, errno.AuthorizationFailedErr
		}
		// 判断关注状态
		if claims.Id != req.UserId {
			isFollow, err = cache.GetFollowState(s.ctx, claims.Id, req.UserId)
			if err != nil {
				return nil, err
			}
		}
	}
	return pack.BuildRespUser(userInfo, isFollow), nil
}
