package main

import "fmt"

func quickSort(nums []int, left, right int) {

	l, r := left, right
	pivot := nums[(l+r)/2]

	for l < r {
		for nums[r] > pivot {
			r--
		}
		for nums[l] < pivot {
			l++
		}
		if l >= r {
			// 说明本轮分割完成
			break
		}

		// 交换
		nums[l], nums[r] = nums[r], nums[l]

		// 优化
		if nums[l] == pivot {
			r--
		}
		if nums[r] == pivot {
			l++
		}
	}
	if l == r {
		l++
		r--
	}

	if left < r {
		quickSort(nums, left, r)
	}

	if l < right {
		quickSort(nums, l, right)
	}

}

func main() {

	nums := []int{10, 2, 8, 4, 5, 0, 0, 1, 7, 3, 11, 11, 18, 0}

	quickSort(nums, 0, len(nums)-1)

	fmt.Println(nums)

}
