package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/navy1125/config"
	"github.com/xuyu/iconv"
	mysql "github.com/ziutek/mymysql/autorc"
	_ "github.com/ziutek/mymysql/thrsafe" // You may also use the native engine
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"git.code4.in/mobilegameserver/logging"
)

var (
	db_account                     *mysql.Conn
	db_login                       *mysql.Conn
	db_monitor                     *mysql.Conn
	zoneid_map                     = make(map[uint32]string)
	last_online_out_string         string
	last_online_min                int
	last_online_country_out_string string
	last_online_country_min        int
)

type SdoRetData struct {
	SndaId      uint `json:"sndaId"`
	MaxLevel    int  `json:"maxLevel,omitempty"`
	Logined     bool `json:"logined,omitempty"`
	IsPreCreate bool `json:"IsPreCreate,omitempty"`
}
type SdoRet struct {
	Return_code    int        `json:"return_code"`
	Return_message string     `json:"return_message"`
	Data           SdoRetData `json:"data,omitempty"`
}
type SdoRetCardData struct {
	SndaId      uint   `json:"sndaId"`
	KeyString   string `json:"keystring,omitempty"`
	Charid      uint   `json:"charid,omitempty"`
	Name        string `json:"name,omitempty"`
	Accid       uint   `json:"accid,omitempty"`
	Zone        uint   `json:"zone,omitempty"`
	Plat_accid  uint   `json:"plat_accid,omitempty"`
	IsPreCreate bool   `json:"IsPreCreate,omitempty"`
}
type SdoRetCard struct {
	Return_code    int            `json:"return_code"`
	Return_message string         `json:"return_message"`
	Data           SdoRetCardData `json:"data,omitempty"`
}
type SdoRetCardCountData struct {
	KeyString string `json:"keystring,omitempty"`
	Count     uint   `json:"count,omitempty"`
}
type SdoRetCardCount struct {
	Return_code    int                 `json:"return_code"`
	Return_message string              `json:"return_message"`
	Data           SdoRetCardCountData `json:"data,omitempty"`
}

