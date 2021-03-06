package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.RemoteAddr)
	//io.WriteString(w, req.PostFormValue("name"))
	io.WriteString(w, req.URL.String())
	//io.WriteString(w, req.Form.Encode())
	//io.WriteString(w, req.PostForm.Encode()+"\n")
	if cookstr, err := req.Cookie("testcookiename"); err == nil {
		io.WriteString(w, " cookie testcookiename:"+cookstr.String()+"\n")
	}

}

func main() {
	//cookstr, _ := req.Cookie("testcookiename")
	//io.WriteString(w, "cookie testcookiename:"+cookstr.String()+"\n")
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(":12346", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
