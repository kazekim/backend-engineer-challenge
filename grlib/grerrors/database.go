package grerrors

import (
	"fmt"
	"github.com/lib/pq"
)

//NewDatabaseError init database error from golang error
func NewDatabaseError(err error) Error {
	if dbErr, ok := err.(*pq.Error); ok {
		message := fmt.Sprintf("(%v) %v", dbErr.Code, dbErr.Message)
		return NewError(ErrCodeDatabaseError, message)
	} else {
		return NewError(ErrCodeDatabaseError, err.Error())
	}
}

//NewDatabaseErrorWithMessage init database error with error message
func NewDatabaseErrorWithMessage(message string) Error {
	return NewError(ErrCodeDatabaseError, message)
}
