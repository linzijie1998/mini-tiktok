package config

type Play struct {
	VideoURL string `mapstructure:"video_url" json:"video_url" yaml:"video_url"`
	CoverURL string `mapstructure:"cover_url" json:"cover_url" yaml:"cover_url"`
}
