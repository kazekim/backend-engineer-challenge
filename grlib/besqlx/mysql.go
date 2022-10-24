package besqlx

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// ConnectMySQL link to database with mysql driver
func (c *defaultClient) ConnectMySQL() error {

	var dataSource = fmt.Sprintf(`%v:%v@(%v:%v)/%v?parseTime=true&locAsia%%2FBangkok`, c.cfg.DatabaseUsername, c.cfg.DatabasePassword, c.cfg.DatabaseUrl, c.cfg.DatabasePort, c.cfg.DatabaseName)

	db, err := sqlx.Connect("mysql", dataSource)
	if err != nil {
		return err
	}
	c.db = db
	return nil
}
