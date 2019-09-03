/*
 * 代码第一行写明package；使用一个文件单独运行的程序必须放在package main下面，否则在go run运行时会报错：”go run: cannot run non-main package“
 */
package main

import "fmt"
// 引入自建包
import uu "main/unit"



func main() {
	fmt.Println("Hello Go!")
	vari1()
	uu.LearnMap();
	var num int= getNum(89);
	fmt.Print(num)
}


// 语言变量 var v_name v_type
func vari1(){
	// string
	var name string  = "Lucas"
	// print
	fmt.Println(name)
	// int
	var age int = 18;
	fmt.Println(name,"的年龄", age)




	// bool 默认false
	var c bool
	fmt.Println(c)
}
/*
 * go 函数
 *
 */
func getNum(arr int) int{
	var varnum int
	for a := 0;arr > 0 ;arr= arr/10 {
		varnum = (varnum * 10 )+ (arr%10)
		a = a+1
	}
	return varnum
}






