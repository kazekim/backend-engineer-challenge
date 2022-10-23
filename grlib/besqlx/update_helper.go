package besqlx

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

type UpdateHelper struct {
	params      map[string]interface{}
	whereParams map[string]interface{}
	db          *sqlx.DB
}

//NewUpdateHelper init code helper to make update database query code more clean
func (c defaultClient) NewUpdateHelper() *UpdateHelper {

	return &UpdateHelper{
		params:      make(map[string]interface{}),
		whereParams: make(map[string]interface{}),
		db:          c.DB(),
	}
}

//SetWhereParam add where parameter to helper
func (h *UpdateHelper) SetWhereParam(key string, value interface{}) {
	h.whereParams[key] = value
}

//SetParam add parameter to helper
func (h *UpdateHelper) SetParam(key string, value interface{}) {
	h.params[key] = value
}

//CommitUpdateQuery do update query
func (h *UpdateHelper) CommitUpdateQuery(model interface{}) grerrors.Error {

	tableName, vErr := parseTableName(model)
	if vErr != nil {
		return vErr
	}

	var args []interface{}

	if !h.HasChanged() {
		return nil
	}

	updateQuery := ""
	isNotFirstUpdateClause := false
	count := 1
	for k, v := range h.params {
		if isNotFirstUpdateClause {
			updateQuery = fmt.Sprintf("%v%v", updateQuery, ", ")
		} else {
			isNotFirstUpdateClause = true
		}
		updateQuery = fmt.Sprintf("%v%v = $%d", updateQuery, k, count)

		args = append(args, v)
		count++
	}

	whereQuery := ""
	if len(h.whereParams) > 0 {
		isNotFirstWhereClause := false
		for k, v := range h.whereParams {
			if isNotFirstWhereClause {
				whereQuery = fmt.Sprintf("%v and ", whereQuery)
			} else {
				isNotFirstWhereClause = true
			}
			whereQuery = fmt.Sprintf("%v%v = $%d", whereQuery, k, count)

			args = append(args, v)
			count++
		}
	} else {
		return grerrors.ErrNoWhereClauseInDatabaseQuery
	}

	query := fmt.Sprintf("update %v set %v where %v", tableName, updateQuery, whereQuery)
	res, err := h.db.Exec(query, args...)
	if err != nil {
		return grerrors.NewDatabaseError(err)
	}
	_, err = res.RowsAffected()
	if err != nil {
		return grerrors.NewDatabaseError(err)
	}
	return nil
}

//HasChanged check if it has any parameters added to helper or not
func (h *UpdateHelper) HasChanged() bool {
	return len(h.params) > 0
}
