package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	fmt.Println(os.Getppid())
	fmt.Println(os.Args)
	if os.Getppid() != 1 {
		//判断当其是否是子进程，当父进程return之后，子进程会被 系统1 号进程接管
		filePath, _ := filepath.Abs(os.Args[0])
		//将命令行参数中执行文件路径转换成可用路径
		cmd := exec.Command(filePath, os.Args[1:]...)
		cmd.Start()
		return
	}
	http.ListenAndServe(":8000", nil)
}
