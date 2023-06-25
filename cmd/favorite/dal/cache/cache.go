package cache

import (
	"context"
	"fmt"
	"github.com/linzijie1998/mini-tiktok/cmd/favorite/global"
	"time"
)

const (
	videoInfoKey       = "video_info_vid%d"
	videoCounterKey    = "video_counter_vid%d"
	publishKey         = "publish_uid%d"
	videoFavoriteKey   = "favorite_uid%d"
	publishInfoNullKey = "publish_null_uid%d"
	publishQueueKey    = "publish_queue"
)

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

func getVideoFavoriteKey(uid int64) string {
	return fmt.Sprintf(videoFavoriteKey, uid)
}
