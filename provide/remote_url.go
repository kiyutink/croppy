package provide

import (
	"image"
	"net/http"
)

func RemoteUrl(url string) (image.Image, string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()
	return image.Decode(resp.Body)
}
