package main

import (
	"image"
	"log"
	"os"
)

func main() {
	file, err := os.Open("E:\\tmp\\7天开服活动_1.04.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, str, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(str)
}
