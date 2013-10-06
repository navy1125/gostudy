package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	//"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	//"net/url"
)

func main() {
	//resp, err := http.Get("http://www.bwgame.com.cn")
	//resp, err := http.Get("http://127.0.0.1:8080/client.go")
	//resp, err := http.Get("http://127.0.0.1:12346/hello")
	//resp, err := http.PostForm("http://192.168.85.71:12346/hello", url.Values{"name": {"whj"}, "age": {"12"}})
	//resp, err := http.Get("http://192.168.85.71:8080/fileservertest.go")
	/*f, _ := os.Open("c:\\test.xml")
	resp, err := http.Post("http://192.168.85.71:8080/post", "go", f)

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
	//*/
	//http.SetCookie(w, cookie)
	//resp.Body.Close()
	if err := postFile("c:\\test.xml", "http://192.168.85.71:8080/post"); err != nil {
		fmt.Println(err)
	}
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
