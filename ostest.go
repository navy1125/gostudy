// You can edit this code!
// Click here and start typing.
package main

import (
	/*
		"flag"
		"fmt"
		"git.code4.in/logging"
		iconv "github.com/hwch/iconv"
		"go/build"
		"os"
		"path"
		"runtime"
		"strconv"
		//"rand"
		//"C"
		"log"
		"time"
		// */
	"fmt"
	"os"
)

var (
	TestMap map[string]string
)

func TestPanic() {
	panic("ddddddddddddd")
}

func main() {
	fmt.Println(os.Stat("xxdf"))
	fmt.Println(os.Stat("ostest.go"))
	/*
		ch := make(chan int, 2)
		ch <- 1
		ch <- 2
		for c := range ch {
			switch c {
			case 1:
				break
			}
			fmt.Println(c)
		}
		//for v := range ch {
		//	fmt.Println("cccccccccc:%v", v)
		//}
		fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()), runtime.NumCPU(), runtime.GOROOT())
		fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()), runtime.NumCPU())
		fmt.Println(len(ch), cap(ch))
		time.Sleep(time.Second)
		ch <- 2
		os.Exit(0)
		ch1 := ch
		if ch1 == ch {
			fmt.Println("aaaaaaaaaaaaaa%v,%v", ch1, ch)
		}
		ch = make(chan int, 3)
		for b1 := range ch1 {

			logging.Debug("%v,%v", b1)
		}
		fmt.Println(len(ch))
		os.Exit(0)
		logging.Debug(path.Dir("c/b/b"))
		log.Println("GOOS:", runtime.GOARCH, runtime.GOOS, build.Default)
		os.Exit(0)
		defer func() {
			if x := recover(); x != nil {
				fmt.Println(x)
			}
		}()
		stra := "a"
		str := fmt.Sprintf("%03o,%d", stra[0], stra[0])
		aa, _ := strconv.ParseInt("141", 8, 32)
		fmt.Println("str:", str, aa, fmt.Sprintf("%c", aa))
		converter, err := iconv.NewCoder(iconv.GBK2312_UTF8_IDX)
		if err != nil {
			fmt.Println("iconv err:", err)
			return
		}
		out := make([]byte, 1024)
		if l, err := converter.CodeConvertFunc([]byte("王海军"), out); err == nil && l > 0 {
			fmt.Println("converter err:", err)
		}
		fmt.Println(string(out))
		return
		fmt.Println("USER", os.Getenv("PATH"))
		defer fmt.Println("Hello, 世界")
		fmt.Printf("text:%d,%d", 100, 100)
		a := 10
		var b = 10
		var c int = 10
		fmt.Println("aaaaaaaaa", a, b, c)
		var host = flag.String("host", "", "Server listen host, default 0.0.0.0")
		fmt.Println("flag test", host)
		fmt.Println("cpunum:", runtime.NumCPU())
		//fmt.Println("aaaaaaaaa",a,b,c,rand.Intn(2))
		//fmt.Println("C.random", int(C.random()))
		TestMap = make(map[string]string)
		TestMap["whj"] = "wanghaijun@ztgame.com"
		fmt.Println(TestMap)
		fmt.Println(TestMap["whj"])
		TestPanic()
		//*/
}
