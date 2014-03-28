package main

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

func main() {
	//cmd, _ := exec.Command("/bin/sh", "/home/whj/gostudy/test.sh").Output()
	//fmt.Println(string(cmd))
	cmd := exec.Command("/bin/sh", "/home/whj/gostudy/test.sh")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
	}
	b := bytes.NewBuffer(nil)
	b.ReadFrom(stdout)
	line := make([]byte, 1023)
	for {
		n, err := stdout.Read(line)
		if n > 0 {
			fmt.Println(n, string(line))
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(n, err)
		}
	}
	if err := cmd.Wait(); err != nil {
		fmt.Println(err)
	}
}
