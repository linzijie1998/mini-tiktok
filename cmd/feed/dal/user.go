package dal

import (
	"context"
	"errors"
	"github.com/linzijie1998/mini-tiktok/cmd/feed/constant"
	"github.com/linzijie1998/mini-tiktok/cmd/feed/dal/cache"
	"github.com/linzijie1998/mini-tiktok/cmd/feed/dal/db"
	"github.com/linzijie1998/mini-tiktok/cmd/feed/global"
	"github.com/linzijie1998/mini-tiktok/model"
	"github.com/linzijie1998/mini-tiktok/pkg/errno"
	"gorm.io/gorm"
)

// QueryUserInfoById 查询用户信息: 包含基本信息 (目前的场景为只读信息, 而实际情况应该为读多写少) 和计数信息 (经常会变更的数据);
// 数据查询顺序为先查Redis, 当Redis未命中时查MySQL, MySQL也未命中时使用空值缓存应对Redis缓存穿透问题;
// 当Redis命中时就更新一次过期时间, 防止出现Redis缓存雪崩的问题
func QueryUserInfoById(ctx context.Context, uid int64) (*model.User, error) {
	// 1. 空值缓存查询
	if err := cache.GetUserInfoNullKey(ctx, uid); err == nil {
		return nil, errno.UserNotRegisterErr
	}
	// 2. 基础信息查询
	userInfo, err := cache.GetUserInfo(ctx, uid)
	if err != nil {
		// 缓存未命中
		userInfo, err = db.QueryUserInfoByID(ctx, uid, constant.UserBaseInfoQueryString)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					// 未找到该UID的用户, 设置空值缓存防止Redis缓存穿透
					_ = cache.NewUserInfoNullKey(ctx, uid, global.ExpireDurationNullKey)
					return nil, errno.UserNotRegisterErr
				}
				return nil, err
			}
		} else {
			// 添加缓存信息（忽略错误信息）
			_ = cache.NewUserInfos(ctx, []*model.User{userInfo}, global.ExpireDurationUserBaseInfo)
		}
	}
	if len(userInfo.Avatar) == 0 {
		userInfo.Avatar = global.Configs.StaticResource.DefaultAvatar
	}
	if len(userInfo.BackgroundImage) == 0 {
		userInfo.BackgroundImage = global.Configs.StaticResource.DefaultBackgroundImage
	}
	// 3. 计数信息查询
	userCounter, err := cache.GetUserCounter(ctx, uid)
	if err != nil {
		userCounter, err = db.QueryUserInfoByID(ctx, uid, constant.UserCounterInfoQueryString)
		if err != nil {
			return nil, err
		} else {
			// 添加缓存信息（忽略错误信息）
			_ = cache.NewUserCounters(ctx, []*model.User{userCounter})
		}
	}
	// 4. 合并基础信息和计数信息
	userInfo.FollowCount = userCounter.FollowCount
	userInfo.FollowerCount = userCounter.FollowerCount
	userInfo.FavoriteCount = userCounter.FavoriteCount
	userInfo.WorkCount = userCounter.WorkCount
	userInfo.TotalFavorited = userCounter.TotalFavorited
	return userInfo, nil
}
