package main

import "fmt"

// 归并排序

func mergeSort(nums []int) []int {
	size := len(nums)
	if size < 2 {
		return nums
	}
	i := size / 2
	left := mergeSort(nums[0:i])
	right := mergeSort(nums[i:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	l, r := len(left), len(right)
	result := make([]int, 0, l+r)
	i, j := 0, 0
	for i < l && j < r {
		if left[i] > right[j] {
			result = append(result, right[j])
			j++
		} else {
			result = append(result, left[i])
			i++
		}
	}
	result = append(result, right[j:]...)
	result = append(result, left[i:]...)
	return result
}

func main() {
	nums := []int{10, 2, 8, 4, 5, 0, 0, 1, 7, 3, 11, 11, 18, 0}
	fmt.Println(mergeSort(nums))
}
