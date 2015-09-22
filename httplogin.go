package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

func main() {

	//paracheck := "http://127.0.0.1:7000/httplogin?game=101&zone=1&cmd=plat-token-login"
	//paracheck := "http://50.97.76.195:7002/shen/user/http"
	paracheck := "http://127.0.0.1:7001/shen/user/http"
	//paracheck := "http://14.17.104.56:7001/shen/user/http"
	fmt.Println(paracheck)
	hash := md5.New()
	io.WriteString(hash, "10722whj@whj.whjmalenavy112567123451072280f47650ba0d2834f54c837d518a8eca")
	//h5.bwgame.com.cn/SuperSlot/?platid=67&uid=10722&account=10722&email=whj@whj.whj&gender=male&nickname=navy1125&timestamp=12345&sign=54ad6e9d12baac23ed73e5613913a6db
	//hashstr := plataccount + email + gender + gameid + nickname + platid + timestamp + plataccid + key
	sign := fmt.Sprintf("%x", hash.Sum(nil))
	fmt.Println(sign)
	//resp, err := http.Post(paracheck, "application/x-www-form-urlencoded", strings.NewReader(`{"do":"register-account","data":{"gameid":170,"mid":"{\"sdk\":\"official\",\"account\":\"{\\\"iso\\\":\\\"operator\\\",\\\"systemversion\\\":\\\"00\\\",\\\"IMEI\\\":\\\"867255022606737\\\"}\"}"}}}`))
	str := fmt.Sprintf(`{"do":"plat-token-login","gameid":301,"zoneid":303,"data":{"gameid":301,"zoneid":303,"platinfo":{"account":"10722","platid":"67","email":"whj@whj.whj","gender":"male","nickname":"navy1125","timestamp":"12345","uid":"10722","sign":"%s"}}}`, sign)
	for i := 1; i < 20000; i++ {
		go func() {
			for j := 0; j < 1000; j += 1 {
				n := uint32(rand.Int() + 1)
				str := fmt.Sprintf(`{"do":"plat-token-login","uid":"%d","account":"67::%d","gameid":301,"zoneid":303,"data":{"gameid":301,"zoneid":303,"platinfo":{"account":"67::%d","platid":"67","email":"%d","gender":"male","nickname":"navy1125","timestamp":"12345","uid":"%d","sign":"%s"}}}`, n, n, n, n, n, sign)
				resp, err := http.Post(paracheck, "application/x-www-form-urlencoded", strings.NewReader(str))
				if err == nil {
					text, err := ioutil.ReadAll(resp.Body)
					if err == nil {
						fmt.Println(string(text))
					}
					resp.Body.Close()
				} else {
					fmt.Println(err)
				}

			}
		}()
	}
	/*
		for i := 1; i < 2000; i++ {
			go func() {
				for j := 1000; j < 2000; j += 1 {
				n := uint32(rand.Int())
					str1 := fmt.Sprintf(`{"do":"plat-token-login","uid":"%d","account":"67::%d","gameid":301,"zoneid":303,"data":{"gameid":301,"zoneid":303,"platinfo":{"account":"67::%d","platid":"67","email":"67::%d","gender":"male","nickname":"navy1125","timestamp":"12345","uid":"%d","sign":"%s"}}}`, n,n, n, n, n, sign)
					resp, err := http.Post(paracheck, "application/x-www-form-urlencoded", strings.NewReader(str1))
					if err == nil {
						text, err := ioutil.ReadAll(resp.Body)
						if err == nil {
							fmt.Println(string(text))
						}
						resp.Body.Close()
					} else {
						fmt.Println(err)
					}

				}
			}()
		}
		//*/
	c := make(chan int, 1)
	fmt.Println(<-c)
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
