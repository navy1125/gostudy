package main

import (
	"fmt"
	"github.com/xuyu/logging"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native" // Native engine
	"io"
	"net/http"
	"strconv"
)

var (
	db mysql.Conn
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	rows, res, err := db.Query("select * from ONLINENUM_TODAY order by id desc limit 2")
	if err != nil {
		logging.Error("select err:", err)
		return
	}
	var out_string string
	for _, row := range rows {
		zoneid := res.Map("zone_id")
		onlinenum := res.Map("online_number")
		id := row.Int(zoneid)
		num := row.Int(onlinenum)
		out_string += strconv.Itoa(id) + "\\" + strconv.Itoa(num) + ";"
	}
	io.WriteString(w, out_string)
	logging.Debug("quest online num:%s", req.RemoteAddr)
}

func main() {
	logger, err := logging.NewRotationLogger("/tmp/sdoonlineserver.log", "060102-15")
	if err != nil {
		fmt.Println(err)
		return
	}
	db = mysql.New("tcp", "", "127.0.0.1:3306", "root", "123", "MonitorServer")
	err = db.Connect()
	if err != nil {
		logging.Error("db connect error:", err)
		return
	}
	logging.SetDefaultLogger(logger)
	logging.SetPrefix("SDO")
	http.HandleFunc("/online", HelloServer)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		logging.Error("ListenAndServe: ", err)
	}
}
