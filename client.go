package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	//"net/url"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
	//resp, err := http.PostForm("http://112.65.197.72:8080/post", url.Values{"uuid": {"1234"}, "data": {"12"}})
	//resp, err := http.Get("http://www.bwgame.com.cn")
	//resp, err := http.Get("http://127.0.0.1:8080/client.go")
	//resp, err := http.Get("http://127.0.0.1:12346/hello")
	//resp, err := http.PostForm("http://192.168.85.71:12346/hello", url.Values{"name": {"whj"}, "age": {"12"}})
	//resp, err := http.Get("http://192.168.85.71:8080/fileservertest.go")
	/*f, _ := os.Open("c:\\test.xml")
	resp, err := http.Post("http://192.168.85.71:8080/post", "go", f)
	//*/

	tsec := time.Now().Unix()
	para := "qid=1234" + "&time=" + strconv.Itoa(int(tsec)) + "&server_id=10351"
	hash := md5.New()
	io.WriteString(hash, para+"344a5ec3dacac264f8603db0f24c9f49")
	para = "http://120.132.152.40:8000/zssj/juxian/auth?" + fmt.Sprintf("%s&sign=%x", para, hash.Sum(nil))
	fmt.Println(para)
	hash.Reset()
	parabill := "qid=1234" + "&order_amount=100" + "&order_id=1000" + "&server_id=10351"
	io.WriteString(hash, "1234100100010351"+"344a5ec3dacac264f8603db0f24c9f49")
	parabill = "http://120.132.152.40:8000/zssj/juxian/bill?" + fmt.Sprintf("%s&sign=%x", parabill, hash.Sum(nil))
	fmt.Println(parabill)
	hash.Reset()
	paracheck := "qid=1234" + "&server_id=10351"
	io.WriteString(hash, paracheck+"344a5ec3dacac264f8603db0f24c9f49")
	paracheck = "http://192.168.85.71:8000/zssj/juxian/check?" + fmt.Sprintf("%s&sign=%x", paracheck, hash.Sum(nil))
	fmt.Println(paracheck)
	hash.Reset()
	parakw := "from_id=kw" + "&game_id=3055" + "&login_name=1234" + "&server_id=10351" + "&time=" + strconv.Itoa(int(tsec))
	io.WriteString(hash, parakw+"&key="+"344a5ec3dacac264f8603db0f24c9f49")
	parakw = "http://192.168.85.71:8000/zssj/kw/auth?" + fmt.Sprintf("%s&token=%x", parakw, hash.Sum(nil))
	fmt.Println(parakw)
	hash.Reset()
	parakwbill := "from_id=kw" + "&game_id=3055" + "&login_name=1234" + "&money=1" + "&order_no=1111" + "&server_id=10351"
	io.WriteString(hash, "kw305512341111110351344a5ec3dacac264f8603db0f24c9f49")
	parakwbill = "http://192.168.85.71:8000/zssj/kw/bill?" + fmt.Sprintf("%s&token=%x", parakwbill, hash.Sum(nil))
	fmt.Println(parakwbill)
	hash.Reset()
	parakwcheck := "from_id=kw" + "&game_id=3055" + "&login_name=1234" + "&server_id=10351"
	io.WriteString(hash, "kw3055123410351344a5ec3dacac264f8603db0f24c9f49")
	parakwcheck = "http://192.168.85.71:8000/zssj/kw/check?" + fmt.Sprintf("%s&token=%x", parakwcheck, hash.Sum(nil))
	fmt.Println(parakwcheck)

	hash.Reset()
	para619 := "from_id=kw" + "&game_id=3055" + "&login_name=1234" + "&server_id=10351" + "&time=" + strconv.Itoa(int(tsec))
	io.WriteString(hash, para619+"&key="+"344a5ec3dacac264f8603db0f24c9f49")
	para619 = "http://192.168.85.71:8000/zssj/kw/auth?" + fmt.Sprintf("%s&token=%x", para619, hash.Sum(nil))
	fmt.Println(para619)
	resp, err := http.Get(paracheck)
	if err != nil {
		log.Fatal("http.Get Err:", err)
	}
	defer resp.Body.Close()
	text, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("httpioutil.ReadAll Err:", err)
	}
	fmt.Println(resp.Header.Get("Server"))
	for key, value := range resp.Header {
		fmt.Println(key, value)
	}
	fmt.Printf("%s", text)
	//http.SetCookie(w, cookie)
	//resp.Body.Close()
	//if err := postFile("c:\\test.xml", "http://192.168.85.71:8080/post"); err != nil {
	//	fmt.Println(err)
	//}
}

func postFile(filename string, taretUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, err := bodyWriter.CreateFormFile("upload", filepath.Base(filename))
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error open file", filename)
		return err
	}
	if err != nil {
		return nil
	}
	size, err := io.Copy(fileWriter, fh)
	fmt.Println(filename, "size:", size)
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	fmt.Println("file type:", contentType)
	resp, err := http.Post(taretUrl, contentType, bodyBuf)
	if err != nil {
		fmt.Println("post err", filename)
		return err
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}
