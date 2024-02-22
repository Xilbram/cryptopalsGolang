package main

//Segui o passo a passo para a resolução desta questão. Primeiro, estabeleci uma funcao para realizar a 
//distancia de hamming. Depois normalizei elas dividindo pelo keysize. Após decidir uma keysize
//é necessário quebrar o conteudo em blocos de tamanho keysize e transpor os blocos
//Após isso, é preciso resolver cada bloco com um algoritmo de xor e produzir um histogramas de chaves resultantes

import (
    "encoding/base64"
	"os"
	"bufio"
	"strconv"
	"log"
    "math"
)

func main() {
	file,err := os.Open("content.txt")
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()

	fileDecriptacao, err := os.Create("decriptacao.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer fileDecriptacao.Close()

	scanner := bufio.NewScanner(file)
	countLinhas := 0

	for scanner.Scan(){
		texto := []byte{}
		message := scanner.Text()
		msgBase64Decifrada,_ := base64.StdEncoding.DecodeString(message)
		keySize := guessKeySize(msgBase64Decifrada)
		key := breakRepeatingKeyXOR(msgBase64Decifrada, keySize)

		for i := range msgBase64Decifrada {
			texto[i] = msgBase64Decifrada[i] ^ key[i%len(key)]
		}

		countLinhas++
		countLinhasAsString := strconv.Itoa(countLinhas)
		fileDecriptacao.WriteString("Linha:" + countLinhasAsString + " Chave:" + string(key) + ": " + string(texto) + "\n")		

	}
}


//Estabelece a distancia entre 2 bytes (total de posições com diferença)
func hammingDistance(b1 []byte, b2 []byte) int {
    if len(b1) != len(b2) {
        return -1
    }

    distance := 0
    for i := range b1 {
		bit1 := (b1[i] >> uint(i)) & 1
		bit2 := (b2[i] >> uint(i)) & 1
		xorResult := (bit1 != bit2)
		var xorNum int; if xorResult {xorNum=1};

        for j := 0; j < 8; j++ {
            distance += int((xorNum >> j) & 1)
        }
    }
    return distance
}

//Normaliza a distancia 
func normalizedEditDistance(b []byte, keySize int) float64 {
    blocos := len(b) / keySize
    distancia := 0
    for i := 0; i < blocos-1; i++ {
        bloco1 := b[i*keySize : (i+1)*keySize]
        bloco2 := b[(i+1)*keySize : (i+2)*keySize]
        distancia += hammingDistance(bloco1, bloco2)
    }
    return float64(distancia) / float64(keySize*blocos)
}

//Deduz o tamanho da key de acordo com a sugestao 2<=key<=40
func guessKeySize(data []byte) int {
    minKey := 2
    maxKey := 40
    minDistance := math.MaxFloat64
    bestKeySize := 0
    for keySize := minKey; keySize <= maxKey; keySize++ {
        distance := normalizedEditDistance(data, keySize)
        if distance < minDistance {
            minDistance = distance
            bestKeySize = keySize
        }
    }
    return bestKeySize
}

//Reorganiza os blocos pra que cada byte forme um novo bloco
func transpose(blocks [][]byte) [][]byte {
    transposed := make([][]byte, len(blocks[0]))
    for i := range transposed {
        transposed[i] = make([]byte, len(blocks))
        for j := range transposed[i] {
            transposed[i][j] = blocks[j][i]
        }
    }
    return transposed
}


//Descobri há pouco tempo que é possivel fazer XOR de um jeito mais simples
//Realiza xor entre SLice de bytes e um byte
func singleByteXOR(bSlice []byte, key byte) []byte {
    res := make([]byte, len(bSlice))
    for i, b := range bSlice {
        res[i] = b ^ key
    }
    return res
}


//Quebra a encriptação usando repetição de xor com uma unica cadeia como contraposição
func breakRepeatingKeyXOR(msgDescobrir []byte, keySize int) []byte {
    transposed := make([][]byte, keySize)
	
    for i := range transposed {
        transposed[i] = make([]byte, len(msgDescobrir)/keySize)
    }

	for i, c := range msgDescobrir {
        blockIndex := i / keySize
        byteIndex := i % keySize
        transposed[byteIndex][blockIndex] = c
    }
    key := make([]byte, keySize)
    for i, block := range transposed {
        _, key[i] = breakSingleByteXOR(block)
    }
    return key
}

//Calcula o melhor score para um texto que possivelmente esteja em ingles
func breakSingleByteXOR(ciphertext []byte) (bestScore float64, key byte) {
    bestScore = -1
    for k := 0; k < 256; k++ {
        plaintext := singleByteXOR(ciphertext, byte(k))
        score := scoreText(plaintext)
        if score > bestScore {
            bestScore = score
            key = byte(k)
        }
    }
    return bestScore, key
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


