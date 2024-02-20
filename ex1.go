package main

import (
    "encoding/hex"
    "fmt"
    "math"
    "strconv"
)

//Pegar 64 caractereces da base64 -> Transformar em Slice (similar a um array) ->
//Obter a mensagem em Hexadecimal -> Transformar em um Slice de Bytes ->
//Iterar slice de bytes -> Pegar bit a bit em inseri-los em uma lista contendo 6 bits ->
//Realizar a correspondência entre os 6 bits e um valor decimal ([0,63])
//Obter caracterece de base64 de acordo com o index fornecido pelo valor decimal


func main() {
    base64Chars := `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/`
    base64SliceIndex := strToSlice(base64Chars)
    hexstr := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
    hexArr := strToHex(hexstr)

    bitHolder := []byte{}
    decodedStr := ""

	//Pega-se o byte, e itera bit a bit (0 até 7)
    for _, b := range hexArr {
        for i := 7; i >= 0; i-- {
            bit := (b >> uint(i)) & 1
            bitHolder = append(bitHolder, bit)

			//Ao obter 6 bits, transformar cadeia em decimal e obter caracterece de base64 através do decimal
            if len(bitHolder) == 6 {
                dec := binToDec(bitHolder)
                strToAppend := base64SliceIndex[dec]
                decodedStr += strToAppend
                bitHolder = []byte{}
            }
        }
    }

    //fmt.Printf("%08b", hexArr)
    fmt.Printf(decodedStr)
}


//Transforma uma string em seu valor hexadecimal
func strToHex(s string) []byte {
    decodedHex, err := hex.DecodeString(s)
    if err != nil {
        fmt.Println("Error:", err)
    }

    return decodedHex
}

//Converte um valor binário em decimal
func binToDec(b []byte) int {
    sum := 0
    for i := 0; i <= 5; i++ {
        var1 := int(b[i])
        var2 := int(math.Pow(2, float64(5-i)))
        sum += var1 * var2
    }

    return sum
}

//Transforma uma cadeia de strings em um Slice
func strToSlice(s string) []string {
    res := make([]string, len(s))

    for i, char := range s {
        newchar := strconv.QuoteRune(char)
        res[i] = newchar
    }
    return res
}