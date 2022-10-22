package besqlx

import (
	"fmt"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

//DeleteExec execute delete query with where clause
func (c *defaultClient) DeleteExec(arg interface{}, whereQuery string, values ...interface{}) grerrors.Error {

	tableName, vErr := parseTableName(arg)
	if vErr != nil {
		return vErr
	}

	query := fmt.Sprintf("delete from %v where %v", tableName, whereQuery)
	_, err := c.db.Exec(query, values...)
	if err != nil {
		return grerrors.NewDatabaseError(err)
	}
	return nil
}