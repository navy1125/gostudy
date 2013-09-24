package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"net/url"
	"strconv"
)

func main() {
	buf, _ := ioutil.ReadFile("E:/tmp/zsyy.html")
	str := html.EscapeString(string(buf))
	//fmt.Print(string(buf))
	fmt.Println(html.UnescapeString(str))
	fmt.Println(url.QueryEscape("http://golang.org/pkg/net/url/"))
	urls, err := url.QueryUnescape("http://golang.org/pkg/net/url/")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(urls)
	uparse, err := url.Parse("http://180.153.244.220/gxtime.svn/gxtsdk")
	if err != nil {
		fmt.Println("err", err)
	}
	q := uparse.Query()
	q.Add("name", "whj")
	q.Add("pwd", "wanghaijun")
	uparse.RawQuery = q.Encode()
	fmt.Println(uparse.Host, uparse.Path)
	fmt.Println(uparse,uparse.RequestURI(),uparse.RawQuery)
	testp := TestPrint{Str: "1111", Int: 12}
	fmt.Println(&testp)
	char := make([]byte , 100 ,200)
	fmt.Println(cap(char),len(char))
}

type TestPrint struct {
	Str string
	Int int
}

func (t *TestPrint) String() string{
	//fmt.Println(t.Str)
	return t.Str + strconv.Itoa(t.Int)
}
