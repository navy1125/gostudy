package main

import (
	"fmt"

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

type Logger struct {
}

func (self *Logger) Output(calldepth int, s string) error {
	fmt.Println(s)
	return nil
}

func main() {
	mgo.SetLogger(&Logger{})
	mgo.SetDebug(true)
	TestServer()
	return
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
		Username:     "whj5",
		Password:     "whj5",
		OtherDBRoles: map[string][]mgo.Role{"whj5": []mgo.Role{mgo.RoleReadWrite}},
	}
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
	coll := db.C("mycoll")
	coll.EnsureIndexKey("n")
	//coll := session.DB("mydb").C("mycoll")
	//err = coll.EnsureIndex(mgo.Index{Key: []string{"n"}, Unique: true})
	coll.EnsureIndexKey("-n")
	coll.Insert(bson.M{"n": 3})
	admindb = session.DB("admin")
	err = admindb.Login("admin", "admin")
	coll = admindb.C("mycoll")
	coll.Insert(bson.M{"n": 3})
	info, err := mgo.ParseURL("mongodb://whj5:whj5@127.0.0.1:27017/admin?maxPoolSize=10")
	if err != nil {
		fmt.Println("mgo.ParseURL", err)
	}
	fmt.Println("mgo.ParseURL", info)
	session, err = mgo.DialWithInfo(info)
	if err != nil {
		fmt.Println("mgo.DialWithInfo", err)
	}
	admindb = session.DB("whj5")
	coll = admindb.C("mycoll")
	coll.Insert(bson.M{"whj": 8})
}
