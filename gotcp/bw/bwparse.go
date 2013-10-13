package main

import (
	"../gotcp"
	"encoding/binary"
	"fmt"
	"io"
)

func handleHeartBeatRequestBw(task *gotcp.Task) {
	cmd := Newt_ClientNullCmd()
	task.SendCmd(cmd)
}
func handleHeartBeatReturnBw(task *gotcp.Task) {
	handleWriteFunBw(task, []byte("return tick"))
}
func handleReadFunBw(task *gotcp.Task) ([]gotcp.ReadData, error) {
	var l uint32
	if err := binary.Read(task.Conn, binary.LittleEndian, &l); err != nil {
		return nil, err
	}
	fmt.Println("readlen:", l)
	data := make([]gotcp.ReadData, 1)
	data[0].Length = int(l)
	data[0].Data = make([]byte, l)
	_, err := io.ReadFull(task.Conn, data[0].Data)
	return data, err
}
func handleWriteFunBw(task *gotcp.Task, data []byte) error {
	fmt.Println("write data:", data[0], data[1])
	err := binary.Write(task.Conn, binary.LittleEndian, uint16(len(data)))
	if err != nil {
		return err
	}
	err = binary.Write(task.Conn, binary.LittleEndian, uint16(0))
	if err != nil {
		return err
	}
	_, err = task.Conn.Write(data)
	return err
}
func handleParseBw(task *gotcp.Task, data []byte) error {
	fmt.Println(data[0], data[1])
	switch byCmd := data[0]; byCmd {
	case CMD_NULL:
		switch byParam := data[1]; byParam {
		case SERVER_PARA_NULL:
			handleWriteFunBw(task, data)
			task.HeartBeatReturn = true
		case CLIENT_PARA_NULL:
			task.HeartBeatReturn = true
		}
	case 2:
		task.HeartBeatReturn = true
	}
	return nil
}
