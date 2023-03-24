package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}
var MySecret = "abc&1*~#^2^#s0^=)^^7%b34"

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

// Add Pad and Unpad functions to a plainttext message
// Allowing message length to block size of byte

func Pad(data []byte, blockSize int) []byte {
	padLen := blockSize - len(data)%blockSize
	padding := make([]byte, padLen)
	for i := 0; i < padLen; i++ {
		padding[i] = byte(padLen)
	}
	return append(data, padding...)
}

func Unpad(data []byte) []byte {
	padLen := int(data[len(data)-1])
	return data[:len(data)-padLen]
}

func ECBEncrypt(data []byte, key []byte) ([]byte, error) {
	// creates a new AES cipher using the given key
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	paddedData := Pad(data, blockSize) // multiple of the block size for message length
	encrypted := make([]byte, len(paddedData))
	// loop each of block of plaintext
	for i := 0; i < len(paddedData); i += blockSize {
		block.Encrypt(encrypted[i:i+blockSize], paddedData[i:i+blockSize])
	}
	return encrypted, nil
}

func ECBDecrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	decrypted := make([]byte, len(data))
	// loop through ciphertext
	for i := 0; i < len(data); i += blockSize {
		block.Decrypt(decrypted[i:i+blockSize], data[i:i+blockSize])
	}
	return Unpad(decrypted), nil // unpad is to remove padding added to plaintext during encode
}

func OFBEncrypt(text, MySecret string) (string, error) {
	//Create the AES block
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	//plaintext must be a byte format
	plainText := []byte(text)
	ofb := cipher.NewOFB(block, bytes)
	cipherText := make([]byte, len(plainText))
	ofb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}

// Decrypt method is to extract back the encrypted text
func OFBDecrypt(text, MySecret string) (string, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	cipherText := Decode(text)
	ofb := cipher.NewOFB(block, bytes)
	plainText := make([]byte, len(cipherText))
	ofb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}

func main() {
	fmt.Println("Enter the string to encrypt: ")
	var StringToEncrypt string
	fmt.Scanln(&StringToEncrypt)
	ECBencText, err := ECBEncrypt([]byte(StringToEncrypt), []byte(MySecret))
	OFBencText, err := OFBEncrypt(StringToEncrypt, MySecret)
	if err != nil {
		fmt.Println("error encrypting your classified text: ", err)
	}
	fmt.Println("ECB Encoding: ", Encode(ECBencText))
	fmt.Println("OFB Encoding", OFBencText)
	ECBdecText, err := ECBDecrypt(ECBencText, []byte(MySecret))
	OFBdecText, err := OFBDecrypt(OFBencText, MySecret)
	if err != nil {
		fmt.Println("error decrypting your encrypted text: ", err)
	}
	fmt.Println("ECB Decoding: ", string(ECBdecText))
	fmt.Println("OFB Decoding: ", OFBdecText)
}
