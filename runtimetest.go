// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	//"git.code4.in/logging"
	"reflect"
	"runtime"
	"runtime/debug"
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
func main() {
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
	i := 100
	fmt.Println(reflect.TypeOf((*int32)(nil)).Elem())
	fmt.Println(reflect.ValueOf(&i).Kind())
	fmt.Println(reflect.TypeOf(&i).Kind())
	RoundPos(3, func(r, x, y int32) {
		println(r, x, y)
	})
}
