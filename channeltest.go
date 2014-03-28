package main

import (
	"fmt"
	"runtime"
)

func main() {
	//a := []int{1, 2, 3, 4, 5, 6, 7}
	c := make(chan int, 2)
	c <- 1
	c <- 2
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
