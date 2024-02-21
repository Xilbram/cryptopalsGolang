package main


import (
	"cryptopals/utils"
	"fmt"
	"math"
)

func main(){
	message := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	messageHex := utils.StringToHex(message)
	byteHolder := []uint8{}
	var xorByte uint8
	var texto string
	var key byte
	maxScore := math.Inf(-1)

	for n := 0; n < 255; n++{
		numAsByte := byte(n)

		for _, char := range messageHex{
			for j := 7; j >= 0 ; j--{
				bit1 := (numAsByte >> uint(j)) & 1
				bit2 := (char >> uint(j)) & 1

				xorResult := (bit1 != bit2)
				var xorNum float64; if xorResult {xorNum=1};

				xorByte += uint8(xorNum * (math.Pow(2, float64(j))))
			}
			byteHolder = append(byteHolder,xorByte)
			xorByte = 0

			score := scoreText(byteHolder)
			if score > maxScore {
				maxScore = score
				texto = string(byteHolder)
				key = byte(n)
			}
		} 

		byteHolder = []uint8{}
		xorByte = 0
	}

	fmt.Println("Texto: ", texto)
	fmt.Printf("Chave: %s \n", string(key))
}


//The values from the freqMap came from wikipedia, and are rounded
//O objetivo do scoreText aqui é evitar que eu precise ler todos os resultados
func scoreText(text []byte) float64 {
    
    freqMap := map[byte]float64{
        'a': 0.082, 'b': 0.015, 'c': 0.028, 'd': 0.043, 'e': 0.127,
        'f': 0.022, 'g': 0.020, 'h': 0.061, 'i': 0.07, 'j': 0.0015,
        'k': 0.0078, 'l': 0.040, 'm': 0.024, 'n': 0.0675, 'o': 0.075,
        'p': 0.019, 'q': 0.00095, 'r': 0.06, 's': 0.063, 't': 0.091,
        'u': 0.028, 'v': 0.0098, 'w': 0.024, 'x': 0.00150, 'y': 0.02,
        'z': 0.00074, ' ': 0.13000, // Espaço foi considerado também
    }

    score := 0.0
    for _, char := range text {
        score += freqMap[char]
    }
    return score
}