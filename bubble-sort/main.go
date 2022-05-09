package main

import "fmt"

func bubbleSort(nums []int) {

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

}

func main() {
	nums := []int{10, 2, 8, 4, 5, 0, 0, 1, 7, 3, 11, 11, 18, 0}
	bubbleSort(nums)
	fmt.Println(nums)
}
