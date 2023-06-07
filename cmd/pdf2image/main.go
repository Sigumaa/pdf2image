package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Sigumaa/pdf2image"
)

func main() {
	pdfName := flag.String("pdf", "", "PDF file name")
	imgName := flag.String("img", "", "Image file name")
	flag.Parse()

	if *pdfName == "" || *imgName == "" {
		flag.Usage()
		log.Fatal("PDF file name or image file name is missing.")
	}

	pdf2image.PdfToImage(*pdfName, *imgName)

	fmt.Println("Success!")
}
