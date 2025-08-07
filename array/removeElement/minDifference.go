package removeElement

import "math"

// 题目描述
// 在一个城市区域内，被划分成了n * m个连续的区块，每个区块都拥有不同的权值，代表着其土地价值。目前，有两家开发公司，A 公司和 B 公司，希望购买这个城市区域的土地。
// 现在，需要将这个城市区域的所有区块分配给 A 公司和 B 公司。
// 然而，由于城市规划的限制，只允许将区域按横向或纵向划分成两个子区域，而且每个子区域都必须包含一个或多个区块。
// 为了确保公平竞争，你需要找到一种分配方式，使得 A 公司和 B 公司各自的子区域内的土地总价值之差最小。
//
// 注意：区块不可再分。
//
// 输入描述
// 第一行输入两个正整数，代表 n 和 m。
//
// 接下来的 n 行，每行输出 m 个正整数。
//
// 输出描述
// 请输出一个整数，代表两个子区域内土地总价值之间的最小差距。
// 输入示例
// 3 3
// 1 2 3
// 2 1 3
// 1 2 3
// 输出示例
// 0
func minDifference(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])

	total := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			total += grid[i][j]
		}
	}

	//最后一行不用算
	minDiff := total
	// 横向划分
	rowTotal := 0
	for i := 0; i < n-1; i++ {
		tmpTotal := 0
		for j := 0; j < m; j++ {
			tmpTotal += grid[i][j]
		}
		rowTotal += tmpTotal
		downTotal := total - rowTotal
		diff := math.Abs(float64(rowTotal - downTotal))
		if int(diff) < minDiff {
			minDiff = int(diff)
		}
	}
	// 纵向划分，最后一列不用算
	leftTotal := 0
	for j := 0; j < m-1; j++ {
		tmpTotal := 0
		for i := 0; i < n; i++ {
			tmpTotal += grid[i][j]
		}
		leftTotal += tmpTotal
		rightTotal := total - leftTotal
		diff := math.Abs(float64(leftTotal - rightTotal))
		if int(diff) < minDiff {
			minDiff = int(diff)
		}
	}
	return minDiff
}

func TestMinDifference() {
	grid := [][]int{{1, 2, 3}, {2, 1, 3}, {1, 2, 3}}
	res := minDifference(grid)
	println(res)
}
