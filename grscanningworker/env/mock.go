package grwenv

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/besqlx"
	"github.com/kazekim/backend-engineer-challenge/grlib/grscankafka"
)

var (
	EnvMock = Environment{
		Release: true,
		ServerHost: "localhost",
		ServerName: "mockserver",
		ServerPort: "9999",
		DatabaseConfig: besqlx.Config{
			DatabaseDriver: "postgres",
			DatabaseUrl: "localhost",
			DatabasePort: "5432",
			DatabaseName: "database",
			DatabaseUsername: "username",
			DatabasePassword: "password",
			DatabaseLogEnabled: true,
			DatabaseDisableSSLMode: true,
		},
		GitScannerMQConfig: grscankafka.Config{
			BrokerUrl: "localhost:9092",
			ServerId: "testserver",
		},
	}
)
