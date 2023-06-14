package cache

import (
	"context"

	"github.com/linzijie1998/mini-tiktok/cmd/favorite/global"
)

func AddFavoriteVideo(ctx context.Context, uid, vid int64) error {
	key := getUserFavoriteKey(uid)
	_, err := global.RedisClient.SAdd(ctx, key, vid).Result()
	return err
}

func RemoveFavoriteVideo(ctx context.Context, uid, vid int64) error {
	key := getUserFavoriteKey(uid)
	_, err := global.RedisClient.SRem(ctx, key, vid).Result()
	return err
}
