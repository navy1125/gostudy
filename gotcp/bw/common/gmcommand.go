package Cmd

import ()

type StGmUserCmd struct {
	StNullUserCmd
}

const REQUEST_LOGIN_GM_USERCMD_PARA_G = 1

type StRequestLoginGmUserCmd struct {
	StGmUserCmd
	Name     [MAX_NAMESIZE]byte
	Password [MAX_PASSWORD]byte
	Version  uint32
}

func NewStRequestLoginGmUserCmd() *StRequestLoginGmUserCmd {
	cmd := &StRequestLoginGmUserCmd{}
	cmd.ByCmd = GM_USERCMD
	cmd.ByParam = REQUEST_LOGIN_GM_USERCMD_PARA_G
	return cmd
}

const RETURN_LOGIN_GM_USERCMD_PARA_S = 2

type StReturnLoginGmUserCmd struct {
	StGmUserCmd
	Retcode   byte
	Pri       uint32
	QMaxNum   uint16
	AutoRecv  byte
	WorkState byte
	WinNum    uint16
}

func NewStReturnLoginGmUserCmd() *StReturnLoginGmUserCmd {
	cmd := &StReturnLoginGmUserCmd{}
	cmd.ByCmd = GM_USERCMD
	cmd.ByParam = RETURN_LOGIN_GM_USERCMD_PARA_S
	return cmd
}
