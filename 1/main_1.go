package main

import (
	"fmt"
	"crypto/sha256"
	"encoding/hex"
	"crypto/rand"
)

func main() {
	decimal := 12
	octal := 0o34
	hex := 0x56

	float := 7.8
	str := "string"
	boolean := true
	compl := complex64(9)

	getType(decimal)
	getType(octal)
	getType(hex)
	getType(float)
	getType(str)
	getType(boolean)
	getType(compl)

	oneStr := createString(decimal, octal, hex, float, str, boolean, compl)

	convertToRune(oneStr)

	hasher := Hasher{}

	salt := generateSalt(32)
	result := hasher.HashStringSHA256(str, salt)

	fmt.Println(result)

}

func getType(arg interface{}) {
	fmt.Printf("%T\n", arg)
}

func createString(decimal, octal, hex int, args ...interface{}) string {
	str := fmt.Sprintf("%d%o%X", decimal, octal, hex)
	for _, arg := range args {
		switch arg.(type) {
		case float64:
			str += fmt.Sprintf("%f", arg.(float64))
		case string:
			str += arg.(string)
		case bool:
			str += fmt.Sprintf("%t", arg.(bool))
		case complex64:
			str += fmt.Sprintf("%g", arg.(complex64))
		}
	}

	fmt.Println(str)
	return str
}

func convertToRune(str string) []rune {
	return []rune(str)
}

type Hasher struct {
}

func (h *Hasher) HashStringSHA256(str, salt string) string {
	hasher := sha256.New()
	hasher.Write([]byte(str + salt))
	return hex.EncodeToString(hasher.Sum(nil))
}

func generateSalt(length int) string {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(salt)
}
