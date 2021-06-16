package main

import "fmt"

// 双指针算法, 最大面积
func maxArea(l []int) int {

	length := len(l)
	fmt.Println(length)

	if length < 2 {
		return 0
	}

	max := 0

	for i, j := 0, length - 1; i < j; {
		var temMax int
		if l[i] < l[j] {
			temMax = (j - i) * l[i]
			i ++
		} else {
			temMax =  (j - i) * l[j]
			j --
		}
		if temMax > max {
			max = temMax
		}
		fmt.Println(temMax)
	}
	return max
}

// 三数之和
func threeSum(nums []int) [][]int {

	length := len(nums)
	return [][]int{}

}