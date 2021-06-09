package leetcode

import (
	"fmt"
	"testing"
)

type question88 struct {
	para88
	ans88
}

// para 是参数
// one 代表第一个参数
type para88 struct {
	one []int
	l1  int
	two []int
	l2  int
}

// ans 是答案
// one 代表第一个答案
type ans88 struct {
	ans []int
}

func TestProblem88(t *testing.T) {

	qs := []question88{
		{
			para88{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3},
			ans88{[]int{1, 2, 2, 3, 5, 6}},
		},

		{
			para88{[]int{1}, 1, []int{}, 0},
			ans88{[]int{1}},
		},

		{
			para88{[]int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3},
			ans88{[]int{1, 2, 3, 4, 5, 6}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 88------------------------\n")

	for _, q := range qs {
		a, p := q.ans88, q.para88
		fmt.Printf("【input】:%v       【answer】:%v      【output】:%v\n", p, a, merge(p.one, p.l1, p.two, p.l2))
	}
	fmt.Printf("\n\n\n")
}
