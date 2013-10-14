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
