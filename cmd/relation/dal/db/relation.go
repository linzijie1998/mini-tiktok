package db

import (
	"context"

	"github.com/linzijie1998/mini-tiktok/cmd/relation/global"
	"github.com/linzijie1998/mini-tiktok/model"
	"gorm.io/gorm"
)

func CreateRelationInfos(ctx context.Context, relationInfos []*model.Relation) error {
	return global.GormDB.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Create(&relationInfos).Error; err != nil {
			return err
		}
		for _, relationInfo := range relationInfos {
			if err := tx.WithContext(ctx).Model(&model.User{}).Where("id = ?", relationInfo.UserId).Update("follow_count", gorm.Expr("follow_count + 1")).Error; err != nil {
				return err
			}
			if err := tx.WithContext(ctx).Model(&model.User{}).Where("id = ?", relationInfo.FollowUserId).Update("follower_count", gorm.Expr("follower_count + 1")).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func DeleteRelationInfo(ctx context.Context, userId, followUserId int64) error {
	return global.GormDB.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Where("user_id = ? AND follow_user_id = ?", userId, followUserId).Delete(&model.Relation{}).Error; err != nil {
			return err
		}
		if err := tx.WithContext(ctx).Model(&model.User{}).Where("id = ?", userId).Update("follow_count", gorm.Expr("follow_count - 1")).Error; err != nil {
			return err
		}
		if err := tx.WithContext(ctx).Model(&model.User{}).Where("id = ?", followUserId).Update("follower_count", gorm.Expr("follower_count - 1")).Error; err != nil {
			return err
		}
		return nil
	})
}

func QueryFollowInfos(ctx context.Context, uid int64, query string) ([]*model.Relation, error) {
	res := make([]*model.Relation, 0)
	if err := global.GormDB.WithContext(ctx).Select(query).Where("user_id = ?", uid).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func QueryFollowerInfos(ctx context.Context, uid int64, query string) ([]*model.Relation, error) {
	res := make([]*model.Relation, 0)
	if err := global.GormDB.WithContext(ctx).Select(query).Where("follow_user_id = ?", uid).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func QueryFollowInfo(ctx context.Context, userId, followUserId int64, query string) error {
	return global.GormDB.WithContext(ctx).Select(query).Where("user_id = ? AND follow_user_id = ?", userId, followUserId).First(&model.Relation{}).Error
}
