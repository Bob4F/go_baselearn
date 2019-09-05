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
 go 关键字
```cassandraql
go关键字开启 goroutine 即可  
goroutine 是轻量级线程，goroutine(协程) 的调度是由 Golang 运行时进行管理的。
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
+ Go interface 接口编程
+ Go 指针Learn  
获取变量的指针  
指针变量赋值  
获取指针变量对应的值

## 09-05
+ Go 错误处理
+ Go 数据库操作
```cassandraql
database/sql 首先是 Go 标准库提供的一个包，用于和 SQL/SQL-Like 数据库(关系或类似关系数据库）通讯   

database/sql 提供的是抽象概念，和具体数据库无关  
具体的数据库实现，由驱动来做，这样可以很方便的更换数据库

该包提供了一些类型（概括性的），每个类型可能包括一个或多个概念。
1.DB  
sql.DB 类型代表了一个数据库。这点和很多其他语言不同，它并不代表一个到数据库的具体连接，而是一个能操作的数据库对象，具体的连接在内部通过连接池来管理，对外不暴露。这点是很多人容易误解的：每一次数据库操作，都产生一个 sql.DB 实例，操作完 Close。
2.Results  
定义了三种结果类型：sql.Rows、sql.Row 和 sql.Result，分别用于获取多个多行结果、一行结果和修改数据库影响的行数（或其返回last insert id）。
3.Statements  
sql.Stmt 代表一个语句，如：DDL、DML等。
4.Transactions  
sql.Tx 代表带有特定属性的一个事务。

mysql 的驱动：github.com/go-sql-driver/mysql
```



