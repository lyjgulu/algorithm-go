package leetcode

import (
	"fmt"
	"testing"
)

type question46 struct {
	para46
	ans46
}

// para 是参数
// one 代表第一个参数
type para46 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans46 struct {
	ans [][]int
}

func TestProblem46(t *testing.T) {

	qs := []question46{
		{
			para46{[]int{1, 2, 3}},
			ans46{[][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
		},

		{
			para46{[]int{0, 1}},
			ans46{[][]int{{0, 1}, {1, 0}}},
		},

		{
			para46{[]int{1}},
			ans46{[][]int{{1}}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 46------------------------\n")

	for _, q := range qs {
		a, p := q.ans46, q.para46
		fmt.Printf("【input】:%v       【answer】:%v      【output】:%v\n", p, a, permute(p.one))
	}
	fmt.Printf("\n\n\n")
}
