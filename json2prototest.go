package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"code.google.com/p/goprotobuf/proto"
)

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
