package besqlx

import (
	"database/sql"
	"fmt"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

// NamedExec create row in database with query from struct
func (c *defaultClient) NamedExec(paramQuery, valueQuery string, arg interface{}) (sql.Result, grerrors.Error) {

	tableName, vErr := parseTableName(arg)
	if vErr != nil {
		return nil, vErr
	}

	query := fmt.Sprintf("insert into %v (%v) values (%v)", tableName, paramQuery, valueQuery)

	if obj, ok := arg.(interface{ BeforeCreate() (err error) }); ok {
		err := obj.BeforeCreate()
		if err != nil {
			return nil, grerrors.NewDatabaseError(err)
		}
	}

	result, err := c.DB().NamedExec(query, arg)
	if err != nil {
		return nil, grerrors.NewDatabaseError(err)
	}

	if obj, ok := arg.(interface{ AfterCreate() (err error) }); ok {
		err := obj.AfterCreate()
		if err != nil {
			return nil, grerrors.NewDatabaseError(err)
		}
	}

	return result, nil
}
