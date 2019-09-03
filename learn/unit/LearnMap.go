package unit

/*
 Go 数据结构 map 无序集合
 hash表实现
*/
// 创建
var MapDemo = make(map[string]int)

// 赋值
func Init() {
	MapDemo["one"] = 1
	MapDemo["two"] = 2
}

//删除 map的元素 通过delete函数
func remove(hashmap map[string]int, key string) {
	// 删除元素
	delete(hashmap, key)
}

// map 实现 hashmap
