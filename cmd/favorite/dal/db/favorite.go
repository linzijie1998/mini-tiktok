package db

import (
	"context"

	"github.com/linzijie1998/mini-tiktok/cmd/favorite/global"
	"github.com/linzijie1998/mini-tiktok/model"
	"gorm.io/gorm"
)

func AddFavoriteInfo(ctx context.Context, favoriteInfos []*model.Favorite) error {
	return global.GormDB.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Create(&favoriteInfos).Error; err != nil {
			return err
		}
		for _, favoriteInfo := range favoriteInfos {
			if err := tx.WithContext(ctx).Model(&model.Video{}).Where("id = ?", favoriteInfo.VideoId).Update("favorite_count", gorm.Expr("favorite_count + 1")).Error; err != nil {
				return err
			}
			if err := tx.WithContext(ctx).Model(&model.User{}).Where("id = ?", favoriteInfo.UserId).Update("favorite_count", gorm.Expr("favorite_count + 1")).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func DeleteFavoriteInfo(ctx context.Context, uid, vid int64) error {
	return global.GormDB.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Where("user_id = ? AND video_id = ?", uid, vid).Delete(&model.Favorite{}).Error; err != nil {
			return err
		}
		if err := tx.WithContext(ctx).Model(&model.Video{}).Where("id = ?", vid).Update("favorite_count", gorm.Expr("favorite_count - 1")).Error; err != nil {
			return err
		}
		if err := tx.WithContext(ctx).Model(&model.User{}).Where("id = ?", uid).Update("favorite_count", gorm.Expr("favorite_count - 1")).Error; err != nil {
			return err
		}
		return nil
	})
}

func QueryFavoriteInfosByUserId(ctx context.Context, uid int64, query string) ([]*model.Favorite, error) {
	res := make([]*model.Favorite, 0)
	if err := global.GormDB.WithContext(ctx).Select(query).Where("user_id = ?", uid).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func QueryFavoriteInfo(ctx context.Context, uid, vid int64) error {
	return global.GormDB.WithContext(ctx).Select("id").Where("user_id = ? AND video_id = ?", uid, vid).First(&model.Favorite{}).Error
}
