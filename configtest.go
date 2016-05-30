package main

import (
	"flag"
	"fmt"
)

func main() {
	var gopherType string
	const (
		defaultGopher = "pocket"
		usage         = "the variety of gopher"
	)
	flag.StringVar(&gopherType, "gopher_type", defaultGopher, usage)
	flag.StringVar(&gopherType, "g", defaultGopher, usage+" (shorthand)")
	flag.Parse()
	fmt.Println(gopherType)
}
