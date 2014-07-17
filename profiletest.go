package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	f, err := os.Create(*cpuprofile)
	if err != nil {
		log.Fatal(err)
	}
	runtime.Breakpoint()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	Print("aaaaaaaaaaaaaaaaaaa")
}

func Print(str string) {
	for i := 0; i < 1000; i++ {
		log.Println(str)
	}
}
