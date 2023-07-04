package main

import (
	"fmt"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ifc <inputfile> <outputfile>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	img, err := imaging.Open(inputFile)
	if err != nil {
		fmt.Printf("failed to open image: %v", err)
		os.Exit(1)
	}

	outputFileExtension := strings.ToLower(strings.TrimPrefix(outputFile, strings.TrimSuffix(outputFile, filepath.Ext(outputFile))))

	switch outputFileExtension {
	case ".jpg", ".jpeg":
		err = imaging.Save(img, outputFile, imaging.JPEGQuality(95))
	case ".png":
		err = imaging.Save(img, outputFile, imaging.PNGCompressionLevel(png.BestCompression))
	default:
		fmt.Printf("Unsupported output format: %s\n", outputFileExtension)
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("failed to save image: %v", err)
		os.Exit(1)
	}

	fmt.Println("Image converted successfully!")
}
