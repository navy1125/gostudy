package main

import (
	"flag"
	"fmt"
	"github.com/navy1125/config"
	"github.com/xuyu/logging"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native" // Native engine
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	db mysql.Conn
)

// hello world, the web server
func OnlineServer(w http.ResponseWriter, req *http.Request) {
	rows, res, err := db.Query("show tables")
	if err != nil {
		err = db.Connect()
		if err != nil {
			logging.Error("db connect error:%s", err.Error())
			return
		}
	}
	rows, res, err = db.Query("show tables")
	if err != nil {
		logging.Error("select err:%s", err.Error())
		return
	}
	zones := make(map[int]int)
	for _, row := range rows {
		tname := res.Map("Tables_in_" + config.GetConfigStr("db"))
		if ok, _ := regexp.MatchString("USER_DATA_", row.Str(tname)); ok == true {
			zoneid := strings.Replace(row.Str(tname), "USER_DATA_", "", 1)
			id, _ := strconv.Atoi(zoneid)
			zones[id] = 0
		}
	}
	now := time.Now()
	_, offset := now.Zone()
	min := int((now.Unix()+int64(offset))/60) - 1
	query_string := "select * from ONLINENUM_TODAY where timestamp_min = " + strconv.Itoa(min)
	rows, res, err = db.Query(query_string)
	if err != nil {
		logging.Error("select err:%s", err.Error())
	}
	var out_string string
	for _, row := range rows {
		zoneid := res.Map("zone_id")
		onlinenum := res.Map("online_number")
		zones[row.Int(zoneid)] = row.Int(onlinenum)
	}
	for id, num := range zones {
		out_string += strconv.Itoa(id) + "\\" + strconv.Itoa(num) + ";"
	}
	io.WriteString(w, out_string+"\n")
	logging.Debug("quest online num:%s", req.RemoteAddr)
}

func main() {
	flag.Parse()
	config.SetConfig("config", *flag.String("config", "config.xml", "config xml file for start"))
	config.SetConfig("logfilename", *flag.String("logfilename", "/log/logfilename.log", "log file name"))
	config.SetConfig("deamon", *flag.String("deamon", "false", "need run as demo"))
	config.SetConfig("port", *flag.String("port", "8000", "http port "))
	config.SetConfig("log", *flag.String("log", "debug", "logger level "))
	if err := config.LoadFromFile(config.GetConfigStr("config"), "SdoOnlineServer"); err != nil {
		fmt.Println(err)
		return
	}
	logger, err := logging.NewTimeRotationHandler(config.GetConfigStr("logfilename"), "060102-15")
	if err != nil {
		fmt.Println(err)
		return
	}
	logger.SetLevel(logging.DEBUG)
	logging.AddHandler("SDO", logger)
	mysqlurl := config.GetConfigStr("mysql")
	if ok, err := regexp.MatchString("^mysql://.*:.*@.*/.*$", mysqlurl); ok == false || err != nil {
		logging.Error("mysql config syntax err:%s", mysqlurl)
		return
	}
	mysqlurl = strings.Replace(mysqlurl, "mysql://", "", 1)
	mysqlurl = strings.Replace(mysqlurl, "@", ":", 1)
	mysqlurl = strings.Replace(mysqlurl, "/", ":", 1)
	mysqlurls := strings.Split(mysqlurl, ":")
	config.SetConfig("db", mysqlurls[4])
	db = mysql.New("tcp", "", mysqlurls[2]+":"+mysqlurls[3], mysqlurls[0], mysqlurls[1], mysqlurls[4])
	err = db.Connect()
	if err != nil {
		logging.Error("db connect error:%s", err.Error())
		return
	}
	http.HandleFunc("/online", OnlineServer)
	err = http.ListenAndServe(":"+config.GetConfigStr("port"), nil)
	if err != nil {
		logging.Error("ListenAndServe:%s", err.Error())
	}
}
