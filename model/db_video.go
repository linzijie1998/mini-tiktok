package model

type Video struct {
	DefaultModel
	AuthorId      int64  `mapstructure:"author_id" gorm:"comment:作者ID"`
	Title         string `mapstructure:"title" gorm:"comment:标题"`
	VideoPath     string `mapstructure:"video_path" gorm:"comment:视频地址/路径"`
	CoverPath     string `mapstructure:"cover_path" gorm:"comment:封面地址/路径"`
	FavoriteCount int64  `mapstructure:"favorite_count" gorm:"default:0;comment:获赞数"`
	CommentCount  int64  `mapstructure:"comment_count" gorm:"default:0;comment:评论数"`
}
