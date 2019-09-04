package unit
/*
  Go  Range 范围
  range关键字用于for循环中 迭代数组(array),slice,channel,map
  slice和array中返回元素 即 索引和对应值
  集合中返回 key-value 返回key
 */
import "fmt"


// 计算数组总值
func LearnRangeArr(){
	println("Start Method LearnRange1")
	nums := []int{1,2,3}
	sum:= 0
	// range将传入 index 和value  '_' 代表省略index传入
	for _, num:=range nums {
		sum+=num
	}
	println("sum:",sum)

	for i, num:=range nums {
		if(num==3){
			fmt.Print("index:",i)
		}
	}
}


func LearnRangeMap(){
	println("Start Method LearnRangeMap")
	hashmap := map[string]string{"a":"apple","b":"banana"}
	for k,v :=range hashmap {
		fmt.Printf("%s -> %s\n",k,v)
	}


}