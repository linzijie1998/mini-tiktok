package pack

import (
	"errors"
	"github.com/linzijie1998/mini-tiktok/model"
)

func MergeUserInfo(baseInfo, counterInfo *model.User) error {
	if baseInfo == nil || counterInfo == nil {
		return errors.New("found nil")
	}
	baseInfo.FollowCount = counterInfo.FollowCount
	baseInfo.FollowerCount = counterInfo.FollowerCount
	baseInfo.FavoriteCount = counterInfo.FavoriteCount
	baseInfo.WorkCount = counterInfo.WorkCount
	baseInfo.TotalFavorited = counterInfo.TotalFavorited
	return nil
}
