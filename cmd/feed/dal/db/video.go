package db

import (
	"context"
	"time"

	"github.com/linzijie1998/mini-tiktok/cmd/feed/global"
	"github.com/linzijie1998/mini-tiktok/model"
)

func QueryVideoInfoWithLimit(ctx context.Context, limit int, latestTime time.Time, query string) ([]*model.Video, error) {
	res := make([]*model.Video, 0)
	if err := global.GormDB.WithContext(ctx).Select(query).Order("created_at desc").Where("created_at <= ?", latestTime).Limit(limit).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
