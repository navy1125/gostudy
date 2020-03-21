package main

import (
	"fmt"
	"reflect"
	"strings"
)

func RoundPos(dep int32, cb func(r, x, y int32)) {
	cb(0, 0, 0)
	var n, m int32
	for n = 1; n <= dep; n++ {
		for m = -n; m < n; m++ {
			cb(n, m, n)
		}
		for m = n; m > -n; m-- {
			cb(n, n, m)
		}
		for m = n; m > -n; m-- {
			cb(n, m, -n)
		}
		for m = -n; m < n; m++ {
			cb(n, -n, m)
		}
	}
}

type Test struct {
	X *int
	y int
}

func (t *Test) Sstring() {
	fmt.Println(t.y)
}

func main() {
	test := &Test{}
	tv := reflect.ValueOf(test)
	tt := reflect.New(reflect.ValueOf(tv).Type())
	fmt.Println(tt.Type())
	X := 2
	fff := 1.3000
	fmt.Printf("xxxxxx:%v,%s\n", fff, strings.Replace(".000000", ".000000$", "", -1))
	xc := reflect.ValueOf(test).MethodByName("Sstring")
	xc.Call([]reflect.Value{})
	//fmt.Println("----------------------------", xc.CanAddr(), xc.CanSet(), xc.IsNil(), xc.IsValid(), xc.Type(), xc.Kind(), xc.Elem().Kind())
	in := []reflect.Value{}
	xc.Call(in)
	xv := reflect.ValueOf(test).Elem().FieldByName("X")
	fmt.Println("----------------------------", xv.CanAddr(), xv.CanSet(), xv.IsNil(), xv.IsValid(), xv.Type(), xv.Kind(), xv.Elem().Kind())
	iv := reflect.ValueOf(&X)
	//test.X = &i
	//reflect.ValueOf(test).Elem().FieldByName("X").Elem().SetInt(1)
	//reflect.New(dst.Type().Elem()).Elem()
	xxv := reflect.New(xv.Type().Elem())
	fmt.Println("----------------------------", xxv.CanAddr(), xxv.CanSet(), xxv.IsNil(), xxv.IsValid(), xxv.Type(), xxv.Kind(), xxv.Elem().Kind())
	xxv.Elem().SetInt(4)
	fmt.Println("reflect.New xxv", xxv.Type(), xxv.String())
	fmt.Println(xv.Type(), reflect.PtrTo(xxv.Elem().Type()), reflect.ValueOf(xxv).Type())
	//xxv.Elem().Set(xv)
	//xv.Set(reflect.New(reflect.ValueOf(xv).Type()))
	xv.Set(xxv)
	fmt.Println("----------------------------", xv.CanAddr(), xv.CanSet(), xv.IsNil(), xv.IsValid(), xv.Type(), xv.Kind(), xv.Elem().Kind())
	iv.Elem().SetInt(5)
	fmt.Println("xxv", xxv.Elem().Type(), xxv.String())
	ivn := reflect.New(iv.Type())
	fmt.Println(ivn.Type())
	//ivn.Elem().SetInt(6)
	//xv.Set(iv)
	//xv.Elem().SetInt(3)
	fmt.Println(reflect.ValueOf(test).Elem().FieldByName("y"), *test.X)
	xv.Set(iv)
	/*
		test1 := Test{}
		test1.next = &Test{x: 2}
		test1.x = 1
		test1.Func()
		fmt.Println("aaaa", reflect.TypeOf(test1).String(), len(strings.Split("aaa", "3")))
		fmt.Println("bbbb", reflect.TypeOf(test1).Name())
		//reflect.ValueOf(test1).FieldByName("x").SetInt(111)
		fmt.Println(reflect.ValueOf(test1).FieldByName("next").Elem().Elem().FieldByName("x"))
		//fmt.Println(((interface{})(reflect.ValueOf(test1).FieldByName("next").Elem().Pointer())).(*Test))
		//fmt.Println(((*Test)((interface{})(reflect.ValueOf(test1).FieldByName("next").Elem().Pointer()))))
		//fmt.Println(reflect.ValueOf(test1).FieldByName("next").Convert(reflect.TypeOf(test1)).Elem().FieldByName("x"))
		//fmt.Println(fmt.Sprintf("aaaa:%d", reflect.ValueOf(test1).FieldByName("next").Elem().FieldByName("x").Int()))
		fmt.Println("aaaa:", runtime.NumCPU())
		select {
		case <-time.After(time.Second):
			fmt.Println(time.Minute.Seconds(), time.Second)
		}
		select {
		case <-time.After(time.Minute):
			fmt.Println(time.Minute.Seconds(), time.Second)
		}
		var kvlist []int
		kvlist = make([]int, 0, 10)
		kvlist = append(kvlist, 1)
		fmt.Println(kvlist)
		m := make(map[int]int, 10)
		m[1] += 1
		fmt.Println("aaaa:", string(runtime.CPUProfile()), m, cap(kvlist))
		for i := 0; ; i++ {
			pc, file, line, ok := runtime.Caller(i)
			if ok == false {
				break
			}
			fmt.Println(pc, file, line, ok)
			//runtime.Breakpoint()
		}
		fmt.Println(runtime.Version())
		buf := make([]byte, 1025)
		len := runtime.Stack(buf, true)
		fmt.Println(string(runtime.CPUProfile()))
		fmt.Println(len, string(buf[:len]))
		// */

	//RoundPos(3, func(r, x, y int32) {
	//	println(r, x, y)
	//})
}
