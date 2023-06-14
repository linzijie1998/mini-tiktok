package db

import (
	"context"

	"github.com/linzijie1998/mini-tiktok/cmd/message/global"
	"github.com/linzijie1998/mini-tiktok/model"
)

func CreateMessageInfos(ctx context.Context, messageInfos []*model.Message) error {
	return global.GormDB.WithContext(ctx).Create(&messageInfos).Error
}

func QueryMessageByUserIDAndToUserID(ctx context.Context, userID int64, toUserID int64, query string) (messages []model.Message, err error) {
	err = global.GormDB.WithContext(ctx).Select(query).Find(&messages, "user_id = ? and to_user_id = ?", userID, toUserID).Error
	return
}

func QueryMessageByUserIDAndToUserIDWithLimit(ctx context.Context, userID int64, toUserID int64, limit int64, query string) (messages []model.Message, err error) {
	err = global.GormDB.WithContext(ctx).Select(query).Where("publish_date > ?", limit).Find(&messages, "user_id = ? and to_user_id = ?", userID, toUserID).Error
	return
}