func HandleGonghuiKeyCard(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	var keystring string
	var count int
	rows, res, err := db_monitor.Query("SELECT COUNT(USERID) AS COUNT FROM SDO_GONGHUI_CARD_USER WHERE KEYSTRING = '" + req.FormValue("KeyString") + "'")
	if err != nil {
		logging.Error("select err:%s", err.Error())
		return
	}
	if len(rows) != 0 {
		COUNT := res.Map("COUNT")
		count = rows[0].Int(COUNT)
	}
	ret := SdoRetCardCount{
		Return_code:    0,
		Return_message: "",
		Data: SdoRetCardCountData{
			KeyString: keystring,
			Count:     uint(count),
		},
	}
	b, _ := json.Marshal(ret)
	w.Write(b)
	logging.Debug("gonghui keystring:%s", req.FormValue("sndaId"))
}
func HandleGonghuiUserCard(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	sndaid, _ := strconv.Atoi(req.FormValue("sndaId"))
	maxlevel := 0
	keystring := ""
	userid := 0
	name := ""
	accid := 0
	plat_accid := 0
	zone := 0
	rows, res, err := db_monitor.Query("SELECT KEYSTRING,USERID,NAME,`ZONE`,ACCID,PLAT_ACCID FROM SDO_GONGHUI_CARD_USER WHERE PLAT_ACCID = " + req.FormValue("sndaId"))
	if err != nil {
		logging.Error("select err:%s", err.Error())
		return
	}
	convert, err1 := iconv.Open("GB2312", "UTF-8")
	if len(rows) != 0 {
		KEYSTRING := res.Map("KEYSTRING")
		USERID := res.Map("USERID")
		NAME := res.Map("NAME")
		ZONE := res.Map("ZONE")
		ACCID := res.Map("ACCID")
		PLAT_ACCID := res.Map("PLAT_ACCID")
		keystring = rows[0].Str(KEYSTRING)
		userid = rows[0].Int(USERID)
		if err1 == nil {
			var err2 error
			name, err2 = convert.ConvString(rows[0].Str(NAME))
			if err2 != nil {
				name = rows[0].Str(NAME)
			}
		} else {
			name = rows[0].Str(NAME)
		}
		accid = rows[0].Int(ACCID)
		zone = rows[0].Int(ZONE)
		plat_accid = rows[0].Int(PLAT_ACCID)
	}
	ret := SdoRetCard{
		Return_code:    0,
		Return_message: "",
		Data: SdoRetCardData{
			SndaId:     uint(sndaid),
			KeyString:  keystring,
			Charid:     uint(userid),
			Name:       name,
			Accid:      uint(accid),
			Zone:       uint(zone),
			Plat_accid: uint(plat_accid),
		},
	}
	b, _ := json.Marshal(ret)
	w.Write(b)
	logging.Debug("gonghui card:%s,%d", req.FormValue("sndaId"), maxlevel)
}
func HandleMaxlevel(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	maxlevel := 0
	sndaid, _ := strconv.Atoi(req.FormValue("sndaId"))
	rows, res, err := db_monitor.Query("SELECT MAXLEVEL FROM USERMAXLEVEL WHERE ACCOUNT = 'S:" + req.FormValue("sndaId") + ":" + req.FormValue("sndaId") + "'")
	if err != nil {
		logging.Error("select err:%s", err.Error())
		return
	}
	if len(rows) != 0 {
		level := res.Map("MAXLEVEL")
		maxlevel = rows[0].Int(level)
	}
	ret := SdoRet{
		Return_code:    0,
		Return_message: "",
		Data: SdoRetData{
			SndaId:   uint(sndaid),
			MaxLevel: maxlevel,
		},
	}
	b, _ := json.Marshal(ret)
	w.Write(b)
	logging.Debug("maxlevel:%s,%d", req.FormValue("sndaId"), maxlevel)
}
func HandleIsPrecreate(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	accid, sndaid := GetAccidBySndaId(req.FormValue("sndaId"))
	ret := SdoRet{
		Return_code:    0,
		Return_message: "",
		Data: SdoRetData{
			SndaId:      uint(sndaid),
			IsPreCreate: false,
		},
	}
	for zoneid, _ := range zoneid_map {
		query_string := fmt.Sprintf("SELECT FIRSTDAY FROM USER_DATA_%d where ACCID=%d", zoneid, accid)
		rows, res, err := db_monitor.Query(query_string)
		if err != nil {
			continue
		}
		if len(rows) != 0 {
			FIRSTDAY := res.Map("FIRSTDAY")
			firstday := rows[0].Int(FIRSTDAY)
			//if firstday < int(GetUnixTime()/60) {
			if firstday < int(1406181600/60) {
				ret.Data.IsPreCreate = true
				break
			}
		}
	}
	b, _ := json.Marshal(ret)
	w.Write(b)
	logging.Debug("isprecreate:%s,%v", req.FormValue("sndaId"), ret.Data.IsPreCreate)
}
func HandleIsOnline(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	accid, sndaid := GetAccidBySndaId(req.FormValue("sndaId"))
	ret := SdoRet{
		Return_code:    0,
		Return_message: "",
		Data: SdoRetData{
			SndaId:  uint(sndaid),
			Logined: false,
		},
	}
	for zoneid, _ := range zoneid_map {
		query_string := fmt.Sprintf("SELECT LASTDAY FROM USER_DATA_%d where ACCID=%d", zoneid, accid)
		rows, res, err := db_monitor.Query(query_string)
		if err != nil {
			continue
		}
		if len(rows) != 0 {
			LASTDAY := res.Map("LASTDAY")
			lastday := rows[0].Int(LASTDAY)
			if int(lastday/1440) == int(GetUnixTime()/86400) {
				ret.Data.Logined = true
				break
			}
		}
	}
	b, _ := json.Marshal(ret)
	w.Write(b)
	logging.Debug("isonline:%s,%v", req.FormValue("sndaId"), ret.Data.Logined)
}
func GetAccidBySndaId(sndastr string) (uint32, uint32) {
	sndaid, _ := strconv.Atoi(sndastr)
	accid := uint32(0)
	rows, res, err := db_account.Query("SELECT ACCID FROM ACCOUNT WHERE ACCOUNTID = " + strconv.Itoa(int(sndaid)))
	if err != nil {
		logging.Error("select err:%s", err.Error())
		return accid, uint32(sndaid)
	}
	if len(rows) != 0 {
		id := res.Map("ACCID")
		accid = uint32(rows[0].Int(id))
	}
	logging.Debug("accid:%d,%d", sndaid, accid)
	return accid, uint32(sndaid)
}
func GetUnixTime() int64 {
	now := time.Now()
	_, offset := now.Zone()
	return now.Unix() + int64(offset)
}
func OnlineServer(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	min := int((GetUnixTime())/60) - 1
	if last_online_min == min && last_online_out_string != "" {
		io.WriteString(w, last_online_out_string+"\n")
		logging.Debug("quest online num:%s,%s,%d,%s", req.RemoteAddr, req.URL.Path, min, last_online_out_string)
		return
	}
	last_online_min = min
	rows, res, err := db_login.Query("SELECT game,zone,name FROM zoneInfo")
	if err != nil {
		logging.Error("select err:%s", err.Error())
		return
	}
	zones := make(map[int]int)
	for _, row := range rows {
		zoneid := res.Map("zone")
		gameid := res.Map("game")
		name := res.Map("name")
		id := row.Int(zoneid)
		zones[id] = 0
		zoneid_map[uint32(row.Int(gameid)<<16+row.Int(zoneid))] = row.Str(name)
	}
	rows, res, err = db_monitor.Query("select * from ONLINENUM_TODAY where timestamp_min = " + strconv.Itoa(min))
	if err != nil {
		logging.Error("select err:%s", err.Error())
	}
	if len(rows) == 0 {
		min = int((GetUnixTime())/60) - 2
		rows, res, err = db_monitor.Query("select * from ONLINENUM_TODAY where timestamp_min = " + strconv.Itoa(min))
		if err != nil {
			logging.Error("select err:%s", err.Error())
		}

	}
	last_online_out_string = ""
	for _, row := range rows {
		zoneid := res.Map("zone_id")
		onlinenum := res.Map("online_number")
		zones[int(int16(row.Int(zoneid)))] = row.Int(onlinenum)
	}
	for id, num := range zones {
		last_online_out_string += strconv.Itoa(int(int16(id))) + "\\" + strconv.Itoa(num) + ";"
	}
	io.WriteString(w, last_online_out_string+"\n")
	logging.Debug("quest online num:%s,%s,%d,%s", req.RemoteAddr, req.URL.Path, min, last_online_out_string)
}
func OnlineCountryServer(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	min := int((GetUnixTime())/60) - 1
	if last_online_country_min == min && last_online_country_out_string != "" {
		io.WriteString(w, last_online_country_out_string+"\n")
		logging.Debug("quest online num:%s,%s,%d,%s", req.RemoteAddr, req.URL.Path, min, last_online_country_out_string)
		return
	}
	last_online_country_min = min
	rows, res, err := db_login.Query("SELECT zone FROM zoneInfo")
	if err != nil {
		logging.Error("select err:%s", err.Error())
		return
	}
	zones := make(map[int]int)
	for _, row := range rows {
		zoneid := res.Map("zone")
		id := row.Int(zoneid)
		zones[(id<<16)+3] = 0
		zones[(id<<16)+4] = 0
		zones[(id<<16)+5] = 0
		zones[(id<<16)+6] = 0
	}
	rows, res, err = db_monitor.Query("select * from COUNTRYONLINENUM_TODAY where timestamp_min = " + strconv.Itoa(min))
	if err != nil {
		logging.Error("select err:%s", err.Error())
	}
	if len(rows) == 0 {
		min = int((GetUnixTime())/60) - 2
		rows, res, err = db_monitor.Query("select * from COUNTRYONLINENUM_TODAY where timestamp_min = " + strconv.Itoa(min))
		if err != nil {
			logging.Error("select err:%s", err.Error())
		}

	}
	last_online_country_out_string = ""
	for _, row := range rows {
		zoneid := res.Map("ZONEID")
		country := res.Map("COUNTRY")
		onlinenum := res.Map("ONLINENUM")
		zones[int(((row.Int(zoneid)&0X0000FFFF)<<16)+row.Int(country))] = row.Int(onlinenum)
	}
	for id, num := range zones {
		if num > 0 {
			last_online_country_out_string += strconv.Itoa(int(id>>16)) + "\\" + strconv.Itoa(int(int16(id))) + "\\" + strconv.Itoa(num) + ";"
		}
	}
	io.WriteString(w, last_online_country_out_string+"\n")
	logging.Debug("quest online num:%s,%s,%d,%s", req.RemoteAddr, req.URL.Path, min, last_online_country_out_string)
}
func main() {
	flag.Parse()
	config.SetConfig("config", *flag.String("config", "config.xml", "config xml file for start"))
	config.SetConfig("logfilename", *flag.String("logfilename", "/log/logfilename.log", "log file name"))
	config.SetConfig("deamon", *flag.String("deamon", "false", "need run as demo"))
	config.SetConfig("port", *flag.String("port", "8000", "http port "))
	config.SetConfig("log", *flag.String("log", "debug", "logger level "))
	config.LoadFromFile(config.GetConfigStr("config"), "global")
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
	logging.Info("server starting...")
	mysqlurl = strings.Replace(mysqlurl, "mysql://", "", 1)
	mysqlurl = strings.Replace(mysqlurl, "@", ":", 1)
	mysqlurl = strings.Replace(mysqlurl, "/", ":", 1)
	mysqlurls := strings.Split(mysqlurl, ":")
	config.SetConfig("dbname", mysqlurls[4])
	db_login = mysql.New("tcp", "", mysqlurls[2]+":"+mysqlurls[3], mysqlurls[0], mysqlurls[1], mysqlurls[4])
	if err != nil {
		logging.Error("db connect error:%s", err.Error())
		return
	}
	mysqlurl = config.GetConfigStr("mysql_monitor")
	if ok, err := regexp.MatchString("^mysql://.*:.*@.*/.*$", mysqlurl); ok == false || err != nil {
		logging.Error("mysql config syntax err:%s", mysqlurl)
		return
	}
	mysqlurl = strings.Replace(mysqlurl, "mysql://", "", 1)
	mysqlurl = strings.Replace(mysqlurl, "@", ":", 1)
	mysqlurl = strings.Replace(mysqlurl, "/", ":", 1)
	mysqlurls = strings.Split(mysqlurl, ":")
	config.SetConfig("dbname", mysqlurls[4])
	db_monitor = mysql.New("tcp", "", mysqlurls[2]+":"+mysqlurls[3], mysqlurls[0], mysqlurls[1], mysqlurls[4])
	if err != nil {
		logging.Error("db connect error:%s", err.Error())
		return
	}
	mysqlurl = config.GetConfigStr("mysql_account")
	if ok, err := regexp.MatchString("^mysql://.*:.*@.*/.*$", mysqlurl); ok == false || err != nil {
		logging.Error("mysql config syntax err:%s", mysqlurl)
		return
	}
	mysqlurl = strings.Replace(mysqlurl, "mysql://", "", 1)
	mysqlurl = strings.Replace(mysqlurl, "@", ":", 1)
	mysqlurl = strings.Replace(mysqlurl, "/", ":", 1)
	mysqlurls = strings.Split(mysqlurl, ":")
	config.SetConfig("dbname", mysqlurls[4])
	db_account = mysql.New("tcp", "", mysqlurls[2]+":"+mysqlurls[3], mysqlurls[0], mysqlurls[1], mysqlurls[4])
	if err != nil {
		logging.Error("db connect error:%s", err.Error())
		return
	}
	http.HandleFunc("/online", OnlineServer)
	http.HandleFunc("/online/country", OnlineCountryServer)
	http.HandleFunc("/gjia/maxlevel", HandleMaxlevel)
	http.HandleFunc("/gjia/isprecreate", HandleIsPrecreate)
	http.HandleFunc("/gjia/isonline", HandleIsOnline)
	http.HandleFunc("/card/gonghuiuser", HandleGonghuiUserCard)
	http.HandleFunc("/card/gonghuikey", HandleGonghuiKeyCard)
	err = http.ListenAndServe(":"+config.GetConfigStr("port"), nil)
	if err != nil {
		logging.Error("ListenAndServe:%s", err.Error())
	}
	logging.Info("server stop...")
}
