package w32api

/*
#include <stdio.h>
*/
import "C"

func PrintHello() {
	C.puts(C.CString("Hello, world\n"))

}
