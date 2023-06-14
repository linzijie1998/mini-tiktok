package db

import (
	"context"

	"github.com/linzijie1998/mini-tiktok/cmd/publish/global"
	"github.com/linzijie1998/mini-tiktok/model"
)

func QueryVideoInfoByHash(ctx context.Context, hash string, query string) (*model.Video, error) {
	var video model.Video
	err := global.GormDB.WithContext(ctx).Select(query).Where("hash = ?", hash).First(&video).Error
	if err != nil {
		return nil, err
	}
	return &video, nil
}

func CreateVideoInfo(ctx context.Context, videoInfos []*model.Video) error {
	return global.GormDB.WithContext(ctx).Create(&videoInfos).Error
}

func QueryVideoInfoByUserId(ctx context.Context, uid int64, query string) ([]*model.Video, error) {
	var video []*model.Video
	err := global.GormDB.WithContext(ctx).Select(query).Where("author_id = ?", uid).Find(&video).Error
	if err != nil {
		return nil, err
	}
	return video, nil
}

func QueryVideoInfoById(ctx context.Context, vid int64, query string) (*model.Video, error) {
	var video model.Video
	err := global.GormDB.WithContext(ctx).Select(query).Where("id = ?", vid).First(&video).Error
	if err != nil {
		return nil, err
	}
	return &video, nil
}
