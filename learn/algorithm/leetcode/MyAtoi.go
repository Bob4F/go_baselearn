/**
	请你来实现一个 atoi 函数，使其能将字符串转换成整数。
首先，该函数会根据需要丢弃无用的开头空格字符，直到寻找到第一个非空格的字符为止
当我们寻找到的第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字组合起来，作为该整数的正负号；
假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成整数。
该字符串除了有效的整数部分之后也可能会存在多余的字符，这些字符可以被忽略，它们对于函数不应该造成影响。
注意：假如该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，则你的函数不需要进行转换。
在任何情况下，若函数不能进行有效的转换时，请返回 0。
说明：
假设我们的环境只能存储 32 位大小的有符号整数，那么其数值范围为 [−231,  231 − 1]。如果数值超过这个范围，请返回  INT_MAX (231 − 1) 或 INT_MIN (−231)
*/
package main

import (
_ "fmt"
"math"
"regexp"
"strconv"
"strings"
)

func main() {
	println(atoi("Hello World,3www7kkk8998"))
	println(atoi("  -42"))
}

/*
 	执行用时 :4 ～ 16 ms
    消耗内存 : 6.5~MB
*/
func atoi(str string) int {
	str = strings.TrimLeft(str, " ")
	println(str)
	if str=="" {
		return 0
	}
	sign,i := 1,0
	flysnowRegexp := regexp.MustCompile(`\d`) // 正则工具
	// 取第一个字符
	char0 := str[0]
	p := flysnowRegexp.FindStringSubmatch(string(char0))
	if char0 == 43 { // 第一个字符为 43 ('+')，表示正数，sign设置为1
		sign = 1
	} else if char0 == 45 { // 第一个字符为 45 ('-')，表示正数，sign设置为-1
		sign = -1
	} else if len(p) > 0 { // 第一个字符为数字，设置初始值
		result,_:=strconv.Atoi(p[0])
		i = result
	} else { // 否则说明第一个字符不符合要求，直接返回 0
		return 0
	}
	for _,s:= range str[1:] {
		// 相比用正则表达式匹配 很明显直接比对 ASCII 更快
		p := flysnowRegexp.FindStringSubmatch(string(s))
		if len(p) > 0 {
			if i > math.MaxInt32/10 || (i == math.MaxInt32/10 && (int(s)-48) > 7) {
				if sign == 1 {
					return math.MaxInt32
				}
				// 否则返回最小值
				return math.MinInt32
			}
			result,_:=strconv.Atoi(p[0])
			i = (i*10) + result
		}else {
			break
		}
	}
	return i*sign
}


/*
  通过案例 - 成功案例

 	执行用时 :0 ～ 4 ms
    消耗内存 : 2.5~MB
 */
func myAtoi(str string) int {
	// sign 表示符号
	// i 表示结果的无符号数
	sign, i := 1, 0
	// 去除左侧空格
	trimStr := strings.TrimLeft(str, " ")
	if trimStr == "" {
		return 0
	}
	// 取第一个字符
	char0 := trimStr[0]
	if char0 == 43 { // 第一个字符为 43 ('+')，表示正数，sign设置为1
		// sign = 1
	} else if char0 == 45 { // 第一个字符为 45 ('-')，表示正数，sign设置为-1
		sign = -1
	} else if char0 > 47 && char0 < 58 { // 第一个字符为数字，设置初始值
		i = int(char0) - 48
	} else { // 否则说明第一个字符不符合要求，直接返回 0
		return 0
	}
	// 开始循环遍历剩余的字符
	for _, s := range trimStr[1:] {
		if s > 47 && s < 58 { // 此时每次遍历到的字符都应该是数字
			// 判断溢出，根据第七题的判断方法来实现，由于i是无符号的，直接判断最大值即可
			if i > math.MaxInt32/10 || (i == math.MaxInt32/10 && (int(s)-48) > 7) {
				// 如果sign为1表示正数，返回最大值
				if sign == 1 {
					return math.MaxInt32
				}
				// 否则返回最小值
				return math.MinInt32
			}
			// 溢出检查通过，通过第七题推入数字方法来更新结果值
			i = i*10 + (int(s) - 48)
		} else { // 不是数字直接结束循环
			break
		}
	}
	// 循环正常结束，返回结果值
	return i * sign
}