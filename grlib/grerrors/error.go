package grerrors

import "fmt"

type ErrorCode int64

//Int64 return error code in int64 format
func (e ErrorCode) Int64() int64 {
	return int64(e)
}

//Int return error code in int format
func (e ErrorCode) Int() int {
	return int(e)
}

//String return error code in string format
func (e ErrorCode) String() string {
	return fmt.Sprintf("%d", e.Int64())
}

type Error interface {
	Code() ErrorCode
	Message() string
	Error() string
}

type defaultError struct {
	CodeValue    ErrorCode
	MessageValue string
	Log          string
}

//NewError new error from code and message
func NewError(code ErrorCode, message string) Error {
	return &defaultError{
		CodeValue:    code,
		MessageValue: message,
	}
}

//Code return error code
func (e *defaultError) Code() ErrorCode {
	return e.CodeValue
}

//Message return error message
func (e *defaultError) Message() string {
	return e.MessageValue
}

//Message interface func of golang error to return error
func (e *defaultError) Error() string {
	return fmt.Sprintf("%v: %v", e.CodeValue, e.MessageValue)
}
