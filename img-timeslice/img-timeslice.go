package main

import (
	"github.com/hawx/img/exif"
	"github.com/hawx/img/utils"
	"flag"
	"fmt"
	"image"
	"os"
	"log"

	_ "image/jpeg"
	_ "image/gif"
	_ "image/png"
)

func run(paths []string) {
	if len(paths) < 2 {
		log.Fatal("Require 2 or more images")
	}

	var data *exif.Exif
	images := make([]image.Image, len(paths))

	for i, path := range paths {
		// Open
		file, err := os.Open(path)
		if err != nil {
			log.Fatal(path, "\n", err)
		}
		defer file.Close()

		// Decode
		m, _, err := image.Decode(file)
		if err != nil {
			log.Fatal(path, "\n", err)
		}

		if i == 0 {
			file.Seek(0, 0)
			data = exif.Decode(file)
		}

		images[i] = m
	}

	// Assume all photos are the same size as the first (stupid!)
	bounds := images[0].Bounds()

	// This is the width of a column. It mean that some of the right of an image
	// may be sacrificed. But it does mean all columns are exact pixels, and are
	// all exactly the same size.
	width := bounds.Dx() / len(images)

	out := image.NewRGBA(image.Rect(
		bounds.Min.X, bounds.Min.Y, bounds.Min.X + (len(images) * width), bounds.Max.Y,
	))

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Min.X + (len(images) * width); x++ {
			col := x / width
			out.Set(x, y, images[col].At(x, y))
		}
	}

	utils.WriteStdout(out, data)
}

func main() {
	var (
		long  = flag.Bool("long",  false, "")
		short = flag.Bool("short", false, "")
		usage = flag.Bool("usage", false, "")
	)

	os.Args = utils.GetOutput(os.Args)
	flag.Parse()

	if *long {
		fmt.Println(
			"  Take equal sized slices from a list of images, and glue them together.\n" +
			"  This can be used to create a not-a-timelapse type effect when used with\n" +
			"  photos of the same subject, taken over the course of a month (for instance).\n" +
			"  \n" +
			"  Unlike other img tools, this takes a list of filenames, for instance,\n" +
			"  \n" +
			"    img timelapse Trees/photo-*.png > output.png",
		)

	} else if *short {
		fmt.Println("glue slices of images together")

	} else if *usage {
		fmt.Println("timeslice <images...>")

	} else {
		run(flag.Args())
	}
}
