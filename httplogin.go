package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {

	paracheck := "http://127.0.0.1:7000/httplogin?game=101&zone=1&cmd=register-account"
	fmt.Println(paracheck)
	resp, err := http.Post(paracheck, "application/x-www-form-urlencoded", strings.NewReader(`{"do":"register-account","data":{"gameid":170}}`))
	//resp, err := http.Get(paracheck)
	if err != nil {
		log.Fatal("http.Get Err:", err)
	}
	defer resp.Body.Close()
	text, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("httpioutil.ReadAll Err:", err)
	}
	fmt.Println(resp.Header.Get("Server"))
	for key, value := range resp.Header {
		fmt.Println(key, value)
	}
	fmt.Printf("%s", text)
}
