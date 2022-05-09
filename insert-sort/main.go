package main

import "fmt"

func insertSort(nums []int) {

	for i := 1; i < len(nums); i++ {
		for j := i; j > 0 && nums[j] < nums[j-1]; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}

}

func main() {

	nums := []int{10, 2, 8, 4, 5, 0, 0, 1, 7, 3, 11, 11, 18, 0}
	insertSort(nums)
	fmt.Println(nums)

}
