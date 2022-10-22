package besqlx

import "errors"

func (c *defaultClient) Connect() error {

	switch c.cfg.DatabaseDriver {
	case "postgres":
		return c.ConnectPostgres()
	case "mysql":
		return c.ConnectMySQL()
	default:
		return errors.New("invalid database driver")
	}
}
