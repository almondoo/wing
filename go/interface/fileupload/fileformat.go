package fileupload

import (
	"path"

	"github.com/google/uuid"
)

func FormatFile(fn string) string {
	ext := path.Ext(fn)
	u := uuid.New()

	newFileName := u.String() + ext
	return newFileName
}
