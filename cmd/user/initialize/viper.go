package initialize

import (
	"github.com/linzijie1998/mini-tiktok/cmd/user/global"
	"github.com/spf13/viper"
	"time"
)

func parseDuration() error {
	var err error
	global.ExpireDurationNullKey, err = time.ParseDuration(global.Configs.CacheExpire.NullKey)
	if err != nil {
		return err
	}
	global.ExpireDurationUserBaseInfo, err = time.ParseDuration(global.Configs.CacheExpire.UserBaseInfo)
	if err != nil {
		return err
	}
	return nil
}

func Viper(path string) error {
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	if err := v.Unmarshal(&global.Configs); err != nil {
		return err
	}
	if err := parseDuration(); err != nil {
		panic(err)
	}
	return nil
}
