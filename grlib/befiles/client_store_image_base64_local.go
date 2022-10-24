package befiles

import (
	"bytes"
	"encoding/base64"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"github.com/rs/xid"
	"golang.org/x/image/webp"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"
)

// StoreImageBase64Local store base64 image to local temp path defined in config
func (c *defaultClient) StoreImageBase64Local(data string) (*FileData, grerrors.Error) {

	b64data := data[strings.IndexByte(data, ',')+1:]

	it := PNG
	if strings.Contains(data, "image/jpeg") {
		it = JPEG
	} else if strings.Contains(data, "image/webp") {
		it = WEBP
	}

	unbased, err := base64.StdEncoding.DecodeString(b64data)
	if err != nil {
		log.Println("store image: ", err.Error())
		return nil, grerrors.ErrStorageUploadFileFailed
	}

	r := bytes.NewReader(unbased)

	var im image.Image
	switch it {
	case JPEG:
		im, err = jpeg.Decode(r)
	case PNG:
		im, err = png.Decode(r)
	case WEBP:
		im, err = webp.Decode(r)
	}
	if err != nil {
		log.Println("store image: ", err.Error())
		return nil, grerrors.ErrStorageUploadFileFailed
	}

	tempRootDir := c.rootPath + "/" + xid.New().String()

	vErr := CreateDirectory(tempRootDir)
	if vErr != nil {
		log.Println("store image: ", err.Error())
		return nil, vErr
	}

	fileName := "capture.jpeg"

	dstPath := tempRootDir + "/" + fileName

	f, err := os.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Println("store image: ", err.Error())
		return nil, grerrors.ErrStorageUploadFileFailed
	}

	err = jpeg.Encode(f, im, &jpeg.Options{
		Quality: 100,
	})

	if err != nil {
		log.Println("store image: ", err.Error())
		return nil, grerrors.ErrStorageUploadFileFailed
	}

	fi, err := f.Stat()
	if err != nil {
		log.Println("store image: ", err.Error())
		return nil, grerrors.ErrStorageUploadFileFailed
	}

	contentType := "image/jpeg"

	m := FileData{
		Name:        fileName,
		Path:        tempRootDir,
		FullPath:    dstPath,
		Size:        fi.Size(),
		ContentType: contentType,
	}

	return &m, nil
}
