package leetcode

import (
	"fmt"
	"testing"
)

type question77 struct {
	para77
	ans77
}

// para 是参数
// one 代表第一个参数
type para77 struct {
	one int
	two int
}

// ans 是答案
// one 代表第一个答案
type ans77 struct {
	ans [][]int
}

func TestProblem77(t *testing.T) {

	qs := []question77{
		{
			para77{4, 2},
			ans77{[][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 77------------------------\n")

	for _, q := range qs {
		a, p := q.ans77, q.para77
		fmt.Printf("【input】:%v       【answer】:%v      【output】:%v\n", p, a, combine(p.one, p.two))
	}
	fmt.Printf("\n\n\n")
}
