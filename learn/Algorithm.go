package main

import (
	"fmt"
	"strings"
)

func main() {
	//l1 :=ListNode{2,&ListNode{4,&ListNode{3,nil}}}
	//l2 :=ListNode{5,&ListNode{6,&ListNode{4,nil}}}
	//result :=twonumcount(&l1,&l2)
	//for result!=nil  {
	//	fmt.Print(result.Val,"->")
	//	result=result.Next
	//}
	fmt.Println("Result:",longestPalindromeTrue("aacdefcaa"))
}


func twonumcount(l1 *ListNode,l2 *ListNode)*ListNode{
	if l1 ==nil || l2 ==nil {
		return l2
	}
	var result = &ListNode{0,nil}
	dump := result
	var carry = 0
	for l1 != nil || l2!=nil {
			var sum int
			if l1==nil {
				sum =  l2.Val + carry
			}else if l2==nil{
				sum = l1.Val +  carry
			}else {
				sum = l1.Val + l2.Val + carry
			}

			fmt.Println("init sum:",sum)
			carry = sum / 10
			fmt.Println("carry:",carry)
			sum %= 10
			fmt.Println("sum&:",sum)

			dump.Next = &ListNode{sum,nil}

			dump =dump.Next
			if l1 != nil{
				l1 = l1.Next
			}
			if(l2 != nil){
				l2 = l2.Next
			}
	}
	if carry==1 {
		dump.Next = &ListNode{1,nil}
	}
	return result.Next;
}

type ListNode struct {
	Val int
	Next *ListNode
}

// t通过了 43个例子 卡在了aacdefcaa 返回aac 预期是要aa
func longestPalindrome(s string)string{
	if  &s ==nil || len(s) <= 1 {
		return s
	}
	temp:=""
	// 反转
	for i :=len(s)-1;i > -1;i-- {
		temp = temp + string(s[i])
	}

	fmt.Println("temp:",temp,",len:",len(temp))
	fmt.Println("s:",s)
	result :=false
	resultstr := ""

	// 返回**最长**回文字符串
	for i:=0;i< len(s); i++ {
		for next:= len(s);next > -1;next--  {
			if(next<i){
				continue
			}
			fmt.Println("i:",i,"next:",next)
			cache_temp := temp[i:next]
			if len(cache_temp)<= 1 {
				continue
			}
			if strings.Contains(s,string(cache_temp)) {
				fmt.Println(i,":",cache_temp)
				if len(resultstr) > 0 && len(cache_temp) > len(resultstr) {
					resultstr = cache_temp
				}else if len(resultstr) == 0 {
					resultstr = cache_temp
					result = true
				}
			}
		}

	}
	if result {
		return resultstr
	}
	// 如果没有长回文子串，就返回字符第一个字符 不需要返回""
	return s[0:1]
}

// 动态规划
func longestPalindromeTrue(s string)string{
	strlen := len(s)
	if strlen < 2 {
		return s
	}
	temp := ""

	currLen:=0
	maxLen:=0
	for i:=0;i<strlen ;i++  {
		for j:=0;j<=i ;j++  {
			if((s[i] == s[j])&&((i-j < 2)||(i>0))){
				currLen = i - j + 1;
				if(currLen > maxLen){
					maxLen = currLen;
					temp = s[j:currLen];
				}
			}
		}
	}
	return temp
}