package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

// CompactSpaces преобразует последовательности пробелов и табуляций в одиночные пробелы.
func CompactSpaces(inputBytes []byte) []byte {
	outputBytes := inputBytes[:0] // Используем срез, чтобы избежать копирования исходных данных

	var lastRune rune

	for i := 0; i < len(inputBytes); {
		currentRune, size := utf8.DecodeRune(inputBytes[i:])

		if !unicode.IsSpace(currentRune) {
			outputBytes = append(outputBytes, inputBytes[i:i+size]...)
		} else if unicode.IsSpace(currentRune) && !unicode.IsSpace(lastRune) {
			outputBytes = append(outputBytes, ' ')
		}

		lastRune = currentRune
		i += size
	}

	return outputBytes
}

func main() {
	input := []byte("   Hello, \t   World!   ")
	fmt.Printf("Изначальная: %s\n", input)

	result := CompactSpaces(input)

	fmt.Printf("Преобразованная: %s\n", result)
}
