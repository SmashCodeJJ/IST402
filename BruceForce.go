package IST402

import "fmt"

func caesarCipher(plaintext string, shift int) string {
	ciphertext := ""

	for _, char := range plaintext {
		if char >= 'a' && char <= 'z' {
			ciphertext += string((char-'a'+rune(shift))%26 + 'a')
		} else if char >= 'A' && char <= 'Z' {
			ciphertext += string((char-'A'+rune(shift))%26 + 'A')
		} else {
			ciphertext += string(char)
		}
	}

	return ciphertext
}

func main() {
	plaintext := "hello world"
	shift := 3
	ciphertext := caesarCipher(plaintext, shift)
	fmt.Printf("Plaintext: %s\nShift: %d\nCiphertext: %s\n", plaintext, shift, ciphertext)
}
