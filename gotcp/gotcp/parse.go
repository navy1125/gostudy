package gotcp

import (
	"encoding/binary"
	//"fmt"
	"io"
)

func handleHeartBeatRequestDefault(task *Task) {
	handleWriteFunDefault(task, []byte("tick"))
}
func handleHeartBeatReturnDefault(task *Task) {
	handleWriteFunDefault(task, []byte("return tick"))
}
func handleReadFunDefault(task *Task) ([]ReadData, error) {
	var l uint32
	if err := binary.Read(task.Conn, binary.BigEndian, &l); err != nil {
		return nil, err
	}
	data := make([]ReadData, 1)
	data[0].Length = int(l)
	data[0].Data = make([]byte, l)
	_, err := io.ReadFull(task.Conn, data[0].Data)
	return data, err
}
func handleWriteFunDefault(task *Task, data []byte) error {
	err := binary.Write(task.Conn, binary.BigEndian, uint32(len(data)))
	if err != nil {
		return err
	}
	_, err = task.Conn.Write(data)
	return err
}
func handleParseDefault(task *Task, data []byte) bool {
	switch {
	case string(data) == "tick":
		task.Debug("time tick return")
		handleHeartBeatReturnDefault(task)
		task.HeartBeatReturn = true
	case string(data) == "return tick":
		task.HeartBeatReturn = true
	default:
		return false
	}
	return true
}
