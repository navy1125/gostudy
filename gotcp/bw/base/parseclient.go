package base

import (
	"../../gotcp"
	"../common"
	"encoding/binary"
	"fmt"
	"io"
)

func HandleHeartBeatRequestBw(task *gotcp.Task) {
	cmd := Cmd.NewT_ClientNullCmd()
	task.SendCmd(cmd)
}
func HandleHeartBeatReturnBw(task *gotcp.Task) {
	HandleWriteFunBw(task, []byte("return tick"))
}
func HandleReadFunBw(task *gotcp.Task) ([]gotcp.ReadData, error) {
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
func HandleWriteFunBw(task *gotcp.Task, data []byte) error {
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
func HandleParseBw(task *gotcp.Task, data []byte) error {
	fmt.Println(data[0], data[1])
	switch byCmd := data[0]; byCmd {
	case Cmd.CMD_NULL:
		switch byParam := data[1]; byParam {
		case Cmd.SERVER_PARA_NULL:
			HandleWriteFunBw(task, data)
			task.HeartBeatReturn = true
		case Cmd.CLIENT_PARA_NULL:
			task.HeartBeatReturn = true
		}
	case 2:
		task.HeartBeatReturn = true
	}
	return nil
}
