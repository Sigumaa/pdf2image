package main

import (
	"fmt"
	"image/jpeg"
	"os"
	"path/filepath"

	"github.com/gen2brain/go-fitz"
)

func main() {
	if len(os.Args) != 3 {
		panic("The number of arguments does not match.")
	}

	pdfName := os.Args[1]
	imgName := os.Args[2]

	doc, err := fitz.New(fmt.Sprintf("%s.pdf", pdfName))
	if err != nil {
		panic(err)
	}

	for n := 0; n < doc.NumPage(); n++ {
		img, err := doc.Image(n)
		if err != nil {
			panic(err)
		}

		f, err := os.Create(filepath.Join(fmt.Sprintf("%s%03d.jpg", imgName, n)))
		if err != nil {
			panic(err)
		}

		err = jpeg.Encode(f, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
		if err != nil {
			panic(err)
		}

		f.Close()
	}
}
