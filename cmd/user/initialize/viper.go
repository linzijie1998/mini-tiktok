package initialize

import (
	"github.com/linzijie1998/mini-tiktok/cmd/user/global"
	"github.com/spf13/viper"
)

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
