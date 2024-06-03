package main

import (
	"fmt"
	"sort"
)

func main() {
	//lastStr()
	//strCount()
	HJ3()
}

// 明明的随机数
func HJ3() {
	var count int
	var textNum int
	_, err := fmt.Scan(&count)
	if err != nil {
		fmt.Println("数字错误")
		return
	}
	textArr := make([]int, 0)
	textMap := make(map[int]string)
	for i := 0; i < count; i++ {
		_, err := fmt.Scan(&textNum)
		if err != nil {
			fmt.Println("输入错误")
			return
		}
		textMap[textNum] = ""
	}
	for k, _ := range textMap {
		textArr = append(textArr, k)
	}
	sort.Ints(textArr)
	for _, value := range textArr {
		fmt.Println(value)
	}
}
