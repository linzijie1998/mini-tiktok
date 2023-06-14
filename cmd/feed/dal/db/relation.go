package db

import (
	"context"

	"github.com/linzijie1998/mini-tiktok/cmd/feed/global"
	"github.com/linzijie1998/mini-tiktok/model"
)

func QueryFollowInfo(ctx context.Context, userId, followUserId int64, query string) error {
	return global.GormDB.WithContext(ctx).Select(query).Where("user_id = ? AND follow_user_id = ?", userId, followUserId).First(&model.Relation{}).Error
}
