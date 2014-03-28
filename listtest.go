package main
import (
	l "container/list"
	"fmt"
	"time"
)

var (
	name = "viney"
)

func list() {
	names := l.New()
	t := time.Now()
	for i := 1; i <= 1000000; i++ {
		_ = names.PushFront(name)
	}
	fmt.Println("list: " + time.Now().Sub(t).String())
}

func slice() {
	names := []string{}
	t := time.Now()
	for i := 1; i <= 1000000; i++ {
		names = append(names, name)
	}
	fmt.Println("slice: " + time.Now().Sub(t).String())
}

func main() {
	list()
	slice()
}

