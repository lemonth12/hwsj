package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		text := reader.Text()
		i, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			fmt.Println(err, "cuowu")
		}
		formatInt := strconv.FormatInt(i, 2)
		count := 0
		//fmt.Println(formatInt)
		for i := 0; i < len(formatInt); i++ {
			if string(formatInt[i]) == "1" {
				count++
			}
		}
		fmt.Println(count)
	}

}
