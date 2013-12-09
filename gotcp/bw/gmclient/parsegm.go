package main

import (
	"../../gotcp"
	"../common"
	"fmt"
)

var (
	gmHandleMessageMap gotcp.HanldeMessageMap
)

func init() {
	RegisterMessage(Cmd.TIME_USERCMD, Cmd.GAMETIME_TIMER_USERCMD_PARA, parseStGameTimeTimerUserCmd)
	RegisterMessage(Cmd.TIME_USERCMD, Cmd.REQUESTUSERGAMETIME_TIMER_USERCMD_PARA, parseStRequestUserGameTimeTimerUserCmd)

}

func RegisterMessage(byCmd, byParam byte, fun gotcp.HandleMessageFunc) {
	fmt.Println(byCmd, byParam)
	gmHandleMessageMap[byCmd][byParam] = fun
}

func parseStGameTimeTimerUserCmd(task *gotcp.Task, data []byte) {
	task.Debug("heartBeat")
	cmd := Cmd.NewStGameTimeTimerUserCmd()
	task.SendCmd(cmd)
}

func parseStRequestUserGameTimeTimerUserCmd(task *gotcp.Task, data []byte) {
	cmd := Cmd.NewStRequestUserGameTimeTimerUserCmd()
	err := task.GetCmd(data, cmd)
	if err != nil {
		task.Error(err.Error())
		return
	}
	task.Debug("parseStRequestUserGameTimeTimerUserCmd:%d,%d", cmd.ByCmd, cmd.ByParam)
	task.SendCmd(cmd)
}
