package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

type Interface interface {
	String() string
	GetMe() Interface
}

type StructInterface struct {
}

func (self *StructInterface) String() string {
	return "aaaa"
}
func (self *StructInterface) GetMe() Interface {
	return self
}

type StructInterface2 struct {
	StructInterface
}

func (self *StructInterface2) String() string {
	return "bbbb"
}

func main() {
	st := &StructInterface{}
	st2 := &StructInterface2{}
	var it Interface
	it = st
	fmt.Println(it.GetMe().String())
	it = st2
	fmt.Println(it.GetMe().(*StructInterface2).String())
	return
	begin := time.Now().UnixNano()
	fmt.Println("begin:", begin)
	for i := 0; i < 1000000000; i++ {
		if ss := st; ss != nil {
			ss.String()
		}
	}
	fmt.Println("end:", time.Now().UnixNano()-begin)
	begin = time.Now().UnixNano()
	fmt.Println("begin:", begin)
	for i := 0; i < 1000000000; i++ {
		//if ss, _ := it.(*StructInterface); ss != nil {
		if ss, _ := it.(*StructInterface); ss != nil {
			ss.String()
		}
	}
	fmt.Println("end:", time.Now().UnixNano()-begin)
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
	fmt.Println("type nil:", t, vv)
	go Say("World")
	Say("Hellow")

}
func Say(str string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(str, i)
	}
}
