package main

import (
	gotcp "../gotcp"
	"fmt"
	"math/rand"
	"net"
	"time"
)

func main() {
	raddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:7000")
	conn, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		fmt.Println("conn err:", err)
		return
	}
	fmt.Println("new connection:", conn.RemoteAddr())
	task := gotcp.NewTask(conn, "Client")
	task.Id = rand.Int63()
	task.Name = conn.RemoteAddr().String()
	task.Start()
	tick := time.Tick(time.Second * 5)
	for {
		select {
		case <-tick:
			//case t := <-tick:
			//	fmt.Println(t)
		}
	}
	time.Sleep(time.Second * 10)
	task.Stop()
}
