package main

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"compress/lzw"
	"compress/zlib"
	"fmt"
	"io"
	"io/ioutil"
)

var (
	defaultFlateCompressor = &FlateCompressor{level: flate.DefaultCompression}
	defaultGzipCompressor  = &GzipCompressor{level: flate.DefaultCompression}
	defaultZlibCompressor  = &ZlibCompressor{level: flate.DefaultCompression}
	defaultLzwCompressor   = &LzwCompressor{order: lzw.LSB, litWidth: 8}
)

const (
	CompressType_None  = 0
	CompressType_Flate = 1
	CompressType_Gzip  = 2
	CompressType_Zlib  = 3
	CompressType_Lzw   = 4
)

type CompressFunc func(src []byte) ([]byte, error)
type DeCompressFunc func(src []byte) ([]byte, error)
type DeCompressFromReaderFunc func(src io.Reader) ([]byte, error)

type CompressorInterface interface {
	Compress(src []byte) ([]byte, error)
	DeCompress(src []byte) ([]byte, error)
	DeCompressFromReader(src io.Reader) ([]byte, error)
}

func GetCompressType(compress string) int {
	switch compress {
	case "flate":
		return CompressType_Flate
	case "gzip":
		return CompressType_Gzip
	case "zlib":
		return CompressType_Zlib
	case "lzw":
		return CompressType_Lzw
	}
	return CompressType_None
}

func Compress(ctype int, src []byte) (dst []byte, err error) {
	var tmpstr string
	switch ctype {
	case CompressType_Flate:
		tmpstr = "Compress defaultFlateCompressor"
		dst, err = defaultFlateCompressor.Compress(src)
	case CompressType_Gzip:
		tmpstr = "Compress defaultGzipCompressor"
		dst, err = defaultGzipCompressor.Compress(src)
	case CompressType_Zlib:
		tmpstr = "Compress defaultZlibCompressor"
		dst, err = defaultZlibCompressor.Compress(src)
	case CompressType_Lzw:
		tmpstr = "Compress defaultLzwCompressor"
		dst, err = defaultLzwCompressor.Compress(src)
	default:
		dst = src
	}
	if dst != nil {
		logging.Debug("%s:%d,%d", tmpstr, len(src), len(dst))
	}
	return dst, nil
}
func Decompress(ctype int, src []byte) ([]byte, error) {
	switch ctype {
	case CompressType_Flate:
		logging.Debug("Decompress defaultFlateCompressor")
		return defaultFlateCompressor.Decompress(src)
	case CompressType_Gzip:
		logging.Debug("Decompress defaultGzipCompressor")
		return defaultGzipCompressor.Decompress(src)
	case CompressType_Zlib:
		logging.Debug("Decompress defaultZlibCompressor")
		return defaultZlibCompressor.Decompress(src)
	case CompressType_Lzw:
		logging.Debug("Decompress defaultLzwCompressor")
		return defaultLzwCompressor.Decompress(src)
	}
	return src, nil
}
func DecompressFromReader(ctype int, src io.Reader) ([]byte, error) {
	switch ctype {
	case CompressType_Flate:
		logging.Debug("DecompressFromReader defaultFlateCompressor")
		return defaultFlateCompressor.DecompressFromReader(src)
	case CompressType_Gzip:
		logging.Debug("DecompressFromReader defaultGzipCompressor")
		return defaultGzipCompressor.DecompressFromReader(src)
	case CompressType_Zlib:
		logging.Debug("DecompressFromReader defaultZlibCompressor")
		return defaultZlibCompressor.DecompressFromReader(src)
	case CompressType_Lzw:
		logging.Debug("DecompressFromReader defaultLzwCompressor")
		return defaultLzwCompressor.DecompressFromReader(src)
	}
	return ioutil.ReadAll(src)
}

type FlateCompressor struct {
	level int
	dict  []byte
	//compressor   *flate.Writer
	//decompressor io.ReadCloser
	//cdest        *bytes.Buffer
	//ddest *bytes.Buffer
}

func (self *FlateCompressor) Compress(src []byte) ([]byte, error) {
	var err error
	var compressor *flate.Writer
	cdest := bytes.NewBuffer(make([]byte, 0, len(src)))
	if self.dict == nil {
		compressor, err = flate.NewWriter(cdest, self.level)
	} else {
		compressor, err = flate.NewWriterDict(cdest, self.level, self.dict)
	}
	//compressor.Reset(cdest)
	compressor.Write(src)
	err = compressor.Close()
	if err != nil {
		fmt.Println("Compress Close err:%s", err.Error())
	}
	return cdest.Bytes(), err
}

