package db

import (
	"context"

	"github.com/linzijie1998/mini-tiktok/cmd/feed/global"
	"github.com/linzijie1998/mini-tiktok/model"
)

func QueryFirstUserInfoByID(ctx context.Context, id int64, query string) (*model.User, error) {
	var user model.User
	err := global.GormDB.WithContext(ctx).Select(query).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
