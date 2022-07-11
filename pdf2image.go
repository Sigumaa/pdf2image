package pdf2image

import (
	"fmt"
	"image/jpeg"
	"os"
	"path/filepath"

	"github.com/gen2brain/go-fitz"
)

func PdfToImage(pdfName string, imgName string) {

	doc, err := fitz.New(fmt.Sprintf("%s.pdf", pdfName))
	if err != nil {
		panic(err)
	}

	for n := 0; n < doc.NumPage(); n++ {
		img, err := doc.Image(n)
		if err != nil {
			panic(err)
		}

		f, err := os.Create(filepath.Join(fmt.Sprintf("%s.jpg", imgName)))
		if err != nil {
			panic(err)
		}

		defer f.Close()

		err = jpeg.Encode(f, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
		if err != nil {
			panic(err)
		}
	}
}
