package main

import (
	"fmt"
	"strings"
)

func main() {
	//checkString()
	useIndex()

}

// Contains函数 和 ContainsAny函数
// 是否存在某个字符或则子串
func checkString(){
	//Contains： test中是否包含i 有则为true 非则false
	fmt.Println("contains():",strings.Contains("test","i")) // false
	// ContainsAny：第二个参数 chars中任意字符在第一个参数s中存在就返回true
	fmt.Println(strings.ContainsAny("team", "i"))// false
	fmt.Println(strings.ContainsAny("failure", "u & i")) // false
	fmt.Println(strings.ContainsAny("in failure", "sg")) // false
	fmt.Println(strings.ContainsAny("foo", ""))// false
	fmt.Println(strings.ContainsAny("", ""))//false
}

/* 根据索引获取对应值 */
func useIndex(){
	s := "Hello, World!"
	s1 := s[:5]  // Hello
	s2 := s[7:]  // World!
	s3 := s[1:5] // ello
	s4 := s[0:] //  Hello, World!
	s5 := s[3:] //  lo, World!
	s6 := s[0:1]
	fmt.Println(s,"\n", s1,"\n", s2, s3,s4,"\n",s5,"\n",s6)
}