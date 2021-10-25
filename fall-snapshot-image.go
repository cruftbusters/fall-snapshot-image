package main

import (
	"image"
	"image/png"
	"os"
)

func main() {
	grid := image.NewRGBA(image.Rect(0, 0, 400, 600))
	png.Encode(os.Stdout, grid)
}
