package cache

import (
	"context"
	"fmt"
	"github.com/linzijie1998/mini-tiktok/cmd/feed/global"
	"github.com/linzijie1998/mini-tiktok/pkg/errno"
	"strconv"
	"time"
)

const (
	userInfoKey        = "user_info_uid%d"
	userCounterKey     = "user_counter_uid%d"
	followInfoKey      = "relation_follow_uid%d"
	userInfoNullKey    = "user_null_uid%d"
	userLoginNullKey   = "user_null_username_%s"
	videoInfoKey       = "video_info_vid%d"
	videoCounterKey    = "video_counter_vid%d"
	publishKey         = "publish_uid%d"
	videoFavoriteKey   = "favorite_uid%d"
	publishInfoNullKey = "publish_null_uid%d"
	publishQueueKey    = "publish_queue"
)

func getUserInfoKey(uid int64) string {
	return fmt.Sprintf(userInfoKey, uid)
}

func getUserCounterKey(uid int64) string {
	return fmt.Sprintf(userCounterKey, uid)
}

func getFollowInfoKey(uid int64) string {
	return fmt.Sprintf(followInfoKey, uid)
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

func getPublishInfoNullKey(uid int64) string {
	return fmt.Sprintf(publishInfoNullKey, uid)
}

func getVideoInfoKey(vid int64) string {
	return fmt.Sprintf(videoInfoKey, vid)
}

func getVideoCounterKey(vid int64) string {
	return fmt.Sprintf(videoCounterKey, vid)
}

func getPublishKey(uid int64) string {
	return fmt.Sprintf(publishKey, uid)
}

func getVideoFavoriteKey(uid int64) string {
	return fmt.Sprintf(videoFavoriteKey, uid)
}

func change(ctx context.Context, key, field string, incr int64) error {
	before, err := global.RedisClient.HGet(ctx, key, field).Result()
	if err != nil {
		return err
	}
	beforeInt, err := strconv.ParseInt(before, 10, 64)
	if err != nil {
		return err
	}
	if beforeInt+incr < 0 {
		return errno.ParamErr
	}
	if _, err := global.RedisClient.HIncrBy(ctx, key, field, incr).Result(); err != nil {
		return err
	}
	return nil
}
