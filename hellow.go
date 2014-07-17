// You can edit this code!
// Click here and start typing.
package main

import (
	"flag"
	"fmt"
	iconv "github.com/hwch/iconv"
	"os"
	"reflect"
	"runtime"
	"strconv"
	//"rand"
	//"C"
)

var (
	TestMap map[string]string
)

type Struct struct {
	TestI int
}

func (self *Struct) Print() {
	fmt.Println("aaaaaaaaaaaa:", self.TestI)
}
func (self *Struct) Print1() {
	fmt.Println("xxxxxxxxxxxxx:")
}

type StructInterface interface {
	Print()
}

func TestStruct(st StructInterface) {
	s := st.(*Struct)
	s.Print()
}

func TestPanic() {
	panic("ddddddddddddd")
}

func parseGatewayTaskNullGateCmd(v interface{}) bool {
	call := reflect.ValueOf(v).MethodByName("Print1")
	if call.IsValid() == true {
		call.Call([]reflect.Value{})
		fmt.Println(reflect.TypeOf(v))
		return true
	}
	return true
}
func main() {
	st := &Struct{}
	parseGatewayTaskNullGateCmd(st)
	TestStruct(st)
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	stra := "a"
	str := fmt.Sprintf("%03o,%d", stra[0], stra[0])
	aa, _ := strconv.ParseInt("141", 8, 32)
	fmt.Println("str:", str, aa, fmt.Sprintf("%c", aa))
	converter, err := iconv.NewCoder(iconv.GBK2312_UTF8_IDX)
	if err != nil {
		fmt.Println("iconv err:", err)
		return
	}
	out := make([]byte, 1024)
	if l, err := converter.CodeConvertFunc([]byte("王海军"), out); err == nil && l > 0 {
		fmt.Println("converter err:", err)
	}
	fmt.Println(string(out))
	return
	fmt.Println("USER", os.Getenv("PATH"))
	defer fmt.Println("Hello, 世界")
	fmt.Printf("text:%d,%d", 100, 100)
	a := 10
	var b = 10
	var c int = 10
	fmt.Println("aaaaaaaaa", a, b, c)
	var host = flag.String("host", "", "Server listen host, default 0.0.0.0")
	fmt.Println("flag test", host)
	fmt.Println("cpunum:", runtime.NumCPU())
	//fmt.Println("aaaaaaaaa",a,b,c,rand.Intn(2))
	//fmt.Println("C.random", int(C.random()))
	TestMap = make(map[string]string)
	TestMap["whj"] = "wanghaijun@ztgame.com"
	fmt.Println(TestMap)
	fmt.Println(TestMap["whj"])
	TestPanic()
}
