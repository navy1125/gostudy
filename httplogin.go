package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {

	paracheck := "http://127.0.0.1:7000/httplogin?game=101&zone=1&cmd=plat-token-login"
	fmt.Println(paracheck)
	hash := md5.New()
	io.WriteString(hash, "10722whj@whj.whjmalenavy112567123451072280f47650ba0d2834f54c837d518a8eca")
	//h5.bwgame.com.cn/SuperSlot/?platid=67&uid=10722&account=10722&email=whj@whj.whj&gender=male&nickname=navy1125&timestamp=12345&sign=54ad6e9d12baac23ed73e5613913a6db
	//hashstr := plataccount + email + gender + gameid + nickname + platid + timestamp + plataccid + key
	sign := fmt.Sprintf("%x", hash.Sum(nil))
	fmt.Println(sign)
	//resp, err := http.Post(paracheck, "application/x-www-form-urlencoded", strings.NewReader(`{"do":"register-account","data":{"gameid":170,"mid":"{\"sdk\":\"official\",\"account\":\"{\\\"iso\\\":\\\"operator\\\",\\\"systemversion\\\":\\\"00\\\",\\\"IMEI\\\":\\\"867255022606737\\\"}\"}"}}}`))
	str := fmt.Sprintf(`{"do":"plat-token-login","data":{"gameid":170,"platinfo":{"account":"10722","platid":"67","email":"whj@whj.whj","gender":"male","nickname":"navy1125","timestamp":"12345","uid":"10722","sign":"%s"}}}`, sign)
	resp, err := http.Post(paracheck, "application/x-www-form-urlencoded", strings.NewReader(str))
	//resp, err := http.Get(paracheck)
	if err != nil {
		log.Fatal("http.Get Err:", err)
	}
	defer resp.Body.Close()
	text, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("httpioutil.ReadAll Err:", err)
	}
	fmt.Println("request:%s", str)
	fmt.Println("cookie:", resp.Cookies())
	//fmt.Println(resp.Header.Get("Server"))
	//for key, value := range resp.Header {
	//	fmt.Println(key, value)
	//}
	fmt.Printf("return:%s\n", text)
}
