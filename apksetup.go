package main

import (
	"bytes"
	"code.google.com/p/go.net/websocket"
	"flag"
	"fmt"
	iconv "github.com/hwch/iconv"
	"github.com/navy1125/config"
	"io"
	"log"
	"net/http"
	"os/exec"
	"sync"
)

var (
	logNutex  sync.Mutex
	setupMap  map[*websocket.Conn]*websocket.Conn
	converter *iconv.Converter
	//utf8
)

func SetupServer(ws *websocket.Conn) {
	var message string
	//world := createWorld()

	for {
		err := websocket.Message.Receive(ws, &message)
		if err != nil {
			delete(setupMap, ws)
			log.Print("Receive error - stopping worker: ", err)
			break
		}
		if message == "setup apk" {
			ws.Write([]byte("start setup apk now..."))
			if len(setupMap) == 0 {
				setupMap[ws] = ws
				//execBat("publish_android.bat", ws)
				//execBat("copy_apk.bat", ws)
				execBat(config.GetConfigStr("bat_apk"), ws)
				for k, _ := range setupMap {
					k.Write([]byte("setup finish apk"))
					k.Close()
					delete(setupMap, k)
				}
			} else {
				setupMap[ws] = ws
			}
		}
		if message == "setup win" {
			ws.Write([]byte("start setup win now..."))
			if len(setupMap) == 0 {
				setupMap[ws] = ws
				//execBat("publish_android.bat", ws)
				//execBat("copy_apk.bat", ws)
				execBat(config.GetConfigStr("bat_win"), ws)
				for k, _ := range setupMap {
					k.Write([]byte("setup finish win"))
					k.Close()
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
		//k.Write(b)
		out := make([]byte, len(b)*2)
		if l, err := converter.CodeConvertFunc(b, out); err == nil && l > 0 {
			//k.Write([]byte("wanghaijun"))
			k.Write([]byte(out))
		}
	}
}
func execBat(bat string, ws *websocket.Conn) {
	cmd := exec.Command(bat)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		Broadcask([]byte(err.Error()))
		fmt.Print(string([]byte(err.Error())))
	}
	stdin, err := cmd.StderrPipe()
	if err != nil {
		Broadcask([]byte(err.Error()))
		fmt.Print(string([]byte(err.Error())))
	}
	if err := cmd.Start(); err != nil {
		Broadcask([]byte(err.Error()))
		fmt.Print(string([]byte(err.Error())))
	}
	go outputFunc(stdout, ws)
	go outputFunc(stdin, ws)
	cmd.Wait()
}
func downloadWinFunc(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, config.GetConfigStr("download_win"), 303)
}
func downloadApkFunc(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, config.GetConfigStr("download_apk"), 303)
}
func outputFunc(r io.ReadCloser, w io.Writer) {
	b := bytes.NewBuffer(make([]byte, 1024))
	for {
		n, err := r.Read(b.Bytes())
		if n > 0 {
			logNutex.Lock()
			Broadcask(b.Bytes()[0:n])
			fmt.Print(string(b.Bytes()[0:n]))
			logNutex.Unlock()
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			Broadcask([]byte(err.Error()))
			fmt.Print(string([]byte(err.Error())))
			break
		}
	}
}
func main() {
	flag.Parse()
	config.SetConfig("config", *flag.String("config", "config.xml", "config xml file for start"))
	config.SetConfig("logfilename", *flag.String("logfilename", "/log/logfilename.log", "log file name"))
	config.SetConfig("deamon", *flag.String("deamon", "false", "need run as demo"))
	config.SetConfig("port", *flag.String("port", "8000", "http port "))
	config.SetConfig("log", *flag.String("log", "debug", "logger level "))
	config.SetConfig("loginServerList", *flag.String("loginServerList", "loginServerList.xml", "server list config"))
	config.LoadFromFile(config.GetConfigStr("config"), "global")
	if err := config.LoadFromFile(config.GetConfigStr("config"), "ApkServer"); err != nil {
		fmt.Println(err)
		return
	}
	var err error
	converter, err = iconv.NewCoder(iconv.GBK2312_UTF8_IDX)
	if err != nil {
		fmt.Println("iconv err:", err)
		return
	}
	setupMap = make(map[*websocket.Conn]*websocket.Conn)
	http.Handle("/ws", websocket.Handler(SetupServer))
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/download_win", downloadWinFunc)
	http.HandleFunc("/download_apk", downloadApkFunc)
	//err := http.ListenAndServe("180.168.197.87:18080", nil)
	err = http.ListenAndServe(config.GetConfigStr("ip")+":"+config.GetConfigStr("port"), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
