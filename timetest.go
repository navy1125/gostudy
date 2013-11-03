package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	//time.Sleep(1000000000)
	t := time.Date(2009, time.December, 30, 13, 13, 13, 13, time.UTC)
	fmt.Printf("%s,%d\n", t.Local(), time.Millisecond)
	seconds := 10
	fmt.Println(time.Duration(seconds) * time.Second)
	zone, _ := now.Z
	fmt.Println(now.Unix(), zone)
	timer := time.AfterFunc(time.Second*2, timefunc)
	defer timer.Stop()
	for i := 0; i < 10; i++ {
		//time.Sleep(time.Second * 2)
		fmt.Println("sleep:", time.Now().Unix())
	}
}
func timefunc() {
	fmt.Println("timer:", time.Now().Unix())
	time.AfterFunc(time.Second*2, timefunc)
}
