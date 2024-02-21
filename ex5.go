package main

import (
	"math"
	"fmt"
)

func main(){
	messageToEncrypt := "Burning 'em, if you ain't quick and nimble I go crazy when I hear a cymbal"
	key := "ICE"
	result := ""

	for i := 0; i < len(messageToEncrypt); i++{
		currentByte := key[i%len(key)]
		xorRes := xor2Bytes(messageToEncrypt[i], currentByte) 
		result += fmt.Sprintf("%02x", xorRes)
	}
	fmt.Printf(result)
}

//Recebe 2 bytes e realiza uma operação xor entre eles
func xor2Bytes(b1 byte, b2 byte) uint8{
	var xorByte uint8 = 0
	for j := 7; j >= 0 ; j--{
		bit1 := (b1 >> uint(j)) & 1
		bit2 := (b2 >> uint(j)) & 1
		xorResult := (bit1 != bit2)
		var xorNum float64; if xorResult {xorNum=1};
		xorByte += uint8(xorNum * (math.Pow(2, float64(j))))
	}
	return xorByte
}