package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"
)

func main() {
	imagePath := os.Args[1]
	img, err := loadImage(imagePath)
	logFatalIfError(err)

	density := []string{" ", ".", "=", "!", "░", "▒", "▓", "█"}
	lengthDensity := len(density)
	maxDensityIndex := lengthDensity - 1

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			densityValue := getDensityValue(int(c.Y), lengthDensity)
			if densityValue < 0 {
				densityValue = 0
			}
			if densityValue > maxDensityIndex {
				densityValue = maxDensityIndex
			}
			densityString := " " + string(density[densityValue])
			fmt.Print(densityString)
		}
		fmt.Print("\n")
	}
}

func loadImage(imagePath string) (image.Image, error) {
	var img image.Image
	var err error
	sl := strings.Split(imagePath, ".")
	ext := sl[len(sl)-1]

	image, err := os.Open(imagePath)
	logFatalIfError(err)
	defer image.Close()

	switch strings.ToLower(ext) {
	case "png":
		{
			img, err = png.Decode(image)
			break
		}

	case "jpg", "jpeg":
		{
			img, err = jpeg.Decode(image)
			break
		}
	default:
		err = errors.New("format foto tidak didukung")
	}
	return img, err
}

func getDensityValue(grayScaleValue int, lengthDensity int) int {
	const maxColorValue = 255
	return grayScaleValue/(maxColorValue/lengthDensity) - 1
}

func logFatalIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
