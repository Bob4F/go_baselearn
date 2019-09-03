package main

// 引用自定义包
import "go_baselearn/learn/unit"

/*
 * 程序启动入口类
 */
func main() {
	// 调用其他包下方法
	unit.LearnMap()
	unit.LearnDefer()
}
