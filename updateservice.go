package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func HandleRestart(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	os.Exit(0)
}

func main() {
	var port int
	flag.IntVar(&port, "port", 18081, "listen port")
	flag.Parse()
	http.HandleFunc("/restart", HandleRestart)
	pwd, _ := os.Getwd()
	exename := os.Args[0]
	exename = strings.Replace(exename, "./", "/", 1)
	params := ""
	for _, param := range os.Args[1:] {
		params = params + " " + param
	}
	cmd := exec.Command(pwd+exename, "-port=18082")
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
	}

	//if err := cmd.Wait(); err != nil {
	//	fmt.Println(err)
	//}

	fmt.Println("ssssssssssss")
	err := http.ListenAndServe("127.0.0.1"+":"+strconv.Itoa(port), nil)
	if err != nil {
		fmt.Println(err)
	}
}
