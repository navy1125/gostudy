package main

import (
	"bytes"
	"container/list"
	"encoding/binary"
	"fmt"
	"time"
)

var (
	name = "viney"
)

func listTest() {
	names := list.New()
	t := time.Now()
	for i := 1; i <= 1000000; i++ {
		_ = names.PushBack(name)
	}
	names.Init()
	for i := 1; i <= 1000000; i++ {
		_ = names.PushBack(name)
	}
	fmt.Println("list: " + time.Now().Sub(t).String())
}

func slice() {
	//names := []string{}
	//names := make([]string, 0, 1000000)
	var names []string
	var names1 []string
	names = append(names, "whj")
	names1 = append(names1, "whj")
	names1 = append(names1, "whj1")
	names = append(names, names1...)
	fmt.Println(names)
	t := time.Now()
	for i := 1; i <= 1000000; i++ {
		names = append(names, name)
	}
	for name := range names[0:10] {
		names = append(names, names[0])
		fmt.Println(name)
	}
	names = names[0:0]
	for i := 1; i <= 1000000; i++ {
		names = append(names, name)
	}
	fmt.Println("slice: " + time.Now().Sub(t).String())
}
func sendDataList(datalist *list.List) {
	sendbuf := bytes.NewBuffer(nil)
	binary.Write(sendbuf, binary.LittleEndian, []byte("["))
	for e := datalist.Front(); e != nil; e = e.Next() {
		binary.Write(sendbuf, binary.LittleEndian, e.Value)
		binary.Write(sendbuf, binary.LittleEndian, []byte(","))
	}
	binary.Write(sendbuf, binary.LittleEndian, []byte("]"))
	fmt.Println(string(sendbuf.Bytes()))
}

func main() {
	listTest()
	slice()
	l := list.New()
	l.PushBack([]byte(`{}`))
	l.PushBack([]byte(`{}`))
	l.PushBack([]byte(`{}`))
	l.PushBack([]byte(`{}`))
	sendDataList(l)
}
