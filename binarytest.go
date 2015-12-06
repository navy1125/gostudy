package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	//"reflect"
	"time"
)

type TestStruct struct {
	i int
}

func TestParm(i int) {
	fmt.Println(i)
}

func main() {
	whj := "wanghaijun"
	bb := [32]byte{}
	copy(bb[:len(whj)], []byte(whj))
	//b[:len(whj)] = []byte(whj)
	//bbuf := bytes.NewBuffer(b)
	//bbuf.Write([]byte(whj))
	fmt.Println(bb)

	//test := &TestStruct{1}
	conn := &net.TCPConn{}
	var pi float64
	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	buf := bytes.NewBuffer(b)
	buf.WriteByte(0x09)
	buf.WriteByte(0x40)
	err := binary.Read(buf, binary.LittleEndian, &pi)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	//fmt.Print(pi)
	go func() {
		err := binary.Read(conn, binary.LittleEndian, &pi)
		if err != nil {
			fmt.Println("binary.Read failed:", err)
		}
	}()
	go func() {
		tick := time.Tick(time.Second)
		for {
			select {
			case t := <-tick:
				fmt.Println(t)
				//tt := time.Time
			}
		}
	}()
	for _, c := range b {
		cb := []byte{c}
		conn.Write(cb)
		time.Sleep(time.Second)
	}
	//fmt.Println(reflect.TypeOf().Size())
}
