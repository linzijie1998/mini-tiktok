package config

import "github.com/linzijie1998/mini-tiktok/config"

type ServiceConfigs struct {
	ETCD      config.ETCD      `mapstructure:"etcd" yaml:"etcd"`
	Hertz     config.Hertz     `mapstructure:"hertz" yaml:"hertz"`
	RPCClient config.RPCClient `mapstructure:"rpc_client" yaml:"rpc_client"`
}
