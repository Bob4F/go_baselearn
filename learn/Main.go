package main

// 引用自定义包
import (
	"fmt"
	method "go_baselearn/learn/methods"
	"go_baselearn/learn/unit"
)

/*
 * 程序启动入口类
 */
func main() {
	// 调用其他包下方法
	// 关于map
	unit.LearnMap()
	// 关于defer
	unit.LearnDefer()
	unit.PringDevName()
	// 直接调用常量
	var sex int = unit.Female
	println("性别：", sex)
	unit.InitMember()
	var member unit.Struct_Member
	member.Id = 5
	fmt.Println("member id:", member.Id)
	// 返回多参数
	H, N := method.MoreReturn()
	fmt.Println(H, N)
	jsonmap := make(map[string]int)
	method.Method3("key", 1, 2, jsonmap)
	fmt.Println("JSONMAP", jsonmap["key"])
	// 数组
	unit.Arr()
	unit.LearnRangeArr()
	unit.LearnRangeMap()
	// unit.Recursion()
	// 递归：阶乘
	unit.Factorial()
	// 递归：斐波那契数列
	unit.Fibonacci()
	//
	unit.PublicGetUserName()
	unit.Pointer()
    unit.ThrowException()
}
