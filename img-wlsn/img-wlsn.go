package main

import (
	"github.com/hawx/img/blend"
	"github.com/hawx/img/crop"
	"github.com/hawx/img/utils"
	"flag"
	"fmt"
	"image"
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

func run(scale float64) {
	img := utils.ReadStdin()

	b := img.Bounds()
	s := int(float64(min(b.Dx(), b.Dy())) * scale)

	crc := crop.Circle(img, s, utils.Centre)
	img  = blend.Normal(img, flip(crc))

	utils.WriteStdout(img)
}

func main() {
	var (
		long  = flag.Bool("long",  false, "")
		short = flag.Bool("short", false, "")
		usage = flag.Bool("usage", false, "")

		scale = flag.Float64("scale", DEFAULT_SCALE, "")
	)

	flag.Parse()

	if *long {
		fmt.Println(
			"  Flips a central circle of the image (size of which can be controlled\n" +
			"  with the --scale flag), and overlays it on the original image.\n" +
			"  \n" +
			"    --scale <n>     # Scale factor for central circle (default: 0.66)",
		)

	} else if *short {
		fmt.Println("spin a circle")

	} else if *usage {
		fmt.Println("wlsn [options]")

	} else {
		run(*scale)
	}
}
