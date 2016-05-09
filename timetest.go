package main

import (
	"errors"
	"fmt"
	"git.code4.in/mobilegameserver/unibase"
	"net"
	"os"
	"strings"
	"time"
)

func ParseTime(datetime string) (time.Time, error) {
	dt := strings.Split(datetime, " ")
	if len(dt) != 2 {
		return time.Time{}, errors.New("format datetime err")
	}
	ds := strings.Split(dt[0], "-")
	if len(ds) != 3 {
		return time.Time{}, errors.New("format date string err")
	}
	ts := strings.Split(dt[1], ":")
	if len(ts) != 3 {
		return time.Time{}, errors.New("format time string err")
	}
	year := unibase.Atoi(ds[0])
	month := unibase.Atoi(ds[1])
	day := unibase.Atoi(ds[2])
	hour := unibase.Atoi(ts[0])
	min := unibase.Atoi(ts[1])
	sec := unibase.Atoi(ts[2])
	return time.Date(year, time.Month(month), day, hour, min, sec, 0, time.Now().Location()), nil
}

type Test struct {
	offset int
}

func (self *Test) Add() int {
	self.offset += 1
	return self.offset
}

func main() {
	tt := &Test{}
	fmt.Println(tt.offset)
	fmt.Println(tt.Add())
	fmt.Println(tt.offset)
	return
	fmt.Println(unibase.Atoi("192.158.68.2"))
	d, e := time.ParseDuration("1s")
	fmt.Println(d, e)
	fmt.Println(int(time.Millisecond), int(time.Millisecond.Nanoseconds()))
	year, month, day := time.Now().Date()
	fmt.Println(year, month, day)
	ip := net.ParseIP("1.1.1.11").To4()
	fmt.Println(uint32(ip[0])<<24 + uint32(ip[1])<<16)
	loc, _ := time.LoadLocation("")
	fmt.Println(time.Now().Unix(), time.Now().Local().Unix(), time.Now().Second(), time.Now().Location().String(), loc.String())
	pt, err := ParseTime("2012-12-12 12:12:12")
	fmt.Println(pt.Unix(), err)
	now := time.Now()
	_, offset := now.Zone()
	fmt.Println(int(time.Duration(offset)))
	begin := time.Duration(offset) * time.Second
	zonenow := now.Add(-begin)
	fmt.Println(now.Unix(), zonenow.Unix())
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
