package cache

import (
	"context"
	"errors"
	"github.com/linzijie1998/mini-tiktok/cmd/relation/global"
	"github.com/linzijie1998/mini-tiktok/model"
	"strconv"
	"time"
)

func buildUserInfoMap(user *model.User) map[string]interface{} {
	if user == nil {
		return nil
	}
	return map[string]interface{}{
		"id":               user.Id,
		"nickname":         user.Nickname,
		"avatar":           user.Avatar,
		"background_image": user.BackgroundImage,
		"signature":        user.Signature,
	}
}

func parserUserInfoMap(userInfoMap map[string]string) (*model.User, error) {
	userInfo := new(model.User)
	var ok bool
	if value, ok := userInfoMap["id"]; ok {
		parseInt, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, err
		}
		userInfo.Id = parseInt
	} else {
		return nil, errors.New("missing field")
	}
	if userInfo.Nickname, ok = userInfoMap["nickname"]; !ok {
		return nil, errors.New("missing field")
	}
	if userInfo.Avatar, ok = userInfoMap["avatar"]; !ok {
		return nil, errors.New("missing field")
	}
	if userInfo.BackgroundImage, ok = userInfoMap["background_image"]; !ok {
		return nil, errors.New("missing field")
	}
	if userInfo.Signature, ok = userInfoMap["signature"]; !ok {
		return nil, errors.New("missing field")
	}
	return userInfo, nil
}

func buildUserCounterMap(user *model.User) map[string]interface{} {
	if user == nil {
		return nil
	}
	return map[string]interface{}{
		"id":              user.Id,
		"follow_count":    user.FollowCount,
		"follower_count":  user.FollowerCount,
		"favorite_count":  user.FavoriteCount,
		"work_count":      user.WorkCount,
		"total_favorited": user.TotalFavorited,
	}
}

func parserUserCounterMap(userCounterMap map[string]string) (*model.User, error) {
	userCounter := new(model.User)
	var err error
	if value, ok := userCounterMap["id"]; ok {
		userCounter.Id, err = strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("missing field")
	}

	if value, ok := userCounterMap["follow_count"]; ok {
		userCounter.FollowCount, err = strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("missing field")
	}

	if value, ok := userCounterMap["follower_count"]; ok {
		userCounter.FollowerCount, err = strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("missing field")
	}

	if value, ok := userCounterMap["favorite_count"]; ok {
		userCounter.FavoriteCount, err = strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("missing field")
	}

	if value, ok := userCounterMap["work_count"]; ok {
		userCounter.WorkCount, err = strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("missing field")
	}

	if value, ok := userCounterMap["total_favorited"]; ok {
		userCounter.TotalFavorited, err = strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("missing field")
	}

	return userCounter, nil
}

func NewUserInfos(ctx context.Context, users []*model.User, duration time.Duration) error {
	pipe := global.RedisClient.Pipeline()
	for _, user := range users {
		key := getUserInfoKey(user.Id)
		if _, err := pipe.Del(ctx, key).Result(); err != nil {
			return err
		}
		userInfoMap := buildUserInfoMap(user)
		if _, err := pipe.HMSet(ctx, key, userInfoMap).Result(); err != nil {
			return err
		}
		if _, err := pipe.Expire(ctx, key, duration).Result(); err != nil {
			return err
		}
	}
	if _, err := pipe.Exec(ctx); err != nil {
		// 报错后进行一次额外尝试
		if _, err = pipe.Exec(ctx); err != nil {
			return err
		}
	}
	return nil
}

func NewUserCounters(ctx context.Context, users []*model.User) error {
	pipe := global.RedisClient.Pipeline()
	for _, user := range users {
		key := getUserCounterKey(user.Id)
		if _, err := pipe.Del(ctx, key).Result(); err != nil {
			return err
		}
		userInfoMap := buildUserCounterMap(user)
		if _, err := pipe.HMSet(ctx, key, userInfoMap).Result(); err != nil {
			return err
		}
	}
	if _, err := pipe.Exec(ctx); err != nil {
		// 报错后进行一次额外尝试
		if _, err = pipe.Exec(ctx); err != nil {
			return err
		}
	}
	return nil
}

func GetUserInfo(ctx context.Context, uid int64) (*model.User, error) {
	key := getUserInfoKey(uid)
	userInfoMap, err := global.RedisClient.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return parserUserInfoMap(userInfoMap)
}

func GetUserCounter(ctx context.Context, uid int64) (*model.User, error) {
	key := getUserCounterKey(uid)
	userCounterMap, err := global.RedisClient.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return parserUserCounterMap(userCounterMap)
}

func NewUserInfoNullKey(ctx context.Context, uid int64, duration time.Duration) error {
	key := getUserInfoNullKey(uid)
	return addNullKey(ctx, key, duration)
}

func GetUserInfoNullKey(ctx context.Context, uid int64) error {
	key := getUserInfoNullKey(uid)
	return getNullKey(ctx, key)
}

func DelUserInfoNullKey(ctx context.Context, uid int64) error {
	key := getUserInfoNullKey(uid)
	return delNullKey(ctx, key)
}

func NewUserLoginNullKey(ctx context.Context, username string, duration time.Duration) error {
	key := getUserLoginNullKey(username)
	return addNullKey(ctx, key, duration)
}

func GetUserLoginNullKey(ctx context.Context, username string) error {
	key := getUserLoginNullKey(username)
	return getNullKey(ctx, key)
}

func DelUserLoginNullKey(ctx context.Context, username string) error {
	key := getUserLoginNullKey(username)
	return delNullKey(ctx, key)
}
