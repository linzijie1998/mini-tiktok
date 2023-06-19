package config

type FileAccess struct {
	UploadPath string `mapstructure:"upload_path" json:"upload_path" yaml:"upload_path"` // 文件存储地址
	NginxUrl   string `mapstructure:"nginx_url" json:"nginx_url" yaml:"nginx_url"`       // nginx访问地址
}
