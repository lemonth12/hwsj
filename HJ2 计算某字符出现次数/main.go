package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//lastStr()
	strCount()
}

// 写出一个程序，接受一个由字母、数字和空格组成的字符串，和一个字符，然后输出输入字符串中该字符的出现次数。（不区分大小写字母)
func strCount() {
	text := ""
	char := ""
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text = scanner.Text()
	} else {
		fmt.Println("错误1")
	}
	if scanner.Scan() {
		char = scanner.Text()
	} else {
		fmt.Println("错误2")
	}

	text = strings.ToUpper(text)
	char = strings.ToUpper(char)
	count := 0
	for _, v := range []byte(text) {
		if string(v) != " " && string(v) < "A" && string(v) > "Z" {
			fmt.Println("输入错误")
			return
		}
		if strings.EqualFold(string(v), char) {
			count += 1
		}
	}
	fmt.Println(count)
}
