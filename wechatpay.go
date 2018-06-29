package main

import (
	"image"
	"image/color"
	"io/ioutil"
	"log"
	"os"
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
	indir := "./"   //输入
	outdir := indir //输出
	comdir := indir //公共资源目录
	if len(os.Args) == 2 {
		indir = os.Args[1]
		outdir = indir
		comdir = indir
	} else if len(os.Args) == 3 {
		indir = os.Args[1]
		outdir = os.Args[2]
		comdir = indir
	} else if len(os.Args) == 4 {
		indir = os.Args[1]
		outdir = os.Args[2]
		comdir = os.Args[3]
	}
	top, err := imaging.Open(comdir + "/" + "back.png")
	if err != nil {
		log.Fatalf("failed to open image:", err)
	}
	center, _ := imaging.Open(comdir + "/" + "center.jpg")
	bottom, _ := imaging.Open(comdir + "/" + "bottom.jpg")
	files, _ := ioutil.ReadDir(indir)
	pt := image.Pt(0, 0)
	for _, file := range files {
		if file.IsDir() {
			log.Println("ignore dir:", file.Name())
			continue
		}
		if filepath.Ext(file.Name()) != ".jpg" {
			log.Println("ignore file:", file.Name())
			continue
		}
		path := filepath.Join(indir, file.Name())

		src, err := imaging.Open(path)
		if err != nil {
			log.Println("failed to open image:", err, file.Name())
			continue
		}
		if src.Bounds().Dx() != top.Bounds().Dx() || src.Bounds().Dy() != top.Bounds().Dy() {
			log.Println("siz err image:", file.Name(), src.Bounds().Dx(), top.Bounds().Dx(), src.Bounds().Dy(), top.Bounds().Dy())
			continue
		}

		dst := imaging.New(top.Bounds().Dx(), top.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
		dst = imaging.Paste(dst, src, pt)
		if top != nil {
			log.Println("top ok")
			dst = imaging.Overlay(dst, top, pt, 1)
		}
		if center != nil {
			dst = imaging.Overlay(dst, center, pt, 1)
		}
		if bottom != nil {
			dst = imaging.Overlay(dst, bottom, pt, 1)
		}
		err = imaging.Save(dst, filepath.Join(outdir, file.Name()), imaging.JPEGQuality(50))
		if err != nil {
			log.Fatalf("failed to save image: %v", err)
		}
		log.Println("out file:", filepath.Join(outdir, file.Name()))
	}
	/*
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
						src, err = imaging.Open("ludan/fenshupai_dian_02.png")
						if err != nil {
							log.Fatalf("failed to save image: %v", err)
						}
						dst = imaging.Paste(dst, src, pt)
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
		//*/
}
