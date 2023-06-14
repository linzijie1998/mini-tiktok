package config

type Upload struct {
	VideoPath string `mapstructure:"video_path" json:"video_path"`
	CoverPath string `mapstructure:"cover_path" json:"cover_path"`
}
