package main

import (
	"fmt"
)

type Rotator struct {
	data []int
}

func NewRotator(data []int) *Rotator {
	return &Rotator{data: data}
}

func (r *Rotator) Rotate() {
	first := r.data[0]
	copy(r.data, r.data[1:])
	r.data[len(r.data)-1] = first
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	rotator := NewRotator(slice)

	rotator.Rotate()
	fmt.Println(rotator.data)
}
