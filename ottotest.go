package main

import (
	"fmt"
	"github.com/robertkrimen/otto"
	//"github.com/robertkrimen/otto/parser"
	"reflect"
	"time"
)

func main() {
	fmt.Println("aaaaa")
	vm := otto.New()
	vm.Run(`
	    abc = 2 + 2;
			console.log("The value of abc is " + abc); // 4
			`)
	value, _ := vm.Get("abc")
	value1, _ := value.ToInteger()
	fmt.Println(value, reflect.TypeOf(value).String(), value1)
	vm.Set("def", 11)
	vm.Set("PrintGo", 11)
	vm.Run(`
	console.log("The value of def is " + def);
	// The value of def is 11
	`)
	_, err := vm.Run(`
		function GoPrint(para1,para2){
		console.log("GoPrint:"+para1+para2);
		}
		console.log("GoPrint:");
		GoPrint("wwwwwwwwwww");
		`)
	vm.Set("xyzzy", "Nothing happens.")
	vm.Run(`
	console.log(xyzzy.length); // 16
	`)
	fmt.Println("err")
	if err != nil {
		fmt.Println("aaaaaa:", err.Error())
	}
	vm.Run(`
	console.log(xyzzy.length); // 16
	`)
	_, err = vm.Call("GoPrint", nil, "GoPrint", "test2")
	if err != nil {
		fmt.Println("val", err)
	}
	//program, err := parser.ParseFile(nil, filename, src, 0)
	//program.
	//parse ,_ := "js/apksetup.js"
	vm.Set("LoadJsFile", func(call otto.FunctionCall) otto.Value {
		arg0, _ := call.Argument(0).ToString()
		ret := InitJs(vm, arg0)
		val, _ := otto.ToValue(ret)
		return val
	})
	InitJs(vm, "ottotest.js")
	_, err = vm.Call("init", nil, "dir")
	if err != nil {
		fmt.Println(err)
	}
	//object, _ := vm.Object(`({ xyzzy: "Nothing happens." })`)
	object, err := vm.Object(`aaa=1`)
	if err != nil {
		fmt.Println("aaaaaaaaaaaaaaaaaaaa")
		fmt.Println(object.Value().ToString())
		fmt.Println(err)
		fmt.Println("bbbbbbbbbbbbbbbbbbbb")
	} else {
		object.Set("volume", "11")
		fmt.Println(object.Get("volume"))
		fmt.Println(object.Get("xyzzy"))
	}
	vm.Interrupt = make(chan func(), 1)
	go func() {
		for true {
			time.Sleep(time.Second)
			vm.Interrupt <- func() {
				fmt.Println("timeout")
			}
		}
	}()
	for true {
		vm.Call("DeadLoop", nil)
		time.Sleep(time.Second)
	}
}

func InitJs(vm *otto.Otto, filename string) bool {
	if vm == nil {
		vm = otto.New()
	}
	script, err := vm.Compile(filename, nil)
	if err != nil {
		fmt.Println("InitJS err:", err.Error())
		return false
	}
	_, err = vm.Run(script)
	if err != nil {
		fmt.Println("InitJS err:", err.Error())
		return false
	}
	return true
}
