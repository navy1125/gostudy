package main

import (
	gotcp "../gotcp"
	//"bytes"
	"fmt"
	"github.com/navy1125/config"
	"math/rand"
	"net"
	"time"
	//"unsafe"
)

func main() {
	if err := config.LoadFromFile("loginServerList.xml", "GMServerList"); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(config.GetConfigStr("ip") + ":" + config.GetConfigStr("port"))
	raddr, _ := net.ResolveTCPAddr("tcp", config.GetConfigStr("ip")+":"+config.GetConfigStr("port"))
	conn, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		fmt.Println("conn err:", err)
		return
	}
	fmt.Println("new connection:", conn.RemoteAddr())
	task := gotcp.NewTask(conn, "Client")
	task.SetHandleReadFun(handleReadFunBw)
	task.SetHandleWriteFun(handleWriteFunBw)
	task.SetHandleParseFun(handleParseBw)
	task.SetHandleHeartBteaFun(handleHeartBeatRequestBw, time.Second*10)
	cmd := NewstRequestLoginGmUserCmd()
	cmd.version = 2012102901
	for i, b := range "webmaster" {
		cmd.name[i] = byte(b)
	}
	for i, b := range "TZ95d5KV" {
		cmd.password[i] = byte(b)
	}
	//buf := bytes.NewBuffer(cmd.name)
	//buf.WriteString("webmaster")
	//buf = bytes.NewBuffer(cmd.password)
	//buf.WriteString("TZ95d5KV")
	task.SendCmd(cmd)
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
