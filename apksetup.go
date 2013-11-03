package main

import (
	"bytes"
	"code.google.com/p/go.net/websocket"
	"io"
	"log"
	"net/http"
	"os/exec"
	"sync"
)

var (
	logNutex sync.Mutex
	setupMap map[*websocket.Conn]*websocket.Conn
)

func SetupServer(ws *websocket.Conn) {
	var message string
	//world := createWorld()

	for {
		err := websocket.Message.Receive(ws, &message)
		if err != nil {
			log.Print("Receive error - stopping worker: ", err)
			break
		}
		if message == "setup" {
			if len(setupMap) == 0 {
				setupMap[ws] = ws
				execBat("publish_android.bat", ws)
				execBat("copy_apk.bat", ws)
				for k, _ := range setupMap {
					delete(setupMap, k)
				}
			} else {
				setupMap[ws] = ws
			}
		}

		//err = websocket.JSON.Send(ws, world.Update(message))
		//if err != nil {
		//	log.Print("Send error - stopping worker: ", err)
		//	break
		//}
	}
}
func Broadcask(b []byte) {
	for k, _ := range setupMap {
		k.Write(b)
	}
}
func execBat(bat string, ws *websocket.Conn) {
	cmd := exec.Command(bat)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		Broadcask([]byte(err.Error()))
	}
	stdin, err := cmd.StderrPipe()
	if err != nil {
		Broadcask([]byte(err.Error()))
	}
	if err := cmd.Start(); err != nil {
		Broadcask([]byte(err.Error()))
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
			logNutex.Lock()
			w.Write(b.Bytes()[0:n])
			logNutex.Unlock()
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			w.Write([]byte(err.Error()))
			break
		}
	}
}
func main() {
	setupMap = make(map[*websocket.Conn]*websocket.Conn)
	http.Handle("/ws", websocket.Handler(SetupServer))
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/download", downloadFunc)
	err := http.ListenAndServe("180.168.197.87:18080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
