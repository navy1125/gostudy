package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"net/url"
)

func main() {
	//resp, err := http.Get("http://www.bwgame.com.cn")
	//resp, err := http.Get("http://127.0.0.1:8080/client.go")
	//resp, err := http.Get("http://127.0.0.1:12346/hello")
	//resp, err := http.PostForm("http://127.0.0.1:12346/hello",url.Values{"name":{"whj"},"age":{"12"}})
	resp, err := http.Get("http://112.65.197.72:8080/hello")

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
	http.SetCookie(w, cookie)
	//resp.Body.Close()
}
