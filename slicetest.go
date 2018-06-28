package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}
	if len(a) > 5 {
	}
	b := []int{4, 5, 6}
	copy(b[2:], a)
	fmt.Println(b)
}
