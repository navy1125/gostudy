package main

import (
	//"code.google.com/p/go.net/websocket"
	"io"
	"log"
	"net"
	"time"
)

type Entry struct {
	Id uint64
}

type MyEntry struct {
	Entry
}

func (self *MyEntry) Foo() {

}

type FooInterface interface {
	Foo()
}

func Foo(v FooInterface) {
	//log.Println(v.(MyEntry).Id)
}
func main() {
	/*
		origin := "http://14.17.104.56:7000/shen/user"
		url := "ws://14.17.104.56:7000/shen/user"
		//origin := "http://103.44.146.57:20057"
		//url := "ws://103.44.146.57:20057"
		//origin := "http://echo.websocket.org"
		//url := "ws://echo.websocket.org:80"

		ws, err := websocket.Dial(url, "", origin)
		if err != nil {
			log.Println("websocket.Dial:", err.Error())
			return
		}
		var msg = make([]byte, 10)
		msg[0] = 1
		msg[1] = 1
		msg[2] = 10
		msg[3] = 1
		msg[4] = 49
		websocket.Message.Send(ws, msg)
		if _, err = ws.Read(msg); err != nil {
			log.Println("ws.Read: %s", err.Error())
			return
		}
		// */
	var tcpAddr *net.TCPAddr
	//tcpAddr, _ = net.ResolveTCPAddr("tcp", "103.44.146.57:20057")
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "103.44.146.42:10514")
	//tcpAddr, _ = net.ResolveTCPAddr("tcp", "14.17.104.56:7000")
	cnn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err == nil {
		log.Println("connect ok")
		n, err := cnn.Write([]byte("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"))
		if err == nil {
			log.Println("xxxxx0:%v", n)
		} else {
			log.Println("xxxxx0:%s", err.Error())
		}
		buf := make([]byte, 1024*1000)
		_, err = io.ReadFull(cnn, buf)
		if err == nil {
			log.Println("xxxxx1:%s", string(buf))
		} else {
			log.Println("xxxxx1:%s", err.Error())
		}
	} else {
		log.Println("xxxxx2:%s", err.Error())
	}
	time.Sleep(time.Second)
}
