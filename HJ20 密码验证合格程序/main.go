package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		countA := 0
		counta := 0
		countNum := 0
		countOther := 0
		if len(text) > 8 && check(text) {
			for i := 0; i < len(text); i++ {
				if text[i] >= '0' && text[i] <= '9' {
					if countNum != 0 {
						continue
					}
					countNum++
					continue
				} else if text[i] >= 'A' && text[i] <= 'Z' {
					if countA != 0 {
						continue
					}
					countA++
					continue
				} else if text[i] >= 'a' && text[i] <= 'z' {
					if counta != 0 {
						continue
					}
					counta++
					continue
				} else {
					if countOther != 0 {
						continue
					}
					countOther++
					continue
				}
			}
			if counta+countA+countOther+countNum >= 3 {
				fmt.Println("OK")
			} else {
				fmt.Println("NG")
			}
		} else {
			fmt.Println("NG")
		}
	}
}

func check(str string) bool {
	m := make(map[string]int)
	for i := 0; i < len(str)-3; i++ {
		if _, ok := m[str[i:i+3]]; ok {
			return false
		} else {
			m[str[i:i+3]] = 1
		}
	}
	return true
}
