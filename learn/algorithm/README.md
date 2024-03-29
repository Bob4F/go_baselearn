# 算法简介

## 数据结构

### 图
图由节点和边组成，非线性结构

常见用例: 广度优先搜索(BFS),


### 队列(queue)
First In First Out 先进先出，与栈(后进先出)相反


## 对数 / logrithm
*幂运算的逆运算*

在算法中 *log* 时间经常出现，需要了解


## 大 O 表示法
指出算法的速度有多快，表示时间复杂度  
他的单位并不是时间单位，而是操作次数

> O(log n) 对数时间 （二分查找）
> O(n) 线性时间 （简单查找）
> O(n^2) 通常被称为暴力破解法





# 动态规划
忘记过去的人被谴责重复的事

![](https://user-gold-cdn.xitu.io/2019/9/16/16d3923b83a0bb27?w=359&h=140&f=png&s=17620)

## 
边界条件：即最简单的，可以直接得出的局部最优解。  
全局最优解中一定包含某个局部最优解，但不一定包含前一个局部最优解，因此需要记录之前的所有最优解； # 记住已经解决过的局部的解 
动态规划的关键是**状态转移方程**，即如何由以求出的局部最优解来推导全局最优解；

通过加法例子说明:
```cassandraql
1+1+1+= ？ 答案是 3  
那么在上面等式上 +1 ，问此时等值的值为多少  
答案是4，你不可能在重新通过 1+1+1+1得到4 而是在3的基础上+1    
这就是动态规划，记住求过的解来节省时间  
```
## 算法案例


## 参考文献
[算法-动态规划 Dynamic Programming--从菜鸟到老鸟](https://blog.csdn.net/u013309870/article/details/75193592): 强烈推荐

# 中心扩散算法
顾名思义就是以某一个位置为中心，向周围扩散，直到满足条件或到达边界

## 算法案例
[最长回文子串](https://github.com/lucasIsfyf/go_baselearn/blob/master/learn/algorithm/leetcode/LongestPalindrome.go)
用时 0ms/2.3mb



# History
9.30 RomaToint [罗马数转整数](https://github.com/lucasIsfyf/go_baselearn/blob/master/learn/algorithm/leetcode/RomaToint.go)

