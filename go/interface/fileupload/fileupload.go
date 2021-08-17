package fileupload

import (
	"image"
	"image/jpeg"
	"io"
	"mime/multipart"
	"os"
)

type FileUpload interface {
	UploadFile(file *multipart.FileHeader) error
	Crop(file *multipart.FileHeader, crop map[string]int) error
}

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

type fileUpload struct{}

func NewFileUpload() FileUpload {
	return &fileUpload{}
}

func (nfu *fileUpload) UploadFile(file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}

func (ntf *fileUpload) Crop(file *multipart.FileHeader, crop map[string]int) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	img, _, err := image.Decode(src)
	if err != nil {
		return err
	}

	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	cimg := img.(SubImager).SubImage(image.Rect(crop["x"], crop["y"], crop["width"], crop["height"]))

	jpeg.Encode(dst, cimg, &jpeg.Options{Quality: 80})

	return nil
}
