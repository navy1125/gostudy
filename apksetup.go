package main

import (
	"bytes"
	"code.google.com/p/go.net/websocket"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
)

func RugServer(ws *websocket.Conn) {
	var message string
	//world := createWorld()

	fmt.Print("Launched worker\n")
	log.Print("Launched worker\n")

	for {
		err := websocket.Message.Receive(ws, &message)
		if err != nil {
			log.Print("Receive error - stopping worker: ", err)
			break
		}
		if message == "setup" {
			execBat("publish_android.bat", ws)
			execBat("copy_apk.bat", ws)
		}

		//err = websocket.JSON.Send(ws, world.Update(message))
		//if err != nil {
		//	log.Print("Send error - stopping worker: ", err)
		//	break
		//}
	}
}
func execBat(bat string, ws *websocket.Conn) {
	cmd := exec.Command(bat)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		ws.Write([]byte(err.Error()))
		fmt.Print([]byte(err.Error()))
	}
	stdin, err := cmd.StderrPipe()
	if err != nil {
		ws.Write([]byte(err.Error()))
		fmt.Print([]byte(err.Error()))
	}
	if err := cmd.Start(); err != nil {
		ws.Write([]byte(err.Error()))
		fmt.Print([]byte(err.Error()))
	}
	go outputFunc(stdout, ws)
	go outputFunc(stdin, ws)
}
func downloadFunc(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "proj.android\\bin\\hxsg-debug.apk", 303)
}
func outputFunc(r io.ReadCloser, w io.Writer) {
	b := bytes.NewBuffer(make([]byte, 1024))
	for {
		n, err := r.Read(b.Bytes())
		if n > 0 {
			w.Write(b.Bytes()[0:n])
			fmt.Print(b.Bytes()[0:n])
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			w.Write([]byte(err.Error()))
			fmt.Print([]byte(err.Error()))
			break
		}
	}
}
func main() {
	http.Handle("/ws", websocket.Handler(RugServer))
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/download", downloadFunc)
	err := http.ListenAndServe("180.168.197.87:18080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
