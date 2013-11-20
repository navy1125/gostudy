package main

import (
	"bytes"
	"code.google.com/p/go.net/websocket"
	"flag"
	"fmt"
	"github.com/navy1125/config"
	//iconv "github.com/navy1125/iconv"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"sync"
)

type GameTempateData struct {
	Name    string
	UrlName string
}

var (
	logNutex sync.Mutex
	setupMap map[*websocket.Conn]*websocket.Conn
	gtdMap   map[string]GameTempateData
	//converter *iconv.Converter
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
		fmt.Println(message)
		if strings.Contains(message, "setup apk") {
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
		if strings.Contains(message, "setup win") {
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
	/*
		out := make([]byte, len(b)*4)
		if l, err := converter.CodeConvertFunc(b, out); err == nil && l > 0 {
			for k, _ := range setupMap {
				//k.Write(b)
				//k.Write([]byte("wanghaijun"))
				k.Write([]byte(out))
			}
		}*/
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
	fmt.Println(r.URL.String())
	gameurl := strings.Split(r.URL.String(), "/")[1]
	if len(setupMap) != 0 {
		w.Write([]byte("有人正在做版本,请稍后再试"))
		defer r.Body.Close()
	} else {
		http.Redirect(w, r, strings.Replace(r.URL.String(), "download_win", "win", 1)+"/"+config.GetConfigStr("download_win"+"_"+gameurl+"_file"), 303)
	}
}
func downloadApkFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.String())
	gameurl := strings.Split(r.URL.String(), "/")[1]
	if len(setupMap) != 0 {
		w.Write([]byte("有人正在做版本,请稍后再试"))
		defer r.Body.Close()
	} else {
		http.Redirect(w, r, strings.Replace(r.URL.String(), "download_apk", "apk", 1)+"/"+config.GetConfigStr("download_apk"+"_"+gameurl+"_file"), 303)
	}
}
func mobileGameFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.String())
	pwd, _ := os.Getwd()
	tmpl, err := template.ParseFiles(pwd + "/templates/index.html")
	if err != nil {
		fmt.Println("open game page error:%s", err.Error())
		return
	}
	err = tmpl.Execute(w, gtdMap[strings.Replace(r.URL.String(), "/", "", 1)])
	if err != nil {
		fmt.Println("excute game page error:%s", err.Error())
		return
	}
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
	fmt.Println("server starting...")
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
	/*
		converter, err = iconv.NewCoder(iconv.GBK18030_UTF8_IDX)
		if err != nil {
			fmt.Println("iconv err:", err)
			return
		}*/
	setupMap = make(map[*websocket.Conn]*websocket.Conn)
	gtdMap = make(map[string]GameTempateData)

	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js/"))))
	//http.Handle("/", http.FileServer(http.Dir(".")))
	InitGames()
	//err := http.ListenAndServe("180.168.197.87:18080", nil)
	err = http.ListenAndServe(config.GetConfigStr("ip")+":"+config.GetConfigStr("port"), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	fmt.Println("server stop...")
}
func InitGames() {
	cfg := config.GetConfig()
	for k, _ := range *(*map[string]string)(cfg) {
		if strings.Contains(k, "mobile_game_") {
			url := strings.Split(k, "_")
			if len(url) == 3 {
				http.Handle("/"+url[2]+"/ws", websocket.Handler(SetupServer))
				http.HandleFunc("/"+url[2], mobileGameFunc)
				http.HandleFunc("/"+url[2]+"/download_win", downloadWinFunc)
				http.HandleFunc("/"+url[2]+"/download_apk", downloadApkFunc)
				http.Handle("/"+url[2]+"/apk/", http.StripPrefix("/"+url[2]+"/apk/", http.FileServer(http.Dir(config.GetConfigStr("download_apk"+"_"+url[2])))))
				http.Handle("/"+url[2]+"/win/", http.StripPrefix("/"+url[2]+"/win/", http.FileServer(http.Dir(config.GetConfigStr("download_win"+"_"+url[2])))))
				fmt.Println("/"+url[2]+"/win", config.GetConfigStr("download_win"+"_"+url[2]))
				gamedata := GameTempateData{Name: config.GetConfigStr(k + "_name"), UrlName: url[2]}
				gtdMap[url[2]] = gamedata
			}
			//http.Handle("/"+url[2], http.StripPrefix("/"+url[2], http.FileServer(http.Dir(v))))

		}
	}
}
