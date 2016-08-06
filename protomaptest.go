package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/generator"
)

func main() {
	// Begin by allocating a generator. The request and response structures are stored there
	// so we can do error handling easily - the response structure contains the field to
	// report failure.
	str := `{"cmd_name":"Pmd.UserJsMessageForwardUserPmd_CS","msg":{"errno":"0","st":1470339434,"data":{"ret":0,"uid":1126119,"errno":"0","userInfo":{"roundNum":41,"signature":"","gender":"男","giftCoupon":0,"nickName":"1126119","headUrl":"http:\/\/1251210123.cdn.myqcloud.com\/1251210123\/QZONESTATIC\/headImages\/97.jpg","uid":1126119,"remainder":17187041,"maxMulti":59,"sumRecharge":0,"platId":0,"subPlatId":0,"bankChips":0},"gmLevel":0,"roomInfo":[{"lowestCarry":1,"maxSeat":20,"bankerConfig":{"selectChips":[3000000,5000000,8000000,10000000],"lowestBankerChips":3000000},"lowestBet":100,"roomName":"房间10002","roomId":10002,"maxUser":10000},{"lowestCarry":500000,"maxSeat":20,"bankerConfig":{"selectChips":[30000000,50000000,80000000,100000000],"lowestBankerChips":30000000},"lowestBet":100,"roomName":"房间10004","roomId":10004,"maxUser":10000}]},"do":"Cmd.UserInfoSynRequestLobbyCmd_S"}}`
	var jsmap map[string]interface{}
	err1 := json.Unmarshal([]byte(str), &jsmap)
	if err1 != nil {
		fmt.Println(err1)
	}
	proto.MarshalMessageSetJSON
	g := generator.New()

	file, er := os.Open("/tmp/pmd.pb")
	if er != nil {
		g.Error(er, "reading input")
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		g.Error(err, "reading input")
	}

	if err := proto.Unmarshal(data, g.Request); err != nil {
		g.Error(err, "parsing input proto")
	}

	if len(g.Request.FileToGenerate) == 0 {
		g.Fail("no files to generate")
	}
	fmt.Println(g.Request.ProtoFile)
	fmt.Println(g.Request.FileToGenerate)
	//g.Request.FileToGenerate[0] = "pmd.proto"
	//g.Request.FileToGenerate = append(g.Request.FileToGenerate, "pmd.proto")
	//fmt.Println(g.Request.FileToGenerate[0])
	//fmt.Println(g.Request.Parameter)
	//fmt.Println(g.Request.ProtoFile)
	//fmt.Println(g.Request.XXX_unrecognized)

	//g.CommandLineParameters(g.Request.GetParameter())

	// Create a wrapped version of the Descriptors and EnumDescriptors that
	// point to the file that defines them.
	//g.WrapTypes()
	fmt.Println("XXXXXXXXXXXXXXXXXXXXX")

	//g.SetPackageNames()
	fmt.Println("XXXXXXXXXXXXXXXXXXXXX")
	//g.BuildTypeNameMap()

	//g.GenerateAllFiles()

	// Send back the results.
	data, err = proto.Marshal(g.Response)
	//if err != nil {
	//	g.Error(err, "failed to marshal output proto")
	//}
	//_, err = os.Stdout.Write(data)
	if err != nil {
		g.Error(err, "failed to write output proto")
	}
	//file.Write(data)
	file.Close()
}
