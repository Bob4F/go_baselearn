package main

import "fmt"

/*
  动态规划
  最长回文子串
 */
func main() {
	fmt.Println("result:",longestPalindrome("b"))
}

func longestPalindrome(s string)string{
	if len(s)  <= 1  {
		return s
	}else if len(s)==2  && s[0]!=s[1]{
		return string(s[0])
	}
	
	return ""
}