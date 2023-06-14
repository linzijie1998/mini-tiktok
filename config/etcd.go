package config

import "fmt"

type ETCD struct {
	Host string `mapstructure:"host" json:"host"`
	Port string `mapstructure:"port" json:"port"`
}

func (e *ETCD) Addr() string {
	return fmt.Sprintf("%s:%s", e.Host, e.Port)
}
