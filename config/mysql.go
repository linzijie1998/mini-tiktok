package config

import "fmt"

type MySQL struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`             // 服务器地址
	Port     string `mapstructure:"port" json:"port" yaml:"port"`             // 端口
	Config   string `mapstructure:"config" json:"config" yaml:"config"`       // 高级配置
	DBName   string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`    // 数据库名
	Username string `mapstructure:"username" json:"username" yaml:"username"` // 数据库用户名
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 数据库密码
}

func (m *MySQL) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", m.Username, m.Password, m.Host, m.Port, m.DBName, m.Config)
}
