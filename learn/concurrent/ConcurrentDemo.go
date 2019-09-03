package concurrent

import (
	f "fmt"
	"time"
)

//  在入口类通过使用go和不使用go语句 使用并发
func Say(str string) {
	for i := 0; i < 5; i++ {
		f.Println(i, ":", str)
		time.Sleep(100 * time.Millisecond)
	}
}

// 学习 channel
// 声明创建channel类型 channel := make(chan int)
func Sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 放入通道
}

// 循环channel 并关闭
func Channel2(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x // 入通道
		x, y = y, x+y
	}
	// 收不到ok就close 关闭函数
	close(c)
}
