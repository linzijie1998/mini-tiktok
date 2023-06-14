package model

type Message struct {
	DefaultModel
	ToUserId   int64  `mapstructure:"to_user_id" gorm:"comment:接收用户ID"`
	FromUserId int64  `mapstructure:"from_user_id" gorm:"comment:发送用户ID"`
	Content    string `mapstructure:"content" gorm:"comment:消息内容"`
	CreateTime string `mapstructure:"create_time" gorm:"comment:发送时间"`
}
