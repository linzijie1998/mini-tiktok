package service

import (
	"context"
	"errors"
	"github.com/linzijie1998/mini-tiktok/cmd/user/constant"
	"github.com/linzijie1998/mini-tiktok/cmd/user/dal/cache"
	"github.com/linzijie1998/mini-tiktok/cmd/user/dal/db"
	"github.com/linzijie1998/mini-tiktok/cmd/user/global"
	"github.com/linzijie1998/mini-tiktok/cmd/user/pack"
	"github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/user"
	"github.com/linzijie1998/mini-tiktok/model"
	"github.com/linzijie1998/mini-tiktok/pkg/errno"
	"github.com/linzijie1998/mini-tiktok/pkg/jwt"
	"gorm.io/gorm"
)

type UserInfoService struct {
	ctx context.Context
}

func NewUserInfoService(ctx context.Context) *UserInfoService {
	return &UserInfoService{ctx: ctx}
}

// UserInfo 查询用户信息: 包含基本信息 (目前的场景为只读信息, 而实际情况应该为读多写少) 和计数信息 (经常会变更的数据);
// 数据查询顺序为先查Redis, 当Redis未命中时查MySQL, MySQL也未命中时使用空值缓存应对Redis缓存穿透问题;
// 当Redis命中时就更新一次过期时间, 防止出现Redis缓存雪崩的问题
func (s *UserInfoService) UserInfo(req *user.InfoRequest) (*user.User, error) {
	// 1. 查询该UID是否在存在空值缓存
	if err := cache.GetUserInfoNullKey(s.ctx, req.UserId); err == nil {
		// 空值缓存命中, 该UID的用户不存在
		return nil, errno.UserNotRegisterErr
	}
	// 2. 用户基础信息查询（先查Redis，缓存未命中再查询MySQL）
	userInfo, err := cache.GetUserInfo(s.ctx, req.UserId)
	if err != nil {
		// 缓存未命中, 需要从数据库读取用户信息
		userInfo, err = db.QueryUserInfosByID(s.ctx, req.UserId, constant.UserBaseInfoQueryString)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 未找到该UID的用户, 设置空值缓存防止Redis缓存穿透
				_ = cache.NewUserInfoNullKey(s.ctx, req.UserId, global.Configs.CacheExpire.ParseNullKeyExpireDuration())
				return nil, errno.UserNotRegisterErr
			}
			return nil, err
		}
		// 添加缓存信息（忽略错误信息）
		_ = cache.NewUserInfos(s.ctx, []*model.User{userInfo}, global.Configs.CacheExpire.ParseUserBaseInfoExpireDuration())
	}
	if len(userInfo.Avatar) == 0 {
		userInfo.Avatar = global.Configs.StaticResource.DefaultAvatar
	}
	if len(userInfo.BackgroundImage) == 0 {
		userInfo.BackgroundImage = global.Configs.StaticResource.DefaultBackgroundImage
	}
	// 3. 计数信息查询（先查Redis，缓存未命中再查询MySQL）
	userCounter, err := cache.GetUserCounter(s.ctx, req.UserId)
	if err != nil {
		// 缓存未命中, 需要从数据库读取用户信息
		userCounter, err = db.QueryUserInfosByID(s.ctx, req.UserId, constant.UserCounterInfoQueryString)
		if err != nil {
			return nil, err
		}
		// 添加缓存信息（忽略错误信息）
		_ = cache.NewUserCounters(s.ctx, []*model.User{userCounter})
	}
	if err = pack.MergeUserInfo(userInfo, userCounter); err != nil {
		return nil, err
	}
	// 4. 解析token, 判断关注状态
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
