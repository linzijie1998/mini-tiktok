package model

type Comment struct {
	DefaultModel
	AuthorId   int64  `mapstructure:"author_id" gorm:"comment:作者ID"`
	VideoId    int64  `mapstructure:"video_id" gorm:"comment:视频ID"`
	Content    string `mapstructure:"content" gorm:"comment:评论内容"`
	CreateTime string `mapstructure:"create_time" gorm:"comment:评论时间"`
}
