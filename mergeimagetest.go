package main

import (
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"log"
	"math/rand"
	"path/filepath"

	"github.com/disintegration/imaging"
)

var (
	Point2Png map[int]string
)

type Road struct {
	Point   []int //筛子或者,扑克牌的明细信息
	WinType int
}

type RoadLine struct {
	List []*Road //一列
}

type RoadMap struct {
	List []*RoadLine //筛子或者,扑克牌的明细信息
}

func main() {
	Point2Png = make(map[int]string)
	Point2Png[0] = "ludan/fenshupai_zhuang.png"
	Point2Png[1] = "ludan/fenshupai_xian.png"
	Point2Png[2] = "ludan/fenshupai_he.png"
	/*
		src, err := imaging.Open("testdata/flowers.png")
		if err != nil {
			log.Fatalf("failed to open image: %v", err)
		}

		// Crop the original image to 300x300px size using the center anchor.
		src = imaging.CropAnchor(src, 300, 300, imaging.Center)
		// */

	// Create a new image and paste the four produced images into it.
	dst := imaging.New(1400, 442, color.NRGBA{0, 0, 0, 0})
	//dst = imaging.Paste(dst, img1, image.Pt(0, 0))
	//dst = imaging.Paste(dst, img2, image.Pt(0, 200))
	//dst = imaging.Paste(dst, img3, image.Pt(200, 0))
	//dst = imaging.Paste(dst, img4, image.Pt(200, 200))

	// Save the resulting image as JPEG.
	dir := "ludan/"
	files, _ := ioutil.ReadDir(dir)
	pt := image.Pt(1, 1)
	for _, file := range files {
		path := filepath.Join(dir, file.Name())
		src, err := imaging.Open(path)
		if err != nil {
			log.Fatalf("failed to save image: %v", err)
		}
		dst = imaging.Paste(dst, src, pt)
		pt.X += src.Bounds().Dx() + 1
		fmt.Println(path, src.Bounds().Size())
		if pt.X+src.Bounds().Dx()+1 > dst.Rect.Dx() {
			pt.Y += src.Bounds().Dy() + 1
			pt.X = 0
		}
	}
	err := imaging.Save(dst, "ludan.jpg")
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
	{
		m := RoadMap{}
		for i := 0; i < 20; i++ {
			l := &RoadLine{}
			m.List = append(m.List, l)
			r := &Road{WinType: rand.Intn(2)}
			for j := 0; j < 10; j++ {
				l.List = append(l.List, r)
				if rand.Intn(3) == 1 {
					break
				}
			}
		}
		dst := imaging.New(1400, 442, color.NRGBA{0, 0, 0, 0})

		pt := image.Pt(1, 1)
		for _, l := range m.List {
			x := 0
			for _, r := range l.List {
				path, _ := Point2Png[r.WinType]
				src, err := imaging.Open(path)
				if err != nil {
					log.Fatalf("failed to save image: %v", err)
				}
				dst = imaging.Overlay(dst, src, pt, 1)
				/*
					src, err = imaging.Open("ludan/fenshupai_dian_02.png")
					if err != nil {
						log.Fatalf("failed to save image: %v", err)
					}
					dst = imaging.Paste(dst, src, pt)
					/*/
				pt.Y += src.Bounds().Dy() + 1 - 9
				x = src.Bounds().Dx()
			}
			pt.X += x + 1
			pt.Y = 1
		}
		err := imaging.Save(dst, "ludan1.jpg", imaging.JPEGQuality(100))
		if err != nil {
			log.Fatalf("failed to save image: %v", err)
		}
	}
}
