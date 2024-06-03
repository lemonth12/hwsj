package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 编写一个程序，将输入字符串中的字符按如下规则排序。
// 规则 1 ：英文字母从 A 到 Z 排列，不区分大小写。
// 如，输入： Type 输出： epTy
// 规则 2 ：同一个英文字母的大小写同时存在时，按照输入顺序排列。
// 如，输入： BabA 输出： aABb
// 规则 3 ：非英文字母的其它字符保持原来的位置。
// 如，输入： By?e 输出： Be?y
// 数据范围：输入的字符串长度满足
// 1≤n≤1000

type sort struct {
	str string
	num int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		sorts := make([]sort, 0)
		var s sort
		for i := 0; i < len(text); i++ {
			s.num = i
			s.str = string(text[i])
			sorts = append(sorts, s)
		}
		//min := sorts[0].str
		for i := 0; i < len(sorts)-1; i++ {
			//fmt.Println(sorts[i].str, sorts[i].str > 'z', "222222222aa22")
			if (sorts[i].str > "z" || sorts[i].str < "a") && (sorts[i].str > "Z" || sorts[i].str < "A") {
				//fmt.Println("1")
				continue
			}
			for j := i + 1; j < len(sorts); j++ {
				//fmt.Println(sorts[j].str, (sorts[j].str > 'z' || sorts[j].str < 'a'), (sorts[j].str > '9' && sorts[j].str < '0'), (sorts[j].str > 'Z' && sorts[j].str < 'A'))
				if (sorts[j].str > "z" || sorts[j].str < "a") && (sorts[j].str > "Z" || sorts[j].str < "A") {
					//fmt.Println("12")
					continue
				}
				if strings.ToLower(string(sorts[i].str)) > strings.ToLower(string(sorts[j].str)) {
					sorts[i], sorts[j] = sorts[j], sorts[i]
				}
				if strings.ToLower(string(sorts[i].str)) == strings.ToLower(string(sorts[j].str)) {
					if sorts[i].num > sorts[j].num {
						sorts[i], sorts[j] = sorts[j], sorts[i]
					}
				}
			}
		}

		str := ""
		for _, v := range sorts {
			str += v.str
		}
		fmt.Println(str)
	}
}
