package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	split := strings.Split(text, " ")
	atoi, err := strconv.Atoi(split[0])
	if err != nil {
		fmt.Println(err)
	}
	atoi2, err := strconv.Atoi(split[2])
	count := 0
	count += atoi2
	fmt.Println(count)
	atoi1, err := strconv.Atoi(split[1])
	if err != nil {
		fmt.Println(err)
	}
	atoi1--
	for i := 1; i <= atoi1; i++ {
		if i != 2 {
			count += 30
		}
		if i <= 7 && i%2 == 1 {
			count++
		}
		if i >= 8 && i%2 == 0 {
			count++
		}
	}
	fmt.Println(count)
	if atoi1 > 2 {
		if atoi%4 == 0 && atoi%100 != 0 || atoi%400 == 0 {
			count += 29
		} else {
			count += 28
		}
	}
	fmt.Println(count)

}
