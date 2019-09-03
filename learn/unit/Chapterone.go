package unit

// 自命名包ff
import ff "fmt"

// 方法名首字母小写 私有方法
func learnmap(){
	ff.Println("process private method")
}

func LearnMap(){
	// init map
	_dict := make(map[string]int)
	_dict["key"] = 0
	println("_dict key:",_dict["key"])

	// init map
	m := map[string]int{"one":1}
	// print map key:one
	println("map:",m["one"])
	learnmap()
}