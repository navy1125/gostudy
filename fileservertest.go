package main

import (
	"fmt"
	"github.com/simonz05/godis/redis"
	"log"
	"net/http"
)

var (
	redis_handle *redis.Client
)

func main() {
	redis_handle := redis.New("tcp:112.65.197.72:6379", 0, "")
	//log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("D:\\work\\gostudy"))))
	//http.Handle("/gostudy/", http.StripPrefix("/gostudy/", http.FileServer(http.Dir("D:\\work\\gostudy"))))
	http.Handle("/", http.FileServer(http.Dir("e:\\tmp")))
	sm1 := http.NewServeMux()
	sm2 := http.NewServeMux()
	sm1.HandleFunc("/", hf1)
	sm2.HandleFunc("/", hf2)
	http.Handle("/test1/", sm1)
	http.Handle("/test2/", sm2)
	http.HandleFunc("/hijack", hf3)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hf1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world from host1"))
}

func hf2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world from host2"))
}
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not found :" + r.Host))
}
func hf3(w http.ResponseWriter, r *http.Request) {
	hj, ok := w.(http.Hijacker)
	if !ok {
		w.Write([]byte("webserver doesn't support hijacking"))
		return
	}
	w.Write([]byte("hijack :" + r.Host))
	return
	conn, bufrw, err := hj.Hijack()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	// Don't forget to close the connection:
	defer conn.Close()
	bufrw.WriteString("Now we're speaking raw TCP. Say hi: ")
	bufrw.Flush()
	s, err := bufrw.ReadString('\n')
	if err != nil {
		log.Printf("error reading string: %v", err)
		return
	}
	fmt.Fprintf(bufrw, "You said: %q\nBye.\n", s)
	bufrw.Flush()
}
