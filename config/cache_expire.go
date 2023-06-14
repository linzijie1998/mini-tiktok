package config

import "time"

type CacheExpire struct {
	NullKey       string `mapstructure:"null-key" json:"null-key" yaml:"null-key"`                   // 空值缓存过期时间
	UserBaseInfo  string `mapstructure:"user-base-info" json:"user-base-info" yaml:"user-base-info"` // 用户基本信息过期时间
	VideoBaseInfo string `mapstructure:"video-info" json:"video-info" yaml:"video-info"`             // 视频基本信息过期时间
}

func (c *CacheExpire) ParseNullKeyExpireDuration() time.Duration {
	duration, err := time.ParseDuration(c.NullKey)
	if err != nil {
		panic(err)
	}
	return duration
}

func (c *CacheExpire) ParseUserBaseInfoExpireDuration() time.Duration {
	duration, err := time.ParseDuration(c.UserBaseInfo)
	if err != nil {
		panic(err)
	}
	return duration
}

func (c *CacheExpire) ParseVideoBaseInfoExpireDuration() time.Duration {
	duration, err := time.ParseDuration(c.VideoBaseInfo)
	if err != nil {
		panic(err)
	}
	return duration
}
