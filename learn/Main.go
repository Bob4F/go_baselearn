package main

// 引用自定义包
import (
	"fmt"
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

}
