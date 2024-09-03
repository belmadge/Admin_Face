package utils

import (
	"image"
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

// ReduceImageSize reduces the size of an image
func ReduceImageSize(filePath string, width uint) (string, error) {
	// Open the image file
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		return "", err
	}

	// Resize the image
	m := resize.Resize(width, 0, img, resize.Lanczos3)

	// Save the resized image
	out, err := os.Create("resized_" + filePath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	jpeg.Encode(out, m, nil)

	return out.Name(), nil
}
