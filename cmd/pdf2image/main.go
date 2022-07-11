package main

import (
	"os"

	"github.com/Sigumaa/pdf2image"
)

func main() {
	if len(os.Args) != 3 {
		panic("The number of arguments does not match.")
	}

	pdfName := os.Args[1]
	imgName := os.Args[2]

	pdf2image.PdfToImage(pdfName, imgName)
}
