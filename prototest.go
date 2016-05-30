package main

import (
	"code.google.com/p/goprotobuf/proto"
	"git.code4.in/mobilegameserver/platcommon"
	"github.com/dedis/protobuf"
	//"github.com/golang/protobuf/ptypes"

	"fmt"
	"reflect"
)

func main() {
	m := make(map[int]string)
	m[1] = "wabghaijun"
	a := 1
	fmt.Println(protobuf.Encode(&a))
	mset := make(map[int32]proto.Extension)
	//mset[1] = proto.Extension{enc: []byte("sss")}
	//umset, err := proto.MarshalMessageSet(mset)
	var b []byte
	fmt.Println(len(b))
	nmd := &Pmd.ForwardNullUserPmd_CS{}
	nmd1 := &Pmd.ForwardNullUserPmd_CS{}
	nmd2 := &Pmd.ForwardNullUserPmd_CS{}
	cmd3 := &Pmd.RequestCloseNullUserPmd_CS{}
	cmd4 := &Pmd.RequestCloseNullUserPmd_CS{}
	cmd3.Reason = proto.String("2222")
	fmt.Println(proto.GetProperties(reflect.TypeOf(cmd3).Elem()))
	cmd3byte, err1 := proto.Marshal(cmd3)
	if err1 != nil {
		fmt.Println("xxxxxxxxxxx", err1)
	}
	fmt.Println(proto.Unmarshal(cmd3byte, cmd3))
	fmt.Println(mset)
	cmd3test := proto.MarshalTextString(cmd3)
	fmt.Println(cmd3test)
	fmt.Println(proto.UnmarshalText(cmd3test, nmd))
	//nmd.Prototype = proto.Uint64(2)
	nmd.ByCmd = proto.Uint32(0)
	//nmd.ByParam = proto.Uint32(0)
	//nmd.ByCmd = append(nmd.ByCmd, 0)
	//nmd.ByParam = append(nmd.ByParam, 0)
	sendbuf := proto.NewBuffer(nil)
	err := sendbuf.Marshal(nmd)
	if err != nil {
		fmt.Println("1", err)
	}
	nmd.Data = sendbuf.Bytes()
	fmt.Println(nmd, proto.Size(nmd), len(sendbuf.Bytes()))
	fmt.Println(len(sendbuf.Bytes()), sendbuf.Bytes())
	//data := sendbuf.Bytes()
	err = sendbuf.Marshal(cmd3)
	if err != nil {
		fmt.Println("2", err)
	}
	fmt.Println(len(sendbuf.Bytes()), sendbuf.Bytes())
	data := sendbuf.Bytes()
	fmt.Println(len(data), data)
	//data = append(data, byte(1))
	databuf := proto.NewBuffer(data)
	err = databuf.Unmarshal(nmd1)
	if err != nil {
		fmt.Println("3", err)
	}
	//err = databuf.Unmarshal(nmd2)
	err = proto.Unmarshal(data[:2], nmd2)
	if err != nil {
		fmt.Println("4", err)
	}
	err = proto.Unmarshal(data[2:], cmd4)
	//err = databuf.Unmarshal(cmd4)
	if err != nil {
		fmt.Println("5", err)
	}
	fmt.Println(nmd, proto.Size(nmd))
	fmt.Println(nmd1)
	fmt.Println(nmd2)
	fmt.Println(cmd4)
}
