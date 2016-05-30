package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/generator"
)

func main() {
	// Begin by allocating a generator. The request and response structures are stored there
	// so we can do error handling easily - the response structure contains the field to
	// report failure.
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
