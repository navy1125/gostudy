package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"code.google.com/p/goprotobuf/proto"
)

type AccountTokenVerifyLoginUserPmd struct {
	Account          *string `protobuf:"bytes,1,req,name=account" json:"account,omitempty"`
	Token            *string `protobuf:"bytes,2,req,name=token" json:"token,omitempty"`
	Version          *uint32 `protobuf:"varint,3,req,name=version" json:"version,omitempty"`
	Gameid           *uint32 `protobuf:"varint,4,opt,name=gameid" json:"gameid,omitempty"`
	Mid              *string `protobuf:"bytes,5,opt,name=mid" json:"mid,omitempty"`
	Platid           *uint32 `protobuf:"varint,6,opt,name=platid" json:"platid,omitempty"`
	Zoneid           *uint32 `protobuf:"varint,7,opt,name=zoneid" json:"zoneid,omitempty"`
	Gameversion      *uint32 `protobuf:"varint,8,opt,name=gameversion" json:"gameversion,omitempty"`
	Compress         *string `protobuf:"bytes,9,opt,name=compress" json:"compress,omitempty"`
	Encrypt          *string `protobuf:"bytes,10,opt,name=encrypt" json:"encrypt,omitempty"`
	Encryptkey       *string `protobuf:"bytes,11,opt,name=encryptkey" json:"encryptkey,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

var ()

func BuildProtoFromJson(typ reflect.Type, cmdjson string) []byte {
	proto_cmd := reflect.New(typ).Interface().(proto.Message)
	rawdata := []byte(cmdjson)
	json.Unmarshal(rawdata, proto_cmd)
	sendbuf := proto.NewBuffer(nil)
	sendbuf.Marshal(proto_cmd)
	return sendbuf.Bytes()
}

func BuildJsonFromProto(cmdname string, cmddata []byte) string {
	recvbuf := proto.NewBuffer(cmddata)
	recv := &AccountTokenVerifyLoginUserPmd_CS{} //难点,这里这个结构是不确定的,只能动态描述
	recvbuf.Unmarshal(recv)
	recv_json, _ := json.Marshal(recv)
	return string(recv_json)
}

func main() {
	account := "wangaijun"
	token := "Token"
	send := &AccountTokenVerifyLoginUserPmd_CS{ //难点,这里这个结构是不确定的,只能动态描述
		Account: &account,
		Token:   &token,
	}
	send_json, _ := json.Marshal(send)
	fmt.Println(string(send_json))
	bytes := BuildProtoFromJson(reflect.TypeOf(*send), string(send_json))

	//网络发送部分

	recv := bytes
	recv_json := BuildJsonFromProto("AccountTokenVerifyLoginUserPmd_CS", recv)
	fmt.Println("BuildJsonFromProto", string(recv_json))
}
