package model

type Relation struct {
	DefaultModel
	UserId       int64 `mapstructure:"user_id" gorm:"comment:用户ID"`
	FollowUserId int64 `mapstructure:"follow_user_id" gorm:"comment:关注的用户ID"`
}
