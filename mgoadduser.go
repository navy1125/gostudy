package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		fmt.Println(err)
		return
	}
	admindb := session.DB("admin")
	err = admindb.Login("admin", "admin")
	if err != nil {
		err = admindb.AddUser("admin", "admin", false)
		if err != nil {
			fmt.Println(err)
		}
	}
	admindb = session.DB("sss")
	err = admindb.AddUser("sss", "sss", false)
	if err != nil {
		fmt.Println("RoleReadWrite", err)
	}
}
