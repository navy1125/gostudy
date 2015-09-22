package main

//import "github.com/aarzilli/golua/lua"
import (
	"fmt"
	"git.code4.in/mobilegameserver/golua/lua"
)

func main() {
	L := lua.NewState()
	L.OpenLibs()
	if err := L.DoFile("error.lua"); err != nil {
		panic(err)
	}
	var s string
	s = nil

	// 120 worked on my machine, might be different for you
	for i := 0; i < 120; i++ {
		top := L.GetTop()
		L.GetGlobal("doError")
		fmt.Println("CheckStack:", L.CheckStack(300))
		err := L.Call(0, lua.LUA_MULTRET)
		if err != nil {
			L.Remove(-1)
			L.Remove(-1)
			L.Remove(-1)
			L.Remove(-1)
			fmt.Println("error:", err)
			fmt.Println("stack length:", len(L.StackTrace()))
			//L.Pop(3)
			// fmt.Println("stack:", L.StackTrace())
		}
		L.SetTop(top)
	}
}
