package removeElement

// 力扣  904
// 找一个最长连续子数组，满足子数组中至多有两种数字。返回子数组的长度。
func totalFruit(fruits []int) int {
	mp := make(map[int]int)
	l, r := 0, 0
	res := 0
	for r < len(fruits) {
		mp[fruits[r]]++
		for len(mp) > 2 {
			mp[fruits[l]]--
			if mp[fruits[l]] == 0 {
				delete(mp, fruits[l])
			}
			l++
		}
		res = max(res, r-l+1)

		r++
	}
	return res
}
