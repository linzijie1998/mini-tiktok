package model

type Favorite struct {
	DefaultModel
	UserId  int64 `mapstructure:"user_id" gorm:"comment:用户ID"`
	VideoId int64 `mapstructure:"video_id" gorm:"comment:点赞的视频ID"`
}
