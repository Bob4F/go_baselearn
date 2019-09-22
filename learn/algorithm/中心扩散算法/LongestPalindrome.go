package main

import (
	"fmt"
)

/*
  动态规划
  最长回文子串
 */
func main() {
	//fmt.Println("result:",longestPalindrome("acccccid"))
	fmt.Println(5/2)
	fmt.Println("result:",longestPalindrome("abcba"))
}

func longestPalindrome(s string)string{
	length := len(s)
	if length < 2 {
		return s
	}
	// 最长回文的首字符索引，和最长回文的长度
	begin, maxLen := 0, 1
	// 在 for 循环中
	// start 代表回文的**首**字符索引号
	// end 代表回文的**尾**字符索引号
	// i 代表回文`正中间段`首字符的索引号
	// 在每一次for循环中
	// 先从i开始，利用`正中间段`所有字符相同的特性，让start，end分别指向`正中间段`的首尾
	// 再从`正中间段`向两边扩张，让start，end分别指向此`正中间段`为中心的最长回文的首尾
	for i := 0; i < length; { // 以s[i]为`正中间段`首字符开始寻找最长回文。
		fmt.Println("第",i,"次:")
		if length -i <= maxLen/2 {
			break
			// 因为i是回文`正中间段`首字符的索引号
			// 假设此时能找到的最长回文的长度为l, 则，l <= (len(s)-i)*2 - 1( l不可能大于s的长度 )
			// 如果，len(s)-i <= maxLen/2 成立
			// 则，l <= maxLen - 1
			// 则，l < maxLen
			// 所以，无需再找下去了
		}

		start,end :=i,i
		for end < length-1 && s[end+1] == s[end] {
			// 循环结束后，s[start:end+1]是一串相同的字符串
			end++
		}

		// 下一个回文的`正中间段`的首字符只会是s[end+1]
		// 为下一次循环做准备
		i = end + 1
		// 前两个条件是防止数组指针溢出
		for end < length -1 && start > 0 && s[end+1] == s[start-1] {
			end++
			start--
			// 循环结束后，s[start:end+1]是这次能找到的最长回文。
		}

		newLen := end + 1 - start
		// 创新记录的话，就更新记录
		if newLen > maxLen {
			begin = start
			maxLen = newLen
		}
	}
	return s[begin : begin+maxLen]
}