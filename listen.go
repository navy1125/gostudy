package main

import (
	"fmt"
	"github.com/xuyu/logging"
	"io"
	"log"
	"net/http"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.RemoteAddr)
	io.WriteString(w, req.PostFormValue("name"))
	io.WriteString(w, req.URL.String())
	io.WriteString(w, req.Form.Encode())
	io.WriteString(w, req.PostForm.Encode()+"\n")
	if cookstr, err := req.Cookie("testcookiename"); err == nil {
		io.WriteString(w, " cookie testcookiename:"+cookstr.String()+"\n")
	}

}

func main() {
	logger, err := logging.NewRotationLogger("/tmp/listenserver.log", "060102-15")
	//cookstr, _ := req.Cookie("testcookiename")
	//io.WriteString(w, "cookie testcookiename:"+cookstr.String()+"\n")
	if err != nil {
		fmt.Println(err)
		return
	}
	logging.SetDefaultLogger(logger)
	logging.SetPrefix("LT")
	http.HandleFunc("/hello", HelloServer)
	err = http.ListenAndServe(":12346", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
