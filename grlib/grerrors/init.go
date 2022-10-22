package grerrors

import "fmt"

//NewDefaultError new error with internal error code and specify error
func NewDefaultError(err error) Error {
	return &defaultError{
		CodeValue:    ErrCodeInternalServerError,
		MessageValue: err.Error(),
	}
}

//NewDefaultErrorWithMessage new error with internal error code and specify message
func NewDefaultErrorWithMessage(message string) Error {
	return &defaultError{
		CodeValue:    ErrCodeInternalServerError,
		MessageValue: message,
	}
}

//NewBadInputError new error with bad input error code and specify error
func NewBadInputError(err string) Error {
	vErr := NewError(ErrCodeBadInput, fmt.Sprintf("%v: %v", ErrMessageBadInput, err))
	return vErr
}

//NewBadInputErrorWithMessage new error with bad input error code and specify message
func NewBadInputErrorWithMessage(message string) Error {
	vErr := NewError(ErrCodeBadInput, fmt.Sprintf("%v: %v", ErrMessageBadInput, message))
	return vErr
}
