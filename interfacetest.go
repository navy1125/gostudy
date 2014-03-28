package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	var i interface{}
	a := 5
	i = a
	if v, ok := i.(int); ok {
		fmt.Println(v)
	}
	switch v := i.(type) {
	case int:
		fmt.Println(v)
	case string:
		fmt.Println("string", v)
	}
	t := reflect.TypeOf(nil)
	vv := reflect.New(t)
	v := reflect.ValueOf(i)
	fmt.Println("tyep:", v.Type())
	fmt.Println("type nil:", t,vv)
	go Say("World")
	Say("Hellow")

}
func Say(str string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(str, i)
	}
}
