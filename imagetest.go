package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func main() {
	//image.RegisterFormat(name, magic, decode, decodeConfig)
	file, err := os.Open("E:\\psds18291.jpg")
	if err != nil {
		log.Fatal("open file err:", err)
	}
	defer file.Close()
	img, str, err := image.Decode(file)
	if err != nil {
		log.Fatal("decode image err:", err)
	}
	log.Println(img.Bounds(), str)
	outfilepng, err := os.OpenFile("c:\\test1.jpg", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	outfilejpg, err := os.OpenFile("c:\\test.jpg", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	defer outfilejpg.Close()
	defer outfilepng.Close()
	rect := img.Bounds()
	rect = rect.Inset(10)
	m0 := image.NewRGBA(image.Rect(0, 0, 80, 50))
	for i := 0; i < m0.Bounds().Dx(); i++ {
		for j := 0; j < m0.Bounds().Dy(); j++ {
			m0.Set(i, j, color.RGBA{255, 0, 0, 255})
		}
	}
	m1, _ := m0.SubImage(image.Rect(1, 2, 5, 5)).(*image.RGBA)
	fmt.Println(m0.Bounds().Dx(), m1.Bounds().Dx())
	fmt.Println(m0.Stride == m1.Stride)
	m2 := m0.SubImage(image.Rect(0, 0, m0.Bounds().Dx(), m0.Bounds().Dy()))
	fmt.Println(m2.Bounds().Dx(), m2.Bounds().Dy())
	opt := jpeg.Options{Quality: 10}
	jpeg.Encode(outfilejpg, img, &opt)
	//png.Encode(outfilepng, img)
	//png.Encode(outfilepng, m0)
	uniform := image.NewUniform(color.RGBA{255, 255, 0, 0})
	//png.Encode(outfilepng, uniform)
	opt.Quality = 100
	jpeg.Encode(outfilepng, uniform, &opt)
}
