package main

import (
	"github.com/simonz05/godis/redis"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	redis_handle *redis.Client
)

func main() {
	redis_handle = redis.New("tcp:112.65.197.72:6379", 0, "")
	redis_handle.Set("111", "111")
	//log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("D:\\work\\gostudy"))))
	//http.Handle("/gostudy/", http.StripPrefix("/gostudy/", http.FileServer(http.Dir("D:\\work\\gostudy"))))
	http.HandleFunc("/post", postfunc)
	http.HandleFunc("/get", getfunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func postfunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Write([]byte("hello world from host1" + r.Method + "\n"))
		return
	}
	uuid := r.PostFormValue("uuid")
	if uuid == "" {
		w.Write([]byte("uuid can not empty"))
		return
	}
	data := r.PostFormValue("data")
	if data != "" {
		if err := redis_handle.Set(uuid, data); err != nil {
			w.Write([]byte("save redis err"))
			return
		}
		w.Write([]byte("ok"))
		return
	}
	r.ParseMultipartForm(1 << 20)
	//w.Write([]byte("hello world from host1"))
	//w.Write([]byte(r.Header.Get("Content-Type")))
	f, _, err := r.FormFile("client_data")
	if err != nil {
		w.Write([]byte("save data not find client_data"))
		return
	}
	text, err := ioutil.ReadAll(f)
	if err != nil {
		w.Write([]byte("read client_data err"))
		return
	}
	if err := redis_handle.Set(uuid, text); err != nil {
		w.Write([]byte("save redis err"))
		return
	}
	w.Write([]byte("ok"))
	defer f.Close()

}
func getfunc(w http.ResponseWriter, r *http.Request) {
	uuid := r.FormValue("uuid")
	if uuid == "" {
		w.Write([]byte("uuid can not empty"))
		return
	}
	elem, err := redis_handle.Get(uuid)
	if err != nil {
		w.Write([]byte("redis find uuid err"))
		return
	}
	if len(elem.Bytes()) == 0 {
		w.Write([]byte("no data for uuid in redis"))
		return
	}

	w.Write(elem.Bytes())
}
