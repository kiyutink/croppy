package crop

import (
	"image"
	"image/draw"
)

func NewBoundingRect(x, y, w, h int) image.Rectangle {
	rect := image.Rectangle{
		Min: image.Point{
			X: x,
			Y: y,
		},
		Max: image.Point{
			X: x + w,
			Y: y + h,
		},
	}
	return rect
}

// Rectangle takes a source image, and an image.Rectangle and returns a cropped image
func Rectangle(img image.Image, rect image.Rectangle) image.Image {
	newImg := image.NewRGBA(rect)
	draw.Draw(newImg, rect, img, image.Pt(0, 0), draw.Src)
	return newImg
}
