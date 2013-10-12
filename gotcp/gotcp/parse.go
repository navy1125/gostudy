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
func handleReadFunDefault(task *Task) ([]byte, error) {
	var length uint32
	if err := binary.Read(task.Conn, binary.BigEndian, &length); err != nil {
		return nil, err
	}
	buf := make([]byte, length)
	_, err := io.ReadFull(task.Conn, buf)
	return buf, err
}
func handleWriteFunDefault(task *Task, data []byte) error {
	err := binary.Write(task.Conn, binary.BigEndian, uint32(len(data)))
	if err != nil {
		return err
	}
	_, err = task.Conn.Write(data)
	return err
}
func handleParseDefault(task *Task, data []byte) error {
	switch {
	case string(data) == "tick":
		task.Debug("time tick return")
		handleHeartBeatReturnDefault(task)
		task.HeartBeatReturn = true
	case string(data) == "return tick":
		task.HeartBeatReturn = true
	}
	return nil
}
