package main
/*
  Go语言学习 接口 / interface
 */

import "fmt"

// 定义接口
type MemberService interface {
	// 接口方法
	getUserName()
}
// 结构体
type member struct {
	id int
	name string
}
// 实现的接口
func (mem member) getUserName(){
	fmt.Println("name:Lucas Fong")
}


func PublicGetUserName(){
	fmt.Println("Start Interface Demo")
	var service MemberService
	service = new(member)
	service.getUserName()

}

