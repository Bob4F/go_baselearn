package unit

/*
   Go 数组 长度不可变
*/
import _ "fmt"

// var variable_name [size] variable_type
// var table_size [10] int
// 初始化 元素不能多余[]的数字
var balance = [5]float32{1000.0, 5.4, 2, 4}

func Arr() {
	// 赋值
	balance[4] = 50.0
	// 访问并打印
	println("balance4", balance[4])
	var item2 = balance[1]
	println("item2", item2)

}
