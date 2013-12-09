package gotcp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

type HandleMessageFunc func(task *Task, data []byte)

type HanldeMessageMap [255][256]HandleMessageFunc

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
	handleParse     func(task *Task, data []byte) bool
	handleHeartBeat func(task *Task)
	heartBeatTime   time.Duration
	HeartBeatReturn bool
	handleMessage   *HanldeMessageMap
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
func (self *Task) SetHandleMessage(handle *HanldeMessageMap) {
	self.handleMessage = handle
}
func (self *Task) ParseMessage(data []byte) {
	if self.handleMessage == nil {
		self.Error("no handleMessage:%d,%d", data[0], data[1])
		return
	}
	if len(data) < 2 {
		self.Error("message can not less then 2 bytes")
		return
	}
	switch fun := self.handleMessage[data[0]][data[1]]; fun {
	case nil:
		self.Error("no parse func for message2:%d,%d", data[0], data[1])
		return
	default:
		fun(self, data)
	}
}
func (self *Task) SetHandleReadFun(fun func(task *Task) ([]ReadData, error)) {
	self.handleReadFun = fun
}
func (self *Task) SetHandleWriteFun(fun func(task *Task, data []byte) error) {
	self.handleWriteFun = fun
}
func (self *Task) SetHandleParseFun(fun func(task *Task, data []byte) bool) {
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
		self.Error("SendCmd err:", err.Error())
		return
	}
	self.handleWriteFun(self, buf.Bytes())

}

func (self *Task) GetCmd(data []byte, v interface{}) error {
	buf := bytes.NewBuffer(data)
	if err := binary.Read(buf, binary.BigEndian, v); err != nil {
		self.Error("GetCmd err:", err.Error())
		return err
	}
	return nil
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
				if ok := self.handleParse(self, data.Data); !ok {
					self.ParseMessage(data.Data)
				}
			}
		}
	}
}
