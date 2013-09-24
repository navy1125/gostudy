package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

func main() {
	type Address struct {
		City, State string
	}
	type Person struct {
		XMLName   xml.Name `xml:"person"`
		Id        int      `xml:"id,attr"`
		FirstName string   `xml:"name>first"`
		LastName  string   `xml:"name>last"`
		Age       int      `xml:"age"`
		Height    float32  //`xml:"height,omitempty"`
		Married   bool
		Address
		Comment string `xml:",comment"`
	}

	v := &Person{Id: 13, FirstName: "John", LastName: "Doe", Age: 42}
	v.Comment = " Need more details. "
	v.Address = Address{"Hanga Roa", "Easter Island"}
	v.Height = 12.2

	f, err := os.OpenFile("c:\\test.xml", os.O_CREATE, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
	}
	enc := xml.NewEncoder(f)
	enc.Indent("  ", "    ")
	if err := enc.Encode(v); err != nil {
		fmt.Printf("error: %v\n", err)
	}
	f.Close()
	var vin Person
	fin, _ := os.Open("c:\\test.xml")
	defer fin.Close()
	dec := xml.NewDecoder(fin)
	dec.Decode(&vin)
	fmt.Printf("%d,%s,%s,%s,%f\n", vin.Id, vin.FirstName, vin.LastName, vin.Comment, vin.Height)

	type Text struct {
		World string `xml:"word,attr"`
	}
	type Filter struct {
		XMLName xml.Name `xml:"filter"`
		Text    []Text   `xml:"text"`
	}
	var vin2 Filter
	fin2, _ := os.Open("e:\\tmp\\filters_chat_tw_V3.xml")
	dec2 := xml.NewDecoder(fin2)
	dec2.Decode(&vin2)
	fmt.Println(len(vin2.Text))
	//for i, v2 := range vin2.Text {
	//	fmt.Println(i, v2.World)
	//}

	type Map struct {
		MapID   int `xml:"mapID,attr"`
		NewUser int `xml:"newusermap,attr"`
	}
	type Country struct {
		ID  int   `xml:"id,attr"`
		Map []Map `xml:"map"`
	}
	type Server struct {
		ID      int       `xml:"id,attr"`
		Country []Country `xml:"country"`
	}
	type ServerInfo struct {
		Server []Server `xml:"server"`
	}
	var serverinfo ServerInfo
	fs, _ := os.Open("e:\\tmp\\serverinfo.xml")
	decs := xml.NewDecoder(fs)
	decs.CharsetReader = CharsetReader
	if err := decs.Decode(&serverinfo); err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(serverinfo.Server))
	for _, server := range serverinfo.Server {
		fmt.Println("server:", server.ID)
		for _, country := range server.Country {
			fmt.Println("country:", country.ID)
			for _, m := range country.Map {
				fmt.Println("map:", m.MapID)
			}
		}
	}

}
func CharsetReader(charset string, input io.Reader) (io.Reader, error) {
	return input, nil
}
