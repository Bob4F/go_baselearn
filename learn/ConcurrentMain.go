package main

import (
	"fmt"
	"go_baselearn/learn/concurrent"
)

func main() {
	// 没有固定先后顺序 下面两个方法 并行执行

	//go concurrent.Say("HELLO")
	//concurrent.Say("world")

	learnchannel()

	fmt.Println("start for channel")
	channel1 := make(chan int, 10) // 缓冲10
	concurrent.Channel2(cap(channel1), channel1)
	// 没加go 会报错 死锁
	// go concurrent.Channel2(11,channel1)
	for i := range channel1 {
		fmt.Println(i)
	}
}

func learnchannel() {
	// channel
	s := []int{7, 2, 4, 2, 6}
	// 创建channel
	c := make(chan int)
	go concurrent.Sum(s, c)
	go concurrent.Sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从 c 通道中取出
	fmt.Println(x, y, x+y)
	// channel buffer
	ch := make(chan int, 4)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	// 在多放会运行中报错
	fmt.Println("start channel buffer")
	// 先入先出
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
