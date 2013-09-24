package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	text, err := ioutil.ReadFile("D:\\work\\gostudy\\client.go")
	if err != nil {
		log.Fatal("ioutil.ReadFile err:", err)
	}
	fmt.Printf("%s", text)

	test := &TestInfo{}
	test.str = "whj"
	Info(test)

}

type TestInterface interface {
	Test()
}
type TestInfo struct {
	str string
}

func (t *TestInfo) Test() {
	fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", t.str)
}
func Info(v TestInterface) {
	f, ok := v.(TestInterface)
	if ok {
		f.Test()
	}
}
