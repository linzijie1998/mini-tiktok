package config

import "github.com/linzijie1998/mini-tiktok/config"

type ServiceConfigs struct {
	MySQL          config.MySQL          `mapstructure:"mysql" yaml:"mysql"`
	Redis          config.Redis          `mapstructure:"redis" yaml:"redis"`
	JWT            config.JWT            `mapstructure:"jwt" yaml:"jwt"`
	ETCD           config.ETCD           `mapstructure:"etcd" yaml:"etcd"`
	RPCServer      config.RPCServer      `mapstructure:"rpc_server" yaml:"rpc_server"`
	StaticResource config.StaticResource `mapstructure:"static_resource" yaml:"static_resource"`
	FileAccess     config.FileAccess     `mapstructure:"file_access" yaml:"file_access"`
}
