// You can edit this code!
// Click here and start typing.
package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	//"rand"
	//"C"
)

func main() {
	fmt.Println("USER", os.Getenv("PATH"))
	defer fmt.Println("Hello, 世界")
	fmt.Printf("text:%d,%d", 100, 100)
	a := 10
	var b = 10
	var c int = 10
	fmt.Println("aaaaaaaaa", a, b, c)
	var host = flag.String("host", "", "Server listen host, default 0.0.0.0")
	fmt.Println("flag test", host)
	fmt.Println("cpunum:",runtime.NumCPU())
	//fmt.Println("aaaaaaaaa",a,b,c,rand.Intn(2))
	//fmt.Println("C.random", int(C.random()))
}
