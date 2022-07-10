package provide

import (
	"image"
	"image/jpeg"
	"net/http"
)

func RemoteUrl(url string) (image.Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	img, err := jpeg.Decode(resp.Body)
	if err != nil {
		return nil, err
	}
	return img, nil
}
