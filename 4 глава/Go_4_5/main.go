package main

import "fmt"

func main() {
	s := []string{"asd", "asd", "dfg", "ghj"}
	s = uniq(s)
	fmt.Printf("%v\n", s)
}

func uniq(s []string) []string {
	u := s[:0]
	for i, v := range s {
		if i == 0 || v != s[i-1] {
			u = append(u, v)
		}
	}
	return s[:len(u)]
}
