package main

import (
	gotcp "../gotcp"
	"fmt"
	"net"
	"time"
)

func main() {
	raddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:7000")
	listen, err := net.ListenTCP("tcp", raddr)
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
	go func() {
		heartBeat := time.Tick(time.Second * 5)
		for {
			_, ok := <-heartBeat
			if !ok {
				fmt.Println("server heatBteat error")
				break
			}
			//fmt.Println("server heatBteat", h)
		}
	}()
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println("listen error:", err)
			continue
		}
		fmt.Println("new connection:", conn.RemoteAddr())
		task := gotcp.NewTask(conn, "Server")
		task.Start()
	}
}
