package main

import (
	"crypto/sha256"
	"fmt"
)

var popCountTable [256]byte

func init() {
	for i := range popCountTable {
		popCountTable[i] = popCountTable[i/2] + byte(i&1)
	}
}

// BitDifference возвращает разницу в количестве различных битов между двумя хэшами.
func BitDifference(hashA, hashB [32]byte) int {
	diffCount := 0
	for i := range hashA {
		diffCount += int(popCountTable[hashA[i]^hashB[i]])
	}
	return diffCount
}

func main() {
	string1, string2 := "x", "x"
	hash1 := sha256.Sum256([]byte(string1))
	hash2 := sha256.Sum256([]byte(string2))

	fmt.Printf("Разница количества битов: %d\n", BitDifference(hash1, hash2))
}
