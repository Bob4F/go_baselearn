/*
	D&C / divide and conquer(分而治之思想)：递归方式解决问题的方法
	要点:
	1. 找到基线条件，尽可能简单(退出循环)
	2. 讲问题分解(缩小规模)，直到符合基线条件

	应用algorithm:快速排序 二分查找
 */
package main

import (
	"fmt"
)

func main() {
	s :=[] int {1,2,3,7,4}
	fmt.Println(MaxNum(s,0))

	arr :=[] int {1,2,3,7,4,6}
	quicksortresult:=quickSort(arr)
	fmt.Println(quicksortresult)

}

/*
 递归解决数组之和
 */
func Sum(arr []int) int {
	if len(arr)!=0 {
		// 基线条件 缩小规模
		return arr[0]+Sum(arr[1:len(arr)])
	}else{
		// 基线条件
		return 0
	}
}

/**
   找出数组中最大的数字
 */
func MaxNum(arr []int,num int)int{
	if len(arr)==0 {
		// 基线条件
		return num
	}else{
		if arr[0] > num {
			num = arr[0]
		}
		// 缩小规模
		return MaxNum(arr[1:len(arr)],num)
	}
}


/*
  快速排序
 */
func quickSort(arr []int)[]int{
	if len(arr)<2 {
		return arr
	}else {
		zero:=arr[0]
		left:=make([]int, 0)
		right:=make([]int, 0)
		for _,num:= range arr[1:len(arr)]  {
			if zero > num {
				left = append(left,num)
			}
		}
		for _,num:= range arr[1:len(arr)]  {
			if zero < num {
				right = append(right,num)
			}
		}
		left=quickSort(left)
		right=quickSort(right)
		left = append(left, zero)
		for _,num:= range right  {
			left = append(left,num )
		}
		return left
	}
}