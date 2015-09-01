package main

import (
	"bytes"
	"compress/flate"
	"fmt"
	//"git.code4.in/mobilegameserver/libgame/libcompress"
)

//import "io/ioutil"
import "io"

func main() {
	//inData, _ := ioutil.ReadFile("stuff.dat")
	dict := []byte(``)
	//inData := []byte(`{"do":"plat-token-login","data":{"gameid":170,"platinfo":{"account":"10722","platid":"67","email":"whj@whj.whj","gender":"male","nickname":"navy1125","timestamp":"12345","uid":"10722","sign":"%s"}}}{"do":"plat-token-login","data":{"gameid":170,"platinfo":{"account":"10722","platid":"67","email":"whj@whj.whj","gender":"male","nickname":"navy1125","timestamp":"12345","uid":"10722","sign":"%s"}}}{"do":"plat-token-login","data":{"gameid":170,"platinfo":{"account":"10722","platid":"67","email":"whj@whj.whj","gender":"male","nickname":"navy1125","timestamp":"12345","uid":"10722","sign":"%s"}}}{"do":"plat-token-login","data":{"gameid":170,"platinfo":{"account":"10722","platid":"67","email":"whj@whj.whj","gender":"male","nickname":"navy1125","timestamp":"12345","uid":"10722","sign":"%s"}}}`)
	inData := []byte(`{"do":"request-zone-list","data":{"gameid":170},"gameid":170}`)
	compressedData := new(bytes.Buffer)
	//compress(inData, compressedData, 9)
	compressdict(inData, compressedData, 9, dict)
	data := compressedData.Bytes()
	fmt.Println(fmt.Sprintf("%o,%s,%d,%x", data, string(data), data, data))

	//data, _ := libcompress.Compress(inData)
	fmt.Println("commpress len:", len(inData), compressedData.Len())
	//ioutil.WriteFile("compressed.dat", compressedData.Bytes(), os.ModeAppend)

	deCompressedData := new(bytes.Buffer)
	//decompress(compressedData, deCompressedData)
	decompressdict(compressedData, deCompressedData, dict)
	fmt.Println(deCompressedData)
	//outdata, _ := libcompress.Decompress(data)
	//fmt.Println(deCompressedData)
	//fmt.Println(string(outdata))

}
func compress(src []byte, dest io.Writer, level int) {
	compressor, _ := flate.NewWriter(dest, level)
	compressor.Write(src)
	compressor.Close()
}
func decompress(src io.Reader, dest io.Writer) {
	decompressor := flate.NewReader(src)
	io.Copy(dest, decompressor)
	decompressor.Close()
}
func compressdict(src []byte, dest io.Writer, level int, dict []byte) {
	compressor, _ := flate.NewWriterDict(dest, level, dict)
	compressor.Write(src)
	compressor.Close()
}
func decompressdict(src io.Reader, dest io.Writer, dict []byte) {
	decompressor := flate.NewReaderDict(src, dict)
	io.Copy(dest, decompressor)
	decompressor.Close()
}
