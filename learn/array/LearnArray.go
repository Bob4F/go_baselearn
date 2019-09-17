package main

import "fmt"

func main() {
	initNums()
}

// 初始化一维数组
func initNums(){
	// 不设置默认为0
	a := [3]int{1,2}
	fmt.Println(a)
}
