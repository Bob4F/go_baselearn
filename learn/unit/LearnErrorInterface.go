package main

import (
	"errors"
	"fmt"
)

/*
  go语言通过内置的错误接口提供了非常简单的错误处理机制
  error 类型是一个接口类型
 */
type error interface{
	Error() string
}

func ThrowException(){
	number,error:=throw()
	fmt.Println("error info:",error)
	number++
}

func throw() (int,error){
	return 0,errors.New("math: method root of negative number")

}




