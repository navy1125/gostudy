package gotcp

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"time"
)

type Task struct {
	Id              int64
	Name            string
	Conn            *net.TCPConn
	in              chan []byte
	stop            chan bool
	handleReadFun   func(conn *net.TCPConn) ([]byte, error)
	handleWriteFun  func(conn *net.TCPConn, data []byte) error
	handleParse     func(conn *net.TCPConn, data []byte) (int, error)
	handleHeartBeat func(conn *net.TCPConn)
	heartBeatTime   time.Duration
	HeartBeatReturn bool
}

func NewTask(c *net.TCPConn) *Task {
	task := &Task{
		conn:            c,
		in:              make(chan []byte),
		stop:            make(chan bool),
		handleReadFun:   handleReadFunDefault,
		handleWriteFun:  handleWriteFunDefault,
		handleParse:     handleParseDefault,
		handleHeartBeat: handleHeartBeatRequestDefault,
		heartBeatTime:   time.Second * 10,
		heartBeatReturn: true,
	}
	return task
}
func handleHeartBeatRequestDefault(conn *net.TCPConn) {
	handleWriteFunDefault(conn, []byte("tick"))
}
func handleHeartBeatReturnDefault(conn *net.TCPConn) {
	handleWriteFunDefault(conn, []byte("return tick"))
}
func handleReadFunDefault(conn *net.TCPConn) ([]byte, error) {
	var length uint32
	if err := binary.Read(conn, binary.BigEndian, &length); err != nil {
		return nil, err
	}
	buf := make([]byte, length)
	_, err := io.ReadFull(conn, buf)
	return buf, err
}
func handleWriteFunDefault(conn *net.TCPConn, data []byte) error {
	err := binary.Write(conn, binary.BigEndian, uint32(len(data)))
	if err != nil {
		return err
	}
	_, err = conn.Write(data)
	return err
}
func handleParseDefault(conn *net.TCPConn, data []byte) (int, error) {
	switch {
	case string(data) == "tick":
		fmt.Println("time tick return")
		handleHeartBeatReturnDefault(conn)
		return 1, nil
	case string(data) == "return tick":
		return 1, nil
	}
	return 0, nil
}
func (self *Task) Start() {
	go self.startRead()
	go self.startWrite()
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
			fmt.Println("read error:", err)
			break
		}
		self.in <- data
	}
}
func (self *Task) startWrite() {
	heartBeat := time.Tick(self.heartBeatTime)
LOOP:
	for {
		select {
		case <-heartBeat:
			self.handleHeartBeat(self.conn)
			if !self.heartBeatReturn {
				fmt.Println("timeout error")
				break LOOP
			}
			self.heartBeatReturn = false
		case data, ok := <-self.in:
			if !ok {
				fmt.Println("self.in chan err")
				break LOOP
			}
			if ret, _ := self.handleParse(self.conn, data); ret == 1 {
				self.heartBeatReturn = true
			}
		}
	}
}
