package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//a := []int{1, 2, 3, 4, 5, 6, 7}
	m := make(map[int]chan int)
	for i := 1; i < 1000; i++ {
		c := make(chan int, 1)
		m[i] = c
		m[i] <- i
		go func() {
			tick := time.NewTicker(time.Millisecond * 50)
			tickmin := time.NewTicker(time.Second)
			loop := true
			for loop {
				select {
				case cc, ok := <-c:
					fmt.Println(cc, i, ok)
					loop = ok
				case <-tick.C:
				case <-tickmin.C:
					loop = false
				}
			}
			tick.Stop()
			tickmin.Stop()
			fmt.Println("aaaaaaaaaaa", i, loop)
		}()
	}
	time.Sleep(time.Second * 10)
	for _, v := range m {
		close(v)
	}

	//c = nil
	for true {
		time.Sleep(time.Second)
	}
	//fmt.Println(<-d)
	return
	//c <- 1
	//c <- 2
	//go Sum(a[:len(a)/2], c)
	//go Sum(a[len(a)/2:], c)
	//x, y := <-c, <-c
	//fmt.Println(x, y, x+y)
	fmt.Println(runtime.NumCPU())
}
func Sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}
