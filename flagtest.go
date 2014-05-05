package main

import (
	"flag"
	"fmt"
	"os"
)

type Test struct{
	Items []struct{
		Id uint32
		Num uint32
		Str string
	}
}

func main() {
	var port int
	flag.IntVar(&port, "port", 80, "usage")
	host := flag.String("host", "localhost", "addr")
	needTest := flag.Bool("test", true, "usage")
	fmt.Println("args num:", flag.NArg())
	fmt.Println("host:", flag.Arg(1))
	fmt.Println("host:", flag.Args())
	flag.PrintDefaults()
	flag.Parse()
	fmt.Println("needTest:", *needTest)
	fmt.Println("host:", *host)
	fmt.Println("port:", port)
	fmt.Println("args:", os.Args[1])
	flag.Visit(Visit)
	f := flag.Lookup("port")
	fmt.Println(f.Name, f.DefValue)
	flag.Usage()
}
func Visit(f *flag.Flag) {
	fmt.Println(f.Name, f.Usage, f.DefValue, f.Value.String())
}
