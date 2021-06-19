package main

import "fmt"

// 算法练习
func main() {
	fmt.Println("算法练习进行中。。。。。")

	l := []int{1,8,6,2,5,4,8,3,7}

	//BubbleSort(l)

	QuickSort(l, 0, len(l) - 1)
	fmt.Println(l)
	fmt.Println(FindEr(l, 0, len(l) - 1, 8))


	//fmt.Println(maxArea(l))
}
