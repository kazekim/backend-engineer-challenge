package besqlx

import (
	"fmt"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

type SelectQueryBuilder struct {
	SelectQuery string `json:"select_query"`
	WhereQuery string `json:"where_query"`
	OrderByQuery *string `json:"order_by_query"`
	Limit *int64 `json:"limit"`
	Page *int64 `json:"page"`
	Args []interface{} `json:"args"`
}

//Get query normal select query and return first value found in 'dest' struct
func (c *defaultClient) Get(dest interface{}, query string, args ...interface{}) grerrors.Error {
	err := c.db.Get(dest, query, args...)
	if err != nil {
		return grerrors.NewDatabaseError(err)
	}
	return nil
}

//Select do select query with select query builder
func (c *defaultClient) Select(model interface{}, dest interface{}, qb SelectQueryBuilder) grerrors.Error {

	tableName, vErr := parseTableName(model)
	if vErr != nil {
		return vErr
	}

	orderByQuery := ""
	if qb.OrderByQuery != nil {
		orderByQuery = fmt.Sprintf(" order by %v", *qb.OrderByQuery)
	}

	limitQuery := ""
	if qb.Page != nil && qb.Limit != nil {
		limit := *qb.Limit
		offset := ((*qb.Page -1) * *qb.Limit)
		limitQuery = fmt.Sprintf(" limit %v offset %v", limit, offset)
	}else if qb.Limit != nil {
		limitQuery = fmt.Sprintf(" limit %v", *qb.Limit)
	}

	query := fmt.Sprintf("select %v from %v where %v%v%v", qb.SelectQuery, tableName, qb.WhereQuery, orderByQuery, limitQuery)
	err := c.db.Select(dest, query, qb.Args...)
	if err != nil {
		return grerrors.NewDatabaseError(err)
	}

	return nil
}

//Count count all data in database with where clause
func (c *defaultClient) Count(model interface{}, whereQuery string, args []interface{}) (int64, grerrors.Error) {

	if whereQuery != "" {
		whereQuery = fmt.Sprintf(" where %v", whereQuery)
	}

	tableName, vErr := parseTableName(model)
	if vErr != nil {
		return 0, vErr
	}

	var count int64
	query := fmt.Sprintf("select count(*) as count from %v%v", tableName, whereQuery)
	rows, err := c.db.Query(query, args...)
	if err != nil {
		return 0, grerrors.NewDatabaseError(err)
	}
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return 0, grerrors.NewDatabaseError(err)
		}
	}

	return count, nil
}