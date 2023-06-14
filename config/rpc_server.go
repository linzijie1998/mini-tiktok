package config

import "fmt"

type RPCServer struct {
	ServiceName string `mapstructure:"service-name" json:"service-name"`
	Host        string `mapstructure:"host" json:"host"`
	Port        string `mapstructure:"port" json:"port"`
}

func (r *RPCServer) Addr() string {
	return fmt.Sprintf("%s:%s", r.Host, r.Port)
}
