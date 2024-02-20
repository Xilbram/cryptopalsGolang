package main

import(
	"fmt"
	"encoding/hex"
	"math"
	"strings"
)

func main(){
	string1 := "1c0111001f010100061a024b53535009181c"
	string2 := "686974207468652062756c6c277320657965"
	hex1 := stringToHex(string1)
	hex2 := stringToHex(string2)
	
	var xorByte uint8
	ByteHolder := []uint8{}

	for i, b := range hex1{
		for j := 7; j >= 0; j-- {
			bit1 := (b >> uint(j)) & 1
			bit2 := (hex2[i] >> uint(j)) & 1

			//Resultado da operação de xor é um bool, portanto deve ser convertido. O uso float64 é devido ao requisito de math.pow
			xorBool := (bit1 != bit2)
			var xorInt float64
			if xorBool {
				xorInt = 1
			}
			
			//Ao iterar cada bit do byte, optei por guardar os resultados em decimal em um uint8, então devo converter
			//os bits de acordo com sua casa decimal
			xorByte += uint8(xorInt * (math.Pow(2, float64(j))))
		}
		ByteHolder = append(ByteHolder,xorByte)
		xorByte = 0
	}

	output := stringSliceToString(byteToHexString(ByteHolder))
	fmt.Printf(output)
}


//Transforma uma string em cadeia hexadecimal
func stringToHex(s string) []byte {
    decodedHex, err := hex.DecodeString(s)
    if err != nil {
        fmt.Println("Error:", err)
    }

    return decodedHex
}

//Transforma um slice de bytes em uma cadeia Hexadecimal
func byteToHexString(byteSlice []byte) []string{
	hexSlice := make([]string, len(byteSlice))
    for i, b := range byteSlice {
        hexSlice[i] = strings.ToLower(fmt.Sprintf("%02X", b))
    }

	return hexSlice
}

//Transforma um slice de strings em uma unica string
func stringSliceToString(s []string) string{
	res := ""
	for _, item := range s{
		res += item
	}
	return res
} 