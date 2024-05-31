package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	count := 0
	for i := 0; i < len(text); i++ {
		if text[i] >= 'A' && text[i] <= 'Z' {
			count++
		}
	}
	fmt.Println(count)
}
