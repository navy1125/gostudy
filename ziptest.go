package main

import (
	"archive/zip"
	"fmt"
	"log"
)

func main() {
	file, err := zip.OpenReader("e:\\tmp\\LoginServer.zip")
	if err != nil {
		log.Fatal(err)
	}
	for i,f := range file.File {
		fmt.Println(i,f.Name)
	}
}
