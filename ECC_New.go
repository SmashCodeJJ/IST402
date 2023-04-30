package main

import (
	"crypto/ecdsa"    // Provides the ECDSA algorithm for key generation.
	"crypto/elliptic" // Provides elliptic curve cryptography.
	"crypto/rand"     // Provides cryptographically secure random number generation.
	"crypto/sha256"   // Provides SHA256 for key derivation.
	"fmt"
	"golang.org/x/crypto/chacha20poly1305" // Provides ChaCha20-Poly1305 AEAD cipher for encryption and decryption.
)

// Encrypt function encrypts plaintext using the public key of the receiver.
func Encrypt(plaintext []byte, publicKey *ecdsa.PublicKey) ([]byte, error) {
	curve := publicKey.Curve // Get the elliptic curve from the public key.

	// Generate a random ephemeral key.
	ephemeralKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, err
	}

	// Perform ECDH key exchange.
	x, _ := curve.ScalarMult(publicKey.X, publicKey.Y, ephemeralKey.D.Bytes())

	// Derive a symmetric key from the shared secret.
	keyMaterial := sha256.Sum256(x.Bytes())
	// Initialize the AEAD cipher with the derived symmetric key.
	aead, err := chacha20poly1305.NewX(keyMaterial[:])
	if err != nil {
		return nil, err
	}

	// Generate a random nonce for the AEAD cipher.
	nonce := make([]byte, aead.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	// Encrypt the plaintext using the AEAD cipher.
	ciphertext := aead.Seal(nil, nonce, plaintext, nil)

	// Serialize the ephemeral public key.
	ephemeralPubKey := elliptic.Marshal(curve, ephemeralKey.PublicKey.X, ephemeralKey.PublicKey.Y)
	// Combine the serialized ephemeral public key, nonce, and ciphertext.
	return append(ephemeralPubKey, append(nonce, ciphertext...)...), nil
}

// Decrypt function decrypts the ciphertext using the private key of the receiver.
func Decrypt(ciphertext []byte, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	curve := privateKey.Curve
	params := curve.Params()

	// Split the input into ephemeral public key, nonce, and ciphertext.
	ephemeralPubKey, rest := ciphertext[:2*params.BitSize/8+1], ciphertext[2*params.BitSize/8+1:]
	// Deserialize the ephemeral public key.
	x, y := elliptic.Unmarshal(curve, ephemeralPubKey)

	// Perform ECDH key exchange.
	sharedX, _ := curve.ScalarMult(x, y, privateKey.D.Bytes())

	// Derive a symmetric key from the shared secret.
	keyMaterial := sha256.Sum256(sharedX.Bytes())
	// Initialize the AEAD cipher with the derived symmetric key.
	aead, err := chacha20poly1305.NewX(keyMaterial[:])
	if err != nil {
		return nil, err
	}

	// Extract the nonce and encrypted data from the input.
	nonce, ciphertextWithTag := rest[:aead.NonceSize()], rest[aead.NonceSize():]

	// Decrypt the ciphertext using the AEAD cipher.
	plaintext, err := aead.Open(nil, nonce, ciphertextWithTag, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func main() {
	// Generate an ECDSA private key.
	curve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		panic(err)
	}

	// Input string
	input := "This is a ECC assignment."
	fmt.Printf("Original text: %s\n", input)

	// Encrypt the input
	encrypted, err := Encrypt([]byte(input), &privateKey.PublicKey)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Encrypted text: %x\n", encrypted)

	// Decrypt the encrypted input
	decrypted, err := Decrypt(encrypted, privateKey)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decrypted text: %s\n", string(decrypted))
}
