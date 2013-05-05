package main

import (
	"github.com/hawx/img/blend"
	"github.com/hawx/img/crop"
	"github.com/hawx/img/utils"
	"flag"
	"fmt"
	"image"
	"os"
)

const DEFAULT_SCALE = 0.66

func flip(img image.Image) image.Image {
	b := img.Bounds()
	o := image.NewRGBA(b)

	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			yinv := b.Max.Y - (y - b.Min.Y)

			o.Set(x, yinv, img.At(x, y))
		}
	}

	return o
}

func min(x,y int) int {
	if x > y { return y }
	return x
}

type Cropper func(image.Image, int, utils.Direction) image.Image

func run(scale float64, f Cropper) {
	img, data := utils.ReadStdin()

	b := img.Bounds()
	s := int(float64(min(b.Dx(), b.Dy())) * scale)

	inner := f(img, s, utils.Centre)

	img = blend.Normal(img, flip(inner))

	utils.WriteStdout(img, data)
}

func main() {
	var (
		long  = flag.Bool("long",  false, "")
		short = flag.Bool("short", false, "")
		usage = flag.Bool("usage", false, "")

		scale = flag.Float64("scale", DEFAULT_SCALE, "")

		triangle = flag.Bool("triangle", false, "")
		square   = flag.Bool("square",   false, "")
	)

	os.Args = utils.GetOutput(os.Args)
	flag.Parse()

	f := crop.Circle
	if *triangle {
		f = crop.Triangle
	} else if *square {
		f = crop.Square
	}

	if *long {
		fmt.Println(
			"  Flips a central circle of the image (size of which can be controlled\n" +
			"  with the --scale flag), and overlays it on the original image.\n" +
			"  \n" +
			"    --scale <n>     # Scale factor for central circle (default: 0.66)\n" +
		  "    --triangle      # Use a triangle instead\n" +
			"    --square        # Use a square instead",
		)

	} else if *short {
		fmt.Println("spin a circle")

	} else if *usage {
		fmt.Println("wlsn [options]")

	} else {
		run(*scale, f)
	}
}
