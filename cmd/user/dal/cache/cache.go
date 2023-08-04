package cache

import (
	"context"
	"fmt"
	"github.com/linzijie1998/mini-tiktok/cmd/user/global"
	"time"
)

const (
	userInfoKey      = "user_info_uid%d"
	userCounterKey   = "user_counter_uid%d"
	userInfoNullKey  = "user_null_uid%d"
	userLoginNullKey = "user_null_username_%s"
)

func getUserInfoKey(uid int64) string {
	return fmt.Sprintf(userInfoKey, uid)
}

func getUserCounterKey(uid int64) string {
	return fmt.Sprintf(userCounterKey, uid)
}

func getUserInfoNullKey(uid int64) string {
	return fmt.Sprintf(userInfoNullKey, uid)
}

func getUserLoginNullKey(username string) string {
	return fmt.Sprintf(userLoginNullKey, username)
}

func addNullKey(ctx context.Context, key string, duration time.Duration) error {
	_, err := global.RedisClient.Set(ctx, key, "", duration).Result()
	return err
}

func getNullKey(ctx context.Context, key string) error {
	_, err := global.RedisClient.Get(ctx, key).Result()
	return err
}

func delNullKey(ctx context.Context, key string) error {
	_, err := global.RedisClient.Del(ctx, key).Result()
	return err
}
