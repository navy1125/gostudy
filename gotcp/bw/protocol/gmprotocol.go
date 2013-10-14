package GM

import (
	"../common"
	"reflect"
)

const MAX_IP_LENGTH = 16
const LOGIN_GMCMD = 1

const REQUEST_LOGIN_LOGIN_GMCMD_PARA_C = 1

type stRequestLoginLoginGMCmd struct {
	t_NullCmd
	zoneid   uint32
	zonename [MAX_NAMESIZE]byte
	port     uint16
	strIP    [MAX_IP_LENGTH]byte
}

func (self *stRequestLoginLoginGMCmd) size() int {
	return int(reflect.TypeOf(*self).Size())
}

func NewstRequestLoginLoginGMCmd() *stRequestLoginLoginGMCmd {
	cmd := &stRequestLoginLoginGMCmd{}
	cmd.byCmd = LOGIN_GMCMD
	cmd.byParam = REQUEST_LOGIN_LOGIN_GMCMD_PARA_C
	return cmd
}
