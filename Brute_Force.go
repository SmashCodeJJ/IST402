package main

import (
	"fmt"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func encrypt(n int, plaintext string) string {
	var result strings.Builder

	for _, l := range plaintext {
		if index := strings.IndexRune(ALPHABET, l); index != -1 {
			i := (index + n) % 26
			result.WriteString(string(ALPHABET[i]))
		} else if index := strings.IndexRune(alphabet, l); index != -1 {
			i := (index + n) % 26
			result.WriteString(string(alphabet[i]))
		} else {
			result.WriteString(string(l))
		}
	}

	return result.String()
}

//func decrypt(n int, ciphertext string) string {
//	var result strings.Builder
//
//	for _, l := range ciphertext {
//		if index := strings.IndexRune(ALPHABET, l); index != -1 {
//			i := (index - n + 26) % 26
//			result.WriteString(string(ALPHABET[i]))
//		} else if index := strings.IndexRune(alphabet, l); index != -1 {
//			i := (index - n + 26) % 26
//			result.WriteString(string(alphabet[i]))
//		} else {
//			result.WriteString(string(l))
//		}
//	}

//return result.String()
//}

func main() {
	message := "Encryption Is An Interesting Topic"
	//key := 7

	for i := 0; i < 26; i++ {
		enc := encrypt(i, message)
		fmt.Printf("%d, %s\n", i, enc)

		//dec := decrypt(key, enc)
		//fmt.Printf("%d, %s\n", key, dec)
	}
}
