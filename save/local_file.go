package save

import (
	"image"
	"image/jpeg"
	"os"
)

func LocalFile(img image.Image, path string) error {
	newFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer newFile.Close()
	err = jpeg.Encode(newFile, img, nil)

	return err
}
