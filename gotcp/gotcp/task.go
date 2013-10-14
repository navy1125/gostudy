package gotcp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	//"reflect"
	"time"
)

type ReadData struct {
	Length int
	Data   []byte
}
type Task struct {
	Entry
	Conn            *net.TCPConn
	in              chan []ReadData
	stop            chan bool
	handleReadFun   func(task *Task) ([]ReadData, error)
	handleWriteFun  func(task *Task, data []byte) error
	handleParse     func(task *Task, data []byte) error
	handleHeartBeat func(task *Task)
	heartBeatTime   time.Duration
	HeartBeatReturn bool
}

func NewTask(c *net.TCPConn, name string) *Task {
	task := &Task{
		Conn:            c,
		in:              make(chan []ReadData),
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
func (self *Task) SetHandleReadFun(fun func(task *Task) ([]ReadData, error)) {
	self.handleReadFun = fun
}
func (self *Task) SetHandleWriteFun(fun func(task *Task, data []byte) error) {
	self.handleWriteFun = fun
}
func (self *Task) SetHandleParseFun(fun func(task *Task, data []byte) error) {
	self.handleParse = fun
}
func (self *Task) SetHandleHeartBteaFun(fun func(task *Task), dur time.Duration) {
	self.handleHeartBeat = fun
	self.heartBeatTime = dur
}

/*
func (self *Task) SendCmd(ptr unsafe.Pointer, length int) {
	data := make([]byte, 0, 0)
	header := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	header.Data = (*reflect.SliceHeader)(ptr).Data
	header.Len = length
	header.Cap = length
	self.handleWriteFun(self, data)

}
//*/
func (self *Task) SendCmd(v interface{}) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, v); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len(buf.Bytes()))
	self.handleWriteFun(self, buf.Bytes())

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
		case datas, ok := <-self.in:
			if !ok {
				self.Debug("self.in chan err")
				break LOOP
			}
			for _, data := range datas {
				self.handleParse(self, data.Data)
			}
		}
	}
}