package grgitscanner

import "github.com/kazekim/backend-engineer-challenge/grlib/begit"

type Config struct {
	RootPath string `json:"root_path" mapstructure:"root_path"`
	ChannelAmountPerRequest string `json:"channel_amount_per_request" mapstructure:"channel_amount_per_request"`
}

// GitConfig return begit config from grcgitscanner config
func (cfg *Config) GitConfig() *begit.Config {
	return &begit.Config{
		RootPath: cfg.RootPath,
	}
}
