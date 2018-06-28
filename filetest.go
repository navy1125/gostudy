package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	runname := "run.bat"
	fr, err := os.Open(runname)
	if err != nil {
		fmt.Println("aa", err)
	}
	b := []byte{}
	//b := make([]byte, 1024)
	n, err := fr.Read(b)
	if err != nil {
		fmt.Println("bb", string(b), err)
	}
	str := string(b[:n])
	str = strings.Replace(str, ".json", ".xml", -1)
	fr.Close()
	fw, err := os.Create(runname)
	if err != nil {
		fmt.Println("dd", err)
	}
	_, err = fw.WriteString(str)
	if err != nil {
		fmt.Println(err)
	}
	fw.Close()

}
