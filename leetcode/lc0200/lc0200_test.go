package leetcode

import (
	"fmt"
	"testing"
)

type question200 struct {
	para200
	ans200
}

// para 是参数
// one 代表第一个参数
type para200 struct {
	nums   []int
	target int
}

// ans 是答案
// one 代表第一个答案
type ans200 struct {
	one []int
}

func TestProblem200(t *testing.T) {

	qs := []question200{
		{
			para200{[]int{1, 2, 3, 4, 5, 6, 7}, 3},
			ans200{[]int{5, 6, 7, 1, 2, 3, 4}},
		},

		{
			para200{[]int{-1, -100, 3, 99}, 2},
			ans200{[]int{3, 99, -1, -100}},
		},

		{
			para200{[]int{1, 2, 3, 4, 5, 6, 7}, 8},
			ans200{[]int{7, 1, 2, 3, 4, 5, 6}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 200------------------------\n")

	for _, q := range qs {
		a, p := q.ans200, q.para200
		fmt.Printf("【input】:%v       【answer】:%v      【output】:%v\n", p, a, rotate(p.nums, p.target))
	}
	fmt.Printf("\n\n\n")
}
