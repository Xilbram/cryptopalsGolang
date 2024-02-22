package main

//Similar ao exercicio 3, aqui pensei em aplicar um filtro sobre filtro. Basicamente 
//para cada linha seleciono a iteração com a chave mais frequente. Tenho então uma série de linhas
//Dessa serie de linha busco por palavras chaves comuns, e adiciono as linhas encontradas
//Ao arquivo "possiveisResultados"

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
	fileDecriptacao, err := os.Create("decriptacao.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer fileDecriptacao.Close()

	filePossiveisResultados, err := os.Create("possiveisResultados.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer filePossiveisResultados.Close()

	file,err := os.Open("message.txt")
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()
	
	byteHolder := []uint8{}
	countLinhas := 0

	//Iterar linha a linha do arquivo
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		message := scanner.Text()
		hexMessage := utils.StringToHex(message)
		var textoEncontrado string
		var key byte
		maxScore := math.Inf(-1)

		//Iterar todos bytes como possiveis candidatos, realizar xor entre mensagem e byte, avaliar score
		for n := 0; n < 255; n++{
			for _,char := range hexMessage{
				xorResult := xor2Bytes(byte(n), char)
				byteHolder = append(byteHolder,xorResult)

				score := scoreText(byteHolder)
				if score > maxScore {
					maxScore = score
					textoEncontrado = string(byteHolder)
					key = byte(n)
				}
			}
			byteHolder = []uint8{}
		}

		countLinhas++
		countLinhasAsString := strconv.Itoa(countLinhas)
		strings.ToLower(textoEncontrado)
		words := strings.Fields(textoEncontrado)

		//Verificar se o texto da iteração contém uma palavra chave em ingles (sendo um bom indicativo de que a linha
		//decriptografada está em ingles, mas sujeito a ruídos)
		for _, palavra := range words {
			palavrasEncontradas := contains(string(palavra))
			if palavrasEncontradas > 0{
				filePossiveisResultados.WriteString("Linha:" + countLinhasAsString +
				 " Chave:" + string(key) + " Texto:" + textoEncontrado + "\n")
			}
		}
		fileDecriptacao.WriteString("Linha:" + countLinhasAsString + " Chave:" + string(key) + ": " + textoEncontrado + "\n")		
		if err := scanner.Err(); err != nil{log.Fatal(err)}
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

//Os valores do mapa de frequẽncia foram copiados da wikipedia
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

//Busca uma palavra chave em uma série de palavras comuns pré determinadas
func contains(palavra string) uint {
	commonWords := []string{"the", "and", "is", "in", "it", "to"}
	var palavrasEncontradas uint = 0 
	for _, item := range commonWords {
		if item == palavra {
			palavrasEncontradas++
		}
	}
	return palavrasEncontradas
}