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
	// Check that the user has provided the correct number of command line arguments
	if len(os.Args) != 3 {
		fmt.Println("Usage: ifc <inputfile> <outputfile>")
		os.Exit(1) // Exit the program if the number of arguments is incorrect
	}

	// Extract the input and output file paths from the command line arguments
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Open the input image file
	img, err := imaging.Open(inputFile)
	if err != nil {
		fmt.Printf("failed to open image: %v", err)
		os.Exit(1) // Exit the program if the image file cannot be opened
	}

	// Determine the output file extension (i.e., the desired output format)
	outputFileExtension := strings.ToLower(strings.TrimPrefix(outputFile, strings.TrimSuffix(outputFile, filepath.Ext(outputFile))))

	// Based on the desired output format, save the image in the correct format
	switch outputFileExtension {
	case ".jpg", ".jpeg":
		err = imaging.Save(img, outputFile, imaging.JPEGQuality(95)) // Save as JPEG with quality level 95
	case ".png":
		err = imaging.Save(img, outputFile, imaging.PNGCompressionLevel(png.BestCompression)) // Save as PNG with the highest compression
	default:
		fmt.Printf("Unsupported output format: %s\n", outputFileExtension)
		os.Exit(1) // Exit the program if the output format is not supported
	}

	// If there was an error saving the image, print an error message and exit the program
	if err != nil {
		fmt.Printf("failed to save image: %v", err)
		os.Exit(1)
	}

	// If the image was saved successfully, print a success message
	fmt.Println("Image converted successfully!")
}