func (self *FlateCompressor) Decompress(src []byte) ([]byte, error) {
	data, err := self.DecompressFromReader(bytes.NewBuffer(src))
	if err != nil {
		fmt.Println("Decompress err:%s,%p", err.Error(), src)
	}
	return data, err
}
func (self *FlateCompressor) DecompressFromReader(src io.Reader) ([]byte, error) {
	ddest := bytes.NewBuffer(nil)
	var decompressor io.ReadCloser
	if self.dict == nil {
		decompressor = flate.NewReader(src)
	} else {
		decompressor = flate.NewReaderDict(src, self.dict)
	}
	err := decompressor.Close()
	if err != nil {
		fmt.Println("DecompressFromReader err:%s", err.Error())
	} else {
		_, err = io.Copy(ddest, decompressor)
	}
	return ddest.Bytes(), err
}

type ZlibCompressor struct {
	level int
	dict  []byte
}

func (self *ZlibCompressor) Compress(src []byte) ([]byte, error) {
	var err error
	var compressor *zlib.Writer
	cdest := bytes.NewBuffer(make([]byte, 0, len(src)))
	if self.dict == nil {
		compressor, err = zlib.NewWriterLevel(cdest, self.level)
	} else {
		compressor, err = zlib.NewWriterLevelDict(cdest, self.level, self.dict)
	}
	compressor.Write(src)
	err = compressor.Close()
	if err != nil {
		fmt.Println("Compress Close err:%s", err.Error())
	}
	return cdest.Bytes(), err
}

func (self *ZlibCompressor) Decompress(src []byte) ([]byte, error) {
	data, err := self.DecompressFromReader(bytes.NewBuffer(src))
	if err != nil {
		fmt.Println("Decompress err:%s,%p", err.Error(), src)
	}
	return data, err
}
func (self *ZlibCompressor) DecompressFromReader(src io.Reader) ([]byte, error) {
	ddest := bytes.NewBuffer(nil)
	decompressor, err := zlib.NewReaderDict(src, self.dict)
	if err != nil {
		fmt.Println("DecompressFromReader err:%s", err.Error())
	} else {
		_, err = io.Copy(ddest, decompressor)
	}
	if err != nil {
		fmt.Println("DecompressFromReader err:%s", err.Error())
	}
	return ddest.Bytes(), err
}

type LzwCompressor struct {
	litWidth int
	order    lzw.Order
}

func (self *LzwCompressor) Compress(src []byte) ([]byte, error) {
	cdest := bytes.NewBuffer(make([]byte, 0, len(src)))
	compressor := lzw.NewWriter(cdest, self.order, self.litWidth)
	compressor.Write(src)
	err := compressor.Close()
	if err != nil {
		fmt.Println("Compress Close err:%s", err.Error())
	}
	return cdest.Bytes(), err
}

func (self *LzwCompressor) Decompress(src []byte) ([]byte, error) {
	data, err := self.DecompressFromReader(bytes.NewBuffer(src))
	if err != nil {
		fmt.Println("Decompress err:%s,%p", err.Error(), src)
	}
	return data, err
}
func (self *LzwCompressor) DecompressFromReader(src io.Reader) ([]byte, error) {
	ddest := bytes.NewBuffer(nil)
	decompressor := lzw.NewReader(src, self.order, self.litWidth)
	_, err := io.Copy(ddest, decompressor)
	if err != nil {
		fmt.Println("DecompressFromReader err:%s", err.Error())
	}
	return ddest.Bytes(), err
}

type GzipCompressor struct {
	level int
}

func (self *GzipCompressor) Compress(src []byte) ([]byte, error) {
	cdest := bytes.NewBuffer(make([]byte, 0, len(src)))
	compressor, err := gzip.NewWriterLevel(cdest, self.level)
	compressor.Write(src)
	err = compressor.Close()
	if err != nil {
		fmt.Println("Compress Close err:%s", err.Error())
	}
	return cdest.Bytes(), err
}

func (self *GzipCompressor) Decompress(src []byte) ([]byte, error) {
	data, err := self.DecompressFromReader(bytes.NewBuffer(src))
	if err != nil {
		fmt.Println("Decompress err:%s,%p", err.Error(), src)
	}
	return data, err
}
func (self *GzipCompressor) DecompressFromReader(src io.Reader) ([]byte, error) {
	ddest := bytes.NewBuffer(nil)
	decompressor, err := gzip.NewReader(src)
	if err != nil {
		fmt.Println("DecompressFromReader err:%s", err.Error())
	}
	//err = decompressor.Close()
	if err != nil {
		fmt.Println("DecompressFromReader err:%s", err.Error())
	} else {
		_, err = io.Copy(ddest, decompressor)
	}
	return ddest.Bytes(), err
}
