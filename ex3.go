package main

//Primeira abordagem: Força bruta

import (
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main(){
	chars := `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789`
	message := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	hexMessage := stringToHex(message)
	uncoveredMessage := ""
	byteHolder := []uint8{}
	var xorByte uint8

	for _, char := range chars{
		charAsByte := runeToHex(char)

		for _, item := range hexMessage{
			for j, b := range charAsByte{
				bit := (item >> uint(j)) & 1
				charBit := (b >> uint(j)) & 1
				

				//Resultado da operação de xor é um bool, portanto deve ser convertido. O uso float64 é devido ao requisito de math.pow
				xorBool := (bit != charBit)
				var xorInt float64
				if xorBool {
					xorInt = 1
				}
				
				//Ao iterar cada bit do byte, optei por guardar os resultados em decimal em um uint8, então devo converter
				//os bits de acordo com sua casa decimal
				xorByte += uint8(xorInt * (math.Pow(2, float64(j))))

			}
			byteHolder = append(byteHolder,xorByte)
			xorByte = 0
			uncoveredMessage += byteSliceToString(byteHolder)
		}

		fmt.Printf("Caractere:" + string(char) + ":" + uncoveredMessage + "\n")
		uncoveredMessage = ""
		byteHolder = []uint8{}
		xorByte = 0
	
	}


}


func byteSliceToString(b []byte) string{
	res := ""
    for _, item := range b {
		res += string(item)
    }
	return res	
}

//Rune é equivalente a um int32
//Como no cenário temos de A...9, um unico byte é o suficiente
func runeToHex(r rune) []byte{
	charByte := byte(r)
	res := []byte{}

	for i := 7; i >= 0; i-- {
		bit := (charByte >> uint(i)) & 1
		res = append(res, bit)
	}
	return res
}

//Transforma uma string em cadeia hexadecimal
func stringToHex(s string) []byte {
    decodedHex, err := hex.DecodeString(s)
    if err != nil {
        fmt.Println("Error:", err)
    }

    return decodedHex
}

//Transforma uma cadeia de strings em um Slice
func stringToSlice(s string) []string {
    res := make([]string, len(s))

    for i, char := range s {
        newchar := strconv.QuoteRune(char)
        res[i] = newchar
    }
    return res
}

func stringSliceToString(s []string) string{
	res := ""
	for _, item := range s{
		res += item
	}
	return res
} 

//Transforma um slice de bytes em uma cadeia Hexadecimal
func byteToHexString(byteSlice []byte) []string{
	hexSlice := make([]string, len(byteSlice))
    for i, b := range byteSlice {
        hexSlice[i] = strings.ToLower(fmt.Sprintf("%02X", b))
    }

	return hexSlice
}