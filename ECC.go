package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto/ecies"
)

func main() {
	// Get user input
	var input string
	fmt.Print("Enter text to encrypt: ")
	fmt.Scanln(&input)

	// Generate an ECC key pair
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	// Convert the private key to an ECIES public key
	eciesPublicKey := ecies.ImportECDSAPublic(&privateKey.PublicKey)

	// Encrypt the message
	ciphertext, err := ecies.Encrypt(rand.Reader, eciesPublicKey, []byte(input), nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Encrypted message: %x\n", ciphertext)

	// Convert the private key to an ECIES private key
	eciesPrivateKey := ecies.ImportECDSA(privateKey)

	// Decrypt the message
	plaintext, err := eciesPrivateKey.Decrypt(ciphertext, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Decrypted message: %s\n", string(plaintext))
}
