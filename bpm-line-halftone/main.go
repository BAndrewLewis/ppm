package main

import (
	"image"
	"log"
	"os"

	"image/color"
	"image/jpeg"
	_ "image/jpeg"
)

func main() {
	// open file
	reader, err := os.Open("test.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	originalImage, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := originalImage.Bounds()
	newImage := image.NewGray(bounds)
	mode := 0
	counter := 0
	limit := 0
	// 0 normal row
	// 1 white row
	// 2 normal row
	// 3 black row
	for y := 0; y < bounds.Max.Y; y++ {
		counter += 1
		if counter > limit {
			if mode == 0 || mode == 2 {
				limit = 1
			} else if mode == 1 {
				limit = 0
			} else if mode == 3 {
				limit = 0
			}
			counter = 0
		}

		if mode >= 3 {
			mode = 0
		} else {
			mode += 1
		}

		for x := 0; x < bounds.Max.X; x++ {
			c := originalImage.At(x, y)
			r, g, b, _ := c.RGBA()
			if mode == 3 {
				newImage.SetGray(x, y, color.Gray{255})
			} else if mode == 0 || mode == 2 {
				if (r + g + b) > 70000 {
					newImage.SetGray(x, y, color.Gray{255})
				}
			}
		}

	}

	f, err := os.Create("outimage.jpg")
	if err != nil {
		// Handle error
	}
	defer f.Close()

	// Specify the quality, between 0-100
	// Higher is better
	opt := jpeg.Options{
		Quality: 100,
	}
	err = jpeg.Encode(f, newImage, &opt)
	if err != nil {
		// Handle error
	}
}
