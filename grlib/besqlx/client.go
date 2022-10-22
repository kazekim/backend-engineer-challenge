package besqlx

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

type Client interface {
	DB() *sqlx.DB
	ConnectPostgres() error
	ConnectMySQL() error
	Connect() error

	NewUpdateHelper() *UpdateHelper

	NamedExec(paramQuery, valueQuery string, arg interface{}) (sql.Result, grerrors.Error)
	DeleteExec(arg interface{}, whereQuery string, values ...interface{}) grerrors.Error
	Get(dest interface{}, query string, args ...interface{}) error
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
