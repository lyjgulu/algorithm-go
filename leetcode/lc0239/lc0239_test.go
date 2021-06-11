package leetcode

import (
	"fmt"
	"testing"
)

type question239 struct {
	para239
	ans239
}

// para 是参数
// one 代表第一个参数
type para239 struct {
	nums   []int
	target int
}

// ans 是答案
// one 代表第一个答案
type ans239 struct {
	one []int
}

func TestProblem239(t *testing.T) {

	qs := []question239{
		{
			para239{[]int{1, 2, 3, 4, 5, 6, 7}, 3},
			ans239{[]int{5, 6, 7, 1, 2, 3, 4}},
		},

		{
			para239{[]int{-1, -100, 3, 99}, 2},
			ans239{[]int{3, 99, -1, -100}},
		},

		{
			para239{[]int{1, 2, 3, 4, 5, 6, 7}, 8},
			ans239{[]int{7, 1, 2, 3, 4, 5, 6}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 239------------------------\n")

	for _, q := range qs {
		a, p := q.ans239, q.para239
		fmt.Printf("【input】:%v       【answer】:%v      【output】:%v\n", p, a, rotate(p.nums, p.target))
	}
	fmt.Printf("\n\n\n")
}
