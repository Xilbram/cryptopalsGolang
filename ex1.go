package main	

import (
	"fmt"
    "strings"
    "cryptopals/utils"
)
//Embora existam bibliotecas para conversão direta, decidi aproveitar
//o desafio pra aprender algo novo.
//Pegar 64 caractereces da base64 -> Transformar em Slice (similar a um array) ->
//Obter a mensagem em Hexadecimal -> Transformar em um Slice de Bytes ->
//Iterar slice de bytes -> Pegar bit a bit em inseri-los em uma lista contendo 6 bits ->
//Realizar a correspondência entre os 6 bits e um valor decimal ([0,63])
//Obter caracterece de base64 de acordo com o index fornecido pelo valor decimal


func main() {
    base64Chars := `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/`
    alvo := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
    base64SliceIndex := utils.StringToSlice(base64Chars)
    hexstr := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
    hexArr := utils.StringToHex(hexstr)
    bitHolder := []byte{}
    decodedStr := ""

	//Pega-se o byte, e itera bit a bit (0 até 7)
    for _, b := range hexArr {
        for i := 7; i >= 0; i-- {
            bit := (b >> uint(i)) & 1
            bitHolder = append(bitHolder, bit)

			//Ao obter 6 bits, transformar cadeia em decimal e obter caracterece de base64 através do decimal
            if len(bitHolder) == 6 {
                dec := utils.BinToDec(bitHolder)
                strToAppend := base64SliceIndex[dec]
                decodedStr += strToAppend
                bitHolder = []byte{}
            }
        }
    }

    //fmt.Printf("%08b", hexArr)
    parts := strings.Split(strings.ReplaceAll(decodedStr, "''", ""), "'")
    resultCorrigido := strings.Join(parts, "")

    equal,pos := utils.AreEqual(resultCorrigido, alvo)
    if equal {
		fmt.Println("Iguais")
	} else {
		fmt.Printf("Diferentes em index %d e bytes \n", pos)
	}
    fmt.Printf(resultCorrigido)
}
