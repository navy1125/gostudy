package main

import (
	"os"

	ipingo "github.com/bslizon/ipin-go"
)

func main() {
	b, err := ipingo.GetNormalizedPNG(os.Args[1])
	if err != nil {
		panic(err)
	}

	f, err := os.Create(os.Args[2])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err := f.Write(b); err != nil {
		panic(err)
	}
}
