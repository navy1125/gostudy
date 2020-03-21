package main

import (
	"encoding/json"
	"fmt"
	sjson "github.com/bitly/go-simplejson"
)

func main() {
	str := `{"cmd_name":"Pmd.UserLoginReconnectOkLoginUserPmd_S"}`
	//str := `{"cmd_name":"Pmd.UserJsMessageForwardUserPmd_CS","msg":{"errno":"0","st":1470339434,"data":{"ret":0,"uid":1126119,"errno":"0","userInfo":{"roundNum":41,"signature":"","gender":"男","giftCoupon":0,"nickName":"1126119","headUrl":"http:\/\/1251210123.cdn.myqcloud.com\/1251210123\/QZONESTATIC\/headImages\/97.jpg","uid":1126119,"remainder":17187041,"maxMulti":59,"sumRecharge":0,"platId":0,"subPlatId":0,"bankChips":0},"gmLevel":0,"roomInfo":[{"lowestCarry":1,"maxSeat":20,"bankerConfig":{"selectChips":[3000000,5000000,8000000,10000000],"lowestBankerChips":3000000},"lowestBet":100,"roomName":"房间10002","roomId":10002,"maxUser":10000},{"lowestCarry":500000,"maxSeat":20,"bankerConfig":{"selectChips":[30000000,50000000,80000000,100000000],"lowestBankerChips":30000000},"lowestBet":100,"roomName":"房间10004","roomId":10004,"maxUser":10000}]},"do":"Cmd.UserInfoSynRequestLobbyCmd_S"}}`
	var jsmap map[string]interface{}
	err1 := json.Unmarshal([]byte(str), &jsmap)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(jsmap["msg"].(map[string]interface{})["errno"])
	fmt.Println(jsmap["cmd_name"])
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
	sjs, _ = sjson.NewJson([]byte(`[{"do":"register-auto", "gameid":301, "data":{"mid":"-24912448"}},{"do":"register-auto", "gameid":301, "data":{"mid":"-24912448"}}]`))
	d, _ := sjs.GetIndex(0).MarshalJSON()
	fmt.Println(string(d))
}
func test(args ...string) {
	switch len(args) {
	case 0:
	case 1:
		fmt.Println(args[0])
	}
}
