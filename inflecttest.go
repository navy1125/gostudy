package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type MyStruct struct {
	name string
}

func (this *MyStruct) GetName() string {
	return this.name
}

type Point struct {
	MyStruct
	x int
	y int
}

func main() {
	fmt.Println("--------------")
	var a MyStruct
	b := new(MyStruct)
	t := reflect.TypeOf(b)
	fmt.Println("t.Name():", t, t.String(), t.Elem().Field(0).Name)
	ListTypeName(b)
	fmt.Println(reflect.ValueOf(a))
	fmt.Println(reflect.ValueOf(b), reflect.TypeOf(b).Size())

	fmt.Println("--------------")
	a.name = "yejianfeng"
	b.name = "yejianfeng"
	fmt.Println(reflect.ValueOf(b), reflect.TypeOf(b).Size())
	var pp Point
	pp.name = "whj"
	pp.x = 1
	pp.y = 1
	var bp *MyStruct
	bp = (*MyStruct)(unsafe.Pointer(&pp))
	bp.name = "wanghaijun"
	//pp := Point{1, 1}
	//val := reflect.ValueOf(a).FieldByName("name")
	//val := reflect.ValueOf(a).Field(0)

	//painc: val := reflect.ValueOf(b).FieldByName("name")
	fmt.Println("1111111111111", pp.name)

	fmt.Println("--------------")
	fmt.Println(reflect.ValueOf(a).FieldByName("name").CanSet())
	fmt.Println(reflect.ValueOf(&(a.name)).Elem().CanSet())

	fmt.Println("--------------")
	var c string = "yejianfeng"
	p := reflect.ValueOf(&c)
	fmt.Println(p.CanSet())        //false
	fmt.Println(p.Elem().CanSet()) //true
	p.Elem().SetString("newName")
	fmt.Println(c)
	slice := make([]int, 1)
	fmt.Println(slice)
}
func ListTypeName(v interface{}) {
	t := reflect.TypeOf(v)
	fmt.Println("t.Name():", t.String(), t.Elem().Field(0).Name)
}
