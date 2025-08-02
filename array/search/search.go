package search

import "fmt"

// 二分查找
// 左闭右闭 [l, r] l <= r
func Search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}

func SearchTest() {
	nums := []int{1, 3, 5, 6}
	res := Search(nums, 7)
	fmt.Println(res)
}
