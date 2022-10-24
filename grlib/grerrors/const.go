package grerrors

var (
	HttpErrorCode = 599

	ErrInternalServerError          = NewError(ErrCodeInternalServerError, ErrMessageInternalServerError)
	ErrBadInput                     = NewError(ErrCodeBadInput, ErrMessageBadInput)
	ErrInvalidUsernameOrPassword    = NewError(ErrCodeInvalidUsernameOrPassword, ErrMessageInvalidUsernameOrPassword)
	ErrGenerateAccessTokenError     = NewError(ErrCodeGenerateAccessTokenError, ErrMessageGenerateAccessTokenError)
	ErrAccessTokenNotExist          = NewError(ErrCodeAccessTokenNotExist, ErrMessageAccessTokenNotExist)
	ErrUserNotExist                 = NewError(ErrCodeUserNotExist, ErrMessageUserNotExist)
	ErrFileExceedSizeLimit          = NewError(ErrCodeFileExceedSizeLimit, ErrMessageFileExceedSizeLimit)
	ErrStorageUploadFileFailed      = NewError(ErrCodeStorageUploadFileFailed, ErrMessageStorageUploadFileFailed)
	ErrFileNotFound                 = NewError(ErrCodeFileNotFound, ErrMessageFileNotFound)
	ErrCreateDirectoryFailed        = NewError(ErrCodeCreateDirectoryFailed, ErrMessageCreateDirectoryFailed)
	ErrUnsupportedFileType          = NewError(ErrCodeUnsupportedFileType, ErrMessageUnsupportedFileType)
	ErrNoWhereClauseInDatabaseQuery = NewError(ErrCodeNoWhereClauseInDatabaseQuery, ErrMessageNoWhereClauseInDatabaseQuery)
	ErrGitHasNotBeenLoadedYet       = NewError(ErrCodeGitHasNotBeenLoadedYet, ErrMessageGitHasNotBeenLoadedYet)
	ErrDataNotFound                 = NewError(ErrCodeDataNotFound, ErrMessageDataNotFound)
)
