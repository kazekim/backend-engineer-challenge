package besqlx

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//ConnectPostgres link to database with postgres driver
func (c *defaultClient) ConnectPostgres() error {
	sslMode := "require"
	if c.cfg.DatabaseDisableSSLMode {
		sslMode = "disable"
	}
	var dataSource = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", c.cfg.DatabaseUrl, c.cfg.DatabasePort, c.cfg.DatabaseUsername, c.cfg.DatabasePassword, c.cfg.DatabaseName, sslMode)
	db, err := sqlx.Connect("postgres", dataSource)
	if err != nil {
		return err
	}
	c.db = db
	return nil
}
