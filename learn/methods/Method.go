package methods

// 返回多个string 函数
func MoreReturn() (string, string) {
	return "你好", "GoLang"
}

// 无入参 返回单个string
func Method1() string {
	return ""
}

// 入参int类型 返回int
func Method2(x, y int) int {
	return x + y
}

// 入参 无返回值
// 引用传递
func Method3(str string, x, y int, result map[string]int) {
	result[str] = x
}
