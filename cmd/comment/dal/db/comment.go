package db

import (
	"context"

	"github.com/linzijie1998/mini-tiktok/cmd/comment/global"
	"github.com/linzijie1998/mini-tiktok/model"
	"gorm.io/gorm"
)

func CreateCommentInfos(ctx context.Context, commentInfos []*model.Comment) error {
	return global.GormDB.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Create(&commentInfos).Error; err != nil {
			return err
		}
		for _, commentInfo := range commentInfos {
			if err := tx.WithContext(ctx).Model(&model.Video{}).Where("id = ?", commentInfo.VideoId).Update("comment_count", gorm.Expr("comment_count + 1")).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func DeleteCommentInfo(ctx context.Context, cid, vid int64) error {
	return global.GormDB.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Where("id = ?", cid).Delete(&model.Comment{}).Error; err != nil {
			return err
		}
		if err := tx.WithContext(ctx).Model(&model.Video{}).Where("id = ?", vid).Update("comment_count", gorm.Expr("comment_count - 1")).Error; err != nil {
			return err
		}
		return nil
	})
}

func QueryCommentInfos(ctx context.Context, vid int64, limit int, query string) ([]*model.Comment, error) {
	res := make([]*model.Comment, 0)
	if err := global.GormDB.WithContext(ctx).Select(query).Where("video_id = ?", vid).Order("created_at desc").Limit(limit).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
