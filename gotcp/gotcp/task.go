package gotcp

import (
	"fmt"
	"net"
	"time"
)

type Task struct {
	Entry
	Conn            *net.TCPConn
	in              chan []byte
	stop            chan bool
	handleReadFun   func(task *Task) ([]byte, error)
	handleWriteFun  func(task *Task, data []byte) error
	handleParse     func(task *Task, data []byte) error
	handleHeartBeat func(task *Task)
	heartBeatTime   time.Duration
	HeartBeatReturn bool
}

func NewTask(c *net.TCPConn, name string) *Task {
	task := &Task{
		Conn:            c,
		in:              make(chan []byte),
		stop:            make(chan bool),
		handleReadFun:   handleReadFunDefault,
		handleWriteFun:  handleWriteFunDefault,
		handleParse:     handleParseDefault,
		handleHeartBeat: handleHeartBeatRequestDefault,
		heartBeatTime:   time.Second * 10,
		HeartBeatReturn: true,
	}
	task.GetEntryName = func() string { return name }
	return task
}
func (self *Task) SetHhandleReadFun(fun func(task *Task) ([]byte, error)) {
	self.handleReadFun = fun
}
func (self *Task) SetHhandleWriteFun(fun func(task *Task, data []byte) error) {
	self.handleWriteFun = fun
}
func (self *Task) SetHhandleParseFun(fun func(task *Task, data []byte) error) {
	self.handleParse = fun
}
func (self *Task) SetHhandleHeartBteaFun(fun func(task *Task), dur time.Duration) {
	self.handleHeartBeat = fun
	self.heartBeatTime = dur
}
func (self *Task) Start() {
	go self.startRead()
	go self.startWrite()
}
func (self *Task) Stop() {
	self.Conn.Close()
	close(self.stop)
}
func (self *Task) startRead() {
	defer close(self.in)
	for {
		data, err := self.handleReadFun(self)
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
			self.handleHeartBeat(self)
			if !self.HeartBeatReturn {
				self.Debug("timeout error")
				break LOOP
			}
			self.HeartBeatReturn = false
		case data, ok := <-self.in:
			if !ok {
				self.Debug("self.in chan err")
				break LOOP
			}
			self.handleParse(self, data)
		}
	}
}
