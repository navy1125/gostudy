package main

import (
	"encoding/json"
	"fmt"
	sjson "github.com/bitly/go-simplejson"
)

func main() {
	_, err := sjson.NewJson([]byte(`{"do":"register-auto", "gameid":301, "data":{"mid":"-24912448"}}`))
	fmt.Println(err)
	sjs, _ := sjson.NewJson([]byte(`{
	"curequip":[1,2,3,4,5],
	"effect":70,
	"gold":50,
	"item":{},
	"lottery":{"item":[0,0,0,0,0,0,0,0,0],"itemgot":[0,0,0,0,0,0,0,0,0],"refreshtime":0},
	"music":1,
	"name":"强哥",
	"silver":60,
	"stage":{},
	"strength":100,
	"version":1
		}`))
	sjs2, _ := sjson.NewJson([]byte(`{}`))
	var i json.Number
	i = "100"
	sjs2.Set("cur", i)
	//sjs2.Set("cur", "100")
	sjs.Set("sjs2", sjs2)
	sjs.Get("curequip").MustArray()[2] = 10
	jsstr, _ := sjs.MarshalJSON()
	sjs, _ = sjson.NewJson(jsstr)
	//fmt.Println(string(jsstr))
	sjs3 := sjs.Get("sjs2")
	sjs.Get("effect").UnmarshalJSON([]byte(`{}`))
	//fmt.Println(sjs.Map())
	fmt.Println(sjs3.Get("cur"))
	//fmt.Println("aaa", sjs.Get("lottery").Get("item").MustArray()[0])
	//fmt.Println("bbb", sjs.Get("sjs2").Get("cur").MustString(""))
	sjs2.Set("key", make(map[string]interface{}))
	//b, _ := sjs2.MarshalJSON()
	fmt.Println(sjs.Get("sjs2").Get("cur").Int())
	test()
}
func test(args ...string) {
	switch len(args) {
	case 0:
	case 1:
		fmt.Println(args[0])
	}
}
