package removeElement

import (
	"fmt"
)

// 你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素。元素的顺序可能发生改变。
// 然后返回 nums 中与 val 不同的元素的数量。
func removeElement(nums []int, val int) int {
	slow, fast := 0, 0
	for fast <= len(nums)-1 {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

func TestRemoveElement() {
	nums := []int{2}
	result := removeElement(nums, 3)
	fmt.Println(result)
	fmt.Println(nums)
}
