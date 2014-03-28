package main

import (
	"code.google.com/p/go.net/websocket"
	"git.code4.in/logging"

	"log"
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
	var i FooInterface
	my := &MyEntry{}
	i = my
	Foo(i)
	m := map[int]*websocket.Conn{}
	a, _ := m[1]
	log.Println("xxxxxx", a)
	logging.Info("server start...")

	origin := "http://192.168.85.71:8000/shen/user"
	url := "ws://192.168.85.71:8000/shen/user"
	//origin := "http://echo.websocket.org"
	//url := "ws://echo.websocket.org:80"

	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Println("websocket.Dial: %s", err.Error())
		logging.Error("websocket.Dial: %s", err.Error())
		return
	}

	var msg = make([]byte, 10)
	msg[0] = 1
	msg[1] = 1
	msg[2] = 10
	msg[3] = 1
	msg[4] = 49
	websocket.Message.Send(ws, msg)
	var n int
	if n, err = ws.Read(msg); err != nil {
		log.Println("ws.Read: %s", err.Error())
		logging.Error("ws.Read: %s", err.Error())
		return
	}
	logging.Debug("Received: %s.\n", msg[:n])
	logging.Info("server stop...")
	time.Sleep(time.Second)
}
