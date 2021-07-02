package leetcode

import (
	"fmt"
	"testing"
)

type question17 struct {
	para17
	ans17
}

// para 是参数
// one 代表第一个参数
type para17 struct {
	one int
	two int
}

// ans 是答案
// one 代表第一个答案
type ans17 struct {
	ans [][]int
}

func TestProblem17(t *testing.T) {

	qs := []question17{
		{
			para17{4, 2},
			ans17{[][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 17------------------------\n")

	for _, q := range qs {
		a, p := q.ans17, q.para17
		fmt.Printf("【input】:%v       【answer】:%v      【output】:%v\n", p, a, combine(p.one, p.two))
	}
	fmt.Printf("\n\n\n")
}
