package provide

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

func SaveToLocalFile(img image.Image, path string, as string) (string, error) {
	filePath := path + "." + as
	newFile, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer newFile.Close()
	switch as {
	case "jpeg", "jpg":
		err = jpeg.Encode(newFile, img, nil)
	case "png":
		err = png.Encode(newFile, img)
	default:
		return "", errors.New("only jpeg/jpg or png formats are supported")
	}

	return filePath, err
}
