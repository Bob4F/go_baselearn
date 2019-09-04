# go_baselearn
Go 的基础学习   
预期时间一周  
开始时间:2019-09-03  
结束时间:2019-09-08

## 2019-09-03
+ 基础语法 for / if   
+ 基础数据 ：
string ，int ，float，map，bool ，slice   
+  go / 并发：  
 go关键字
```cassandraql
go关键字开启 goroutine 即可  
goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。
```
channel / 通道 实现线程间通信数据结构
```cassandraql
ch:=make(chan int)
ch <- v    // 把 v 发送到通道 ch
v := <-ch  // 从 ch 接收数据  并把值赋给 v
// 无缓冲区，发送端发送数据，必须有接收端接收数据
//            
```

+ struct 结构体：类似实体类  

+ go语言的函数

## 09-04
+ go 语言范围：Range 关键字：  
循环数组  / 循环map集合
+ Go语言递归  
实现阶乘 / 实现斐波那契数列

