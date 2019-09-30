package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Test struct {
}

func (self *Test) Foo(test string) {
	fmt.Println(test)
}

type TestFunc func(t *Test, test string)

func main() {
	name := "wang"
	fmt.Println(fmt.Sprintf("xxxxxxxxx:%s", strings.ToUpper(name[:4])))
	fmt.Println(fmt.Sprintf("xxxxxxxxx:%s", strings.ToLower(name[:4])))
	fmt.Println(fmt.Sprintf("xxxxxxxxx:%d,%d", int64(-1), int32(40649595150)))
	var listi []int
	listi = append(listi, 1)
	for i, v := range listi {
		fmt.Println(i, len(listi), v)
	}
	var b []byte
	s := string(b)
	fmt.Println(s)
	test := &Test{}
	var tf TestFunc
	tf = (*Test).Foo
	tf(test, "wanghi")
	//var m map[int]string
	fmt.Println(reflect.TypeOf(test).Kind())
	fmt.Println(reflect.ValueOf(test).Kind(), reflect.ValueOf(test).Elem().Kind())
	fmt.Println(strings.Replace("aa.json", ".json", "aa.xml", 1))
	var sa *string
	sa = new(string)
	*sa = "whj"
	var ss string
	ss = *sa
	*sa = "xxx"
	fmt.Println(ss, *sa)
}
