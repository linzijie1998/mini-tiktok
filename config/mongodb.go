package config

import "fmt"

type MongoDB struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Database string `mapstructure:"database" json:"database" yaml:"database"`
}

func (m *MongoDB) Url() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s", m.Username, m.Password, m.Host, m.Port)
}
