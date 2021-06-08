package lc0189

import (
	"fmt"
	"testing"
)

type question189 struct {
	para189
	ans189
}

// para 是参数
// one 代表第一个参数
type para189 struct {
	nums   []int
	target int
}

// ans 是答案
// one 代表第一个答案
type ans189 struct {
	one []int
}

func TestProblem189(t *testing.T) {

	qs := []question189{
		{
			para189{[]int{1, 2, 3, 4, 5, 6, 7}, 3},
			ans189{[]int{5, 6, 7, 1, 2, 3, 4}},
		},

		{
			para189{[]int{-1, -100, 3, 99}, 2},
			ans189{[]int{3, 99, -1, -100}},
		},

		{
			para189{[]int{1, 2, 3, 4, 5, 6, 7}, 8},
			ans189{[]int{7, 1, 2, 3, 4, 5, 6}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 189------------------------\n")

	for _, q := range qs {
		a, p := q.ans189, q.para189
		fmt.Printf("【input】:%v       【answer】:%v      【output】:%v\n", p, a, rotate(p.nums, p.target))
	}
	fmt.Printf("\n\n\n")
}
