package db

import (
	"context"

	"github.com/linzijie1998/mini-tiktok/cmd/publish/global"
	"github.com/linzijie1998/mini-tiktok/model"
	"gorm.io/gorm"
)

func AddPublishInfo(ctx context.Context, uid int64, videoInfo *model.Video) error {
	return global.GormDB.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Create(&videoInfo).Error; err != nil {
			return err
		}
		if err := tx.WithContext(ctx).Model(&model.User{}).Where("id = ?", uid).Update("work_count", gorm.Expr("work_count + 1")).Error; err != nil {
			return err
		}
		return nil
	})
}
