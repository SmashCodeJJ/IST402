package main

import (
	"bufio"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/chacha20"
	"io"
	"os"
)

func main() {
	// Get user input string
	fmt.Print("Enter a string to encrypt: ")
	input := getInput()

	// Generate a random 256-bit key
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		fmt.Println("Error generating key:", err)
		return
	}

	// Generate a random 96-bit nonce
	nonce := make([]byte, chacha20.NonceSize)
	_, err = rand.Read(nonce)
	if err != nil {
		fmt.Println("Error generating nonce:", err)
		return
	}

	// Encrypt the input string using ChaCha20
	ciphertext := encrypt([]byte(input), key, nonce)

	// Convert the key and nonce to hex strings for output
	keyStr := hex.EncodeToString(key)
	nonceStr := hex.EncodeToString(nonce)

	// Print the encrypted ciphertext, key, and nonce
	fmt.Println("Ciphertext:", hex.EncodeToString(ciphertext))
	fmt.Println("Key:", keyStr)
	fmt.Println("Nonce:", nonceStr)

	// Decrypt the ciphertext using ChaCha20
	plaintext := decrypt(ciphertext, key, nonce)

	// Print the decrypted plaintext
	fmt.Println("Decrypted plaintext:", string(plaintext))
}

func encrypt(plaintext []byte, key []byte, nonce []byte) []byte {
	// Create a new ChaCha20 cipher with the given key and nonce
	c, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		fmt.Println("Error creating cipher:", err)
		return nil
	}

	// Encrypt the plaintext using the cipher
	ciphertext := make([]byte, len(plaintext))
	c.XORKeyStream(ciphertext, plaintext)

	return ciphertext
}

func decrypt(ciphertext []byte, key []byte, nonce []byte) []byte {
	// Create a new ChaCha20 cipher with the given key and nonce
	c, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		fmt.Println("Error creating cipher:", err)
		return nil
	}

	// Decrypt the ciphertext using the cipher
	plaintext := make([]byte, len(ciphertext))
	c.XORKeyStream(plaintext, ciphertext)

	return plaintext
}

func getInput() string {
	// Read user input from stdin
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Println("Error reading input:", err)
		return ""
	}

	// Strip newline character from input
	input = input[:len(input)-1]

	return input
}
