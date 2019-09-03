package main

import "go_baselearn/learn/concurrent"


func main() {
	// 没有固定先后顺序 下面两个方法 并行执行

	go concurrent.Say("HELLO")
	concurrent.Say("world")


}
