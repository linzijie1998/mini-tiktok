package db

import (
	"context"

	"github.com/linzijie1998/mini-tiktok/cmd/feed/global"
	"github.com/linzijie1998/mini-tiktok/model"
)

func QueryFavoriteInfo(ctx context.Context, uid, vid int64) error {
	return global.GormDB.WithContext(ctx).Select("id").Where("user_id = ? AND video_id = ?", uid, vid).First(&model.Favorite{}).Error
}
