package db

import (
	"context"

	"github.com/linzijie1998/mini-tiktok/cmd/favorite/global"
	"github.com/linzijie1998/mini-tiktok/model"
)

func QueryVideoInfoById(ctx context.Context, id int64, query string) (*model.Video, error) {
	var res model.Video
	if err := global.GormDB.WithContext(ctx).Select(query).Where("id = ?", id).First(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}
