package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"
	"fmt"
)

func main() {
	//inData := []byte(`{"do":"plat-token-login","data":{"gameid":170,"platinfo":{"account":"10722","platid":"67","email":"whj@whj.whj","gender":"male","nickname":"navy1125","timestamp":"12345","uid":"10722","sign":"%s"}}}{"do":"plat-token-login","data":{"gameid":170,"platinfo":{"account":"10722","platid":"67","email":"whj@whj.whj","gender":"male","nickname":"navy1125","timestamp":"12345","uid":"10722","sign":"%s"}}}{"do":"plat-token-login","data":{"gameid":170,"platinfo":{"account":"10722","platid":"67","email":"whj@whj.whj","gender":"male","nickname":"navy1125","timestamp":"12345","uid":"10722","sign":"%s"}}}{"do":"plat-token-login","data":{"gameid":170,"platinfo":{"account":"10722","platid":"67","email":"whj@whj.whj","gender":"male","nickname":"navy1125","timestamp":"12345","uid":"10722","sign":"%s"}}}`)
	//key1 := []byte("1234567890abcdef")
	key1 := []byte("12345678")
	//inData := []byte(`{"do":"request-zone-list","data":{"gameid":170},"gameid":170}`)
	inData := []byte(`{"do":"request-zone-list","data":{"gameid":170},"gameid":170}`)
	block, err := des.NewCipher(key1)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("cipher block size:", block.BlockSize())
	fmt.Println("aes block size:", aes.BlockSize)

	for i := 0; i < int(len(inData)/block.BlockSize()); i++ {
		inData = append(inData, 0)
	}
	blocknum := int(len(inData) / block.BlockSize())
	for i := 0; i < blocknum; i++ {
		block.Encrypt(inData[i*block.BlockSize():(i+1)*block.BlockSize()], inData[i*block.BlockSize():(i+1)*block.BlockSize()])
	}
	fmt.Println(string(inData))
	for i := 0; i < int(len(inData)/block.BlockSize()); i++ {
		inData[len(inData)-i-1] = 0
	}
	for i := 0; i < blocknum; i++ {
		block.Decrypt(inData[i*block.BlockSize():(i+1)*block.BlockSize()], inData[i*block.BlockSize():(i+1)*block.BlockSize()])
	}
	fmt.Println(string(inData))

	key := []byte("1234567890abcdef")
	ciphertext, _ := hex.DecodeString("f363f3ccdcb12bb883abf484ba77d9cd7d32b5baecb3d4b1b3e0e4beffdb3ded")

	block, err = aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	block.Encrypt(inData, inData)
	fmt.Println(string(inData))
	block.Decrypt(inData, inData)
	fmt.Println(string(inData))

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// CBC mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(ciphertext, ciphertext)

	// If the original plaintext lengths are not a multiple of the block
	// size, padding would have to be added when encrypting, which would be
	// removed at this point. For an example, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. However, it's
	// critical to note that ciphertexts must be authenticated (i.e. by
	// using crypto/hmac) before being decrypted in order to avoid creating
	// a padding oracle.

	fmt.Printf("%s\n", ciphertext)

}
