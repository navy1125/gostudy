package Cmd

import ()

//max name size
const MAX_NAMESIZE = 32
const MAX_PASSWORD = 16

type T_NullCmd struct {
	ByCmd   byte
	ByParam byte
}
type StNullUserCmd struct {
	T_NullCmd
	DwTimestamp uint32
}

const CMD_NULL = 0
const SERVER_PARA_NULL = 0

type T_ServerNullCmd struct {
	T_NullCmd
}

func NewT_ServerNullCmd() *T_ServerNullCmd {
	cmd := &T_ServerNullCmd{}
	cmd.ByCmd = CMD_NULL
	cmd.ByParam = SERVER_PARA_NULL
	return cmd
}

const CLIENT_PARA_NULL = 1

type T_ClientNullCmd struct {
	T_NullCmd
}

func NewT_ClientNullCmd() *T_ClientNullCmd {
	cmd := &T_ClientNullCmd{}
	cmd.ByCmd = CMD_NULL
	cmd.ByParam = CLIENT_PARA_NULL
	return cmd
}

const GM_USERCMD = 62

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
