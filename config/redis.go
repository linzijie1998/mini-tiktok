package config

import "fmt"

type Redis struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"` // 服务器地址
	Port     string `mapstructure:"port" json:"port" yaml:"port"` // 端口
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}

func (r *Redis) Addr() string {
	return fmt.Sprintf("%s:%s", r.Host, r.Port)
}
