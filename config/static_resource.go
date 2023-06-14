package config

type StaticResource struct {
	DefaultAvatar          string `mapstructure:"default-avatar" json:"default-avatar"`
	DefaultBackgroundImage string `mapstructure:"default-background-image" json:"default-background-image"`
}
