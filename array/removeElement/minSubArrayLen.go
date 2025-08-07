package removeElement

//力扣 209
//给定一个含有 n 个正整数的数组和一个正整数 target 。
//
//找出该数组中满足其总和大于等于 target 的长度最小的 子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，
//并返回其长度。如果不存在符合条件的子数组，返回 0 。
//
//
//
//示例 1：
//
//输入：target = 7, nums = [2,3,1,2,4,3]
//输出：2
//解释：子数组 [4,3] 是该条件下的长度最小的子数组。

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	sum := 0
	res := n + 1
	l, r := 0, 0
	for ; r < n; r++ {
		sum += nums[r]
		for sum >= target {
			res = min(res, r-l+1)
			sum -= nums[l]
			l++
		}
	}
	if res > n {
		return 0
	}
	return res
}
