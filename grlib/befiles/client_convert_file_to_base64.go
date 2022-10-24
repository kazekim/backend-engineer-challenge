package befiles

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"image"
	"image/png"
	"log"
	"os"
)

// ConvertFileToBase64 convert local file to base64 string
func (c *defaultClient) ConvertFileToBase64(file FileData) (*string, grerrors.Error) {

	f, err := os.Open(file.FullPath)
	if err != nil {
		log.Println("store image: ", err.Error())
		return nil, grerrors.ErrStorageUploadFileFailed
	}
	defer f.Close()

	r := bufio.NewReader(f)

	img, _, err := image.Decode(r)
	if err != nil {
		return nil, grerrors.NewDefaultError(err)
	}
	//m := resize.Resize(125, 125, img, resize.Lanczos3)
	buf := new(bytes.Buffer)
	err = png.Encode(buf, img)
	imageBit := buf.Bytes()

	fileBase64 := base64.StdEncoding.EncodeToString([]byte(imageBit))
	return &fileBase64, nil
}
