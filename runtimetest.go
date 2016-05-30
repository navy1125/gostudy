// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	//"git.code4.in/logging"
	"reflect"
	"runtime"
	"runtime/debug"
	"strings"
	"time"
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
	x    int
	next interface{}
}

func (self *Test) Func() {
	fname, _, _, _ := runtime.Caller(0)
	f := runtime.FuncForPC(fname)
	fmt.Println(f.Name())
}

func main() {
	fmt.Println(debug.SetMaxThreads(10))
	fmt.Println(debug.SetMaxStack(1 << 20))
	fmt.Println(debug.SetMaxStack(1 << 20))
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
	//debug.PrintStack()
	fmt.Println(debug.SetMaxThreads(10))
	fmt.Println(debug.SetMaxStack(10))

	i := 100
	fmt.Println(reflect.TypeOf((*int32)(nil)).Elem())
	fmt.Println(reflect.ValueOf(&i).Kind())
	fmt.Println(reflect.TypeOf(&i).Kind())
	RoundPos(3, func(r, x, y int32) {
		println(r, x, y)
	})
}
