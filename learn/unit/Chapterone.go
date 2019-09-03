package unit

// 自命名包ff
import ff "fmt"

const DEV_NAME = "Lucas"

const (
	Unknow = 0
	Female = 1
	Male   = 2
)

func PringDevName() {
	ff.Println(DEV_NAME)
}

// 方法名首字母小写 私有方法
func learnmap() {
	ff.Println("process private method")
}

func LearnMap() {
	// init map
	_dict := make(map[string]int)
	_dict["key"] = 0
	println("_dict key:", _dict["key"])

	// init map
	m := map[string]int{"one": 1}
	// print map key:one
	println("map:", m["one"])
	learnmap()
}

// defer 会延迟函数的执行 直到上层函数返回
// 多个defer修饰的命令  后入先出
func LearnDefer() {
	ff.Println("start LearnDefer")
	for i := 0; i < 10; i++ {
		defer ff.Println(i)
	}
	ff.Println("end")
}

// 结构体
type Struct_Member struct {
	Id int
	// 小写不可访问 私有
	name string
	age  int
}

func InitMember() {
	member := Struct_Member{1, "fyf", 18}
	ff.Println(member)
}
