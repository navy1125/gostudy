package main

import ()

//max name size
const MAX_NAMESIZE = 32
const MAX_PASSWORD = 16

type t_NullCmd struct {
	byCmd   byte
	byParam byte
}
type stNullUserCmd struct {
	t_NullCmd
	dwTimestamp uint32
}

const CMD_NULL = 0
const SERVER_PARA_NULL = 0

type t_ServerNullCmd struct {
	t_NullCmd
}

func Newt_ServerNullCmd() *t_ServerNullCmd {
	cmd := &t_ServerNullCmd{}
	cmd.byCmd = CMD_NULL
	cmd.byParam = SERVER_PARA_NULL
	return cmd
}

const CLIENT_PARA_NULL = 1

type t_ClientNullCmd struct {
	t_NullCmd
}

func Newt_ClientNullCmd() *t_ClientNullCmd {
	cmd := &t_ClientNullCmd{}
	cmd.byCmd = CMD_NULL
	cmd.byParam = CLIENT_PARA_NULL
	return cmd
}

const GM_USERCMD = 62

type stGmUserCmd struct {
	stNullUserCmd
}

const REQUEST_LOGIN_GM_USERCMD_PARA_G = 1

type stRequestLoginGmUserCmd struct {
	stGmUserCmd
	name     [MAX_NAMESIZE]byte
	password [MAX_PASSWORD]byte
	version  uint32
}

func NewstRequestLoginGmUserCmd() *stRequestLoginGmUserCmd {
	cmd := &stRequestLoginGmUserCmd{}
	cmd.byCmd = GM_USERCMD
	cmd.byParam = REQUEST_LOGIN_GM_USERCMD_PARA_G
	return cmd
}

const RETURN_LOGIN_GM_USERCMD_PARA_S = 2

type stReturnLoginGmUserCmd struct {
	stGmUserCmd
	retcode   byte
	pri       uint32
	qMaxNum   uint16
	autoRecv  byte
	workState byte
	winNum    uint16
}

func NewstReturnLoginGmUserCmd() *stReturnLoginGmUserCmd {
	cmd := &stReturnLoginGmUserCmd{}
	cmd.byCmd = GM_USERCMD
	cmd.byParam = RETURN_LOGIN_GM_USERCMD_PARA_S
	return cmd
}
