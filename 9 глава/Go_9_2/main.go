package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

var initializeOnce sync.Once
var pc [256]byte

func initialize() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func popCount(x uint64) int {
	initializeOnce.Do(initialize)
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func main() {
	var wg sync.WaitGroup

	for _, str := range os.Args[1:] {
		str := str
		wg.Add(1)
		go func() {
			defer wg.Done()
			n, err := strconv.ParseUint(str, 10, 64)
			if err != nil {
				log.Print(err)
			} else {
				fmt.Printf("PopCount of %d = %d\n", n, popCount(n))
			}
		}()
	}
	wg.Wait()
}
