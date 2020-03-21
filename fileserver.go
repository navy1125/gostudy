package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	_, err := os.Stat("hellow")
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Fatal(http.ListenAndServe("127.0.0.1:8087", http.FileServer(http.Dir("/Users/whj/mygo/src/github.com/nebula-chat/nebula-chat.github.io/"))))
	//http.Handle("/gostudy/", http.StripPrefix("/gostudy/", http.FileServer(http.Dir("D:\\work\\gostudy"))))
	http.Handle("/bw/", http.StripPrefix("/bw/", http.FileServer(http.Dir("/home/whj/gogos/gameauth/game/bw"))))
	http.HandleFunc("/post", postfunc)
	http.HandleFunc("/upload", uploadfunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func uploadfunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Write([]byte("hello world from host1" + r.Method + "\n"))
		return
	}
	r.ParseMultipartForm(1 << 20)
	//w.Write([]byte("hello world from host1"))
	//w.Write([]byte(r.Header.Get("Content-Type")))
	f, fh, err := r.FormFile("upload")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	filename := fh.Filename
	if strings.Contains(fh.Filename, ":") {
		strtmp := strings.Split(fh.Filename, ":")
		filename = strtmp[len(strtmp)-1]
		fmt.Fprintf(w, "\nchange file name:%s\n", filename)
	} else if strings.Contains(fh.Filename, "/") {
		strtmp := strings.Split(fh.Filename, "/")
		filename = strtmp[len(strtmp)-1]
		fmt.Fprintf(w, "\nchange file name:%s\n", filename)
	}
	fmt.Fprintf(w, "\n%+v\n%s\n", fh.Header, filename)
	file, err := os.OpenFile("/tmp/"+filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	io.Copy(file, f)
}
func postfunc(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("/tmp/upload.html")
	if err != nil {
		fmt.Println("open upload.html err:", err)
		return
	}
	text, _ := ioutil.ReadAll(file)
	w.Write([]byte(text))
}
