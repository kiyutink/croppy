package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kiyutink/croppy/crop"
	"github.com/kiyutink/croppy/provide"
	"github.com/kiyutink/croppy/save"
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

	img, err := provide.LocalFile(*s)
	if err != nil {
		panic(err)
	}
	croppedImg := crop.Rectangle(img, crop.NewBoundingRect(*x, *y, *x+*w, *y+*h))
	save.LocalFile(croppedImg, *o)

	fmt.Println(*o)
}
