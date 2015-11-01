package main

import (
	"fmt"
	"reflect"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/dbtest"
)

func TestServer() {
	server := &dbtest.DBServer{}
	server.SetPath("/tmp/db/")
	session := server.Session()
	session.Close()
	server.Wipe()
	server.Stop()
	if session == nil {
		server.Wipe()
	}
	server.Wipe()
}
func TestNewDB() {
	m := bson.M{"id": 1}
	fmt.Println(reflect.TypeOf(m["id"]))
	info, err := mgo.ParseURL("mongodb://whj6:whj6@127.0.0.1:27017/whj6?maxPoolSize=10")
	pwd := true
	if err != nil {
		fmt.Println("mgo.ParseURL", err)
	}
	fmt.Println("mgo.ParseURL", info)
	session, err := mgo.DialWithInfo(info)
	if err != nil {
		info, err = mgo.ParseURL("mongodb://127.0.0.1:27017/whj6?maxPoolSize=10")
		pwd = false
		session, err = mgo.DialWithInfo(info)
		if err != nil {
			fmt.Println("mgo.DialWithInfo", err)
			return
		}
	}
	time.Sleep(time.Second * 1)
	admindb := session.DB("whj6")
	if pwd == false {
		err = admindb.AddUser("whj6", "whj6", false)
		if err != nil {
			fmt.Println("admindb.AddUser", err)
			//return
		}
	}
	err = admindb.Login("whj6", "whj6")
	if err != nil {
		fmt.Println("admindb.Login", err)
		//return
	}
	names, _ := admindb.CollectionNames()
	fmt.Println(names)
}

type Logger struct {
}

func (self *Logger) Output(calldepth int, s string) error {
	fmt.Println(s)
	return nil
}

func main() {
	TestNewDB()
	//mgo.SetLogger(&Logger{})
	//mgo.SetDebug(true)
	//TestServer()
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		fmt.Println(err)
		return
	}
	//db.createUser({user: "admin",pwd: "admin",roles: [ { role: "userAdminAnyDatabase", db: "admin" } ]})
	//use whj6
	//db.createUser({user: "whj6",pwd: "whj6",roles: [ { role: "readWrite", db: "whj6" } ]})
	//info, err := mgo.ParseURL("mongodb://whj6:whj6@127.0.0.1:27017/whj6?maxPoolSize=10")
	admindb := session.DB("admin")
	err = admindb.Login("admin", "admin")
	if err != nil {
		err = admindb.AddUser("admin", "admin", false)
		if err != nil {
			fmt.Println(err)
		}
	}
	ruser := &mgo.User{
		Username:     "test",
		Password:     "test",
		OtherDBRoles: map[string][]mgo.Role{"test": []mgo.Role{mgo.RoleReadWrite}},
	}
	admindb = session.DB("test")
	err = admindb.UpsertUser(ruser)
	if err != nil {
		fmt.Println("RoleReadWrite", err)
	}
	admindb.Logout()
	db := session.DB("whj5")
	err = admindb.Login("whj5", "whj5")
	if err != nil {
		fmt.Println("db.Login", err)
	}
	cresult := struct{ ErrMsg string }{}
	err = db.Run(bson.D{{"create", "mycoll"}, {"capped", true}, {"size", 1024}}, &cresult)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("time.Sleep")
	coll := db.C("mycoll")
	fmt.Println("time.Sleep1", err)
	//coll := session.DB("mydb").C("mycoll")
	//err = coll.EnsureIndex(mgo.Index{Key: []string{"n"}, Unique: true})
	err = coll.EnsureIndexKey("-n")
	fmt.Println("time.Sleep2", err)
	err = coll.Insert(bson.M{"n": 3})
	fmt.Println("time.Sleep3", err)
	admindb = session.DB("admin")
	err = admindb.Login("admin", "admin")
	coll = admindb.C("mycoll")
	coll.Insert(bson.M{"n": 3})
	info, err := mgo.ParseURL("mongodb://whj7:whj7@127.0.0.1:27017/whj7?maxPoolSize=10")
	pwd := true
	if err != nil {
		fmt.Println("mgo.ParseURL", err)
	}
	fmt.Println("mgo.ParseURL", info)
	session, err = mgo.DialWithInfo(info)
	if err != nil {
		info, err = mgo.ParseURL("mongodb://127.0.0.1:27017/whj7?maxPoolSize=10")
		pwd = false
		session, err = mgo.DialWithInfo(info)
		if err != nil {
			fmt.Println("mgo.DialWithInfo", err)
			return
		}
	}
	time.Sleep(time.Second * 1)
	admindb = session.DB("whj7")
	if pwd == false {
		err = admindb.AddUser("whj7", "whj7", false)
		if err != nil {
			fmt.Println("admindb.AddUser", err)
			//return
		}
	}
	err = admindb.Login("whj7", "whj7")
	coll = admindb.C("mycoll")
	coll.Insert(bson.M{"whj": 8})
	m := map[string]interface{}{}
	m["whj1"] = 88
	err = coll.Insert(m)
	fmt.Println("time.Sleep3", err)
	result := make([]bson.M, 5)
	err = coll.Find(nil).Select(bson.M{"_id": 1}).All(&result)
	fmt.Println("time.Sleep3", result, err)
	//one := bson.M{}
	//one := string
	one := struct{ WHJ int32 }{}
	//one := struct{ _id string }{}
	err = coll.Find(nil).Select(bson.M{"_id": 0}).One(&one)
	fmt.Println("time.Sleep3", one, err)
	fmt.Println(db, coll)
	db = session.DB("whj9")
	err = db.Login("whj9", "whj9")
	if err != nil {
		err = db.AddUser("whj9", "whj9", false)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	db = session.DB("whj6")
	err = db.Login("whj6", "whj6")
}
