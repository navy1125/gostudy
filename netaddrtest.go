package main

import (
	"log"
	"net"
)

type FOO func() string

func main() {
	eth0, _ := net.InterfaceByName("eth0")
	addrs, _ := eth0.Addrs()
	log.Println(addrs[0].String())
}
