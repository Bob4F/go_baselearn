package leetcode

import (
	"fmt"
)

/**
	最长无重复子回复
 */
func main() {
	fmt.Println(lengthOfLongSubstring("a"))		   // a 1
	fmt.Println(lengthOfLongSubstring("aab")) // AAB 2
	fmt.Println(lengthOfLongSubstring("abcabcbb")) // abc 3
	fmt.Println(lengthOfLongSubstring("pwwkew")) // abc 3
}

// 返回 最长无重复子串 长度
func lengthOfLongSubstring(s string)int{
	lastOccurred := make(map[int32]int) //记录字符最后出现的位置
	start,maxLength := 0,0
	for i, ch := range []int32(s) {
		// 如果ch在map中 并且 该ch的索引大于等于start
		// 假设 s = "abcabcbb" i = 5  ch = b
		// 在map中 存了上一个b的位置 abc  那么lastI = 2 那么此时 start 是 3
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		// 如果没有重复的 那么 start 一直没变，i在增长，我们可以通过下面的等式得到这个无重复子串的长度
		// +1是因为索引是从0开始
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		// ch 其实是底层的ascii值
		lastOccurred[ch] = i
	}
	return maxLength
}