package main
/*
 Go 实现递归
  在使用递归时，开发者需要提供退出时间 否则无限循环
  计算阶乘 生成斐波那契数列
 */
import (
	"fmt"
	"time"
)

// 对外提供调用 递归函数
func Recursion(){
	recursion()
}

func recursion()  {
	recursion()
	fmt.Println("start method recursion")
	time.Sleep(100)
}


func Factorial(){
	fmt.Println("Start Method Factorial!!!")
	var i int = 6
	fmt.Printf("%d 的阶乘是 %d\n",i,factorial(uint64(i)))
}

/*
  阶乘
  5! = 1 * 2 * 3 * 4 * 5
*/
func factorial(n uint64) (result uint64) {
	// 结束条件
	if ( n > 0 ) {
		result = n * factorial(n-1)
		return result
	}
	// 退出条件
	return 1
}


// 对外提供斐波那契数列
func Fibonacci(){
	fmt.Println("Start Method Fibonacci")
    var i int
    for i=0;i<10;i++{
    	fmt.Printf("%d\t",fibonacci(i))
	}
}


func fibonacci(n int) int{
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}