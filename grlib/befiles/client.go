package befiles

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"mime/multipart"
)

type Client interface {
	StoreFileLocal(file *multipart.FileHeader) (*FileData, grerrors.Error)
	StoreImageBase64Local(data string) (*FileData, grerrors.Error)
	ForceDeleteDirectory(name string) grerrors.Error
	SetMaxFileSize(size int64)
	ConvertFileToBase64(file FileData) (*string, grerrors.Error)
}
type defaultClient struct {
	rootPath    string
	maxFileSize int64
}

//NewClient init new file client
func NewClient(rootPath string) (Client, grerrors.Error) {

	vErr := CreateDirectory(rootPath)
	if vErr != nil {
		return nil, vErr
	}
	return &defaultClient{
		rootPath: rootPath,
	}, nil
}
