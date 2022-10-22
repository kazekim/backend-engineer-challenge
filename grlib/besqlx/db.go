package besqlx

import "github.com/jmoiron/sqlx"

//DB get sqlx client
func (c *defaultClient) DB() *sqlx.DB {
	return c.db
}

//Close shutdown db
func (c *defaultClient) Close() error {
	return c.db.Close()
}
