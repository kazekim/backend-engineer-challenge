package befiles

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"github.com/rs/xid"
	"io"
	"mime/multipart"
	"os"
)

// StoreFileLocal store file from request to local temp path in config
func (c *defaultClient) StoreFileLocal(file *multipart.FileHeader) (*FileData, grerrors.Error) {

	if c.maxFileSize > 0 && file.Size > c.maxFileSize {
		return nil, grerrors.ErrFileExceedSizeLimit
	}
	tempRootDir := c.rootPath + "/" + xid.New().String()

	vErr := CreateDirectory(tempRootDir)
	if vErr != nil {
		return nil, vErr
	}

	// Source
	src, err := file.Open()
	if err != nil {
		return nil, grerrors.ErrStorageUploadFileFailed
	}
	defer src.Close()

	contentType := file.Header.Get("Content-Type")

	// Destination
	dstPath := tempRootDir + "/" + file.Filename
	dst, err := os.Create(dstPath)
	if err != nil {
		return nil, grerrors.ErrFileNotFound
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return nil, grerrors.ErrStorageUploadFileFailed
	}

	m := FileData{
		Name:        file.Filename,
		Path:        tempRootDir,
		FullPath:    dstPath,
		Size:        file.Size,
		ContentType: contentType,
	}

	return &m, nil
}
