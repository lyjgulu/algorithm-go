package leetcode

import (
	"fmt"
	"testing"
)

type question322 struct {
	para322
	ans322
}

// para 是参数
// one 代表第一个参数
type para322 struct {
	nums   []int
	target int
}

// ans 是答案
// one 代表第一个答案
type ans322 struct {
	one int
}

func TestProblem322(t *testing.T) {

	qs := []question322{
		{
			para322{[]int{1, 2, 5}, 100},
			ans322{20},
		},

		{
			para322{[]int{2}, 3},
			ans322{-1},
		},

		{
			para322{[]int{1}, 0},
			ans322{0},
		},

		{
			para322{[]int{190, 80, 457, 111, 240}, 6159},
			ans322{17},
		},

		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 322------------------------\n")

	for _, q := range qs {
		a, p := q.ans322, q.para322
		fmt.Printf("【input】:%v       【answer】:%v      【output】:%v\n", p, a, coinChange(p.nums, p.target))
	}
	fmt.Printf("\n\n\n")
}
