package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	fmt.Scanln(&num)
	itoa := strconv.Itoa(num)
	runes := []rune(itoa)

	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-i-1] = runes[len(runes)-i-1], runes[i]
	}
	m := make(map[rune]int)
	ints := make([]rune, 0)
	for _, r := range runes {
		_, ok := m[r]
		if !ok {
			m[r] = 1
			ints = append(ints, r)
		}
	}

	fmt.Println(string(ints))
}
