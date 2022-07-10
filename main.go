package main

import (
	"flag"
	"fmt"
	"image"
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
	out := flag.String("out", "./cropped.jpeg", "output path")
	file := flag.String("file", "", "source path of an image")
	url := flag.String("url", "", "remote url of an image")

	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	var img image.Image
	var err error

	switch {
	case *file != "" && *url != "":
		fmt.Fprintln(os.Stderr, "only on of --file and --url can be provided")
		return
	case *file != "":
		img, err = provide.LocalFile(*file)
	case *url != "":
		img, err = provide.RemoteUrl(*url)
	default:
		fmt.Fprintln(os.Stderr, "either --file or --url has to be provided!")
		return
	}

	if err != nil {
		panic(err)
	}
	croppedImg := crop.Rectangle(img, crop.NewBoundingRect(*x, *y, *x+*w, *y+*h))
	save.LocalFile(croppedImg, *out)

	fmt.Println(*out)
}
