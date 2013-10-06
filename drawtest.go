package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
)

func main() {
	f1, err := os.Open("c:\\1.jpg")
	if err != nil {
		panic(err)
	}
	defer f1.Close()

	f2, err := os.Open("c:\\2.jpg")
	if err != nil {
		panic(err)
	}
	defer f2.Close()

	f3, err := os.Create("c:\\3.jpg")
	if err != nil {
		panic(err)
	}
	defer f3.Close()

	m1, err := jpeg.Decode(f1)
	if err != nil {
		panic(err)
	}
	bounds := m1.Bounds()

	m2, err := jpeg.Decode(f2)
	if err != nil {
		panic(err)
	}

	m := image.NewRGBA(bounds)
	white := color.RGBA{255, 255, 255, 255}
	draw.Draw(m, bounds, &image.Uniform{white}, image.ZP, draw.Src)
	draw.Draw(m, bounds, m1, image.ZP, draw.Src)
	draw.Draw(m, image.Rect(0, 0, 800, 800), m2, image.Pt(0, 0), draw.Src)

	err = jpeg.Encode(f3, m, &jpeg.Options{90})
	if err != nil {
		panic(err)
	}

	fmt.Printf("ok\n")
}
