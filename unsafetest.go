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
