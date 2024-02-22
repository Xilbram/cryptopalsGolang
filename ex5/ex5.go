package main

//Este exercicio é mais simples que o anterior, consiste em iterar a chave e ir realizando xor com a mensagem 
//a ser criptografada. Tive uma diferença de resultados em 2 bits, e preciso debugar por iterações para ver o motivo

import (
	"math"
	"fmt"
)

func main(){
	messageToEncrypt := "Burning 'em, if you ain't quick and nimble I go crazy when I hear a cymbal"
	key := "ICE"
	messageToAchieve := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	result := ""

	for i := 0; i < len(messageToEncrypt); i++{
		currentByte := key[i%len(key)]
		xorRes := xor2Bytes(messageToEncrypt[i], currentByte) 
		result += fmt.Sprintf("%02x", xorRes)
	}

	fmt.Printf(result + "\n")
	fmt.Printf(messageToAchieve + "\n")
	equal, pos,char1,char2 := areEqual(result, messageToAchieve)
	if equal {
		fmt.Println("Iguais")
	} else {
		fmt.Printf("Diferentes em index %d e bytes %s %s\n", pos, char1, char2)
	}

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

func areEqual(str1, str2 string) (bool, int, byte, byte) {
	minLength := len(str1)
	if len(str2) < minLength {
		minLength = len(str2)
	}

	for i := 0; i < minLength; i++ {
		if str1[i] != str2[i] {
			return false, i, str1[i], str2[i]
		}
	}

	if len(str1) != len(str2) {
		return false, minLength, str1[0], str2[0]
	}

	return true, -1, str1[0], str2[0]
}
