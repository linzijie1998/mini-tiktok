package config

import "fmt"

type Hertz struct {
	Host string `mapstructure:"host" json:"host"`
	Port string `mapstructure:"port" json:"port"`
}

func (h *Hertz) Addr() string {
	return fmt.Sprintf("%s:%s", h.Host, h.Port)
}
