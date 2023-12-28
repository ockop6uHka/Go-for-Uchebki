package main

import (
	"fmt"
	"unicode/utf8"
)

// ReverseUTF8Bytes обращает порядок байт в каждом руне UTF-8 в байтовом срезе.
func ReverseUTF8Bytes(bytes []byte) []byte {
	for i := 0; i < len(bytes); {
		_, size := utf8.DecodeRune(bytes[i:])
		reverseBytes(bytes[i : i+size])
		i += size
	}
	reverseBytes(bytes)
	return bytes
}

// reverseBytes обращает порядок байт в байтовом срезе.
func reverseBytes(bytes []byte) []byte {
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return bytes
}

func main() {
	inputString := "Hello World!"
	reversedString := string(ReverseUTF8Bytes([]byte(inputString)))

	fmt.Println(reversedString)
}
