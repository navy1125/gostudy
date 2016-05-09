package main

import (
	"code.google.com/p/goprotobuf/proto"
	"git.code4.in/mobilegameserver/platcommon"

	"fmt"
)

func main() {
	var b []byte
	fmt.Println(len(b))
	nmd := &Pmd.ForwardNullUserPmd_CS{}
	nmd1 := &Pmd.ForwardNullUserPmd_CS{}
	nmd2 := &Pmd.ForwardNullUserPmd_CS{}
	cmd3 := &Pmd.RequestCloseNullUserPmd_CS{}
	cmd4 := &Pmd.RequestCloseNullUserPmd_CS{}
	cmd3.Reason = proto.String("2222")
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
