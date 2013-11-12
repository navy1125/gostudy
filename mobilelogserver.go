package main

import (
	"flag"
	"fmt"
	"github.com/GXTime/logging"
	"github.com/navy1125/config"
	mysql "github.com/ziutek/mymysql/autorc"
	_ "github.com/ziutek/mymysql/native" // Native engine
	"io/ioutil"
	"net/http"
)

var (
	db         *mysql.Conn
	db_monitor *mysql.Conn
)

// hello world, the web server
func LogServer(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	text, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logging.Debug("log err:%s", err.Error())
	}
	logging.Debug("%s,%s", req.URL.String(), text)
}

func main() {
	flag.Parse()
	config.SetConfig("config", *flag.String("config", "config.xml", "config xml file for start"))
	config.SetConfig("logfilename", *flag.String("logfilename", "/log/logfilename.log", "log file name"))
	config.SetConfig("deamon", *flag.String("deamon", "false", "need run as demo"))
	config.SetConfig("port", *flag.String("port", "8000", "http port "))
	config.SetConfig("log", *flag.String("log", "debug", "logger level "))
	config.LoadFromFile(config.GetConfigStr("config"), "global")
	if err := config.LoadFromFile(config.GetConfigStr("config"), "MobileLogServer"); err != nil {
		fmt.Println(err)
		return
	}
	logger, err := logging.NewTimeRotationHandler(config.GetConfigStr("logfilename"), "060102-15")
	if err != nil {
		fmt.Println(err)
		return
	}
	logger.SetLevel(logging.DEBUG)
	logging.AddHandler("MLOG", logger)
	http.HandleFunc("/log/fxsj", LogServer)
	err = http.ListenAndServe(config.GetConfigStr("ip")+":"+config.GetConfigStr("port"), nil)
	if err != nil {
		fmt.Println(err)
		logging.Error("ListenAndServe:%s", err.Error())
	}
}
