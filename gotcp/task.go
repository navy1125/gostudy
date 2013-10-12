package gotcp

import (
	"fmt"
	"net"
	"time"
)

type Task struct {
	Id             int64
	Name           string
	conn           *net.TCPConn
	in             chan []byte
	stop           chan bool
	handleReadFun  func(conn *net.TCPConn) ([]byte, error)
	handleWriteFun func(conn *net.TCPConn, buf []byte) error
}

func NewTask(c *net.TCPConn) *Task {
	task := &Task{
		conn:           c,
		in:             make(chan []byte),
		stop:           make(chan bool),
		handleReadFun:  handleReadFunDefault,
		handleWriteFun: handleWriteFunDefault,
	}
	go task.startRead()
	go task.startWrite()
	return task
}
func handleReadFunDefault(conn *net.TCPConn) ([]byte, error) {
	return nil, nil
}
func handleWriteFunDefault(conn *net.TCPConn, buf []byte) error {
	return nil
}
func (self *Task) SendCmd(cmd []byte, len int) error {
	return nil
}
func (self *Task) Stop() {
	self.conn.Close()
	close(self.stop)
}
func (self *Task) startRead() {
	defer close(self.in)
	for {
		data, err := self.handleReadFun(self.conn)
		if err != nil {
			break
		}
		self.in <- data
	}
}
func (self *Task) startWrite() {
	heartBeat := time.Tick(time.Second)
	for {
		select {
		case <-heartBeat:
			self.SendCmd([]byte("tick"), len("tick"))
		case data := <-self.in:
			fmt.Println(string(data))
		}
	}
}
