// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	lim := &syscall.Rlimit{}

	syscall.Getrlimit(syscall.RLIMIT_NOFILE, lim)
	lim.Cur = 15
	syscall.Setrlimit(syscall.RLIMIT_NOFILE, lim)
	fmt.Println(lim)
	var s []*os.File
	for i := 0; i < int(lim.Max+10); i++ {
		f, err := os.Open("/tmp/whj.txt")
		fmt.Println(f, err)
		s = append(s, f)
		err = syscall.Getrlimit(syscall.RLIMIT_NOFILE, lim)
		fmt.Println(lim, err)
	}
	err := syscall.Getrlimit(syscall.RLIMIT_CORE, lim)
	fmt.Println(lim, err)
	rusage := &syscall.Rusage{}
	err = syscall.Getrusage(0, rusage)
	fmt.Println(rusage, err)
	err = syscall.Getrusage(1, rusage)
	fmt.Println(rusage, err)
	err = syscall.Getrusage(2, rusage)
	fmt.Println(rusage, err)
	err = syscall.Getrusage(3, rusage)
	fmt.Println(rusage, err)
	err = syscall.Getrusage(4, rusage)
	fmt.Println(rusage, err)
}
