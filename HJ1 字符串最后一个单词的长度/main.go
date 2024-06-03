package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	lastStr()
}

// 字符串最后一个单词的长度
func lastStr() {
	var str string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		str = scanner.Text()
	} else {
		fmt.Println("输出错误")
		return
	}
	if len(str) < 0 || len(str) >= 5000 {
		fmt.Println("输入值不符合标准")
		return
	}
	split := strings.Split(str, " ")
	fmt.Println(len(split[len(split)-1]))
}
