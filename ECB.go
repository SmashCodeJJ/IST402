package main

import (
	"fmt"
)

/* an array with 4 rows and 2 columns*/
var codebook = [4][2]int{{0b00, 0b01}, {0b01, 0b10}, {0b10, 0b11}, {0b11, 0b00}}
var message = [4]int{0b00, 0b01, 0b10, 0b11}

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

func main() {
	var x, i int = 0, 0
	var xor int = 0
	var lookupValue int = 0
	for i = 0; i < 4; i++ {
		xor = message[x] // use message block as-is without IV or previous block's ciphertext
		lookupValue = codebookLookup(xor)
		x++
		fmt.Printf("The ciphered value of a is %b\n", lookupValue)
	}
}
