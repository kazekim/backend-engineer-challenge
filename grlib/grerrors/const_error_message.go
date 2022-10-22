package grerrors

const (
	SuccessMessage                         = "Success"
	ErrMessageInternalServerError          = "internal server error"
	ErrMessageBadInput                     = "bad input"
	ErrMessageInvalidUsernameOrPassword    = "invalid username or password"
	ErrMessageGenerateAccessTokenError     = "generate access token error"
	ErrMessageAccessTokenNotExist          = "access token not exist"
	ErrMessageUserNotExist                 = "user is not exist"
	ErrMessageFileExceedSizeLimit          = "file exceed size limit"
	ErrMessageStorageUploadFileFailed      = "storage upload file failed"
	ErrMessageFileNotFound                 = "file not found"
	ErrMessageCreateDirectoryFailed        = "create directory failed"
	ErrMessageUnsupportedFileType          = "unsupported file type"
	ErrMessageCoinNotFound                 = "coin not found"
	ErrMessageNoWhereClauseInDatabaseQuery = "no where clause in database query"
	ErrMessageGitHasNotBeenLoadedYet       = "git has not been loaded yet"
)
