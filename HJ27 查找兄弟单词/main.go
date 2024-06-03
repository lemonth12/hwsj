package main

import (
	"fmt"
	"strings"
)

//
//输入描述：
//输入只有一行。 先输入字典中单词的个数n，再输入n个单词作为字典单词。 然后输入一个单词x 最后后输入一个整数k
//输出描述：
//第一行输出查找到x的兄弟单词的个数m 第二行输出查找到的按照字典顺序排序后的第k个兄弟单词，没有符合第k个的话则不用输出
//
//输入：
//6 cab ad abcd cba abc bca abc 1
//输出：
//3
//bca

func main() {
	var num int
	fmt.Scan(&num)
	strs := make([]string, 0)
	for i := 0; i < num; i++ {
		var str string
		fmt.Scan(&str)
		strs = append(strs, str)
	}
	var strm string
	fmt.Scan(&strm)

	var k int
	fmt.Scan(&k)

	count := 0
	arr := make([]string, 0)
	for i := 0; i < len(strs); i++ {
		flat := false

		if strs[i] != strm && len(strs[i]) == len(strm) {
			fmt.Println(strs[i], strm)
			for i := 0; i < len(strm); i++ {
				contains := strings.Contains(strs[i], string(strm[i]))
				if !contains {
					flat = true
				}
			}
		}
		if !flat {
			arr = append(arr, strs[i])
			fmt.Scan(strs[i], count)
			count++
		}
	}
	if k < len(strs) {
		fmt.Println(count, strs[k])
	} else {
		fmt.Println(count)
	}

	//scanner := bufio.NewScanner(os.Stdin)
	//for scanner.Scan() {
	//	text := scanner.Text()
	//}
}
