package befiles

import (
	"github.com/h2non/filetype"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"os"
	"strings"
)

//ReadFileInfo read file info of file in full path
func ReadFileInfo(fullPath string) (*FileData, error) {
	f, err := os.Open(fullPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Get the content
	ct, err := GetFileContentType(f)

	if !filetype.IsMIMESupported(ct) {
		return nil, grerrors.ErrUnsupportedFileType
	}

	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}

	basePath := strings.Replace(fullPath, fi.Name(), "", -1)

	return &FileData{
		Name:        fi.Name(),
		Path:        basePath,
		FullPath:    fullPath,
		Size:        fi.Size(),
		ContentType: ct,
	}, err
}
