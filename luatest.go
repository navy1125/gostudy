package main

import (
	"fmt"
	"github.com/stevedonovan/luar"
	"reflect"
	"strings"
)

const test = `
for i = 1,10 do
    Print(MSG,i)
end
`

type MyStruct struct {
	Name string
	Age  int
}

func TestCall(v interface{}) {
	fmt.Println("aaaaaaaaaaaaaaaaaaaaaaa", reflect.TypeOf(v))
}

const code = `
Print(#M)
Print(M.one)
Print 'pairs over Go maps'
for k,v in pairs(M) do
    Print(k,v)
end
Print 'ipairs over Go slices'
for i,v in ipairs(S) do
    Print(i,v)
end
ST.Name="dddddddddddddd"
Print(ST.Name)
testcall({a=1,b=2})
`

func main() {
	whj := "wang::hai::jun"
	fmt.Println(strings.Split(whj, "::"))
	TestCall("111")
	L := luar.Init()
	defer L.Close()
	M := luar.Map{
		"one":   "ein",
		"two":   "zwei",
		"three": "drei",
	}

	S := []string{"alfred", "alice", "bob", "frodo"}

	ST := &MyStruct{"Dolly", 46}

	luar.Register(L, "", luar.Map{
		"Print":    fmt.Println,
		"testcall": TestCall,
		"print":    fmt.Println,
		"MSG":      "hello", // can also register constants
		"M":        M,
		"S":        S,
		"ST":       ST,
	})

	//L.DoString(test)
	L.DoString(code)
	L.GetGlobal("print")
	print := luar.NewLuaObject(L, -1)
	print.Call("one two", 12)

	L.GetGlobal("package")
	pack := luar.NewLuaObject(L, -1)
	fmt.Println(pack.Get("path"))

	lcopy := luar.NewLuaObjectFromValue

	gsub := luar.NewLuaObjectFromName(L, "string.gsub")

	rmap := lcopy(L, luar.Map{
		"NAME": "Dolly",
		"HOME": "where you belong",
	})

	res, err := gsub.Call("hello $NAME go $HOME", "%$(%u+)", rmap)
	if res == nil {
		fmt.Println("error", err)
	} else {
		fmt.Println("\033[0;31mresult\033[0m", res)
	}

}
