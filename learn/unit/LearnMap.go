package main

import "fmt"

/*
 Go 数据结构 map 无序集合
*/


func main(){
	test1()
	query("a")
	query("b")
	m["k1"] = 2
	m["k2"] = 2
	m["k3"] = 2
	m["k4"] = 2
	fmt.Println("m的长度",getLength())
	foreachMap()

}

var m map[string]int = make(map[string]int)

// 赋值 / insert
func test1()  {
	m["a"] = 1
	fmt.Printf("全局变量 map m: %v\n" ,m)
}

// 查找
func query(s string){
	val,ok := m[s]
	if ok {
		fmt.Println("Key:",s,"\tresult:",val)
	} else {
		fmt.Println("map中没有",s,"Key")
	}
}

//删除 map的元素 先判断是否存在 通过delete函数
func remove(hashmap map[string]int, key string) {
	if _,ok := hashmap[key];ok {
		delete(hashmap, key)
	}

}

// 获取长度
func getLength()int{
	return len(m)
}

// 循环遍历map
func foreachMap(){
	fmt.Println("start foreach Map")
	for key,val := range m {
		fmt.Println("key:",key,"val:",val)
	}
}


