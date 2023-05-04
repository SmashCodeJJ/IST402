package main

import (
	"fmt"
	"strings"
)

// Rotor is a struct representing a single rotor in the Enigma machine.
type Rotor struct {
	shift int
}

// NewRotor creates a new Rotor with the given initial shift.
func NewRotor(shift int) *Rotor {
	return &Rotor{shift: shift}
}

// Rotate advances the rotor by one position, wrapping around after 26 positions.
func (r *Rotor) Rotate() {
	r.shift = (r.shift + 1) % 26
}

// ShiftChar applies the rotor's shift to the given character. If forward is true, it
// shifts the character forward; if false, it shifts the character backward.
func (r *Rotor) ShiftChar(c byte, forward bool) byte {
	if forward {
		return byte((int(c-'A')+r.shift)%26 + 'A')
	}
	return byte((int(c-'A')-r.shift+26)%26 + 'A')
}

// Enigma is a struct representing the simplified Enigma machine.
type Enigma struct {
	rotors []*Rotor
}

// NewEnigma creates a new Enigma machine with the given rotor shifts.
func NewEnigma(rotorShifts []int) *Enigma {
	rotors := make([]*Rotor, len(rotorShifts))
	for i, shift := range rotorShifts {
		rotors[i] = NewRotor(shift)
	}
	return &Enigma{rotors: rotors}
}

// Encrypt encrypts the given text using the Enigma machine.
func (e *Enigma) Encrypt(text string) string {
	upperText := strings.ToUpper(text)
	encrypted := make([]byte, len(upperText))

	// Process each character in the input text
	for i, c := range upperText {
		// Only process alphabetic characters
		if c >= 'A' && c <= 'Z' {
			// Apply each rotor's shift to the character, moving forward
			for _, rotor := range e.rotors {
				c = rune(rotor.ShiftChar(byte(c), true))
			}
			// Advance the rotors after processing the character
			e.advanceRotors()
		}
		encrypted[i] = byte(c) // Store the processed character
	}

	return string(encrypted)
}

// Decrypt decrypts the given text using the Enigma machine.
func (e *Enigma) Decrypt(text string) string {
	upperText := strings.ToUpper(text)
	decrypted := make([]byte, len(upperText))

	// Process each character in the input text
	for i, c := range upperText {
		// Only process alphabetic characters
		if c >= 'A' && c <= 'Z' {
			// Apply each rotor's shift to the character, moving backward
			for j := len(e.rotors) - 1; j >= 0; j-- {
				c = rune(e.rotors[j].ShiftChar(byte(c), false))
			}
			// Advance the rotors after processing the character
			e.advanceRotors()
		}
		decrypted[i] = byte(c)
	}

	return string(decrypted)
}

// advanceRotors advances the rotors in the Enigma machine.
func (e *Enigma) advanceRotors() {
	for i := 0; i < len(e.rotors); i++ {
		e.rotors[i].Rotate()
		if e.rotors[i].shift != 0 {
			break
		}
	}
}

// ResetRotors resets the rotors to their original positions.
func (e *Enigma) ResetRotors(rotorShifts []int) {
	for i, shift := range rotorShifts {
		e.rotors[i].shift = shift
	}
}

func main() {
	rotorShifts := []int{1, 3, 5}
	enigma := NewEnigma(rotorShifts)

	plaintext := "This is personalized Enigma machine using Caesar cipher."
	encrypted := enigma.Encrypt(plaintext)
	enigma.ResetRotors(rotorShifts) // Reset rotors before decryption
	decrypted := enigma.Decrypt(encrypted)

	fmt.Printf("Plaintext: %s\n", plaintext)
	fmt.Printf("Encrypted: %s\n", encrypted)
	fmt.Printf("Decrypted: %s\n", decrypted)
}
