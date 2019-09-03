package unit

/*
 Go 切片（Slice）
 Go 数组长度不可变，不够灵活
 切片(动态数组) 长度不固定 可以追加元素
*/
import "fmt"

// var identifier []type  = make([]type,len)
//  type 数据类型;  len 初始化长度
var sliceone []int = make([]int, 10)

func init() {
	// 直接初始化 []表示是切片类型
	// cap=len = 3
	slice3 := []int{1, 2, 3}
	fmt.Println(slice3)
}

// 返回slice3的下表为2的值
func ReturnSlice3Index2() int {

	return 0
}
