package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
	"unsafe"
)

type Point struct {
	x int8
	y int16
}
type Rect struct {
	pt     Point
	with   int16
	height int16
	data   [2]byte
}

func main() {
	buf := new(bytes.Buffer)
	rect := *new(Rect)
	i := int16(rect.data[:1])
	fmt.Println(reflect.TypeOf(rect).Size(),reflect.TypeOf(rect).Name())
	reflect.Type
	fmt.Println(unsafe.Sizeof(rect))
	pt := (*Point)(unsafe.Pointer(&rect))
	fmt.Println(rect, pt)
	err := binary.Write(buf, binary.LittleEndian, rect)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	fmt.Println(len(buf.Bytes()), len(toBytes(rect)))
}

func toBytes(v interface{}) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, v)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	return buf.Bytes()
}
