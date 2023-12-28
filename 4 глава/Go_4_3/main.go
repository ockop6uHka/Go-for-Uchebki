package main

import "fmt"

type Reverser struct {
	data *[6]int
}

func NewReverser(data *[6]int) *Reverser {
	return &Reverser{data: data}
}

func (r *Reverser) Reverse() {
	for i, j := 0, len(*r.data)-1; i < j; i, j = i+1, j-1 {
		r.data[i], r.data[j] = r.data[j], r.data[i]
	}
}

func main() {
	array := [...]int{0, 1, 2, 3, 4, 5}
	reverser := NewReverser(&array)

	reverser.Reverse()
	fmt.Println(array)
}
