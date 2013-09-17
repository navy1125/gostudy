package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//resp, err := http.Get("http://www.bwgame.com.cn")
	resp, err := http.Get("http://127.0.0.1:8080/client.go")
	if err != nil {
		log.Fatal("http.Get Err:", err)
	}
	text, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("httpioutil.ReadAll Err:", err)
	}
	fmt.Printf("%s", text)
	//resp.Body.Close()
}
