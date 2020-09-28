package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type color struct {
	r int
	g int
	b int
}

const pictureWidth = 1000
const pictureHeight = 1000
const colorMaximum = 255
const pixelSize = 100

func main() {
	// pictureWidth, err := strconv.Atoi(os.Args[1])
	// check(err)
	// pictureHeight, err := strconv.Atoi(os.Args[2])
	// check(err)
	// colorMaximum, err := strconv.Atoi(os.Args[3])
	// check(err)
	// pixelSize, err := strconv.Atoi(os.Args[4])
	// check(err)

	ppm := generatePPM()

	f := prepFile(pictureWidth, pictureHeight, colorMaximum)
	for verticalIndex := 0; verticalIndex < pixelSize; verticalIndex++ {
		for pixel := 0; pixel < (pictureWidth*pictureHeight)/pixelSize; pixel++ {
			for horizontalIndex := 0; horizontalIndex < pixelSize; horizontalIndex++ {
				color := ppm[pixel][horizontalIndex][verticalIndex]
				f.WriteString(fmt.Sprintf("%d %d %d ", color.r, color.g, color.b))
			}
		}
	}
	f.Close()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func generateRandomColor(maximum int) color {
	rand.Seed(time.Now().UnixNano())
	return color{rand.Intn(maximum), rand.Intn(maximum), rand.Intn(maximum)}
}

func prepFile(pictureWidth int, pictureHeight int, colorMaximum int) *os.File {
	f, err := os.Create("test.ppm")
	check(err)

	f.WriteString(fmt.Sprintf("P3 %d %d %d ", pictureWidth, pictureHeight, colorMaximum))
	return f
}

func generatePPM() [pictureWidth * pictureHeight / pixelSize][pictureWidth][pictureHeight]color {
	var ppm [pictureWidth * pictureHeight / pixelSize][pictureWidth][pictureHeight]color
	for pixel := 0; pixel < (pictureWidth*pictureHeight)/pixelSize; pixel++ {
		color := generateRandomColor(colorMaximum)
		for horizontalIndex := 0; horizontalIndex < pixelSize; horizontalIndex++ {
			for verticalIndex := 0; verticalIndex < pixelSize; verticalIndex++ {
				ppm[pixel][horizontalIndex][verticalIndex] = color
			}
		}
	}
	return ppm
}
