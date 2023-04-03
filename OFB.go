//package main
//
//import (
//	"fmt"
//	"math/rand"
//	"time"
//)
//
//func main() {
//	plaintextMatch := []string{
//		"a", "b", "c", "d", "e", "f",
//		"g", "h", "i", "j", "k", "l",
//		"m", "n", "o", "p", "q", "r",
//		"s", "t", "u", "v", "w", "x",
//		"y", "z"}
//
//	message := "attackatdawn"
//	var str string
//	fmt.Println("Input Message is:", message)
//	rand.Seed(time.Now().UnixNano())
//
//	for _, i := range message {
//		key := rand.Intn(26)
//		fmt.Println("The random KeyStream is:", key)
//		cipherTextNum := index(plaintextMatch, string(i)) + key
//		if cipherTextNum > 25 {
//			cipherTextNum -= 26
//		}
//
//		fmt.Println("The cipher Text Number is:", cipherTextNum)
//		str += plaintextMatch[cipherTextNum]
//		fmt.Println("The cipher Text Number corresponding letter is:", str)
//	}
//
//	fmt.Println(str)
//}
//
//func index(slice []string, value string) int {
//	for i, v := range slice {
//		if v == value {
//			return i
//		}
//	}
//	return -1
//}

package main

import (
	"crypto/aes"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func encryptECB(plaintext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, len(plaintext))
	for i := 0; i < len(plaintext); i += block.BlockSize() {
		block.Encrypt(ciphertext[i:i+block.BlockSize()], plaintext[i:i+block.BlockSize()])
	}

	return ciphertext, nil
}

func decryptECB(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	plaintext := make([]byte, len(ciphertext))
	for i := 0; i < len(ciphertext); i += block.BlockSize() {
		block.Decrypt(plaintext[i:i+block.BlockSize()], ciphertext[i:i+block.BlockSize()])
	}

	return plaintext, nil
}

func main() {
	// Generate a random AES key
	key := make([]byte, 16)
	if _, err := rand.Read(key); err != nil {
		panic(err)
	}

	// Plaintext message to encrypt
	plaintext := []byte("This is a secret message!")

	// Encrypt the message using ECB mode and the random key
	ciphertext, err := encryptECB(plaintext, key)
	if err != nil {
		panic(err)
	}

	// Print the ciphertext as a hex string
	fmt.Println("Ciphertext:", hex.EncodeToString(ciphertext))

	// Decrypt the ciphertext using ECB mode and the same key
	decrypted, err := decryptECB(ciphertext, key)
	if err != nil {
		panic(err)
	}

	// Print the decrypted plaintext
	fmt.Println("Decrypted plaintext:", string(decrypted))
}
