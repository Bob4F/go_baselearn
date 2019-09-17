package main

import (
	"fmt"
	"strings"
)

func main() {
	checkString()
}

// Contains函数 和 ContainsAny函数
// 是否存在某个字符或则子串
func checkString(){
	//Contains： test中是否包含i 有则为true 非则false
	fmt.Println("contains():",strings.Contains("test","i"))
	// ContainsAny：第二个参数 chars中任意字符在第一个参数s中存在就返回true
	fmt.Println(strings.ContainsAny("team", "i"))// false
	fmt.Println(strings.ContainsAny("failure", "u & i")) // true
	fmt.Println(strings.ContainsAny("in failure", "s g")) // true
	fmt.Println(strings.ContainsAny("foo", ""))// false
	fmt.Println(strings.ContainsAny("", ""))//false
}