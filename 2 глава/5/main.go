package main

import (
	"fmt"
	"testing"
)

type BitCounter struct {
	table [256]byte
}

func NewBitCounter(table [256]byte) BitCounter {
	return BitCounter{table: table}
}

func (bc *BitCounter) CountBits(x uint64) int {
	count := 0
	for x > 0 {
		x = x & (x - 1)
		count++
	}
	return count
}

func BenchmarkBitCount(b *testing.B, bc *BitCounter, x uint64) {
	for i := 0; i < b.N; i++ {
		_ = bc.CountBits(x)
	}
}

func main() {
	popCountTable := [256]byte{}
	for i := range popCountTable {
		popCountTable[i] = popCountTable[i/2] + byte(i&1)
	}

	bitCounter := NewBitCounter(popCountTable)
	x := uint64(0x0123456789ABCDEF) // Пример входных данных

	// Используем метод Benchmark для оценки производительности
	result := testing.Benchmark(func(b *testing.B) {
		BenchmarkBitCount(b, &bitCounter, x)
	})

	// Выводим результаты бенчмарка
	fmt.Println(result)
}
