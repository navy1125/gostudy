package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("GOTRACEBACK:", os.Getenv("RACEBACK"))
	//now := time.Now()
	//time.Sleep(1000000000)
	fmt.Println(time.Now().UnixNano(), time.Now().Unix())
	t := time.Date(2009, time.December, 30, 13, 13, 13, 13, time.UTC)
	fmt.Printf("%s,%d\n", t.Local(), time.Millisecond)
	seconds := 10
	var dur time.Duration
	fmt.Println(dur.Hours())
	fmt.Printf("%d\n", time.Duration(seconds)*time.Second)
	//zone, sec := now.Zone()
	//fmt.Println(now.Unix(), now.Sub(1111), zone, sec)
	go func() {
		for {
			select {
			case <-time.After(2 * time.Second):
				fmt.Println("2time.After:", time.Now().Unix())
			case <-time.After(1 * time.Second):
				fmt.Println("1time.After:", time.Now().Unix())
			}
		}
	}()
	timer := time.AfterFunc(time.Second*2, timefunc)
	defer timer.Stop()
	for i := 0; i < 30; i++ {
		time.Sleep(time.Second * 2)
		fmt.Println("sleep:", time.Now().Unix())
	}
}
func timefunc() {
	fmt.Println("timer:", time.Now().Unix())
	time.AfterFunc(time.Second*2, timefunc)
}
