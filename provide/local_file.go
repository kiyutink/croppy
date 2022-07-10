package provide

import (
	"errors"
	"image"
	"os"
	"strings"
)

// LocalFile reads the image from path src and returns an image.Image
func LocalFile(src string) (image.Image, error) {
	srcChunks := strings.Split(src, ".")
	ext := srcChunks[len(srcChunks)-1]
	if ext != "jpeg" && ext != "jpg" {
		return nil, errors.New("only .jpg and .jpeg files are supported")
	}
	file, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}
