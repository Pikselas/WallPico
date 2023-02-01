package wallpico

import (
	"image"
	"image/color"
	"math"
)

type CircularEffect func(float64, int) color.Color

func EffectSolid(cPoint float64, rPoint int) color.Color {
	return color.RGBA{255, 0, 0, 255}
}

func EffectDiam(cPoint float64, rPoint int) color.Color {
	return color.RGBA{uint8((rPoint % 255) * 5), 0, 0, 255}
}

func EffectMongo(cPoint float64, rPoint int) color.Color {
	return color.RGBA{uint8(255 * cPoint / 2 * math.Pi), 0, 0, 255}
}

type Circle struct {
	Image  *image.RGBA
	Effect CircularEffect
}

func (c *Circle) UseEffect(f CircularEffect) {
	c.Effect = f
}

func (c *Circle) Draw(x, y, radIn, radOut int, col color.Color) {
	// Draw a circular path on the image (2*pi = 360 degrees)
	for i := float64(0); i < 2*math.Pi; i += 0.001 {

		// loop from inner radius to outer radius
		for j := radIn; j < radOut; j++ {

			col := c.Effect(i, j)

			//calculate x and y coordinates
			x1 := int(x + int(math.Cos(i)*float64(j))) // get the unit point and scale them by j for actual point
			y1 := int(y + int(math.Sin(i)*float64(j)))

			//fmt.Println(col, c)

			//set the point on image
			c.Image.Set(x1, y1, col)
		}
	}
}
