package main

import (
	"fmt"
	"sort"
)

/*
	给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

	请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。(性能很重要)

	你可以假设 nums1 和 nums2 不会同时为空。
 */
func main() {
	var nums1 = []int{1,2}
	var nums2 = []int{3,4}
	var result = findMedianSortedArrays(nums1,nums2)
	fmt.Println(result)
}


/*
  没有实现 O(log(m+n))
 */
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var to []int
	var f float64
	to=append(to,nums1...)
	to=append(to,nums2...)
	sort.Ints(to) //直接调用sort包排序
	if len(to)%2==0 { //判断奇偶数区分中位数
		f= (float64(to[len(to)/2-1]) + float64(to[len(to)/2]))/2
	}else {
		f= float64(to[(len(to)-1)/2])
	}
	return f
}