package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	copy(b[2:], a)
	fmt.Println(b)
}
