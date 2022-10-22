package grcgitscanner

import "github.com/kazekim/backend-engineer-challenge/grlib/begit"

type Config struct {
	RootPath string `json:"root_path" mapstructure:"root_path"`
}

//GitConfig return begit config from grcgitscanner config
func (cfg *Config) GitConfig() *begit.Config {
	return &begit.Config{
		RootPath: cfg.RootPath,
	}
}
