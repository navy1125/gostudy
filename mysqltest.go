package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const driverName = "mysql"

type MySQLDB struct {
	s  string
	db *sql.DB
}

//初始化后不释放
var db_mysql *MySQLDB

func NewMySQLDB(s string) (*MySQLDB, error) {
	m := &MySQLDB{s: s}
	db, err := sql.Open(driverName, s)
	if err != nil {
		return nil, err
	}
	m.db = db
	err = db.Ping()
	return m, err
}

func (m *MySQLDB) DB() *sql.DB {
	return m.db
}

type cusField struct {
	dst interface{}
}

func (f *cusField) Scan(src interface{}) error {
	switch s := src.(type) {
	case nil:
		f.dst = nil
	case []byte:
		f.dst = string(s)
	default:
		f.dst = src
	}
	return nil
}

type QueryResult struct {
	D         []map[string]interface{}
	E         error
	Sql       string
	Luafunc   interface{}
	Ownerid   uint64
	Ownerdata interface{}
	W         http.ResponseWriter
	R         *http.Request
	ChWait    chan bool
}

var (
	ChanQueryResult = make(chan *QueryResult, 1024) //异步查询返回
)

func (m *MySQLDB) Query(sql string, args []interface{}) *QueryResult {
	r := &QueryResult{}
	rows, err := m.db.Query(sql, args...)
	if err != nil {
		r.E = err
		return r
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		r.E = err
		return r
	}
	n := len(columns)
	for rows.Next() {
		dest := make([]interface{}, n)
		for i := 0; i < n; i++ {
			dest[i] = &cusField{}
		}
		if err := rows.Scan(dest...); err != nil {
			r.E = err
			return r
		}
		row := make(map[string]interface{})
		for i := 0; i < n; i++ {
			row[columns[i]] = dest[i].(*cusField).dst
		}
		r.D = append(r.D, row)
	}
	if err := rows.Err(); err != nil {
		r.E = err
	}
	return r
}
func (m *MySQLDB) Query2(sql string, w http.ResponseWriter, request *http.Request, ch chan bool, args []interface{}) {
	r := &QueryResult{Sql: sql, W: w, R: request, ChWait: ch}
	go func() {
		rows, err := m.db.Query(sql, args...)
		if err != nil {
			r.E = err
			ChanQueryResult <- r
			return
		}
		defer rows.Close()
		columns, err := rows.Columns()
		if err != nil {
			r.E = err
			ChanQueryResult <- r
			return
		}
		n := len(columns)
		for rows.Next() {
			dest := make([]interface{}, n)
			for i := 0; i < n; i++ {
				dest[i] = &cusField{}
			}
			if err := rows.Scan(dest...); err != nil {
				r.E = err
				ChanQueryResult <- r
				return
			}
			row := make(map[string]interface{})
			for i := 0; i < n; i++ {
				row[columns[i]] = dest[i].(*cusField).dst
			}
			r.D = append(r.D, row)
		}
		if err := rows.Err(); err != nil {
			r.E = err
			ChanQueryResult <- r
			return
		}
		ChanQueryResult <- r
	}()
}
func main() {
	db_mysql, _ = NewMySQLDB("root:12345678@tcp(127.0.0.1:3306)/ScrumStudy")

	go func() {
		for true {
			select {
			case result := <-ChanQueryResult:
				result.W.Write([]byte(fmt.Sprintf("mysql_select return:%v", result.D)))
				result.ChWait <- true
			}
		}
	}()
	http.HandleFunc("/mysql_select", HttpHandleFuncWrapper(mysql_select))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func HttpHandleFuncWrapper(fun func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fun(w, req)
	}
}
func HandleHttpChanRecive(w http.ResponseWriter, r *http.Request) {
}
func mysql_select(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ChWait := make(chan bool, 1)
	db_mysql.Query2("select * from rms_role", w, r, ChWait, nil)
	fmt.Println("mysql_select begin")
	select {
	case <-ChWait:
		break
	}
}
