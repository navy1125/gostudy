package main

import (
	"bytes"
	"code.google.com/p/go.net/websocket"
	"flag"
	"fmt"
	"github.com/navy1125/config"
	iconv "github.com/navy1125/iconv"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

type GameTempateData struct {
	Name    string
	UrlName string
}

var (
	logNutex  sync.Mutex
	socketMap map[*websocket.Conn]string
	gamesMap  map[string]int
	gtdMap    map[string]GameTempateData
	converter *iconv.Converter
	//utf8
)

type JSONCommand struct {
	Id   string `json:"Id,omitempty"`
	Data string `json:"Data,omitempty"`
}

func SetupServer(ws *websocket.Conn) {
	var message JSONCommand
	//world := createWorld()

	for {
		err := websocket.JSON.Receive(ws, &message)
		if err != nil {
			if v, ok := socketMap[ws]; ok {
				gamesMap[v]--
				Broadcask("0", []byte("退出减少后当前用户数量:"+v+":"+strconv.Itoa(gamesMap[v])), ws)
			}
			delete(socketMap, ws)
			log.Print("Receive error - stopping worker: ", err)
			break
		}
		fmt.Println(message.Id, message.Data)
		if strings.Contains(message.Id, "setup apk") {
			game := message.Data
			if _, ok := socketMap[ws]; !ok {
				if v, ok := gamesMap[game]; ok {
					gamesMap[game]++
					socketMap[ws] = game
					Broadcask("0", []byte("增加后当前用户数量:"+game+":"+strconv.Itoa(gamesMap[game])), ws)
					if v == 0 {
						SendMessage("0", []byte("start setup apk:"+game+" now..."), ws)
						//execBat("publish_android.bat", ws)
						//execBat("copy_apk.bat", ws)
						SendMessage("0", []byte(config.GetConfigStr("bat_apk_"+game)), ws)
						execBat(config.GetConfigStr("bat_apk_"+game), ws)
						Broadcask("setup finish apk", []byte(game), ws)
						for k, v2 := range socketMap {
							if v2 == game {
								gamesMap[v2]--
								delete(socketMap, k)
								Broadcask("0", []byte("完成减少后当前用户数量:"+v2+":"+strconv.Itoa(gamesMap[v2])), ws)
							}
						}
						for _, v2 := range socketMap {
							if v2 == game {
								//k.Close()
							}
						}
					} else {
						SendMessage("0", []byte("wait setup apk:"+game+" now..."), ws)
					}
				}
			} else {
				SendMessage("0", []byte("重复加入"), ws)
			}
		}
		if strings.Contains(message.Id, "setup win") {
			game := message.Data
			if _, ok := socketMap[ws]; !ok {
				if v, ok := gamesMap[game]; ok {
					gamesMap[game]++
					socketMap[ws] = game
					Broadcask("0", []byte("增加后当前用户数量:"+game+":"+strconv.Itoa(gamesMap[game])), ws)
					if v == 0 {
						SendMessage("0", []byte("start setup win:"+game+" now..."), ws)
						//execBat("publish_android.bat", ws)
						//execBat("copy_apk.bat", ws)
						SendMessage("0", []byte(config.GetConfigStr("bat_apk_"+game)), ws)
						execBat(config.GetConfigStr("bat_win_"+game), ws)
						Broadcask("setup finish win", []byte(game), ws)
						for k, v2 := range socketMap {
							if v2 == game {
								gamesMap[v2]--
								delete(socketMap, k)
								Broadcask("0", []byte("完成减少后当前用户数量:"+v2+":"+strconv.Itoa(gamesMap[v2])), ws)
							}
						}
						for _, v2 := range socketMap {
							if v2 == game {
								//k.Close()
							}
						}
					} else {
						SendMessage("0", []byte("wait setup win:"+game+" now..."), ws)
					}
				}
			} else {
				SendMessage("0", []byte("重复加入"), ws)
			}
		}
		if strings.Contains(message.Id, "reset huodong") {
			game := message.Data
			if _, ok := socketMap[ws]; !ok {
				if v, ok := gamesMap[game]; ok {
					gamesMap[game]++
					socketMap[ws] = game
					Broadcask("0", []byte("增加后当前用户数量:"+game+":"+strconv.Itoa(gamesMap[game])), ws)
					if v == 0 {
						SendMessage("0", []byte("start reset huodong:"+game+" now..."), ws)
						//execBat("publish_android.bat", ws)
						//execBat("copy_apk.bat", ws)
						SendMessage("0", []byte(config.GetConfigStr("bat_huodong_"+game)), ws)
						execBat(config.GetConfigStr("bat_huodong_"+game), ws)
						Broadcask("setup finish reset huodong", []byte(game), ws)
						for k, v2 := range socketMap {
							if v2 == game {
								gamesMap[v2]--
								delete(socketMap, k)
								Broadcask("0", []byte("完成减少后当前用户数量:"+v2+":"+strconv.Itoa(gamesMap[v2])), ws)
							}
						}
						for _, v2 := range socketMap {
							if v2 == game {
								//k.Close()
							}
						}
					} else {
						SendMessage("0", []byte("wait reset huodong:"+game+" now..."), ws)
					}
				}
			} else {
				SendMessage("0", []byte("重复加入"), ws)
			}
		}

		//err = websocket.JSON.Send(ws, world.Update(message))
		//if err != nil {
		//	log.Print("Send error - stopping worker: ", err)
		//	break
		//}
	}
}
func SendMessage(id string, b []byte, ws *websocket.Conn) {
	cmd := JSONCommand{Id: id}
	out := make([]byte, len(b)*4)
	if l, err := converter.CodeConvertFunc(b, out); err == nil && l > 0 {
		cmd.Data = string(out[0:l])
	} else {
		cmd.Data = string(b)
	}
	fmt.Println(cmd.Data, len(cmd.Data))
	SendJSON(&cmd, ws)
}
func Broadcask(id string, b []byte, ws *websocket.Conn) {
	cmd := JSONCommand{Id: id}
	out := make([]byte, len(b)*4)
	if l, err := converter.CodeConvertFunc(b, out); err == nil && l > 0 {
		cmd.Data = string(out[0:l])
	} else {
		cmd.Data = string(b)
	}
	fmt.Println(cmd.Data, len(cmd.Data))
	BroadcaskJSON(&cmd, ws)
}

func BroadcaskJSON(msg *JSONCommand, ws *websocket.Conn) {
	game, ok := socketMap[ws]
	if !ok {
		return
	}
	for k, v := range socketMap {
		if v == game {
			SendJSON(msg, k)
		}
	}
}
func SendJSON(msg *JSONCommand, ws *websocket.Conn) {
	websocket.JSON.Send(ws, msg)
}
func execBat(bat string, ws *websocket.Conn) {
	cmd := exec.Command(bat)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		Broadcask("1", []byte(err.Error()), ws)
		fmt.Print(string([]byte(err.Error())))
	}
	stdin, err := cmd.StderrPipe()
	if err != nil {
		Broadcask("1", []byte(err.Error()), ws)
		fmt.Print(string([]byte(err.Error())))
	}
	if err := cmd.Start(); err != nil {
		Broadcask("1", []byte(err.Error()), ws)
		fmt.Print(string([]byte(err.Error())))
	}
	go outputFunc(stdout, ws)
	go outputFunc(stdin, ws)
	cmd.Wait()
}
func downloadWinFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.String())
	gameurl := strings.Split(r.URL.String(), "/")[1]
	num, ok := gamesMap[gameurl]
	if !ok {
		w.Write([]byte("不可识别的游戏"))
		defer r.Body.Close()
		return
	}
	if _, err := os.Open(config.GetConfigStr("download_win"+"_"+gameurl) + "/" + config.GetConfigStr("download_win"+"_"+gameurl+"_file")); err != nil {
		//w.Write([]byte(config.GetConfigStr("download_win"+"_"+gameurl) + "/" + config.GetConfigStr("download_win"+"_"+gameurl+"_file")))
		if num != 0 {
			w.Write([]byte("有人正在做版本,请稍后再试:" + gameurl + ":" + strconv.Itoa(num)))
		} else {
			w.Write([]byte("目前没有版本可以下载,请先做版本:" + gameurl + ":" + strconv.Itoa(num)))
		}
		defer r.Body.Close()
		return
	}
	http.Redirect(w, r, strings.Replace(r.URL.String(), "download_win", "win", 1)+"/"+config.GetConfigStr("download_win"+"_"+gameurl+"_file"), 303)
}
func downloadApkFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.String())
	gameurl := strings.Split(r.URL.String(), "/")[1]
	num, ok := gamesMap[gameurl]
	if !ok {
		w.Write([]byte("不可识别的游戏"))
		defer r.Body.Close()
		return
	}
	if _, err := os.Open(config.GetConfigStr("download_apk"+"_"+gameurl) + "/" + config.GetConfigStr("download_apk"+"_"+gameurl+"_file")); err != nil {
		//w.Write([]byte(config.GetConfigStr("download_apk"+"_"+gameurl) + "/" + config.GetConfigStr("download_apk"+"_"+gameurl+"_file")))
		if num != 0 {
			w.Write([]byte("有人正在做版本,请稍后再试:" + gameurl + ":" + strconv.Itoa(num)))
		} else {
			w.Write([]byte("目前没有版本可以下载,请先做版本:" + gameurl + ":" + strconv.Itoa(num)))
		}
		return
	}
	http.Redirect(w, r, strings.Replace(r.URL.String(), "download_apk", "apk", 1)+"/"+config.GetConfigStr("download_apk"+"_"+gameurl+"_file"), 303)
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

func outputFunc(r io.ReadCloser, ws *websocket.Conn) {
	b := bytes.NewBuffer(make([]byte, 1024))
	for {
		n, err := r.Read(b.Bytes())
		if n > 0 {
			logNutex.Lock()
			Broadcask("0", b.Bytes()[0:n], ws)
			fmt.Print(string(b.Bytes()[0:n]))
			logNutex.Unlock()
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			Broadcask("1", []byte(err.Error()), ws)
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
	converter, err = iconv.NewCoder(iconv.GBK18030_UTF8_IDX)
	if err != nil {
		fmt.Println("iconv err:", err)
		return
	}
	socketMap = make(map[*websocket.Conn]string)
	gamesMap = make(map[string]int)
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
				gamesMap[url[2]] = 0
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
