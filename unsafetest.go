package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
	"unsafe"
)

type Point struct {
	x int8
	y int16
}
type Rect struct {
	Point
	with   int16
	height int16
	data   [2]byte
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rect := Rect{}
	fmt.Println(reflect.TypeOf(rect).Size())
	fmt.Println(unsafe.Sizeof(rect))
	pt := (*Point)(unsafe.Pointer(&rect))
	fmt.Println(rect, pt)
	c := make(chan int, 2)//修改2为1就报错，修改2为3可以正常运行
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
	ch := make(chan []byte)
	go func() {
		buf := make([]byte, 1024*1000)
		ch <- buf
	}()
	go func() {
		buf := <-ch
		if len(buf) != 1024*1000 {
			fmt.Println("chan read err:", len(buf))

		}
	}()
	tick := time.Tick(time.Second)
	for {
		t := <-tick
		fmt.Println(t)
	}
}
