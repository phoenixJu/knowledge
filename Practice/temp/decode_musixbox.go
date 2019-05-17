package main

import (
	"bytes"
	"crypto/des"
	"encoding/base64"
	"fmt"
	"net/url"
)
//params=3Z0J8Pw5BEpwRIASZdqRFkFo%2Fj%2FYzyT7yev1KExjlMcuYwpGpV8tc0hPwrYqfYHO1eCKXxQVkS0LA9%2F1JaEWYtaDMNo1CxYRgBGhdrWynQZz0tedZu1NftQwynvPcyIx&encSecKey=be845e002b057710ddb96ab2c7f2037630368122b3f63379d8d6d540e7274a2c65c4dcff4bc93bb6e3e2ba97289144c41cfdf879c866ad64e9b2b1d2ac31e1dc0762a854ea187b3d79a0ead37999348c4e10a306a79ba43165f23ccb11745af5f1881a7662855b1a6136be5d790484a9e701e76b71e29b3652680aa5e9af9932
var PASSWORD = []byte("be845e002b057710ddb96ab2c7f2037630368122b3f63379d8d6d540e7274a2c65c4dcff4bc93bb6e3e2ba97289144c41cfdf879c866ad64e9b2b1d2ac31e1dc0762a854ea187b3d79a0ead37999348c4e10a306a79ba43165f23ccb11745af5f1881a7662855b1a6136be5d790484a9e701e76b71e29b3652680aa5e9af9932")
var IV = []byte{0, 0, 0, 0, 0, 0, 0, 0}

func main() {
	//file11, _ := ioutil.ReadFile("/Users/zhuhongquan/GoProject/src/personal/executable/musicbox/register_request")
	encyptByte:= "3Z0J8Pw5BEpwRIASZdqRFkFo%2Fj%2FYzyT7yev1KExjlMcuYwpGpV8tc0hPwrYqfYHO1eCKXxQVkS0LA9%2F1JaEWYtaDMNo1CxYRgBGhdrWynQZz0tedZu1NftQwynvPcyIx&"
	result11 := Encrypt(encyptByte)
	fmt.Println("result:" + string(result11))

	//file, _ := ioutil.ReadFile("/Users/wuwangwen/Desktop/hotwords_response_plaintext")
	//result := Encrypt(string(file))
	//fmt.Println("result:" + string(result))
	//
	////file, _ := ioutil.ReadFile("/Users/wuwangwen/Desktop/getHotwords_response")
	//result1 := Decrypt(result)
	//fmt.Println("result:" + result1)
}


func Decrypt(src []byte) string {
	origData := DecryptDES_ECB(src, PASSWORD)
	decodeBytes, _ := base64.StdEncoding.DecodeString(string(origData))
	text1, _ := url.Parse(string(decodeBytes))
	text2, _ := url.Parse(text1.Path)
	result := text2.Path
	return result
}

func Encrypt(src string) []byte {
	text2 := url.QueryEscape(src)
	text1 := url.QueryEscape(text2)
	decodeString := base64.StdEncoding.EncodeToString([]byte(text1))
	origData := EncryptDES_ECB([]byte(decodeString), PASSWORD)
	return origData
}

func EncryptDES_ECB(data, keyByte []byte) []byte {
	block, err := des.NewCipher(keyByte)
	if err != nil {
		panic(err)
	}
	bs := block.BlockSize()
	data = PKCS5Padding(data, bs)
	if len(data)%bs != 0 {
		panic("Need a multiple of the blocksize")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return out
}

func DecryptDES_ECB(data, keyByte []byte) []byte {
	block, err := des.NewCipher(keyByte)
	if err != nil {
		panic(err)
	}
	bs := block.BlockSize()
	if len(data)%bs != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	out = PKCS5UnPadding(out)
	return out
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}



