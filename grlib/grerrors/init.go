package grerrors

import "fmt"

//NewDefaultError new error with internal error code and specify message
func NewDefaultError(message error) Error {
	return &defaultError{
		CodeValue:    ErrCodeInternalServerError,
		MessageValue: message.Error(),
	}
}

//NewBadInputError new error with bad input error code and specify message
func NewBadInputError(message string) Error {
	vErr := NewError(ErrCodeBadInput, fmt.Sprintf("%v: %v", ErrMessageBadInput, message))
	return vErr
}
