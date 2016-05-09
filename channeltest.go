package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

func CheckChan() {
	v := make(chan bool, 5)
	v2 := make(chan bool, 5)
	v <- true
	v <- true
	v <- true
	v <- true
	v <- true
	//v = append(v, false)
	if true {

		defer func() {
			if x := recover(); x != nil {
				fmt.Println(x)
				fmt.Println(reflect.TypeOf(x))
			}
		}()
	}

	l := len(v)
	fmt.Println("ccccccccccccc", l, cap(v))
	for i := 0; i < l; i++ {
		switch i {
		case 1:
			fmt.Println("xxxxxxx")
			break
		}
		fmt.Println("xxxxxxx")
		v2 <- <-v
		if r, ok := <-v2; ok == true {
			fmt.Println("ccccccccccccc", r)
		}
	}
	close(v)
	if r, ok := <-v; ok == false {
		fmt.Println("bbbbbbbbbbbbbbbbbbbbbbb", r)
		close(v)
	}
	for b := range v {
		fmt.Println(b)
	}
}

func main() {
	str := fmt.Sprintf(`'%%sddd'`)
	fmt.Println(fmt.Sprintf("%s", str))
	m1 := make(map[chan []byte]bool)
	for i := 1; i < 1000000; i++ {
		ch := make(chan []byte, 128)
		m1[ch] = false
	}
	time.Sleep(time.Second * 100)
	return
	ints := []int{}
	ints = append(ints, 1)
	ints = append(ints, 2)
	ints = append(ints, 3)
	ints = append(ints, ints...)
	for v := range ints {
		fmt.Println(v)
	}
	fmt.Println(ints)
	CheckChan()
	v := make(chan bool, 1)
	v <- true
	if v != nil {
		close(v)
	}
	for b := range v {
		fmt.Println(b)
	}
	return
	//a := []int{1, 2, 3, 4, 5, 6, 7}
	m := make(map[int]chan int)
	for i := 1; i < 1000; i++ {
		c := make(chan int, 1)
		m[i] = c
		m[i] <- i
		go func() {
			loop := true
			for loop {
				select {
				case cc, ok := <-c:
					fmt.Println(cc, i, ok)
					//loop = ok
				default:
					time.Sleep(time.Millisecond * 50)
				}
			}
			fmt.Println("aaaaaaaaaaa", i, loop)
		}()
	}
	time.Sleep(time.Second * 10)
	//for _, v := range m {
	//close(v)
	//}

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
