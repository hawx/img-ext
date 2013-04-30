package main

import (
	"code.google.com/p/graphics-go/graphics"
	"github.com/nfnt/resize"

	"github.com/hawx/img/blend"
	"github.com/hawx/img/channel"
	"github.com/hawx/img/contrast"
	"github.com/hawx/img/utils"

	"flag"
	"fmt"
	"image"
	"image/color"
	"os"
)

func maskFor(in image.Image) image.Image {
	b := in.Bounds()

	// thumb size, corresponds to first convert command
	thumbWidth  := int(b.Dx() / 5)
	thumbHeight := int(b.Dy() / 5)

	thumb := image.NewRGBA(image.Rect(0, 0, thumbWidth, thumbHeight))
	final := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))

	// fill black (xc:black)
	// draw white rectangle  (1,1 : thumbWidth-1,thumbHeight-1)
	for y := 0; y < thumbHeight; y++ {
		for x := 0; x < thumbWidth; x++ {
			if (y > 0 && y < thumbHeight-1) && (x > 0 && x < thumbWidth-1) {
				thumb.Set(x, y, color.White)
			} else {
				thumb.Set(x, y, color.Black)
			}
		}
	}

	// apply Gaussian blur with radius=7, sigma=15
	graphics.Blur(thumb, thumb, &graphics.BlurOptions{2, 7})

	// now resize to original image size
	resized := resize.Resize(uint(b.Dx()), uint(b.Dy()), thumb, resize.Bilinear)

	// with gaussian blur radius=0, sigma=5
	graphics.Blur(final, resized, &graphics.BlurOptions{5, 0})

	return final
}

func main() {
	var (
		long  = flag.Bool("long", false, "")
		short = flag.Bool("short", false, "")
		usage = flag.Bool("usage", false, "")
	)

	os.Args = utils.GetOutput(os.Args)
	flag.Parse()

	if *long {
		fmt.Println(
			"  Applies a simple lomo effect to the image, boosting its saturation and\n" +
			"  composing with a black edged mask.",
		)

	} else if *short {
		fmt.Println("applies a simple lomo effect to the image")

	} else if *usage {
		fmt.Println("lomo [options]")

	} else {
		img, data := utils.ReadStdin()

		// http://the.taoofmac.com/space/blog/2005/08/23/2359
		img = contrast.Adjust(img, 1.2)
		img = channel.Adjust(img, utils.Multiplier(1.2), channel.Saturation)
		img = blend.Multiply(img, maskFor(img))

		utils.WriteStdout(img, data)
	}
}
