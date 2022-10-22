package besqlx

import (
	"database/sql"
	"fmt"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"reflect"
)

//NamedExec create tow in database with query from struct
func (c *defaultClient) NamedExec(paramQuery, valueQuery string, arg interface{}) (sql.Result, grerrors.Error) {
	st := reflect.TypeOf(arg)
	tableName := ""
	if obj, ok := arg.(interface{TableName() string}); ok {
		tableName = obj.TableName()
	}else{
		eMsg := fmt.Sprintf("dao %v: has no TableName function defined", st.Name())
		return nil, grerrors.NewDatabaseErrorWithMessage(eMsg)
	}

	query := fmt.Sprintf("insert into %v (%v) values (%v)", tableName, paramQuery, valueQuery)

	if obj, ok := arg.(interface{BeforeCreate() (err error)}); ok {
		err := obj.BeforeCreate()
		if err != nil {
			return nil, grerrors.NewDatabaseError(err)
		}
	}

	result, err := c.DB().NamedExec(query, arg)
	if err != nil {
		return nil, grerrors.NewDatabaseError(err)
	}



	return result, nil
}
