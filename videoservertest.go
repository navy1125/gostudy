package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
//redis_handle *redis.Client
)

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

func main() {
	http.Handle("/video/", http.StripPrefix("/video/", http.FileServer(http.Dir("video/"))))
	http.HandleFunc("/upload", uploadfunc)
	log.Fatal(http.ListenAndServe("127.0.0.1:12345", nil))
}
