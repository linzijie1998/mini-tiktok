package db

import (
	"context"
	"github.com/linzijie1998/mini-tiktok/cmd/favorite/global"
	"github.com/linzijie1998/mini-tiktok/model"
)

func QueryVideoInfoById(ctx context.Context, vid int64, query string) (*model.Video, error) {
	var video model.Video
	err := global.GormDB.WithContext(ctx).Select(query).Where("id = ?", vid).First(&video).Error
	if err != nil {
		return nil, err
	}
	return &video, nil
}
