package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	text, err := ioutil.ReadFile("D:\\work\\gotest\\client.go")
	if err != nil {
		log.Fatal("ioutil.ReadFile err:", err)
	}
	fmt.Printf("%s", text)
}
