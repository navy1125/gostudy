package main

import (
	"fmt"
	"log"
	"os/exec"
	//"encoding/json"
)

func main() {
	//cmd := exec.Command("echo", "-n", `{"Name": "Bob", "Age": 32}`)
	//cmd := exec.Command("curl", "-s -O ipa-1257333576.cos.ap-chengdu.myqcloud.com/9517-2019-09-06-09-48-10.ipa")
	cmd := exec.Command("/bin/sh", "-c", `curl -s -O ipa-1257333576.cos.ap-chengdu.myqcloud.com/9517-2019-09-06-09-48-10.ipa`)
	//cmd := exec.Command("ls", "-l -a ziptest.go")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	p := make([]byte, 1024)
	n, err := stdout.Read(p)
	fmt.Println(string(p[:n]), err)

	/*
		var person struct {
			Name string
			Age  int
		}
		if err := json.NewDecoder(stdout).Decode(&person); err != nil {
			log.Fatal(err)
		}
		if err := cmd.Wait(); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s is %d years old\n", person.Name, person.Age)
		// */
}
