package main

import (
	"fmt"
	"github.com/hit9/reuseport"
	"net/http"
	"os"
)

func main() {
	ln, err := reuseport.Listener("tcp", ":8085")
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	go udp()
	server := &http.Server{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from process %d\n", os.Getpid())
	})
	panic(server.Serve(ln))
}

func udp() {
	conn, err := reuseport.PacketConn("udp", ":8085")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	buf := make([]byte, 2048)
	for {
		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			break
		}
		fmt.Printf("Got message: %s\n", string(buf), n, addr)
	}
}
