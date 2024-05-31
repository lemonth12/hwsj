package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Mg(arr1 []int) string {
	sort.Slice(arr1, func(i, j int) bool {
		return arr1[i] < arr1[j]
	})
	strArr := make([]string, 0)
	for _, v := range arr1 {
		itoa := strconv.Itoa(v)
		strArr = append(strArr, itoa)
	}
	join := strings.Join(strArr, "")
	return join
}

func main() {
	var num1 int
	fmt.Scan(&num1)
	arr1 := make([]int, 0)
	map1 := make(map[int]int)
	for i := 0; i < num1; i++ {
		var num2 int
		fmt.Scan(&num2)
		if _, ok := map1[num2]; !ok {
			arr1 = append(arr1, num2)
			map1[num2] = 1
		}
	}

	//var num1 int
	fmt.Scan(&num1)

	for i := 0; i < num1; i++ {
		var num2 int
		fmt.Scan(&num2)
		if _, ok := map1[num2]; !ok {
			arr1 = append(arr1, num2)
			map1[num2] = 1
		}
	}

	mg := Mg(arr1)
	fmt.Println(mg)
}
