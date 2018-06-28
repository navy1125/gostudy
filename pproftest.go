// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"runtime/pprof"
)

func main() {
	fm, err := os.OpenFile(".mem.out", os.O_RDWR|os.O_CREATE, 0644)
	if err == nil {
		pprof.WriteHeapProfile(fm)
	}
	fm.Close()
}
