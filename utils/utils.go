package utils

import (
	"bytes"
	"os"
	"strings"
)

/**
 * convert string to bytes
*/
func StringToBytes(str string) []byte {
	return bytes.NewBufferString(str).Bytes()
}

func Capitalize(word string) string {
	return strings.ToUpper(string(word[0])) + word[1:]
}

func ToTitleCase(str string) string {
	words := strings.Split(str, " ")
	for i := 0; i < len(words); i++ {
		words[i] = Capitalize(words[i])
	}
	return strings.Join(words, " ")
}

func Exit(code int){
	os.Exit(code)
}
