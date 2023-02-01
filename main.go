package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func drawCirlce(img *image.RGBA, x, y, radIn, radOut int, col color.Color) {

	// Draw a circular path on the image (2*pi = 360 degrees)
	for i := float64(0); i < 2*math.Pi; i += 0.01 {

		// loop from inner radius to outer radius
		for j := radIn; j < radOut; j++ {

			//calculate x and y coordinates
			x1 := int(x + int(math.Cos(i)*float64(j))) // get the unit point and scale them by j for actual point
			y1 := int(y + int(math.Sin(i)*float64(j)))

			//set the point on image
			img.Set(x1, y1, col)
		}
	}
}

func main() {
	// Create a new 100x100 image
	img := image.NewRGBA(image.Rect(0, 0, 1000, 1000))

	// Draw a circular path on the image
	drawCirlce(img, 500, 500, 0, 500, color.RGBA{0, 255, 0, 255})

	// Save the image to a file
	f, _ := os.Create("image.png")
	defer f.Close()
	png.Encode(f, img)
}
