/*
   罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。
	字符          数值
	I             1
	V             5
	X             10
	L             50
	C             100
	D             500
	M             1000
 给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。
 romaToInt1 成功

 */
package main

import (
	"fmt"
)

func main() {
	fmt.Println(romaToInt1("IV")) // 4
	fmt.Println(romaToInt1("IX"))   // 9
	fmt.Println()
}

var romaMapper map[string]int= map[string]int{
	"I" : 1,
	"V" : 5,
	"X" : 10,
	"L" : 50,
	"C" : 100,
	"D" : 500,
	"M" : 1000,
}

func romaToInt1(roma string)int{
	var count = 0
	strLen := len(roma)
	for i:=0;i < strLen; i++ {
		countNum := romaMapper[string(roma[i])]
		fmt.Println(i,",countNum:",countNum)
		if i < strLen-1 {
			dumpNext := romaMapper[string(roma[i+1])]
			if dumpNext > countNum {
				count=count+(dumpNext-countNum)
				i++
				continue
			}
		}
		count=count+countNum
	}
	return count
}
