package main

import (
	"fmt"
	"github.com/simonz05/godis/redis"
	"os"
)

func main() {
	// new client on default port 6379, select db 0 and use no password
	c := redis.New("tcp:112.65.197.72:6379", 0, "")

	// set the key "foo" to "Hello Redis"
	if err := c.Set("foo1", "Hello Redis"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	c.Set("hxsg_uid_max", 300000)
	c.Set("hxsg_sid_max", 3000000)
	// retrieve the value of "foo". Returns an Elem obj
	elem, _ := c.Get("max_id")
	elem1, _ := c.Get("foo1")

	// convert the obj to a string and print it
	fmt.Println("foo:", elem.String())
	fmt.Println("foo1:", elem1.String())
	if _, err := c.Hset("myhash", "field", "field"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	elem2, _ := c.Get("name")
	elem3, _ := c.Hget("myhash", "field1")
	size, _ := c.Hlen("myhash")
	fmt.Println("name:", elem2.String())
	fmt.Println("field:", elem3.String())
	fmt.Println("size:", size)
	TestType("asdfdf")
}

func TestType(value interface{}) {
	switch v := value.(type) {
	case string:
		fmt.Println("type:", v, value)
	}

}
