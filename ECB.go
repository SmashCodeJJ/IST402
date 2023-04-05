package main

import (
	"fmt"
)

/* an array with 4 rows and 2 columns*/
var codebook = [4][2]int{{0b00, 0b01}, {0b01, 0b10}, {0b10, 0b11}, {0b11, 0b00}}
var message = [4]int{0b00, 0b01, 0b10, 0b11}

var codebook2 = [4][2]int{{0b00, 0b010}, {0b01, 0b00}, {0b10, 0b01}, {0b11, 0b11}}

func codebookLookup(xor int) (lookupValue int) {
	var i, j int = 0, 0
	for i = 0; i < 4; i++ {
		if codebook[i][j] == xor {
			j++
			lookupValue = codebook[i][j]
			break
		}
	}
	return lookupValue
}

func codebookReverseLookup(value int) (lookupValue int) {
	var i, j int = 0, 1
	for i = 0; i < 4; i++ {
		if codebook[i][j] == value {
			j--
			lookupValue = codebook[i][j]
			break
		}
	}
	return lookupValue
}

func main() {

	fmt.Println("Now in the code book page 1.")
	var i int = 0
	var lookupValue int = 0
	for i = 0; i < len(message); i++ {
		lookupValue = codebookLookup(message[i])
		fmt.Printf("The ciphered value of block %b is %b\n", message[i], lookupValue)
	}

	fmt.Println("Decrypting...")

	for i = 0; i < len(message); i++ {
		decodedValue := codebookReverseLookup(codebook2[i][1])
		fmt.Printf("The decrypted value of block %b is %b\n", codebook2[i][0], decodedValue)
	}
}
