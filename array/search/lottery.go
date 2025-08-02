package search

import (
	"fmt"
	"math/rand"
	"time"
)

//某抽奖系统有 N 种奖品，每种奖品有一个权重，权重越高被抽中的概率越大。请实现一个抽奖算法，要求每次调用 Draw () 方法时，
//能按照权重概率随机返回一个奖品编号。
//函数签名：
//type Lottery struct {}
//func NewLottery (weights [] int) *Lottery
//func (l *Lottery) Draw () int
//
//输入输出说明：
//
//输入：weights = [1,3,6]，表示奖品 0 权重 1，奖品 1 权重 3，奖品 2 权重 6
//输出：多次调用 Draw ()，奖品 2 被抽中的概率约为 60%，奖品 1 约为 30%，奖品 0 约为 10%

// 算法思路
//1. 将权重数组转化为连续的数组区间
//2. 生成一个随机数，判断随机数落在哪一个区间，区间的编号就是中奖的奖品编号

type Lottery struct {
	arry  []int
	total int
}

func NewLottery(weights []int) *Lottery {
	arry := make([]int, len(weights))
	sum := 0
	for i, v := range weights {
		sum += v
		arry[i] = sum
	}
	return &Lottery{arry: arry, total: sum}
}

func (l *Lottery) Draw() int {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(l.total) + 1
	left, right := 0, len(l.arry)-1
	for left <= right {
		mid := left + (right-left)/2
		if l.arry[mid] == target {
			return mid
		} else if l.arry[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	//此时的 left 就指向了目标值应该所在的区间
	return left
}

func TestLottry() {
	l := NewLottery([]int{1, 3, 6})
	count := make([]int, 3)
	for i := 0; i < 100; i++ {
		res := l.Draw()
		count[res]++
	}
	// 打印概率分布
	for i, cnt := range count {
		probability := float64(cnt) / 100 * 100
		fmt.Printf("奖品%d: %.2f%%\n", i, probability)
	}
}
