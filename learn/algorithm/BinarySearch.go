/**
	二分查找实现
 */
package main

import "fmt"

func main() {
	a := []int{1,2,4,5,6,7,8,9,10}
	flag:=binary_search(a,5)
	fmt.Println(flag)
}

/*
	接受一个有序数组 和 元素
	如果数组中有该元素，则返回元素所在位置(下标)
*/
func binary_search(list []int,num int)int{
	low := 0
	high := len(list)
	for low <= high   {
		mid := (low+high)/2
		guess := list[mid]
		if guess == num {
			return mid;
		}
		if guess > num {
			high = mid-1
		}else {
			low =mid +1
		}
	}
	return 0
}

