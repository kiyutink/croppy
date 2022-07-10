package provide

import (
	"image"
	"os"
)

// LocalFile reads the image from path src and returns an image.Image
func LocalFile(src string) (image.Image, string, error) {
	file, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	return image.Decode(file)
}
