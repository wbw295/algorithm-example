package main

import "fmt"

func selectSort(nums []int) {

	for i := 0; i < len(nums)-1; i++ {
		minIndex := i
		minValue := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < minValue {
				minIndex = j
				minValue = nums[j]
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}

}

func main() {
	nums := []int{10, 2, 8, 4, 5, 0, 0, 1, 7, 3, 11, 11, 18, 0}
	selectSort(nums)
	fmt.Println(nums)
}
