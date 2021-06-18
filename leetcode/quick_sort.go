package main

// 快速排序
func QuickSort(l []int, start, end int) {
	if start >= end {
		return
	}
	base := l[start]
	left := start
	right := end

	for true {
		if left >= right {
			break
		}
		if left < right && l[right] >= base {
			right --
		}
		l[left] = l[right]

		if left < right && l[left] < base {
			left ++
		}
		l[right] = l[left]
	}

	l[left] = base
	QuickSort(l, start, right - 1)
	QuickSort(l, right + 1, end)

}

// 冒泡排序

func BubbleSort(l []int) {

	length := len(l)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if l[j] < l[i] {
				l[i], l[j] = l[j], l[i]
			}
		}
	}
}

// 二分法查找
func FindEr(l []int, start, end, data int) int {
	if start > end {
		return -1
	}
	var mid int
	mid = (start + end) / 2
	if l[mid] == data {
		return mid
	} else if l[mid] > data {
		return FindEr(l, start, mid - 1, data)
	} else {
		return FindEr(l, mid + 1, end, data)
	}
}
