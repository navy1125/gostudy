package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//a := []int{1, 2, 3, 4, 5, 6, 7}
	c := make(chan int, 2)
	d := c
	//d := make(chan int, 1)
	go func() {
		tick := time.Tick(time.Millisecond)
		loop := true
		for loop {
			select {
			case cc := <-c:
				fmt.Println(cc)
			case <-d:
				fmt.Println("dddddd")
				loop = false
			case <-tick:
				close(c)
				close(d)
			}
		}
		fmt.Println("aaaaaaaaaaa")
	}()

	//c = nil
	time.Sleep(time.Second)
	//fmt.Println(<-d)
	return
	//c <- 1
	//c <- 2
	//go Sum(a[:len(a)/2], c)
	//go Sum(a[len(a)/2:], c)
	//x, y := <-c, <-c
	//fmt.Println(x, y, x+y)
	fmt.Println(<-c)
	fmt.Println(<-c)
	close(c)
	fmt.Println(runtime.NumCPU())
}
func Sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}
