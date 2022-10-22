package besqlx

import (
	"github.com/jmoiron/sqlx"
)

type Client interface {
	DB() *sqlx.DB
	ConnectPostgres() error
	ConnectMySQL() error

	NewUpdateHelper() *UpdateHelper
}

type defaultClient struct {
	cfg *Config
	db  *sqlx.DB
}

//NewClient init besqlx client
func NewClient(cfg *Config) Client {

	return &defaultClient{
		cfg: cfg,
	}
}
