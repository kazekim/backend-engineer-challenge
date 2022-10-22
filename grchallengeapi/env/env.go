package grcenv

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/begit"
	"github.com/kazekim/backend-engineer-challenge/grlib/besqlx"
)

// Environment environment
type Environment struct {
	Release        bool          `mapstructure:"release"`
	ServerName     string        `mapstructure:"server_name"`
	ServerHost     string        `mapstructure:"server_host"`
	ServerPort     string        `mapstructure:"server_port"`
	DatabaseConfig besqlx.Config `mapstructure:"legacy_database_config"`
	GitConfig      begit.Config  `mapstructure:"git_config"`
}
