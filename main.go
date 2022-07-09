package main

import (
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"os"
)

func main() {
	help := flag.Bool("help", false, "help")

	x := flag.Int("x", 0, "where to start cropping from, x coordinate")
	y := flag.Int("y", 0, "where to start cropping from, y coordinate")
	w := flag.Int("w", 100, "width of the cropped area in pixels, defaults to 100")
	h := flag.Int("h", 100, "height of the cropped area in pixels, defaults to 100")
	o := flag.String("o", "./cropped.jpeg", "output path")
	s := flag.String("s", "", "source path")

	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	if *s == "" {
		fmt.Fprintln(os.Stderr, "the source path (-s) is required!")
		return
	}

	file, err := os.Open(*s)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	croppedImg := image.NewRGBA(image.Rect(*x, *y, *x+*w, *y+*h))
	draw.Draw(croppedImg, croppedImg.Bounds(), img, image.Point{X: *x, Y: *y}, draw.Src)
	newFile, err := os.Create(*o)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	err = jpeg.Encode(newFile, croppedImg, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(*o)
}
