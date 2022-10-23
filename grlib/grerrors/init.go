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
func NewBadInputError(err error) Error {
	return &defaultError{
		CodeValue:    ErrCodeBadInput,
		MessageValue: fmt.Sprintf("%v: %v", ErrMessageBadInput, err.Error()),
	}
}

//NewBadInputErrorWithMessage new error with bad input error code and specify message
func NewBadInputErrorWithMessage(message string) Error {
	return &defaultError{
		CodeValue:    ErrCodeBadInput,
		MessageValue: fmt.Sprintf("%v: %v", ErrMessageBadInput, message),
	}
}

//NewKafkaProducerError new error with kafka producer error code and specify message
func NewKafkaProducerError(err error) Error {
	return &defaultError{
		CodeValue:    ErrCodeKafkaProducerFailed,
		MessageValue: err.Error(),
	}
}

//NewKafkaConsumerError new error with kafka consumer error code and specify message
func NewKafkaConsumerError(err error) Error {
	return &defaultError{
		CodeValue:    ErrCodeKafkaConsumerFailed,
		MessageValue: err.Error(),
	}
}