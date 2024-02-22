package main

//Consiste em iterar a string (já que são de mesmo tamanho), e realizar xor em seus bytes
import(
	"fmt"
	"cryptopals/utils"
	"math"
)

func main(){
	alvo := "746865206b696420646f6e277420706c6179"
	string1 := "1c0111001f010100061a024b53535009181c"
	string2 := "686974207468652062756c6c277320657965"
	hex1 := utils.StringToHex(string1)
	hex2 := utils.StringToHex(string2)
	
	var xorByte uint8
	ByteHolder := []uint8{}

	for i, b := range hex1{
		for j := 7; j >= 0; j-- {
			bit1 := (b >> uint(j)) & 1
			bit2 := (hex2[i] >> uint(j)) & 1

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

	output := utils.StringSliceToString(utils.ByteToHexString(ByteHolder))
	fmt.Printf(output + "\n")

	equal,pos := utils.AreEqual(output, alvo)
    if equal {
		fmt.Println("Iguais")
	} else {
		fmt.Printf("Diferentes em index %d e bytes \n", pos)
	}
}
