package main

import (
	"fmt"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	b, _ := regexp.MatchString("[a-z]", "1222")
	fmt.Println(b)
	b, _ = regexp.MatchString("[:alpha:]", "asdc")
	fmt.Println(b)
	b, _ = regexp.MatchString("f?", "asdc")
	fmt.Println(b)
	b, _ = regexp.MatchString("a{1}", "asdc")
	fmt.Println(b)
	b, _ = regexp.MatchString("[:alpha:]", "asdc")
	fmt.Println(b)
	b, _ = regexp.MatchString("[:alpha:]", "asdc")
	fmt.Println(b)
	fmt.Println(path.Base("/tmp/test.xml"))
	fmt.Println(path.Dir("/tmp/test.xml"))
	fmt.Println(path.Ext("/tmp/test.xml"))
	fmt.Println(filepath.Base("c:\\test.xml"))
	fmt.Println(filepath.Dir("c:\\test.xml"))
	fmt.Println(filepath.Ext("c:\\test.xml"))
	fmt.Println(strings.Split("c:\\test.xml", "\\")[1])
	fmt.Println(strings.Contains("c:\\test.xml", "\\"))
	resizeRe := regexp.MustCompile("([0-9]+)x([0-9]+)(.+)")
	cropRe := regexp.MustCompile("^(x+?)(.+)")
	//cropRe := regexp.MustCompile("([0-9]+)x([0-9]+)(\\+([-0-9]+)\\+([-0-9])+)?(/(northwest|northeast|southwest|southeast|north|west|south|east|center))?(.+)")
	//thumbnailRe := regexp.MustCompile("([0-9]+)x([0-9]+)(/(northwest|northeast|southwest|southeast|north|west|south|east|center))?(.+)")
	fmt.Println(string(resizeRe.Find([]byte("1sdfsdf1x1sddf"))))
	fmt.Println(string(cropRe.Find([]byte("x/north11"))))

}
