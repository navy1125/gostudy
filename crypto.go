package main

import (
	"crypto"
	"crypto/md5"
	"fmt"
	//"strconv"
)

const (
	TestFlag0 = 1 << iota
	TestFlag1
	TestFlag2
	TestFlag3
)

func main() {
	m := md5.New()
	str := "test"
	m.Write([]byte(str))
	fmt.Printf("%x\n", m.Sum(nil))

	crypto.RegisterHash(crypto.MD5, md5.New)
	fmt.Println(crypto.MD5.Available())
	cm := crypto.MD5.New()
	cm.Write([]byte(str))
	//fmt.Printf("%x\n", cm.Sum(nil))
	fmt.Printf("%x\n", cm.Sum([]byte("test")))
	fmt.Printf("TestFlag0:%d,TestFlag1:%d,TestFlag2:%d,TestFlag3:%d\n", TestFlag0, TestFlag1, TestFlag2, TestFlag3)
}
