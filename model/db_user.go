package model

type User struct {
	DefaultModel
	//UUID            uuid.UUID `mapstructure:"-" gorm:"index;comment:用户UUID"`
	Username        string `mapstructure:"-" gorm:"comment:用户登录名"`
	Password        string `mapstructure:"-" gorm:"comment:用户登录密码"`
	Nickname        string `mapstructure:"nickname" gorm:"comment:用户昵称"`
	Avatar          string `mapstructure:"avatar" gorm:"comment:头像"`
	BackgroundImage string `mapstructure:"background_image" gorm:"comment:主页背景图"`
	Signature       string `mapstructure:"signature" gorm:"default:系统原装签名，送给每个小可爱。;comment:个性签名"`
	FollowCount     int64  `mapstructure:"follow_count" gorm:"default:0;comment:关注数"`
	FollowerCount   int64  `mapstructure:"follower_count" gorm:"default:0;comment:粉丝数"`
	TotalFavorited  int64  `mapstructure:"total_favorited" gorm:"default:0;comment:获赞数"`
	FavoriteCount   int64  `mapstructure:"favorite_count" gorm:"default:0;comment:点赞数"`
	WorkCount       int64  `mapstructure:"work_count" gorm:"default:0;comment:作品数"`
}
