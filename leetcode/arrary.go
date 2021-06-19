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

// 两数之和
func twoSum(nums []int, target int) []int {
	targetList := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		subTarget := target - nums[i]
		for j := 0; j < len(nums); j++ {
			if nums[j] == subTarget && nums[i] != nums[j] {
				targetList[0] = i
				targetList[1] = j
				return targetList
			}
		}
	}
	return []int{}
}

// 两数之和hash
func TwoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		subTarget := target - nums[i]
		if value, ok := hashMap[subTarget]; ok {
			return []int{i, value}
		}
		hashMap[nums[i]] = i
	}
	return []int{}
}

// 两个顺序数组的中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length1 := len(nums1)
	length2 := len(nums2)

	index1 := 0
	index2 := 0
	var data float64

	midIndex := (length1 + length2) / 2

	for i := 0; i < midIndex + 1; i++ {
		if nums2[index2] < nums1[index1] {
			index1 ++
			data = float64(nums1[index1])
		} else {
			index2 ++
			data = float64(nums2[index2])
		}
	}

	return data
}
// 三数之和
//func threeSum(nums []int) [][]int {
//
//	length := len(nums)
//	return [][]int{}
//
//}