// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	//"git.code4.in/logging"
	"runtime"
)

func main() {
	fmt.Println("aaaa:", string(runtime.CPUProfile()))
	runtime.Breakpoint()
}
