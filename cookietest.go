package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func DrawMenu(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, "<a href='/'>HOME <ba><br/>"+"\n")
	io.WriteString(w, "<a href='/readcookie'>Read Cookie <ba><br/>"+"\n")
	io.WriteString(w, "<a href='/writecookie'>Write Cookie <ba><br/>"+"\n")
	io.WriteString(w, "<a href='/deletecookie'>Delete Cookie <ba><br/>"+"\n")

}

func IndexServer(w http.ResponseWriter, req *http.Request) {
	// draw menu
	DrawMenu(w)
}

func ReadCookieServer(w http.ResponseWriter, req *http.Request) {

	// draw menu
	DrawMenu(w)

	// read cookie
	var cookie, err = req.Cookie("testcookiename")
	if err == nil {
		var cookievalue = cookie.Value
		io.WriteString(w, "<b>get cookie value is "+cookievalue+"</b>\n")
	}

}

func WriteCookieServer(w http.ResponseWriter, req *http.Request) {
	// set cookies.
	expire := time.Now().AddDate(0, 0, 1)
	cookie := http.Cookie{Name: "testcookiename", Value: "testcookievalue", Path: "/", Expires: expire, MaxAge: 86400}

	http.SetCookie(w, &cookie)

	//
	// we can not set cookie after writing something to ResponseWriter
	// if so ,we cannot set cookie succefully.
	//
	// so we have draw menu after set cookie
	DrawMenu(w)

}

func DeleteCookieServer(w http.ResponseWriter, req *http.Request) {

	// set cookies.
	cookie := http.Cookie{Name: "testcookiename", Path: "/", MaxAge: -1}
	http.SetCookie(w, &cookie)

	// ABOUT MaxAge
	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	// MaxAge>0 means Max-Age attribute present and given in seconds

	// draw menu
	DrawMenu(w)

}

func main() {

	http.HandleFunc("/", IndexServer)
	http.HandleFunc("/readcookie", ReadCookieServer)
	http.HandleFunc("/writecookie", WriteCookieServer)
	http.HandleFunc("/deletecookie", DeleteCookieServer)

	fmt.Println("listen on 3000")
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
