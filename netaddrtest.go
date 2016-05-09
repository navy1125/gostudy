package main

import (
	"fmt"
	"net"
)

type FOO func() string

func main() {
	eth0, err := net.InterfaceByName("bridge0")
	if err == nil {
		addrs, err := eth0.Addrs()
		if err == nil && len(addrs) > 0 {
			fmt.Println(addrs[0].String())
		}
	}
	//ips, err := net.LookupIP("www.baidu.com")
	ips, err := net.LookupIP("1.1.1.1")
	if err == nil {
		fmt.Println(ips)
	}
	//ip, err := net.ResolveIPAddr("ip", "localhost")
	ip, err := net.ResolveIPAddr("ip", "www.baidu.com")
	fmt.Println(ip)
}
