// http://blog.animalswithinanimals.com/2008/08/databending-and-glitch-art-primer-part.html
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"hawx.me/code/img/utils"
)

const DEFAULT_SKIP = 4

func run(skip int) {
	inData, _ := ioutil.ReadAll(os.Stdin)
	outData := []byte{}

	for i, b := range inData {
		if i > skip {
			// replace BEEP (0x07) with SPACE (0x20)
			if b == 0x07 {
				outData = append(outData, 0x20)

				// replace VERTICAL TAB (0x0B), NEWLINE (0x0A) and RETURN (0x0D) with
				// NEWLINE RETURN (0x0A 0x0D)
			} else if b == 0x0B || b == 0x0A || b == 0x0D {
				outData = append(outData, 0x0A, 0x0D)

			} else {
				outData = append(outData, b)
			}
		} else {
			outData = append(outData, b)
		}
	}

	os.Stdout.Write(outData)
}

func main() {
	var (
		long  = flag.Bool("long", false, "")
		short = flag.Bool("short", false, "")
		usage = flag.Bool("usage", false, "")

		skip = flag.Int("skip", DEFAULT_SKIP, "")
	)

	os.Args = utils.GetOutput(os.Args)
	flag.Parse()

	if *long {
		fmt.Println(
			"  Databends the image by applying 'The Wordpad Effect' to it. This should\n" +
				"  be used on BMP or TIFF files, it may not play well with images that use\n" +
				"  some form of compression. (If you do want to go down that rabbit hole play\n" +
				"  with the --skip parameter, it will generally need to be larger).\n" +
				"  \n" +
				"    --skip <n>      # Bytes to skip (default: 4)",
		)

	} else if *short {
		fmt.Println("applies 'the wordpad effect' to an image")

	} else if *usage {
		fmt.Println("databend [options]")

	} else {
		run(*skip)
	}
}
