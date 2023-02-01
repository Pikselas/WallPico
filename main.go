package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"wallpico/wallpico"
)

func main() {
	// Create a new 1000x1000 image
	img := image.NewRGBA(image.Rect(0, 0, 1000, 1000))

	rect := image.Rect(0, 0, 1000, 500)
	draw.Draw(img, rect, &image.Uniform{color.RGBA{0, 0, 0, 255}}, image.Point{}, draw.Src)
	rect = image.Rect(0, 500, 1000, 1000)
	draw.Draw(img, rect, &image.Uniform{color.RGBA{255, 255, 255, 255}}, image.Point{}, draw.Src)

	c := wallpico.Circle{Image: img}
	c.UseEffect(wallpico.EffectDiam)
	c.Draw(500, 500, 0, 200, color.RGBA{0, 0, 0, 255})
	c.UseEffect(wallpico.EffectMongo)
	c.Draw(500, 500, 200, 400, color.RGBA{0, 0, 0, 255})

	f, _ := os.Create("image.png")
	defer f.Close()
	png.Encode(f, img)
}
