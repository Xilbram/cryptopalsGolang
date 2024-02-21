package main

//Similar ao exercicio 3, aqui pensei em aplicar um filtro sobre filtro. Basicamente 
//para cada linha seleciono a iteração com a chave mais frequente. Tenho então uma série de linhas
//Dessa serie de linha busco por palavras chaves comuns, e adiciono as linhas encontradas
//Ao arquivo "possiveis"

import(
	"os"
	"log"
	"bufio"
	"math"
	"cryptopals/utils"
	"strconv"
	"strings"
)

func main(){
	fileEscrever, err := os.Create("resultado.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer fileEscrever.Close()

	filePossiveis, err := os.Create("possiveis.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer fileEscrever.Close()

	file,err := os.Open("message.txt")
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()
	
	byteHolder := []uint8{}
	var xorByte uint8
	count := 0
	commonWords := []string{"the", "and", "is", "in", "it", "to"}

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		message := scanner.Text()
		hexMessage := utils.StringToHex(message)
		var texto string
		var key byte
		maxScore := math.Inf(-1)

		for n := 0; n < 255; n++{
			numAsByte := byte(n)

			for _,char := range hexMessage{
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
		count++
		chaveString := string(key)
		countString := strconv.Itoa(count)
		strings.ToLower(texto)
		words := strings.Fields(texto)

		//Se o texto tiver alguma palavra muito comum, escrever eles nas possiveis. Dado a possibilidade
		//de que o texto pode não conter nenhuma delas, elas ainda são gravadas em Geral
		for _, palavra := range words {
			if contains(commonWords, string(palavra)) {
				filePossiveis.WriteString("Linha:" + countString + " Chave:" + chaveString + " " + texto + "\n")
			}
		}


		fileEscrever.WriteString("Linha:" + countString + " Chave:" + chaveString + " " + texto + "\n")		

		if err := scanner.Err(); err != nil{log.Fatal(err)}
	}
}


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

//Busca uma palavra chave em um slice de strings
func contains(palavrasChave []string, palavra string) bool {
	for _, item := range palavrasChave {
		if item == palavra {
			return true
		}
	}
	return false
}