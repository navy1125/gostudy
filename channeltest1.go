package main

import (
	"fmt"
	"runtime"
	"time"
)

var cmap map[int]chan int
var count int

func fibonacci(chanint chan int) {
	ticker_sec := time.NewTicker(time.Second)
	for {
		select {
		case <-chanint:
		case <-ticker_sec.C:
			fmt.Println(len(chanint), count)
		}
	}
}
func main() {
	runtime.GOMAXPROCS(4)
	chanint := make(chan int, 10240)
	go func() {
		for {
			chanint <- count
			count++
		}
	}()
	fibonacci(chanint)
}
