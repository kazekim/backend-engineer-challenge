package besqlx

import (
	"fmt"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"reflect"
)

func parseTableName(arg interface{}) (string, grerrors.Error) {
	st := reflect.TypeOf(arg)
	tableName := ""
	if obj, ok := arg.(interface{TableName() string}); ok {
		tableName = obj.TableName()
	}else{
		eMsg := fmt.Sprintf("dao %v: has no TableName function defined", st.Name())
		return "", grerrors.NewDatabaseErrorWithMessage(eMsg)
	}
	return tableName, nil
}
