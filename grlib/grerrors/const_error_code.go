package grerrors

const (
	SuccessCode                         ErrorCode = 0
	ErrCodeInternalServerError          ErrorCode = 1000
	ErrCodeBadInput                     ErrorCode = 1001
	ErrCodeInvalidUsernameOrPassword    ErrorCode = 1002
	ErrCodeGenerateAccessTokenError     ErrorCode = 1003
	ErrCodeAccessTokenNotExist          ErrorCode = 1004
	ErrCodeUserNotExist                 ErrorCode = 1005
	ErrCodeFileExceedSizeLimit          ErrorCode = 1006
	ErrCodeStorageUploadFileFailed      ErrorCode = 1007
	ErrCodeFileNotFound                 ErrorCode = 1008
	ErrCodeCreateDirectoryFailed        ErrorCode = 1009
	ErrCodeUnsupportedFileType          ErrorCode = 1010
	ErrCodeDatabaseError                ErrorCode = 1011
	ErrCodeNoWhereClauseInDatabaseQuery ErrorCode = 1012
	ErrCodeGitHasNotBeenLoadedYet       ErrorCode = 1013
)
