package grwenv

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/besqlx"
	"github.com/kazekim/backend-engineer-challenge/grlib/grgitscanner"
	"github.com/kazekim/backend-engineer-challenge/grlib/grscankafka"
)

// Environment environment
type Environment struct {
	Release        bool                `mapstructure:"release"`
	ServerName     string              `mapstructure:"server_name"`
	ServerHost     string              `mapstructure:"server_host"`
	ServerPort     string              `mapstructure:"server_port"`
	DatabaseConfig besqlx.Config       `mapstructure:"database_config"`
	GitConfig      grgitscanner.Config `mapstructure:"git_config"`
	GitScannerMQConfig grscankafka.Config `mapstructure:"git_scanner_mq_config"`
}
